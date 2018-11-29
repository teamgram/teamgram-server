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

type UserNotifySettingsDAO struct {
	db *sqlx.DB
}

func NewUserNotifySettingsDAO(db *sqlx.DB) *UserNotifySettingsDAO {
	return &UserNotifySettingsDAO{db}
}

// insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) InsertOrUpdate(do *dataobject.UserNotifySettingsDO) int64 {
	var query = "insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0"
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

// select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = :user_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) SelectList(user_id int32) []dataobject.UserNotifySettingsDO {
	var query = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and deleted = 0"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserNotifySettingsDO
	for rows.Next() {
		v := dataobject.UserNotifySettingsDO{}

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

// select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) Select(user_id int32, peer_type int8, peer_id int32) *dataobject.UserNotifySettingsDO {
	var query = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
	rows, err := dao.db.Queryx(query, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in Select(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UserNotifySettingsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in Select(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in Select(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// update user_notify_settings set deleted = 1 where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) DeleteAll(user_id int32) int64 {
	var query = "update user_notify_settings set deleted = 1 where user_id = ?"
	r, err := dao.db.Exec(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in DeleteAll(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in DeleteAll(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}
