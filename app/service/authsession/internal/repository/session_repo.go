// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package repository

import (
	"context"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
)

// All write helpers in this file translate ClientSession / init-connection
// payloads keyed by the *caller* auth_key_id into rows keyed by the
// permanent auth_key_id. The conversion functions take permAuthKeyId
// explicitly so input structs are never mutated in place.

func authsFromClientSession(permAuthKeyId int64, session *authsession.ClientSession) *model.Auths {
	row := &model.Auths{
		AuthKeyId:      permAuthKeyId,
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

func authsFromInitConnection(permAuthKeyId int64, in *authsession.TLAuthsessionSetInitConnection) *model.Auths {
	row := &model.Auths{
		AuthKeyId:      permAuthKeyId,
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

// SetClientSessionInfoByAuthKeyId persists the client session metadata under
// the permanent auth_key_id. The provided ClientSession is treated as
// read-only — its AuthKeyId is not rewritten.
func (r *Repository) SetClientSessionInfoByAuthKeyId(ctx context.Context, session *authsession.ClientSession) error {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, session.AuthKeyId)
	if err != nil {
		return err
	}
	row := authsFromClientSession(permAuthKeyId, session)
	_, _, err = r.CachedConn.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		return r.model.AuthsModel.InsertOrUpdate(ctx, row)
	}, authDataCacheKey(permAuthKeyId))
	return wrapStorage(err)
}

// SetLayerByAuthKeyId records the protocol layer reported by the caller.
func (r *Repository) SetLayerByAuthKeyId(ctx context.Context, authKeyId int64, ip string, layer int32) error {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return err
	}
	row := &model.Auths{
		AuthKeyId:  permAuthKeyId,
		Layer:      layer,
		ClientIp:   ip,
		DateActive: time.Now().Unix(),
	}
	_, _, err = r.CachedConn.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		return r.model.AuthsModel.InsertOrUpdateLayer(ctx, row)
	}, authDataCacheKey(permAuthKeyId))
	return wrapStorage(err)
}

// SetInitConnectionByAuthKeyId persists init-connection metadata. The input
// payload is not mutated; only the resolved permanent key is written.
func (r *Repository) SetInitConnectionByAuthKeyId(ctx context.Context, in *authsession.TLAuthsessionSetInitConnection) error {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, in.AuthKeyId)
	if err != nil {
		return err
	}
	row := authsFromInitConnection(permAuthKeyId, in)
	_, _, err = r.CachedConn.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		return r.model.AuthsModel.InsertOrUpdate(ctx, row)
	}, authDataCacheKey(permAuthKeyId))
	return wrapStorage(err)
}
