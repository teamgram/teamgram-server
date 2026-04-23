/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *logx.Logger

type (
	bizAuthsModel interface {
		InsertOrUpdateLayer(ctx context.Context, data *Auths) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateLayerTx(tx *sqlx.Tx, data *Auths) (lastInsertId, rowsAffected int64, err error)

		InsertOrUpdate(ctx context.Context, data *Auths) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *Auths) (lastInsertId, rowsAffected int64, err error)

		SelectByAuthKeyId(ctx context.Context, authKeyId int64) (*Auths, error)
	}
)

// InsertOrUpdateLayer
// insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on duplicate key update layer = values(layer), client_ip = values(client_ip), date_active = values(date_active)
func (m *defaultAuthsModel) InsertOrUpdateLayer(ctx context.Context, data *Auths) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on duplicate key update layer = values(layer), client_ip = values(client_ip), date_active = values(date_active)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdateLayer(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdateLayer(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdateLayer(%v)_error: %v", data, err)
	}

	return

}

// InsertOrUpdateLayerTx
// insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on duplicate key update layer = values(layer), client_ip = values(client_ip), date_active = values(date_active)
func (m *defaultAuthsModel) InsertOrUpdateLayerTx(tx *sqlx.Tx, data *Auths) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on duplicate key update layer = values(layer), client_ip = values(client_ip), date_active = values(date_active)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdateLayer(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdateLayer(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdateLayer(%v)_error: %v", data, err)
	}

	return
}

// InsertOrUpdate
// insert into auths(auth_key_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update api_id = values(api_id), device_model = values(device_model), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)
func (m *defaultAuthsModel) InsertOrUpdate(ctx context.Context, data *Auths) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auths(auth_key_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update api_id = values(api_id), device_model = values(device_model), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", data, err)
	}

	return

}

// InsertOrUpdateTx
// insert into auths(auth_key_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update api_id = values(api_id), device_model = values(device_model), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)
func (m *defaultAuthsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *Auths) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auths(auth_key_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update api_id = values(api_id), device_model = values(device_model), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", data, err)
	}

	return
}

// SelectByAuthKeyId
// select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id = :auth_key_id and deleted = 0 limit 1
func (m *defaultAuthsModel) SelectByAuthKeyId(ctx context.Context, authKeyId int64) (rValue *Auths, err error) {

	var (
		query = "select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id = ? and deleted = 0 limit 1"
		do    = &Auths{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, authKeyId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByAuthKeyId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}
