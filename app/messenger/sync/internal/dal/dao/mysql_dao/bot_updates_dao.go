/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type BotUpdatesDAO struct {
	db *sqlx.DB
}

func NewBotUpdatesDAO(db *sqlx.DB) *BotUpdatesDAO {
	return &BotUpdatesDAO{db}
}

// Insert
// insert into bot_updates(bot_id, update_id, update_type, update_data, date2) values (:bot_id, :update_id, :update_type, :update_data, :date2)
// TODO(@benqi): sqlmap
func (dao *BotUpdatesDAO) Insert(ctx context.Context, do *dataobject.BotUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_updates(bot_id, update_id, update_type, update_data, date2) values (:bot_id, :update_id, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertTx
// insert into bot_updates(bot_id, update_id, update_type, update_data, date2) values (:bot_id, :update_id, :update_type, :update_data, :date2)
// TODO(@benqi): sqlmap
func (dao *BotUpdatesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.BotUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_updates(bot_id, update_id, update_type, update_data, date2) values (:bot_id, :update_id, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// SelectByLastUpdateId
// select bot_id, update_id, update_type, update_data, date2 from bot_updates where bot_id = :bot_id order by update_id desc limit 1
// TODO(@benqi): sqlmap
func (dao *BotUpdatesDAO) SelectByLastUpdateId(ctx context.Context, bot_id int64) (rValue *dataobject.BotUpdatesDO, err error) {
	var (
		query = "select bot_id, update_id, update_type, update_data, date2 from bot_updates where bot_id = ? order by update_id desc limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, bot_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByLastUpdateId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.BotUpdatesDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByLastUpdateId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectByGtUpdateId
// select bot_id, update_id, update_type, update_data, date2 from bot_updates where bot_id = :bot_id and update_id > :update_id order by update_id asc limit :limit
// TODO(@benqi): sqlmap
func (dao *BotUpdatesDAO) SelectByGtUpdateId(ctx context.Context, bot_id int64, update_id int32, limit int32) (rList []dataobject.BotUpdatesDO, err error) {
	var (
		query = "select bot_id, update_id, update_type, update_data, date2 from bot_updates where bot_id = ? and update_id > ? order by update_id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, bot_id, update_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtUpdateId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.BotUpdatesDO
	for rows.Next() {
		v := dataobject.BotUpdatesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByGtUpdateId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByGtUpdateIdWithCB
// select bot_id, update_id, update_type, update_data, date2 from bot_updates where bot_id = :bot_id and update_id > :update_id order by update_id asc limit :limit
// TODO(@benqi): sqlmap
func (dao *BotUpdatesDAO) SelectByGtUpdateIdWithCB(ctx context.Context, bot_id int64, update_id int32, limit int32, cb func(i int, v *dataobject.BotUpdatesDO)) (rList []dataobject.BotUpdatesDO, err error) {
	var (
		query = "select bot_id, update_id, update_type, update_data, date2 from bot_updates where bot_id = ? and update_id > ? order by update_id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, bot_id, update_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtUpdateId(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.BotUpdatesDO
	for rows.Next() {
		v := dataobject.BotUpdatesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByGtUpdateId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
