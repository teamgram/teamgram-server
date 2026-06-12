package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeAuthKeysModel struct {
	model.AuthKeysModel
	findOneByAuthKeyId  func(ctx context.Context, authKeyId int64) (*model.AuthKeys, error)
	findListByAuthKeyId func(ctx context.Context, authKeyId ...int64) ([]model.AuthKeys, error)
	insertIgnore        func(ctx context.Context, data *model.AuthKeys) (int64, int64, error)
	updatePermBinding   func(ctx context.Context, permAuthKeyId int64, authKeyId int64) (int64, error)
	updateTempBinding   func(ctx context.Context, tempAuthKeyId int64, authKeyId int64) (int64, error)
}

func (m fakeAuthKeysModel) FindOneByAuthKeyId(ctx context.Context, authKeyId int64) (*model.AuthKeys, error) {
	return m.findOneByAuthKeyId(ctx, authKeyId)
}

func (m fakeAuthKeysModel) FindListByAuthKeyIdList(ctx context.Context, authKeyId ...int64) ([]model.AuthKeys, error) {
	return m.findListByAuthKeyId(ctx, authKeyId...)
}

func (m fakeAuthKeysModel) InsertIgnore(ctx context.Context, data *model.AuthKeys) (int64, int64, error) {
	if m.insertIgnore == nil {
		return 0, 1, nil
	}
	return m.insertIgnore(ctx, data)
}

func (m fakeAuthKeysModel) UpdatePermBinding(ctx context.Context, permAuthKeyId int64, authKeyId int64) (int64, error) {
	if m.updatePermBinding == nil {
		return 1, nil
	}
	return m.updatePermBinding(ctx, permAuthKeyId, authKeyId)
}

func (m fakeAuthKeysModel) UpdateTempBinding(ctx context.Context, tempAuthKeyId int64, authKeyId int64) (int64, error) {
	if m.updateTempBinding == nil {
		return 1, nil
	}
	return m.updateTempBinding(ctx, tempAuthKeyId, authKeyId)
}

type fakeAuthKeyLifecycleModel struct {
	activate func(ctx context.Context, authKeyId int64, ttlSeconds int) error
	isActive func(ctx context.Context, authKeyId int64) (bool, error)
	revoke   func(ctx context.Context, authKeyId int64) error
}

func (m *fakeAuthKeyLifecycleModel) Activate(ctx context.Context, authKeyId int64, ttlSeconds int) error {
	return m.activate(ctx, authKeyId, ttlSeconds)
}

func (m *fakeAuthKeyLifecycleModel) IsActive(ctx context.Context, authKeyId int64) (bool, error) {
	return m.isActive(ctx, authKeyId)
}

func (m *fakeAuthKeyLifecycleModel) Revoke(ctx context.Context, authKeyId int64) error {
	if m.revoke == nil {
		return nil
	}
	return m.revoke(ctx, authKeyId)
}

func TestAuthKeyInfoMapping(t *testing.T) {
	authKey := tg.MakeTLAuthKeyInfo(&tg.TLAuthKeyInfo{
		AuthKeyId:          1001,
		AuthKey:            []byte("test-auth-key"),
		AuthKeyType:        tg.AuthKeyTypeTemp,
		PermAuthKeyId:      2002,
		TempAuthKeyId:      3003,
		MediaTempAuthKeyId: 4004,
	})

	row := fromAuthKeyInfo(authKey)
	got, err := toAuthKeyInfo(row)
	if err != nil {
		t.Fatalf("toAuthKeyInfo() error = %v", err)
	}

	if got.AuthKeyId != authKey.AuthKeyId ||
		string(got.AuthKey) != string(authKey.AuthKey) ||
		got.AuthKeyType != authKey.AuthKeyType ||
		got.PermAuthKeyId != authKey.PermAuthKeyId ||
		got.TempAuthKeyId != authKey.TempAuthKeyId ||
		got.MediaTempAuthKeyId != authKey.MediaTempAuthKeyId {
		t.Fatalf("mapped auth key mismatch: got %#v want %#v", got, authKey)
	}
}

func TestAuthKeyInfoInvalidBody(t *testing.T) {
	_, err := toAuthKeyInfo(&model.AuthKeys{
		AuthKeyId: 1,
		Body:      "not base64!",
	})
	if err == nil {
		t.Fatal("toAuthKeyInfo() error = nil, want invalid base64 error")
	}
}

func TestRepositoryErrorHelpers(t *testing.T) {
	if !errors.Is(wrapStorage(sqlx.ErrNotFound), authsession.ErrAuthSessionStorage) {
		t.Fatal("wrapStorage() does not preserve storage sentinel")
	}
	if !isNotFound(&model.NotFoundError{Resource: "auth_keys", Key: "auth_key_id=1", Cause: sqlx.ErrNotFound}) {
		t.Fatal("isNotFound() did not recognize model not-found errors")
	}
	if isNotFound(sqlx.ErrNotFound) {
		t.Fatal("isNotFound() should not recognize storage not-found directly")
	}
}

func TestQueryAuthKeyReturnsDomainNotFoundOnModelNotFound(t *testing.T) {
	repo := &Repository{
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{
				findOneByAuthKeyId: func(ctx context.Context, authKeyId int64) (*model.AuthKeys, error) {
					return nil, &model.NotFoundError{
						Resource: "auth_keys",
						Key:      "auth_key_id=1001",
						Cause:    sqlx.ErrNotFound,
					}
				},
			},
		},
	}

	_, err := repo.QueryAuthKey(context.Background(), 1001)
	if !errors.Is(err, authsession.ErrAuthKeyNotFound) {
		t.Fatalf("QueryAuthKey() error = %v, want ErrAuthKeyNotFound", err)
	}
}

func TestQueryAuthKeyReturnsDomainNotFoundOnNilModelRow(t *testing.T) {
	repo := &Repository{
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{
				findOneByAuthKeyId: func(ctx context.Context, authKeyId int64) (*model.AuthKeys, error) {
					return nil, nil
				},
			},
		},
	}

	_, err := repo.QueryAuthKey(context.Background(), 1001)
	if !errors.Is(err, authsession.ErrAuthKeyNotFound) {
		t.Fatalf("QueryAuthKey() error = %v, want ErrAuthKeyNotFound", err)
	}
}

func TestSaveAuthKeyDefaultsTempTTLToSevenDays(t *testing.T) {
	const sevenDays = 7 * 24 * 60 * 60

	var (
		gotTTL int
		gotId  int64
	)
	repo := &Repository{
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{},
		},
		authKeyLifecycleModel: &fakeAuthKeyLifecycleModel{
			activate: func(ctx context.Context, authKeyId int64, ttlSeconds int) error {
				gotId = authKeyId
				gotTTL = ttlSeconds
				return nil
			},
		},
	}

	err := repo.SaveAuthKey(context.Background(), tg.MakeTLAuthKeyInfo(&tg.TLAuthKeyInfo{
		AuthKeyId:   1001,
		AuthKey:     []byte("body"),
		AuthKeyType: tg.AuthKeyTypeTemp,
	}), 0)
	if err != nil {
		t.Fatalf("SaveAuthKey() error = %v", err)
	}
	if gotId != 1001 {
		t.Fatalf("Activate authKeyId = %d, want 1001", gotId)
	}
	if gotTTL != sevenDays {
		t.Fatalf("Activate ttl = %d, want %d", gotTTL, sevenDays)
	}
}

func TestSaveAuthKeyHonorsExplicitTTL(t *testing.T) {
	var gotTTL int
	repo := &Repository{
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{},
		},
		authKeyLifecycleModel: &fakeAuthKeyLifecycleModel{
			activate: func(ctx context.Context, authKeyId int64, ttlSeconds int) error {
				gotTTL = ttlSeconds
				return nil
			},
		},
	}

	err := repo.SaveAuthKey(context.Background(), tg.MakeTLAuthKeyInfo(&tg.TLAuthKeyInfo{
		AuthKeyId:   1001,
		AuthKey:     []byte("body"),
		AuthKeyType: tg.AuthKeyTypeMediaTemp,
	}), 3600)
	if err != nil {
		t.Fatalf("SaveAuthKey() error = %v", err)
	}
	if gotTTL != 3600 {
		t.Fatalf("Activate ttl = %d, want 3600", gotTTL)
	}
}

func TestSaveAuthKeySkipsLifecycleForPerm(t *testing.T) {
	called := false
	repo := &Repository{
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{},
		},
		authKeyLifecycleModel: &fakeAuthKeyLifecycleModel{
			activate: func(ctx context.Context, authKeyId int64, ttlSeconds int) error {
				called = true
				return nil
			},
		},
	}

	err := repo.SaveAuthKey(context.Background(), tg.MakeTLAuthKeyInfo(&tg.TLAuthKeyInfo{
		AuthKeyId:   1001,
		AuthKey:     []byte("body"),
		AuthKeyType: tg.AuthKeyTypePerm,
	}), 0)
	if err != nil {
		t.Fatalf("SaveAuthKey() error = %v", err)
	}
	if called {
		t.Fatal("Activate must not be called for perm keys")
	}
}

func TestQueryAuthKeyEvictsExpiredTempKey(t *testing.T) {
	repo := &Repository{
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{
				findOneByAuthKeyId: func(ctx context.Context, authKeyId int64) (*model.AuthKeys, error) {
					return &model.AuthKeys{
						AuthKeyId:   authKeyId,
						Body:        "YWJj",
						AuthKeyType: tg.AuthKeyTypeTemp,
					}, nil
				},
			},
		},
		authKeyLifecycleModel: &fakeAuthKeyLifecycleModel{
			isActive: func(ctx context.Context, authKeyId int64) (bool, error) {
				return false, nil
			},
		},
	}

	_, err := repo.QueryAuthKey(context.Background(), 1001)
	if !errors.Is(err, authsession.ErrAuthKeyNotFound) {
		t.Fatalf("QueryAuthKey() error = %v, want ErrAuthKeyNotFound", err)
	}
}

func TestQueryAuthKeyKeepsUnexpiredTempKeyWhenLifecycleMarkerMissing(t *testing.T) {
	repo := &Repository{
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{
				findOneByAuthKeyId: func(ctx context.Context, authKeyId int64) (*model.AuthKeys, error) {
					return &model.AuthKeys{
						AuthKeyId:   authKeyId,
						Body:        "YWJj",
						AuthKeyType: tg.AuthKeyTypeTemp,
						ExpiresAt:   time.Now().UTC().Unix() + 3600,
					}, nil
				},
			},
		},
		authKeyLifecycleModel: &fakeAuthKeyLifecycleModel{
			isActive: func(ctx context.Context, authKeyId int64) (bool, error) {
				return false, nil
			},
		},
	}

	got, err := repo.QueryAuthKey(context.Background(), 1001)
	if err != nil {
		t.Fatalf("QueryAuthKey() error = %v", err)
	}
	if got.AuthKeyId != 1001 {
		t.Fatalf("got auth key id = %d, want 1001", got.AuthKeyId)
	}
}

func TestQueryAuthKeySkipsLifecycleForPerm(t *testing.T) {
	repo := &Repository{
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{
				findOneByAuthKeyId: func(ctx context.Context, authKeyId int64) (*model.AuthKeys, error) {
					return &model.AuthKeys{
						AuthKeyId:   authKeyId,
						Body:        "YWJj",
						AuthKeyType: tg.AuthKeyTypePerm,
					}, nil
				},
			},
		},
		authKeyLifecycleModel: &fakeAuthKeyLifecycleModel{
			isActive: func(ctx context.Context, authKeyId int64) (bool, error) {
				t.Fatal("perm keys must not consult lifecycle store")
				return false, nil
			},
		},
	}

	got, err := repo.QueryAuthKey(context.Background(), 1001)
	if err != nil {
		t.Fatalf("QueryAuthKey() error = %v", err)
	}
	if got.AuthKeyId != 1001 {
		t.Fatalf("got auth key id = %d, want 1001", got.AuthKeyId)
	}
}

func TestBindKeyIdRejectsZeroBindKey(t *testing.T) {
	repo := &Repository{
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{
				updateTempBinding: func(ctx context.Context, tempAuthKeyId int64, authKeyId int64) (int64, error) {
					t.Fatal("UpdateTempBinding must not be called when bindKeyId is zero")
					return 0, nil
				},
			},
		},
	}

	err := repo.BindKeyId(context.Background(), 1001, tg.AuthKeyTypeTemp, 0)
	if !errors.Is(err, authsession.ErrAuthKeyInvalid) {
		t.Fatalf("BindKeyId() error = %v, want ErrAuthKeyInvalid", err)
	}

	err = repo.BindKeyId(context.Background(), 0, tg.AuthKeyTypeTemp, 1)
	if !errors.Is(err, authsession.ErrAuthKeyInvalid) {
		t.Fatalf("BindKeyId() error = %v, want ErrAuthKeyInvalid", err)
	}
}

func TestBindKeyIdReturnsDomainNotFoundOnModelNotFound(t *testing.T) {
	repo := &Repository{
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{
				updateTempBinding: func(ctx context.Context, tempAuthKeyId int64, authKeyId int64) (int64, error) {
					return 0, &model.NotFoundError{
						Resource: "auth_keys",
						Key:      "auth_key_id=1001",
						Cause:    model.ErrNotFound,
					}
				},
			},
		},
	}

	err := repo.BindKeyId(context.Background(), 1001, tg.AuthKeyTypeTemp, 2002)
	if !errors.Is(err, authsession.ErrAuthKeyNotFound) {
		t.Fatalf("BindKeyId() error = %v, want ErrAuthKeyNotFound", err)
	}
}

func TestBindKeyIdReturnsDomainNotFoundOnZeroRows(t *testing.T) {
	repo := &Repository{
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{
				updateTempBinding: func(ctx context.Context, tempAuthKeyId int64, authKeyId int64) (int64, error) {
					return 0, nil
				},
			},
		},
	}

	err := repo.BindKeyId(context.Background(), 1001, tg.AuthKeyTypeTemp, 2002)
	if !errors.Is(err, authsession.ErrAuthKeyNotFound) {
		t.Fatalf("BindKeyId() error = %v, want ErrAuthKeyNotFound", err)
	}
}

func TestExpandAuthKeyIdsUsesBatchRows(t *testing.T) {
	repo := &Repository{
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{
				findListByAuthKeyId: func(ctx context.Context, authKeyId ...int64) ([]model.AuthKeys, error) {
					if len(authKeyId) != 2 || authKeyId[0] != 1001 || authKeyId[1] != 1002 {
						t.Fatalf("FindListByAuthKeyIdList ids = %v, want [1001 1002]", authKeyId)
					}
					return []model.AuthKeys{
						{AuthKeyId: 1001, Body: "YWJj", TempAuthKeyId: 2001},
						{AuthKeyId: 1002, Body: "ZGVm"},
					}, nil
				},
			},
		},
	}

	got, err := repo.ExpandAuthKeyIds(context.Background(), []int64{1001, 1002})
	if err != nil {
		t.Fatalf("ExpandAuthKeyIds() error = %v", err)
	}
	if len(got) != 2 || got[0] != 2001 || got[1] != 1002 {
		t.Fatalf("ExpandAuthKeyIds() = %v, want [2001 1002]", got)
	}
}
