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

type AuthChannelUpdatesStateDAO struct {
	db *sqlx.DB
}

func NewAuthChannelUpdatesStateDAO(db *sqlx.DB) *AuthChannelUpdatesStateDAO {
	return &AuthChannelUpdatesStateDAO{db}
}

// insert into auth_channel_updates_state(auth_key_id, user_id, channel_id, pts, pts2, `date`) values (:auth_key_id, :user_id, :channel_id, :pts, :pts2, :date)
// TODO(@benqi): sqlmap
func (dao *AuthChannelUpdatesStateDAO) Insert(do *dataobject.AuthChannelUpdatesStateDO) int64 {
	var query = "insert into auth_channel_updates_state(auth_key_id, user_id, channel_id, pts, pts2, `date`) values (:auth_key_id, :user_id, :channel_id, :pts, :pts2, :date)"
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

// update auth_channel_updates_state set pts = :pts, pts2 = :pts where auth_key_id = :auth_key_id and channel_id = :channel_id
// TODO(@benqi): sqlmap
func (dao *AuthChannelUpdatesStateDAO) UpdateChannelPts(pts int32, auth_key_id int64, channel_id int32) int64 {
	var query = "update auth_channel_updates_state set pts = ?, pts2 = ? where auth_key_id = ? and channel_id = ?"
	r, err := dao.db.Exec(query, pts, pts, auth_key_id, channel_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateChannelPts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateChannelPts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update auth_channel_updates_state set pts2 = :pts2 where auth_key_id = :auth_key_id and channel_id = :channel_id
// TODO(@benqi): sqlmap
func (dao *AuthChannelUpdatesStateDAO) UpdateChannelPts2(pts2 int32, auth_key_id int64, channel_id int32) int64 {
	var query = "update auth_channel_updates_state set pts2 = ? where auth_key_id = ? and channel_id = ?"
	r, err := dao.db.Exec(query, pts2, auth_key_id, channel_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateChannelPts2(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateChannelPts2(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// select auth_key_id, user_id, channel_id, pts, pts2, `date` from auth_channel_updates_state where auth_key_id = :auth_key_id and channel_id = :channel_id
// TODO(@benqi): sqlmap
func (dao *AuthChannelUpdatesStateDAO) SelectByAuthId(auth_key_id int64, channel_id int32) *dataobject.AuthChannelUpdatesStateDO {
	var query = "select auth_key_id, user_id, channel_id, pts, pts2, `date` from auth_channel_updates_state where auth_key_id = ? and channel_id = ?"
	rows, err := dao.db.Queryx(query, auth_key_id, channel_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByAuthId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.AuthChannelUpdatesStateDO{}
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
