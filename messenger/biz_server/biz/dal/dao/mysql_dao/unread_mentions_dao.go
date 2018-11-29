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

type UnreadMentionsDAO struct {
	db *sqlx.DB
}

func NewUnreadMentionsDAO(db *sqlx.DB) *UnreadMentionsDAO {
	return &UnreadMentionsDAO{db}
}

// insert ignore into unread_mentions(user_id, peer_type, peer_id, mentioned_message_id, deleted) values (:user_id, :peer_type, :peer_id, :mentioned_message_id, 0)
// TODO(@benqi): sqlmap
func (dao *UnreadMentionsDAO) InsertIgnore(do *dataobject.UnreadMentionsDO) int64 {
	var query = "insert ignore into unread_mentions(user_id, peer_type, peer_id, mentioned_message_id, deleted) values (:user_id, :peer_type, :peer_id, :mentioned_message_id, 0)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in InsertIgnore(%v), error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in InsertIgnore(%v)_error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return id
}

// update unread_mentions set deleted = 1 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and mentioned_message_id = :mentioned_message_id
// TODO(@benqi): sqlmap
func (dao *UnreadMentionsDAO) Delete(user_id int32, peer_type int8, peer_id int32, mentioned_message_id int32) int64 {
	var query = "update unread_mentions set deleted = 1 where user_id = ? and peer_type = ? and peer_id = ? and mentioned_message_id = ?"
	r, err := dao.db.Exec(query, user_id, peer_type, peer_id, mentioned_message_id)

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
