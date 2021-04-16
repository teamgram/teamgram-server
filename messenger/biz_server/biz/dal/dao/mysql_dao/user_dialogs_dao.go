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

type UserDialogsDAO struct {
	db *sqlx.DB
}

func NewUserDialogsDAO(db *sqlx.DB) *UserDialogsDAO {
	return &UserDialogsDAO{db}
}

// insert ignore into user_dialogs(user_id, peer_type, peer_id, top_message, unread_count, unread_mentions_count, draft_message_data, date2, created_at) values (:user_id, :peer_type, :peer_id, :top_message, :unread_count, :unread_mentions_count, :draft_message_data, :date2, :created_at)
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) Insert(do *dataobject.UserDialogsDO) int64 {
	var query = "insert ignore into user_dialogs(user_id, peer_type, peer_id, top_message, unread_count, unread_mentions_count, draft_message_data, date2, created_at) values (:user_id, :peer_type, :peer_id, :top_message, :unread_count, :unread_mentions_count, :draft_message_data, :date2, :created_at)"
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

// insert into user_dialogs(user_id, peer_type, peer_id, top_message, unread_count, unread_mentions_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :top_message, :unread_count, :unread_mentions_count, '', :date2) on duplicate key update top_message = values(top_message), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), date2 = values(date2)
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) InsertOrUpdate(do *dataobject.UserDialogsDO) int64 {
	var query = "insert into user_dialogs(user_id, peer_type, peer_id, top_message, unread_count, unread_mentions_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :top_message, :unread_count, :unread_mentions_count, '', :date2) on duplicate key update top_message = values(top_message), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), date2 = values(date2)"
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

// select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = :user_id and is_pinned = 1 order by top_message desc
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectPinnedDialogs(user_id int32) []dataobject.UserDialogsDO {
	var query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = ? and is_pinned = 1 order by top_message desc"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectPinnedDialogs(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserDialogsDO
	for rows.Next() {
		v := dataobject.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectPinnedDialogs(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectPinnedDialogs(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select id from user_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) CheckExists(user_id int32, peer_type int8, peer_id int32) *dataobject.UserDialogsDO {
	var query = "select id from user_dialogs where user_id = ? and peer_type = ? and peer_id = ?"
	rows, err := dao.db.Queryx(query, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in CheckExists(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UserDialogsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in CheckExists(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in CheckExists(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectByPeer(user_id int32, peer_type int8, peer_id int32) *dataobject.UserDialogsDO {
	var query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = ? and peer_type = ? and peer_id = ?"
	rows, err := dao.db.Queryx(query, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByPeer(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UserDialogsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByPeer(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByPeer(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectDialogsByUserID(user_id int32) []dataobject.UserDialogsDO {
	var query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = ?"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectDialogsByUserID(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserDialogsDO
	for rows.Next() {
		v := dataobject.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectDialogsByUserID(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectDialogsByUserID(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = :user_id and is_pinned = :is_pinned and top_message < :top_message order by top_message desc limit :limit
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectByPinnedAndOffset(user_id int32, is_pinned int8, top_message int32, limit int32) []dataobject.UserDialogsDO {
	var query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = ? and is_pinned = ? and top_message < ? and top_message > 0 order by top_message desc limit ?"
	rows, err := dao.db.Queryx(query, user_id, is_pinned, top_message, limit)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByPinnedAndOffset(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserDialogsDO
	for rows.Next() {
		v := dataobject.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByPinnedAndOffset(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByPinnedAndOffset(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = :user_id and is_pinned = :is_pinned and date2 > :date2 order by date2 desc limit :limit
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectDialogsByPinnedAndOffsetDate(user_id int32, is_pinned int8, date2 int32, limit int32) []dataobject.UserDialogsDO {
	var query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = ? and is_pinned = ? and date2 > ? order by date2 desc limit ?"
	rows, err := dao.db.Queryx(query, user_id, is_pinned, date2, limit)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectDialogsByPinnedAndOffsetDate(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserDialogsDO
	for rows.Next() {
		v := dataobject.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectDialogsByPinnedAndOffsetDate(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectDialogsByPinnedAndOffsetDate(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = :user_id and peer_type = :peer_type
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectDialogsByPeerType(user_id int32, peer_type int8) []dataobject.UserDialogsDO {
	var query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = ? and peer_type = ?"
	rows, err := dao.db.Queryx(query, user_id, peer_type)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectDialogsByPeerType(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserDialogsDO
	for rows.Next() {
		v := dataobject.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectDialogsByPeerType(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectDialogsByPeerType(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectListByPeerList(user_id int32, peer_type int8, idList []int32) []dataobject.UserDialogsDO {
	var q = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, show_previews, silent, mute_until, sound, pts, draft_type, draft_message_data, date2 from user_dialogs where user_id = ? and peer_type = ? and peer_id in (?)"
	query, a, err := sqlx.In(q, user_id, peer_type, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectListByPeerList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserDialogsDO
	for rows.Next() {
		v := dataobject.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectListByPeerList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectListByPeerList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// update user_dialogs set top_message = :top_message, date2 = :date2 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) UpdateTopMessage(top_message int32, date2 int32, user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set top_message = ?, date2 = ? where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, top_message, date2, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateTopMessage(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateTopMessage(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_dialogs set top_message = :top_message, draft_type = 0, draft_message_data = '', date2 = :date2 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) UpdateTopMessageAndClearDraft(top_message int32, date2 int32, user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set top_message = ?, draft_type = 0, draft_message_data = '', date2 = ? where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, top_message, date2, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateTopMessageAndClearDraft(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateTopMessageAndClearDraft(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_dialogs set top_message = :top_message, unread_count = unread_count + 1, date2 = :date2 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) UpdateTopMessageAndUnread(top_message int32, date2 int32, user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set top_message = ?, unread_count = unread_count + 1, date2 = ? where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, top_message, date2, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateTopMessageAndUnread(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateTopMessageAndUnread(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_dialogs set top_message = :top_message, unread_mentions_count = unread_mentions_count + 1, date2 = :date2 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) UpdateTopMessageAndMentions(top_message int32, date2 int32, user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set top_message = ?, unread_mentions_count = unread_mentions_count + 1, date2 = ? where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, top_message, date2, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateTopMessageAndMentions(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateTopMessageAndMentions(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_dialogs set top_message = :top_message, unread_mentions_count = unread_mentions_count + 1, draft_type = 0, draft_message_data = '', date2 = :date2 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) UpdateTopMessageAndMentionsAndClearDraft(top_message int32, date2 int32, user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set top_message = ?, unread_mentions_count = unread_mentions_count + 1, draft_type = 0, draft_message_data = '', date2 = ? where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, top_message, date2, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateTopMessageAndMentionsAndClearDraft(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateTopMessageAndMentionsAndClearDraft(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_dialogs set top_message = :top_message, unread_count = unread_count + 1, unread_mentions_count = unread_mentions_count + 1, date2 = :date2 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) UpdateTopMessageAndUnreadAndMentions(top_message int32, date2 int32, user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set top_message = ?, unread_count = unread_count + 1, unread_mentions_count = unread_mentions_count + 1, date2 = ? where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, top_message, date2, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateTopMessageAndUnreadAndMentions(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateTopMessageAndUnreadAndMentions(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_dialogs set unread_count = 0, read_inbox_max_id = :read_inbox_max_id where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) UpdateUnreadByPeer(read_inbox_max_id int32, user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set unread_count = 0, read_inbox_max_id = ? where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, read_inbox_max_id, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateUnreadByPeer(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateUnreadByPeer(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_dialogs set read_outbox_max_id = :read_outbox_max_id where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) UpdateReadOutboxMaxIdByPeer(read_outbox_max_id int32, user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set read_outbox_max_id = ? where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, read_outbox_max_id, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateReadOutboxMaxIdByPeer(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateReadOutboxMaxIdByPeer(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_dialogs set draft_type = 2, draft_message_data = :draft_message_data where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SaveDraft(draft_message_data string, user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set draft_type = 2, draft_message_data = ? where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, draft_message_data, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in SaveDraft(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in SaveDraft(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_dialogs set draft_type = 0, draft_message_data = '' where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) ClearDraft(user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set draft_type = 0, draft_message_data = '' where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in ClearDraft(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in ClearDraft(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_dialogs set is_pinned = :is_pinned where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) UpdatePinned(is_pinned int8, user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set is_pinned = ? where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, is_pinned, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdatePinned(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdatePinned(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_dialogs set top_message = :top_message, unread_count = unread_count + :unreadCount, unread_mentions_count = unread_mentions_count + :unreadMentionCount where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) UpdateDialog(top_message int32, unreadCount int32, unreadMentionCount int32, user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set top_message = ?, unread_count = unread_count + ?, unread_mentions_count = unread_mentions_count + ? where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, top_message, unreadCount, unreadMentionCount, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateDialog(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateDialog(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_dialogs set unread_mentions_count = unread_mentions_count - 1 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) UpdateUnreadMentionCount(user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "update user_dialogs set unread_mentions_count = unread_mentions_count - 1 where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateUnreadMentionCount(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateUnreadMentionCount(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// delete from user_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) Delete(user_id int32, peer_type int8, peer_id int32) int64 {
	var query = "delete from user_dialogs where user_id = ? and peer_type = ? and peer_id = ?"
	r, err := dao.db.Exec(query, user_id, peer_type, peer_id)

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
