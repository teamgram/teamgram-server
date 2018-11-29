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

type MessageBoxesDAO struct {
	db *sqlx.DB
}

func NewMessageBoxesDAO(db *sqlx.DB) *MessageBoxesDAO {
	return &MessageBoxesDAO{db}
}

// insert ignore into message_boxes(user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2) values (:user_id, :user_message_box_id, :dialog_id, :dialog_message_id, :message_data_id, :message_box_type, :reply_to_msg_id, :mentioned, :media_unread, :date2)
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) Insert(do *dataobject.MessageBoxesDO) int64 {
	var query = "insert ignore into message_boxes(user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2) values (:user_id, :user_message_box_id, :dialog_id, :dialog_message_id, :message_data_id, :message_box_type, :reply_to_msg_id, :mentioned, :media_unread, :date2)"
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

// select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = :user_id and deleted = 0 and user_message_box_id in (:idList) order by user_message_box_id desc
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectByMessageIdList(user_id int32, idList []int32) []dataobject.MessageBoxesDO {
	var q = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = ? and deleted = 0 and user_message_box_id in (?) order by user_message_box_id desc"
	query, a, err := sqlx.In(q, user_id, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByMessageIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = :user_id and user_message_box_id = :user_message_box_id and deleted = 0 limit 1
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectByMessageId(user_id int32, user_message_box_id int32) *dataobject.MessageBoxesDO {
	var query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1"
	rows, err := dao.db.Queryx(query, user_id, user_message_box_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageBoxesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByMessageId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where deleted = 0 and message_data_id in (:idList) order by user_message_box_id desc
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectByMessageDataIdList(idList []int64) []dataobject.MessageBoxesDO {
	var q = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where deleted = 0 and message_data_id in (?) order by user_message_box_id desc"
	query, a, err := sqlx.In(q, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByMessageDataIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByMessageDataIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByMessageDataIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where message_data_id = :message_data_id and deleted = 0 limit 1
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectByMessageDataId(message_data_id int64) *dataobject.MessageBoxesDO {
	var query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where message_data_id = ? and deleted = 0 limit 1"
	rows, err := dao.db.Queryx(query, message_data_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByMessageDataId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageBoxesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByMessageDataId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByMessageDataId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = :user_id and dialog_id = :dialog_id and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectBackwardByOffsetLimit(user_id int32, dialog_id int64, user_message_box_id int32, limit int32) []dataobject.MessageBoxesDO {
	var query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = ? and dialog_id = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
	rows, err := dao.db.Queryx(query, user_id, dialog_id, user_message_box_id, limit)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectBackwardByOffsetLimit(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectBackwardByOffsetLimit(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectBackwardByOffsetLimit(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = :user_id and dialog_id = :dialog_id and user_message_box_id >= :user_message_box_id and deleted = 0 order by user_message_box_id asc limit :limit
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectForwardByPeerOffsetLimit(user_id int32, dialog_id int64, user_message_box_id int32, limit int32) []dataobject.MessageBoxesDO {
	var query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = ? and dialog_id = ? and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
	rows, err := dao.db.Queryx(query, user_id, dialog_id, user_message_box_id, limit)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectForwardByPeerOffsetLimit(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectForwardByPeerOffsetLimit(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectForwardByPeerOffsetLimit(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select user_message_box_id, message_box_type from message_boxes where user_id = :peerId and deleted = 0 and message_data_id = (select message_data_id from message_boxes where user_id = :user_id and user_message_box_id = :user_message_box_id and deleted = 0 limit 1)
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectPeerMessageId(peerId int32, user_id int32, user_message_box_id int32) *dataobject.MessageBoxesDO {
	var query = "select user_message_box_id, message_box_type from message_boxes where user_id = ? and deleted = 0 and message_data_id = (select message_data_id from message_boxes where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1)"
	rows, err := dao.db.Queryx(query, peerId, user_id, user_message_box_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectPeerMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageBoxesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectPeerMessageId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectPeerMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id != :user_id and dialog_message_id in (select dialog_message_id from message_boxes where user_id = :user_id and user_message_box_id in (:idList)) and deleted = 0
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectPeerDialogMessageIdList(user_id int32, idList []int32) []dataobject.MessageBoxesDO {
	var q = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id != ? and dialog_message_id in (select dialog_message_id from message_boxes where user_id = ? and user_message_box_id in (?)) and deleted = 0"
	query, a, err := sqlx.In(q, user_id, user_id, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectPeerDialogMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectPeerDialogMessageIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectPeerDialogMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where dialog_message_id = (select dialog_message_id from message_boxes where user_id = :user_id and user_message_box_id = :user_message_box_id) and deleted = 0
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectDialogMessageListByMessageId(user_id int32, user_message_box_id int32) []dataobject.MessageBoxesDO {
	var query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where dialog_message_id = (select dialog_message_id from message_boxes where user_id = ? and user_message_box_id = ?) and deleted = 0"
	rows, err := dao.db.Queryx(query, user_id, user_message_box_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectDialogMessageListByMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectDialogMessageListByMessageId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectDialogMessageListByMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id != :user_id and dialog_message_id = (select dialog_message_id from messages where user_id = :user_id and user_message_box_id = :user_message_box_id) and deleted = 0
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectPeerDialogMessageListByMessageId(user_id int32, user_message_box_id int32) []dataobject.MessageBoxesDO {
	var query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id != ? and dialog_message_id = (select dialog_message_id from messages where user_id = ? and user_message_box_id = ?) and deleted = 0"
	rows, err := dao.db.Queryx(query, user_id, user_id, user_message_box_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectPeerDialogMessageListByMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectPeerDialogMessageListByMessageId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectPeerDialogMessageListByMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select user_message_box_id from message_boxes where user_id = :user_id and deleted = 0 order by user_message_box_id desc limit 2
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectLastTwoMessageId(user_id int32) *dataobject.MessageBoxesDO {
	var query = "select user_message_box_id from message_boxes where user_id = ? and deleted = 0 order by user_message_box_id desc limit 2"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectLastTwoMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageBoxesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectLastTwoMessageId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectLastTwoMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = :user_id and user_message_box_id in (:idList) and deleted = 0
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectDialogsByMessageIdList(user_id int32, idList []int32) []dataobject.MessageBoxesDO {
	var q = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = ? and user_message_box_id in (?) and deleted = 0"
	query, a, err := sqlx.In(q, user_id, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectDialogsByMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectDialogsByMessageIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectDialogsByMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// update message_boxes set deleted = 1 where user_id = :user_id and user_message_box_id in (:idList) and deleted = 0
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) DeleteMessagesByMessageIdList(user_id int32, idList []int32) int64 {
	var q = "update message_boxes set deleted = 1 where user_id = ? and user_message_box_id in (?) and deleted = 0"
	query, a, err := sqlx.In(q, user_id, idList)
	r, err := dao.db.Exec(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in DeleteMessagesByMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in DeleteMessagesByMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// select user_message_box_id, date2 from message_boxes where user_id = :user_id and dialog_id = :dialog_id and deleted = 0 order by user_message_box_id desc
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectDialogMessageIdList(user_id int32, dialog_id int64) []dataobject.MessageBoxesDO {
	var query = "select user_message_box_id, date2 from message_boxes where user_id = ? and dialog_id = ? and deleted = 0 order by user_message_box_id desc"
	rows, err := dao.db.Queryx(query, user_id, dialog_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectDialogMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectDialogMessageIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectDialogMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id != :user_id and message_data_id = :message_data_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectPeerMessageList(user_id int32, message_data_id int64) []dataobject.MessageBoxesDO {
	var query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id != ? and message_data_id = ? and deleted = 0"
	rows, err := dao.db.Queryx(query, user_id, message_data_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectPeerMessageList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectPeerMessageList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectPeerMessageList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// update message_boxes set media_unread = 0 where user_id = :user_id and user_message_box_id = :user_message_box_id
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) UpdateMediaUnread(user_id int32, user_message_box_id int32) int64 {
	var query = "update message_boxes set media_unread = 0 where user_id = ? and user_message_box_id = ?"
	r, err := dao.db.Exec(query, user_id, user_message_box_id)

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
