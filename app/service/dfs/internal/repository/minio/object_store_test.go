package minio

import (
	"context"
	"errors"
	"net/http"
	"testing"

	miniogo "github.com/minio/minio-go/v7"
)

func TestAccessHashRoundTripStorageFileType(t *testing.T) {
	storageType := int32(0x007efe0e)
	accessHash := MakeAccessHash(storageType, 0x01020304)
	if got := StorageTypeFromAccessHash(accessHash); got != storageType {
		t.Fatalf("storage type = %#x, want %#x", got, storageType)
	}
}

func TestObjectPathUsesDatSuffix(t *testing.T) {
	if got := ObjectPath(12345); got != "12345.dat" {
		t.Fatalf("ObjectPath() = %q, want %q", got, "12345.dat")
	}
}

func TestGetObjectRangeValidatesRangeBeforeClientUse(t *testing.T) {
	_, err := (&Store{}).GetObjectRange(context.Background(), "documents", "objects/1.dat", -1, 0)
	if !errors.Is(err, ErrInvalidRange) {
		t.Fatalf("GetObjectRange() error = %v, want ErrInvalidRange", err)
	}
}

func TestNormalizeObjectReadErrorMapsObjectMisses(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "no such key",
			err:  miniogo.ErrorResponse{Code: "NoSuchKey"},
		},
		{
			name: "no such object",
			err:  miniogo.ErrorResponse{Code: "NoSuchObject"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := normalizeObjectReadError("minio get object photos/missing.dat", tt.err)
			if !errors.Is(err, ErrObjectNotFound) {
				t.Fatalf("normalizeObjectReadError() error = %v, want ErrObjectNotFound", err)
			}
		})
	}
}

func TestNormalizeObjectReadErrorDoesNotMapStorageErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "no such bucket",
			err:  miniogo.ErrorResponse{Code: "NoSuchBucket"},
		},
		{
			name: "http not found without object miss code",
			err:  miniogo.ErrorResponse{StatusCode: http.StatusNotFound},
		},
		{
			name: "ordinary error",
			err:  errors.New("connection reset"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := normalizeObjectReadError("minio read object photos/file.dat", tt.err)
			if errors.Is(err, ErrObjectNotFound) {
				t.Fatalf("normalizeObjectReadError() error = %v, want ordinary storage error", err)
			}
			if !errors.Is(err, tt.err) {
				t.Fatalf("normalizeObjectReadError() error = %v, want original cause preserved", err)
			}
		})
	}
}

func TestNormalizeAbortMultipartUploadErrorIgnoresMissingUpload(t *testing.T) {
	err := normalizeAbortMultipartUploadError("minio abort multipart upload documents/objects/file.dat upload_id=missing", miniogo.ErrorResponse{Code: miniogo.NoSuchUpload})
	if err != nil {
		t.Fatalf("normalizeAbortMultipartUploadError() error = %v, want nil", err)
	}
}

func TestNormalizeAbortMultipartUploadErrorPreservesStorageErrors(t *testing.T) {
	cause := errors.New("connection reset")
	err := normalizeAbortMultipartUploadError("minio abort multipart upload documents/objects/file.dat upload_id=up-1", cause)
	if !errors.Is(err, cause) {
		t.Fatalf("normalizeAbortMultipartUploadError() error = %v, want original cause preserved", err)
	}
}
