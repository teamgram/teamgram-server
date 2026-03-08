/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"errors"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/authsession/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthsDAO struct {
	db *sqlx.DB
}

func NewAuthsDAO(db *sqlx.DB) *AuthsDAO {
	return &AuthsDAO{
		db: db,
	}
}

// InsertOrUpdateLayer
// insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on duplicate key update layer = values(layer), client_ip = values(client_ip), date_active = values(date_active)
func (dao *AuthsDAO) InsertOrUpdateLayer(ctx context.Context, do *dataobject.AuthsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on duplicate key update layer = values(layer), client_ip = values(client_ip), date_active = values(date_active)"

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdateLayer(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdateLayer(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdateLayer(%v), error: %v", do, err)
	}

	return
}

// InsertOrUpdateLayerTx
// insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on duplicate key update layer = values(layer), client_ip = values(client_ip), date_active = values(date_active)
func (dao *AuthsDAO) InsertOrUpdateLayerTx(tx *sqlx.Tx, do *dataobject.AuthsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on duplicate key update layer = values(layer), client_ip = values(client_ip), date_active = values(date_active)"

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdateLayer(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdateLayer(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdateLayer(%v), error: %v", do, err)
	}

	return
}

// InsertOrUpdate
// insert into auths(auth_key_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update api_id = values(api_id), device_model = values(device_model), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)
func (dao *AuthsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.AuthsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into auths(auth_key_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update api_id = values(api_id), device_model = values(device_model), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)"

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v), error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into auths(auth_key_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update api_id = values(api_id), device_model = values(device_model), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)
func (dao *AuthsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.AuthsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into auths(auth_key_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update api_id = values(api_id), device_model = values(device_model), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)"

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v), error: %v", do, err)
	}

	return
}

// SelectByAuthKeyId
// select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id = :auth_key_id and deleted = 0 limit 1
func (dao *AuthsDAO) SelectByAuthKeyId(ctx context.Context, authKeyId int64) (rValue *dataobject.AuthsDO, err error) {
	var (
		query string
		do    = &dataobject.AuthsDO{}
	)
	query = "select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id = ? and deleted = 0 limit 1"

	err = dao.db.QueryRowPartial(ctx, do, query, authKeyId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByAuthKeyId(_), error: %v", err)
			return
		} else {
			// not found not error, return nil, nil
			err = nil
		}
	} else {
		rValue = do
	}

	return
}
