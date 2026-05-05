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

var _ DialogPreferencesModel = (*customDialogPreferencesModel)(nil)

type (
	// DialogPreferencesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDialogPreferencesModel.
	DialogPreferencesModel interface {
		dialogPreferencesModel
		bizDialogPreferencesModel
		extendDialogPreferencesModel
	}

	customDialogPreferencesModel struct {
		*defaultDialogPreferencesModel
	}
)

// NewDialogPreferencesModel returns a model for the database table.
func NewDialogPreferencesModel(db *sqlx.DB) DialogPreferencesModel {
	return &customDialogPreferencesModel{
		defaultDialogPreferencesModel: newDialogPreferencesModel(db),
	}
}
