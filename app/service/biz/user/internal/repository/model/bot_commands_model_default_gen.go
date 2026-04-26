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
	botCommandsFieldNames          = builder.RawFieldNames(&BotCommands{})
	botCommandsRows                = strings.Join(botCommandsFieldNames, ",")
	botCommandsRowsExpectAutoSet   = strings.Join(stringx.Remove(botCommandsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	botCommandsRowsWithPlaceHolder = strings.Join(stringx.Remove(botCommandsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	botCommandsModel interface {
		Insert2(ctx context.Context, data *BotCommands) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BotCommands, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]BotCommands, error)
		Update2(ctx context.Context, data *BotCommands) error
		Delete2(ctx context.Context, id int64) error

		FindOneByBotIdCommand(ctx context.Context, botId int64, command string) (*BotCommands, error)
	}

	defaultBotCommandsModel struct {
		db *sqlx.DB
	}

	BotCommands struct {
		Id          int64  `db:"id" json:"id"`
		BotId       int64  `db:"bot_id" json:"bot_id"`
		Command     string `db:"command" json:"command"`
		Description string `db:"description" json:"description"`
	}
)

func newBotCommandsModel(db *sqlx.DB) *defaultBotCommandsModel {
	return &defaultBotCommandsModel{
		db: db,
	}
}

func (m *defaultBotCommandsModel) Insert2(ctx context.Context, data *BotCommands) (sql.Result, error) {
	query := fmt.Sprintf("insert into `bot_commands` (%s) values (?, ?, ?)", botCommandsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.BotId, data.Command, data.Description)
	if err != nil {
		return nil, fmt.Errorf("bot_commands.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultBotCommandsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `bot_commands` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("bot_commands.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultBotCommandsModel) FindOne(ctx context.Context, id int64) (*BotCommands, error) {
	query := fmt.Sprintf("select %s from bot_commands where id = ? limit 1", botCommandsRows)
	var resp BotCommands

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "bot_commands",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("bot_commands.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultBotCommandsModel) FindListByIdList(ctx context.Context, id ...int64) ([]BotCommands, error) {
	if len(id) == 0 {
		return []BotCommands{}, nil
	}

	query := fmt.Sprintf("select %s from bot_commands where id in (%s)", botCommandsRows, sqlx.InInt64List(id))

	var resp []BotCommands
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("bot_commands.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultBotCommandsModel) Update2(ctx context.Context, data *BotCommands) error {
	query := fmt.Sprintf("update `bot_commands` set %s where `id` = ?", botCommandsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.BotId, data.Command, data.Description, data.Id)
	if err != nil {
		return fmt.Errorf("bot_commands.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultBotCommandsModel) FindOneByBotIdCommand(ctx context.Context, botId int64, command string) (*BotCommands, error) {
	query := fmt.Sprintf("select %s from bot_commands where bot_id = ? AND command = ? limit 1", botCommandsRows)
	var resp BotCommands

	err := m.db.QueryRowPartial(ctx, &resp, query, botId, command)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "bot_commands",
				Key:      fmt.Sprintf("bot_id=%v,command=%v", botId, command),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("bot_commands.FindOneByBotIdCommand: %w", err)
	}

	return &resp, nil
}
