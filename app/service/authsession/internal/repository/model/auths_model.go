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

var _ AuthsModel = (*customAuthsModel)(nil)

type (
	// AuthsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthsModel.
	AuthsModel interface {
		authsModel
		bizAuthsModel
		extendAuthsModel
	}

	customAuthsModel struct {
		*defaultAuthsModel
	}
)

// NewAuthsModel returns a model for the database table.
func NewAuthsModel(db *sqlx.DB) AuthsModel {
	return &customAuthsModel{
		defaultAuthsModel: newAuthsModel(db),
	}
}
