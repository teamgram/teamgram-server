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

type AuthUsersDAO struct {
	db *sqlx.DB
}

func NewAuthUsersDAO(db *sqlx.DB) *AuthUsersDAO {
	return &AuthUsersDAO{db}
}

// insert into auth_users(auth_id, user_id, hash, device_model, platform, system_version, api_id, app_name, app_version, date_created, date_active, ip, country, region) values (:auth_id, :user_id, :hash, :device_model, :platform, :system_version, :api_id, :app_name, :app_version, :date_created, :date_active, :ip, :country, :region)
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) Insert(do *dataobject.AuthUsersDO) int64 {
	var query = "insert into auth_users(auth_id, user_id, hash, device_model, platform, system_version, api_id, app_name, app_version, date_created, date_active, ip, country, region) values (:auth_id, :user_id, :hash, :device_model, :platform, :system_version, :api_id, :app_name, :app_version, :date_created, :date_active, :ip, :country, :region)"
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

// select auth_id, user_id, hash, device_model, platform, system_version, api_id, app_name, app_version, date_created, date_active, ip, country, region from auth_users where auth_id = :auth_id
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) SelectByAuthId(auth_id int64) *dataobject.AuthUsersDO {
	var query = "select auth_id, user_id, hash, device_model, platform, system_version, api_id, app_name, app_version, date_created, date_active, ip, country, region from auth_users where auth_id = ?"
	rows, err := dao.db.Queryx(query, auth_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByAuthId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.AuthUsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByAuthId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByAuthId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select auth_id, user_id, hash, device_model, platform, system_version, api_id, app_name, app_version, date_created, date_active, ip, country, region from auth_users where user_id = :user_id and hash = :hash
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) SelectByHash(user_id int32, hash int64) *dataobject.AuthUsersDO {
	var query = "select auth_id, user_id, hash, device_model, platform, system_version, api_id, app_name, app_version, date_created, date_active, ip, country, region from auth_users where user_id = ? and hash = ?"
	rows, err := dao.db.Queryx(query, user_id, hash)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByHash(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.AuthUsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByHash(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByHash(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select auth_id, user_id, hash, device_model, platform, system_version, api_id, app_name, app_version, date_created, date_active, ip, country, region from auth_users where user_id = :user_id order by date_active desc limit 6
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) SelectListByUserId(user_id int32) []dataobject.AuthUsersDO {
	var query = "select auth_id, user_id, hash, device_model, platform, system_version, api_id, app_name, app_version, date_created, date_active, ip, country, region from auth_users where user_id = ? order by date_active desc limit 6"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectListByUserId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.AuthUsersDO
	for rows.Next() {
		v := dataobject.AuthUsersDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectListByUserId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectListByUserId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// update auth_users set date_active = :date_active, ip = :ip, country = :country, region = :region where auth_id = :auth_id
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) Update(date_active int32, ip string, country string, region string, auth_id int64) int64 {
	var query = "update auth_users set date_active = ?, ip = ?, country = ?, region = ? where auth_id = ?"
	r, err := dao.db.Exec(query, date_active, ip, country, region, auth_id)

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

// update auth_users set deleted_at = :deleted_at where auth_id = :auth_id
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) Delete(deleted_at int64, auth_id int64) int64 {
	var query = "update auth_users set deleted_at = ? where auth_id = ?"
	r, err := dao.db.Exec(query, deleted_at, auth_id)

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
