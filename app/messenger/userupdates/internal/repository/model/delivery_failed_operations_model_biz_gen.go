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

type bizDeliveryFailedOperationsModel interface {
	Insert(ctx context.Context, data *DeliveryFailedOperations) (lastInsertId, rowsAffected int64, err error)
	SelectByKafkaOffset(ctx context.Context, kafkaTopic string, kafkaPartition int32, kafkaOffset int64) (*DeliveryFailedOperations, error)
	SelectByUserOperation(ctx context.Context, userId int64, operationId string) (*DeliveryFailedOperations, error)
	SelectByBucketStatus(ctx context.Context, bucketId int32, status int32, limit int32) ([]DeliveryFailedOperations, error)
	SelectByBucketStatusWithCB(ctx context.Context, bucketId int32, status int32, limit int32, cb func(sz, i int, v *DeliveryFailedOperations)) ([]DeliveryFailedOperations, error)
	MarkStatus(ctx context.Context, status int32, retryCount int32, failedId int64) (rowsAffected int64, err error)
	MarkReplayed(ctx context.Context, status int32, replayedAt sql.NullTime, replayedBy string, failedId int64) (rowsAffected int64, err error)
}

type DeliveryFailedOperationsTxModel interface {
	Insert(data *DeliveryFailedOperations) (lastInsertId, rowsAffected int64, err error)
	SelectByKafkaOffset(kafkaTopic string, kafkaPartition int32, kafkaOffset int64) (*DeliveryFailedOperations, error)
	SelectByUserOperation(userId int64, operationId string) (*DeliveryFailedOperations, error)
	SelectByBucketStatus(bucketId int32, status int32, limit int32) ([]DeliveryFailedOperations, error)
	MarkStatus(status int32, retryCount int32, failedId int64) (rowsAffected int64, err error)
	MarkReplayed(status int32, replayedAt sql.NullTime, replayedBy string, failedId int64) (rowsAffected int64, err error)
}

type defaultDeliveryFailedOperationsTxModel struct {
	tx *sqlx.Tx
}

func NewDeliveryFailedOperationsTxModel(tx *sqlx.Tx) DeliveryFailedOperationsTxModel {
	return &defaultDeliveryFailedOperationsTxModel{tx: tx}
}

// Insert
// insert into delivery_failed_operations(failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at) values (:failed_id, :user_id, :operation_id, :op_type, :bucket_id, :kafka_topic, :kafka_partition, :kafka_offset, :payload_schema_version, :payload_hash, :failure_category, :failure_code, :failure_message, :retry_count, :status, :failed_at)
func (m *defaultDeliveryFailedOperationsModel) Insert(ctx context.Context, data *DeliveryFailedOperations) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into delivery_failed_operations(failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at) values (:failed_id, :user_id, :operation_id, :op_type, :bucket_id, :kafka_topic, :kafka_partition, :kafka_offset, :payload_schema_version, :payload_hash, :failure_category, :failure_code, :failure_message, :retry_count, :status, :failed_at)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into delivery_failed_operations(failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at) values (:failed_id, :user_id, :operation_id, :op_type, :bucket_id, :kafka_topic, :kafka_partition, :kafka_offset, :payload_schema_version, :payload_hash, :failure_category, :failure_code, :failure_message, :retry_count, :status, :failed_at)
func (m *defaultDeliveryFailedOperationsTxModel) Insert(data *DeliveryFailedOperations) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into delivery_failed_operations(failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at) values (:failed_id, :user_id, :operation_id, :op_type, :bucket_id, :kafka_topic, :kafka_partition, :kafka_offset, :payload_schema_version, :payload_hash, :failure_category, :failure_code, :failure_message, :retry_count, :status, :failed_at)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.Insert rows affected: %w", err)
	}

	return
}

// SelectByKafkaOffset
// select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where kafka_topic = :kafka_topic and kafka_partition = :kafka_partition and kafka_offset = :kafka_offset limit 1
func (m *defaultDeliveryFailedOperationsModel) SelectByKafkaOffset(ctx context.Context, kafkaTopic string, kafkaPartition int32, kafkaOffset int64) (rValue *DeliveryFailedOperations, err error) {

	var (
		query = "select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where kafka_topic = ? and kafka_partition = ? and kafka_offset = ? limit 1"
		do    = &DeliveryFailedOperations{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, kafkaTopic, kafkaPartition, kafkaOffset)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "delivery_failed_operations",
				Key:      fmt.Sprintf("kafka_topic=%v,kafka_partition=%v,kafka_offset=%v", kafkaTopic, kafkaPartition, kafkaOffset),
				Cause:    err,
			}
		}
		err = fmt.Errorf("delivery_failed_operations.SelectByKafkaOffset: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByKafkaOffset
// select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where kafka_topic = :kafka_topic and kafka_partition = :kafka_partition and kafka_offset = :kafka_offset limit 1
func (m *defaultDeliveryFailedOperationsTxModel) SelectByKafkaOffset(kafkaTopic string, kafkaPartition int32, kafkaOffset int64) (rValue *DeliveryFailedOperations, err error) {
	var (
		query = "select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where kafka_topic = ? and kafka_partition = ? and kafka_offset = ? limit 1"
		do    = &DeliveryFailedOperations{}
	)
	err = m.tx.QueryRowPartial(do, query, kafkaTopic, kafkaPartition, kafkaOffset)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "delivery_failed_operations",
				Key:      fmt.Sprintf("kafka_topic=%v,kafka_partition=%v,kafka_offset=%v", kafkaTopic, kafkaPartition, kafkaOffset),
				Cause:    err,
			}
		}
		err = fmt.Errorf("delivery_failed_operations.SelectByKafkaOffset: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByUserOperation
// select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where user_id = :user_id and operation_id = :operation_id limit 1
func (m *defaultDeliveryFailedOperationsModel) SelectByUserOperation(ctx context.Context, userId int64, operationId string) (rValue *DeliveryFailedOperations, err error) {

	var (
		query = "select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where user_id = ? and operation_id = ? limit 1"
		do    = &DeliveryFailedOperations{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "delivery_failed_operations",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("delivery_failed_operations.SelectByUserOperation: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUserOperation
// select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where user_id = :user_id and operation_id = :operation_id limit 1
func (m *defaultDeliveryFailedOperationsTxModel) SelectByUserOperation(userId int64, operationId string) (rValue *DeliveryFailedOperations, err error) {
	var (
		query = "select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where user_id = ? and operation_id = ? limit 1"
		do    = &DeliveryFailedOperations{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "delivery_failed_operations",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("delivery_failed_operations.SelectByUserOperation: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByBucketStatus
// select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where bucket_id = :bucket_id and `status` = :status order by failed_at asc, failed_id asc limit :limit
func (m *defaultDeliveryFailedOperationsModel) SelectByBucketStatus(ctx context.Context, bucketId int32, status int32, limit int32) (rList []DeliveryFailedOperations, err error) {
	var (
		query  = "select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where bucket_id = ? and `status` = ? order by failed_at asc, failed_id asc limit ?"
		values []DeliveryFailedOperations
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, bucketId, status, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DeliveryFailedOperations{}
			err = nil
			return
		}
		err = fmt.Errorf("delivery_failed_operations.SelectByBucketStatus: %w", err)
		return
	}

	rList = values

	return
}

// SelectByBucketStatus
// select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where bucket_id = :bucket_id and `status` = :status order by failed_at asc, failed_id asc limit :limit
func (m *defaultDeliveryFailedOperationsTxModel) SelectByBucketStatus(bucketId int32, status int32, limit int32) (rList []DeliveryFailedOperations, err error) {
	var (
		query  = "select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where bucket_id = ? and `status` = ? order by failed_at asc, failed_id asc limit ?"
		values []DeliveryFailedOperations
	)
	err = m.tx.QueryRowsPartial(&values, query, bucketId, status, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DeliveryFailedOperations{}
			err = nil
			return
		}
		err = fmt.Errorf("delivery_failed_operations.SelectByBucketStatus: %w", err)
		return
	}

	rList = values

	return
}

// SelectByBucketStatusWithCB
// select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where bucket_id = :bucket_id and `status` = :status order by failed_at asc, failed_id asc limit :limit
func (m *defaultDeliveryFailedOperationsModel) SelectByBucketStatusWithCB(ctx context.Context, bucketId int32, status int32, limit int32, cb func(sz, i int, v *DeliveryFailedOperations)) (rList []DeliveryFailedOperations, err error) {
	var (
		query  = "select failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic, kafka_partition, kafka_offset, payload_schema_version, payload_hash, failure_category, failure_code, failure_message, retry_count, `status`, failed_at from delivery_failed_operations where bucket_id = ? and `status` = ? order by failed_at asc, failed_id asc limit ?"
		values []DeliveryFailedOperations
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, bucketId, status, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DeliveryFailedOperations{}
			err = nil
			return
		}
		err = fmt.Errorf("delivery_failed_operations.SelectByBucketStatusWithCB: %w", err)
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

// MarkStatus
// update delivery_failed_operations set `status` = :status, retry_count = :retry_count where failed_id = :failed_id
func (m *defaultDeliveryFailedOperationsModel) MarkStatus(ctx context.Context, status int32, retryCount int32, failedId int64) (rowsAffected int64, err error) {

	var (
		query   = "update delivery_failed_operations set `status` = ?, retry_count = ? where failed_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, retryCount, failedId)

	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.MarkStatus exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.MarkStatus rows affected: %w", err)
		return
	}

	return
}

// MarkStatus
// update delivery_failed_operations set `status` = :status, retry_count = :retry_count where failed_id = :failed_id
func (m *defaultDeliveryFailedOperationsTxModel) MarkStatus(status int32, retryCount int32, failedId int64) (rowsAffected int64, err error) {
	var (
		query   = "update delivery_failed_operations set `status` = ?, retry_count = ? where failed_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, retryCount, failedId)

	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.MarkStatus exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.MarkStatus rows affected: %w", err)
		return
	}

	return
}

// MarkReplayed
// update delivery_failed_operations set `status` = :status, replayed_at = :replayed_at, replayed_by = :replayed_by where failed_id = :failed_id
func (m *defaultDeliveryFailedOperationsModel) MarkReplayed(ctx context.Context, status int32, replayedAt sql.NullTime, replayedBy string, failedId int64) (rowsAffected int64, err error) {

	var (
		query   = "update delivery_failed_operations set `status` = ?, replayed_at = ?, replayed_by = ? where failed_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, replayedAt, replayedBy, failedId)

	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.MarkReplayed exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.MarkReplayed rows affected: %w", err)
		return
	}

	return
}

// MarkReplayed
// update delivery_failed_operations set `status` = :status, replayed_at = :replayed_at, replayed_by = :replayed_by where failed_id = :failed_id
func (m *defaultDeliveryFailedOperationsTxModel) MarkReplayed(status int32, replayedAt sql.NullTime, replayedBy string, failedId int64) (rowsAffected int64, err error) {
	var (
		query   = "update delivery_failed_operations set `status` = ?, replayed_at = ?, replayed_by = ? where failed_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, replayedAt, replayedBy, failedId)

	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.MarkReplayed exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("delivery_failed_operations.MarkReplayed rows affected: %w", err)
		return
	}

	return
}
