// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package dao

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/dal/dao/mysql_dao"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.BotCommandsDAO
	*mysql_dao.BotsDAO
	*mysql_dao.ImportedContactsDAO
	*mysql_dao.PhoneBooksDAO
	*mysql_dao.PopularContactsDAO
	*mysql_dao.PredefinedUsersDAO
	*mysql_dao.UserContactsDAO
	*mysql_dao.UserGlobalPrivacySettingsDAO
	*mysql_dao.UserNotifySettingsDAO
	*mysql_dao.UserPeerBlocksDAO
	*mysql_dao.UserPeerSettingsDAO
	*mysql_dao.UserPresencesDAO
	*mysql_dao.UserPrivaciesDAO
	*mysql_dao.UserSettingsDAO
	*mysql_dao.UsersDAO
	*mysql_dao.UserProfilePhotosDAO
	*mysql_dao.UnregisteredContactsDAO
	*sqlx.CommonDAO
}

func newMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:                           db,
		BotCommandsDAO:               mysql_dao.NewBotCommandsDAO(db),
		BotsDAO:                      mysql_dao.NewBotsDAO(db),
		ImportedContactsDAO:          mysql_dao.NewImportedContactsDAO(db),
		PhoneBooksDAO:                mysql_dao.NewPhoneBooksDAO(db),
		PopularContactsDAO:           mysql_dao.NewPopularContactsDAO(db),
		PredefinedUsersDAO:           mysql_dao.NewPredefinedUsersDAO(db),
		UserContactsDAO:              mysql_dao.NewUserContactsDAO(db),
		UserGlobalPrivacySettingsDAO: mysql_dao.NewUserGlobalPrivacySettingsDAO(db),
		UserNotifySettingsDAO:        mysql_dao.NewUserNotifySettingsDAO(db),
		UserPeerBlocksDAO:            mysql_dao.NewUserPeerBlocksDAO(db),
		UserPeerSettingsDAO:          mysql_dao.NewUserPeerSettingsDAO(db),
		UserPresencesDAO:             mysql_dao.NewUserPresencesDAO(db),
		UserPrivaciesDAO:             mysql_dao.NewUserPrivaciesDAO(db),
		UserSettingsDAO:              mysql_dao.NewUserSettingsDAO(db),
		UsersDAO:                     mysql_dao.NewUsersDAO(db),
		UserProfilePhotosDAO:         mysql_dao.NewUserProfilePhotosDAO(db),
		UnregisteredContactsDAO:      mysql_dao.NewUnregisteredContactsDAO(db),
		CommonDAO:                    sqlx.NewCommonDAO(db),
	}
}
