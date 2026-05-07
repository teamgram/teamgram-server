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

type bizMessageFanoutReceiversModel interface {
	Insert(ctx context.Context, data *MessageFanoutReceivers) (lastInsertId, rowsAffected int64, err error)
	SelectByReceiverOperation(ctx context.Context, receiverUserId int64, operationId string) (*MessageFanoutReceivers, error)
	SelectByManifest(ctx context.Context, manifestId int64) ([]MessageFanoutReceivers, error)
	SelectByManifestWithCB(ctx context.Context, manifestId int64, cb func(sz, i int, v *MessageFanoutReceivers)) ([]MessageFanoutReceivers, error)
	MarkPublished(ctx context.Context, kafkaTopic string, kafkaPartition int32, kafkaOffset int64, status int32, lastAttemptAt sql.NullTime, manifestId int64, receiverUserId int64) (rowsAffected int64, err error)
	MarkRetryableFailure(ctx context.Context, status int32, nextRetryAt sql.NullTime, lastAttemptAt sql.NullTime, lastErrorCode string, manifestId int64, receiverUserId int64) (rowsAffected int64, err error)
}

type MessageFanoutReceiversTxModel interface {
	Insert(data *MessageFanoutReceivers) (lastInsertId, rowsAffected int64, err error)
	SelectByReceiverOperation(receiverUserId int64, operationId string) (*MessageFanoutReceivers, error)
	SelectByManifest(manifestId int64) ([]MessageFanoutReceivers, error)
	MarkPublished(kafkaTopic string, kafkaPartition int32, kafkaOffset int64, status int32, lastAttemptAt sql.NullTime, manifestId int64, receiverUserId int64) (rowsAffected int64, err error)
	MarkRetryableFailure(status int32, nextRetryAt sql.NullTime, lastAttemptAt sql.NullTime, lastErrorCode string, manifestId int64, receiverUserId int64) (rowsAffected int64, err error)
}

type defaultMessageFanoutReceiversTxModel struct {
	tx *sqlx.Tx
}

func NewMessageFanoutReceiversTxModel(tx *sqlx.Tx) MessageFanoutReceiversTxModel {
	return &defaultMessageFanoutReceiversTxModel{tx: tx}
}

// Insert
// insert into message_fanout_receivers(manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code) values (:manifest_id, :receiver_user_id, :operation_id, :operation_payload_schema_version, :operation_payload_codec, :operation_payload, :operation_payload_hash, :kafka_topic, :kafka_partition, :kafka_offset, :status, :retry_count, :next_retry_at, :last_attempt_at, :last_error_code)
func (m *defaultMessageFanoutReceiversModel) Insert(ctx context.Context, data *MessageFanoutReceivers) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_fanout_receivers(manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code) values (:manifest_id, :receiver_user_id, :operation_id, :operation_payload_schema_version, :operation_payload_codec, :operation_payload, :operation_payload_hash, :kafka_topic, :kafka_partition, :kafka_offset, :status, :retry_count, :next_retry_at, :last_attempt_at, :last_error_code)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into message_fanout_receivers(manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code) values (:manifest_id, :receiver_user_id, :operation_id, :operation_payload_schema_version, :operation_payload_codec, :operation_payload, :operation_payload_hash, :kafka_topic, :kafka_partition, :kafka_offset, :status, :retry_count, :next_retry_at, :last_attempt_at, :last_error_code)
func (m *defaultMessageFanoutReceiversTxModel) Insert(data *MessageFanoutReceivers) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_fanout_receivers(manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code) values (:manifest_id, :receiver_user_id, :operation_id, :operation_payload_schema_version, :operation_payload_codec, :operation_payload, :operation_payload_hash, :kafka_topic, :kafka_partition, :kafka_offset, :status, :retry_count, :next_retry_at, :last_attempt_at, :last_error_code)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.Insert rows affected: %w", err)
	}

	return
}

// SelectByReceiverOperation
// select manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code from message_fanout_receivers where receiver_user_id = :receiver_user_id and operation_id = :operation_id limit 1
func (m *defaultMessageFanoutReceiversModel) SelectByReceiverOperation(ctx context.Context, receiverUserId int64, operationId string) (rValue *MessageFanoutReceivers, err error) {

	var (
		query = "select manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code from message_fanout_receivers where receiver_user_id = ? and operation_id = ? limit 1"
		do    = &MessageFanoutReceivers{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, receiverUserId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_fanout_receivers",
				Key:      fmt.Sprintf("receiver_user_id=%v,operation_id=%v", receiverUserId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_fanout_receivers.SelectByReceiverOperation: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByReceiverOperation
// select manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code from message_fanout_receivers where receiver_user_id = :receiver_user_id and operation_id = :operation_id limit 1
func (m *defaultMessageFanoutReceiversTxModel) SelectByReceiverOperation(receiverUserId int64, operationId string) (rValue *MessageFanoutReceivers, err error) {
	var (
		query = "select manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code from message_fanout_receivers where receiver_user_id = ? and operation_id = ? limit 1"
		do    = &MessageFanoutReceivers{}
	)
	err = m.tx.QueryRowPartial(do, query, receiverUserId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_fanout_receivers",
				Key:      fmt.Sprintf("receiver_user_id=%v,operation_id=%v", receiverUserId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_fanout_receivers.SelectByReceiverOperation: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByManifest
// select manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code from message_fanout_receivers where manifest_id = :manifest_id order by receiver_user_id asc
func (m *defaultMessageFanoutReceiversModel) SelectByManifest(ctx context.Context, manifestId int64) (rList []MessageFanoutReceivers, err error) {
	var (
		query  = "select manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code from message_fanout_receivers where manifest_id = ? order by receiver_user_id asc"
		values []MessageFanoutReceivers
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, manifestId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []MessageFanoutReceivers{}
			err = nil
			return
		}
		err = fmt.Errorf("message_fanout_receivers.SelectByManifest: %w", err)
		return
	}

	rList = values

	return
}

// SelectByManifest
// select manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code from message_fanout_receivers where manifest_id = :manifest_id order by receiver_user_id asc
func (m *defaultMessageFanoutReceiversTxModel) SelectByManifest(manifestId int64) (rList []MessageFanoutReceivers, err error) {
	var (
		query  = "select manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code from message_fanout_receivers where manifest_id = ? order by receiver_user_id asc"
		values []MessageFanoutReceivers
	)
	err = m.tx.QueryRowsPartial(&values, query, manifestId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []MessageFanoutReceivers{}
			err = nil
			return
		}
		err = fmt.Errorf("message_fanout_receivers.SelectByManifest: %w", err)
		return
	}

	rList = values

	return
}

// SelectByManifestWithCB
// select manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code from message_fanout_receivers where manifest_id = :manifest_id order by receiver_user_id asc
func (m *defaultMessageFanoutReceiversModel) SelectByManifestWithCB(ctx context.Context, manifestId int64, cb func(sz, i int, v *MessageFanoutReceivers)) (rList []MessageFanoutReceivers, err error) {
	var (
		query  = "select manifest_id, receiver_user_id, operation_id, operation_payload_schema_version, operation_payload_codec, operation_payload, operation_payload_hash, kafka_topic, kafka_partition, kafka_offset, `status`, retry_count, next_retry_at, last_attempt_at, last_error_code from message_fanout_receivers where manifest_id = ? order by receiver_user_id asc"
		values []MessageFanoutReceivers
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, manifestId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []MessageFanoutReceivers{}
			err = nil
			return
		}
		err = fmt.Errorf("message_fanout_receivers.SelectByManifestWithCB: %w", err)
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

// MarkPublished
// update message_fanout_receivers set kafka_topic = :kafka_topic, kafka_partition = :kafka_partition, kafka_offset = :kafka_offset, `status` = :status, last_attempt_at = :last_attempt_at where manifest_id = :manifest_id and receiver_user_id = :receiver_user_id
func (m *defaultMessageFanoutReceiversModel) MarkPublished(ctx context.Context, kafkaTopic string, kafkaPartition int32, kafkaOffset int64, status int32, lastAttemptAt sql.NullTime, manifestId int64, receiverUserId int64) (rowsAffected int64, err error) {

	var (
		query   = "update message_fanout_receivers set kafka_topic = ?, kafka_partition = ?, kafka_offset = ?, `status` = ?, last_attempt_at = ? where manifest_id = ? and receiver_user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, kafkaTopic, kafkaPartition, kafkaOffset, status, lastAttemptAt, manifestId, receiverUserId)

	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.MarkPublished exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.MarkPublished rows affected: %w", err)
		return
	}

	return
}

// MarkPublished
// update message_fanout_receivers set kafka_topic = :kafka_topic, kafka_partition = :kafka_partition, kafka_offset = :kafka_offset, `status` = :status, last_attempt_at = :last_attempt_at where manifest_id = :manifest_id and receiver_user_id = :receiver_user_id
func (m *defaultMessageFanoutReceiversTxModel) MarkPublished(kafkaTopic string, kafkaPartition int32, kafkaOffset int64, status int32, lastAttemptAt sql.NullTime, manifestId int64, receiverUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update message_fanout_receivers set kafka_topic = ?, kafka_partition = ?, kafka_offset = ?, `status` = ?, last_attempt_at = ? where manifest_id = ? and receiver_user_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, kafkaTopic, kafkaPartition, kafkaOffset, status, lastAttemptAt, manifestId, receiverUserId)

	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.MarkPublished exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.MarkPublished rows affected: %w", err)
		return
	}

	return
}

// MarkRetryableFailure
// update message_fanout_receivers set `status` = :status, retry_count = retry_count + 1, next_retry_at = :next_retry_at, last_attempt_at = :last_attempt_at, last_error_code = :last_error_code where manifest_id = :manifest_id and receiver_user_id = :receiver_user_id
func (m *defaultMessageFanoutReceiversModel) MarkRetryableFailure(ctx context.Context, status int32, nextRetryAt sql.NullTime, lastAttemptAt sql.NullTime, lastErrorCode string, manifestId int64, receiverUserId int64) (rowsAffected int64, err error) {

	var (
		query   = "update message_fanout_receivers set `status` = ?, retry_count = retry_count + 1, next_retry_at = ?, last_attempt_at = ?, last_error_code = ? where manifest_id = ? and receiver_user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, nextRetryAt, lastAttemptAt, lastErrorCode, manifestId, receiverUserId)

	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.MarkRetryableFailure exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.MarkRetryableFailure rows affected: %w", err)
		return
	}

	return
}

// MarkRetryableFailure
// update message_fanout_receivers set `status` = :status, retry_count = retry_count + 1, next_retry_at = :next_retry_at, last_attempt_at = :last_attempt_at, last_error_code = :last_error_code where manifest_id = :manifest_id and receiver_user_id = :receiver_user_id
func (m *defaultMessageFanoutReceiversTxModel) MarkRetryableFailure(status int32, nextRetryAt sql.NullTime, lastAttemptAt sql.NullTime, lastErrorCode string, manifestId int64, receiverUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update message_fanout_receivers set `status` = ?, retry_count = retry_count + 1, next_retry_at = ?, last_attempt_at = ?, last_error_code = ? where manifest_id = ? and receiver_user_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, nextRetryAt, lastAttemptAt, lastErrorCode, manifestId, receiverUserId)

	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.MarkRetryableFailure exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_fanout_receivers.MarkRetryableFailure rows affected: %w", err)
		return
	}

	return
}
