package repository

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/objectstore"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/spool"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/xkv"
	"github.com/teamgram/teamgram-server/v2/pkg/filelease"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	phase1FileObjectChunkSize  = 128 * 1024
	defaultReadLeaseTTLSeconds = int64(3600)
	pendingMultipartUploadID   = "__pending_multipart_upload__"
)

type segmentedUploadStateModel interface {
	SegmentCount(info *xkv.DfsFileInfo) (int64, error)
	BuildSegment(ctx context.Context, info *xkv.DfsFileInfo, segmentNo int64, uploadID string, partNumber int) (spool.BuiltSegment, error)
	RebuildSegment(ctx context.Context, info *xkv.DfsFileInfo, segmentNo int64, uploadID string, partNumber int) (spool.BuiltSegment, error)
	ReplaySegment(ctx context.Context, info *xkv.DfsFileInfo, segmentNo int64, dst io.Writer) (spool.ReplayedSegment, error)
	WriteSegmentState(ctx context.Context, info *xkv.DfsFileInfo, state spool.SegmentState) error
	LoadSegmentStates(ctx context.Context, info *xkv.DfsFileInfo) ([]spool.SegmentState, error)
	DeleteSegmentBytes(ctx context.Context, info *xkv.DfsFileInfo, segmentNo int64) error
}

func (r *Repository) CommitUpload(ctx context.Context, uploadSessionID string, ownerID int64, file tg.InputFileClazz, purpose string) (*dfs.FileFinalizedObject, error) {
	if uploadSessionID == "" || ownerID == 0 || file == nil || purpose == "" {
		return nil, dfs.ErrDfsInvalidArgument
	}
	input, err := parseCommitInputFile(file)
	if err != nil {
		return nil, err
	}
	if manifest, err := r.loadUploadManifest(ctx, uploadSessionID); err == nil {
		if !manifestMatchesCommit(manifest, ownerID, input.fileID, purpose) {
			return nil, dfs.ErrDfsInvalidArgument
		}
		if err := r.ensureUploadManifestSidecars(ctx, uploadSessionID, manifest); err != nil {
			return nil, err
		}
		return r.finalizedFromManifest(manifest)
	} else if !errors.Is(err, dfs.ErrDfsFileNotFound) {
		return nil, err
	}
	if manifest, err := r.loadObjectManifest(ctx, commitUploadObjectID(uploadSessionID)); err == nil {
		if !manifestMatchesCommit(manifest, ownerID, input.fileID, purpose) {
			return nil, dfs.ErrDfsInvalidArgument
		}
		if err := r.storeUploadManifest(ctx, uploadSessionID, manifest); err != nil {
			return nil, err
		}
		if len(manifest.Chunks) > 0 {
			if err := r.storeHashManifest(ctx, manifest.ObjectID, manifest.Chunks); err != nil {
				return nil, err
			}
		}
		return r.finalizedFromManifest(manifest)
	} else if !errors.Is(err, dfs.ErrDfsFileNotFound) {
		return nil, err
	}
	if manifest, err := r.loadFinalizingUploadManifest(ctx, uploadSessionID); err == nil {
		if !manifestMatchesCommit(manifest, ownerID, input.fileID, purpose) {
			return nil, dfs.ErrDfsInvalidArgument
		}
		verified, err := r.verifyFinalizedObject(ctx, manifest)
		if err == nil {
			if err := r.storeFinalizedUploadManifests(ctx, uploadSessionID, verified); err != nil {
				return nil, err
			}
			return r.finalizedFromManifest(verified)
		}
		if !errors.Is(err, dfs.ErrDfsFileNotFound) {
			return nil, err
		}
	} else if !errors.Is(err, dfs.ErrDfsFileNotFound) {
		return nil, err
	}

	info, err := r.LoadUploadFileInfo(ctx, ownerID, input.fileID)
	if err != nil {
		return nil, err
	}
	if info == nil || info.Creator != ownerID || info.FileID != input.fileID {
		return nil, dfs.ErrDfsInvalidArgument
	}
	if input.parts > 0 && info.FileTotalParts > 0 && int(input.parts) != info.FileTotalParts {
		return nil, dfs.ErrDfsInvalidArgument
	}
	if segmented, ok := r.uploadStateModel.(segmentedUploadStateModel); ok {
		return r.commitUploadFromSegments(ctx, segmented, uploadSessionID, ownerID, input, purpose, info)
	}
	reader, err := r.OpenUploadFileReader(ctx, info)
	if err != nil {
		return nil, err
	}
	// Legacy non-local upload state fallback. Local spool implements the
	// segmented streaming path below and must not use this compatibility bridge.
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, dfs.WrapDfsStorage("read upload file", err)
	}
	if len(data) == 0 {
		return nil, dfs.ErrDfsInvalidArgument
	}
	if !input.big && input.md5Checksum != "" && fmt.Sprintf("%x", md5.Sum(data)) != strings.ToLower(input.md5Checksum) {
		return nil, dfs.ErrDfsChecksumInvalid
	}
	fileName := input.name
	if fileName == "" {
		fileName = info.FileName
	}
	mimeType := info.MimeType
	if mimeType == "" {
		mimeType = inferMimeType(fileName, data)
	}
	return r.finalizeFileObject(ctx, uploadSessionID, ownerID, input.fileID, purpose, fileName, mimeType, data)
}

func (r *Repository) commitUploadFromSegments(ctx context.Context, segmented segmentedUploadStateModel, uploadSessionID string, ownerID int64, input commitInputFile, purpose string, info *DfsFileInfo) (*dfs.FileFinalizedObject, error) {
	if r == nil || r.objectStore == nil {
		return nil, dfs.WrapDfsStorage("commit upload segments", errors.New("object store unavailable"))
	}
	if r.documentsBucket == "" {
		return nil, dfs.WrapDfsStorage("commit upload segments", errors.New("documents bucket unavailable"))
	}
	xinfo := toXKVDfsFileInfo(info)
	segmentCount, err := segmented.SegmentCount(xinfo)
	if err != nil {
		return nil, mapUploadStateError("count upload segments", err)
	}
	if segmentCount <= 0 {
		return nil, dfs.ErrDfsInvalidArgument
	}
	objectID := commitUploadObjectID(uploadSessionID)
	key := "objects/" + objectID + ".dat"
	states, err := segmented.LoadSegmentStates(ctx, xinfo)
	if err != nil && !errors.Is(err, spool.ErrUploadStateNotFound) {
		return nil, mapUploadStateError("load upload segment states", err)
	}
	uploadID, err := reusableMultipartUploadID(states, key)
	if err != nil {
		return nil, err
	}
	if segmentCount > 1 && uploadID == "" {
		uploadID = pendingMultipartUploadID
	}
	digest := newFileObjectDigest(phase1FileObjectChunkSize)
	var firstBytes []byte
	var completeParts []minioadapter.MultipartPart
	var pendingUploads []pendingMultipartSegment
	var manifest objectstore.ObjectManifest
	completedMultipart := false

	if segmentCount == 1 {
		segment, err := segmented.BuildSegment(ctx, xinfo, 0, "single-put", 1)
		if err != nil {
			return nil, mapBuildSegmentError(segment, err)
		}
		segment.State.ObjectKey = key
		segment.State.MultipartUploadID = ""
		segment.State.MultipartPartNumber = 0
		if err := segmented.WriteSegmentState(ctx, xinfo, segment.State); err != nil {
			return nil, dfs.WrapDfsStorage("write segment uploading state", err)
		}
		replayed, err := segmented.ReplaySegment(ctx, xinfo, 0, digest)
		if err != nil {
			return nil, mapReplaySegmentError(replayed, err)
		}
		if segment.State.Checksum != replayed.Checksum {
			return nil, checksumChangedDuringSegmentBuildError(segmentNoLabel(0), segment.State.Checksum, replayed.Checksum)
		}
		if err := validateCommitDigest(input, digest); err != nil {
			return nil, err
		}
		size, err := r.putSingleSegment(ctx, key, segment.Path, nil)
		if err != nil {
			return nil, err
		}
		segment.State.Status = spool.SegmentStatusUploaded
		segment.State.ETag = "single-put"
		segment.State.Size = size
		if err := segmented.WriteSegmentState(ctx, xinfo, segment.State); err != nil {
			return nil, dfs.WrapDfsStorage("write segment uploaded state", err)
		}
		if err := segmented.DeleteSegmentBytes(ctx, xinfo, 0); err != nil {
			return nil, dfs.WrapDfsStorage("delete uploaded segment bytes", err)
		}
		firstBytes = replayed.FirstBytes
	} else {
		for segmentNo := int64(0); segmentNo < segmentCount; segmentNo++ {
			partNumber := int(segmentNo) + 1
			segment, err := segmented.BuildSegment(ctx, xinfo, segmentNo, uploadID, partNumber)
			if err != nil {
				return nil, mapBuildSegmentError(segment, err)
			}
			replayed, err := segmented.ReplaySegment(ctx, xinfo, segmentNo, digest)
			if err != nil {
				return nil, mapReplaySegmentError(replayed, err)
			}
			if len(firstBytes) == 0 {
				firstBytes = replayed.FirstBytes
			}
			segment.State.ObjectKey = key
			if segment.AlreadyDone && segment.State.Checksum == replayed.Checksum {
				completeParts = append(completeParts, minioadapter.MultipartPart{
					PartNumber: segment.State.MultipartPartNumber,
					ETag:       segment.State.ETag,
					Size:       segment.State.Size,
					Checksum:   segment.State.Checksum,
				})
				continue
			}
			if segment.AlreadyDone {
				rebuilt, err := segmented.RebuildSegment(ctx, xinfo, segmentNo, uploadID, partNumber)
				if err != nil {
					return nil, mapBuildSegmentError(rebuilt, err)
				}
				rebuilt.State.ObjectKey = key
				if rebuilt.State.Checksum != replayed.Checksum {
					return nil, checksumChangedDuringSegmentBuildError(segmentNoLabel(segmentNo), rebuilt.State.Checksum, replayed.Checksum)
				}
				segment = rebuilt
			} else if segment.State.Checksum != replayed.Checksum {
				return nil, checksumChangedDuringSegmentBuildError(segmentNoLabel(segmentNo), segment.State.Checksum, replayed.Checksum)
			}
			pendingUploads = append(pendingUploads, pendingMultipartSegment{
				state:      segment.State,
				path:       segment.Path,
				partNumber: partNumber,
			})
		}
		if err := validateCommitDigest(input, digest); err != nil {
			return nil, err
		}
		if uploadID == pendingMultipartUploadID {
			uploadID = ""
		}
		if uploadID != "" {
			completedMultipart, err = r.completedMultipartObjectExists(ctx, key)
			if err != nil {
				return nil, err
			}
			if completedMultipart {
				pendingUploads = nil
			}
		}
		beganMultipart := false
		multipartTracked := uploadID != ""
		if !completedMultipart && len(pendingUploads) > 0 && uploadID == "" {
			uploadID, err = r.objectStore.BeginMultipartUpload(ctx, r.documentsBucket, key)
			if err != nil {
				return nil, dfs.WrapDfsStorage("begin multipart upload", err)
			}
			beganMultipart = true
		}
		for _, pending := range pendingUploads {
			pending.state.MultipartUploadID = uploadID
			pending.state.ObjectKey = key
			if err := segmented.WriteSegmentState(ctx, xinfo, pending.state); err != nil {
				if beganMultipart && !multipartTracked {
					if abortErr := r.objectStore.AbortMultipartUpload(ctx, r.documentsBucket, key, uploadID); abortErr != nil {
						return nil, dfs.WrapDfsStorage("write segment uploading state", fmt.Errorf("%w; abort multipart upload: %v", err, abortErr))
					}
				}
				return nil, dfs.WrapDfsStorage("write segment uploading state", err)
			}
			multipartTracked = true
			part, err := r.uploadMultipartSegment(ctx, key, uploadID, pending.partNumber, pending.path, pending.state.Checksum, nil)
			if err != nil {
				return nil, err
			}
			pending.state.Status = spool.SegmentStatusUploaded
			pending.state.ETag = part.ETag
			pending.state.Size = part.Size
			if err := segmented.WriteSegmentState(ctx, xinfo, pending.state); err != nil {
				return nil, dfs.WrapDfsStorage("write segment uploaded state", err)
			}
			if err := segmented.DeleteSegmentBytes(ctx, xinfo, pending.state.SegmentNo); err != nil {
				return nil, dfs.WrapDfsStorage("delete uploaded segment bytes", err)
			}
			completeParts = append(completeParts, part)
		}
		fileName := input.name
		if fileName == "" {
			fileName = info.FileName
		}
		mimeType := info.MimeType
		if mimeType == "" {
			mimeType = inferMimeType(fileName, firstBytes)
		}
		manifest = r.newObjectManifest(objectID, uploadSessionID, ownerID, input.fileID, purpose, fileName, mimeType, key, digest)
		if !completedMultipart {
			if err := r.storeFinalizingUploadManifest(ctx, uploadSessionID, manifest); err != nil {
				return nil, err
			}
			sort.Slice(completeParts, func(i, j int) bool { return completeParts[i].PartNumber < completeParts[j].PartNumber })
			if _, err := r.objectStore.CompleteMultipartUpload(ctx, r.documentsBucket, key, uploadID, completeParts); err != nil {
				return nil, dfs.WrapDfsStorage("complete multipart upload", err)
			}
		}
		verified, err := r.verifyFinalizedObject(ctx, manifest)
		if err != nil {
			return nil, err
		}
		manifest = verified
	}
	if err := validateCommitDigest(input, digest); err != nil {
		return nil, err
	}
	if manifest.ObjectID == "" {
		fileName := input.name
		if fileName == "" {
			fileName = info.FileName
		}
		mimeType := info.MimeType
		if mimeType == "" {
			mimeType = inferMimeType(fileName, firstBytes)
		}
		manifest = r.newObjectManifest(objectID, uploadSessionID, ownerID, input.fileID, purpose, fileName, mimeType, key, digest)
	}
	if err := r.storeFinalizedUploadManifests(ctx, uploadSessionID, manifest); err != nil {
		return nil, err
	}
	return r.finalizedFromManifest(manifest)
}

type pendingMultipartSegment struct {
	state      spool.SegmentState
	path       string
	partNumber int
}

func reusableMultipartUploadID(states []spool.SegmentState, objectKey string) (string, error) {
	uploadID := ""
	for _, state := range states {
		if state.MultipartUploadID == "" {
			continue
		}
		if state.Status != spool.SegmentStatusUploading && state.Status != spool.SegmentStatusUploaded {
			continue
		}
		if state.ObjectKey != "" && state.ObjectKey != objectKey {
			return "", dfs.WrapDfsStorage("reuse multipart upload", fmt.Errorf("conflicting object key for upload_id=%s", state.MultipartUploadID))
		}
		if uploadID != "" && uploadID != state.MultipartUploadID {
			return "", dfs.WrapDfsStorage("reuse multipart upload", fmt.Errorf("conflicting multipart upload ids %q and %q", uploadID, state.MultipartUploadID))
		}
		uploadID = state.MultipartUploadID
	}
	return uploadID, nil
}

func validateCommitDigest(input commitInputFile, digest *fileObjectDigest) error {
	if digest.Size() <= 0 {
		return dfs.ErrDfsInvalidArgument
	}
	if !input.big && input.md5Checksum != "" && digest.MD5Hex() != strings.ToLower(input.md5Checksum) {
		return dfs.ErrDfsChecksumInvalid
	}
	return nil
}

func (r *Repository) newObjectManifest(objectID, uploadSessionID string, ownerID, fileID int64, purpose, fileName, mimeType, key string, digest *fileObjectDigest) objectstore.ObjectManifest {
	return objectstore.ObjectManifest{
		ObjectID:        objectID,
		UploadSessionID: uploadSessionID,
		OwnerID:         ownerID,
		FileID:          fileID,
		Purpose:         purpose,
		FileName:        fileName,
		Bucket:          r.documentsBucket,
		Key:             key,
		Size:            digest.Size(),
		MimeType:        mimeType,
		StorageType:     storageTypeForFile(fileName, mimeType),
		SHA256:          digest.SHA256(),
		DCID:            r.localDCID,
		Chunks:          digest.Chunks(),
	}
}

func (r *Repository) ensureUploadManifestSidecars(ctx context.Context, uploadSessionID string, manifest objectstore.ObjectManifest) error {
	if _, err := r.loadObjectManifest(ctx, manifest.ObjectID); err != nil {
		if !errors.Is(err, dfs.ErrDfsFileNotFound) {
			return err
		}
		if err := r.storeObjectManifest(ctx, manifest); err != nil {
			return err
		}
	}
	if len(manifest.Chunks) == 0 {
		return nil
	}
	if _, err := r.loadHashManifest(ctx, manifest.ObjectID); err != nil {
		if !errors.Is(err, dfs.ErrDfsFileNotFound) {
			return err
		}
		if err := r.storeHashManifest(ctx, manifest.ObjectID, manifest.Chunks); err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) storeFinalizedUploadManifests(ctx context.Context, uploadSessionID string, manifest objectstore.ObjectManifest) error {
	if err := r.storeObjectManifest(ctx, manifest); err != nil {
		return err
	}
	if err := r.storeUploadManifest(ctx, uploadSessionID, manifest); err != nil {
		return err
	}
	if len(manifest.Chunks) > 0 {
		if err := r.storeHashManifest(ctx, manifest.ObjectID, manifest.Chunks); err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) verifyFinalizedObject(ctx context.Context, manifest objectstore.ObjectManifest) (objectstore.ObjectManifest, error) {
	digest, err := r.readObjectDigest(ctx, manifest.Bucket, manifest.Key)
	if err != nil {
		return objectstore.ObjectManifest{}, err
	}
	actualSHA := digest.SHA256()
	if digest.Size() != manifest.Size || !bytes.Equal(actualSHA, manifest.SHA256) {
		return objectstore.ObjectManifest{}, dfs.WrapDfsStorage("verify finalized object", fmt.Errorf("object digest mismatch key=%s size=%d want=%d sha256=%x want=%x", manifest.Key, digest.Size(), manifest.Size, actualSHA, manifest.SHA256))
	}
	verified := manifest
	verified.Size = digest.Size()
	verified.SHA256 = append([]byte(nil), actualSHA...)
	verified.Chunks = digest.Chunks()
	return verified, nil
}

func (r *Repository) readObjectDigest(ctx context.Context, bucket, key string) (*fileObjectDigest, error) {
	if r == nil || r.objectStore == nil {
		return nil, dfs.WrapDfsStorage("read finalized object", errors.New("object store unavailable"))
	}
	reader, err := r.objectStore.GetObjectReader(ctx, bucket, key)
	if err != nil {
		return nil, mapObjectReadError("read finalized object", err)
	}
	digest := newFileObjectDigest(phase1FileObjectChunkSize)
	_, copyErr := io.Copy(digest, reader)
	closeErr := reader.Close()
	if copyErr != nil {
		return nil, mapObjectReadError("read finalized object", copyErr)
	}
	if closeErr != nil {
		return nil, dfs.WrapDfsStorage("close finalized object", closeErr)
	}
	return digest, nil
}

func (r *Repository) completedMultipartObjectExists(ctx context.Context, key string) (bool, error) {
	if r == nil || r.objectStore == nil {
		return false, dfs.WrapDfsStorage("probe completed multipart object", errors.New("object store unavailable"))
	}
	if _, err := r.objectStore.GetObjectRange(ctx, r.documentsBucket, key, 0, 1); err != nil {
		if errors.Is(err, minioadapter.ErrObjectNotFound) {
			return false, nil
		}
		return false, dfs.WrapDfsStorage("probe completed multipart object", err)
	}
	return true, nil
}

func mapBuildSegmentError(segment spool.BuiltSegment, err error) error {
	if errors.Is(err, spool.ErrUploadStateNotFound) && segment.MissingPart {
		return &dfs.MissingUploadPartError{Part: segment.MissingPartNo}
	}
	return mapUploadStateError("build upload segment", err)
}

func mapReplaySegmentError(segment spool.ReplayedSegment, err error) error {
	if errors.Is(err, spool.ErrUploadStateNotFound) && segment.MissingPart {
		return &dfs.MissingUploadPartError{Part: segment.MissingPartNo}
	}
	return mapUploadStateError("replay upload segment", err)
}

func segmentNoLabel(segmentNo int64) string {
	return strconv.FormatInt(segmentNo, 10)
}

func checksumChangedDuringSegmentBuildError(segmentNo, builtChecksum, replayChecksum string) error {
	return dfs.WrapDfsStorage("validate upload segment checksum", fmt.Errorf("segment %s checksum changed between build and replay: built=%s replay=%s", segmentNo, builtChecksum, replayChecksum))
}

func (r *Repository) putSingleSegment(ctx context.Context, key, path string, digest *fileObjectDigest) (int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, dfs.WrapDfsStorage("open upload segment", err)
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return 0, dfs.WrapDfsStorage("stat upload segment", err)
	}
	reader := io.Reader(f)
	if digest != nil {
		reader = io.TeeReader(f, digest)
	}
	size, err := r.objectStore.PutObjectReader(ctx, r.documentsBucket, key, reader)
	if err != nil {
		return 0, dfs.WrapDfsStorage("put file object", err)
	}
	if size <= 0 {
		size = info.Size()
	}
	return size, nil
}

func (r *Repository) uploadMultipartSegment(ctx context.Context, key, uploadID string, partNumber int, path, checksum string, digest *fileObjectDigest) (minioadapter.MultipartPart, error) {
	f, err := os.Open(path)
	if err != nil {
		return minioadapter.MultipartPart{}, dfs.WrapDfsStorage("open upload segment", err)
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return minioadapter.MultipartPart{}, dfs.WrapDfsStorage("stat upload segment", err)
	}
	reader := io.Reader(f)
	if digest != nil {
		reader = io.TeeReader(f, digest)
	}
	part, err := r.objectStore.UploadMultipartPart(ctx, r.documentsBucket, key, uploadID, partNumber, reader, info.Size(), checksum)
	if err != nil {
		return minioadapter.MultipartPart{}, dfs.WrapDfsStorage("upload multipart part", err)
	}
	return part, nil
}

type fileObjectDigest struct {
	sha256      hash.Hash
	md5         hash.Hash
	chunkSize   int
	chunkOffset int64
	chunkBuf    []byte
	chunks      []objectstore.HashChunk
	size        int64
}

func newFileObjectDigest(chunkSize int) *fileObjectDigest {
	return &fileObjectDigest{
		sha256:    sha256.New(),
		md5:       md5.New(),
		chunkSize: chunkSize,
	}
}

func (d *fileObjectDigest) Write(p []byte) (int, error) {
	if d == nil {
		return len(p), nil
	}
	_, _ = d.sha256.Write(p)
	_, _ = d.md5.Write(p)
	d.size += int64(len(p))
	remaining := p
	for len(remaining) > 0 && d.chunkSize > 0 {
		need := d.chunkSize - len(d.chunkBuf)
		if need > len(remaining) {
			need = len(remaining)
		}
		d.chunkBuf = append(d.chunkBuf, remaining[:need]...)
		remaining = remaining[need:]
		if len(d.chunkBuf) == d.chunkSize {
			d.flushChunk()
		}
	}
	return len(p), nil
}

func (d *fileObjectDigest) flushChunk() {
	sum := sha256.Sum256(d.chunkBuf)
	d.chunks = append(d.chunks, objectstore.HashChunk{
		Offset: d.chunkOffset,
		Limit:  int32(len(d.chunkBuf)),
		Hash:   append([]byte(nil), sum[:]...),
	})
	d.chunkOffset += int64(len(d.chunkBuf))
	d.chunkBuf = d.chunkBuf[:0]
}

func (d *fileObjectDigest) Size() int64 {
	if d == nil {
		return 0
	}
	return d.size
}

func (d *fileObjectDigest) SHA256() []byte {
	if d == nil {
		return nil
	}
	return d.sha256.Sum(nil)
}

func (d *fileObjectDigest) MD5Hex() string {
	if d == nil {
		return ""
	}
	return fmt.Sprintf("%x", d.md5.Sum(nil))
}

func (d *fileObjectDigest) Chunks() []objectstore.HashChunk {
	if d == nil {
		return nil
	}
	if len(d.chunkBuf) > 0 {
		d.flushChunk()
	}
	return append([]objectstore.HashChunk(nil), d.chunks...)
}

func (r *Repository) PutInternalFile(ctx context.Context, ownerID int64, purpose, fileName, mimeType string, data []byte) (*dfs.FileFinalizedObject, error) {
	if ownerID == 0 || purpose == "" || fileName == "" || len(data) == 0 {
		return nil, dfs.ErrDfsInvalidArgument
	}
	if mimeType == "" {
		mimeType = inferMimeType(fileName, data)
	}
	objectID, err := r.NextDocumentID(ctx)
	if err != nil {
		return nil, err
	}
	objectIDString := strconv.FormatInt(objectID, 10)
	return r.finalizeKnownFileObject(ctx, objectIDString, "internal:"+objectIDString, ownerID, objectID, purpose, fileName, mimeType, data)
}

func (r *Repository) ReadByLease(ctx context.Context, readLease []byte, offset int64, limit int32) ([]byte, int32, error) {
	if len(readLease) == 0 || offset < 0 || limit < 0 {
		return nil, 0, dfs.ErrDfsInvalidArgument
	}
	claims, err := r.verifyReadLease(readLease)
	if err != nil {
		return nil, 0, err
	}
	if claims.ObjectID == "" || claims.Bucket == "" || claims.Key == "" {
		return nil, 0, dfs.ErrDfsInvalidArgument
	}
	data, err := r.getObjectRange(ctx, claims.Bucket, claims.Key, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return data, claims.StorageType, nil
}

func (r *Repository) HashesByLease(ctx context.Context, readLease []byte, offset int64, limit int32) ([]objectstore.HashChunk, error) {
	if len(readLease) == 0 || offset < 0 || limit < 0 {
		return nil, dfs.ErrDfsInvalidArgument
	}
	claims, err := r.verifyReadLease(readLease)
	if err != nil {
		return nil, err
	}
	manifest, err := r.loadObjectManifest(ctx, claims.ObjectID)
	if err != nil {
		return nil, err
	}
	chunks := manifest.Chunks
	if len(chunks) == 0 {
		chunks, err = r.loadHashManifest(ctx, claims.ObjectID)
		if err != nil {
			return nil, err
		}
	}
	return objectstore.FilterHashChunks(chunks, offset, limit), nil
}

func (r *Repository) finalizeFileObject(ctx context.Context, uploadSessionID string, ownerID, fileID int64, purpose, fileName, mimeType string, data []byte) (*dfs.FileFinalizedObject, error) {
	return r.finalizeKnownFileObject(ctx, commitUploadObjectID(uploadSessionID), uploadSessionID, ownerID, fileID, purpose, fileName, mimeType, data)
}

func (r *Repository) finalizeKnownFileObject(ctx context.Context, objectID, uploadSessionID string, ownerID, fileID int64, purpose, fileName, mimeType string, data []byte) (*dfs.FileFinalizedObject, error) {
	if r == nil || r.objectStore == nil {
		return nil, dfs.WrapDfsStorage("finalize file object", errors.New("object store unavailable"))
	}
	if r.documentsBucket == "" {
		return nil, dfs.WrapDfsStorage("finalize file object", errors.New("documents bucket unavailable"))
	}
	if objectID == "" || uploadSessionID == "" || ownerID == 0 || fileID == 0 || purpose == "" || len(data) == 0 {
		return nil, dfs.ErrDfsInvalidArgument
	}
	key := "objects/" + objectID + ".dat"
	size, err := r.objectStore.PutObjectBytes(ctx, r.documentsBucket, key, data)
	if err != nil {
		return nil, dfs.WrapDfsStorage("put file object", err)
	}
	sum := sha256.Sum256(data)
	manifest := objectstore.ObjectManifest{
		ObjectID:        objectID,
		UploadSessionID: uploadSessionID,
		OwnerID:         ownerID,
		FileID:          fileID,
		Purpose:         purpose,
		FileName:        fileName,
		Bucket:          r.documentsBucket,
		Key:             key,
		Size:            size,
		MimeType:        mimeType,
		StorageType:     storageTypeForFile(fileName, mimeType),
		SHA256:          append([]byte(nil), sum[:]...),
		DCID:            r.localDCID,
		Chunks:          objectstore.BuildHashChunks(data, phase1FileObjectChunkSize),
	}
	if err := r.storeObjectManifest(ctx, manifest); err != nil {
		return nil, err
	}
	if err := r.storeUploadManifest(ctx, uploadSessionID, manifest); err != nil {
		return nil, err
	}
	if err := r.storeHashManifest(ctx, objectID, manifest.Chunks); err != nil {
		return nil, err
	}
	return r.finalizedFromManifest(manifest)
}

func manifestMatchesCommit(manifest objectstore.ObjectManifest, ownerID, fileID int64, purpose string) bool {
	return manifest.OwnerID == ownerID && manifest.FileID == fileID && manifest.Purpose == purpose
}

func (r *Repository) finalizedFromManifest(manifest objectstore.ObjectManifest) (*dfs.FileFinalizedObject, error) {
	lease, err := r.signReadLease(manifest)
	if err != nil {
		return nil, err
	}
	return dfs.MakeTLFileFinalizedObject(&dfs.TLFileFinalizedObject{
		ObjectId:        manifest.ObjectID,
		UploadSessionId: manifest.UploadSessionID,
		Bucket:          manifest.Bucket,
		Key:             manifest.Key,
		Size2:           manifest.Size,
		MimeType:        manifest.MimeType,
		Sha256:          append([]byte(nil), manifest.SHA256...),
		ReadLease:       lease,
		DcId:            manifest.DCID,
	}), nil
}

func (r *Repository) signReadLease(manifest objectstore.ObjectManifest) ([]byte, error) {
	ttl := r.readLeaseTTLSeconds
	if ttl <= 0 {
		ttl = defaultReadLeaseTTLSeconds
	}
	token, err := filelease.Sign(r.readLeaseSecret, filelease.Claims{
		ObjectID:    manifest.ObjectID,
		Bucket:      manifest.Bucket,
		Key:         manifest.Key,
		Size:        manifest.Size,
		MimeType:    manifest.MimeType,
		StorageType: manifest.StorageType,
		DCID:        manifest.DCID,
		ExpiresAt:   time.Now().Add(time.Duration(ttl) * time.Second).Unix(),
	})
	if err != nil {
		return nil, mapReadLeaseError("sign read lease", err)
	}
	return token, nil
}

func (r *Repository) verifyReadLease(token []byte) (filelease.Claims, error) {
	claims, err := filelease.Verify(r.readLeaseSecret, token, time.Now())
	if err != nil {
		return filelease.Claims{}, mapReadLeaseError("verify read lease", err)
	}
	return claims, nil
}

func mapReadLeaseError(op string, err error) error {
	if errors.Is(err, filelease.ErrEmptySecret) {
		return dfs.WrapDfsStorage(op, err)
	}
	if errors.Is(err, filelease.ErrInvalidToken) || errors.Is(err, filelease.ErrExpired) {
		return fmt.Errorf("%w: %s: %w", dfs.ErrDfsInvalidArgument, op, err)
	}
	return dfs.WrapDfsStorage(op, err)
}

func (r *Repository) storeObjectManifest(ctx context.Context, manifest objectstore.ObjectManifest) error {
	key, err := r.manifestKeys.Object(manifest.ObjectID)
	if err != nil {
		return mapManifestKeyError(err)
	}
	return r.putJSON(ctx, "put object manifest", key, manifest)
}

func (r *Repository) storeUploadManifest(ctx context.Context, uploadSessionID string, manifest objectstore.ObjectManifest) error {
	key, err := r.manifestKeys.Upload(uploadSessionID)
	if err != nil {
		return mapManifestKeyError(err)
	}
	return r.putJSON(ctx, "put upload manifest", key, manifest)
}

func (r *Repository) storeFinalizingUploadManifest(ctx context.Context, uploadSessionID string, manifest objectstore.ObjectManifest) error {
	key, err := r.manifestKeys.UploadFinalizing(uploadSessionID)
	if err != nil {
		return mapManifestKeyError(err)
	}
	return r.putJSON(ctx, "put finalizing upload manifest", key, manifest)
}

func (r *Repository) storeHashManifest(ctx context.Context, objectID string, chunks []objectstore.HashChunk) error {
	key, err := r.manifestKeys.Hashes(objectID)
	if err != nil {
		return mapManifestKeyError(err)
	}
	return r.putJSON(ctx, "put hash manifest", key, chunks)
}

func (r *Repository) loadObjectManifest(ctx context.Context, objectID string) (objectstore.ObjectManifest, error) {
	key, err := r.manifestKeys.Object(objectID)
	if err != nil {
		return objectstore.ObjectManifest{}, mapManifestKeyError(err)
	}
	var manifest objectstore.ObjectManifest
	if err := r.getJSON(ctx, "get object manifest", key, &manifest); err != nil {
		return objectstore.ObjectManifest{}, err
	}
	return manifest, nil
}

func (r *Repository) loadUploadManifest(ctx context.Context, uploadSessionID string) (objectstore.ObjectManifest, error) {
	key, err := r.manifestKeys.Upload(uploadSessionID)
	if err != nil {
		return objectstore.ObjectManifest{}, mapManifestKeyError(err)
	}
	var manifest objectstore.ObjectManifest
	if err := r.getJSON(ctx, "get upload manifest", key, &manifest); err != nil {
		return objectstore.ObjectManifest{}, err
	}
	return manifest, nil
}

func (r *Repository) loadFinalizingUploadManifest(ctx context.Context, uploadSessionID string) (objectstore.ObjectManifest, error) {
	key, err := r.manifestKeys.UploadFinalizing(uploadSessionID)
	if err != nil {
		return objectstore.ObjectManifest{}, mapManifestKeyError(err)
	}
	var manifest objectstore.ObjectManifest
	if err := r.getJSON(ctx, "get finalizing upload manifest", key, &manifest); err != nil {
		return objectstore.ObjectManifest{}, err
	}
	return manifest, nil
}

func (r *Repository) loadHashManifest(ctx context.Context, objectID string) ([]objectstore.HashChunk, error) {
	key, err := r.manifestKeys.Hashes(objectID)
	if err != nil {
		return nil, mapManifestKeyError(err)
	}
	var chunks []objectstore.HashChunk
	if err := r.getJSON(ctx, "get hash manifest", key, &chunks); err != nil {
		return nil, err
	}
	return chunks, nil
}

func (r *Repository) putJSON(ctx context.Context, op, key string, v any) error {
	if r == nil || r.objectStore == nil {
		return dfs.WrapDfsStorage(op, errors.New("object store unavailable"))
	}
	data, err := json.Marshal(v)
	if err != nil {
		return dfs.WrapDfsStorage(op, err)
	}
	if _, err := r.objectStore.PutObjectBytes(ctx, r.documentsBucket, key, data); err != nil {
		return dfs.WrapDfsStorage(op, err)
	}
	return nil
}

func (r *Repository) getJSON(ctx context.Context, op, key string, out any) error {
	data, err := r.getObjectRange(ctx, r.documentsBucket, key, 0, 0)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, out); err != nil {
		return dfs.WrapDfsStorage(op, err)
	}
	return nil
}

func (r *Repository) getObjectRange(ctx context.Context, bucket, key string, offset int64, limit int32) ([]byte, error) {
	if r == nil || r.objectStore == nil {
		return nil, dfs.WrapDfsStorage("get file object", errors.New("object store unavailable"))
	}
	data, err := r.objectStore.GetObjectRange(ctx, bucket, key, offset, limit)
	if err != nil {
		return nil, mapObjectReadError("get file object", err)
	}
	return data, nil
}

func mapManifestKeyError(err error) error {
	if errors.Is(err, objectstore.ErrInvalidManifestKey) {
		return fmt.Errorf("%w: manifest key: %w", dfs.ErrDfsInvalidArgument, err)
	}
	return dfs.WrapDfsStorage("manifest key", err)
}

type commitInputFile struct {
	fileID      int64
	parts       int32
	name        string
	md5Checksum string
	big         bool
}

func parseCommitInputFile(file tg.InputFileClazz) (commitInputFile, error) {
	switch f := file.(type) {
	case *tg.TLInputFile:
		if f == nil || f.Id == 0 || f.Parts <= 0 || f.Name == "" {
			return commitInputFile{}, dfs.ErrDfsInvalidArgument
		}
		return commitInputFile{fileID: f.Id, parts: f.Parts, name: f.Name, md5Checksum: f.Md5Checksum}, nil
	case *tg.TLInputFileBig:
		if f == nil || f.Id == 0 || f.Parts <= 0 || f.Name == "" {
			return commitInputFile{}, dfs.ErrDfsInvalidArgument
		}
		return commitInputFile{fileID: f.Id, parts: f.Parts, name: f.Name, big: true}, nil
	default:
		return commitInputFile{}, dfs.ErrDfsInvalidArgument
	}
}

func commitUploadObjectID(uploadSessionID string) string {
	sum := sha256.Sum256([]byte(uploadSessionID))
	return fmt.Sprintf("upload-%x", sum)
}

func inferMimeType(fileName string, data []byte) string {
	if ext := strings.ToLower(filepath.Ext(fileName)); ext != "" {
		if mt := mime.TypeByExtension(ext); mt != "" {
			if base, _, err := mime.ParseMediaType(mt); err == nil {
				return base
			}
			return strings.Split(mt, ";")[0]
		}
	}
	if len(data) > 0 {
		return strings.Split(http.DetectContentType(data), ";")[0]
	}
	return "application/octet-stream"
}

func storageTypeForFile(fileName, mimeType string) int32 {
	base := strings.ToLower(strings.TrimSpace(strings.Split(mimeType, ";")[0]))
	switch base {
	case "image/jpeg", "image/jpg":
		return storageTypeID(tg.ClazzID_storage_fileJpeg)
	case "image/png":
		return storageTypeID(tg.ClazzID_storage_filePng)
	case "image/gif":
		return storageTypeID(tg.ClazzID_storage_fileGif)
	case "application/pdf":
		return storageTypeID(tg.ClazzID_storage_filePdf)
	case "audio/mpeg", "audio/mp3":
		return storageTypeID(tg.ClazzID_storage_fileMp3)
	case "video/quicktime", "video/mov":
		return storageTypeID(tg.ClazzID_storage_fileMov)
	case "video/mp4":
		return storageTypeID(tg.ClazzID_storage_fileMp4)
	case "image/webp":
		return storageTypeID(tg.ClazzID_storage_fileWebp)
	}
	switch strings.ToLower(filepath.Ext(fileName)) {
	case ".jpg", ".jpeg":
		return storageTypeID(tg.ClazzID_storage_fileJpeg)
	case ".png":
		return storageTypeID(tg.ClazzID_storage_filePng)
	case ".gif":
		return storageTypeID(tg.ClazzID_storage_fileGif)
	case ".pdf":
		return storageTypeID(tg.ClazzID_storage_filePdf)
	case ".mp3":
		return storageTypeID(tg.ClazzID_storage_fileMp3)
	case ".mov":
		return storageTypeID(tg.ClazzID_storage_fileMov)
	case ".mp4":
		return storageTypeID(tg.ClazzID_storage_fileMp4)
	case ".webp":
		return storageTypeID(tg.ClazzID_storage_fileWebp)
	default:
		return storageTypeID(tg.ClazzID_storage_fileUnknown)
	}
}

func storageTypeID(id uint32) int32 {
	return int32(id)
}
