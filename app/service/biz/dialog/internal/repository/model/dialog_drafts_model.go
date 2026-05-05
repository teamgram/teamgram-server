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

var _ DialogDraftsModel = (*customDialogDraftsModel)(nil)

type (
	// DialogDraftsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDialogDraftsModel.
	DialogDraftsModel interface {
		dialogDraftsModel
		bizDialogDraftsModel
		extendDialogDraftsModel
	}

	customDialogDraftsModel struct {
		*defaultDialogDraftsModel
	}
)

// NewDialogDraftsModel returns a model for the database table.
func NewDialogDraftsModel(db *sqlx.DB) DialogDraftsModel {
	return &customDialogDraftsModel{
		defaultDialogDraftsModel: newDialogDraftsModel(db),
	}
}
