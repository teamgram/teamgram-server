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
	messageSendStatesFieldNames          = builder.RawFieldNames(&MessageSendStates{})
	messageSendStatesRows                = strings.Join(messageSendStatesFieldNames, ",")
	messageSendStatesRowsExpectAutoSet   = strings.Join(stringx.Remove(messageSendStatesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	messageSendStatesRowsWithPlaceHolder = strings.Join(stringx.Remove(messageSendStatesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	messageSendStatesModel interface {
		Insert2(ctx context.Context, data *MessageSendStates) (sql.Result, error)
		FindOne(ctx context.Context, sendStateId int64) (*MessageSendStates, error)
		FindListBySendStateIdList(ctx context.Context, sendStateId ...int64) ([]MessageSendStates, error)
		Update2(ctx context.Context, data *MessageSendStates) error
		Delete2(ctx context.Context, sendStateId int64) error

		FindOneBySenderUserIdPeerTypePeerIdClientRandomId(ctx context.Context, senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*MessageSendStates, error)

		FindOneBySenderOperationId(ctx context.Context, senderOperationId string) (*MessageSendStates, error)
		FindListBySenderOperationIdList(ctx context.Context, senderOperationId ...string) ([]MessageSendStates, error)
	}

	defaultMessageSendStatesModel struct {
		db *sqlx.DB
	}

	MessageSendStates struct {
		SendStateId                 int64  `db:"send_state_id" json:"send_state_id"`
		SenderUserId                int64  `db:"sender_user_id" json:"sender_user_id"`
		PeerType                    int32  `db:"peer_type" json:"peer_type"`
		PeerId                      int64  `db:"peer_id" json:"peer_id"`
		ClientRandomId              int64  `db:"client_random_id" json:"client_random_id"`
		CanonicalMessageId          int64  `db:"canonical_message_id" json:"canonical_message_id"`
		PeerSeq                     int64  `db:"peer_seq" json:"peer_seq"`
		Status                      int32  `db:"status" json:"status"`
		RequestPayloadSchemaVersion int32  `db:"request_payload_schema_version" json:"request_payload_schema_version"`
		RequestPayloadHash          []byte `db:"request_payload_hash" json:"request_payload_hash"`
		SenderOperationId           string `db:"sender_operation_id" json:"sender_operation_id"`
		SenderPts                   int64  `db:"sender_pts" json:"sender_pts"`
		SenderPtsCount              int32  `db:"sender_pts_count" json:"sender_pts_count"`
		SenderUpdateSchemaVersion   int32  `db:"sender_update_schema_version" json:"sender_update_schema_version"`
		SenderUpdatePayload         []byte `db:"sender_update_payload" json:"sender_update_payload"`
		SenderUpdatePayloadHash     []byte `db:"sender_update_payload_hash" json:"sender_update_payload_hash"`
		ReceiverManifestId          int64  `db:"receiver_manifest_id" json:"receiver_manifest_id"`
		LastErrorCategory           int32  `db:"last_error_category" json:"last_error_category"`
		LastErrorCode               string `db:"last_error_code" json:"last_error_code"`
		LastErrorMessage            string `db:"last_error_message" json:"last_error_message"`
		RetryCount                  int32  `db:"retry_count" json:"retry_count"`
		CompletedAt                 int64  `db:"completed_at" json:"completed_at"`
	}
)

func newMessageSendStatesModel(db *sqlx.DB) *defaultMessageSendStatesModel {
	return &defaultMessageSendStatesModel{
		db: db,
	}
}

func (m *defaultMessageSendStatesModel) Insert2(ctx context.Context, data *MessageSendStates) (sql.Result, error) {
	tableName := "message_send_states"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, messageSendStatesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.SenderUserId, data.PeerType, data.PeerId, data.ClientRandomId, data.CanonicalMessageId, data.PeerSeq, data.Status, data.RequestPayloadSchemaVersion, data.RequestPayloadHash, data.SenderOperationId, data.SenderPts, data.SenderPtsCount, data.SenderUpdateSchemaVersion, data.SenderUpdatePayload, data.SenderUpdatePayloadHash, data.ReceiverManifestId, data.LastErrorCategory, data.LastErrorCode, data.LastErrorMessage, data.RetryCount, data.CompletedAt)
	if err != nil {
		return nil, fmt.Errorf("message_send_states.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultMessageSendStatesModel) Delete2(ctx context.Context, sendStateId int64) error {
	tableName := "message_send_states"
	query := fmt.Sprintf("delete from `%s` where `send_state_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, sendStateId)
	if err != nil {
		return fmt.Errorf("message_send_states.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultMessageSendStatesModel) FindOne(ctx context.Context, sendStateId int64) (*MessageSendStates, error) {
	tableName := "message_send_states"
	query := fmt.Sprintf("select %s from %s where send_state_id = ? limit 1", messageSendStatesRows, tableName)
	var resp MessageSendStates

	err := m.db.QueryRowPartial(ctx, &resp, query, sendStateId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_send_states",
				Key:      fmt.Sprintf("send_state_id=%v", sendStateId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("message_send_states.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultMessageSendStatesModel) FindListBySendStateIdList(ctx context.Context, sendStateId ...int64) ([]MessageSendStates, error) {
	if len(sendStateId) == 0 {
		return []MessageSendStates{}, nil
	}
	tableName := "message_send_states"

	query := fmt.Sprintf("select %s from %s where send_state_id in (%s)", messageSendStatesRows, tableName, sqlx.InInt64List(sendStateId))

	var resp []MessageSendStates
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []MessageSendStates{}, nil
		}
		return nil, fmt.Errorf("message_send_states.FindListBySendStateIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultMessageSendStatesModel) Update2(ctx context.Context, data *MessageSendStates) error {
	tableName := "message_send_states"
	query := fmt.Sprintf("update `%s` set %s where `send_state_id` = ?", tableName, messageSendStatesRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.SenderUserId, data.PeerType, data.PeerId, data.ClientRandomId, data.CanonicalMessageId, data.PeerSeq, data.Status, data.RequestPayloadSchemaVersion, data.RequestPayloadHash, data.SenderOperationId, data.SenderPts, data.SenderPtsCount, data.SenderUpdateSchemaVersion, data.SenderUpdatePayload, data.SenderUpdatePayloadHash, data.ReceiverManifestId, data.LastErrorCategory, data.LastErrorCode, data.LastErrorMessage, data.RetryCount, data.CompletedAt, data.SendStateId)
	if err != nil {
		return fmt.Errorf("message_send_states.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultMessageSendStatesModel) FindOneBySenderUserIdPeerTypePeerIdClientRandomId(ctx context.Context, senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*MessageSendStates, error) {
	tableName := "message_send_states"
	query := fmt.Sprintf("select %s from %s where sender_user_id = ? AND peer_type = ? AND peer_id = ? AND client_random_id = ? limit 1", messageSendStatesRows, tableName)
	var resp MessageSendStates

	err := m.db.QueryRowPartial(ctx, &resp, query, senderUserId, peerType, peerId, clientRandomId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_send_states",
				Key:      fmt.Sprintf("sender_user_id=%v,peer_type=%v,peer_id=%v,client_random_id=%v", senderUserId, peerType, peerId, clientRandomId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("message_send_states.FindOneBySenderUserIdPeerTypePeerIdClientRandomId: %w", err)
	}

	return &resp, nil
}

func (m *defaultMessageSendStatesModel) FindOneBySenderOperationId(ctx context.Context, senderOperationId string) (*MessageSendStates, error) {
	tableName := "message_send_states"
	query := fmt.Sprintf("select %s from %s where sender_operation_id = ? limit 1", messageSendStatesRows, tableName)
	var resp MessageSendStates

	err := m.db.QueryRowPartial(ctx, &resp, query, senderOperationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_send_states",
				Key:      fmt.Sprintf("sender_operation_id=%v", senderOperationId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("message_send_states.FindOneBySenderOperationId: %w", err)
	}

	return &resp, nil
}

func (m *defaultMessageSendStatesModel) FindListBySenderOperationIdList(ctx context.Context, senderOperationId ...string) ([]MessageSendStates, error) {
	if len(senderOperationId) == 0 {
		return []MessageSendStates{}, nil
	}
	tableName := "message_send_states"

	query := fmt.Sprintf("select %s from %s where sender_operation_id in (%s)", messageSendStatesRows, tableName, sqlx.InStringList(senderOperationId))
	var resp []MessageSendStates
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []MessageSendStates{}, nil
		}
		return nil, fmt.Errorf("message_send_states.FindListBySenderOperationIdList: %w", err)
	}

	return resp, nil
}
