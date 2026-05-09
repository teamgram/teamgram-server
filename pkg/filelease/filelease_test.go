package filelease

import (
	"errors"
	"testing"
	"time"
)

func TestSignVerifyRoundTrip(t *testing.T) {
	now := time.Unix(1700000000, 0)
	claims := Claims{
		ObjectID:    "123",
		Bucket:      "documents",
		Key:         "objects/123.dat",
		Size:        6,
		MimeType:    "image/png",
		StorageType: 0x0a4f63c0,
		DCID:        2,
		ExpiresAt:   now.Add(time.Minute).Unix(),
	}

	token, err := Sign("secret", claims)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}
	got, err := Verify("secret", token, now)
	if err != nil {
		t.Fatalf("Verify() error = %v", err)
	}
	if got != claims {
		t.Fatalf("Verify() claims = %#v, want %#v", got, claims)
	}
}

func TestVerifyRejectsTamperedToken(t *testing.T) {
	now := time.Unix(1700000000, 0)
	token, err := Sign("secret", Claims{ObjectID: "123", Bucket: "documents", Key: "objects/123.dat", ExpiresAt: now.Add(time.Minute).Unix()})
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}
	token[len(token)-1] ^= 0xff

	if _, err := Verify("secret", token, now); !errors.Is(err, ErrInvalidToken) {
		t.Fatalf("Verify() error = %v, want ErrInvalidToken", err)
	}
}

func TestVerifyRejectsExpiredClaims(t *testing.T) {
	now := time.Unix(1700000000, 0)
	token, err := Sign("secret", Claims{ObjectID: "123", Bucket: "documents", Key: "objects/123.dat", ExpiresAt: now.Add(-time.Second).Unix()})
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	if _, err := Verify("secret", token, now); !errors.Is(err, ErrExpired) {
		t.Fatalf("Verify() error = %v, want ErrExpired", err)
	}
}

func TestSignVerifyRejectEmptySecret(t *testing.T) {
	claims := Claims{ObjectID: "123", Bucket: "documents", Key: "objects/123.dat", ExpiresAt: time.Now().Add(time.Minute).Unix()}
	if _, err := Sign("", claims); !errors.Is(err, ErrEmptySecret) {
		t.Fatalf("Sign() error = %v, want ErrEmptySecret", err)
	}
	if _, err := Verify("", []byte("token"), time.Now()); !errors.Is(err, ErrEmptySecret) {
		t.Fatalf("Verify() error = %v, want ErrEmptySecret", err)
	}
}
