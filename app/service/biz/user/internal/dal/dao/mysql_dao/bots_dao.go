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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type BotsDAO struct {
	db *sqlx.DB
}

func NewBotsDAO(db *sqlx.DB) *BotsDAO {
	return &BotsDAO{db}
}

// Select
// select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id = :bot_id
// TODO(@benqi): sqlmap
func (dao *BotsDAO) Select(ctx context.Context, bot_id int64) (rValue *dataobject.BotsDO, err error) {
	var (
		query = "select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, bot_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.BotsDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in Select(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectByToken
// select bot_id from bots where token = :token
// TODO(@benqi): sqlmap
func (dao *BotsDAO) SelectByToken(ctx context.Context, token string) (rValue int64, err error) {
	var query = "select bot_id from bots where token = ?"
	err = dao.db.Get(ctx, &rValue, query, token)

	if err != nil {
		logx.WithContext(ctx).Errorf("get in SelectByToken(_), error: %v", err)
	}

	return
}

// SelectByIdList
// select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id in (:id_list)
// TODO(@benqi): sqlmap
func (dao *BotsDAO) SelectByIdList(ctx context.Context, id_list []int32) (rList []dataobject.BotsDO, err error) {
	var (
		query = "select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(id_list) == 0 {
		rList = []dataobject.BotsDO{}
		return
	}

	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.BotsDO
	for rows.Next() {
		v := dataobject.BotsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByIdListWithCB
// select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id in (:id_list)
// TODO(@benqi): sqlmap
func (dao *BotsDAO) SelectByIdListWithCB(ctx context.Context, id_list []int32, cb func(i int, v *dataobject.BotsDO)) (rList []dataobject.BotsDO, err error) {
	var (
		query = "select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(id_list) == 0 {
		rList = []dataobject.BotsDO{}
		return
	}

	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
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

	var values []dataobject.BotsDO
	for rows.Next() {
		v := dataobject.BotsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
