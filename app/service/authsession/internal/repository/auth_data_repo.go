package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const authDataCachePrefix = "authsession:auth_data:v1"

type bindUser struct {
	UserId               int64 `json:"user_id"`
	Hash                 int64 `json:"hash"`
	DateCreated          int64 `json:"date_created"`
	DateActivated        int64 `json:"date_activated"`
	AndroidPushSessionId int64 `json:"android_push_session_id"`
}

type cacheAuthData struct {
	Client   *authsession.ClientSession `json:"client"`
	BindUser *bindUser                  `json:"bind_user,omitempty"`
}

func authDataCacheKey(authKeyId int64) string {
	return fmt.Sprintf("%s:%d", authDataCachePrefix, authKeyId)
}

func (c *cacheAuthData) toAuthState() int32 {
	switch {
	case c == nil:
		return tg.AuthStateNew
	case c.Client == nil:
		return tg.AuthStateWaitInit
	case c.BindUser == nil:
		return tg.AuthStateUnauthorized
	default:
		return tg.AuthStateNormal
	}
}

func toClientSession(authKeyId int64, row *model.Auths) *authsession.ClientSession {
	if row == nil {
		return nil
	}
	return authsession.MakeTLClientSession(&authsession.TLClientSession{
		AuthKeyId:      authKeyId,
		Ip:             row.ClientIp,
		Layer:          row.Layer,
		ApiId:          row.ApiId,
		DeviceModel:    row.DeviceModel,
		SystemVersion:  row.SystemVersion,
		AppVersion:     row.AppVersion,
		SystemLangCode: row.SystemLangCode,
		LangPack:       row.LangPack,
		LangCode:       row.LangCode,
		Proxy:          row.Proxy,
		Params:         row.Params,
	}).ToClientSession()
}

func toAuthKeyStateData(authKeyId int64, data *cacheAuthData) *authsession.AuthKeyStateData {
	stateData := authsession.MakeTLAuthKeyStateData(&authsession.TLAuthKeyStateData{
		AuthKeyId: authKeyId,
		KeyState:  tg.AuthStateNew,
	})
	if data == nil {
		return stateData.ToAuthKeyStateData()
	}

	stateData.KeyState = data.toAuthState()
	stateData.Client = data.Client
	if data.BindUser != nil {
		stateData.UserId = data.BindUser.UserId
		stateData.AccessHash = data.BindUser.Hash
		if data.BindUser.AndroidPushSessionId != 0 {
			stateData.AndroidPushSessionId = &data.BindUser.AndroidPushSessionId
		}
	}
	return stateData.ToAuthKeyStateData()
}

func authDataFromRows(authKeyId int64, authRow *model.Auths, userRow *model.AuthUsers) *cacheAuthData {
	cacheData := &cacheAuthData{}
	if authRow != nil {
		cacheData.Client = toClientSession(authKeyId, authRow)
	}
	if userRow != nil {
		cacheData.BindUser = &bindUser{
			UserId:               userRow.UserId,
			Hash:                 userRow.Hash,
			DateCreated:          userRow.DateCreated,
			DateActivated:        userRow.DateActive,
			AndroidPushSessionId: userRow.AndroidPushSessionId,
		}
	}
	return cacheData
}

func (r *Repository) GetAuthData(ctx context.Context, permAuthKeyId int64) (*cacheAuthData, error) {
	var data *cacheAuthData
	err := r.CachedConn.QueryRow(ctx, &data, authDataCacheKey(permAuthKeyId), func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
		authRow, err := r.model.AuthsModel.SelectByAuthKeyId(ctx, permAuthKeyId)
		if err != nil {
			if !isNotFound(err) {
				return err
			}
		}

		userRow, err := r.model.AuthUsersModel.Select(ctx, permAuthKeyId)
		if err != nil {
			if !isNotFound(err) {
				return err
			}
		}

		*v.(**cacheAuthData) = authDataFromRows(permAuthKeyId, authRow, userRow)
		return nil
	})
	if err != nil {
		return nil, wrapStorage(err)
	}
	return data, nil
}

func (r *Repository) GetAuthStateDataByAuthKeyId(ctx context.Context, authKeyId int64) (*authsession.AuthKeyStateData, error) {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return nil, err
	}
	return r.GetAuthStateData(ctx, authKeyId, permAuthKeyId)
}

func (r *Repository) GetAuthStateData(ctx context.Context, authKeyId int64, permAuthKeyId int64) (*authsession.AuthKeyStateData, error) {
	data, err := r.GetAuthData(ctx, permAuthKeyId)
	if err != nil {
		return nil, err
	}
	stateData := toAuthKeyStateData(authKeyId, data)
	if stateData.Client != nil {
		stateData.Client.AuthKeyId = authKeyId
	}
	return stateData, nil
}

func (r *Repository) GetApiLayerByAuthKeyId(ctx context.Context, authKeyId int64) (int32, error) {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return 0, err
	}
	return r.GetApiLayer(ctx, permAuthKeyId)
}

func (r *Repository) GetApiLayer(ctx context.Context, permAuthKeyId int64) (int32, error) {
	data, err := r.GetAuthData(ctx, permAuthKeyId)
	if err != nil || data == nil || data.Client == nil {
		return 0, err
	}
	return data.Client.Layer, nil
}

func (r *Repository) GetLangCodeByAuthKeyId(ctx context.Context, authKeyId int64) (string, error) {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return "", err
	}
	return r.GetLangCode(ctx, permAuthKeyId)
}

func (r *Repository) GetLangCode(ctx context.Context, permAuthKeyId int64) (string, error) {
	data, err := r.GetAuthData(ctx, permAuthKeyId)
	if err != nil || data == nil || data.Client == nil {
		return "en", err
	}
	return data.Client.LangCode, nil
}

func (r *Repository) GetLangPackByAuthKeyId(ctx context.Context, authKeyId int64) (string, error) {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return "", err
	}
	return r.GetLangPack(ctx, permAuthKeyId)
}

func (r *Repository) GetLangPack(ctx context.Context, permAuthKeyId int64) (string, error) {
	data, err := r.GetAuthData(ctx, permAuthKeyId)
	if err != nil || data == nil || data.Client == nil {
		return "", err
	}
	return normalizeLangPack(data.Client.LangPack, data.Client.AppVersion), nil
}

func (r *Repository) GetClientKindByAuthKeyId(ctx context.Context, authKeyId int64) (string, error) {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return "", err
	}
	return r.GetClientKind(ctx, permAuthKeyId)
}

func (r *Repository) GetClientKind(ctx context.Context, permAuthKeyId int64) (string, error) {
	data, err := r.GetAuthData(ctx, permAuthKeyId)
	if err != nil || data == nil || data.Client == nil {
		return "", err
	}
	clientKind := normalizeLangPack(data.Client.LangPack, data.Client.AppVersion)
	if clientKind == "android" && strings.Contains(data.Client.AppVersion, "TDLib") {
		clientKind = "react"
	}
	return clientKind, nil
}

func (r *Repository) GetAuthKeyUserIdByAuthKeyId(ctx context.Context, authKeyId int64) (int64, error) {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return 0, err
	}
	return r.GetAuthKeyUserId(ctx, permAuthKeyId)
}

func (r *Repository) GetAuthKeyUserId(ctx context.Context, permAuthKeyId int64) (int64, error) {
	data, err := r.GetAuthData(ctx, permAuthKeyId)
	if err != nil || data == nil || data.BindUser == nil {
		return 0, err
	}
	return data.BindUser.UserId, nil
}

func (r *Repository) GetAndroidPushSessionIdByAuthKeyId(ctx context.Context, authKeyId int64) (int64, error) {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return 0, err
	}
	return r.GetAndroidPushSessionId(ctx, permAuthKeyId)
}

func (r *Repository) GetAndroidPushSessionId(ctx context.Context, permAuthKeyId int64) (int64, error) {
	data, err := r.GetAuthData(ctx, permAuthKeyId)
	if err != nil || data == nil || data.BindUser == nil {
		return 0, err
	}
	return data.BindUser.AndroidPushSessionId, nil
}

func (r *Repository) BindAuthKeyUserByAuthKeyId(ctx context.Context, authKeyId int64, userId int64) (int64, error) {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return 0, err
	}
	return r.BindAuthKeyUser(ctx, permAuthKeyId, userId)
}

func (r *Repository) BindAuthKeyUser(ctx context.Context, permAuthKeyId int64, userId int64) (int64, error) {
	now := time.Now().Unix()
	hash, err := secureRandInt63()
	if err != nil {
		return 0, err
	}
	authUser := &model.AuthUsers{
		AuthKeyId:   permAuthKeyId,
		UserId:      userId,
		Hash:        hash,
		DateCreated: now,
		DateActive:  now,
	}

	_, _, err = r.CachedConn.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		return r.model.AuthUsersModel.InsertOrUpdates(ctx, authUser)
	}, authDataCacheKey(permAuthKeyId))
	if err != nil {
		return 0, wrapStorage(err)
	}
	return authUser.Hash, nil
}

func (r *Repository) UnbindAuthKeyUserByAuthKeyId(ctx context.Context, authKeyId int64, userId int64) error {
	permAuthKeyId := int64(0)
	if authKeyId != 0 {
		var err error
		permAuthKeyId, err = r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
		if err != nil {
			return err
		}
	}
	return r.UnbindAuthKeyUser(ctx, permAuthKeyId, userId)
}

func (r *Repository) UnbindAuthKeyUser(ctx context.Context, permAuthKeyId int64, userId int64) error {
	cacheKeys := []string{authDataCacheKey(permAuthKeyId)}
	if permAuthKeyId == 0 {
		cacheKeys = nil
		rows, err := r.model.AuthUsersModel.SelectAuthKeyIds(ctx, userId)
		if err != nil {
			return wrapStorage(err)
		}
		for i := range rows {
			cacheKeys = append(cacheKeys, authDataCacheKey(rows[i].AuthKeyId))
		}
	}

	_, _, err := r.CachedConn.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		if permAuthKeyId == 0 {
			rowsAffected, err := r.model.AuthUsersModel.DeleteUser(ctx, userId)
			return 0, rowsAffected, err
		}
		rowsAffected, err := r.model.AuthUsersModel.Delete(ctx, permAuthKeyId, userId)
		return 0, rowsAffected, err
	}, cacheKeys...)
	return wrapStorage(err)
}

func (r *Repository) SetAndroidPushSessionIdByAuthKeyId(ctx context.Context, userId int64, authKeyId int64, sessionId int64) error {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return err
	}
	return r.SetAndroidPushSessionId(ctx, userId, permAuthKeyId, sessionId)
}

func (r *Repository) SetAndroidPushSessionId(ctx context.Context, userId int64, permAuthKeyId int64, sessionId int64) error {
	_, _, err := r.CachedConn.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		rowsAffected, err := r.model.AuthUsersModel.UpdateAndroidPushSessionId(ctx, sessionId, permAuthKeyId, userId)
		return 0, rowsAffected, err
	}, authDataCacheKey(permAuthKeyId))
	return wrapStorage(err)
}

func normalizeLangPack(langPack, appVersion string) string {
	if langPack != "" {
		return langPack
	}
	if strings.HasSuffix(appVersion, " A") || strings.HasSuffix(appVersion, " Z") {
		return "weba"
	}
	return ""
}
