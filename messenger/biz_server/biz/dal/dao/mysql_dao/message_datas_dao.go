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

type MessageDatasDAO struct {
	db *sqlx.DB
}

func NewMessageDatasDAO(db *sqlx.DB) *MessageDatasDAO {
	return &MessageDatasDAO{db}
}

// insert ignore into message_datas(message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date) values (:message_data_id, :dialog_id, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_type, :message_data, :media_unread, :has_media_unread, :date, :edit_message, :edit_date)
// TODO(@benqi): sqlmap
func (dao *MessageDatasDAO) Insert(do *dataobject.MessageDatasDO) int64 {
	var query = "insert ignore into message_datas(message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date) values (:message_data_id, :dialog_id, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_type, :message_data, :media_unread, :has_media_unread, :date, :edit_message, :edit_date)"
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

// select message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where message_data_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *MessageDatasDAO) SelectMessageListByDataIdList(idList []int64) []dataobject.MessageDatasDO {
	var q = "select message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where message_data_id in (?)"
	query, a, err := sqlx.In(q, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectMessageListByDataIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageDatasDO
	for rows.Next() {
		v := dataobject.MessageDatasDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectMessageListByDataIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectMessageListByDataIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where message_data_id = :message_data_id
// TODO(@benqi): sqlmap
func (dao *MessageDatasDAO) SelectMessageByDataId(message_data_id int64) *dataobject.MessageDatasDO {
	var query = "select message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where message_data_id = ?"
	rows, err := dao.db.Queryx(query, message_data_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectMessageByDataId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageDatasDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectMessageByDataId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectMessageByDataId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where dialog_id = :dialog_id and dialog_message_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *MessageDatasDAO) SelectMessageList(dialog_id int64, idList []int32) []dataobject.MessageDatasDO {
	var q = "select message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where dialog_id = ? and dialog_message_id in (?)"
	query, a, err := sqlx.In(q, dialog_id, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectMessageList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageDatasDO
	for rows.Next() {
		v := dataobject.MessageDatasDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectMessageList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectMessageList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where dialog_id = :dialog_id and dialog_message_id = :dialog_message_id limit 1
// TODO(@benqi): sqlmap
func (dao *MessageDatasDAO) SelectMessage(dialog_id int64, dialog_message_id int32) *dataobject.MessageDatasDO {
	var query = "select message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where dialog_id = ? and dialog_message_id = ? limit 1"
	rows, err := dao.db.Queryx(query, dialog_id, dialog_message_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectMessage(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageDatasDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectMessage(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectMessage(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where sender_user_id = :sender_user_id and random_id = :random_id limit 1
// TODO(@benqi): sqlmap
func (dao *MessageDatasDAO) SelectMessageByRandomId(sender_user_id int32, random_id int64) *dataobject.MessageDatasDO {
	var query = "select message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where sender_user_id = ? and random_id = ? limit 1"
	rows, err := dao.db.Queryx(query, sender_user_id, random_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectMessageByRandomId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageDatasDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectMessageByRandomId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectMessageByRandomId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// update message_datas set edit_message = :edit_message, edit_date = :edit_date where dialog_id = :dialog_id and dialog_message_id = :dialog_message_id
// TODO(@benqi): sqlmap
func (dao *MessageDatasDAO) UpdateEditMessage(edit_message string, edit_date int32, dialog_id int64, dialog_message_id int32) int64 {
	var query = "update message_datas set edit_message = ?, edit_date = ? where dialog_id = ? and dialog_message_id = ?"
	r, err := dao.db.Exec(query, edit_message, edit_date, dialog_id, dialog_message_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateEditMessage(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateEditMessage(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update message_datas set media_unread = :media_unread where message_data_id = :message_data_id
// TODO(@benqi): sqlmap
func (dao *MessageDatasDAO) UpdateMediaUnread(media_unread int8, message_data_id int64) int64 {
	var query = "update message_datas set media_unread = ? where message_data_id = ?"
	r, err := dao.db.Exec(query, media_unread, message_data_id)

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
