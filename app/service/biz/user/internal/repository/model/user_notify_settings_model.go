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

var _ UserNotifySettingsModel = (*customUserNotifySettingsModel)(nil)

type (
	// UserNotifySettingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserNotifySettingsModel.
	UserNotifySettingsModel interface {
		userNotifySettingsModel
		bizUserNotifySettingsModel
		extendUserNotifySettingsModel
	}

	customUserNotifySettingsModel struct {
		*defaultUserNotifySettingsModel
	}
)

// NewUserNotifySettingsModel returns a model for the database table.
func NewUserNotifySettingsModel(db *sqlx.DB) UserNotifySettingsModel {
	return &customUserNotifySettingsModel{
		defaultUserNotifySettingsModel: newUserNotifySettingsModel(db),
	}
}
