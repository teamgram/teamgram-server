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

var _ UserGlobalPrivacySettingsModel = (*customUserGlobalPrivacySettingsModel)(nil)

type (
	// UserGlobalPrivacySettingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserGlobalPrivacySettingsModel.
	UserGlobalPrivacySettingsModel interface {
		userGlobalPrivacySettingsModel
		bizUserGlobalPrivacySettingsModel
		extendUserGlobalPrivacySettingsModel
	}

	customUserGlobalPrivacySettingsModel struct {
		*defaultUserGlobalPrivacySettingsModel
	}
)

// NewUserGlobalPrivacySettingsModel returns a model for the database table.
func NewUserGlobalPrivacySettingsModel(db *sqlx.DB) UserGlobalPrivacySettingsModel {
	return &customUserGlobalPrivacySettingsModel{
		defaultUserGlobalPrivacySettingsModel: newUserGlobalPrivacySettingsModel(db),
	}
}
