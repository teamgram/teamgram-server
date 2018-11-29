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

type PopularContactsDAO struct {
	db *sqlx.DB
}

func NewPopularContactsDAO(db *sqlx.DB) *PopularContactsDAO {
	return &PopularContactsDAO{db}
}

// insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) InsertOrUpdate(do *dataobject.PopularContactsDO) int64 {
	var query = "insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1"
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

// update popular_contacts set importers = importers + 1 where phone = :phone
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) IncreaseImporters(phone string) int64 {
	var query = "update popular_contacts set importers = importers + 1 where phone = ?"
	r, err := dao.db.Exec(query, phone)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in IncreaseImporters(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in IncreaseImporters(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update popular_contacts set importers = importers + 1 where phone in (:phoneList)
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) IncreaseImportersList(phoneList []string) int64 {
	var q = "update popular_contacts set importers = importers + 1 where phone in (?)"
	query, a, err := sqlx.In(q, phoneList)
	r, err := dao.db.Exec(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in IncreaseImportersList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in IncreaseImportersList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// select phone, importers from popular_contacts where phone = :phone
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) SelectImporters(phone string) *dataobject.PopularContactsDO {
	var query = "select phone, importers from popular_contacts where phone = ?"
	rows, err := dao.db.Queryx(query, phone)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectImporters(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.PopularContactsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectImporters(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectImporters(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select phone, importers from popular_contacts where phone in (:phoneList)
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) SelectImportersList(phoneList []string) []dataobject.PopularContactsDO {
	var q = "select phone, importers from popular_contacts where phone in (?)"
	query, a, err := sqlx.In(q, phoneList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectImportersList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.PopularContactsDO
	for rows.Next() {
		v := dataobject.PopularContactsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectImportersList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectImportersList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}
