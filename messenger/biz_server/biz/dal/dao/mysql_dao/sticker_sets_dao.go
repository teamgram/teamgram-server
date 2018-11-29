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

type StickerSetsDAO struct {
	db *sqlx.DB
}

func NewStickerSetsDAO(db *sqlx.DB) *StickerSetsDAO {
	return &StickerSetsDAO{db}
}

// insert into sticker_sets(sticker_set_id, access_hash) values (:sticker_set_id, :access_hash)
// TODO(@benqi): sqlmap
func (dao *StickerSetsDAO) Insert(do *dataobject.StickerSetsDO) int64 {
	var query = "insert into sticker_sets(sticker_set_id, access_hash) values (:sticker_set_id, :access_hash)"
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

// select sticker_set_id, access_hash, title, short_name, count, hash from sticker_sets where hash > 0
// TODO(@benqi): sqlmap
func (dao *StickerSetsDAO) SelectAll() []dataobject.StickerSetsDO {
	var query = "select sticker_set_id, access_hash, title, short_name, count, hash from sticker_sets where hash > 0"
	rows, err := dao.db.Queryx(query)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectAll(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.StickerSetsDO
	for rows.Next() {
		v := dataobject.StickerSetsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectAll(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectAll(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select sticker_set_id, access_hash, title, short_name, count, hash from sticker_sets where sticker_set_id = :sticker_set_id and access_hash = :access_hash
// TODO(@benqi): sqlmap
func (dao *StickerSetsDAO) SelectByID(sticker_set_id int64, access_hash int64) *dataobject.StickerSetsDO {
	var query = "select sticker_set_id, access_hash, title, short_name, count, hash from sticker_sets where sticker_set_id = ? and access_hash = ?"
	rows, err := dao.db.Queryx(query, sticker_set_id, access_hash)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByID(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.StickerSetsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByID(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByID(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select sticker_set_id, access_hash, title, short_name, count, hash from sticker_sets where short_name = :short_name
// TODO(@benqi): sqlmap
func (dao *StickerSetsDAO) SelectByShortName(short_name string) *dataobject.StickerSetsDO {
	var query = "select sticker_set_id, access_hash, title, short_name, count, hash from sticker_sets where short_name = ?"
	rows, err := dao.db.Queryx(query, short_name)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByShortName(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.StickerSetsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByShortName(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByShortName(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}
