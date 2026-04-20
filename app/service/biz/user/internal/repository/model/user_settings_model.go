/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026 The Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package model

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ UserSettingsModel = (*customUserSettingsModel)(nil)

type (
	// UserSettingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSettingsModel.
	UserSettingsModel interface {
		userSettingsModel
		bizUserSettingsModel
		extendUserSettingsModel
	}

	customUserSettingsModel struct {
		*defaultUserSettingsModel
	}
)

// NewUserSettingsModel returns a model for the database table.
func NewUserSettingsModel(db *sqlx.DB) UserSettingsModel {
	return &customUserSettingsModel{
		defaultUserSettingsModel: newUserSettingsModel(db),
	}
}
