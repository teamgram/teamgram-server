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

type ChannelMessageBoxesDAO struct {
	db *sqlx.DB
}

func NewChannelMessageBoxesDAO(db *sqlx.DB) *ChannelMessageBoxesDAO {
	return &ChannelMessageBoxesDAO{db}
}

// insert into channel_message_boxes(sender_user_id, channel_id, channel_message_box_id, message_id, `date`) values (:sender_user_id, :channel_id, :channel_message_box_id, :message_id, :date)
// TODO(@benqi): sqlmap
func (dao *ChannelMessageBoxesDAO) Insert(do *dataobject.ChannelMessageBoxesDO) int64 {
	var query = "insert into channel_message_boxes(sender_user_id, channel_id, channel_message_box_id, message_id, `date`) values (:sender_user_id, :channel_id, :channel_message_box_id, :message_id, :date)"
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

// select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = :channel_id and deleted = 0 and channel_message_box_id in (:idList) order by channel_message_box_id desc
// TODO(@benqi): sqlmap
func (dao *ChannelMessageBoxesDAO) SelectByMessageIdList(channel_id int32, idList []int32) []dataobject.ChannelMessageBoxesDO {
	var q = "select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = ? and deleted = 0 and channel_message_box_id in (?) order by channel_message_box_id desc"
	query, a, err := sqlx.In(q, channel_id, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChannelMessageBoxesDO
	for rows.Next() {
		v := dataobject.ChannelMessageBoxesDO{}

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

// select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = :channel_id and channel_message_box_id = :channel_message_box_id and deleted = 0 limit 1
// TODO(@benqi): sqlmap
func (dao *ChannelMessageBoxesDAO) SelectByMessageId(channel_id int32, channel_message_box_id int32) *dataobject.ChannelMessageBoxesDO {
	var query = "select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = ? and channel_message_box_id = ? and deleted = 0 limit 1"
	rows, err := dao.db.Queryx(query, channel_id, channel_message_box_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.ChannelMessageBoxesDO{}
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

// select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = :channel_id and channel_message_box_id < :channel_message_box_id and deleted = 0 order by channel_message_box_id desc limit :limit
// TODO(@benqi): sqlmap
func (dao *ChannelMessageBoxesDAO) SelectBackwardByOffsetLimit(channel_id int32, channel_message_box_id int32, limit int32) []dataobject.ChannelMessageBoxesDO {
	var query = "select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = ? and channel_message_box_id < ? and deleted = 0 order by channel_message_box_id desc limit ?"
	rows, err := dao.db.Queryx(query, channel_id, channel_message_box_id, limit)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectBackwardByOffsetLimit(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChannelMessageBoxesDO
	for rows.Next() {
		v := dataobject.ChannelMessageBoxesDO{}

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

// select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = :channel_id and channel_message_box_id >= :channel_message_box_id and deleted = 0 order by channel_message_box_id asc limit :limit
// TODO(@benqi): sqlmap
func (dao *ChannelMessageBoxesDAO) SelectForwardByOffsetLimit(channel_id int32, channel_message_box_id int32, limit int32) []dataobject.ChannelMessageBoxesDO {
	var query = "select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = ? and channel_message_box_id >= ? and deleted = 0 order by channel_message_box_id asc limit ?"
	rows, err := dao.db.Queryx(query, channel_id, channel_message_box_id, limit)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectForwardByOffsetLimit(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChannelMessageBoxesDO
	for rows.Next() {
		v := dataobject.ChannelMessageBoxesDO{}

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
