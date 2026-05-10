package repository

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
)

func TestFileReferenceGenerateValidateStatefulHandle(t *testing.T) {
	now := time.Unix(1700000000, 0)
	store := newMemoryFileReferenceStore()
	svc := NewFileReferenceService([]byte("ignored-for-opaque-handles"), func() time.Time { return now })
	claims := FileReferenceClaims{
		MediaID:      100,
		ObjectID:     "object-100",
		OriginDomain: "photo",
		OriginID:     10,
		ExpireAt:     now.Add(time.Hour).Unix(),
		AccessHash:   200,
	}

	token, err := svc.Generate(context.Background(), claims, store)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if len(token) != 25 {
		t.Fatalf("len(token) = %d, want 25", len(token))
	}
	got, err := svc.Validate(context.Background(), token, store)
	if err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if got != claims {
		t.Fatalf("Validate() claims = %#v, want %#v", got, claims)
	}
}

func TestFileReferenceValidateLegacyHMACTokenDuringMigration(t *testing.T) {
	now := time.Unix(1700000000, 0)
	legacy := NewLegacyFileReferenceServiceForTest([]byte("legacy-secret"), func() time.Time { return now })
	legacyToken, err := legacy.Generate(FileReferenceClaims{
		MediaID:      101,
		ObjectID:     "legacy-object",
		OriginDomain: "photo",
		OriginID:     10,
		ExpireAt:     now.Add(time.Hour).Unix(),
		AccessHash:   202,
	})
	if err != nil {
		t.Fatalf("legacy Generate() error = %v", err)
	}
	if len(legacyToken) == 25 {
		t.Fatalf("legacy token unexpectedly has new handle length")
	}

	svc := NewFileReferenceService([]byte("legacy-secret"), func() time.Time { return now })
	claims, err := svc.Validate(context.Background(), legacyToken, newMemoryFileReferenceStore())
	if err != nil {
		t.Fatalf("Validate(legacy token) error = %v", err)
	}
	if claims.MediaID != 101 || claims.ObjectID != "legacy-object" || claims.AccessHash != 202 {
		t.Fatalf("legacy claims = %#v, want decoded claims", claims)
	}
}

func TestFileReferenceValidateEmpty(t *testing.T) {
	svc := NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) })

	_, err := svc.Validate(context.Background(), nil, newMemoryFileReferenceStore())
	if !errors.Is(err, media.ErrFileReferenceEmpty) {
		t.Fatalf("Validate(nil) error = %v, want ErrFileReferenceEmpty", err)
	}
}

func TestFileReferenceRejectsNilServiceOrStore(t *testing.T) {
	now := time.Unix(1700000000, 0)
	claims := FileReferenceClaims{
		MediaID:    10,
		ObjectID:   "obj-1",
		ExpireAt:   now.Add(time.Hour).Unix(),
		AccessHash: 30,
	}

	_, err := (*FileReferenceService)(nil).Generate(context.Background(), claims, newMemoryFileReferenceStore())
	if !errors.Is(err, media.ErrFileReferenceInvalid) {
		t.Fatalf("Generate(nil service) error = %v, want ErrFileReferenceInvalid", err)
	}

	svc := NewFileReferenceService([]byte("test-secret"), func() time.Time { return now })
	_, err = svc.Generate(context.Background(), claims, nil)
	if !errors.Is(err, media.ErrFileReferenceInvalid) {
		t.Fatalf("Generate(nil store) error = %v, want ErrFileReferenceInvalid", err)
	}

	_, err = svc.Validate(context.Background(), []byte{fileReferenceOpaqueVersion}, nil)
	if !errors.Is(err, media.ErrFileReferenceInvalid) {
		t.Fatalf("Validate(nil store) error = %v, want ErrFileReferenceInvalid", err)
	}
}

func TestFileReferenceValidateExpired(t *testing.T) {
	now := time.Unix(1700000000, 0)
	store := newMemoryFileReferenceStore()
	svc := NewFileReferenceService([]byte("test-secret"), func() time.Time { return now })
	token, err := svc.Generate(context.Background(), FileReferenceClaims{
		MediaID:    10,
		ObjectID:   "obj-1",
		ExpireAt:   now.Add(-time.Second).Unix(),
		AccessHash: 30,
	}, store)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	_, err = svc.Validate(context.Background(), token, store)
	if !errors.Is(err, media.ErrFileReferenceExpired) {
		t.Fatalf("Validate(expired) error = %v, want ErrFileReferenceExpired", err)
	}
}

func TestFileReferenceValidateInvalidHandle(t *testing.T) {
	now := time.Unix(1700000000, 0)
	svc := NewFileReferenceService([]byte("test-secret"), func() time.Time { return now })
	token := make([]byte, fileReferenceOpaqueLength)
	token[0] = fileReferenceOpaqueVersion
	token[1] = 1

	_, err := svc.Validate(context.Background(), token, newMemoryFileReferenceStore())
	if !errors.Is(err, media.ErrFileReferenceInvalid) {
		t.Fatalf("Validate(unknown handle) error = %v, want ErrFileReferenceInvalid", err)
	}
}

func TestFileReferenceValidateTamperedLegacyToken(t *testing.T) {
	now := time.Unix(1700000000, 0)
	legacy := NewLegacyFileReferenceServiceForTest([]byte("test-secret"), func() time.Time { return now })
	token, err := legacy.Generate(FileReferenceClaims{
		MediaID:    10,
		ObjectID:   "obj-1",
		ExpireAt:   now.Add(time.Hour).Unix(),
		AccessHash: 30,
	})
	if err != nil {
		t.Fatalf("legacy Generate() error = %v", err)
	}
	token[len(token)-1] ^= 1

	svc := NewFileReferenceService([]byte("test-secret"), func() time.Time { return now })
	_, err = svc.Validate(context.Background(), token, newMemoryFileReferenceStore())
	if !errors.Is(err, media.ErrFileReferenceInvalid) {
		t.Fatalf("Validate(tampered legacy) error = %v, want ErrFileReferenceInvalid", err)
	}
}

func TestRepositoryFileReferenceStoreMapsSemanticErrors(t *testing.T) {
	ctx := context.Background()
	token := []byte{fileReferenceOpaqueVersion, 1, 2, 3}
	claims := FileReferenceClaims{
		MediaID:      10,
		ObjectID:     "object-10",
		OriginDomain: "photo",
		OriginID:     20,
		ExpireAt:     time.Unix(1700000000, 0).Add(time.Hour).Unix(),
		AccessHash:   30,
	}
	fileReferences := newCaptureFileReferencesModel()
	repo := &Repository{model: &model.Models{FileReferencesModel: fileReferences}}

	if err := repo.SaveFileReference(ctx, token, claims); err != nil {
		t.Fatalf("SaveFileReference() error = %v", err)
	}
	got, err := repo.LoadFileReference(ctx, token)
	if err != nil {
		t.Fatalf("LoadFileReference() error = %v", err)
	}
	if got != claims {
		t.Fatalf("LoadFileReference() = %#v, want %#v", got, claims)
	}

	_, err = repo.LoadFileReference(ctx, []byte{fileReferenceOpaqueVersion, 9})
	if !errors.Is(err, media.ErrFileReferenceInvalid) {
		t.Fatalf("LoadFileReference(not found) error = %v, want ErrFileReferenceInvalid", err)
	}

	fileReferences.rows[string(token)].RevokedAt = 1
	_, err = repo.LoadFileReference(ctx, token)
	if !errors.Is(err, media.ErrFileReferenceInvalid) {
		t.Fatalf("LoadFileReference(revoked) error = %v, want ErrFileReferenceInvalid", err)
	}

	storageErr := errors.New("db down")
	storageModel := newCaptureFileReferencesModel()
	storageModel.insertErr = storageErr
	repo = &Repository{model: &model.Models{FileReferencesModel: storageModel}}
	err = repo.SaveFileReference(ctx, token, claims)
	if !errors.Is(err, media.ErrMediaStorage) || !errors.Is(err, storageErr) {
		t.Fatalf("SaveFileReference(storage) error = %v, want ErrMediaStorage wrapping db error", err)
	}

	storageModel = newCaptureFileReferencesModel()
	storageModel.selectErr = storageErr
	repo = &Repository{model: &model.Models{FileReferencesModel: storageModel}}
	_, err = repo.LoadFileReference(ctx, token)
	if !errors.Is(err, media.ErrMediaStorage) || !errors.Is(err, storageErr) {
		t.Fatalf("LoadFileReference(storage) error = %v, want ErrMediaStorage wrapping db error", err)
	}
}

type memoryFileReferenceStore struct {
	rows map[string]FileReferenceClaims
}

func newMemoryFileReferenceStore() *memoryFileReferenceStore {
	return &memoryFileReferenceStore{rows: map[string]FileReferenceClaims{}}
}

func (s *memoryFileReferenceStore) SaveFileReference(_ context.Context, token []byte, claims FileReferenceClaims) error {
	s.rows[string(token)] = claims
	return nil
}

func (s *memoryFileReferenceStore) LoadFileReference(_ context.Context, token []byte) (FileReferenceClaims, error) {
	claims, ok := s.rows[string(token)]
	if !ok {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	return claims, nil
}

type captureFileReferencesModel struct {
	model.FileReferencesModel
	inserted  []*model.FileReferences
	rows      map[string]*model.FileReferences
	insertErr error
	selectErr error
}

func newCaptureFileReferencesModel() *captureFileReferencesModel {
	return &captureFileReferencesModel{rows: map[string]*model.FileReferences{}}
}

func (m *captureFileReferencesModel) Insert(_ context.Context, data *model.FileReferences) (int64, int64, error) {
	if m.insertErr != nil {
		return 0, 0, m.insertErr
	}
	row := *data
	row.RefHash = append([]byte(nil), data.RefHash...)
	m.inserted = append(m.inserted, &row)
	m.rows[string(row.RefHash)] = &row
	return 1, 1, nil
}

func (m *captureFileReferencesModel) SelectByRefHash(_ context.Context, refHash []byte) (*model.FileReferences, error) {
	if m.selectErr != nil {
		return nil, fmt.Errorf("select: %w", m.selectErr)
	}
	row := m.rows[string(refHash)]
	if row == nil {
		return nil, &model.NotFoundError{Resource: "file_references", Key: "ref_hash"}
	}
	out := *row
	out.RefHash = append([]byte(nil), row.RefHash...)
	return &out, nil
}

type legacyFileReferenceServiceForTest struct {
	secret []byte
	now    func() time.Time
}

func NewLegacyFileReferenceServiceForTest(secret []byte, now func() time.Time) *legacyFileReferenceServiceForTest {
	return &legacyFileReferenceServiceForTest{secret: append([]byte(nil), secret...), now: now}
}

func (s *legacyFileReferenceServiceForTest) Generate(claims FileReferenceClaims) ([]byte, error) {
	payload, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}
	mac := hmac.New(sha256.New, s.secret)
	_, _ = mac.Write(payload)
	token := make([]byte, 0, base64.RawURLEncoding.EncodedLen(len(payload))+1+base64.RawURLEncoding.EncodedLen(sha256.Size))
	token = base64.RawURLEncoding.AppendEncode(token, payload)
	token = append(token, '.')
	token = base64.RawURLEncoding.AppendEncode(token, mac.Sum(nil))
	return token, nil
}
