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
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("minio store unavailable")
	}
	if offset < 0 || limit < 0 {
		return nil, ErrInvalidRange
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
