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
	messageFanoutReceiversFieldNames          = builder.RawFieldNames(&MessageFanoutReceivers{})
	messageFanoutReceiversRows                = strings.Join(messageFanoutReceiversFieldNames, ",")
	messageFanoutReceiversRowsExpectAutoSet   = strings.Join(stringx.Remove(messageFanoutReceiversFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	messageFanoutReceiversRowsWithPlaceHolder = strings.Join(stringx.Remove(messageFanoutReceiversFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	messageFanoutReceiversModel interface {
		Insert2(ctx context.Context, data *MessageFanoutReceivers) (sql.Result, error)

		FindOneByReceiverUserIdOperationId(ctx context.Context, receiverUserId int64, operationId string) (*MessageFanoutReceivers, error)
	}

	defaultMessageFanoutReceiversModel struct {
		db *sqlx.DB
	}

	MessageFanoutReceivers struct {
		ManifestId                    int64  `db:"manifest_id" json:"manifest_id"`
		ReceiverUserId                int64  `db:"receiver_user_id" json:"receiver_user_id"`
		OperationId                   string `db:"operation_id" json:"operation_id"`
		OperationPayloadSchemaVersion int32  `db:"operation_payload_schema_version" json:"operation_payload_schema_version"`
		OperationPayloadCodec         int32  `db:"operation_payload_codec" json:"operation_payload_codec"`
		OperationPayload              []byte `db:"operation_payload" json:"operation_payload"`
		OperationPayloadHash          string `db:"operation_payload_hash" json:"operation_payload_hash"`
		KafkaTopic                    string `db:"kafka_topic" json:"kafka_topic"`
		KafkaPartition                int32  `db:"kafka_partition" json:"kafka_partition"`
		KafkaOffset                   int64  `db:"kafka_offset" json:"kafka_offset"`
		Status                        int32  `db:"status" json:"status"`
		RetryCount                    int32  `db:"retry_count" json:"retry_count"`
		NextRetryAt                   string `db:"next_retry_at" json:"next_retry_at"`
		LastAttemptAt                 string `db:"last_attempt_at" json:"last_attempt_at"`
		LastErrorCode                 string `db:"last_error_code" json:"last_error_code"`
	}
)

func newMessageFanoutReceiversModel(db *sqlx.DB) *defaultMessageFanoutReceiversModel {
	return &defaultMessageFanoutReceiversModel{
		db: db,
	}
}

func (m *defaultMessageFanoutReceiversModel) Insert2(ctx context.Context, data *MessageFanoutReceivers) (sql.Result, error) {
	tableName := "message_fanout_receivers"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, messageFanoutReceiversRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.ManifestId, data.ReceiverUserId, data.OperationId, data.OperationPayloadSchemaVersion, data.OperationPayloadCodec, data.OperationPayload, data.OperationPayloadHash, data.KafkaTopic, data.KafkaPartition, data.KafkaOffset, data.Status, data.RetryCount, data.NextRetryAt, data.LastAttemptAt, data.LastErrorCode)
	if err != nil {
		return nil, fmt.Errorf("message_fanout_receivers.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultMessageFanoutReceiversModel) FindOneByReceiverUserIdOperationId(ctx context.Context, receiverUserId int64, operationId string) (*MessageFanoutReceivers, error) {
	tableName := "message_fanout_receivers"
	query := fmt.Sprintf("select %s from %s where receiver_user_id = ? AND operation_id = ? limit 1", messageFanoutReceiversRows, tableName)
	var resp MessageFanoutReceivers

	err := m.db.QueryRowPartial(ctx, &resp, query, receiverUserId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_fanout_receivers",
				Key:      fmt.Sprintf("receiver_user_id=%v,operation_id=%v", receiverUserId, operationId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("message_fanout_receivers.FindOneByReceiverUserIdOperationId: %w", err)
	}

	return &resp, nil
}
