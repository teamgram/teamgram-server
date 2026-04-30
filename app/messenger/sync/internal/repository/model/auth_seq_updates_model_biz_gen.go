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

type bizAuthSeqUpdatesModel interface {
	Insert(ctx context.Context, data *AuthSeqUpdates) (lastInsertId, rowsAffected int64, err error)
	SelectLastSeq(ctx context.Context, authId int64, userId int64) (*AuthSeqUpdates, error)
	SelectByGtSeq(ctx context.Context, authId int64, userId int64, seq int32) ([]AuthSeqUpdates, error)
	SelectByGtSeqWithCB(ctx context.Context, authId int64, userId int64, seq int32, cb func(sz, i int, v *AuthSeqUpdates)) ([]AuthSeqUpdates, error)
	SelectByGtDate(ctx context.Context, authId int64, userId int64, date2 int64) ([]AuthSeqUpdates, error)
	SelectByGtDateWithCB(ctx context.Context, authId int64, userId int64, date2 int64, cb func(sz, i int, v *AuthSeqUpdates)) ([]AuthSeqUpdates, error)
}

type AuthSeqUpdatesTxModel interface {
	Insert(data *AuthSeqUpdates) (lastInsertId, rowsAffected int64, err error)
	SelectLastSeq(authId int64, userId int64) (*AuthSeqUpdates, error)
	SelectByGtSeq(authId int64, userId int64, seq int32) ([]AuthSeqUpdates, error)
	SelectByGtDate(authId int64, userId int64, date2 int64) ([]AuthSeqUpdates, error)
}

type defaultAuthSeqUpdatesTxModel struct {
	tx *sqlx.Tx
}

func NewAuthSeqUpdatesTxModel(tx *sqlx.Tx) AuthSeqUpdatesTxModel {
	return &defaultAuthSeqUpdatesTxModel{tx: tx}
}

// Insert
// insert into auth_seq_updates(auth_id, user_id, seq, update_type, update_data, date2) values (:auth_id, :user_id, :seq, :update_type, :update_data, :date2)
func (m *defaultAuthSeqUpdatesModel) Insert(ctx context.Context, data *AuthSeqUpdates) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_seq_updates(auth_id, user_id, seq, update_type, update_data, date2) values (:auth_id, :user_id, :seq, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("auth_seq_updates.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_seq_updates.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_seq_updates.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into auth_seq_updates(auth_id, user_id, seq, update_type, update_data, date2) values (:auth_id, :user_id, :seq, :update_type, :update_data, :date2)
func (m *defaultAuthSeqUpdatesTxModel) Insert(data *AuthSeqUpdates) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_seq_updates(auth_id, user_id, seq, update_type, update_data, date2) values (:auth_id, :user_id, :seq, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("auth_seq_updates.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_seq_updates.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_seq_updates.Insert rows affected: %w", err)
	}

	return
}

// SelectLastSeq
// select seq from auth_seq_updates where auth_id = :auth_id and user_id = :user_id order by seq desc limit 1
func (m *defaultAuthSeqUpdatesModel) SelectLastSeq(ctx context.Context, authId int64, userId int64) (rValue *AuthSeqUpdates, err error) {

	var (
		query = "select seq from auth_seq_updates where auth_id = ? and user_id = ? order by seq desc limit 1"
		do    = &AuthSeqUpdates{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, authId, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_seq_updates",
				Key:      fmt.Sprintf("auth_id=%v,user_id=%v", authId, userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auth_seq_updates.SelectLastSeq: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectLastSeq
// select seq from auth_seq_updates where auth_id = :auth_id and user_id = :user_id order by seq desc limit 1
func (m *defaultAuthSeqUpdatesTxModel) SelectLastSeq(authId int64, userId int64) (rValue *AuthSeqUpdates, err error) {
	var (
		query = "select seq from auth_seq_updates where auth_id = ? and user_id = ? order by seq desc limit 1"
		do    = &AuthSeqUpdates{}
	)
	err = m.tx.QueryRowPartial(do, query, authId, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_seq_updates",
				Key:      fmt.Sprintf("auth_id=%v,user_id=%v", authId, userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auth_seq_updates.SelectLastSeq: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByGtSeq
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and seq > :seq order by seq asc
func (m *defaultAuthSeqUpdatesModel) SelectByGtSeq(ctx context.Context, authId int64, userId int64, seq int32) (rList []AuthSeqUpdates, err error) {
	var (
		query  = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and seq > ? order by seq asc"
		values []AuthSeqUpdates
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, authId, userId, seq)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthSeqUpdates{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_seq_updates.SelectByGtSeq: %w", err)
		return
	}

	rList = values

	return
}

// SelectByGtSeq
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and seq > :seq order by seq asc
func (m *defaultAuthSeqUpdatesTxModel) SelectByGtSeq(authId int64, userId int64, seq int32) (rList []AuthSeqUpdates, err error) {
	var (
		query  = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and seq > ? order by seq asc"
		values []AuthSeqUpdates
	)
	err = m.tx.QueryRowsPartial(&values, query, authId, userId, seq)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthSeqUpdates{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_seq_updates.SelectByGtSeq: %w", err)
		return
	}

	rList = values

	return
}

// SelectByGtSeqWithCB
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and seq > :seq order by seq asc
func (m *defaultAuthSeqUpdatesModel) SelectByGtSeqWithCB(ctx context.Context, authId int64, userId int64, seq int32, cb func(sz, i int, v *AuthSeqUpdates)) (rList []AuthSeqUpdates, err error) {
	var (
		query  = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and seq > ? order by seq asc"
		values []AuthSeqUpdates
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, authId, userId, seq)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthSeqUpdates{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_seq_updates.SelectByGtSeqWithCB: %w", err)
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

// SelectByGtDate
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and date2 > :date2 order by seq asc
func (m *defaultAuthSeqUpdatesModel) SelectByGtDate(ctx context.Context, authId int64, userId int64, date2 int64) (rList []AuthSeqUpdates, err error) {
	var (
		query  = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and date2 > ? order by seq asc"
		values []AuthSeqUpdates
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, authId, userId, date2)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthSeqUpdates{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_seq_updates.SelectByGtDate: %w", err)
		return
	}

	rList = values

	return
}

// SelectByGtDate
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and date2 > :date2 order by seq asc
func (m *defaultAuthSeqUpdatesTxModel) SelectByGtDate(authId int64, userId int64, date2 int64) (rList []AuthSeqUpdates, err error) {
	var (
		query  = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and date2 > ? order by seq asc"
		values []AuthSeqUpdates
	)
	err = m.tx.QueryRowsPartial(&values, query, authId, userId, date2)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthSeqUpdates{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_seq_updates.SelectByGtDate: %w", err)
		return
	}

	rList = values

	return
}

// SelectByGtDateWithCB
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and date2 > :date2 order by seq asc
func (m *defaultAuthSeqUpdatesModel) SelectByGtDateWithCB(ctx context.Context, authId int64, userId int64, date2 int64, cb func(sz, i int, v *AuthSeqUpdates)) (rList []AuthSeqUpdates, err error) {
	var (
		query  = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and date2 > ? order by seq asc"
		values []AuthSeqUpdates
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, authId, userId, date2)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthSeqUpdates{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_seq_updates.SelectByGtDateWithCB: %w", err)
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
