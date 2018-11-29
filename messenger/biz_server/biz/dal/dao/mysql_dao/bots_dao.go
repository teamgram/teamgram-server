// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package mysql_dao

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
)

type BotsDAO struct {
	db *sqlx.DB
}

func NewBotsDAO(db *sqlx.DB) *BotsDAO {
	return &BotsDAO{db}
}

// insert into bots(bot_id, bot_type, description) values (:bot_id, :bot_type, :description)
// TODO(@benqi): sqlmap
func (dao *BotsDAO) InsertOrUpdate(do *dataobject.BotsDO) int64 {
	var query = "insert into bots(bot_id, bot_type, description) values (:bot_id, :bot_type, :description)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in InsertOrUpdate(%v), error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return id
}

// select id, bot_id, bot_type, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id = :bot_id
// TODO(@benqi): sqlmap
func (dao *BotsDAO) Select(bot_id int32) *dataobject.BotsDO {
	var query = "select id, bot_id, bot_type, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id = ?"
	rows, err := dao.db.Queryx(query, bot_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in Select(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.BotsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in Select(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in Select(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select id, bot_id, bot_type, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id in (:id_list)
// TODO(@benqi): sqlmap
func (dao *BotsDAO) SelectByIdList(id_list []int32) []dataobject.BotsDO {
	var q = "select id, bot_id, bot_type, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id in (?)"
	query, a, err := sqlx.In(q, id_list)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.BotsDO
	for rows.Next() {
		v := dataobject.BotsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}
