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

var _ PredefinedUsersModel = (*customPredefinedUsersModel)(nil)

type (
	// PredefinedUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPredefinedUsersModel.
	PredefinedUsersModel interface {
		predefinedUsersModel
		bizPredefinedUsersModel
		extendPredefinedUsersModel
	}

	customPredefinedUsersModel struct {
		*defaultPredefinedUsersModel
	}
)

// NewPredefinedUsersModel returns a model for the database table.
func NewPredefinedUsersModel(db *sqlx.DB) PredefinedUsersModel {
	return &customPredefinedUsersModel{
		defaultPredefinedUsersModel: newPredefinedUsersModel(db),
	}
}
