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

type StickerPacksDAO struct {
	db *sqlx.DB
}

func NewStickerPacksDAO(db *sqlx.DB) *StickerPacksDAO {
	return &StickerPacksDAO{db}
}

// insert into sticker_packs(sticker_set_id, emoticon, document_id) values (:sticker_set_id, :emoticon, :document_id)
// TODO(@benqi): sqlmap
func (dao *StickerPacksDAO) Insert(do *dataobject.StickerPacksDO) int64 {
	var query = "insert into sticker_packs(sticker_set_id, emoticon, document_id) values (:sticker_set_id, :emoticon, :document_id)"
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

// select sticker_set_id, emoticon, document_id from sticker_packs where sticker_set_id = :sticker_set_id
// TODO(@benqi): sqlmap
func (dao *StickerPacksDAO) SelectBySetID(sticker_set_id int64) []dataobject.StickerPacksDO {
	var query = "select sticker_set_id, emoticon, document_id from sticker_packs where sticker_set_id = ?"
	rows, err := dao.db.Queryx(query, sticker_set_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectBySetID(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.StickerPacksDO
	for rows.Next() {
		v := dataobject.StickerPacksDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectBySetID(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectBySetID(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}
