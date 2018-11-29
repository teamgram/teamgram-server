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

type AuthPhoneTransactionsDAO struct {
	db *sqlx.DB
}

func NewAuthPhoneTransactionsDAO(db *sqlx.DB) *AuthPhoneTransactionsDAO {
	return &AuthPhoneTransactionsDAO{db}
}

// insert into auth_phone_transactions(auth_key_id, phone_number, code, code_expired, transaction_hash, sent_code_type, flash_call_pattern, next_code_type, state, api_id, api_hash, created_time) values (:auth_key_id, :phone_number, :code, :code_expired, :transaction_hash, :sent_code_type, :flash_call_pattern, :next_code_type, :state, :api_id, :api_hash, :created_time)
// TODO(@benqi): sqlmap
func (dao *AuthPhoneTransactionsDAO) Insert(do *dataobject.AuthPhoneTransactionsDO) int64 {
	var query = "insert into auth_phone_transactions(auth_key_id, phone_number, code, code_expired, transaction_hash, sent_code_type, flash_call_pattern, next_code_type, state, api_id, api_hash, created_time) values (:auth_key_id, :phone_number, :code, :code_expired, :transaction_hash, :sent_code_type, :flash_call_pattern, :next_code_type, :state, :api_id, :api_hash, :created_time)"
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

// select id, code, code_expired, sent_code_type, flash_call_pattern, next_code_type, attempts, state from auth_phone_transactions where auth_key_id = :auth_key_id and phone_number = :phone_number and transaction_hash = :transaction_hash
// TODO(@benqi): sqlmap
func (dao *AuthPhoneTransactionsDAO) SelectByPhoneCodeHash(auth_key_id int64, phone_number string, transaction_hash string) *dataobject.AuthPhoneTransactionsDO {
	var query = "select id, code, code_expired, sent_code_type, flash_call_pattern, next_code_type, attempts, state from auth_phone_transactions where auth_key_id = ? and phone_number = ? and transaction_hash = ?"
	rows, err := dao.db.Queryx(query, auth_key_id, phone_number, transaction_hash)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByPhoneCodeHash(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.AuthPhoneTransactionsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByPhoneCodeHash(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByPhoneCodeHash(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// update auth_phone_transactions set state = :state where id = :id
// TODO(@benqi): sqlmap
func (dao *AuthPhoneTransactionsDAO) UpdateState(state int8, id int64) int64 {
	var query = "update auth_phone_transactions set state = ? where id = ?"
	r, err := dao.db.Exec(query, state, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateState(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateState(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update auth_phone_transactions set state = :state where auth_key_id = :auth_key_id and phone_number = :phone_number and transaction_hash = :transaction_hash
// TODO(@benqi): sqlmap
func (dao *AuthPhoneTransactionsDAO) Delete(state int8, auth_key_id int64, phone_number string, transaction_hash string) int64 {
	var query = "update auth_phone_transactions set state = ? where auth_key_id = ? and phone_number = ? and transaction_hash = ?"
	r, err := dao.db.Exec(query, state, auth_key_id, phone_number, transaction_hash)

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

// update auth_phone_transactions set attempts = attempts + 1 where id = :id
// TODO(@benqi): sqlmap
func (dao *AuthPhoneTransactionsDAO) UpdateAttempts(id int64) int64 {
	var query = "update auth_phone_transactions set attempts = attempts + 1 where id = ?"
	r, err := dao.db.Exec(query, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateAttempts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateAttempts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}
