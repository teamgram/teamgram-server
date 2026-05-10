package repository

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
)

const (
	fileReferenceOpaqueVersion = byte(0x03)
	fileReferenceOpaqueLength  = 25
)

type FileReferenceClaims struct {
	MediaID      int64  `json:"media_id"`
	ObjectID     string `json:"object_id"`
	OriginDomain string `json:"origin_domain"`
	OriginID     int64  `json:"origin_id"`
	ExpireAt     int64  `json:"expire_at"`
	AccessHash   int64  `json:"access_hash"`
}

type FileReferenceStore interface {
	SaveFileReference(ctx context.Context, token []byte, claims FileReferenceClaims) error
	LoadFileReference(ctx context.Context, token []byte) (FileReferenceClaims, error)
}

type FileReferenceService struct {
	legacySecret []byte
	now          func() time.Time
}

func NewFileReferenceService(legacySecret []byte, now func() time.Time) *FileReferenceService {
	if now == nil {
		now = time.Now
	}
	return &FileReferenceService{
		legacySecret: append([]byte(nil), legacySecret...),
		now:          now,
	}
}

func (s *FileReferenceService) Generate(ctx context.Context, claims FileReferenceClaims, store FileReferenceStore) ([]byte, error) {
	if s == nil || store == nil {
		return nil, media.ErrFileReferenceInvalid
	}
	token := make([]byte, fileReferenceOpaqueLength)
	token[0] = fileReferenceOpaqueVersion
	if _, err := rand.Read(token[1:]); err != nil {
		return nil, fmt.Errorf("%w: random handle: %v", media.ErrFileReferenceInvalid, err)
	}
	if err := store.SaveFileReference(ctx, token, claims); err != nil {
		return nil, err
	}
	return token, nil
}

func (s *FileReferenceService) Validate(ctx context.Context, token []byte, store FileReferenceStore) (FileReferenceClaims, error) {
	if s == nil || store == nil {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	if len(token) == 0 {
		return FileReferenceClaims{}, media.ErrFileReferenceEmpty
	}
	if len(token) != fileReferenceOpaqueLength || token[0] != fileReferenceOpaqueVersion {
		return s.validateLegacy(token)
	}
	claims, err := store.LoadFileReference(ctx, token)
	if err != nil {
		return FileReferenceClaims{}, err
	}
	if claims.ExpireAt <= s.now().Unix() {
		return FileReferenceClaims{}, media.ErrFileReferenceExpired
	}
	return claims, nil
}

func (s *FileReferenceService) validateLegacy(token []byte) (FileReferenceClaims, error) {
	if len(s.legacySecret) == 0 {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	claims, err := validateLegacyHMACFileReference(token, s.legacySecret)
	if err != nil {
		return FileReferenceClaims{}, err
	}
	if claims.ExpireAt <= s.now().Unix() {
		return FileReferenceClaims{}, media.ErrFileReferenceExpired
	}
	return claims, nil
}

func validateLegacyHMACFileReference(token []byte, secret []byte) (FileReferenceClaims, error) {
	parts := bytes.Split(token, []byte{'.'})
	if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	payload, err := base64.RawURLEncoding.DecodeString(string(parts[0]))
	if err != nil {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	signature, err := base64.RawURLEncoding.DecodeString(string(parts[1]))
	if err != nil {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	mac := hmac.New(sha256.New, secret)
	_, _ = mac.Write(payload)
	if !hmac.Equal(signature, mac.Sum(nil)) {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	var claims FileReferenceClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	return claims, nil
}
