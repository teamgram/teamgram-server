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
	messageReadOutboxFieldNames          = builder.RawFieldNames(&MessageReadOutbox{})
	messageReadOutboxRows                = strings.Join(messageReadOutboxFieldNames, ",")
	messageReadOutboxRowsExpectAutoSet   = strings.Join(stringx.Remove(messageReadOutboxFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	messageReadOutboxRowsWithPlaceHolder = strings.Join(stringx.Remove(messageReadOutboxFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	messageReadOutboxModel interface {
		Insert2(ctx context.Context, data *MessageReadOutbox) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*MessageReadOutbox, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]MessageReadOutbox, error)
		Update2(ctx context.Context, data *MessageReadOutbox) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserIdPeerTypePeerIdReadUserIdReadOutboxMaxId(ctx context.Context, userId int64, peerType int32, peerId int64, readUserId int64, readOutboxMaxId int32) (*MessageReadOutbox, error)
	}

	defaultMessageReadOutboxModel struct {
		db *sqlx.DB
	}

	MessageReadOutbox struct {
		Id                int64     `db:"id" json:"id"`
		UserId            int64     `db:"user_id" json:"user_id"`
		PeerType          int32     `db:"peer_type" json:"peer_type"`
		PeerId            int64     `db:"peer_id" json:"peer_id"`
		ReadUserId        int64     `db:"read_user_id" json:"read_user_id"`
		ReadOutboxMaxId   int32     `db:"read_outbox_max_id" json:"read_outbox_max_id"`
		ReadOutboxMaxDate time.Time `db:"read_outbox_max_date" json:"read_outbox_max_date"`
	}
)

func newMessageReadOutboxModel(db *sqlx.DB) *defaultMessageReadOutboxModel {
	return &defaultMessageReadOutboxModel{
		db: db,
	}
}

func (m *defaultMessageReadOutboxModel) Insert2(ctx context.Context, data *MessageReadOutbox) (sql.Result, error) {
	tableName := "message_read_outbox"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?)", tableName, messageReadOutboxRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.ReadUserId, data.ReadOutboxMaxId, data.ReadOutboxMaxDate)
	if err != nil {
		return nil, fmt.Errorf("message_read_outbox.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultMessageReadOutboxModel) Delete2(ctx context.Context, id int64) error {
	tableName := "message_read_outbox"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("message_read_outbox.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultMessageReadOutboxModel) FindOne(ctx context.Context, id int64) (*MessageReadOutbox, error) {
	tableName := "message_read_outbox"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", messageReadOutboxRows, tableName)
	var resp MessageReadOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_read_outbox",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("message_read_outbox.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultMessageReadOutboxModel) FindListByIdList(ctx context.Context, id ...int64) ([]MessageReadOutbox, error) {
	if len(id) == 0 {
		return []MessageReadOutbox{}, nil
	}
	tableName := "message_read_outbox"

	query := fmt.Sprintf("select %s from %s where id in (%s)", messageReadOutboxRows, tableName, sqlx.InInt64List(id))

	var resp []MessageReadOutbox
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []MessageReadOutbox{}, nil
		}
		return nil, fmt.Errorf("message_read_outbox.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultMessageReadOutboxModel) Update2(ctx context.Context, data *MessageReadOutbox) error {
	tableName := "message_read_outbox"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, messageReadOutboxRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.ReadUserId, data.ReadOutboxMaxId, data.ReadOutboxMaxDate, data.Id)
	if err != nil {
		return fmt.Errorf("message_read_outbox.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultMessageReadOutboxModel) FindOneByUserIdPeerTypePeerIdReadUserIdReadOutboxMaxId(ctx context.Context, userId int64, peerType int32, peerId int64, readUserId int64, readOutboxMaxId int32) (*MessageReadOutbox, error) {
	tableName := "message_read_outbox"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND peer_type = ? AND peer_id = ? AND read_user_id = ? AND read_outbox_max_id = ? limit 1", messageReadOutboxRows, tableName)
	var resp MessageReadOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerType, peerId, readUserId, readOutboxMaxId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_read_outbox",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v,read_user_id=%v,read_outbox_max_id=%v", userId, peerType, peerId, readUserId, readOutboxMaxId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("message_read_outbox.FindOneByUserIdPeerTypePeerIdReadUserIdReadOutboxMaxId: %w", err)
	}

	return &resp, nil
}
