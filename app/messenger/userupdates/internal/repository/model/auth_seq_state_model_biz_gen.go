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

type bizAuthSeqStateModel interface {
	InsertIgnore(ctx context.Context, data *AuthSeqState) (lastInsertId, rowsAffected int64, err error)
	SelectByUserAuthKey(ctx context.Context, userId int64, permAuthKeyId int64) (*AuthSeqState, error)
	SelectForUpdate(ctx context.Context, userId int64, permAuthKeyId int64) (*AuthSeqState, error)
	UpdateSeqDate(ctx context.Context, seq int64, date int64, userId int64, permAuthKeyId int64) (rowsAffected int64, err error)
}

type AuthSeqStateTxModel interface {
	InsertIgnore(data *AuthSeqState) (lastInsertId, rowsAffected int64, err error)
	SelectByUserAuthKey(userId int64, permAuthKeyId int64) (*AuthSeqState, error)
	SelectForUpdate(userId int64, permAuthKeyId int64) (*AuthSeqState, error)
	UpdateSeqDate(seq int64, date int64, userId int64, permAuthKeyId int64) (rowsAffected int64, err error)
}

type defaultAuthSeqStateTxModel struct {
	tx *sqlx.Tx
}

func NewAuthSeqStateTxModel(tx *sqlx.Tx) AuthSeqStateTxModel {
	return &defaultAuthSeqStateTxModel{tx: tx}
}

// InsertIgnore
// insert ignore into auth_seq_state(user_id, perm_auth_key_id, seq, `date`, row_version) values (:user_id, :perm_auth_key_id, :seq, :date, :row_version)
func (m *defaultAuthSeqStateModel) InsertIgnore(ctx context.Context, data *AuthSeqState) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into auth_seq_state(user_id, perm_auth_key_id, seq, `date`, row_version) values (:user_id, :perm_auth_key_id, :seq, :date, :row_version)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("auth_seq_state.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_seq_state.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_seq_state.InsertIgnore rows affected: %w", err)
	}

	return

}

// InsertIgnore
// insert ignore into auth_seq_state(user_id, perm_auth_key_id, seq, `date`, row_version) values (:user_id, :perm_auth_key_id, :seq, :date, :row_version)
func (m *defaultAuthSeqStateTxModel) InsertIgnore(data *AuthSeqState) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into auth_seq_state(user_id, perm_auth_key_id, seq, `date`, row_version) values (:user_id, :perm_auth_key_id, :seq, :date, :row_version)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("auth_seq_state.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_seq_state.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_seq_state.InsertIgnore rows affected: %w", err)
	}

	return
}

// SelectByUserAuthKey
// select user_id, perm_auth_key_id, seq, `date`, row_version from auth_seq_state where user_id = :user_id and perm_auth_key_id = :perm_auth_key_id limit 1
func (m *defaultAuthSeqStateModel) SelectByUserAuthKey(ctx context.Context, userId int64, permAuthKeyId int64) (rValue *AuthSeqState, err error) {

	var (
		query = "select user_id, perm_auth_key_id, seq, `date`, row_version from auth_seq_state where user_id = ? and perm_auth_key_id = ? limit 1"
		do    = &AuthSeqState{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, permAuthKeyId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_seq_state",
				Key:      fmt.Sprintf("user_id=%v,perm_auth_key_id=%v", userId, permAuthKeyId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auth_seq_state.SelectByUserAuthKey: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUserAuthKey
// select user_id, perm_auth_key_id, seq, `date`, row_version from auth_seq_state where user_id = :user_id and perm_auth_key_id = :perm_auth_key_id limit 1
func (m *defaultAuthSeqStateTxModel) SelectByUserAuthKey(userId int64, permAuthKeyId int64) (rValue *AuthSeqState, err error) {
	var (
		query = "select user_id, perm_auth_key_id, seq, `date`, row_version from auth_seq_state where user_id = ? and perm_auth_key_id = ? limit 1"
		do    = &AuthSeqState{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, permAuthKeyId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_seq_state",
				Key:      fmt.Sprintf("user_id=%v,perm_auth_key_id=%v", userId, permAuthKeyId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auth_seq_state.SelectByUserAuthKey: %w", err)
		return
	}
	rValue = do

	return
}

// SelectForUpdate
// select user_id, perm_auth_key_id, seq, `date`, row_version from auth_seq_state where user_id = :user_id and perm_auth_key_id = :perm_auth_key_id limit 1 for update
func (m *defaultAuthSeqStateModel) SelectForUpdate(ctx context.Context, userId int64, permAuthKeyId int64) (rValue *AuthSeqState, err error) {

	var (
		query = "select user_id, perm_auth_key_id, seq, `date`, row_version from auth_seq_state where user_id = ? and perm_auth_key_id = ? limit 1 for update"
		do    = &AuthSeqState{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, permAuthKeyId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_seq_state",
				Key:      fmt.Sprintf("user_id=%v,perm_auth_key_id=%v", userId, permAuthKeyId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auth_seq_state.SelectForUpdate: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectForUpdate
// select user_id, perm_auth_key_id, seq, `date`, row_version from auth_seq_state where user_id = :user_id and perm_auth_key_id = :perm_auth_key_id limit 1 for update
func (m *defaultAuthSeqStateTxModel) SelectForUpdate(userId int64, permAuthKeyId int64) (rValue *AuthSeqState, err error) {
	var (
		query = "select user_id, perm_auth_key_id, seq, `date`, row_version from auth_seq_state where user_id = ? and perm_auth_key_id = ? limit 1 for update"
		do    = &AuthSeqState{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, permAuthKeyId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_seq_state",
				Key:      fmt.Sprintf("user_id=%v,perm_auth_key_id=%v", userId, permAuthKeyId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auth_seq_state.SelectForUpdate: %w", err)
		return
	}
	rValue = do

	return
}

// UpdateSeqDate
// update auth_seq_state set seq = :seq, `date` = :date, row_version = row_version + 1 where user_id = :user_id and perm_auth_key_id = :perm_auth_key_id
func (m *defaultAuthSeqStateModel) UpdateSeqDate(ctx context.Context, seq int64, date int64, userId int64, permAuthKeyId int64) (rowsAffected int64, err error) {

	var (
		query   = "update auth_seq_state set seq = ?, `date` = ?, row_version = row_version + 1 where user_id = ? and perm_auth_key_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, seq, date, userId, permAuthKeyId)

	if err != nil {
		err = fmt.Errorf("auth_seq_state.UpdateSeqDate exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_seq_state.UpdateSeqDate rows affected: %w", err)
		return
	}

	return
}

// UpdateSeqDate
// update auth_seq_state set seq = :seq, `date` = :date, row_version = row_version + 1 where user_id = :user_id and perm_auth_key_id = :perm_auth_key_id
func (m *defaultAuthSeqStateTxModel) UpdateSeqDate(seq int64, date int64, userId int64, permAuthKeyId int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_seq_state set seq = ?, `date` = ?, row_version = row_version + 1 where user_id = ? and perm_auth_key_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, seq, date, userId, permAuthKeyId)

	if err != nil {
		err = fmt.Errorf("auth_seq_state.UpdateSeqDate exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_seq_state.UpdateSeqDate rows affected: %w", err)
		return
	}

	return
}
