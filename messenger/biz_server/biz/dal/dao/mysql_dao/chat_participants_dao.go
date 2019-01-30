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

type ChatParticipantsDAO struct {
	db *sqlx.DB
}

func NewChatParticipantsDAO(db *sqlx.DB) *ChatParticipantsDAO {
	return &ChatParticipantsDAO{db}
}

// insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at)
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) Insert(do *dataobject.ChatParticipantsDO) int64 {
	var query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at)"
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

// insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) InsertOrUpdate(do *dataobject.ChatParticipantsDO) int64 {
	var query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0"
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

// select id, chat_id, user_id, participant_type, inviter_user_id, invited_at, state, kicked_at, left_at from chat_participants where chat_id = :chat_id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) SelectList(chat_id int32) []dataobject.ChatParticipantsDO {
	var query = "select id, chat_id, user_id, participant_type, inviter_user_id, invited_at, state, kicked_at, left_at from chat_participants where chat_id = ?"
	rows, err := dao.db.Queryx(query, chat_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChatParticipantsDO
	for rows.Next() {
		v := dataobject.ChatParticipantsDO{}

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

// update chat_participants set participant_type = :participant_type, inviter_user_id = :inviter_user_id, invited_at = :invited_at, state = 0, kicked_at = 0, left_at = 0 where id = :id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) Update(participant_type int8, inviter_user_id int32, invited_at int32, id int32) int64 {
	var query = "update chat_participants set participant_type = ?, inviter_user_id = ?, invited_at = ?, state = 0, kicked_at = 0, left_at = 0 where id = ?"
	r, err := dao.db.Exec(query, participant_type, inviter_user_id, invited_at, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in Update(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in Update(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update chat_participants set state = 2, kicked_at = :kicked_at where id = :id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateKicked(kicked_at int32, id int32) int64 {
	var query = "update chat_participants set state = 2, kicked_at = ? where id = ?"
	r, err := dao.db.Exec(query, kicked_at, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateKicked(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateKicked(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update chat_participants set state = 1, left_at = :kicked_at where id = :id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateLeft(kicked_at int32, id int32) int64 {
	var query = "update chat_participants set state = 1, left_at = ? where id = ?"
	r, err := dao.db.Exec(query, kicked_at, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateLeft(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateLeft(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update chat_participants set participant_type = :participant_type where id = :id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateParticipantType(participant_type int8, id int32) int64 {
	var query = "update chat_participants set participant_type = ? where id = ?"
	r, err := dao.db.Exec(query, participant_type, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateParticipantType(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateParticipantType(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update chat_participants set participant_type = :participant_type where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateParticipantTypeByUserId(participant_type int8, user_id int32) int64 {
	var query = "update chat_participants set participant_type = ? where user_id = ?"
	r, err := dao.db.Exec(query, participant_type, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateParticipantType(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateParticipantType(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}
