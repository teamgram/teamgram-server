package repository

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
)

type FileReferenceClaims struct {
	MediaID      int64  `json:"media_id"`
	ObjectID     string `json:"object_id"`
	OriginDomain string `json:"origin_domain"`
	OriginID     int64  `json:"origin_id"`
	ExpireAt     int64  `json:"expire_at"`
	AccessHash   int64  `json:"access_hash"`
}

type FileReferenceService struct {
	secret []byte
	now    func() time.Time
}

func NewFileReferenceService(secret []byte, now func() time.Time) *FileReferenceService {
	if now == nil {
		now = time.Now
	}
	return &FileReferenceService{
		secret: append([]byte(nil), secret...),
		now:    now,
	}
}

func (s *FileReferenceService) Generate(claims FileReferenceClaims) ([]byte, error) {
	if s == nil || len(s.secret) == 0 {
		return nil, media.ErrFileReferenceInvalid
	}
	payload, err := json.Marshal(claims)
	if err != nil {
		return nil, fmt.Errorf("%w: marshal claims: %v", media.ErrFileReferenceInvalid, err)
	}
	signature := s.sign(payload)
	token := make([]byte, 0, base64.RawURLEncoding.EncodedLen(len(payload))+1+base64.RawURLEncoding.EncodedLen(len(signature)))
	token = base64.RawURLEncoding.AppendEncode(token, payload)
	token = append(token, '.')
	token = base64.RawURLEncoding.AppendEncode(token, signature)
	return token, nil
}

func (s *FileReferenceService) Validate(token []byte) (FileReferenceClaims, error) {
	if s == nil || len(s.secret) == 0 {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	if len(token) == 0 {
		return FileReferenceClaims{}, media.ErrFileReferenceEmpty
	}

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
	if !hmac.Equal(signature, s.sign(payload)) {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}

	var claims FileReferenceClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	if claims.ExpireAt <= s.now().Unix() {
		return FileReferenceClaims{}, media.ErrFileReferenceExpired
	}
	return claims, nil
}

func (s *FileReferenceService) sign(payload []byte) []byte {
	mac := hmac.New(sha256.New, s.secret)
	_, _ = mac.Write(payload)
	return mac.Sum(nil)
}
