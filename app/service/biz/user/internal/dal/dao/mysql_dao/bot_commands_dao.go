/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type BotCommandsDAO struct {
	db *sqlx.DB
}

func NewBotCommandsDAO(db *sqlx.DB) *BotCommandsDAO {
	return &BotCommandsDAO{
		db: db,
	}
}

// InsertBulk
// insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)
func (dao *BotCommandsDAO) InsertBulk(ctx context.Context, doList []*dataobject.BotCommandsDO) (lastInsertId, rowsAffected int64, err error) {
	if len(doList) == 0 {
		return
	}

	var (
		query string
		r     sql.Result
	)
	query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"

	r, err = dao.db.NamedExec(ctx, query, doList)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertBulk(%v), error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertBulk(%v), error: %v", doList, err)
	}

	return
}

// InsertBulkTx
// insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)
func (dao *BotCommandsDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.BotCommandsDO) (lastInsertId, rowsAffected int64, err error) {
	if len(doList) == 0 {
		return
	}

	var (
		query string
		r     sql.Result
	)
	query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertBulk(%v), error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertBulk(%v), error: %v", doList, err)
	}

	return
}

// Delete
// delete from bot_commands where bot_id = :bot_id
func (dao *BotCommandsDAO) Delete(ctx context.Context, botId int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "delete from bot_commands where bot_id = ?"

	rResult, err = dao.db.Exec(ctx, query, botId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

// DeleteTx
// delete from bot_commands where bot_id = :bot_id
func (dao *BotCommandsDAO) DeleteTx(tx *sqlx.Tx, botId int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "delete from bot_commands where bot_id = ?"

	rResult, err = tx.Exec(query, botId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

// InsertOrUpdate
// insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)
func (dao *BotCommandsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.BotCommandsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v), error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)
func (dao *BotCommandsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.BotCommandsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v), error: %v", do, err)
	}

	return
}

// SelectList
// select id, bot_id, command, description from bot_commands where bot_id = :bot_id
func (dao *BotCommandsDAO) SelectList(ctx context.Context, botId int64) (rList []dataobject.BotCommandsDO, err error) {
	var (
		query  string
		values []dataobject.BotCommandsDO
	)
	query = "select id, bot_id, command, description from bot_commands where bot_id = ?"

	err = dao.db.QueryRowsPartial(ctx, &values, query, botId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, bot_id, command, description from bot_commands where bot_id = :bot_id
func (dao *BotCommandsDAO) SelectListWithCB(ctx context.Context, botId int64, cb func(sz, i int, v *dataobject.BotCommandsDO)) (rList []dataobject.BotCommandsDO, err error) {
	var (
		query  string
		values []dataobject.BotCommandsDO
	)
	query = "select id, bot_id, command, description from bot_commands where bot_id = ?"

	err = dao.db.QueryRowsPartial(ctx, &values, query, botId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// SelectListByIdList
// select id, bot_id, command, description from bot_commands where bot_id in (:id_list)
func (dao *BotCommandsDAO) SelectListByIdList(ctx context.Context, idList []int32) (rList []dataobject.BotCommandsDO, err error) {
	if len(idList) == 0 {
		rList = []dataobject.BotCommandsDO{}
		return
	}

	var (
		query  string
		values []dataobject.BotCommandsDO
	)
	query = fmt.Sprintf("select id, bot_id, command, description from bot_commands where bot_id in (%s)", sqlx.InInt32List(idList))

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByIdListWithCB
// select id, bot_id, command, description from bot_commands where bot_id in (:id_list)
func (dao *BotCommandsDAO) SelectListByIdListWithCB(ctx context.Context, idList []int32, cb func(sz, i int, v *dataobject.BotCommandsDO)) (rList []dataobject.BotCommandsDO, err error) {
	if len(idList) == 0 {
		rList = []dataobject.BotCommandsDO{}
		return
	}

	var (
		query  string
		values []dataobject.BotCommandsDO
	)
	query = fmt.Sprintf("select id, bot_id, command, description from bot_commands where bot_id in (%s)", sqlx.InInt32List(idList))

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByIdList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, &rList[i])
		}
	}

	return
}
