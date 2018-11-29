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

type UserPrivacysDAO struct {
	db *sqlx.DB
}

func NewUserPrivacysDAO(db *sqlx.DB) *UserPrivacysDAO {
	return &UserPrivacysDAO{db}
}

// insert into user_privacys(user_id, key_type, rules) values (:user_id, :key_type, :rules)
// TODO(@benqi): sqlmap
func (dao *UserPrivacysDAO) Insert(do *dataobject.UserPrivacysDO) int64 {
	var query = "insert into user_privacys(user_id, key_type, rules) values (:user_id, :key_type, :rules)"
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

// update user_privacys set rules = :rules where user_id = :user_id and :key_type
// TODO(@benqi): sqlmap
func (dao *UserPrivacysDAO) UpdatePrivacy(rules string, user_id int32, key_type int8) int64 {
	var query = "update user_privacys set rules = ? where user_id = ? and ?"
	r, err := dao.db.Exec(query, rules, user_id, key_type)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdatePrivacy(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdatePrivacy(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// select id, user_id, key_type, rules from user_privacys where user_id = :user_id and key_type = :key_type
// TODO(@benqi): sqlmap
func (dao *UserPrivacysDAO) SelectPrivacy(user_id int32, key_type int8) *dataobject.UserPrivacysDO {
	var query = "select id, user_id, key_type, rules from user_privacys where user_id = ? and key_type = ?"
	rows, err := dao.db.Queryx(query, user_id, key_type)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectPrivacy(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UserPrivacysDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectPrivacy(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectPrivacy(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}
