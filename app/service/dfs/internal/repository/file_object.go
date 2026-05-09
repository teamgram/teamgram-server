package repository

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/objectstore"
	"github.com/teamgram/teamgram-server/v2/pkg/filelease"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	phase1FileObjectChunkSize  = 128 * 1024
	defaultReadLeaseTTLSeconds = int64(3600)
)

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
		return r.finalizedFromManifest(manifest)
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
	reader, err := r.OpenUploadFileReader(ctx, info)
	if err != nil {
		return nil, err
	}
	// Task 5b finalizes through the existing upload-state reader bridge. Real
	// segment manifests, streaming finalize, startup scan, TTL cleanup, drain
	// state, and multipart recovery remain Task 5d scope.
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
