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

type UserQtsUpdatesDAO struct {
	db *sqlx.DB
}

func NewUserQtsUpdatesDAO(db *sqlx.DB) *UserQtsUpdatesDAO {
	return &UserQtsUpdatesDAO{db}
}

// insert into user_qts_updates(user_id, qts, update_type, update_data, date2) values (:user_id, :qts, :update_type, :update_data, :date2)
// TODO(@benqi): sqlmap
func (dao *UserQtsUpdatesDAO) Insert(do *dataobject.UserQtsUpdatesDO) int64 {
	var query = "insert into user_qts_updates(user_id, qts, update_type, update_data, date2) values (:user_id, :qts, :update_type, :update_data, :date2)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in Insert(%v), error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in Insert(%v)_error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return id
}

// select qts from user_qts_updates where user_id = :user_id order by qts desc limit 1
// TODO(@benqi): sqlmap
func (dao *UserQtsUpdatesDAO) SelectLastQts(user_id int32) *dataobject.UserQtsUpdatesDO {
	var query = "select qts from user_qts_updates where user_id = ? order by qts desc limit 1"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectLastQts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UserQtsUpdatesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectLastQts(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectLastQts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select user_id, qts, update_type, update_data, date2 from user_qts_updates where user_id = :user_id and qts > :qts order by qts asc
// TODO(@benqi): sqlmap
func (dao *UserQtsUpdatesDAO) SelectByGtQts(user_id int32, qts int32) []dataobject.UserQtsUpdatesDO {
	var query = "select user_id, qts, update_type, update_data, date2 from user_qts_updates where user_id = ? and qts > ? order by qts asc"
	rows, err := dao.db.Queryx(query, user_id, qts)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByGtQts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserQtsUpdatesDO
	for rows.Next() {
		v := dataobject.UserQtsUpdatesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByGtQts(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByGtQts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}
