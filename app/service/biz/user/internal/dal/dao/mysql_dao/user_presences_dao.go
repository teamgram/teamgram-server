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

type UserPresencesDAO struct {
	db *sqlx.DB
}

func NewUserPresencesDAO(db *sqlx.DB) *UserPresencesDAO {
	return &UserPresencesDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into user_presences(user_id, last_seen_at, expires) values (:user_id, :last_seen_at, :expires) on duplicate key update last_seen_at = values(last_seen_at), expires = values(expires)
func (dao *UserPresencesDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserPresencesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_presences(user_id, last_seen_at, expires) values (:user_id, :last_seen_at, :expires) on duplicate key update last_seen_at = values(last_seen_at), expires = values(expires)"
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
// insert into user_presences(user_id, last_seen_at, expires) values (:user_id, :last_seen_at, :expires) on duplicate key update last_seen_at = values(last_seen_at), expires = values(expires)
func (dao *UserPresencesDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserPresencesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_presences(user_id, last_seen_at, expires) values (:user_id, :last_seen_at, :expires) on duplicate key update last_seen_at = values(last_seen_at), expires = values(expires)"
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
// select id, user_id, last_seen_at, expires from user_presences where user_id = :user_id
func (dao *UserPresencesDAO) Select(ctx context.Context, userId int64) (rValue *dataobject.UserPresencesDO, err error) {
	var (
		query = "select id, user_id, last_seen_at, expires from user_presences where user_id = ?"
		do    = &dataobject.UserPresencesDO{}
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

// SelectList
// select id, user_id, last_seen_at, expires from user_presences where user_id in (:idList)
func (dao *UserPresencesDAO) SelectList(ctx context.Context, idList []int64) (rList []dataobject.UserPresencesDO, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, last_seen_at, expires from user_presences where user_id in (%s)", sqlx.InInt64List(idList))
		values []dataobject.UserPresencesDO
	)
	if len(idList) == 0 {
		rList = []dataobject.UserPresencesDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, user_id, last_seen_at, expires from user_presences where user_id in (:idList)
func (dao *UserPresencesDAO) SelectListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *dataobject.UserPresencesDO)) (rList []dataobject.UserPresencesDO, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, last_seen_at, expires from user_presences where user_id in (%s)", sqlx.InInt64List(idList))
		values []dataobject.UserPresencesDO
	)
	if len(idList) == 0 {
		rList = []dataobject.UserPresencesDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

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
func (dao *UserPresencesDAO) UpdateLastSeenAt(ctx context.Context, lastSeenAt int64, expires int32, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_presences set last_seen_at = ?, expires = ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, lastSeenAt, expires, userId)

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
func (dao *UserPresencesDAO) UpdateLastSeenAtTx(tx *sqlx.Tx, lastSeenAt int64, expires int32, userId int64) (rowsAffected int64, err error) {
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
