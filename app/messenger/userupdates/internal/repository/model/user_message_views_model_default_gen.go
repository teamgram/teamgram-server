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
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userMessageViewsFieldNames          = builder.RawFieldNames(&UserMessageViews{})
	userMessageViewsRows                = strings.Join(userMessageViewsFieldNames, ",")
	userMessageViewsRowsExpectAutoSet   = strings.Join(stringx.Remove(userMessageViewsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userMessageViewsRowsWithPlaceHolder = strings.Join(stringx.Remove(userMessageViewsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userMessageViewsModel interface {
		Insert2(ctx context.Context, data *UserMessageViews) (sql.Result, error)

		FindOneByUserIdCanonicalMessageId(ctx context.Context, userId int64, canonicalMessageId int64) (*UserMessageViews, error)
	}

	defaultUserMessageViewsModel struct {
		db *sqlx.DB
	}

	UserMessageViews struct {
		UserId             int64        `db:"user_id" json:"user_id"`
		PeerType           int32        `db:"peer_type" json:"peer_type"`
		PeerId             int64        `db:"peer_id" json:"peer_id"`
		PeerSeq            int64        `db:"peer_seq" json:"peer_seq"`
		CanonicalMessageId int64        `db:"canonical_message_id" json:"canonical_message_id"`
		FromUserId         int64        `db:"from_user_id" json:"from_user_id"`
		Outgoing           bool         `db:"outgoing" json:"outgoing"`
		MessageKind        int32        `db:"message_kind" json:"message_kind"`
		MessageStatus      int32        `db:"message_status" json:"message_status"`
		EditVersion        int32        `db:"edit_version" json:"edit_version"`
		Date               time.Time    `db:"date" json:"date"`
		EditDate           sql.NullTime `db:"edit_date" json:"edit_date"`
		DeletedAt          sql.NullTime `db:"deleted_at" json:"deleted_at"`
		ViewSchemaVersion  int32        `db:"view_schema_version" json:"view_schema_version"`
		ViewPayload        []byte       `db:"view_payload" json:"view_payload"`
	}
)

func newUserMessageViewsModel(db *sqlx.DB) *defaultUserMessageViewsModel {
	return &defaultUserMessageViewsModel{
		db: db,
	}
}

func (m *defaultUserMessageViewsModel) Insert2(ctx context.Context, data *UserMessageViews) (sql.Result, error) {
	tableName := "user_message_views"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, userMessageViewsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.PeerSeq, data.CanonicalMessageId, data.FromUserId, data.Outgoing, data.MessageKind, data.MessageStatus, data.EditVersion, data.Date, data.EditDate, data.DeletedAt, data.ViewSchemaVersion, data.ViewPayload)
	if err != nil {
		return nil, fmt.Errorf("user_message_views.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserMessageViewsModel) FindOneByUserIdCanonicalMessageId(ctx context.Context, userId int64, canonicalMessageId int64) (*UserMessageViews, error) {
	tableName := "user_message_views"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND canonical_message_id = ? limit 1", userMessageViewsRows, tableName)
	var resp UserMessageViews

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, canonicalMessageId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_message_views",
				Key:      fmt.Sprintf("user_id=%v,canonical_message_id=%v", userId, canonicalMessageId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_message_views.FindOneByUserIdCanonicalMessageId: %w", err)
	}

	return &resp, nil
}
