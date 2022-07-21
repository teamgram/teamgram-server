/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/authsession/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type AuthKeyInfosDAO struct {
	db *sqlx.DB
}

func NewAuthKeyInfosDAO(db *sqlx.DB) *AuthKeyInfosDAO {
	return &AuthKeyInfosDAO{db}
}

// Insert
// insert into auth_key_infos(auth_key_id, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id) values (:auth_key_id, :auth_key_type, :perm_auth_key_id, :temp_auth_key_id, :media_temp_auth_key_id)
// TODO(@benqi): sqlmap
func (dao *AuthKeyInfosDAO) Insert(ctx context.Context, do *dataobject.AuthKeyInfosDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_key_infos(auth_key_id, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id) values (:auth_key_id, :auth_key_type, :perm_auth_key_id, :temp_auth_key_id, :media_temp_auth_key_id)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertTx
// insert into auth_key_infos(auth_key_id, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id) values (:auth_key_id, :auth_key_type, :perm_auth_key_id, :temp_auth_key_id, :media_temp_auth_key_id)
// TODO(@benqi): sqlmap
func (dao *AuthKeyInfosDAO) InsertTx(tx *sqlx.Tx, do *dataobject.AuthKeyInfosDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_key_infos(auth_key_id, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id) values (:auth_key_id, :auth_key_type, :perm_auth_key_id, :temp_auth_key_id, :media_temp_auth_key_id)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// SelectByAuthKeyId
// select auth_key_id, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id from auth_key_infos where auth_key_id = :auth_key_id limit 1
// TODO(@benqi): sqlmap
func (dao *AuthKeyInfosDAO) SelectByAuthKeyId(ctx context.Context, auth_key_id int64) (rValue *dataobject.AuthKeyInfosDO, err error) {
	var (
		query = "select auth_key_id, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id from auth_key_infos where auth_key_id = ? limit 1"
		do    = &dataobject.AuthKeyInfosDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, auth_key_id)

	if err != nil {
		if err != sqlx.ErrNotFound {
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

// UpdateCustomMap
// update auth_key_infos set %s where auth_key_id = :auth_key_id
// TODO(@benqi): sqlmap
func (dao *AuthKeyInfosDAO) UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, auth_key_id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update auth_key_infos set %s where auth_key_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, auth_key_id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateCustomMap(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateCustomMap(_), error: %v", err)
	}

	return
}

// UpdateCustomMapTx
// update auth_key_infos set %s where auth_key_id = :auth_key_id
// TODO(@benqi): sqlmap
func (dao *AuthKeyInfosDAO) UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, auth_key_id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update auth_key_infos set %s where auth_key_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, auth_key_id)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateCustomMap(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateCustomMap(_), error: %v", err)
	}

	return
}
