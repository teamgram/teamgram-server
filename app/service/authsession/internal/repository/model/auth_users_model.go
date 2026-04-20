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

var _ AuthUsersModel = (*customAuthUsersModel)(nil)

type (
	// AuthUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthUsersModel.
	AuthUsersModel interface {
		authUsersModel
		bizAuthUsersModel
		extendAuthUsersModel
	}

	customAuthUsersModel struct {
		*defaultAuthUsersModel
	}
)

// NewAuthUsersModel returns a model for the database table.
func NewAuthUsersModel(db *sqlx.DB) AuthUsersModel {
	return &customAuthUsersModel{
		defaultAuthUsersModel: newAuthUsersModel(db),
	}
}
