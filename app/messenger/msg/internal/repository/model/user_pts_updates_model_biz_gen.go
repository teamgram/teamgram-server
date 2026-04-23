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
	bizUserPtsUpdatesModel interface {
		Insert(ctx context.Context, data *UserPtsUpdates) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *UserPtsUpdates) (lastInsertId, rowsAffected int64, err error)

		SelectLastPts(ctx context.Context, userId int64) (*UserPtsUpdates, error)

		SelectByGtPts(ctx context.Context, userId int64, pts int32) ([]UserPtsUpdates, error)
		SelectByGtPtsWithCB(ctx context.Context, userId int64, pts int32, cb func(sz, i int, v *UserPtsUpdates)) ([]UserPtsUpdates, error)
	}
)

// Insert
// insert into user_pts_updates(user_id, pts, pts_count, update_type, update_data, date2) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date2)
func (m *defaultUserPtsUpdatesModel) Insert(ctx context.Context, data *UserPtsUpdates) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_pts_updates(user_id, pts, pts_count, update_type, update_data, date2) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
	}

	return

}

// InsertTx
// insert into user_pts_updates(user_id, pts, pts_count, update_type, update_data, date2) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date2)
func (m *defaultUserPtsUpdatesModel) InsertTx(tx *sqlx.Tx, data *UserPtsUpdates) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_pts_updates(user_id, pts, pts_count, update_type, update_data, date2) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
	}

	return
}

// SelectLastPts
// select pts from user_pts_updates where user_id = :user_id order by pts desc limit 1
func (m *defaultUserPtsUpdatesModel) SelectLastPts(ctx context.Context, userId int64) (rValue *UserPtsUpdates, err error) {

	var (
		query = "select pts from user_pts_updates where user_id = ? order by pts desc limit 1"
		do    = &UserPtsUpdates{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectLastPts(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectByGtPts
// select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = :user_id and pts > :pts order by pts asc
func (m *defaultUserPtsUpdatesModel) SelectByGtPts(ctx context.Context, userId int64, pts int32) (rList []UserPtsUpdates, err error) {
	var (
		query  = "select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = ? and pts > ? order by pts asc"
		values []UserPtsUpdates
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, pts)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtPts(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByGtPtsWithCB
// select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = :user_id and pts > :pts order by pts asc
func (m *defaultUserPtsUpdatesModel) SelectByGtPtsWithCB(ctx context.Context, userId int64, pts int32, cb func(sz, i int, v *UserPtsUpdates)) (rList []UserPtsUpdates, err error) {
	var (
		query  = "select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = ? and pts > ? order by pts asc"
		values []UserPtsUpdates
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, pts)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtPts(_), error: %v", err)
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
