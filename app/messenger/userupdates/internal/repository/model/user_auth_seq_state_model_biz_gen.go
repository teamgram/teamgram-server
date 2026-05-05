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

type bizUserAuthSeqStateModel interface {
	InsertIgnore(ctx context.Context, data *UserAuthSeqState) (lastInsertId, rowsAffected int64, err error)
	SelectByUserId(ctx context.Context, userId int64) (*UserAuthSeqState, error)
	SelectForUpdate(ctx context.Context, userId int64) (*UserAuthSeqState, error)
	UpdateSeqDate(ctx context.Context, seq int64, date int32, userId int64) (rowsAffected int64, err error)
}

type UserAuthSeqStateTxModel interface {
	InsertIgnore(data *UserAuthSeqState) (lastInsertId, rowsAffected int64, err error)
	SelectByUserId(userId int64) (*UserAuthSeqState, error)
	SelectForUpdate(userId int64) (*UserAuthSeqState, error)
	UpdateSeqDate(seq int64, date int32, userId int64) (rowsAffected int64, err error)
}

type defaultUserAuthSeqStateTxModel struct {
	tx *sqlx.Tx
}

func NewUserAuthSeqStateTxModel(tx *sqlx.Tx) UserAuthSeqStateTxModel {
	return &defaultUserAuthSeqStateTxModel{tx: tx}
}

// InsertIgnore
// insert ignore into user_auth_seq_state(user_id, seq, `date`, row_version) values (:user_id, :seq, :date, :row_version)
func (m *defaultUserAuthSeqStateModel) InsertIgnore(ctx context.Context, data *UserAuthSeqState) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into user_auth_seq_state(user_id, seq, `date`, row_version) values (:user_id, :seq, :date, :row_version)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_auth_seq_state.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_auth_seq_state.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_auth_seq_state.InsertIgnore rows affected: %w", err)
	}

	return

}

// InsertIgnore
// insert ignore into user_auth_seq_state(user_id, seq, `date`, row_version) values (:user_id, :seq, :date, :row_version)
func (m *defaultUserAuthSeqStateTxModel) InsertIgnore(data *UserAuthSeqState) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into user_auth_seq_state(user_id, seq, `date`, row_version) values (:user_id, :seq, :date, :row_version)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_auth_seq_state.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_auth_seq_state.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_auth_seq_state.InsertIgnore rows affected: %w", err)
	}

	return
}

// SelectByUserId
// select user_id, seq, `date`, row_version from user_auth_seq_state where user_id = :user_id limit 1
func (m *defaultUserAuthSeqStateModel) SelectByUserId(ctx context.Context, userId int64) (rValue *UserAuthSeqState, err error) {

	var (
		query = "select user_id, seq, `date`, row_version from user_auth_seq_state where user_id = ? limit 1"
		do    = &UserAuthSeqState{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_auth_seq_state",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_auth_seq_state.SelectByUserId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUserId
// select user_id, seq, `date`, row_version from user_auth_seq_state where user_id = :user_id limit 1
func (m *defaultUserAuthSeqStateTxModel) SelectByUserId(userId int64) (rValue *UserAuthSeqState, err error) {
	var (
		query = "select user_id, seq, `date`, row_version from user_auth_seq_state where user_id = ? limit 1"
		do    = &UserAuthSeqState{}
	)
	err = m.tx.QueryRowPartial(do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_auth_seq_state",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_auth_seq_state.SelectByUserId: %w", err)
		return
	}
	rValue = do

	return
}

// SelectForUpdate
// select user_id, seq, `date`, row_version from user_auth_seq_state where user_id = :user_id limit 1 for update
func (m *defaultUserAuthSeqStateModel) SelectForUpdate(ctx context.Context, userId int64) (rValue *UserAuthSeqState, err error) {

	var (
		query = "select user_id, seq, `date`, row_version from user_auth_seq_state where user_id = ? limit 1 for update"
		do    = &UserAuthSeqState{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_auth_seq_state",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_auth_seq_state.SelectForUpdate: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectForUpdate
// select user_id, seq, `date`, row_version from user_auth_seq_state where user_id = :user_id limit 1 for update
func (m *defaultUserAuthSeqStateTxModel) SelectForUpdate(userId int64) (rValue *UserAuthSeqState, err error) {
	var (
		query = "select user_id, seq, `date`, row_version from user_auth_seq_state where user_id = ? limit 1 for update"
		do    = &UserAuthSeqState{}
	)
	err = m.tx.QueryRowPartial(do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_auth_seq_state",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_auth_seq_state.SelectForUpdate: %w", err)
		return
	}
	rValue = do

	return
}

// UpdateSeqDate
// update user_auth_seq_state set seq = :seq, `date` = :date, row_version = row_version + 1 where user_id = :user_id
func (m *defaultUserAuthSeqStateModel) UpdateSeqDate(ctx context.Context, seq int64, date int32, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update user_auth_seq_state set seq = ?, `date` = ?, row_version = row_version + 1 where user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, seq, date, userId)

	if err != nil {
		err = fmt.Errorf("user_auth_seq_state.UpdateSeqDate exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_auth_seq_state.UpdateSeqDate rows affected: %w", err)
		return
	}

	return
}

// UpdateSeqDate
// update user_auth_seq_state set seq = :seq, `date` = :date, row_version = row_version + 1 where user_id = :user_id
func (m *defaultUserAuthSeqStateTxModel) UpdateSeqDate(seq int64, date int32, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_auth_seq_state set seq = ?, `date` = ?, row_version = row_version + 1 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, seq, date, userId)

	if err != nil {
		err = fmt.Errorf("user_auth_seq_state.UpdateSeqDate exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_auth_seq_state.UpdateSeqDate rows affected: %w", err)
		return
	}

	return
}
