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

var _ UserDialogsModel = (*customUserDialogsModel)(nil)

type (
	// UserDialogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserDialogsModel.
	UserDialogsModel interface {
		userDialogsModel
		bizUserDialogsModel
		extendUserDialogsModel
	}

	customUserDialogsModel struct {
		*defaultUserDialogsModel
	}
)

// NewUserDialogsModel returns a model for the database table.
func NewUserDialogsModel(db *sqlx.DB) UserDialogsModel {
	return &customUserDialogsModel{
		defaultUserDialogsModel: newUserDialogsModel(db),
	}
}
