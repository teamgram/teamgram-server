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

type AuthUsersDAO struct {
	db *sqlx.DB
}

func NewAuthUsersDAO(db *sqlx.DB) *AuthUsersDAO {
	return &AuthUsersDAO{db}
}

// insert into auth_users(auth_key_id, user_id, hash, date_created, date_actived) values (:auth_key_id, :user_id, :hash, :date_created, :date_actived) on duplicate key update hash = values(hash), date_actived = values(date_actived), deleted = 0
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) InsertOrUpdates(do *dataobject.AuthUsersDO) int64 {
	var query = "insert into auth_users(auth_key_id, user_id, hash, date_created, date_actived) values (:auth_key_id, :user_id, :hash, :date_created, :date_actived) on duplicate key update hash = values(hash), date_actived = values(date_actived), deleted = 0"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in InsertOrUpdates(%v), error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in InsertOrUpdates(%v)_error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return id
}

// select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where auth_key_id = :auth_key_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) Select(auth_key_id int64) *dataobject.AuthUsersDO {
	var query = "select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where auth_key_id = ? and deleted = 0"
	rows, err := dao.db.Queryx(query, auth_key_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in Select(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.AuthUsersDO{}
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

// select id, auth_key_id, user_id, hash from auth_users where user_id = :user_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) SelectAuthKeyIds(user_id int32) []dataobject.AuthUsersDO {
	var query = "select id, auth_key_id, user_id, hash from auth_users where user_id = ? and deleted = 0"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectAuthKeyIds(_), error: %v", err)
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
			errDesc := fmt.Sprintf("StructScan in SelectAuthKeyIds(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectAuthKeyIds(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// update auth_users set deleted = 1, date_created = 0, date_actived = 0 where id = :id
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) DeleteByHash(id int32) int64 {
	var query = "update auth_users set deleted = 1, date_created = 0, date_actived = 0 where id = ?"
	r, err := dao.db.Exec(query, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in DeleteByHash(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in DeleteByHash(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// select id, auth_key_id, user_id, hash from auth_users where hash = :hash and deleted = 0
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) SelectByHash(hash int64) *dataobject.AuthUsersDO {
	var query = "select id, auth_key_id, user_id, hash from auth_users where hash = ? and deleted = 0"
	rows, err := dao.db.Queryx(query, hash)

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

// update auth_users set deleted = 1, date_actived = 0 where auth_key_id = :auth_key_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) Delete(auth_key_id int64, user_id int32) int64 {
	var query = "update auth_users set deleted = 1, date_actived = 0 where auth_key_id = ? and user_id = ?"
	r, err := dao.db.Exec(query, auth_key_id, user_id)

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
