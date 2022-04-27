/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type BotCommandsDAO struct {
	db *sqlx.DB
}

func NewBotCommandsDAO(db *sqlx.DB) *BotCommandsDAO {
	return &BotCommandsDAO{db}
}

// InsertBulk
// insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)
// TODO(@benqi): sqlmap
func (dao *BotCommandsDAO) InsertBulk(ctx context.Context, doList []*dataobject.BotCommandsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = dao.db.NamedExec(ctx, query, doList)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

// InsertBulkTx
// insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)
// TODO(@benqi): sqlmap
func (dao *BotCommandsDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.BotCommandsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

// Delete
// delete from bot_commands where bot_id = :bot_id
// TODO(@benqi): sqlmap
func (dao *BotCommandsDAO) Delete(ctx context.Context, bot_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from bot_commands where bot_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, bot_id)

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
// TODO(@benqi): sqlmap
func (dao *BotCommandsDAO) DeleteTx(tx *sqlx.Tx, bot_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from bot_commands where bot_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, bot_id)

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
// TODO(@benqi): sqlmap
func (dao *BotCommandsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.BotCommandsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)
// TODO(@benqi): sqlmap
func (dao *BotCommandsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.BotCommandsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// SelectList
// select id, bot_id, command, description from bot_commands where bot_id = :bot_id
// TODO(@benqi): sqlmap
func (dao *BotCommandsDAO) SelectList(ctx context.Context, bot_id int64) (rList []dataobject.BotCommandsDO, err error) {
	var (
		query  = "select id, bot_id, command, description from bot_commands where bot_id = ?"
		values []dataobject.BotCommandsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, bot_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, bot_id, command, description from bot_commands where bot_id = :bot_id
// TODO(@benqi): sqlmap
func (dao *BotCommandsDAO) SelectListWithCB(ctx context.Context, bot_id int64, cb func(i int, v *dataobject.BotCommandsDO)) (rList []dataobject.BotCommandsDO, err error) {
	var (
		query  = "select id, bot_id, command, description from bot_commands where bot_id = ?"
		values []dataobject.BotCommandsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, bot_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// SelectListByIdList
// select id, bot_id, command, description from bot_commands where bot_id in (:id_list)
// TODO(@benqi): sqlmap
func (dao *BotCommandsDAO) SelectListByIdList(ctx context.Context, id_list []int32) (rList []dataobject.BotCommandsDO, err error) {
	var (
		query  = "select id, bot_id, command, description from bot_commands where bot_id in (?)"
		a      []interface{}
		values []dataobject.BotCommandsDO
	)
	if len(id_list) == 0 {
		rList = []dataobject.BotCommandsDO{}
		return
	}

	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByIdListWithCB
// select id, bot_id, command, description from bot_commands where bot_id in (:id_list)
// TODO(@benqi): sqlmap
func (dao *BotCommandsDAO) SelectListByIdListWithCB(ctx context.Context, id_list []int32, cb func(i int, v *dataobject.BotCommandsDO)) (rList []dataobject.BotCommandsDO, err error) {
	var (
		query  = "select id, bot_id, command, description from bot_commands where bot_id in (?)"
		a      []interface{}
		values []dataobject.BotCommandsDO
	)
	if len(id_list) == 0 {
		rList = []dataobject.BotCommandsDO{}
		return
	}

	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByIdList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}
