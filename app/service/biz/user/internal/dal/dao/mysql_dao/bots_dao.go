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
	"fmt"
	"strings"

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
		do    = &dataobject.BotsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, bot_id)

	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectByToken
// select bot_id from bots where token = :token
// TODO(@benqi): sqlmap
func (dao *BotsDAO) SelectByToken(ctx context.Context, token string) (rValue int64, err error) {
	var query = "select bot_id from bots where token = ?"
	err = dao.db.QueryRowPartial(ctx, &rValue, query, token)

	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.WithContext(ctx).Errorf("get in SelectByToken(_), error: %v", err)
			return
		} else {
			err = nil
		}
	}

	return
}

// SelectByIdList
// select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id in (:id_list)
// TODO(@benqi): sqlmap
func (dao *BotsDAO) SelectByIdList(ctx context.Context, id_list []int32) (rList []dataobject.BotsDO, err error) {
	var (
		query  = "select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id in (?)"
		a      []interface{}
		values []dataobject.BotsDO
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
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByIdListWithCB
// select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id in (:id_list)
// TODO(@benqi): sqlmap
func (dao *BotsDAO) SelectByIdListWithCB(ctx context.Context, id_list []int32, cb func(i int, v *dataobject.BotsDO)) (rList []dataobject.BotsDO, err error) {
	var (
		query  = "select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id in (?)"
		a      []interface{}
		values []dataobject.BotsDO
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
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
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

// Update
// update bots set %s where bot_id = :bot_id
// TODO(@benqi): sqlmap
func (dao *BotsDAO) Update(ctx context.Context, cMap map[string]interface{}, bot_id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update bots set %s where bot_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, bot_id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

// UpdateTx
// update bots set %s where bot_id = :bot_id
// TODO(@benqi): sqlmap
func (dao *BotsDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, bot_id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update bots set %s where bot_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, bot_id)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}
