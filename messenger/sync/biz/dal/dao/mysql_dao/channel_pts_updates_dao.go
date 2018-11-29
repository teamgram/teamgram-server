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

type ChannelPtsUpdatesDAO struct {
	db *sqlx.DB
}

func NewChannelPtsUpdatesDAO(db *sqlx.DB) *ChannelPtsUpdatesDAO {
	return &ChannelPtsUpdatesDAO{db}
}

// insert into channel_pts_updates(channel_id, pts, pts_count, update_type, update_data, date2) values (:channel_id, :pts, :pts_count, :update_type, :update_data, :date2)
// TODO(@benqi): sqlmap
func (dao *ChannelPtsUpdatesDAO) Insert(do *dataobject.ChannelPtsUpdatesDO) int64 {
	var query = "insert into channel_pts_updates(channel_id, pts, pts_count, update_type, update_data, date2) values (:channel_id, :pts, :pts_count, :update_type, :update_data, :date2)"
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

// select pts from channel_pts_updates where channel_id = :channel_id order by pts desc limit 1
// TODO(@benqi): sqlmap
func (dao *ChannelPtsUpdatesDAO) SelectLastPts(channel_id int32) *dataobject.ChannelPtsUpdatesDO {
	var query = "select pts from channel_pts_updates where channel_id = ? order by pts desc limit 1"
	rows, err := dao.db.Queryx(query, channel_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectLastPts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.ChannelPtsUpdatesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectLastPts(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectLastPts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select channel_id, pts, pts_count, update_type, update_data from channel_pts_updates where channel_id = :channel_id and pts > :pts order by pts asc
// TODO(@benqi): sqlmap
func (dao *ChannelPtsUpdatesDAO) SelectByGtPts(channel_id int32, pts int32) []dataobject.ChannelPtsUpdatesDO {
	var query = "select channel_id, pts, pts_count, update_type, update_data from channel_pts_updates where channel_id = ? and pts > ? order by pts asc"
	rows, err := dao.db.Queryx(query, channel_id, pts)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByGtPts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChannelPtsUpdatesDO
	for rows.Next() {
		v := dataobject.ChannelPtsUpdatesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByGtPts(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByGtPts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}
