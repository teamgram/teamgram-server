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

type BannedDAO struct {
	db *sqlx.DB
}

func NewBannedDAO(db *sqlx.DB) *BannedDAO {
	return &BannedDAO{db}
}

// insert into banned(phone, banned_time, expires, banned_reason, log, state) values (:phone, :banned_time, :expires, :banned_reason, :log, :state) on duplicate key update banned_time = values(banned_time), expires = values(expires), banned_reason = values(banned_reason), log = values(log), state = values(state)
// TODO(@benqi): sqlmap
func (dao *BannedDAO) InsertOrUpdate(do *dataobject.BannedDO) int64 {
	var query = "insert into banned(phone, banned_time, expires, banned_reason, log, state) values (:phone, :banned_time, :expires, :banned_reason, :log, :state) on duplicate key update banned_time = values(banned_time), expires = values(expires), banned_reason = values(banned_reason), log = values(log), state = values(state)"
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

// select id from banned where phone = :phone and state > 0
// TODO(@benqi): sqlmap
func (dao *BannedDAO) CheckBannedByPhone(phone string) *dataobject.BannedDO {
	var query = "select id from banned where phone = ? and state > 0"
	rows, err := dao.db.Queryx(query, phone)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in CheckBannedByPhone(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.BannedDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in CheckBannedByPhone(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in CheckBannedByPhone(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// update banned set expires = :expires, state = 0, log = :log, state = :state where phone = :phone
// TODO(@benqi): sqlmap
func (dao *BannedDAO) Update(expires int64, log string, state int8, phone string) int64 {
	var query = "update banned set expires = ?, state = 0, log = ?, state = ? where phone = ?"
	r, err := dao.db.Exec(query, expires, log, state, phone)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in Update(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in Update(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}
