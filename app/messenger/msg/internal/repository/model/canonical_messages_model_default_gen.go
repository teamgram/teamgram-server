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
	canonicalMessagesFieldNames          = builder.RawFieldNames(&CanonicalMessages{})
	canonicalMessagesRows                = strings.Join(canonicalMessagesFieldNames, ",")
	canonicalMessagesRowsExpectAutoSet   = strings.Join(stringx.Remove(canonicalMessagesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	canonicalMessagesRowsWithPlaceHolder = strings.Join(stringx.Remove(canonicalMessagesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	canonicalMessagesModel interface {
		Insert2(ctx context.Context, data *CanonicalMessages) (sql.Result, error)
		FindOne(ctx context.Context, canonicalMessageId int64) (*CanonicalMessages, error)
		FindListByCanonicalMessageIdList(ctx context.Context, canonicalMessageId ...int64) ([]CanonicalMessages, error)
		Update2(ctx context.Context, data *CanonicalMessages) error
		Delete2(ctx context.Context, canonicalMessageId int64) error

		FindOneByPeerTypePeerIdPeerSeq(ctx context.Context, peerType int32, peerId int64, peerSeq int64) (*CanonicalMessages, error)
	}

	defaultCanonicalMessagesModel struct {
		db *sqlx.DB
	}

	CanonicalMessages struct {
		CanonicalMessageId           int64  `db:"canonical_message_id" json:"canonical_message_id"`
		PeerType                     int32  `db:"peer_type" json:"peer_type"`
		PeerId                       int64  `db:"peer_id" json:"peer_id"`
		PeerSeq                      int64  `db:"peer_seq" json:"peer_seq"`
		FromUserId                   int64  `db:"from_user_id" json:"from_user_id"`
		MessageKind                  int32  `db:"message_kind" json:"message_kind"`
		MessageText                  string `db:"message_text" json:"message_text"`
		EntitiesPayloadSchemaVersion int32  `db:"entities_payload_schema_version" json:"entities_payload_schema_version"`
		EntitiesPayload              []byte `db:"entities_payload" json:"entities_payload"`
		MediaRefSchemaVersion        int32  `db:"media_ref_schema_version" json:"media_ref_schema_version"`
		MediaRefPayload              []byte `db:"media_ref_payload" json:"media_ref_payload"`
		ServiceActionSchemaVersion   int32  `db:"service_action_schema_version" json:"service_action_schema_version"`
		ServiceActionPayload         []byte `db:"service_action_payload" json:"service_action_payload"`
		MessageStatus                int32  `db:"message_status" json:"message_status"`
		EditVersion                  int32  `db:"edit_version" json:"edit_version"`
		Date                         string `db:"date" json:"date"`
		EditDate                     string `db:"edit_date" json:"edit_date"`
		DeletedAt                    string `db:"deleted_at" json:"deleted_at"`
		StorageSchemaVersion         int32  `db:"storage_schema_version" json:"storage_schema_version"`
	}
)

func newCanonicalMessagesModel(db *sqlx.DB) *defaultCanonicalMessagesModel {
	return &defaultCanonicalMessagesModel{
		db: db,
	}
}

func (m *defaultCanonicalMessagesModel) Insert2(ctx context.Context, data *CanonicalMessages) (sql.Result, error) {
	tableName := "canonical_messages"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, canonicalMessagesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.PeerType, data.PeerId, data.PeerSeq, data.FromUserId, data.MessageKind, data.MessageText, data.EntitiesPayloadSchemaVersion, data.EntitiesPayload, data.MediaRefSchemaVersion, data.MediaRefPayload, data.ServiceActionSchemaVersion, data.ServiceActionPayload, data.MessageStatus, data.EditVersion, data.Date, data.EditDate, data.DeletedAt, data.StorageSchemaVersion)
	if err != nil {
		return nil, fmt.Errorf("canonical_messages.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultCanonicalMessagesModel) Delete2(ctx context.Context, canonicalMessageId int64) error {
	tableName := "canonical_messages"
	query := fmt.Sprintf("delete from `%s` where `canonical_message_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, canonicalMessageId)
	if err != nil {
		return fmt.Errorf("canonical_messages.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultCanonicalMessagesModel) FindOne(ctx context.Context, canonicalMessageId int64) (*CanonicalMessages, error) {
	tableName := "canonical_messages"
	query := fmt.Sprintf("select %s from %s where canonical_message_id = ? limit 1", canonicalMessagesRows, tableName)
	var resp CanonicalMessages

	err := m.db.QueryRowPartial(ctx, &resp, query, canonicalMessageId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "canonical_messages",
				Key:      fmt.Sprintf("canonical_message_id=%v", canonicalMessageId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("canonical_messages.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultCanonicalMessagesModel) FindListByCanonicalMessageIdList(ctx context.Context, canonicalMessageId ...int64) ([]CanonicalMessages, error) {
	if len(canonicalMessageId) == 0 {
		return []CanonicalMessages{}, nil
	}
	tableName := "canonical_messages"

	query := fmt.Sprintf("select %s from %s where canonical_message_id in (%s)", canonicalMessagesRows, tableName, sqlx.InInt64List(canonicalMessageId))

	var resp []CanonicalMessages
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []CanonicalMessages{}, nil
		}
		return nil, fmt.Errorf("canonical_messages.FindListByCanonicalMessageIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultCanonicalMessagesModel) Update2(ctx context.Context, data *CanonicalMessages) error {
	tableName := "canonical_messages"
	query := fmt.Sprintf("update `%s` set %s where `canonical_message_id` = ?", tableName, canonicalMessagesRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.PeerType, data.PeerId, data.PeerSeq, data.FromUserId, data.MessageKind, data.MessageText, data.EntitiesPayloadSchemaVersion, data.EntitiesPayload, data.MediaRefSchemaVersion, data.MediaRefPayload, data.ServiceActionSchemaVersion, data.ServiceActionPayload, data.MessageStatus, data.EditVersion, data.Date, data.EditDate, data.DeletedAt, data.StorageSchemaVersion, data.CanonicalMessageId)
	if err != nil {
		return fmt.Errorf("canonical_messages.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultCanonicalMessagesModel) FindOneByPeerTypePeerIdPeerSeq(ctx context.Context, peerType int32, peerId int64, peerSeq int64) (*CanonicalMessages, error) {
	tableName := "canonical_messages"
	query := fmt.Sprintf("select %s from %s where peer_type = ? AND peer_id = ? AND peer_seq = ? limit 1", canonicalMessagesRows, tableName)
	var resp CanonicalMessages

	err := m.db.QueryRowPartial(ctx, &resp, query, peerType, peerId, peerSeq)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "canonical_messages",
				Key:      fmt.Sprintf("peer_type=%v,peer_id=%v,peer_seq=%v", peerType, peerId, peerSeq),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("canonical_messages.FindOneByPeerTypePeerIdPeerSeq: %w", err)
	}

	return &resp, nil
}
