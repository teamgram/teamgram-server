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

		FindOneByUserIdPeerDialogIdReadUserIdReadOutboxMaxId(ctx context.Context, userId int64, peerDialogId int64, readUserId int64, readOutboxMaxId int32) (*MessageReadOutbox, error)
	}

	defaultMessageReadOutboxModel struct {
		db *sqlx.DB
	}

	MessageReadOutbox struct {
		Id                int64 `db:"id" json:"id"`
		UserId            int64 `db:"user_id" json:"user_id"`
		PeerDialogId      int64 `db:"peer_dialog_id" json:"peer_dialog_id"`
		ReadUserId        int64 `db:"read_user_id" json:"read_user_id"`
		ReadOutboxMaxId   int32 `db:"read_outbox_max_id" json:"read_outbox_max_id"`
		ReadOutboxMaxDate int64 `db:"read_outbox_max_date" json:"read_outbox_max_date"`
	}
)

func newMessageReadOutboxModel(db *sqlx.DB) *defaultMessageReadOutboxModel {
	return &defaultMessageReadOutboxModel{
		db: db,
	}
}

func (m *defaultMessageReadOutboxModel) Insert2(ctx context.Context, data *MessageReadOutbox) (sql.Result, error) {
	query := fmt.Sprintf("insert into `message_read_outbox` (%s) values (?, ?, ?, ?, ?)", messageReadOutboxRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerDialogId, data.ReadUserId, data.ReadOutboxMaxId, data.ReadOutboxMaxDate)
	if err != nil {
		return nil, fmt.Errorf("message_read_outbox.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultMessageReadOutboxModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `message_read_outbox` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("message_read_outbox.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultMessageReadOutboxModel) FindOne(ctx context.Context, id int64) (*MessageReadOutbox, error) {
	query := fmt.Sprintf("select %s from message_read_outbox where id = ? limit 1", messageReadOutboxRows)
	var resp MessageReadOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("message_read_outbox.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultMessageReadOutboxModel) FindListByIdList(ctx context.Context, id ...int64) ([]MessageReadOutbox, error) {
	if len(id) == 0 {
		return []MessageReadOutbox{}, nil
	}

	query := fmt.Sprintf("select %s from message_read_outbox where id in (%s)", messageReadOutboxRows, sqlx.InInt64List(id))

	var resp []MessageReadOutbox
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("message_read_outbox.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultMessageReadOutboxModel) Update2(ctx context.Context, data *MessageReadOutbox) error {
	query := fmt.Sprintf("update `message_read_outbox` set %s where `id` = ?", messageReadOutboxRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.PeerDialogId, data.ReadUserId, data.ReadOutboxMaxId, data.ReadOutboxMaxDate, data.Id)
	if err != nil {
		return fmt.Errorf("message_read_outbox.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultMessageReadOutboxModel) FindOneByUserIdPeerDialogIdReadUserIdReadOutboxMaxId(ctx context.Context, userId int64, peerDialogId int64, readUserId int64, readOutboxMaxId int32) (*MessageReadOutbox, error) {
	query := fmt.Sprintf("select %s from message_read_outbox where user_id = ? AND peer_dialog_id = ? AND read_user_id = ? AND read_outbox_max_id = ? limit 1", messageReadOutboxRows)
	var resp MessageReadOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerDialogId, readUserId, readOutboxMaxId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("message_read_outbox.FindOneByUserIdPeerDialogIdReadUserIdReadOutboxMaxId: %w", err)
	}

	return &resp, nil
}
