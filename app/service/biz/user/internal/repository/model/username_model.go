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

var _ UsernameModel = (*customUsernameModel)(nil)

type (
	// UsernameModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsernameModel.
	UsernameModel interface {
		usernameModel
		bizUsernameModel
		extendUsernameModel
	}

	customUsernameModel struct {
		*defaultUsernameModel
	}
)

// NewUsernameModel returns a model for the database table.
func NewUsernameModel(db *sqlx.DB) UsernameModel {
	return &customUsernameModel{
		defaultUsernameModel: newUsernameModel(db),
	}
}
