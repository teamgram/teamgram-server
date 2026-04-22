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

var _ DialogsModel = (*customDialogsModel)(nil)

type (
	// DialogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDialogsModel.
	DialogsModel interface {
		dialogsModel
		bizDialogsModel
		extendDialogsModel
	}

	customDialogsModel struct {
		*defaultDialogsModel
	}
)

// NewDialogsModel returns a model for the database table.
func NewDialogsModel(db *sqlx.DB) DialogsModel {
	return &customDialogsModel{
		defaultDialogsModel: newDialogsModel(db),
	}
}
