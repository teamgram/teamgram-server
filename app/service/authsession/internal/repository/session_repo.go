package repository

import (
	"context"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
)

func authsFromClientSession(session *authsession.ClientSession) *model.Auths {
	row := &model.Auths{
		AuthKeyId:      session.AuthKeyId,
		Layer:          session.Layer,
		ApiId:          session.ApiId,
		DeviceModel:    session.DeviceModel,
		SystemVersion:  session.SystemVersion,
		AppVersion:     session.AppVersion,
		SystemLangCode: session.SystemLangCode,
		LangPack:       session.LangPack,
		LangCode:       session.LangCode,
		ClientIp:       session.Ip,
		Proxy:          session.Proxy,
		Params:         session.Params,
		DateActive:     time.Now().Unix(),
	}
	normalizeAuthsParams(row)
	return row
}

func authsFromInitConnection(in *authsession.TLAuthsessionSetInitConnection) *model.Auths {
	row := &model.Auths{
		AuthKeyId:      in.AuthKeyId,
		ApiId:          in.ApiId,
		DeviceModel:    in.DeviceModel,
		SystemVersion:  in.SystemVersion,
		AppVersion:     in.AppVersion,
		SystemLangCode: in.SystemLangCode,
		LangPack:       in.LangPack,
		LangCode:       in.LangCode,
		ClientIp:       in.Ip,
		Proxy:          in.Proxy,
		Params:         in.Params,
		DateActive:     time.Now().Unix(),
	}
	normalizeAuthsParams(row)
	return row
}

func normalizeAuthsParams(row *model.Auths) {
	if row.Params == "" {
		row.Params = "null"
	}
}

func (r *Repository) SetClientSessionInfo(ctx context.Context, session *authsession.ClientSession) error {
	row := authsFromClientSession(session)
	_, _, err := r.CachedConn.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		return r.model.AuthsModel.InsertOrUpdate(ctx, row)
	}, authDataCacheKey(row.AuthKeyId))
	return wrapStorage(err)
}

func (r *Repository) SetLayer(ctx context.Context, authKeyId int64, ip string, layer int32) error {
	row := &model.Auths{
		AuthKeyId:  authKeyId,
		Layer:      layer,
		ClientIp:   ip,
		DateActive: time.Now().Unix(),
	}
	_, _, err := r.CachedConn.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		return r.model.AuthsModel.InsertOrUpdateLayer(ctx, row)
	}, authDataCacheKey(authKeyId))
	return wrapStorage(err)
}

func (r *Repository) SetInitConnection(ctx context.Context, in *authsession.TLAuthsessionSetInitConnection) error {
	row := authsFromInitConnection(in)
	_, _, err := r.CachedConn.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		return r.model.AuthsModel.InsertOrUpdate(ctx, row)
	}, authDataCacheKey(row.AuthKeyId))
	return wrapStorage(err)
}
