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
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is

type UserSettingsDAO struct {
	db *sqlx.DB
}

func NewUserSettingsDAO(db *sqlx.DB) *UserSettingsDAO {
	return &UserSettingsDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into user_settings(user_id, key2, value) values (:user_id, :key2, :value) on duplicate key update value = values(value)
func (dao *UserSettingsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserSettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_settings(user_id, key2, value) values (:user_id, :key2, :value) on duplicate key update value = values(value)"
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
// insert into user_settings(user_id, key2, value) values (:user_id, :key2, :value) on duplicate key update value = values(value)
func (dao *UserSettingsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserSettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_settings(user_id, key2, value) values (:user_id, :key2, :value) on duplicate key update value = values(value)"
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

// SelectByKey
// select id, user_id, key2, value from user_settings where user_id = :user_id and key2 = :key2 and deleted = 0 limit 1
func (dao *UserSettingsDAO) SelectByKey(ctx context.Context, userId int64, key2 string) (rValue *dataobject.UserSettingsDO, err error) {
	var (
		query = "select id, user_id, key2, value from user_settings where user_id = ? and key2 = ? and deleted = 0 limit 1"
		do    = &dataobject.UserSettingsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, userId, key2)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByKey(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// Update
// update user_settings set value = :value, deleted = 0 where user_id = :user_id and key2 = :key2
func (dao *UserSettingsDAO) Update(ctx context.Context, value string, userId int64, key2 string) (rowsAffected int64, err error) {
	var (
		query   = "update user_settings set value = ?, deleted = 0 where user_id = ? and key2 = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, value, userId, key2)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

// UpdateTx
// update user_settings set value = :value, deleted = 0 where user_id = :user_id and key2 = :key2
func (dao *UserSettingsDAO) UpdateTx(tx *sqlx.Tx, value string, userId int64, key2 string) (rowsAffected int64, err error) {
	var (
		query   = "update user_settings set value = ?, deleted = 0 where user_id = ? and key2 = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, value, userId, key2)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}
