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
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB

type (
	bizBotCommandsModel interface {
		InsertBulk(ctx context.Context, doList []*BotCommands) (lastInsertId, rowsAffected int64, err error)
		InsertBulkTx(tx *sqlx.Tx, doList []*BotCommands) (lastInsertId, rowsAffected int64, err error)

		Delete(ctx context.Context, botId int64) (rowsAffected int64, err error)
		DeleteTx(tx *sqlx.Tx, botId int64) (rowsAffected int64, err error)

		InsertOrUpdate(ctx context.Context, data *BotCommands) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *BotCommands) (lastInsertId, rowsAffected int64, err error)

		SelectList(ctx context.Context, botId int64) ([]BotCommands, error)
		SelectListWithCB(ctx context.Context, botId int64, cb func(sz, i int, v *BotCommands)) ([]BotCommands, error)

		SelectListByIdList(ctx context.Context, idList []int32) ([]BotCommands, error)
		SelectListByIdListWithCB(ctx context.Context, idList []int32, cb func(sz, i int, v *BotCommands)) ([]BotCommands, error)
	}
)

// InsertBulk
// insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)
func (m *defaultBotCommandsModel) InsertBulk(ctx context.Context, doList []*BotCommands) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = m.db.NamedExec(ctx, query, doList)
	if err != nil {
		err = fmt.Errorf("bot_commands.InsertBulk named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("bot_commands.InsertBulk last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("bot_commands.InsertBulk rows affected: %w", err)
	}

	return
}

// InsertBulkTx
// insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)
func (m *defaultBotCommandsModel) InsertBulkTx(tx *sqlx.Tx, doList []*BotCommands) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		err = fmt.Errorf("bot_commands.InsertBulkTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("bot_commands.InsertBulkTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("bot_commands.InsertBulkTx rows affected: %w", err)
	}

	return
}

// Delete
// delete from bot_commands where bot_id = :bot_id
func (m *defaultBotCommandsModel) Delete(ctx context.Context, botId int64) (rowsAffected int64, err error) {

	var (
		query   = "delete from bot_commands where bot_id = ?"
		rResult sql.Result
	)
	rResult, err = m.db.Exec(ctx, query, botId)

	if err != nil {
		err = fmt.Errorf("bot_commands.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("bot_commands.Delete rows affected: %w", err)
	}

	return
}

// DeleteTx
// delete from bot_commands where bot_id = :bot_id
func (m *defaultBotCommandsModel) DeleteTx(tx *sqlx.Tx, botId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from bot_commands where bot_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, botId)

	if err != nil {
		err = fmt.Errorf("bot_commands.DeleteTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("bot_commands.DeleteTx rows affected: %w", err)
	}

	return
}

// InsertOrUpdate
// insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)
func (m *defaultBotCommandsModel) InsertOrUpdate(ctx context.Context, data *BotCommands) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("bot_commands.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("bot_commands.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("bot_commands.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdateTx
// insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)
func (m *defaultBotCommandsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *BotCommands) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("bot_commands.InsertOrUpdateTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("bot_commands.InsertOrUpdateTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("bot_commands.InsertOrUpdateTx rows affected: %w", err)
	}

	return
}

// SelectList
// select id, bot_id, command, description from bot_commands where bot_id = :bot_id
func (m *defaultBotCommandsModel) SelectList(ctx context.Context, botId int64) (rList []BotCommands, err error) {
	var (
		query  = "select id, bot_id, command, description from bot_commands where bot_id = ?"
		values []BotCommands
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, botId)

	if err != nil {
		err = fmt.Errorf("bot_commands.SelectList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, bot_id, command, description from bot_commands where bot_id = :bot_id
func (m *defaultBotCommandsModel) SelectListWithCB(ctx context.Context, botId int64, cb func(sz, i int, v *BotCommands)) (rList []BotCommands, err error) {
	var (
		query  = "select id, bot_id, command, description from bot_commands where bot_id = ?"
		values []BotCommands
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, botId)

	if err != nil {
		err = fmt.Errorf("bot_commands.SelectListWithCB: %w", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// SelectListByIdList
// select id, bot_id, command, description from bot_commands where bot_id in (:id_list)
func (m *defaultBotCommandsModel) SelectListByIdList(ctx context.Context, idList []int32) (rList []BotCommands, err error) {
	var (
		query  = fmt.Sprintf("select id, bot_id, command, description from bot_commands where bot_id in (%s)", sqlx.InInt32List(idList))
		values []BotCommands
	)
	if len(idList) == 0 {
		rList = []BotCommands{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		err = fmt.Errorf("bot_commands.SelectListByIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByIdListWithCB
// select id, bot_id, command, description from bot_commands where bot_id in (:id_list)
func (m *defaultBotCommandsModel) SelectListByIdListWithCB(ctx context.Context, idList []int32, cb func(sz, i int, v *BotCommands)) (rList []BotCommands, err error) {
	var (
		query  = fmt.Sprintf("select id, bot_id, command, description from bot_commands where bot_id in (%s)", sqlx.InInt32List(idList))
		values []BotCommands
	)
	if len(idList) == 0 {
		rList = []BotCommands{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		err = fmt.Errorf("bot_commands.SelectListByIdListWithCB: %w", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}
