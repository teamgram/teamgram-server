package repository

import (
	"errors"
	"testing"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

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
	if !errors.Is(wrapStorage(sqlc.ErrNotFound), authsession.ErrAuthSessionStorage) {
		t.Fatal("wrapStorage() does not preserve storage sentinel")
	}
	if !isNotFound(model.ErrNotFound) || !isNotFound(sqlx.ErrNotFound) || !isNotFound(sqlc.ErrNotFound) {
		t.Fatal("isNotFound() did not recognize repository not-found errors")
	}
}
