/*
 *  Copyright (c) 2018, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mysql_dao

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
)

type ChannelMediaUnreadDAO struct {
	db *sqlx.DB
}

func NewChannelMediaUnreadDAO(db *sqlx.DB) *ChannelMediaUnreadDAO {
	return &ChannelMediaUnreadDAO{db}
}

// insert into channel_media_unread(user_id, channel_id, channel_message_id, media_unread) values (:user_id, :channel_id, :channel_message_id, :media_unread)
// TODO(@benqi): sqlmap
func (dao *ChannelMediaUnreadDAO) Insert(do *dataobject.ChannelMediaUnreadDO) int64 {
	var query = "insert into channel_media_unread(user_id, channel_id, channel_message_id, media_unread) values (:user_id, :channel_id, :channel_message_id, :media_unread)"
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

// select media_unread from channel_media_unread where user_id = :user_id and channel_id = :channel_id and channel_message_id = :channel_message_id
// TODO(@benqi): sqlmap
func (dao *ChannelMediaUnreadDAO) SelectMediaUnread(user_id int32, channel_id int32, channel_message_id int32) *dataobject.ChannelMediaUnreadDO {
	var query = "select media_unread from channel_media_unread where user_id = ? and channel_id = ? and channel_message_id = ?"
	rows, err := dao.db.Queryx(query, user_id, channel_id, channel_message_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectMediaUnread(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.ChannelMediaUnreadDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectMediaUnread(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectMediaUnread(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// update channel_media_unread set media_unread = 0 where user_id = :user_id and channel_id = :channel_id and channel_message_id = :channel_message_id
// TODO(@benqi): sqlmap
func (dao *ChannelMediaUnreadDAO) UpdateMediaUnread(user_id int32, channel_id int32, channel_message_id int32) int64 {
	var query = "update channel_media_unread set media_unread = 0 where user_id = ? and channel_id = ? and channel_message_id = ?"
	r, err := dao.db.Exec(query, user_id, channel_id, channel_message_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateMediaUnread(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateMediaUnread(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}
