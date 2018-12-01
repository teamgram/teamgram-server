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

type UserBlocksDAO struct {
	db *sqlx.DB
}

func NewUserBlocksDAO(db *sqlx.DB) *UserBlocksDAO {
	return &UserBlocksDAO{db}
}

// insert into user_blocks(user_id, block_id, `date`) values (:user_id, :block_id, :date) on duplicate key update `date` = values(`date`), deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserBlocksDAO) InsertOrUpdate(do *dataobject.UserBlocksDO) int64 {
	var query = "insert into user_blocks(user_id, block_id, `date`) values (:user_id, :block_id, :date) on duplicate key update `date` = values(`date`), deleted = 0"
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

// select user_id, block_id from user_blocks where user_id = :user_id and deleted = 0 order by id asc limit :limit
// TODO(@benqi): sqlmap
func (dao *UserBlocksDAO) SelectList(user_id int32, limit int32) []dataobject.UserBlocksDO {
	var query = "select user_id, block_id from user_blocks where user_id = ? and deleted = 0 order by id asc limit ?"
	rows, err := dao.db.Queryx(query, user_id, limit)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserBlocksDO
	for rows.Next() {
		v := dataobject.UserBlocksDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select user_id, block_id from user_blocks where user_id = :user_id and block_id = :block_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserBlocksDAO) Select(user_id int32, block_id int32) *dataobject.UserBlocksDO {
	var query = "select user_id, block_id from user_blocks where user_id = ? and block_id = ? and deleted = 0"
	rows, err := dao.db.Queryx(query, user_id, block_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in Select(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UserBlocksDO{}
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

// update user_blocks set deleted = 1, `date` = 0 where user_id = :user_id and block_id = :block_id
// TODO(@benqi): sqlmap
func (dao *UserBlocksDAO) Delete(user_id int32, block_id int32) int64 {
	var query = "update user_blocks set deleted = 1, `date` = 0 where user_id = ? and block_id = ?"
	r, err := dao.db.Exec(query, user_id, block_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in Delete(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in Delete(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}
