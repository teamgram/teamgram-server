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

type bizAuthSeqDeliveriesModel interface {
	InsertIgnore(ctx context.Context, data *AuthSeqDeliveries) (lastInsertId, rowsAffected int64, err error)
	SelectByOperation(ctx context.Context, userId int64, permAuthKeyId int64, operationId string) (*AuthSeqDeliveries, error)
	SelectReplayableAfterDate(ctx context.Context, userId int64, permAuthKeyId int64, date int64, now int64, limit int32) ([]AuthSeqDeliveries, error)
	SelectReplayableAfterDateWithCB(ctx context.Context, userId int64, permAuthKeyId int64, date int64, now int64, limit int32, cb func(sz, i int, v *AuthSeqDeliveries)) ([]AuthSeqDeliveries, error)
}

type AuthSeqDeliveriesTxModel interface {
	InsertIgnore(data *AuthSeqDeliveries) (lastInsertId, rowsAffected int64, err error)
	SelectByOperation(userId int64, permAuthKeyId int64, operationId string) (*AuthSeqDeliveries, error)
	SelectReplayableAfterDate(userId int64, permAuthKeyId int64, date int64, now int64, limit int32) ([]AuthSeqDeliveries, error)
}

type defaultAuthSeqDeliveriesTxModel struct {
	tx *sqlx.Tx
}

func NewAuthSeqDeliveriesTxModel(tx *sqlx.Tx) AuthSeqDeliveriesTxModel {
	return &defaultAuthSeqDeliveriesTxModel{tx: tx}
}

// InsertIgnore
// insert ignore into auth_seq_deliveries(user_id, perm_auth_key_id, seq, `date`, payload_id, replay_policy, source_perm_auth_key_id, visibility_policy, operation_id, expire_at) values (:user_id, :perm_auth_key_id, :seq, :date, :payload_id, :replay_policy, :source_perm_auth_key_id, :visibility_policy, :operation_id, :expire_at)
func (m *defaultAuthSeqDeliveriesModel) InsertIgnore(ctx context.Context, data *AuthSeqDeliveries) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into auth_seq_deliveries(user_id, perm_auth_key_id, seq, `date`, payload_id, replay_policy, source_perm_auth_key_id, visibility_policy, operation_id, expire_at) values (:user_id, :perm_auth_key_id, :seq, :date, :payload_id, :replay_policy, :source_perm_auth_key_id, :visibility_policy, :operation_id, :expire_at)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("auth_seq_deliveries.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_seq_deliveries.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_seq_deliveries.InsertIgnore rows affected: %w", err)
	}

	return

}

// InsertIgnore
// insert ignore into auth_seq_deliveries(user_id, perm_auth_key_id, seq, `date`, payload_id, replay_policy, source_perm_auth_key_id, visibility_policy, operation_id, expire_at) values (:user_id, :perm_auth_key_id, :seq, :date, :payload_id, :replay_policy, :source_perm_auth_key_id, :visibility_policy, :operation_id, :expire_at)
func (m *defaultAuthSeqDeliveriesTxModel) InsertIgnore(data *AuthSeqDeliveries) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into auth_seq_deliveries(user_id, perm_auth_key_id, seq, `date`, payload_id, replay_policy, source_perm_auth_key_id, visibility_policy, operation_id, expire_at) values (:user_id, :perm_auth_key_id, :seq, :date, :payload_id, :replay_policy, :source_perm_auth_key_id, :visibility_policy, :operation_id, :expire_at)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("auth_seq_deliveries.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_seq_deliveries.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_seq_deliveries.InsertIgnore rows affected: %w", err)
	}

	return
}

// SelectByOperation
// select user_id, perm_auth_key_id, seq, `date`, payload_id, replay_policy, source_perm_auth_key_id, visibility_policy, operation_id, expire_at from auth_seq_deliveries where user_id = :user_id and perm_auth_key_id = :perm_auth_key_id and operation_id = :operation_id limit 1
func (m *defaultAuthSeqDeliveriesModel) SelectByOperation(ctx context.Context, userId int64, permAuthKeyId int64, operationId string) (rValue *AuthSeqDeliveries, err error) {

	var (
		query = "select user_id, perm_auth_key_id, seq, `date`, payload_id, replay_policy, source_perm_auth_key_id, visibility_policy, operation_id, expire_at from auth_seq_deliveries where user_id = ? and perm_auth_key_id = ? and operation_id = ? limit 1"
		do    = &AuthSeqDeliveries{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, permAuthKeyId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_seq_deliveries",
				Key:      fmt.Sprintf("user_id=%v,perm_auth_key_id=%v,operation_id=%v", userId, permAuthKeyId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auth_seq_deliveries.SelectByOperation: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByOperation
// select user_id, perm_auth_key_id, seq, `date`, payload_id, replay_policy, source_perm_auth_key_id, visibility_policy, operation_id, expire_at from auth_seq_deliveries where user_id = :user_id and perm_auth_key_id = :perm_auth_key_id and operation_id = :operation_id limit 1
func (m *defaultAuthSeqDeliveriesTxModel) SelectByOperation(userId int64, permAuthKeyId int64, operationId string) (rValue *AuthSeqDeliveries, err error) {
	var (
		query = "select user_id, perm_auth_key_id, seq, `date`, payload_id, replay_policy, source_perm_auth_key_id, visibility_policy, operation_id, expire_at from auth_seq_deliveries where user_id = ? and perm_auth_key_id = ? and operation_id = ? limit 1"
		do    = &AuthSeqDeliveries{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, permAuthKeyId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_seq_deliveries",
				Key:      fmt.Sprintf("user_id=%v,perm_auth_key_id=%v,operation_id=%v", userId, permAuthKeyId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auth_seq_deliveries.SelectByOperation: %w", err)
		return
	}
	rValue = do

	return
}

// SelectReplayableAfterDate
// select d.user_id, d.perm_auth_key_id, d.seq, d.`date`, d.payload_id, d.replay_policy, d.source_perm_auth_key_id, d.visibility_policy, d.operation_id, d.expire_at from auth_seq_deliveries as d where d.user_id = :user_id and d.perm_auth_key_id = :perm_auth_key_id and d.`date` > :date and (d.expire_at = 0 or d.expire_at > :now) order by d.`date` asc, d.seq asc limit :limit
func (m *defaultAuthSeqDeliveriesModel) SelectReplayableAfterDate(ctx context.Context, userId int64, permAuthKeyId int64, date int64, now int64, limit int32) (rList []AuthSeqDeliveries, err error) {
	var (
		query  = "select d.user_id, d.perm_auth_key_id, d.seq, d.`date`, d.payload_id, d.replay_policy, d.source_perm_auth_key_id, d.visibility_policy, d.operation_id, d.expire_at from auth_seq_deliveries as d where d.user_id = ? and d.perm_auth_key_id = ? and d.`date` > ? and (d.expire_at = 0 or d.expire_at > ?) order by d.`date` asc, d.seq asc limit ?"
		values []AuthSeqDeliveries
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, permAuthKeyId, date, now, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthSeqDeliveries{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_seq_deliveries.SelectReplayableAfterDate: %w", err)
		return
	}

	rList = values

	return
}

// SelectReplayableAfterDate
// select d.user_id, d.perm_auth_key_id, d.seq, d.`date`, d.payload_id, d.replay_policy, d.source_perm_auth_key_id, d.visibility_policy, d.operation_id, d.expire_at from auth_seq_deliveries as d where d.user_id = :user_id and d.perm_auth_key_id = :perm_auth_key_id and d.`date` > :date and (d.expire_at = 0 or d.expire_at > :now) order by d.`date` asc, d.seq asc limit :limit
func (m *defaultAuthSeqDeliveriesTxModel) SelectReplayableAfterDate(userId int64, permAuthKeyId int64, date int64, now int64, limit int32) (rList []AuthSeqDeliveries, err error) {
	var (
		query  = "select d.user_id, d.perm_auth_key_id, d.seq, d.`date`, d.payload_id, d.replay_policy, d.source_perm_auth_key_id, d.visibility_policy, d.operation_id, d.expire_at from auth_seq_deliveries as d where d.user_id = ? and d.perm_auth_key_id = ? and d.`date` > ? and (d.expire_at = 0 or d.expire_at > ?) order by d.`date` asc, d.seq asc limit ?"
		values []AuthSeqDeliveries
	)
	err = m.tx.QueryRowsPartial(&values, query, userId, permAuthKeyId, date, now, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthSeqDeliveries{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_seq_deliveries.SelectReplayableAfterDate: %w", err)
		return
	}

	rList = values

	return
}

// SelectReplayableAfterDateWithCB
// select d.user_id, d.perm_auth_key_id, d.seq, d.`date`, d.payload_id, d.replay_policy, d.source_perm_auth_key_id, d.visibility_policy, d.operation_id, d.expire_at from auth_seq_deliveries as d where d.user_id = :user_id and d.perm_auth_key_id = :perm_auth_key_id and d.`date` > :date and (d.expire_at = 0 or d.expire_at > :now) order by d.`date` asc, d.seq asc limit :limit
func (m *defaultAuthSeqDeliveriesModel) SelectReplayableAfterDateWithCB(ctx context.Context, userId int64, permAuthKeyId int64, date int64, now int64, limit int32, cb func(sz, i int, v *AuthSeqDeliveries)) (rList []AuthSeqDeliveries, err error) {
	var (
		query  = "select d.user_id, d.perm_auth_key_id, d.seq, d.`date`, d.payload_id, d.replay_policy, d.source_perm_auth_key_id, d.visibility_policy, d.operation_id, d.expire_at from auth_seq_deliveries as d where d.user_id = ? and d.perm_auth_key_id = ? and d.`date` > ? and (d.expire_at = 0 or d.expire_at > ?) order by d.`date` asc, d.seq asc limit ?"
		values []AuthSeqDeliveries
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, permAuthKeyId, date, now, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthSeqDeliveries{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_seq_deliveries.SelectReplayableAfterDateWithCB: %w", err)
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
