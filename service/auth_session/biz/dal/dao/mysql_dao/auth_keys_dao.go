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
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/service/auth_session/biz/dal/dataobject"
)

type AuthKeysDAO struct {
	db *sqlx.DB
}

func NewAuthKeysDAO(db *sqlx.DB) *AuthKeysDAO {
	return &AuthKeysDAO{db}
}

// insert into auth_keys(auth_key_id, body) values (:auth_key_id, :body)
// TODO(@benqi): sqlmap
func (dao *AuthKeysDAO) Insert(do *dataobject.AuthKeysDO) int64 {
	var query = "insert into auth_keys(auth_key_id, body) values (:auth_key_id, :body)"
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

// select auth_key_id, body from auth_keys where auth_key_id = :auth_key_id
// TODO(@benqi): sqlmap
func (dao *AuthKeysDAO) SelectByAuthKeyId(auth_key_id int64) *dataobject.AuthKeysDO {
	var query = "select auth_key_id, body from auth_keys where auth_key_id = ?"
	rows, err := dao.db.Queryx(query, auth_key_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByAuthKeyId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.AuthKeysDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByAuthKeyId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByAuthKeyId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}
