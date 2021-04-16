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

type ChannelMessagesDAO struct {
	db *sqlx.DB
}

func NewChannelMessagesDAO(db *sqlx.DB) *ChannelMessagesDAO {
	return &ChannelMessagesDAO{db}
}

// insert ignore into channel_messages(channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, has_media_unread, edit_message, edit_date, views, `date`) values (:channel_id, :channel_message_id, :sender_user_id, :random_id, :message_data_id, :message_type, :message_data, :has_media_unread, :edit_message, :edit_date, :views, :date)
// TODO(@benqi): sqlmap
func (dao *ChannelMessagesDAO) Insert(do *dataobject.ChannelMessagesDO) int64 {
	var query = "insert ignore into channel_messages(channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, has_media_unread, edit_message, edit_date, views, `date`) values (:channel_id, :channel_message_id, :sender_user_id, :random_id, :message_data_id, :message_type, :message_data, :has_media_unread, :edit_message, :edit_date, :views, :date)"
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

// select channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, has_media_unread, edit_message, edit_date, views, `date` from channel_messages where channel_id = :channel_id and deleted = 0 and channel_message_id in (:idList) order by channel_message_id desc
// TODO(@benqi): sqlmap
func (dao *ChannelMessagesDAO) SelectByMessageIdList(channel_id int32, idList []int32) []dataobject.ChannelMessagesDO {
	var q = "select channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, has_media_unread, edit_message, edit_date, views, `date` from channel_messages where channel_id = ? and deleted = 0 and channel_message_id in (?) order by channel_message_id desc"
	query, a, err := sqlx.In(q, channel_id, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChannelMessagesDO
	for rows.Next() {
		v := dataobject.ChannelMessagesDO{}

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

// select channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, has_media_unread, edit_message, edit_date, views, `date` from channel_messages where channel_id = :channel_id and channel_message_id = :channel_message_id and deleted = 0 limit 1
// TODO(@benqi): sqlmap
func (dao *ChannelMessagesDAO) SelectByMessageId(channel_id int32, channel_message_id int32) *dataobject.ChannelMessagesDO {
	var query = "select channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, has_media_unread, edit_message, edit_date, views, `date` from channel_messages where channel_id = ? and channel_message_id = ? and deleted = 0 limit 1"
	rows, err := dao.db.Queryx(query, channel_id, channel_message_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.ChannelMessagesDO{}
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

// select channel_message_id, views from channel_messages where channel_id = :channel_id and channel_message_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *ChannelMessagesDAO) SelectMessagesViews(channel_id int32, idList []int32) []dataobject.ChannelMessagesDO {
	var q = "select channel_message_id, views from channel_messages where channel_id = ? and channel_message_id in (?)"
	query, a, err := sqlx.In(q, channel_id, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectMessagesViews(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChannelMessagesDO
	for rows.Next() {
		v := dataobject.ChannelMessagesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectMessagesViews(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectMessagesViews(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// update channel_messages set views = views + 1 where channel_id = :channel_id and channel_message_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *ChannelMessagesDAO) UpdateMessagesViews(channel_id int32, idList []int32) int64 {
	var q = "update channel_messages set views = views + 1 where channel_id = ? and channel_message_id in (?)"
	query, a, err := sqlx.In(q, channel_id, idList)
	r, err := dao.db.Exec(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateMessagesViews(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateMessagesViews(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// select channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, has_media_unread, edit_message, edit_date, views, `date` from channel_messages where sender_user_id = :sender_user_id and random_id = :random_id
// TODO(@benqi): sqlmap
func (dao *ChannelMessagesDAO) SelectByRandomId(sender_user_id int32, random_id int64) *dataobject.ChannelMessagesDO {
	var query = "select channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, has_media_unread, edit_message, edit_date, views, `date` from channel_messages where sender_user_id = ? and random_id = ?"
	rows, err := dao.db.Queryx(query, sender_user_id, random_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByRandomId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.ChannelMessagesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByRandomId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByRandomId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, has_media_unread, edit_message, edit_date, views, `date` from channel_messages where channel_id = :channel_id and channel_message_id < :channel_message_id and deleted = 0 order by channel_message_id desc limit :limit
// TODO(@benqi): sqlmap
func (dao *ChannelMessagesDAO) SelectBackwardByOffsetLimit(channel_id int32, channel_message_id int32, limit int32) []dataobject.ChannelMessagesDO {
	var query = "select channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, has_media_unread, edit_message, edit_date, views, `date` from channel_messages where channel_id = ? and channel_message_id < ? and deleted = 0 order by channel_message_id desc limit ?"
	rows, err := dao.db.Queryx(query, channel_id, channel_message_id, limit)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectBackwardByOffsetLimit(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChannelMessagesDO
	for rows.Next() {
		v := dataobject.ChannelMessagesDO{}

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

// select channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, has_media_unread, edit_message, edit_date, views, `date` from channel_messages where channel_id = :channel_id and channel_message_id >= :channel_message_id and deleted = 0 order by channel_message_id asc limit :limit
// TODO(@benqi): sqlmap
func (dao *ChannelMessagesDAO) SelectForwardByOffsetLimit(channel_id int32, channel_message_id int32, limit int32) []dataobject.ChannelMessagesDO {
	var query = "select channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, has_media_unread, edit_message, edit_date, views, `date` from channel_messages where channel_id = ? and channel_message_id >= ? and deleted = 0 order by channel_message_id asc limit ?"
	rows, err := dao.db.Queryx(query, channel_id, channel_message_id, limit)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectForwardByOffsetLimit(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChannelMessagesDO
	for rows.Next() {
		v := dataobject.ChannelMessagesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectForwardByOffsetLimit(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectForwardByOffsetLimit(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}
