package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeAuthKeysModel struct {
	model.AuthKeysModel
	findOneByAuthKeyId func(ctx context.Context, authKeyId int64) (*model.AuthKeys, error)
}

func (m fakeAuthKeysModel) FindOneByAuthKeyId(ctx context.Context, authKeyId int64) (*model.AuthKeys, error) {
	return m.findOneByAuthKeyId(ctx, authKeyId)
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
