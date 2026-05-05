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

var _ DialogPreferenceVersionsModel = (*customDialogPreferenceVersionsModel)(nil)

type (
	// DialogPreferenceVersionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDialogPreferenceVersionsModel.
	DialogPreferenceVersionsModel interface {
		dialogPreferenceVersionsModel
		bizDialogPreferenceVersionsModel
		extendDialogPreferenceVersionsModel
	}

	customDialogPreferenceVersionsModel struct {
		*defaultDialogPreferenceVersionsModel
	}
)

// NewDialogPreferenceVersionsModel returns a model for the database table.
func NewDialogPreferenceVersionsModel(db *sqlx.DB) DialogPreferenceVersionsModel {
	return &customDialogPreferenceVersionsModel{
		defaultDialogPreferenceVersionsModel: newDialogPreferenceVersionsModel(db),
	}
}
