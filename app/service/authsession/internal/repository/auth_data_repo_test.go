package repository

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/teamgram/marmota/pkg/stores/cache"
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type noOpBatchCache struct{}

func (noOpBatchCache) Del(keys ...string) error { return nil }
func (noOpBatchCache) DelCtx(ctx context.Context, keys ...string) error {
	return nil
}
func (noOpBatchCache) Get(key string, val any) error { return model.ErrNotFound }
func (noOpBatchCache) GetCtx(ctx context.Context, key string, val any) error {
	return model.ErrNotFound
}
func (noOpBatchCache) IsNotFound(err error) bool     { return errors.Is(err, model.ErrNotFound) }
func (noOpBatchCache) Set(key string, val any) error { return nil }
func (noOpBatchCache) SetCtx(ctx context.Context, key string, val any) error {
	return nil
}
func (noOpBatchCache) SetWithExpire(key string, val any, expire time.Duration) error {
	return nil
}
func (noOpBatchCache) SetWithExpireCtx(ctx context.Context, key string, val any, expire time.Duration) error {
	return nil
}
func (noOpBatchCache) Take(val any, key string, query func(val any) error) error {
	return query(val)
}
func (noOpBatchCache) TakeCtx(ctx context.Context, val any, key string, query func(val any) error) error {
	return query(val)
}
func (noOpBatchCache) TakeWithExpire(val any, key string, query func(val any, expire time.Duration) error) error {
	return query(val, 0)
}
func (noOpBatchCache) TakeWithExpireCtx(ctx context.Context, val any, key string, query func(val any, expire time.Duration) error) error {
	return query(val, 0)
}
func (noOpBatchCache) Takes(query func(keys ...string) (map[string]any, error), cacheF func(k, v string) (any, error), keys ...string) error {
	_, err := query(keys...)
	return err
}
func (noOpBatchCache) TakesCtx(ctx context.Context, query func(keys ...string) (map[string]any, error), cacheF func(k, v string) (any, error), keys ...string) error {
	_, err := query(keys...)
	return err
}
func (noOpBatchCache) TakesWithExpire(query func(expire time.Duration, keys ...string) (map[string]any, error), cacheF func(k, v string) (any, error), keys ...string) error {
	_, err := query(0, keys...)
	return err
}
func (noOpBatchCache) TakesWithExpireCtx(ctx context.Context, query func(expire time.Duration, keys ...string) (map[string]any, error), cacheF func(k, v string) (any, error), keys ...string) error {
	_, err := query(0, keys...)
	return err
}

var _ cache.BatchCache = noOpBatchCache{}

type fakeAuthUsersPushModel struct {
	model.AuthUsersModel
	updateAndroidPushSessionId func(ctx context.Context, androidPushSessionId int64, authKeyId int64, userId int64) (int64, error)
	selectListByUserId         func(ctx context.Context, userId int64) ([]model.AuthUsers, error)
}

func (m fakeAuthUsersPushModel) UpdateAndroidPushSessionId(ctx context.Context, androidPushSessionId int64, authKeyId int64, userId int64) (int64, error) {
	return m.updateAndroidPushSessionId(ctx, androidPushSessionId, authKeyId, userId)
}

func (m fakeAuthUsersPushModel) SelectListByUserId(ctx context.Context, userId int64) ([]model.AuthUsers, error) {
	return m.selectListByUserId(ctx, userId)
}

func TestAuthDataStateMapping(t *testing.T) {
	tests := []struct {
		name string
		data *cacheAuthData
		want int32
	}{
		{name: "nil aggregate", data: nil, want: tg.AuthStateNew},
		{name: "nil client", data: &cacheAuthData{}, want: tg.AuthStateWaitInit},
		{name: "client without user", data: &cacheAuthData{Client: &clientSessionCacheData{}}, want: tg.AuthStateUnauthorized},
		{name: "client and user", data: &cacheAuthData{Client: &clientSessionCacheData{}, BindUser: &bindUser{UserId: 42}}, want: tg.AuthStateNormal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toAuthKeyStateData(1001, tt.data)
			if got.KeyState != tt.want {
				t.Fatalf("KeyState = %d, want %d", got.KeyState, tt.want)
			}
		})
	}
}

func TestAuthDataCachePayloadDoesNotUseTLDebugJSON(t *testing.T) {
	data := &cacheAuthData{
		Client: &clientSessionCacheData{
			AuthKeyId: 1001,
			Ip:        "127.0.0.1",
			Layer:     158,
			Params:    "{}",
		},
	}
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal cacheAuthData error: %v", err)
	}
	payload := string(b)
	if strings.Contains(payload, `"_name"`) || strings.Contains(payload, `"_id"`) {
		t.Fatalf("cache payload contains TL debug JSON metadata: %s", payload)
	}
	if !strings.Contains(payload, `"auth_key_id":1001`) || !strings.Contains(payload, `"params":"{}"`) {
		t.Fatalf("cache payload missing service-owned client fields: %s", payload)
	}
}

func TestAuthDataClientSessionMapping(t *testing.T) {
	got := toClientSession(1001, &model.Auths{
		AuthKeyId:      2002,
		ClientIp:       "127.0.0.1",
		Layer:          158,
		ApiId:          9,
		DeviceModel:    "device",
		SystemVersion:  "system",
		AppVersion:     "app",
		SystemLangCode: "en-US",
		LangPack:       "android",
		LangCode:       "en",
		Proxy:          "proxy",
		Params:         "{}",
	})
	if got.AuthKeyId != 1001 || got.Ip != "127.0.0.1" || got.Layer != 158 || got.Params != "{}" {
		t.Fatalf("mapped client session mismatch: %#v", got)
	}
}

func TestClientKindAndLangPackMapping(t *testing.T) {
	if got := normalizeLangPack("", "Telegram A"); got != "weba" {
		t.Fatalf("normalizeLangPack() = %q, want weba", got)
	}
	if got := normalizeLangPack("android", "Telegram TDLib"); got != "android" {
		t.Fatalf("normalizeLangPack() = %q, want android", got)
	}
}

func TestGetPermAuthKeyIdsByUserIdReturnsActiveBindings(t *testing.T) {
	repo := &Repository{
		model: &model.Models{
			AuthUsersModel: fakeAuthUsersPushModel{
				selectListByUserId: func(ctx context.Context, userId int64) ([]model.AuthUsers, error) {
					if userId != 42 {
						t.Fatalf("user_id = %d, want 42", userId)
					}
					return []model.AuthUsers{
						{AuthKeyId: 1001, UserId: userId},
						{AuthKeyId: 1002, UserId: userId},
					}, nil
				},
			},
		},
	}

	got, err := repo.GetPermAuthKeyIdsByUserId(context.Background(), 42)
	if err != nil {
		t.Fatalf("GetPermAuthKeyIdsByUserId() error = %v", err)
	}
	if len(got) != 2 || got[0] != 1001 || got[1] != 1002 {
		t.Fatalf("GetPermAuthKeyIdsByUserId() = %#v", got)
	}
}

func TestSetAndroidPushSessionIdReturnsAuthorizationNotFoundOnModelNotFound(t *testing.T) {
	repo := &Repository{
		CachedConn: sqlc.NewConnWithCache(nil, noOpBatchCache{}),
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{
				findOneByAuthKeyId: func(ctx context.Context, authKeyId int64) (*model.AuthKeys, error) {
					return &model.AuthKeys{
						AuthKeyId:     authKeyId,
						Body:          "YWJj",
						AuthKeyType:   tg.AuthKeyTypePerm,
						PermAuthKeyId: authKeyId,
					}, nil
				},
			},
			AuthUsersModel: fakeAuthUsersPushModel{
				updateAndroidPushSessionId: func(ctx context.Context, androidPushSessionId int64, authKeyId int64, userId int64) (int64, error) {
					return 0, &model.NotFoundError{
						Resource: "auth_users",
						Key:      "auth_key_id=1001,user_id=42",
						Cause:    model.ErrNotFound,
					}
				},
			},
		},
	}

	err := repo.SetAndroidPushSessionIdByAuthKeyId(context.Background(), 42, 1001, 3003)
	if !errors.Is(err, authsession.ErrAuthorizationNotFound) {
		t.Fatalf("SetAndroidPushSessionIdByAuthKeyId() error = %v, want ErrAuthorizationNotFound", err)
	}
}

func TestSetAndroidPushSessionIdReturnsAuthorizationNotFoundOnZeroRows(t *testing.T) {
	repo := &Repository{
		CachedConn: sqlc.NewConnWithCache(nil, noOpBatchCache{}),
		model: &model.Models{
			AuthKeysModel: fakeAuthKeysModel{
				findOneByAuthKeyId: func(ctx context.Context, authKeyId int64) (*model.AuthKeys, error) {
					return &model.AuthKeys{
						AuthKeyId:     authKeyId,
						Body:          "YWJj",
						AuthKeyType:   tg.AuthKeyTypePerm,
						PermAuthKeyId: authKeyId,
					}, nil
				},
			},
			AuthUsersModel: fakeAuthUsersPushModel{
				updateAndroidPushSessionId: func(ctx context.Context, androidPushSessionId int64, authKeyId int64, userId int64) (int64, error) {
					return 0, nil
				},
			},
		},
	}

	err := repo.SetAndroidPushSessionIdByAuthKeyId(context.Background(), 42, 1001, 3003)
	if !errors.Is(err, authsession.ErrAuthorizationNotFound) {
		t.Fatalf("SetAndroidPushSessionIdByAuthKeyId() error = %v, want ErrAuthorizationNotFound", err)
	}
}
