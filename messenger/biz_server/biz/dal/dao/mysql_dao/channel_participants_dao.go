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

type ChannelParticipantsDAO struct {
	db *sqlx.DB
}

func NewChannelParticipantsDAO(db *sqlx.DB) *ChannelParticipantsDAO {
	return &ChannelParticipantsDAO{db}
}

// insert into channel_participants(channel_id, user_id, is_creator, inviter_user_id, invited_at, joined_at, promoted_by, admin_rights, promoted_at, kicked_by, banned_rights, banned_until_date, banned_at) values (:channel_id, :user_id, :is_creator, :inviter_user_id, :invited_at, :joined_at, :promoted_by, :admin_rights, :promoted_at, :kicked_by, :banned_rights, :banned_until_date, :banned_at)
// TODO(@benqi): sqlmap
func (dao *ChannelParticipantsDAO) Insert(do *dataobject.ChannelParticipantsDO) int64 {
	var query = "insert into channel_participants(channel_id, user_id, is_creator, inviter_user_id, invited_at, joined_at, promoted_by, admin_rights, promoted_at, kicked_by, banned_rights, banned_until_date, banned_at) values (:channel_id, :user_id, :is_creator, :inviter_user_id, :invited_at, :joined_at, :promoted_by, :admin_rights, :promoted_at, :kicked_by, :banned_rights, :banned_until_date, :banned_at)"
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

// insert into channel_participants(channel_id, user_id, inviter_user_id, invited_at, joined_at) values (:channel_id, :user_id, :inviter_user_id, :invited_at, :joined_at) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(inviter_user_id), joined_at = values(joined_at), promoted_by = 0, admin_rights = 0, promoted_at = 0, is_left = 0, left_at = 0, kicked_by = 0, banned_rights = 0, banned_until_date = 0, banned_at = 0
// TODO(@benqi): sqlmap
func (dao *ChannelParticipantsDAO) InsertOrUpdate(do *dataobject.ChannelParticipantsDO) int64 {
	var query = "insert into channel_participants(channel_id, user_id, inviter_user_id, invited_at, joined_at) values (:channel_id, :user_id, :inviter_user_id, :invited_at, :joined_at) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(inviter_user_id), joined_at = values(joined_at), promoted_by = 0, admin_rights = 0, promoted_at = 0, is_left = 0, left_at = 0, kicked_by = 0, banned_rights = 0, banned_until_date = 0, banned_at = 0"
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

// select channel_id, user_id, is_creator, inviter_user_id, invited_at, joined_at, promoted_by, admin_rights, promoted_at, is_left, left_at, kicked_by, banned_rights, banned_until_date, banned_at from channel_participants where channel_id = :channel_id
// TODO(@benqi): sqlmap
func (dao *ChannelParticipantsDAO) SelectByChannelId(channel_id int32) []dataobject.ChannelParticipantsDO {
	var query = "select channel_id, user_id, is_creator, inviter_user_id, invited_at, joined_at, promoted_by, admin_rights, promoted_at, is_left, left_at, kicked_by, banned_rights, banned_until_date, banned_at from channel_participants where channel_id = ?"
	rows, err := dao.db.Queryx(query, channel_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByChannelId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByChannelId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByChannelId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select channel_id, user_id, is_creator, inviter_user_id, invited_at, joined_at, promoted_by, admin_rights, promoted_at, is_left, left_at, kicked_by, banned_rights, banned_until_date, banned_at from channel_participants where channel_id = :channel_id and user_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *ChannelParticipantsDAO) SelectByUserIdList(channel_id int32, idList []int32) []dataobject.ChannelParticipantsDO {
	var q = "select channel_id, user_id, is_creator, inviter_user_id, invited_at, joined_at, promoted_by, admin_rights, promoted_at, is_left, left_at, kicked_by, banned_rights, banned_until_date, banned_at from channel_participants where channel_id = ? and user_id in (?)"
	query, a, err := sqlx.In(q, channel_id, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByUserIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByUserIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByUserIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select channel_id, user_id, is_creator, inviter_user_id, invited_at, joined_at, promoted_by, admin_rights, promoted_at, is_left, left_at, kicked_by, banned_rights, banned_until_date, banned_at from channel_participants where channel_id = :channel_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChannelParticipantsDAO) SelectByUserId(channel_id int32, user_id int32) *dataobject.ChannelParticipantsDO {
	var query = "select channel_id, user_id, is_creator, inviter_user_id, invited_at, joined_at, promoted_by, admin_rights, promoted_at, is_left, left_at, kicked_by, banned_rights, banned_until_date, banned_at from channel_participants where channel_id = ? and user_id = ?"
	rows, err := dao.db.Queryx(query, channel_id, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByUserId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.ChannelParticipantsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByUserId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByUserId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// update channel_participants set banned_rights = :banned_rights, banned_until_date = :banned_until_date where channel_id = :channel_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChannelParticipantsDAO) UpdateBannedRights(banned_rights int32, banned_until_date int32, channel_id int32, user_id int32) int64 {
	var query = "update channel_participants set banned_rights = ?, banned_until_date = ? where channel_id = ? and user_id = ?"
	r, err := dao.db.Exec(query, banned_rights, banned_until_date, channel_id, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateBannedRights(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateBannedRights(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update channel_participants set admin_rights = :admin_rights where channel_id = :channel_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChannelParticipantsDAO) UpdateAdminRights(admin_rights int32, channel_id int32, user_id int32) int64 {
	var query = "update channel_participants set admin_rights = ? where channel_id = ? and user_id = ?"
	r, err := dao.db.Exec(query, admin_rights, channel_id, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateAdminRights(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateAdminRights(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update channel_participants set is_left = 1, left_at = :left_at where channel_id = :channel_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChannelParticipantsDAO) UpdateLeave(left_at int32, channel_id int32, user_id int32) int64 {
	var query = "update channel_participants set is_left = 1, left_at = ? where channel_id = ? and user_id = ?"
	r, err := dao.db.Exec(query, left_at, channel_id, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateLeave(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateLeave(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update channel_participants set read_inbox_max_id = :read_inbox_max_id where channel_id = :channel_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChannelParticipantsDAO) UpdateReadInboxMaxId(read_inbox_max_id int32, channel_id int32, user_id int32) int64 {
	var query = "update channel_participants set read_inbox_max_id = ? where channel_id = ? and user_id = ?"
	r, err := dao.db.Exec(query, read_inbox_max_id, channel_id, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateReadInboxMaxId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateReadInboxMaxId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update channel_participants set read_outbox_max_id = :read_inbox_max_id where channel_id = :channel_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChannelParticipantsDAO) UpdateReadOutboxMaxId(read_inbox_max_id int32, channel_id int32, user_id int32) int64 {
	var query = "update channel_participants set read_outbox_max_id = ? where channel_id = ? and user_id = ?"
	r, err := dao.db.Exec(query, read_inbox_max_id, channel_id, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateReadOutboxMaxId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateReadOutboxMaxId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}
