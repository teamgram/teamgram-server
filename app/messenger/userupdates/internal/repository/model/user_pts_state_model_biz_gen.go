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

type bizUserPtsStateModel interface {
	InsertIgnore(ctx context.Context, data *UserPtsState) (lastInsertId, rowsAffected int64, err error)
	SelectByUserId(ctx context.Context, userId int64) (*UserPtsState, error)
	SelectForUpdate(ctx context.Context, userId int64) (*UserPtsState, error)
	UpdatePts(ctx context.Context, pts int64, ptsUpdatedAt string, partitionId int32, ownerEpoch int64, userId int64) (rowsAffected int64, err error)
}

type UserPtsStateTxModel interface {
	InsertIgnore(data *UserPtsState) (lastInsertId, rowsAffected int64, err error)
	SelectByUserId(userId int64) (*UserPtsState, error)
	SelectForUpdate(userId int64) (*UserPtsState, error)
	UpdatePts(pts int64, ptsUpdatedAt string, partitionId int32, ownerEpoch int64, userId int64) (rowsAffected int64, err error)
}

type defaultUserPtsStateTxModel struct {
	tx *sqlx.Tx
}

func NewUserPtsStateTxModel(tx *sqlx.Tx) UserPtsStateTxModel {
	return &defaultUserPtsStateTxModel{tx: tx}
}

// InsertIgnore
// insert ignore into user_pts_state(user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version) values (:user_id, :pts, :pts_updated_at, :partition_id, :owner_epoch, :row_version)
func (m *defaultUserPtsStateModel) InsertIgnore(ctx context.Context, data *UserPtsState) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into user_pts_state(user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version) values (:user_id, :pts, :pts_updated_at, :partition_id, :owner_epoch, :row_version)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_pts_state.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_pts_state.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_pts_state.InsertIgnore rows affected: %w", err)
	}

	return

}

// InsertIgnore
// insert ignore into user_pts_state(user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version) values (:user_id, :pts, :pts_updated_at, :partition_id, :owner_epoch, :row_version)
func (m *defaultUserPtsStateTxModel) InsertIgnore(data *UserPtsState) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into user_pts_state(user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version) values (:user_id, :pts, :pts_updated_at, :partition_id, :owner_epoch, :row_version)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_pts_state.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_pts_state.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_pts_state.InsertIgnore rows affected: %w", err)
	}

	return
}

// SelectByUserId
// select user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version from user_pts_state where user_id = :user_id limit 1
func (m *defaultUserPtsStateModel) SelectByUserId(ctx context.Context, userId int64) (rValue *UserPtsState, err error) {

	var (
		query = "select user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version from user_pts_state where user_id = ? limit 1"
		do    = &UserPtsState{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_state",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_pts_state.SelectByUserId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUserId
// select user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version from user_pts_state where user_id = :user_id limit 1
func (m *defaultUserPtsStateTxModel) SelectByUserId(userId int64) (rValue *UserPtsState, err error) {
	var (
		query = "select user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version from user_pts_state where user_id = ? limit 1"
		do    = &UserPtsState{}
	)
	err = m.tx.QueryRowPartial(do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_state",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_pts_state.SelectByUserId: %w", err)
		return
	}
	rValue = do

	return
}

// SelectForUpdate
// select user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version from user_pts_state where user_id = :user_id limit 1 for update
func (m *defaultUserPtsStateModel) SelectForUpdate(ctx context.Context, userId int64) (rValue *UserPtsState, err error) {

	var (
		query = "select user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version from user_pts_state where user_id = ? limit 1 for update"
		do    = &UserPtsState{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_state",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_pts_state.SelectForUpdate: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectForUpdate
// select user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version from user_pts_state where user_id = :user_id limit 1 for update
func (m *defaultUserPtsStateTxModel) SelectForUpdate(userId int64) (rValue *UserPtsState, err error) {
	var (
		query = "select user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version from user_pts_state where user_id = ? limit 1 for update"
		do    = &UserPtsState{}
	)
	err = m.tx.QueryRowPartial(do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_state",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_pts_state.SelectForUpdate: %w", err)
		return
	}
	rValue = do

	return
}

// UpdatePts
// update user_pts_state set pts = :pts, pts_updated_at = :pts_updated_at, partition_id = :partition_id, owner_epoch = :owner_epoch, row_version = row_version + 1 where user_id = :user_id
func (m *defaultUserPtsStateModel) UpdatePts(ctx context.Context, pts int64, ptsUpdatedAt string, partitionId int32, ownerEpoch int64, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update user_pts_state set pts = ?, pts_updated_at = ?, partition_id = ?, owner_epoch = ?, row_version = row_version + 1 where user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, pts, ptsUpdatedAt, partitionId, ownerEpoch, userId)

	if err != nil {
		err = fmt.Errorf("user_pts_state.UpdatePts exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_pts_state.UpdatePts rows affected: %w", err)
		return
	}

	return
}

// UpdatePts
// update user_pts_state set pts = :pts, pts_updated_at = :pts_updated_at, partition_id = :partition_id, owner_epoch = :owner_epoch, row_version = row_version + 1 where user_id = :user_id
func (m *defaultUserPtsStateTxModel) UpdatePts(pts int64, ptsUpdatedAt string, partitionId int32, ownerEpoch int64, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_pts_state set pts = ?, pts_updated_at = ?, partition_id = ?, owner_epoch = ?, row_version = row_version + 1 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, pts, ptsUpdatedAt, partitionId, ownerEpoch, userId)

	if err != nil {
		err = fmt.Errorf("user_pts_state.UpdatePts exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_pts_state.UpdatePts rows affected: %w", err)
		return
	}

	return
}
