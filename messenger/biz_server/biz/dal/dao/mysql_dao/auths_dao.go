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

type AuthsDAO struct {
	db *sqlx.DB
}

func NewAuthsDAO(db *sqlx.DB) *AuthsDAO {
	return &AuthsDAO{db}
}

// insert into auths(auth_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, connection_hash) values (:auth_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :connection_hash)
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) Insert(do *dataobject.AuthsDO) int64 {
	var query = "insert into auths(auth_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, connection_hash) values (:auth_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :connection_hash)"
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

// select connection_hash from auths where auth_id = :auth_id
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) SelectConnectionHashByAuthId(auth_id int64) *dataobject.AuthsDO {
	var query = "select connection_hash from auths where auth_id = ?"
	rows, err := dao.db.Queryx(query, auth_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectConnectionHashByAuthId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.AuthsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectConnectionHashByAuthId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectConnectionHashByAuthId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}
