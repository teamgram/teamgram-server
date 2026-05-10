package filelease

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var (
	ErrEmptySecret  = errors.New("filelease: empty secret")
	ErrInvalidToken = errors.New("filelease: invalid token")
	ErrExpired      = errors.New("filelease: expired")
)

type Claims struct {
	ObjectID    string `json:"object_id"`
	Bucket      string `json:"bucket"`
	Key         string `json:"key"`
	Size        int64  `json:"size"`
	MimeType    string `json:"mime_type"`
	StorageType int32  `json:"storage_type"`
	DCID        int32  `json:"dc_id"`
	ExpiresAt   int64  `json:"expires_at"`
}

func Sign(secret string, claims Claims) ([]byte, error) {
	if secret == "" {
		return nil, ErrEmptySecret
	}
	payload, err := json.Marshal(claims)
	if err != nil {
		return nil, fmt.Errorf("%w: marshal claims: %w", ErrInvalidToken, err)
	}
	sig := signPayload(secret, payload)
	token := make([]byte, 0, base64.RawURLEncoding.EncodedLen(len(payload))+1+base64.RawURLEncoding.EncodedLen(len(sig)))
	token = base64.RawURLEncoding.AppendEncode(token, payload)
	token = append(token, '.')
	token = base64.RawURLEncoding.AppendEncode(token, sig)
	return token, nil
}

func Verify(secret string, token []byte, now time.Time) (Claims, error) {
	if secret == "" {
		return Claims{}, ErrEmptySecret
	}
	payloadPart, sigPart, ok := bytes.Cut(token, []byte("."))
	if !ok || len(payloadPart) == 0 || len(sigPart) == 0 {
		return Claims{}, ErrInvalidToken
	}
	payload, err := base64.RawURLEncoding.DecodeString(string(payloadPart))
	if err != nil {
		return Claims{}, fmt.Errorf("%w: decode payload: %w", ErrInvalidToken, err)
	}
	gotSig, err := base64.RawURLEncoding.DecodeString(string(sigPart))
	if err != nil {
		return Claims{}, fmt.Errorf("%w: decode signature: %w", ErrInvalidToken, err)
	}
	wantSig := signPayload(secret, payload)
	if !hmac.Equal(gotSig, wantSig) {
		return Claims{}, ErrInvalidToken
	}
	var claims Claims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return Claims{}, fmt.Errorf("%w: unmarshal claims: %w", ErrInvalidToken, err)
	}
	if claims.ExpiresAt <= now.Unix() {
		return Claims{}, ErrExpired
	}
	return claims, nil
}

func signPayload(secret string, payload []byte) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write(payload)
	return mac.Sum(nil)
}
