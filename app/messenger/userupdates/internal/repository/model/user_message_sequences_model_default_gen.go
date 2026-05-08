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
	userMessageSequencesFieldNames          = builder.RawFieldNames(&UserMessageSequences{})
	userMessageSequencesRows                = strings.Join(userMessageSequencesFieldNames, ",")
	userMessageSequencesRowsExpectAutoSet   = strings.Join(stringx.Remove(userMessageSequencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userMessageSequencesRowsWithPlaceHolder = strings.Join(stringx.Remove(userMessageSequencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userMessageSequencesModel interface {
		Insert2(ctx context.Context, data *UserMessageSequences) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*UserMessageSequences, error)
		FindListByUserIdList(ctx context.Context, userId ...int64) ([]UserMessageSequences, error)
		Update2(ctx context.Context, data *UserMessageSequences) error
		Delete2(ctx context.Context, userId int64) error
	}

	defaultUserMessageSequencesModel struct {
		db *sqlx.DB
	}

	UserMessageSequences struct {
		UserId            int64 `db:"user_id" json:"user_id"`
		NextUserMessageId int64 `db:"next_user_message_id" json:"next_user_message_id"`
	}
)

func newUserMessageSequencesModel(db *sqlx.DB) *defaultUserMessageSequencesModel {
	return &defaultUserMessageSequencesModel{
		db: db,
	}
}

func (m *defaultUserMessageSequencesModel) Insert2(ctx context.Context, data *UserMessageSequences) (sql.Result, error) {
	tableName := "user_message_sequences"
	query := fmt.Sprintf("insert into `%s` (%s) values (?)", tableName, userMessageSequencesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.NextUserMessageId)
	if err != nil {
		return nil, fmt.Errorf("user_message_sequences.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserMessageSequencesModel) Delete2(ctx context.Context, userId int64) error {
	tableName := "user_message_sequences"
	query := fmt.Sprintf("delete from `%s` where `user_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, userId)
	if err != nil {
		return fmt.Errorf("user_message_sequences.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserMessageSequencesModel) FindOne(ctx context.Context, userId int64) (*UserMessageSequences, error) {
	tableName := "user_message_sequences"
	query := fmt.Sprintf("select %s from %s where user_id = ? limit 1", userMessageSequencesRows, tableName)
	var resp UserMessageSequences

	err := m.db.QueryRowPartial(ctx, &resp, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_message_sequences",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_message_sequences.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUserMessageSequencesModel) FindListByUserIdList(ctx context.Context, userId ...int64) ([]UserMessageSequences, error) {
	if len(userId) == 0 {
		return []UserMessageSequences{}, nil
	}
	tableName := "user_message_sequences"

	query := fmt.Sprintf("select %s from %s where user_id in (%s)", userMessageSequencesRows, tableName, sqlx.InInt64List(userId))

	var resp []UserMessageSequences
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []UserMessageSequences{}, nil
		}
		return nil, fmt.Errorf("user_message_sequences.FindListByUserIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUserMessageSequencesModel) Update2(ctx context.Context, data *UserMessageSequences) error {
	tableName := "user_message_sequences"
	query := fmt.Sprintf("update `%s` set %s where `user_id` = ?", tableName, userMessageSequencesRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.NextUserMessageId, data.UserId)
	if err != nil {
		return fmt.Errorf("user_message_sequences.Update2 exec: %w", err)
	}

	return nil
}
