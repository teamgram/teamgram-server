package minio

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/config"
)

var (
	ErrInvalidRange   = errors.New("invalid minio range")
	ErrObjectNotFound = errors.New("minio object not found")
)

type ObjectStore interface {
	PutObjectBytes(ctx context.Context, bucket, key string, data []byte) (int64, error)
	PutObjectReader(ctx context.Context, bucket, key string, r io.Reader) (int64, error)
	GetObjectReader(ctx context.Context, bucket, key string) (io.ReadCloser, error)
	GetObjectRange(ctx context.Context, bucket, key string, offset int64, limit int32) ([]byte, error)
	BeginMultipartUpload(ctx context.Context, bucket, key string) (string, error)
	UploadMultipartPart(ctx context.Context, bucket, key, uploadID string, partNumber int, r io.Reader, size int64, checksum string) (MultipartPart, error)
	ListMultipartParts(ctx context.Context, bucket, key, uploadID string) ([]MultipartPart, error)
	CompleteMultipartUpload(ctx context.Context, bucket, key, uploadID string, parts []MultipartPart) (int64, error)
	AbortMultipartUpload(ctx context.Context, bucket, key, uploadID string) error
	PutPhotoBytes(ctx context.Context, path string, data []byte) (int64, error)
	PutPhotoReader(ctx context.Context, path string, r io.Reader) (int64, error)
	GetPhotoFile(ctx context.Context, path string, offset int64, limit int32) ([]byte, error)
	PutVideoBytes(ctx context.Context, path string, data []byte) (int64, error)
	GetVideoFile(ctx context.Context, path string, offset int64, limit int32) ([]byte, error)
	PutDocumentReader(ctx context.Context, path string, r io.Reader) (int64, error)
	GetDocumentFile(ctx context.Context, path string, offset int64, limit int32) ([]byte, error)
	PutEncryptedFileReader(ctx context.Context, path string, r io.Reader) (int64, error)
	GetEncryptedFile(ctx context.Context, path string, offset int64, limit int32) ([]byte, error)
}

type MultipartPart struct {
	PartNumber int
	ETag       string
	Size       int64
	Checksum   string
}

type Store struct {
	client    *minio.Client
	photos    string
	videos    string
	documents string
	encrypted string
}

func NewObjectStore(c config.MinioConf) (*Store, error) {
	cli, err := minio.New(c.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKeyID, c.SecretAccessKey, ""),
		Secure: c.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("minio new client: %w", err)
	}
	return &Store{
		client:    cli,
		photos:    c.PhotosBucket,
		videos:    c.VideosBucket,
		documents: c.DocumentsBucket,
		encrypted: c.EncryptedBucket,
	}, nil
}

func MakeAccessHash(storageType int32, rand32 uint32) int64 {
	return int64(storageType)<<32 | int64(rand32)
}

func StorageTypeFromAccessHash(accessHash int64) int32 {
	return int32(accessHash >> 32)
}

func ObjectPath(id int64) string {
	return fmt.Sprintf("%d.dat", id)
}

func (s *Store) PutObjectBytes(ctx context.Context, bucket, key string, data []byte) (int64, error) {
	return s.PutObjectReader(ctx, bucket, key, bytes.NewReader(data))
}

func (s *Store) PutObjectReader(ctx context.Context, bucket, key string, r io.Reader) (int64, error) {
	return s.put(ctx, bucket, key, r)
}

func (s *Store) GetObjectReader(ctx context.Context, bucket, key string) (io.ReadCloser, error) {
	return s.getReader(ctx, bucket, key)
}

func (s *Store) GetObjectRange(ctx context.Context, bucket, key string, offset int64, limit int32) ([]byte, error) {
	return s.get(ctx, bucket, key, offset, limit)
}

func (s *Store) BeginMultipartUpload(ctx context.Context, bucket, key string) (string, error) {
	if s == nil || s.client == nil {
		return "", fmt.Errorf("minio store unavailable")
	}
	core := minio.Core{Client: s.client}
	uploadID, err := core.NewMultipartUpload(ctx, bucket, key, minio.PutObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("minio begin multipart upload %s/%s: %w", bucket, key, err)
	}
	return uploadID, nil
}

func (s *Store) UploadMultipartPart(ctx context.Context, bucket, key, uploadID string, partNumber int, r io.Reader, size int64, checksum string) (MultipartPart, error) {
	if s == nil || s.client == nil {
		return MultipartPart{}, fmt.Errorf("minio store unavailable")
	}
	core := minio.Core{Client: s.client}
	part, err := core.PutObjectPart(ctx, bucket, key, uploadID, partNumber, r, size, minio.PutObjectPartOptions{})
	if err != nil {
		return MultipartPart{}, fmt.Errorf("minio upload multipart part %s/%s upload_id=%s part=%d: %w", bucket, key, uploadID, partNumber, err)
	}
	return MultipartPart{PartNumber: part.PartNumber, ETag: part.ETag, Size: part.Size, Checksum: checksum}, nil
}

func (s *Store) ListMultipartParts(ctx context.Context, bucket, key, uploadID string) ([]MultipartPart, error) {
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("minio store unavailable")
	}
	core := minio.Core{Client: s.client}
	var out []MultipartPart
	partNumberMarker := 0
	for {
		result, err := core.ListObjectParts(ctx, bucket, key, uploadID, partNumberMarker, 1000)
		if err != nil {
			return nil, fmt.Errorf("minio list multipart parts %s/%s upload_id=%s: %w", bucket, key, uploadID, err)
		}
		for _, part := range result.ObjectParts {
			out = append(out, MultipartPart{PartNumber: part.PartNumber, ETag: part.ETag, Size: part.Size})
		}
		if !result.IsTruncated {
			return out, nil
		}
		partNumberMarker = result.NextPartNumberMarker
	}
}

func (s *Store) CompleteMultipartUpload(ctx context.Context, bucket, key, uploadID string, parts []MultipartPart) (int64, error) {
	if s == nil || s.client == nil {
		return 0, fmt.Errorf("minio store unavailable")
	}
	completeParts := make([]minio.CompletePart, 0, len(parts))
	var size int64
	for _, part := range parts {
		completeParts = append(completeParts, minio.CompletePart{PartNumber: part.PartNumber, ETag: part.ETag})
		size += part.Size
	}
	core := minio.Core{Client: s.client}
	if _, err := core.CompleteMultipartUpload(ctx, bucket, key, uploadID, completeParts, minio.PutObjectOptions{}); err != nil {
		return 0, fmt.Errorf("minio complete multipart upload %s/%s upload_id=%s: %w", bucket, key, uploadID, err)
	}
	return size, nil
}

func (s *Store) AbortMultipartUpload(ctx context.Context, bucket, key, uploadID string) error {
	if s == nil || s.client == nil {
		return fmt.Errorf("minio store unavailable")
	}
	core := minio.Core{Client: s.client}
	if err := core.AbortMultipartUpload(ctx, bucket, key, uploadID); err != nil {
		return fmt.Errorf("minio abort multipart upload %s/%s upload_id=%s: %w", bucket, key, uploadID, err)
	}
	return nil
}

func (s *Store) PutPhotoBytes(ctx context.Context, path string, data []byte) (int64, error) {
	return s.PutPhotoReader(ctx, path, bytes.NewReader(data))
}

func (s *Store) PutPhotoReader(ctx context.Context, path string, r io.Reader) (int64, error) {
	return s.put(ctx, s.photos, path, r)
}

func (s *Store) GetPhotoFile(ctx context.Context, path string, offset int64, limit int32) ([]byte, error) {
	return s.get(ctx, s.photos, path, offset, limit)
}

func (s *Store) PutVideoBytes(ctx context.Context, path string, data []byte) (int64, error) {
	return s.PutVideoReader(ctx, path, bytes.NewReader(data))
}

func (s *Store) PutVideoReader(ctx context.Context, path string, r io.Reader) (int64, error) {
	return s.put(ctx, s.videos, path, r)
}

func (s *Store) GetVideoFile(ctx context.Context, path string, offset int64, limit int32) ([]byte, error) {
	return s.get(ctx, s.videos, path, offset, limit)
}

func (s *Store) PutDocumentReader(ctx context.Context, path string, r io.Reader) (int64, error) {
	return s.put(ctx, s.documents, path, r)
}

func (s *Store) GetDocumentFile(ctx context.Context, path string, offset int64, limit int32) ([]byte, error) {
	return s.get(ctx, s.documents, path, offset, limit)
}

func (s *Store) PutEncryptedFileReader(ctx context.Context, path string, r io.Reader) (int64, error) {
	return s.put(ctx, s.encrypted, path, r)
}

func (s *Store) GetEncryptedFile(ctx context.Context, path string, offset int64, limit int32) ([]byte, error) {
	return s.get(ctx, s.encrypted, path, offset, limit)
}

func (s *Store) put(ctx context.Context, bucket, path string, r io.Reader) (int64, error) {
	if s == nil || s.client == nil {
		return 0, fmt.Errorf("minio store unavailable")
	}
	info, err := s.client.PutObject(ctx, bucket, path, r, -1, minio.PutObjectOptions{})
	if err != nil {
		return 0, fmt.Errorf("minio put object %s/%s: %w", bucket, path, err)
	}
	return info.Size, nil
}

func (s *Store) get(ctx context.Context, bucket, path string, offset int64, limit int32) ([]byte, error) {
	if offset < 0 || limit < 0 {
		return nil, ErrInvalidRange
	}
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("minio store unavailable")
	}
	opts := minio.GetObjectOptions{}
	if offset > 0 || limit > 0 {
		end := int64(0)
		if limit > 0 {
			end = offset + int64(limit) - 1
		}
		if err := opts.SetRange(offset, end); err != nil {
			return nil, fmt.Errorf("minio range %s/%s: %w", bucket, path, err)
		}
	}
	obj, err := s.client.GetObject(ctx, bucket, path, opts)
	if err != nil {
		return nil, normalizeObjectReadError(fmt.Sprintf("minio get object %s/%s", bucket, path), err)
	}
	defer obj.Close()
	b, err := io.ReadAll(obj)
	if err != nil {
		return nil, normalizeObjectReadError(fmt.Sprintf("minio read object %s/%s", bucket, path), err)
	}
	return b, nil
}

func (s *Store) getReader(ctx context.Context, bucket, path string) (io.ReadCloser, error) {
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("minio store unavailable")
	}
	obj, err := s.client.GetObject(ctx, bucket, path, minio.GetObjectOptions{})
	if err != nil {
		return nil, normalizeObjectReadError(fmt.Sprintf("minio get object %s/%s", bucket, path), err)
	}
	if _, err := obj.Stat(); err != nil {
		_ = obj.Close()
		return nil, normalizeObjectReadError(fmt.Sprintf("minio stat object %s/%s", bucket, path), err)
	}
	return obj, nil
}

func normalizeObjectReadError(op string, err error) error {
	if err == nil {
		return nil
	}
	if isObjectNotFound(err) {
		return fmt.Errorf("%w: %s: %w", ErrObjectNotFound, op, err)
	}
	return fmt.Errorf("%s: %w", op, err)
}

func isObjectNotFound(err error) bool {
	if isObjectNotFoundResponse(minio.ToErrorResponse(err)) {
		return true
	}
	var resp minio.ErrorResponse
	if errors.As(err, &resp) {
		return isObjectNotFoundResponse(resp)
	}
	return false
}

func isObjectNotFoundResponse(resp minio.ErrorResponse) bool {
	switch resp.Code {
	case "NoSuchKey", "NoSuchObject":
		return true
	}
	return false
}
