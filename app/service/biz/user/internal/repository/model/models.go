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

type TxModels struct {
	BotCommandsModel               BotCommandsTxModel
	BotsModel                      BotsTxModel
	DefaultHistoryTtlModel         DefaultHistoryTtlTxModel
	ImportedContactsModel          ImportedContactsTxModel
	PhoneBooksModel                PhoneBooksTxModel
	PopularContactsModel           PopularContactsTxModel
	PredefinedUsersModel           PredefinedUsersTxModel
	UnregisteredContactsModel      UnregisteredContactsTxModel
	UserContactsModel              UserContactsTxModel
	UserGlobalPrivacySettingsModel UserGlobalPrivacySettingsTxModel
	UserNotifySettingsModel        UserNotifySettingsTxModel
	UserPeerBlocksModel            UserPeerBlocksTxModel
	UserPeerSettingsModel          UserPeerSettingsTxModel
	UserPresencesModel             UserPresencesTxModel
	UserPrivaciesModel             UserPrivaciesTxModel
	UserProfilePhotosModel         UserProfilePhotosTxModel
	UserSavedMusicModel            UserSavedMusicTxModel
	UserSettingsModel              UserSettingsTxModel
	UsernameModel                  UsernameTxModel
	UsersModel                     UsersTxModel
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

func (m *Models) WithTx(tx *sqlx.Tx) *TxModels {
	return &TxModels{
		BotCommandsModel:               NewBotCommandsTxModel(tx),
		BotsModel:                      NewBotsTxModel(tx),
		DefaultHistoryTtlModel:         NewDefaultHistoryTtlTxModel(tx),
		ImportedContactsModel:          NewImportedContactsTxModel(tx),
		PhoneBooksModel:                NewPhoneBooksTxModel(tx),
		PopularContactsModel:           NewPopularContactsTxModel(tx),
		PredefinedUsersModel:           NewPredefinedUsersTxModel(tx),
		UnregisteredContactsModel:      NewUnregisteredContactsTxModel(tx),
		UserContactsModel:              NewUserContactsTxModel(tx),
		UserGlobalPrivacySettingsModel: NewUserGlobalPrivacySettingsTxModel(tx),
		UserNotifySettingsModel:        NewUserNotifySettingsTxModel(tx),
		UserPeerBlocksModel:            NewUserPeerBlocksTxModel(tx),
		UserPeerSettingsModel:          NewUserPeerSettingsTxModel(tx),
		UserPresencesModel:             NewUserPresencesTxModel(tx),
		UserPrivaciesModel:             NewUserPrivaciesTxModel(tx),
		UserProfilePhotosModel:         NewUserProfilePhotosTxModel(tx),
		UserSavedMusicModel:            NewUserSavedMusicTxModel(tx),
		UserSettingsModel:              NewUserSettingsTxModel(tx),
		UsernameModel:                  NewUsernameTxModel(tx),
		UsersModel:                     NewUsersTxModel(tx),
	}
}
