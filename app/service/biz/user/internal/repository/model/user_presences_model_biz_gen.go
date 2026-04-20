/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026 The Teamgram Authors.
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
	bizUserPresencesModel interface {
		InsertOrUpdate(ctx context.Context, data *UserPresences) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *UserPresences) (lastInsertId, rowsAffected int64, err error)

		Select(ctx context.Context, userId int64) (*UserPresences, error)

		SelectList(ctx context.Context, idList []int64) ([]UserPresences, error)
		SelectListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *UserPresences)) ([]UserPresences, error)

		UpdateLastSeenAt(ctx context.Context, lastSeenAt int64, expires int32, userId int64) (rowsAffected int64, err error)
		UpdateLastSeenAtTx(tx *sqlx.Tx, lastSeenAt int64, expires int32, userId int64) (rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into user_presences(user_id, last_seen_at, expires) values (:user_id, :last_seen_at, :expires) on duplicate key update last_seen_at = values(last_seen_at), expires = values(expires)
func (m *defaultUserPresencesModel) InsertOrUpdate(ctx context.Context, data *UserPresences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_presences(user_id, last_seen_at, expires) values (:user_id, :last_seen_at, :expires) on duplicate key update last_seen_at = values(last_seen_at), expires = values(expires)"
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
// insert into user_presences(user_id, last_seen_at, expires) values (:user_id, :last_seen_at, :expires) on duplicate key update last_seen_at = values(last_seen_at), expires = values(expires)
func (m *defaultUserPresencesModel) InsertOrUpdateTx(tx *sqlx.Tx, data *UserPresences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_presences(user_id, last_seen_at, expires) values (:user_id, :last_seen_at, :expires) on duplicate key update last_seen_at = values(last_seen_at), expires = values(expires)"
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

// Select
// select id, user_id, last_seen_at, expires from user_presences where user_id = :user_id
func (m *defaultUserPresencesModel) Select(ctx context.Context, userId int64) (rValue *UserPresences, err error) {
	var (
		query = "select id, user_id, last_seen_at, expires from user_presences where user_id = ?"
		do    = &UserPresences{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId)

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

// SelectList
// select id, user_id, last_seen_at, expires from user_presences where user_id in (:idList)
func (m *defaultUserPresencesModel) SelectList(ctx context.Context, idList []int64) (rList []UserPresences, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, last_seen_at, expires from user_presences where user_id in (%s)", sqlx.InInt64List(idList))
		values []UserPresences
	)
	if len(idList) == 0 {
		rList = []UserPresences{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, user_id, last_seen_at, expires from user_presences where user_id in (:idList)
func (m *defaultUserPresencesModel) SelectListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *UserPresences)) (rList []UserPresences, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, last_seen_at, expires from user_presences where user_id in (%s)", sqlx.InInt64List(idList))
		values []UserPresences
	)
	if len(idList) == 0 {
		rList = []UserPresences{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// UpdateLastSeenAt
// update user_presences set last_seen_at = :last_seen_at, expires = :expires where user_id = :user_id
func (m *defaultUserPresencesModel) UpdateLastSeenAt(ctx context.Context, lastSeenAt int64, expires int32, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_presences set last_seen_at = ?, expires = ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, lastSeenAt, expires, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateLastSeenAt(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateLastSeenAt(_), error: %v", err)
	}

	return
}

// UpdateLastSeenAtTx
// update user_presences set last_seen_at = :last_seen_at, expires = :expires where user_id = :user_id
func (m *defaultUserPresencesModel) UpdateLastSeenAtTx(tx *sqlx.Tx, lastSeenAt int64, expires int32, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_presences set last_seen_at = ?, expires = ? where user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, lastSeenAt, expires, userId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateLastSeenAt(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateLastSeenAt(_), error: %v", err)
	}

	return
}
