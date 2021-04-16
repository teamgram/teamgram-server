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

package dao

import (
	"sync"

	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao/mysql_dao"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao/redis_dao"
	"github.com/nebula-chat/chatengine/pkg/redis_client"
)

const (
	DB_MASTER = "immaster"
	DB_SLAVE  = "imslave"
	CACHE     = "cache"
)

type MysqlDAOList struct {
	// common
	CommonDAO *mysql_dao.CommonDAO

	// auth_key
	AuthKeysDAO  *mysql_dao.AuthKeysDAO
	AuthsDAO     *mysql_dao.AuthsDAO
	AuthUsersDAO *mysql_dao.AuthUsersDAO

	// biz
	UsersDAO                 *mysql_dao.UsersDAO
	DevicesDAO               *mysql_dao.DevicesDAO
	AuthPhoneTransactionsDAO *mysql_dao.AuthPhoneTransactionsDAO
	UserDialogsDAO           *mysql_dao.UserDialogsDAO
	UserContactsDAO          *mysql_dao.UserContactsDAO
	MessageBoxesDAO          *mysql_dao.MessageBoxesDAO
	UserNotifySettingsDAO    *mysql_dao.UserNotifySettingsDAO
	ReportsDAO               *mysql_dao.ReportsDAO
	UserPrivacysDAO          *mysql_dao.UserPrivacysDAO
	TmpPasswordsDAO          *mysql_dao.TmpPasswordsDAO

	ChannelsDAO            *mysql_dao.ChannelsDAO
	ChannelParticipantsDAO *mysql_dao.ChannelParticipantsDAO
	ChannelMessagesDAO     *mysql_dao.ChannelMessagesDAO
	ChannelMessageBoxesDAO *mysql_dao.ChannelMessageBoxesDAO
	ChannelMediaUnreadDAO  *mysql_dao.ChannelMediaUnreadDAO

	ChatsDAO             *mysql_dao.ChatsDAO
	ChatParticipantsDAO  *mysql_dao.ChatParticipantsDAO
	UserPtsUpdatesDAO    *mysql_dao.UserPtsUpdatesDAO
	UserQtsUpdatesDAO    *mysql_dao.UserQtsUpdatesDAO
	AuthSeqUpdatesDAO    *mysql_dao.AuthSeqUpdatesDAO
	AuthUpdatesStateDAO  *mysql_dao.AuthUpdatesStateDAO
	UserPresencesDAO     *mysql_dao.UserPresencesDAO
	UserPasswordsDAO     *mysql_dao.UserPasswordsDAO
	WallPapersDAO        *mysql_dao.WallPapersDAO
	PhoneCallSessionsDAO *mysql_dao.PhoneCallSessionsDAO
	MessageReactDAO      *mysql_dao.MessageReactDAO
	MessageReactDataDAO  *mysql_dao.MessageReactDataDAO
	StickerSetsDAO       *mysql_dao.StickerSetsDAO
	StickerPacksDAO      *mysql_dao.StickerPacksDAO

	MessageDatasDAO *mysql_dao.MessageDatasDAO

	UnregisteredContactsDAO *mysql_dao.UnregisteredContactsDAO
	PopularContactsDAO      *mysql_dao.PopularContactsDAO
	ImportedContactsDAO     *mysql_dao.ImportedContactsDAO
	PhoneBooksDAO           *mysql_dao.PhoneBooksDAO

	UsernameDAO *mysql_dao.UsernameDAO

	BotsDAO        *mysql_dao.BotsDAO
	BotCommandsDAO *mysql_dao.BotCommandsDAO

	UnreadMentionsDAO *mysql_dao.UnreadMentionsDAO
	UserBlocksDAO     *mysql_dao.UserBlocksDAO

	BannedDAO *mysql_dao.BannedDAO
}

// TODO(@benqi): 一主多从
type MysqlDAOManager struct {
	daoListMap map[string]*MysqlDAOList
}

var mysqlDAOManager = &MysqlDAOManager{make(map[string]*MysqlDAOList)}

func InstallMysqlDAOManager(clients sync.Map /*map[string]*sqlx.DB*/) {
	clients.Range(func(key, value interface{}) bool {
		k, _ := key.(string)
		v, _ := value.(*sqlx.DB)

		daoList := &MysqlDAOList{}

		// Common
		daoList.CommonDAO = mysql_dao.NewCommonDAO(v)

		// auth_key
		daoList.AuthKeysDAO = mysql_dao.NewAuthKeysDAO(v)
		daoList.AuthsDAO = mysql_dao.NewAuthsDAO(v)
		daoList.AuthUsersDAO = mysql_dao.NewAuthUsersDAO(v)

		// biz
		daoList.UsersDAO = mysql_dao.NewUsersDAO(v)
		daoList.DevicesDAO = mysql_dao.NewDevicesDAO(v)
		daoList.AuthPhoneTransactionsDAO = mysql_dao.NewAuthPhoneTransactionsDAO(v)
		daoList.UserDialogsDAO = mysql_dao.NewUserDialogsDAO(v)
		daoList.UserContactsDAO = mysql_dao.NewUserContactsDAO(v)
		daoList.MessageBoxesDAO = mysql_dao.NewMessageBoxesDAO(v)
		daoList.AuthUpdatesStateDAO = mysql_dao.NewAuthUpdatesStateDAO(v)
		daoList.UserNotifySettingsDAO = mysql_dao.NewUserNotifySettingsDAO(v)
		daoList.ReportsDAO = mysql_dao.NewReportsDAO(v)
		daoList.UserPrivacysDAO = mysql_dao.NewUserPrivacysDAO(v)
		daoList.TmpPasswordsDAO = mysql_dao.NewTmpPasswordsDAO(v)

		daoList.ChannelsDAO = mysql_dao.NewChannelsDAO(v)
		daoList.ChannelParticipantsDAO = mysql_dao.NewChannelParticipantsDAO(v)
		daoList.ChannelMessagesDAO = mysql_dao.NewChannelMessagesDAO(v)
		daoList.ChannelMessageBoxesDAO = mysql_dao.NewChannelMessageBoxesDAO(v)
		daoList.ChannelMediaUnreadDAO = mysql_dao.NewChannelMediaUnreadDAO(v)

		daoList.ChatsDAO = mysql_dao.NewChatsDAO(v)
		daoList.ChatParticipantsDAO = mysql_dao.NewChatParticipantsDAO(v)
		daoList.UserPtsUpdatesDAO = mysql_dao.NewUserPtsUpdatesDAO(v)
		daoList.UserQtsUpdatesDAO = mysql_dao.NewUserQtsUpdatesDAO(v)
		daoList.AuthSeqUpdatesDAO = mysql_dao.NewAuthSeqUpdatesDAO(v)
		daoList.UserPresencesDAO = mysql_dao.NewUserPresencesDAO(v)
		daoList.UserPasswordsDAO = mysql_dao.NewUserPasswordsDAO(v)
		daoList.WallPapersDAO = mysql_dao.NewWallPapersDAO(v)
		daoList.PhoneCallSessionsDAO = mysql_dao.NewPhoneCallSessionsDAO(v)
		daoList.StickerSetsDAO = mysql_dao.NewStickerSetsDAO(v)
		daoList.StickerPacksDAO = mysql_dao.NewStickerPacksDAO(v)
		daoList.MessageReactDAO = mysql_dao.NewMessageReactDAO(v)
		daoList.MessageReactDataDAO = mysql_dao.NewMessageReactDataDAO(v)
		daoList.MessageDatasDAO = mysql_dao.NewMessageDatasDAO(v)

		daoList.UnregisteredContactsDAO = mysql_dao.NewUnregisteredContactsDAO(v)
		daoList.PopularContactsDAO = mysql_dao.NewPopularContactsDAO(v)
		daoList.ImportedContactsDAO = mysql_dao.NewImportedContactsDAO(v)
		daoList.PhoneBooksDAO = mysql_dao.NewPhoneBooksDAO(v)

		daoList.UsernameDAO = mysql_dao.NewUsernameDAO(v)

		daoList.BotsDAO = mysql_dao.NewBotsDAO(v)
		daoList.BotCommandsDAO = mysql_dao.NewBotCommandsDAO(v)

		daoList.UnreadMentionsDAO = mysql_dao.NewUnreadMentionsDAO(v)

		daoList.UserBlocksDAO = mysql_dao.NewUserBlocksDAO(v)

		daoList.BannedDAO = mysql_dao.NewBannedDAO(v)

		mysqlDAOManager.daoListMap[k] = daoList
		return true
	})
}

func GetMysqlDAOListMap() map[string]*MysqlDAOList {
	return mysqlDAOManager.daoListMap
}

func GetMysqlDAOList(dbName string) (daoList *MysqlDAOList) {
	daoList, ok := mysqlDAOManager.daoListMap[dbName]
	if !ok {
		glog.Errorf("GetMysqlDAOList - Not found daoList: %s", dbName)
	}
	return
}

func GetCommonDAO(dbName string) (dao *mysql_dao.CommonDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.CommonDAO
	}
	return
}
func GetMessageReactDataDAO(dbName string) (dao *mysql_dao.MessageReactDataDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.MessageReactDataDAO
	}
	return
}
func GetMessageReactDAO(dbName string) (dao *mysql_dao.MessageReactDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.MessageReactDAO
	}
	return
}
func GetAuthKeysDAO(dbName string) (dao *mysql_dao.AuthKeysDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.AuthKeysDAO
	}
	return
}

func GetAuthsDAO(dbName string) (dao *mysql_dao.AuthsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.AuthsDAO
	}
	return
}

func GetAuthUsersDAO(dbName string) (dao *mysql_dao.AuthUsersDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.AuthUsersDAO
	}
	return
}

func GetUsersDAO(dbName string) (dao *mysql_dao.UsersDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UsersDAO
	}
	return
}

func GetDevicesDAO(dbName string) (dao *mysql_dao.DevicesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.DevicesDAO
	}
	return
}

func GetAuthPhoneTransactionsDAO(dbName string) (dao *mysql_dao.AuthPhoneTransactionsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.AuthPhoneTransactionsDAO
	}
	return
}

func GetUserDialogsDAO(dbName string) (dao *mysql_dao.UserDialogsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UserDialogsDAO
	}
	return
}

func GetUserContactsDAO(dbName string) (dao *mysql_dao.UserContactsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UserContactsDAO
	}
	return
}

func GetMessageBoxesDAO(dbName string) (dao *mysql_dao.MessageBoxesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.MessageBoxesDAO
	}
	return
}

func GetAuthUpdatesStateDAO(dbName string) (dao *mysql_dao.AuthUpdatesStateDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.AuthUpdatesStateDAO
	}
	return
}

func GetUserNotifySettingsDAO(dbName string) (dao *mysql_dao.UserNotifySettingsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UserNotifySettingsDAO
	}
	return
}

func GetReportsDAO(dbName string) (dao *mysql_dao.ReportsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.ReportsDAO
	}
	return
}

func GetUserPrivacysDAO(dbName string) (dao *mysql_dao.UserPrivacysDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UserPrivacysDAO
	}
	return
}

func GetTmpPasswordsDAO(dbName string) (dao *mysql_dao.TmpPasswordsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.TmpPasswordsDAO
	}
	return
}

func GetChannelsDAO(dbName string) (dao *mysql_dao.ChannelsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.ChannelsDAO
	}
	return
}

func GetChannelParticipantsDAO(dbName string) (dao *mysql_dao.ChannelParticipantsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.ChannelParticipantsDAO
	}
	return
}

func GetChannelMessageBoxesDAO(dbName string) (dao *mysql_dao.ChannelMessageBoxesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.ChannelMessageBoxesDAO
	}
	return
}

func GetChannelMessagesDAO(dbName string) (dao *mysql_dao.ChannelMessagesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.ChannelMessagesDAO
	}
	return
}

func GetChannelMediaUnreadDAO(dbName string) (dao *mysql_dao.ChannelMediaUnreadDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.ChannelMediaUnreadDAO
	}
	return
}

func GetChatsDAO(dbName string) (dao *mysql_dao.ChatsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.ChatsDAO
	}
	return
}

func GetChatParticipantsDAO(dbName string) (dao *mysql_dao.ChatParticipantsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.ChatParticipantsDAO
	}
	return
}

func GetUserPtsUpdatesDAO(dbName string) (dao *mysql_dao.UserPtsUpdatesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UserPtsUpdatesDAO
	}
	return
}

func GetUserQtsUpdatesDAO(dbName string) (dao *mysql_dao.UserQtsUpdatesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UserQtsUpdatesDAO
	}
	return
}

func GetAuthSeqUpdatesDAO(dbName string) (dao *mysql_dao.AuthSeqUpdatesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.AuthSeqUpdatesDAO
	}
	return
}

func GetUserPresencesDAO(dbName string) (dao *mysql_dao.UserPresencesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UserPresencesDAO
	}
	return
}

func GetUserPasswordsDAO(dbName string) (dao *mysql_dao.UserPasswordsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UserPasswordsDAO
	}
	return
}

func GetWallPapersDAO(dbName string) (dao *mysql_dao.WallPapersDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.WallPapersDAO
	}
	return
}

func GetPhoneCallSessionsDAO(dbName string) (dao *mysql_dao.PhoneCallSessionsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.PhoneCallSessionsDAO
	}
	return
}

func GetStickerSetsDAO(dbName string) (dao *mysql_dao.StickerSetsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.StickerSetsDAO
	}
	return
}

func GetStickerPacksDAO(dbName string) (dao *mysql_dao.StickerPacksDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.StickerPacksDAO
	}
	return
}

func GetMessageDatasDAO(dbName string) (dao *mysql_dao.MessageDatasDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.MessageDatasDAO
	}
	return
}

func GetUnregisteredContactsDAO(dbName string) (dao *mysql_dao.UnregisteredContactsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UnregisteredContactsDAO
	}
	return
}

func GetPopularContactsDAO(dbName string) (dao *mysql_dao.PopularContactsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.PopularContactsDAO
	}
	return
}

func GetImportedContactsDAO(dbName string) (dao *mysql_dao.ImportedContactsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.ImportedContactsDAO
	}
	return
}

func GetPhoneBooksDAO(dbName string) (dao *mysql_dao.PhoneBooksDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.PhoneBooksDAO
	}
	return
}

func GetUsernameDAO(dbName string) (dao *mysql_dao.UsernameDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UsernameDAO
	}
	return
}

func GetBotsDAO(dbName string) (dao *mysql_dao.BotsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.BotsDAO
	}
	return
}

func GetBotCommandsDAO(dbName string) (dao *mysql_dao.BotCommandsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.BotCommandsDAO
	}
	return
}

func GetUnreadMentionsDAO(dbName string) (dao *mysql_dao.UnreadMentionsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UnreadMentionsDAO
	}
	return
}

func GetUserBlocksDAO(dbName string) (dao *mysql_dao.UserBlocksDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UserBlocksDAO
	}
	return
}

func GetBannedDAO(dbName string) (dao *mysql_dao.BannedDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.BannedDAO
	}
	return
}

///////////////////////////////////////////////////////////////////////////////////////////
type RedisDAOList struct {
	SequenceDAO *redis_dao.SequenceDAO
}

type RedisDAOManager struct {
	daoListMap map[string]*RedisDAOList
}

var redisDAOManager = &RedisDAOManager{make(map[string]*RedisDAOList)}

func InstallRedisDAOManager(clients map[string]*redis_client.RedisPool) {
	for k, v := range clients {
		daoList := &RedisDAOList{}
		daoList.SequenceDAO = redis_dao.NewSequenceDAO(v)
		redisDAOManager.daoListMap[k] = daoList
	}
}

func GetRedisDAOList(redisName string) (daoList *RedisDAOList) {
	daoList, ok := redisDAOManager.daoListMap[redisName]
	if !ok {
		glog.Errorf("GetRedisDAOList - Not found daoList: %s", redisName)
	}
	return
}

func GetRedisDAOListMap() map[string]*RedisDAOList {
	return redisDAOManager.daoListMap
}

func GetSequenceDAO(redisName string) (dao *redis_dao.SequenceDAO) {
	daoList := GetRedisDAOList(redisName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.SequenceDAO
	}
	return
}
