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

var _ DialogFilterTagsModel = (*customDialogFilterTagsModel)(nil)

type (
	// DialogFilterTagsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDialogFilterTagsModel.
	DialogFilterTagsModel interface {
		dialogFilterTagsModel
		bizDialogFilterTagsModel
		extendDialogFilterTagsModel
	}

	customDialogFilterTagsModel struct {
		*defaultDialogFilterTagsModel
	}
)

// NewDialogFilterTagsModel returns a model for the database table.
func NewDialogFilterTagsModel(db *sqlx.DB) DialogFilterTagsModel {
	return &customDialogFilterTagsModel{
		defaultDialogFilterTagsModel: newDialogFilterTagsModel(db),
	}
}
