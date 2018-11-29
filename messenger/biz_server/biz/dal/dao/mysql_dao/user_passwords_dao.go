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

type UserPasswordsDAO struct {
	db *sqlx.DB
}

func NewUserPasswordsDAO(db *sqlx.DB) *UserPasswordsDAO {
	return &UserPasswordsDAO{db}
}

// insert into user_passwords(user_id, server_salt, hash, salt, hint, email, state) values (:user_id, :server_salt, '', '', '', '', 0)
// TODO(@benqi): sqlmap
func (dao *UserPasswordsDAO) Insert(do *dataobject.UserPasswordsDO) int64 {
	var query = "insert into user_passwords(user_id, server_salt, hash, salt, hint, email, state) values (:user_id, :server_salt, '', '', '', '', 0)"
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

// select user_id, server_salt, hash, salt, hint, email, state from user_passwords where user_id = :user_id limit 1
// TODO(@benqi): sqlmap
func (dao *UserPasswordsDAO) SelectByUserId(user_id int32) *dataobject.UserPasswordsDO {
	var query = "select user_id, server_salt, hash, salt, hint, email, state from user_passwords where user_id = ? limit 1"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByUserId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UserPasswordsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByUserId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByUserId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select code, code_expired, attempts from user_passwords where user_id = :user_id limit 1
// TODO(@benqi): sqlmap
func (dao *UserPasswordsDAO) SelectCode(user_id int32) *dataobject.UserPasswordsDO {
	var query = "select code, code_expired, attempts from user_passwords where user_id = ? limit 1"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectCode(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UserPasswordsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectCode(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectCode(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// update user_passwords set salt = :salt, hash = :hash, hint = :hint, email = :email, state = :state where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserPasswordsDAO) Update(salt string, hash string, hint string, email string, state int8, user_id int32) int64 {
	var query = "update user_passwords set salt = ?, hash = ?, hint = ?, email = ?, state = ? where user_id = ?"
	r, err := dao.db.Exec(query, salt, hash, hint, email, state, user_id)

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
