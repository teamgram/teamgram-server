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

type bizUserPtsUpdatesModel interface {
	Insert(ctx context.Context, data *UserPtsUpdates) (lastInsertId, rowsAffected int64, err error)
	SelectLastPts(ctx context.Context, userId int64) (*UserPtsUpdates, error)
	SelectByGtPts(ctx context.Context, userId int64, pts int32, limit int32) ([]UserPtsUpdates, error)
	SelectByGtPtsWithCB(ctx context.Context, userId int64, pts int32, limit int32, cb func(sz, i int, v *UserPtsUpdates)) ([]UserPtsUpdates, error)
}

type UserPtsUpdatesTxModel interface {
	Insert(data *UserPtsUpdates) (lastInsertId, rowsAffected int64, err error)
	SelectLastPts(userId int64) (*UserPtsUpdates, error)
	SelectByGtPts(userId int64, pts int32, limit int32) ([]UserPtsUpdates, error)
}

type defaultUserPtsUpdatesTxModel struct {
	tx *sqlx.Tx
}

func NewUserPtsUpdatesTxModel(tx *sqlx.Tx) UserPtsUpdatesTxModel {
	return &defaultUserPtsUpdatesTxModel{tx: tx}
}

// Insert
// insert into user_pts_updates(user_id, pts, pts_count, update_type, update_data, date2) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date2)
func (m *defaultUserPtsUpdatesModel) Insert(ctx context.Context, data *UserPtsUpdates) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_pts_updates(user_id, pts, pts_count, update_type, update_data, date2) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_pts_updates.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_pts_updates.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_pts_updates.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into user_pts_updates(user_id, pts, pts_count, update_type, update_data, date2) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date2)
func (m *defaultUserPtsUpdatesTxModel) Insert(data *UserPtsUpdates) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_pts_updates(user_id, pts, pts_count, update_type, update_data, date2) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_pts_updates.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_pts_updates.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_pts_updates.Insert rows affected: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_updates",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_pts_updates.SelectLastPts: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectLastPts
// select pts from user_pts_updates where user_id = :user_id order by pts desc limit 1
func (m *defaultUserPtsUpdatesTxModel) SelectLastPts(userId int64) (rValue *UserPtsUpdates, err error) {
	var (
		query = "select pts from user_pts_updates where user_id = ? order by pts desc limit 1"
		do    = &UserPtsUpdates{}
	)
	err = m.tx.QueryRowPartial(do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_updates",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_pts_updates.SelectLastPts: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByGtPts
// select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = :user_id and pts > :pts order by pts asc limit :limit
func (m *defaultUserPtsUpdatesModel) SelectByGtPts(ctx context.Context, userId int64, pts int32, limit int32) (rList []UserPtsUpdates, err error) {
	var (
		query  = "select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = ? and pts > ? order by pts asc limit ?"
		values []UserPtsUpdates
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, pts, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPtsUpdates{}
			err = nil
			return
		}
		err = fmt.Errorf("user_pts_updates.SelectByGtPts: %w", err)
		return
	}

	rList = values

	return
}

// SelectByGtPts
// select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = :user_id and pts > :pts order by pts asc limit :limit
func (m *defaultUserPtsUpdatesTxModel) SelectByGtPts(userId int64, pts int32, limit int32) (rList []UserPtsUpdates, err error) {
	var (
		query  = "select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = ? and pts > ? order by pts asc limit ?"
		values []UserPtsUpdates
	)
	err = m.tx.QueryRowsPartial(&values, query, userId, pts, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPtsUpdates{}
			err = nil
			return
		}
		err = fmt.Errorf("user_pts_updates.SelectByGtPts: %w", err)
		return
	}

	rList = values

	return
}

// SelectByGtPtsWithCB
// select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = :user_id and pts > :pts order by pts asc limit :limit
func (m *defaultUserPtsUpdatesModel) SelectByGtPtsWithCB(ctx context.Context, userId int64, pts int32, limit int32, cb func(sz, i int, v *UserPtsUpdates)) (rList []UserPtsUpdates, err error) {
	var (
		query  = "select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = ? and pts > ? order by pts asc limit ?"
		values []UserPtsUpdates
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, pts, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPtsUpdates{}
			err = nil
			return
		}
		err = fmt.Errorf("user_pts_updates.SelectByGtPtsWithCB: %w", err)
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
