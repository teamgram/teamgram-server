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

type ImportedContactsDAO struct {
	db *sqlx.DB
}

func NewImportedContactsDAO(db *sqlx.DB) *ImportedContactsDAO {
	return &ImportedContactsDAO{db}
}

// insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) InsertOrUpdate(do *dataobject.ImportedContactsDO) int64 {
	var query = "insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0"
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

// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) SelectList(user_id int32) []dataobject.ImportedContactsDO {
	var query = "select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ImportedContactsDO
	for rows.Next() {
		v := dataobject.ImportedContactsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id and deleted = 0 and imported_user_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) SelectListByImportedList(user_id int32, idList []int32) []dataobject.ImportedContactsDO {
	var q = "select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0 and imported_user_id in (?)"
	query, a, err := sqlx.In(q, user_id, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectListByImportedList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ImportedContactsDO
	for rows.Next() {
		v := dataobject.ImportedContactsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectListByImportedList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectListByImportedList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) SelectAllList(user_id int32) []dataobject.ImportedContactsDO {
	var query = "select id, user_id, imported_user_id from imported_contacts where user_id = ?"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectAllList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ImportedContactsDO
	for rows.Next() {
		v := dataobject.ImportedContactsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectAllList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectAllList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// update imported_contacts set deleted = 1 where user_id = :user_id and imported_user_id = :imported_user_id
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) Delete(user_id int32, imported_user_id int32) int64 {
	var query = "update imported_contacts set deleted = 1 where user_id = ? and imported_user_id = ?"
	r, err := dao.db.Exec(query, user_id, imported_user_id)

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
