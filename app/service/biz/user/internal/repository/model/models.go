/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package model

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

type Models struct {
	BotCommandsModel               BotCommandsModel
	BotsModel                      BotsModel
	DefaultHistoryTtlModel         DefaultHistoryTtlModel
	ImportedContactsModel          ImportedContactsModel
	PhoneBooksModel                PhoneBooksModel
	PopularContactsModel           PopularContactsModel
	PredefinedUsersModel           PredefinedUsersModel
	UnregisteredContactsModel      UnregisteredContactsModel
	UserContactsModel              UserContactsModel
	UserGlobalPrivacySettingsModel UserGlobalPrivacySettingsModel
	UserNotifySettingsModel        UserNotifySettingsModel
	UserPeerBlocksModel            UserPeerBlocksModel
	UserPeerSettingsModel          UserPeerSettingsModel
	UserPresencesModel             UserPresencesModel
	UserPrivaciesModel             UserPrivaciesModel
	UserProfilePhotosModel         UserProfilePhotosModel
	UserSavedMusicModel            UserSavedMusicModel
	UserSettingsModel              UserSettingsModel
	UsernameModel                  UsernameModel
	UsersModel                     UsersModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		BotCommandsModel:               NewBotCommandsModel(db),
		BotsModel:                      NewBotsModel(db),
		DefaultHistoryTtlModel:         NewDefaultHistoryTtlModel(db),
		ImportedContactsModel:          NewImportedContactsModel(db),
		PhoneBooksModel:                NewPhoneBooksModel(db),
		PopularContactsModel:           NewPopularContactsModel(db),
		PredefinedUsersModel:           NewPredefinedUsersModel(db),
		UnregisteredContactsModel:      NewUnregisteredContactsModel(db),
		UserContactsModel:              NewUserContactsModel(db),
		UserGlobalPrivacySettingsModel: NewUserGlobalPrivacySettingsModel(db),
		UserNotifySettingsModel:        NewUserNotifySettingsModel(db),
		UserPeerBlocksModel:            NewUserPeerBlocksModel(db),
		UserPeerSettingsModel:          NewUserPeerSettingsModel(db),
		UserPresencesModel:             NewUserPresencesModel(db),
		UserPrivaciesModel:             NewUserPrivaciesModel(db),
		UserProfilePhotosModel:         NewUserProfilePhotosModel(db),
		UserSavedMusicModel:            NewUserSavedMusicModel(db),
		UserSettingsModel:              NewUserSettingsModel(db),
		UsernameModel:                  NewUsernameModel(db),
		UsersModel:                     NewUsersModel(db),
	}
}
