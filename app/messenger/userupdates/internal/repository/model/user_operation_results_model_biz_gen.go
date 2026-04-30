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

type bizUserOperationResultsModel interface {
	Insert(ctx context.Context, data *UserOperationResults) (lastInsertId, rowsAffected int64, err error)
	SelectByOperation(ctx context.Context, userId int64, operationId string) (*UserOperationResults, error)
	SelectByStatusCompletedBefore(ctx context.Context, status int32, beforeCompletedAt string, limit int32) ([]UserOperationResults, error)
	SelectByStatusCompletedBeforeWithCB(ctx context.Context, status int32, beforeCompletedAt string, limit int32, cb func(sz, i int, v *UserOperationResults)) ([]UserOperationResults, error)
}

type UserOperationResultsTxModel interface {
	Insert(data *UserOperationResults) (lastInsertId, rowsAffected int64, err error)
	SelectByOperation(userId int64, operationId string) (*UserOperationResults, error)
	SelectByStatusCompletedBefore(status int32, beforeCompletedAt string, limit int32) ([]UserOperationResults, error)
}

type defaultUserOperationResultsTxModel struct {
	tx *sqlx.Tx
}

func NewUserOperationResultsTxModel(tx *sqlx.Tx) UserOperationResultsTxModel {
	return &defaultUserOperationResultsTxModel{tx: tx}
}

// Insert
// insert into user_operation_results(user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at) values (:user_id, :operation_id, :op_type, :status, :pts, :pts_count, :payload_hash, :response_schema_version, :response_codec, :response_payload, :response_payload_hash, :terminal_error_code, :completed_at)
func (m *defaultUserOperationResultsModel) Insert(ctx context.Context, data *UserOperationResults) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_operation_results(user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at) values (:user_id, :operation_id, :op_type, :status, :pts, :pts_count, :payload_hash, :response_schema_version, :response_codec, :response_payload, :response_payload_hash, :terminal_error_code, :completed_at)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_operation_results.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_operation_results.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_operation_results.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into user_operation_results(user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at) values (:user_id, :operation_id, :op_type, :status, :pts, :pts_count, :payload_hash, :response_schema_version, :response_codec, :response_payload, :response_payload_hash, :terminal_error_code, :completed_at)
func (m *defaultUserOperationResultsTxModel) Insert(data *UserOperationResults) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_operation_results(user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at) values (:user_id, :operation_id, :op_type, :status, :pts, :pts_count, :payload_hash, :response_schema_version, :response_codec, :response_payload, :response_payload_hash, :terminal_error_code, :completed_at)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_operation_results.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_operation_results.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_operation_results.Insert rows affected: %w", err)
	}

	return
}

// SelectByOperation
// select user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at from user_operation_results where user_id = :user_id and operation_id = :operation_id limit 1
func (m *defaultUserOperationResultsModel) SelectByOperation(ctx context.Context, userId int64, operationId string) (rValue *UserOperationResults, err error) {

	var (
		query = "select user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at from user_operation_results where user_id = ? and operation_id = ? limit 1"
		do    = &UserOperationResults{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_operation_results",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_operation_results.SelectByOperation: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByOperation
// select user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at from user_operation_results where user_id = :user_id and operation_id = :operation_id limit 1
func (m *defaultUserOperationResultsTxModel) SelectByOperation(userId int64, operationId string) (rValue *UserOperationResults, err error) {
	var (
		query = "select user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at from user_operation_results where user_id = ? and operation_id = ? limit 1"
		do    = &UserOperationResults{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_operation_results",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_operation_results.SelectByOperation: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByStatusCompletedBefore
// select user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at from user_operation_results where `status` = :status and completed_at < :beforeCompletedAt order by completed_at asc, user_id asc, operation_id asc limit :limit
func (m *defaultUserOperationResultsModel) SelectByStatusCompletedBefore(ctx context.Context, status int32, beforeCompletedAt string, limit int32) (rList []UserOperationResults, err error) {
	var (
		query  = "select user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at from user_operation_results where `status` = ? and completed_at < ? order by completed_at asc, user_id asc, operation_id asc limit ?"
		values []UserOperationResults
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, status, beforeCompletedAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserOperationResults{}
			err = nil
			return
		}
		err = fmt.Errorf("user_operation_results.SelectByStatusCompletedBefore: %w", err)
		return
	}

	rList = values

	return
}

// SelectByStatusCompletedBefore
// select user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at from user_operation_results where `status` = :status and completed_at < :beforeCompletedAt order by completed_at asc, user_id asc, operation_id asc limit :limit
func (m *defaultUserOperationResultsTxModel) SelectByStatusCompletedBefore(status int32, beforeCompletedAt string, limit int32) (rList []UserOperationResults, err error) {
	var (
		query  = "select user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at from user_operation_results where `status` = ? and completed_at < ? order by completed_at asc, user_id asc, operation_id asc limit ?"
		values []UserOperationResults
	)
	err = m.tx.QueryRowsPartial(&values, query, status, beforeCompletedAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserOperationResults{}
			err = nil
			return
		}
		err = fmt.Errorf("user_operation_results.SelectByStatusCompletedBefore: %w", err)
		return
	}

	rList = values

	return
}

// SelectByStatusCompletedBeforeWithCB
// select user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at from user_operation_results where `status` = :status and completed_at < :beforeCompletedAt order by completed_at asc, user_id asc, operation_id asc limit :limit
func (m *defaultUserOperationResultsModel) SelectByStatusCompletedBeforeWithCB(ctx context.Context, status int32, beforeCompletedAt string, limit int32, cb func(sz, i int, v *UserOperationResults)) (rList []UserOperationResults, err error) {
	var (
		query  = "select user_id, operation_id, op_type, `status`, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, completed_at from user_operation_results where `status` = ? and completed_at < ? order by completed_at asc, user_id asc, operation_id asc limit ?"
		values []UserOperationResults
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, status, beforeCompletedAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserOperationResults{}
			err = nil
			return
		}
		err = fmt.Errorf("user_operation_results.SelectByStatusCompletedBeforeWithCB: %w", err)
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
