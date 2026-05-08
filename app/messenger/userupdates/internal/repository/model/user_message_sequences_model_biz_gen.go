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

type bizUserMessageSequencesModel interface {
	InsertIgnore(ctx context.Context, data *UserMessageSequences) (lastInsertId, rowsAffected int64, err error)
	SelectForUpdate(ctx context.Context, userId int64) (*UserMessageSequences, error)
	UpdateNext(ctx context.Context, nextUserMessageId int64, userId int64) (rowsAffected int64, err error)
}

type UserMessageSequencesTxModel interface {
	InsertIgnore(data *UserMessageSequences) (lastInsertId, rowsAffected int64, err error)
	SelectForUpdate(userId int64) (*UserMessageSequences, error)
	UpdateNext(nextUserMessageId int64, userId int64) (rowsAffected int64, err error)
}

type defaultUserMessageSequencesTxModel struct {
	tx *sqlx.Tx
}

func NewUserMessageSequencesTxModel(tx *sqlx.Tx) UserMessageSequencesTxModel {
	return &defaultUserMessageSequencesTxModel{tx: tx}
}

// InsertIgnore
// insert ignore into user_message_sequences(user_id, next_user_message_id) values (:user_id, :next_user_message_id)
func (m *defaultUserMessageSequencesModel) InsertIgnore(ctx context.Context, data *UserMessageSequences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into user_message_sequences(user_id, next_user_message_id) values (:user_id, :next_user_message_id)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_message_sequences.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_message_sequences.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_message_sequences.InsertIgnore rows affected: %w", err)
	}

	return

}

// InsertIgnore
// insert ignore into user_message_sequences(user_id, next_user_message_id) values (:user_id, :next_user_message_id)
func (m *defaultUserMessageSequencesTxModel) InsertIgnore(data *UserMessageSequences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into user_message_sequences(user_id, next_user_message_id) values (:user_id, :next_user_message_id)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_message_sequences.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_message_sequences.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_message_sequences.InsertIgnore rows affected: %w", err)
	}

	return
}

// SelectForUpdate
// select user_id, next_user_message_id from user_message_sequences where user_id = :user_id limit 1 for update
func (m *defaultUserMessageSequencesModel) SelectForUpdate(ctx context.Context, userId int64) (rValue *UserMessageSequences, err error) {

	var (
		query = "select user_id, next_user_message_id from user_message_sequences where user_id = ? limit 1 for update"
		do    = &UserMessageSequences{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_message_sequences",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_message_sequences.SelectForUpdate: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectForUpdate
// select user_id, next_user_message_id from user_message_sequences where user_id = :user_id limit 1 for update
func (m *defaultUserMessageSequencesTxModel) SelectForUpdate(userId int64) (rValue *UserMessageSequences, err error) {
	var (
		query = "select user_id, next_user_message_id from user_message_sequences where user_id = ? limit 1 for update"
		do    = &UserMessageSequences{}
	)
	err = m.tx.QueryRowPartial(do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_message_sequences",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_message_sequences.SelectForUpdate: %w", err)
		return
	}
	rValue = do

	return
}

// UpdateNext
// update user_message_sequences set next_user_message_id = :next_user_message_id where user_id = :user_id
func (m *defaultUserMessageSequencesModel) UpdateNext(ctx context.Context, nextUserMessageId int64, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update user_message_sequences set next_user_message_id = ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, nextUserMessageId, userId)

	if err != nil {
		err = fmt.Errorf("user_message_sequences.UpdateNext exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_message_sequences.UpdateNext rows affected: %w", err)
		return
	}

	return
}

// UpdateNext
// update user_message_sequences set next_user_message_id = :next_user_message_id where user_id = :user_id
func (m *defaultUserMessageSequencesTxModel) UpdateNext(nextUserMessageId int64, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_message_sequences set next_user_message_id = ? where user_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, nextUserMessageId, userId)

	if err != nil {
		err = fmt.Errorf("user_message_sequences.UpdateNext exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_message_sequences.UpdateNext rows affected: %w", err)
		return
	}

	return
}
