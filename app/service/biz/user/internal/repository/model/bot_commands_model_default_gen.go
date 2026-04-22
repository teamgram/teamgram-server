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

	cacheTBotCommandsIdPrefix = "cache:t:bot_commands:id:"

	cacheBotCommandsIdPrefix = "cache#BotCommands#id"

	cacheBotCommandsBotIdCommandPrefix = "cache#BotId#Command"
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
	return m.db.Exec(ctx, query, data.BotId, data.Command, data.Description)
}

func (m *defaultBotCommandsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `bot_commands` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultBotCommandsModel) FindOne(ctx context.Context, id int64) (*BotCommands, error) {
	query := fmt.Sprintf("select %s from bot_commands where id = ? limit 1", botCommandsRows)
	var resp BotCommands
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return resp, nil
}

func (m *defaultBotCommandsModel) Update2(ctx context.Context, data *BotCommands) error {
	query := fmt.Sprintf("update `bot_commands` set %s where `id` = ?", botCommandsRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.BotId, data.Command, data.Description, data.Id)
	return err
}

func (m *defaultBotCommandsModel) FindOneByBotIdCommand(ctx context.Context, botId int64, command string) (*BotCommands, error) {
	query := fmt.Sprintf("select %s from bot_commands where bot_id = ? AND command = ? limit 1", botCommandsRows)
	var resp BotCommands
	err := m.db.QueryRowPartial(ctx, &resp, query, botId, command)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultBotCommandsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheBotCommandsIdPrefix, primary)
}

func (m *defaultBotCommandsModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from bot_commands where id = ? limit 1", botCommandsRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
