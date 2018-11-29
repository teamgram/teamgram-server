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
	"github.com/nebula-chat/chatengine/service/document/biz/dal/dataobject"
)

type PhotoDatasDAO struct {
	db *sqlx.DB
}

func NewPhotoDatasDAO(db *sqlx.DB) *PhotoDatasDAO {
	return &PhotoDatasDAO{db}
}

// insert into photo_datas(photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext) values (:photo_id, :photo_type, :dc_id, :volume_id, :local_id, :access_hash, :width, :height, :file_size, :file_path, :ext)
// TODO(@benqi): sqlmap
func (dao *PhotoDatasDAO) Insert(do *dataobject.PhotoDatasDO) int64 {
	var query = "insert into photo_datas(photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext) values (:photo_id, :photo_type, :dc_id, :volume_id, :local_id, :access_hash, :width, :height, :file_size, :file_path, :ext)"
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

// select id, photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext from photo_datas where dc_id = 2 and volume_id = :volume_id and local_id = :local_id and access_hash = :access_hash
// TODO(@benqi): sqlmap
func (dao *PhotoDatasDAO) SelectByFileLocation(volume_id int64, local_id int32, access_hash int64) *dataobject.PhotoDatasDO {
	var query = "select id, photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext from photo_datas where dc_id = 2 and volume_id = ? and local_id = ? and access_hash = ?"
	rows, err := dao.db.Queryx(query, volume_id, local_id, access_hash)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByFileLocation(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.PhotoDatasDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByFileLocation(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByFileLocation(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select id, photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext from photo_datas where photo_id = :photo_id and local_id != 0
// TODO(@benqi): sqlmap
func (dao *PhotoDatasDAO) SelectListByPhotoId(photo_id int64) []dataobject.PhotoDatasDO {
	var query = "select id, photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext from photo_datas where photo_id = ? and local_id != 0"
	rows, err := dao.db.Queryx(query, photo_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectListByPhotoId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.PhotoDatasDO
	for rows.Next() {
		v := dataobject.PhotoDatasDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectListByPhotoId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectListByPhotoId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}
