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
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *sqlx.Tx

type bizAuthsModel interface {
	InsertOrUpdateLayer(ctx context.Context, data *Auths) (lastInsertId, rowsAffected int64, err error)
	InsertOrUpdate(ctx context.Context, data *Auths) (lastInsertId, rowsAffected int64, err error)
	SelectByAuthKeyId(ctx context.Context, authKeyId int64) (*Auths, error)
}

type AuthsTxModel interface {
	InsertOrUpdateLayer(data *Auths) (lastInsertId, rowsAffected int64, err error)
	InsertOrUpdate(data *Auths) (lastInsertId, rowsAffected int64, err error)
	SelectByAuthKeyId(authKeyId int64) (*Auths, error)
}

type defaultAuthsTxModel struct {
	tx *sqlx.Tx
}

func NewAuthsTxModel(tx *sqlx.Tx) AuthsTxModel {
	return &defaultAuthsTxModel{tx: tx}
}

// InsertOrUpdateLayer
// insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on duplicate key update layer = values(layer), client_ip = values(client_ip), date_active = values(date_active)
func (m *defaultAuthsModel) InsertOrUpdateLayer(ctx context.Context, data *Auths) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on duplicate key update layer = values(layer), client_ip = values(client_ip), date_active = values(date_active)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("auths.InsertOrUpdateLayer named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auths.InsertOrUpdateLayer last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auths.InsertOrUpdateLayer rows affected: %w", err)
	}

	return

}

// InsertOrUpdateLayer
// insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on duplicate key update layer = values(layer), client_ip = values(client_ip), date_active = values(date_active)
func (m *defaultAuthsTxModel) InsertOrUpdateLayer(data *Auths) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on duplicate key update layer = values(layer), client_ip = values(client_ip), date_active = values(date_active)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("auths.InsertOrUpdateLayer named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auths.InsertOrUpdateLayer last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auths.InsertOrUpdateLayer rows affected: %w", err)
	}

	return
}

// InsertOrUpdate
// insert into auths(auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :layer, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update layer = values(layer), api_id = values(api_id), device_model = values(device_model), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)
func (m *defaultAuthsModel) InsertOrUpdate(ctx context.Context, data *Auths) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auths(auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :layer, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update layer = values(layer), api_id = values(api_id), device_model = values(device_model), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("auths.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auths.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auths.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into auths(auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :layer, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update layer = values(layer), api_id = values(api_id), device_model = values(device_model), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)
func (m *defaultAuthsTxModel) InsertOrUpdate(data *Auths) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auths(auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :layer, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update layer = values(layer), api_id = values(api_id), device_model = values(device_model), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("auths.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auths.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auths.InsertOrUpdate rows affected: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auths",
				Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auths.SelectByAuthKeyId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByAuthKeyId
// select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id = :auth_key_id and deleted = 0 limit 1
func (m *defaultAuthsTxModel) SelectByAuthKeyId(authKeyId int64) (rValue *Auths, err error) {
	var (
		query = "select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id = ? and deleted = 0 limit 1"
		do    = &Auths{}
	)
	err = m.tx.QueryRowPartial(do, query, authKeyId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auths",
				Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auths.SelectByAuthKeyId: %w", err)
		return
	}
	rValue = do

	return
}
