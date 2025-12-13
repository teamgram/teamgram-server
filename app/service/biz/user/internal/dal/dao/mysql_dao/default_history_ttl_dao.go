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

type DefaultHistoryTtlDAO struct {
	db *sqlx.DB
}

func NewDefaultHistoryTtlDAO(db *sqlx.DB) *DefaultHistoryTtlDAO {
	return &DefaultHistoryTtlDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into default_history_ttl(user_id, period) values (:user_id, :period) on duplicate key update period = values(period)
func (dao *DefaultHistoryTtlDAO) InsertOrUpdate(ctx context.Context, do *dataobject.DefaultHistoryTtlDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into default_history_ttl(user_id, period) values (:user_id, :period) on duplicate key update period = values(period)"
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
// insert into default_history_ttl(user_id, period) values (:user_id, :period) on duplicate key update period = values(period)
func (dao *DefaultHistoryTtlDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.DefaultHistoryTtlDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into default_history_ttl(user_id, period) values (:user_id, :period) on duplicate key update period = values(period)"
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

// Select
// select id, user_id, period from default_history_ttl where user_id = :user_id
func (dao *DefaultHistoryTtlDAO) Select(ctx context.Context, userId int64) (rValue *dataobject.DefaultHistoryTtlDO, err error) {
	var (
		query = "select id, user_id, period from default_history_ttl where user_id = ?"
		do    = &dataobject.DefaultHistoryTtlDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, userId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}
