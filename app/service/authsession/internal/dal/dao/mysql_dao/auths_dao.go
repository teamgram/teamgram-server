/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/authsession/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type AuthsDAO struct {
	db *sqlx.DB
}

func NewAuthsDAO(db *sqlx.DB) *AuthsDAO {
	return &AuthsDAO{db}
}

// InsertOrUpdate
// insert into auths(auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :layer, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update layer = values(layer), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.AuthsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auths(auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :layer, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update layer = values(layer), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into auths(auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :layer, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update layer = values(layer), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.AuthsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auths(auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :layer, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on duplicate key update layer = values(layer), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip), date_active = values(date_active)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// SelectSessions
// select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) SelectSessions(ctx context.Context, idList []int64) (rList []dataobject.AuthsDO, err error) {
	var (
		query = "select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.AuthsDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectSessions(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectSessions(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.AuthsDO
	for rows.Next() {
		v := dataobject.AuthsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectSessions(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectSessionsWithCB
// select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) SelectSessionsWithCB(ctx context.Context, idList []int64, cb func(i int, v *dataobject.AuthsDO)) (rList []dataobject.AuthsDO, err error) {
	var (
		query = "select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.AuthsDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectSessions(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectSessions(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.AuthsDO
	for rows.Next() {
		v := dataobject.AuthsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectSessions(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByAuthKeyId
// select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id = :auth_key_id and deleted = 0 limit 1
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) SelectByAuthKeyId(ctx context.Context, auth_key_id int64) (rValue *dataobject.AuthsDO, err error) {
	var (
		query = "select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_key_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByAuthKeyId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.AuthsDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByAuthKeyId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectLayer
// select layer from auths where auth_key_id = :auth_key_id limit 1
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) SelectLayer(ctx context.Context, auth_key_id int64) (rValue int32, err error) {
	var query = "select layer from auths where auth_key_id = ? limit 1"
	err = dao.db.Get(ctx, &rValue, query, auth_key_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("get in SelectLayer(_), error: %v", err)
	}

	return
}

// SelectLangCode
// select lang_code from auths where auth_key_id = :auth_key_id limit 1
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) SelectLangCode(ctx context.Context, auth_key_id int64) (rValue string, err error) {
	var query = "select lang_code from auths where auth_key_id = ? limit 1"
	err = dao.db.Get(ctx, &rValue, query, auth_key_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("get in SelectLangCode(_), error: %v", err)
	}

	return
}

// SelectLangPack
// select lang_pack from auths where auth_key_id = :auth_key_id limit 1
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) SelectLangPack(ctx context.Context, auth_key_id int64) (rValue string, err error) {
	var query = "select lang_pack from auths where auth_key_id = ? limit 1"
	err = dao.db.Get(ctx, &rValue, query, auth_key_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("get in SelectLangPack(_), error: %v", err)
	}

	return
}
