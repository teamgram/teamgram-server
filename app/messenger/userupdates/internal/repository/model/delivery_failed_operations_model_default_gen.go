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

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	deliveryFailedOperationsFieldNames          = builder.RawFieldNames(&DeliveryFailedOperations{})
	deliveryFailedOperationsRows                = strings.Join(deliveryFailedOperationsFieldNames, ",")
	deliveryFailedOperationsRowsExpectAutoSet   = strings.Join(stringx.Remove(deliveryFailedOperationsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	deliveryFailedOperationsRowsWithPlaceHolder = strings.Join(stringx.Remove(deliveryFailedOperationsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	deliveryFailedOperationsModel interface {
		Insert2(ctx context.Context, data *DeliveryFailedOperations) (sql.Result, error)
		FindOne(ctx context.Context, failedId int64) (*DeliveryFailedOperations, error)
		FindListByFailedIdList(ctx context.Context, failedId ...int64) ([]DeliveryFailedOperations, error)
		Update2(ctx context.Context, data *DeliveryFailedOperations) error
		Delete2(ctx context.Context, failedId int64) error

		FindOneByKafkaTopicKafkaPartitionKafkaOffset(ctx context.Context, kafkaTopic string, kafkaPartition int32, kafkaOffset int64) (*DeliveryFailedOperations, error)
	}

	defaultDeliveryFailedOperationsModel struct {
		db *sqlx.DB
	}

	DeliveryFailedOperations struct {
		FailedId             int64  `db:"failed_id" json:"failed_id"`
		UserId               int64  `db:"user_id" json:"user_id"`
		OperationId          string `db:"operation_id" json:"operation_id"`
		OpType               int32  `db:"op_type" json:"op_type"`
		BucketId             int32  `db:"bucket_id" json:"bucket_id"`
		KafkaTopic           string `db:"kafka_topic" json:"kafka_topic"`
		KafkaPartition       int32  `db:"kafka_partition" json:"kafka_partition"`
		KafkaOffset          int64  `db:"kafka_offset" json:"kafka_offset"`
		PayloadSchemaVersion int32  `db:"payload_schema_version" json:"payload_schema_version"`
		PayloadHash          string `db:"payload_hash" json:"payload_hash"`
		FailureCategory      int32  `db:"failure_category" json:"failure_category"`
		FailureCode          string `db:"failure_code" json:"failure_code"`
		FailureMessage       string `db:"failure_message" json:"failure_message"`
		RetryCount           int32  `db:"retry_count" json:"retry_count"`
		Status               int32  `db:"status" json:"status"`
		ReplayedAt           string `db:"replayed_at" json:"replayed_at"`
		ReplayedBy           string `db:"replayed_by" json:"replayed_by"`
	}
)

func newDeliveryFailedOperationsModel(db *sqlx.DB) *defaultDeliveryFailedOperationsModel {
	return &defaultDeliveryFailedOperationsModel{
		db: db,
	}
}

func (m *defaultDeliveryFailedOperationsModel) Insert2(ctx context.Context, data *DeliveryFailedOperations) (sql.Result, error) {
	tableName := "delivery_failed_operations"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, deliveryFailedOperationsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.OperationId, data.OpType, data.BucketId, data.KafkaTopic, data.KafkaPartition, data.KafkaOffset, data.PayloadSchemaVersion, data.PayloadHash, data.FailureCategory, data.FailureCode, data.FailureMessage, data.RetryCount, data.Status, data.ReplayedAt, data.ReplayedBy)
	if err != nil {
		return nil, fmt.Errorf("delivery_failed_operations.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultDeliveryFailedOperationsModel) Delete2(ctx context.Context, failedId int64) error {
	tableName := "delivery_failed_operations"
	query := fmt.Sprintf("delete from `%s` where `failed_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, failedId)
	if err != nil {
		return fmt.Errorf("delivery_failed_operations.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultDeliveryFailedOperationsModel) FindOne(ctx context.Context, failedId int64) (*DeliveryFailedOperations, error) {
	tableName := "delivery_failed_operations"
	query := fmt.Sprintf("select %s from %s where failed_id = ? limit 1", deliveryFailedOperationsRows, tableName)
	var resp DeliveryFailedOperations

	err := m.db.QueryRowPartial(ctx, &resp, query, failedId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "delivery_failed_operations",
				Key:      fmt.Sprintf("failed_id=%v", failedId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("delivery_failed_operations.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultDeliveryFailedOperationsModel) FindListByFailedIdList(ctx context.Context, failedId ...int64) ([]DeliveryFailedOperations, error) {
	if len(failedId) == 0 {
		return []DeliveryFailedOperations{}, nil
	}
	tableName := "delivery_failed_operations"

	query := fmt.Sprintf("select %s from %s where failed_id in (%s)", deliveryFailedOperationsRows, tableName, sqlx.InInt64List(failedId))

	var resp []DeliveryFailedOperations
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []DeliveryFailedOperations{}, nil
		}
		return nil, fmt.Errorf("delivery_failed_operations.FindListByFailedIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultDeliveryFailedOperationsModel) Update2(ctx context.Context, data *DeliveryFailedOperations) error {
	tableName := "delivery_failed_operations"
	query := fmt.Sprintf("update `%s` set %s where `failed_id` = ?", tableName, deliveryFailedOperationsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.OperationId, data.OpType, data.BucketId, data.KafkaTopic, data.KafkaPartition, data.KafkaOffset, data.PayloadSchemaVersion, data.PayloadHash, data.FailureCategory, data.FailureCode, data.FailureMessage, data.RetryCount, data.Status, data.ReplayedAt, data.ReplayedBy, data.FailedId)
	if err != nil {
		return fmt.Errorf("delivery_failed_operations.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultDeliveryFailedOperationsModel) FindOneByKafkaTopicKafkaPartitionKafkaOffset(ctx context.Context, kafkaTopic string, kafkaPartition int32, kafkaOffset int64) (*DeliveryFailedOperations, error) {
	tableName := "delivery_failed_operations"
	query := fmt.Sprintf("select %s from %s where kafka_topic = ? AND kafka_partition = ? AND kafka_offset = ? limit 1", deliveryFailedOperationsRows, tableName)
	var resp DeliveryFailedOperations

	err := m.db.QueryRowPartial(ctx, &resp, query, kafkaTopic, kafkaPartition, kafkaOffset)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "delivery_failed_operations",
				Key:      fmt.Sprintf("kafka_topic=%v,kafka_partition=%v,kafka_offset=%v", kafkaTopic, kafkaPartition, kafkaOffset),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("delivery_failed_operations.FindOneByKafkaTopicKafkaPartitionKafkaOffset: %w", err)
	}

	return &resp, nil
}
