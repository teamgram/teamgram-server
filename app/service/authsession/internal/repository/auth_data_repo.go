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

type clientSessionCacheData struct {
	AuthKeyId      int64  `json:"auth_key_id"`
	Ip             string `json:"ip"`
	Layer          int32  `json:"layer"`
	ApiId          int32  `json:"api_id"`
	DeviceModel    string `json:"device_model"`
	SystemVersion  string `json:"system_version"`
	AppVersion     string `json:"app_version"`
	SystemLangCode string `json:"system_lang_code"`
	LangPack       string `json:"lang_pack"`
	LangCode       string `json:"lang_code"`
	Proxy          string `json:"proxy"`
	Params         string `json:"params"`
}

type cacheAuthData struct {
	Client   *clientSessionCacheData `json:"client"`
	BindUser *bindUser               `json:"bind_user,omitempty"`
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

func toClientSessionCacheData(authKeyId int64, row *model.Auths) *clientSessionCacheData {
	if row == nil {
		return nil
	}
	return &clientSessionCacheData{
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
	}
}

func toClientSession(authKeyId int64, row *model.Auths) *authsession.ClientSession {
	return toClientSessionCacheData(authKeyId, row).toClientSession()
}

func (c *clientSessionCacheData) toClientSession() *authsession.ClientSession {
	if c == nil {
		return nil
	}
	return authsession.MakeTLClientSession(&authsession.TLClientSession{
		AuthKeyId:      c.AuthKeyId,
		Ip:             c.Ip,
		Layer:          c.Layer,
		ApiId:          c.ApiId,
		DeviceModel:    c.DeviceModel,
		SystemVersion:  c.SystemVersion,
		AppVersion:     c.AppVersion,
		SystemLangCode: c.SystemLangCode,
		LangPack:       c.LangPack,
		LangCode:       c.LangCode,
		Proxy:          c.Proxy,
		Params:         c.Params,
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
	stateData.Client = data.Client.toClientSession()
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
		cacheData.Client = toClientSessionCacheData(authKeyId, authRow)
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

// GetAuthData loads the cached aggregate (auth row + bound user row) for the
// given permanent auth_key_id. It is the single read-through path used by the
// *ByAuthKeyId helpers in this file; downstream code should reach for those
// methods rather than passing perm ids around.
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

// authDataByAuthKeyId resolves the permanent auth_key_id for the caller's
// auth_key_id and returns the joined cache record. Read helpers below all
// route through it instead of duplicating the resolve+load boilerplate.
func (r *Repository) authDataByAuthKeyId(ctx context.Context, authKeyId int64) (permAuthKeyId int64, data *cacheAuthData, err error) {
	permAuthKeyId, err = r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return 0, nil, err
	}
	data, err = r.GetAuthData(ctx, permAuthKeyId)
	return permAuthKeyId, data, err
}

func (r *Repository) GetAuthStateDataByAuthKeyId(ctx context.Context, authKeyId int64) (*authsession.AuthKeyStateData, error) {
	_, data, err := r.authDataByAuthKeyId(ctx, authKeyId)
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
	_, data, err := r.authDataByAuthKeyId(ctx, authKeyId)
	if err != nil || data == nil || data.Client == nil {
		return 0, err
	}
	return data.Client.Layer, nil
}

// GetLangCodeByAuthKeyId returns the negotiated lang_code or an empty string
// when no client metadata is on record. The caller is responsible for
// applying any defaulting (e.g. "en") so storage failures stay observable.
func (r *Repository) GetLangCodeByAuthKeyId(ctx context.Context, authKeyId int64) (string, error) {
	_, data, err := r.authDataByAuthKeyId(ctx, authKeyId)
	if err != nil || data == nil || data.Client == nil {
		return "", err
	}
	return data.Client.LangCode, nil
}

func (r *Repository) GetLangPackByAuthKeyId(ctx context.Context, authKeyId int64) (string, error) {
	_, data, err := r.authDataByAuthKeyId(ctx, authKeyId)
	if err != nil || data == nil || data.Client == nil {
		return "", err
	}
	return normalizeLangPack(data.Client.LangPack, data.Client.AppVersion), nil
}

func (r *Repository) GetClientKindByAuthKeyId(ctx context.Context, authKeyId int64) (string, error) {
	_, data, err := r.authDataByAuthKeyId(ctx, authKeyId)
	if err != nil || data == nil || data.Client == nil {
		return "", err
	}
	clientKind := normalizeLangPack(data.Client.LangPack, data.Client.AppVersion)
	// "android" + a TDLib build string is how Telegram identifies the
	// React-based client; preserve that legacy mapping until callers move
	// to a richer client-kind enum.
	if clientKind == "android" && strings.Contains(data.Client.AppVersion, "TDLib") {
		clientKind = "react"
	}
	return clientKind, nil
}

func (r *Repository) GetAuthKeyUserIdByAuthKeyId(ctx context.Context, authKeyId int64) (int64, error) {
	_, data, err := r.authDataByAuthKeyId(ctx, authKeyId)
	if err != nil || data == nil || data.BindUser == nil {
		return 0, err
	}
	return data.BindUser.UserId, nil
}

func (r *Repository) GetAndroidPushSessionIdByAuthKeyId(ctx context.Context, authKeyId int64) (int64, error) {
	_, data, err := r.authDataByAuthKeyId(ctx, authKeyId)
	if err != nil || data == nil || data.BindUser == nil {
		return 0, err
	}
	return data.BindUser.AndroidPushSessionId, nil
}

func (r *Repository) GetPermAuthKeyIdsByUserId(ctx context.Context, userId int64) ([]int64, error) {
	rows, err := r.model.AuthUsersModel.SelectListByUserId(ctx, userId)
	if err != nil {
		return nil, wrapStorage(err)
	}
	authKeyIds := make([]int64, 0, len(rows))
	for i := range rows {
		authKeyIds = append(authKeyIds, rows[i].AuthKeyId)
	}
	return authKeyIds, nil
}

// BindAuthKeyUserByAuthKeyId binds the caller's auth_key to a user and
// returns the freshly minted access hash.
func (r *Repository) BindAuthKeyUserByAuthKeyId(ctx context.Context, authKeyId int64, userId int64) (int64, error) {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return 0, err
	}

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

// UnbindAuthKeyUserByAuthKeyId removes the binding for the given user. When
// authKeyId is 0 every binding owned by the user is deleted; callers use that
// path during full account-wide logout.
func (r *Repository) UnbindAuthKeyUserByAuthKeyId(ctx context.Context, authKeyId int64, userId int64) error {
	permAuthKeyId := int64(0)
	if authKeyId != 0 {
		var err error
		permAuthKeyId, err = r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
		if err != nil {
			return err
		}
	}

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
	_, rowsAffected, err := r.CachedConn.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		rowsAffected, err := r.model.AuthUsersModel.UpdateAndroidPushSessionId(ctx, sessionId, permAuthKeyId, userId)
		return 0, rowsAffected, err
	}, authDataCacheKey(permAuthKeyId))
	if err != nil {
		if isNotFound(err) {
			return authsession.ErrAuthorizationNotFound
		}
		return wrapStorage(err)
	}
	if rowsAffected == 0 {
		return authsession.ErrAuthorizationNotFound
	}
	return nil
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
