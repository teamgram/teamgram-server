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
	"github.com/nebula-chat/chatengine/messenger/sync/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
)

type AuthSeqUpdatesDAO struct {
	db *sqlx.DB
}

func NewAuthSeqUpdatesDAO(db *sqlx.DB) *AuthSeqUpdatesDAO {
	return &AuthSeqUpdatesDAO{db}
}

// insert into auth_seq_updates(auth_id, user_id, seq, update_type, update_data, date2) values (:auth_id, :user_id, :seq, :update_type, :update_data, :date2)
// TODO(@benqi): sqlmap
func (dao *AuthSeqUpdatesDAO) Insert(do *dataobject.AuthSeqUpdatesDO) int64 {
	var query = "insert into auth_seq_updates(auth_id, user_id, seq, update_type, update_data, date2) values (:auth_id, :user_id, :seq, :update_type, :update_data, :date2)"
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

// select seq from auth_seq_updates where auth_id = :auth_id and user_id = :user_id order by seq desc limit 1
// TODO(@benqi): sqlmap
func (dao *AuthSeqUpdatesDAO) SelectLastSeq(auth_id int64, user_id int32) *dataobject.AuthSeqUpdatesDO {
	var query = "select seq from auth_seq_updates where auth_id = ? and user_id = ? order by seq desc limit 1"
	rows, err := dao.db.Queryx(query, auth_id, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectLastSeq(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.AuthSeqUpdatesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectLastSeq(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectLastSeq(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select auth_id, user_id, seq, update_type, update_data, date2 from user_seq_updates where auth_id = :auth_id and user_id = :user_id and seq > :seq order by seq asc
// TODO(@benqi): sqlmap
func (dao *AuthSeqUpdatesDAO) SelectByGtSeq(auth_id int64, user_id int32, seq int32) []dataobject.AuthSeqUpdatesDO {
	var query = "select auth_id, user_id, seq, update_type, update_data, date2 from user_seq_updates where auth_id = ? and user_id = ? and seq > ? order by seq asc"
	rows, err := dao.db.Queryx(query, auth_id, user_id, seq)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByGtSeq(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.AuthSeqUpdatesDO
	for rows.Next() {
		v := dataobject.AuthSeqUpdatesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByGtSeq(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByGtSeq(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}
