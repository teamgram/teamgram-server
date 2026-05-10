package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
)

func TestFileReferenceGenerateValidate(t *testing.T) {
	now := time.Unix(1700000000, 0)
	svc := NewFileReferenceService([]byte("test-secret"), func() time.Time { return now })
	claims := FileReferenceClaims{
		MediaID:      10,
		ObjectID:     "obj-1",
		OriginDomain: "photo",
		OriginID:     20,
		ExpireAt:     now.Add(time.Hour).Unix(),
		AccessHash:   30,
	}

	token, err := svc.Generate(claims)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	got, err := svc.Validate(token)
	if err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if got != claims {
		t.Fatalf("Validate() claims = %#v, want %#v", got, claims)
	}
}

func TestFileReferenceValidateEmpty(t *testing.T) {
	svc := NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) })

	_, err := svc.Validate(nil)
	if !errors.Is(err, media.ErrFileReferenceEmpty) {
		t.Fatalf("Validate(nil) error = %v, want ErrFileReferenceEmpty", err)
	}
}

func TestFileReferenceRejectsEmptySecret(t *testing.T) {
	now := time.Unix(1700000000, 0)
	svc := NewFileReferenceService(nil, func() time.Time { return now })

	_, err := svc.Generate(FileReferenceClaims{
		MediaID:    10,
		ObjectID:   "obj-1",
		ExpireAt:   now.Add(time.Hour).Unix(),
		AccessHash: 30,
	})
	if !errors.Is(err, media.ErrFileReferenceInvalid) {
		t.Fatalf("Generate(empty secret) error = %v, want ErrFileReferenceInvalid", err)
	}

	_, err = svc.Validate([]byte("payload.signature"))
	if !errors.Is(err, media.ErrFileReferenceInvalid) {
		t.Fatalf("Validate(empty secret) error = %v, want ErrFileReferenceInvalid", err)
	}
}

func TestFileReferenceValidateExpired(t *testing.T) {
	now := time.Unix(1700000000, 0)
	svc := NewFileReferenceService([]byte("test-secret"), func() time.Time { return now })
	token, err := svc.Generate(FileReferenceClaims{
		MediaID:    10,
		ObjectID:   "obj-1",
		ExpireAt:   now.Add(-time.Second).Unix(),
		AccessHash: 30,
	})
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	_, err = svc.Validate(token)
	if !errors.Is(err, media.ErrFileReferenceExpired) {
		t.Fatalf("Validate(expired) error = %v, want ErrFileReferenceExpired", err)
	}
}

func TestFileReferenceValidateTampered(t *testing.T) {
	now := time.Unix(1700000000, 0)
	svc := NewFileReferenceService([]byte("test-secret"), func() time.Time { return now })
	token, err := svc.Generate(FileReferenceClaims{
		MediaID:    10,
		ObjectID:   "obj-1",
		ExpireAt:   now.Add(time.Hour).Unix(),
		AccessHash: 30,
	})
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	token[len(token)-1] ^= 1

	_, err = svc.Validate(token)
	if !errors.Is(err, media.ErrFileReferenceInvalid) {
		t.Fatalf("Validate(tampered) error = %v, want ErrFileReferenceInvalid", err)
	}
}
