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
	chatCreateOperationsFieldNames          = builder.RawFieldNames(&ChatCreateOperations{})
	chatCreateOperationsRows                = strings.Join(chatCreateOperationsFieldNames, ",")
	chatCreateOperationsRowsExpectAutoSet   = strings.Join(stringx.Remove(chatCreateOperationsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	chatCreateOperationsRowsWithPlaceHolder = strings.Join(stringx.Remove(chatCreateOperationsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	chatCreateOperationsModel interface {
		Insert2(ctx context.Context, data *ChatCreateOperations) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ChatCreateOperations, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]ChatCreateOperations, error)
		Update2(ctx context.Context, data *ChatCreateOperations) error
		Delete2(ctx context.Context, id int64) error

		FindOneByReplayKey(ctx context.Context, replayKey string) (*ChatCreateOperations, error)
		FindListByReplayKeyList(ctx context.Context, replayKey ...string) ([]ChatCreateOperations, error)

		FindOneByOperationId(ctx context.Context, operationId string) (*ChatCreateOperations, error)
		FindListByOperationIdList(ctx context.Context, operationId ...string) ([]ChatCreateOperations, error)
	}

	defaultChatCreateOperationsModel struct {
		db *sqlx.DB
	}

	ChatCreateOperations struct {
		Id                  int64  `db:"id" json:"id"`
		OperationId         string `db:"operation_id" json:"operation_id"`
		ReplayKey           string `db:"replay_key" json:"replay_key"`
		ActorUserId         int64  `db:"actor_user_id" json:"actor_user_id"`
		ClientMsgId         int64  `db:"client_msg_id" json:"client_msg_id"`
		Title               string `db:"title" json:"title"`
		InviteeIds          string `db:"invitee_ids" json:"invitee_ids"`
		TtlPeriod           int32  `db:"ttl_period" json:"ttl_period"`
		ChatId              int64  `db:"chat_id" json:"chat_id"`
		ParticipantsVersion int32  `db:"participants_version" json:"participants_version"`
		Status              int32  `db:"status" json:"status"`
		Date                int64  `db:"date" json:"date"`
		UpdatedAtSec        int64  `db:"updated_at_sec" json:"updated_at_sec"`
		ExpiresAt           int64  `db:"expires_at" json:"expires_at"`
	}
)

func newChatCreateOperationsModel(db *sqlx.DB) *defaultChatCreateOperationsModel {
	return &defaultChatCreateOperationsModel{
		db: db,
	}
}

func (m *defaultChatCreateOperationsModel) Insert2(ctx context.Context, data *ChatCreateOperations) (sql.Result, error) {
	tableName := "chat_create_operations"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, chatCreateOperationsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.OperationId, data.ReplayKey, data.ActorUserId, data.ClientMsgId, data.Title, data.InviteeIds, data.TtlPeriod, data.ChatId, data.ParticipantsVersion, data.Status, data.Date, data.UpdatedAtSec, data.ExpiresAt)
	if err != nil {
		return nil, fmt.Errorf("chat_create_operations.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultChatCreateOperationsModel) Delete2(ctx context.Context, id int64) error {
	tableName := "chat_create_operations"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("chat_create_operations.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultChatCreateOperationsModel) FindOne(ctx context.Context, id int64) (*ChatCreateOperations, error) {
	tableName := "chat_create_operations"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", chatCreateOperationsRows, tableName)
	var resp ChatCreateOperations

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chat_create_operations",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("chat_create_operations.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultChatCreateOperationsModel) FindListByIdList(ctx context.Context, id ...int64) ([]ChatCreateOperations, error) {
	if len(id) == 0 {
		return []ChatCreateOperations{}, nil
	}
	tableName := "chat_create_operations"

	query := fmt.Sprintf("select %s from %s where id in (%s)", chatCreateOperationsRows, tableName, sqlx.InInt64List(id))

	var resp []ChatCreateOperations
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []ChatCreateOperations{}, nil
		}
		return nil, fmt.Errorf("chat_create_operations.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultChatCreateOperationsModel) Update2(ctx context.Context, data *ChatCreateOperations) error {
	tableName := "chat_create_operations"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, chatCreateOperationsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.OperationId, data.ReplayKey, data.ActorUserId, data.ClientMsgId, data.Title, data.InviteeIds, data.TtlPeriod, data.ChatId, data.ParticipantsVersion, data.Status, data.Date, data.UpdatedAtSec, data.ExpiresAt, data.Id)
	if err != nil {
		return fmt.Errorf("chat_create_operations.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultChatCreateOperationsModel) FindOneByReplayKey(ctx context.Context, replayKey string) (*ChatCreateOperations, error) {
	tableName := "chat_create_operations"
	query := fmt.Sprintf("select %s from %s where replay_key = ? limit 1", chatCreateOperationsRows, tableName)
	var resp ChatCreateOperations

	err := m.db.QueryRowPartial(ctx, &resp, query, replayKey)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chat_create_operations",
				Key:      fmt.Sprintf("replay_key=%v", replayKey),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("chat_create_operations.FindOneByReplayKey: %w", err)
	}

	return &resp, nil
}

func (m *defaultChatCreateOperationsModel) FindListByReplayKeyList(ctx context.Context, replayKey ...string) ([]ChatCreateOperations, error) {
	if len(replayKey) == 0 {
		return []ChatCreateOperations{}, nil
	}
	tableName := "chat_create_operations"

	query := fmt.Sprintf("select %s from %s where replay_key in (%s)", chatCreateOperationsRows, tableName, sqlx.InStringList(replayKey))
	var resp []ChatCreateOperations
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []ChatCreateOperations{}, nil
		}
		return nil, fmt.Errorf("chat_create_operations.FindListByReplayKeyList: %w", err)
	}

	return resp, nil
}

func (m *defaultChatCreateOperationsModel) FindOneByOperationId(ctx context.Context, operationId string) (*ChatCreateOperations, error) {
	tableName := "chat_create_operations"
	query := fmt.Sprintf("select %s from %s where operation_id = ? limit 1", chatCreateOperationsRows, tableName)
	var resp ChatCreateOperations

	err := m.db.QueryRowPartial(ctx, &resp, query, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chat_create_operations",
				Key:      fmt.Sprintf("operation_id=%v", operationId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("chat_create_operations.FindOneByOperationId: %w", err)
	}

	return &resp, nil
}

func (m *defaultChatCreateOperationsModel) FindListByOperationIdList(ctx context.Context, operationId ...string) ([]ChatCreateOperations, error) {
	if len(operationId) == 0 {
		return []ChatCreateOperations{}, nil
	}
	tableName := "chat_create_operations"

	query := fmt.Sprintf("select %s from %s where operation_id in (%s)", chatCreateOperationsRows, tableName, sqlx.InStringList(operationId))
	var resp []ChatCreateOperations
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []ChatCreateOperations{}, nil
		}
		return nil, fmt.Errorf("chat_create_operations.FindListByOperationIdList: %w", err)
	}

	return resp, nil
}
