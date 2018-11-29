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

type AuthsDAO struct {
	db *sqlx.DB
}

func NewAuthsDAO(db *sqlx.DB) *AuthsDAO {
	return &AuthsDAO{db}
}

// insert into auths(auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip) values (:auth_key_id, :layer, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :client_ip) on duplicate key update layer = values(layer), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), client_ip = values(client_ip)
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) InsertOrUpdate(do *dataobject.AuthsDO) int64 {
	var query = "insert into auths(auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip) values (:auth_key_id, :layer, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :client_ip) on duplicate key update layer = values(layer), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), client_ip = values(client_ip)"
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

// select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip from auths where auth_key_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) SelectSessions(idList []int64) []dataobject.AuthsDO {
	var q = "select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip from auths where auth_key_id in (?)"
	query, a, err := sqlx.In(q, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectSessions(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.AuthsDO
	for rows.Next() {
		v := dataobject.AuthsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectSessions(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectSessions(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select layer from auths where auth_key_id = :auth_key_id
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) SelectLayer(auth_key_id int64) *dataobject.AuthsDO {
	var query = "select layer from auths where auth_key_id = ?"
	rows, err := dao.db.Queryx(query, auth_key_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectLayer(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.AuthsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectLayer(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectLayer(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}
