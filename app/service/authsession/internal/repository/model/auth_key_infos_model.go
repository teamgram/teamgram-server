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

var _ AuthKeyInfosModel = (*customAuthKeyInfosModel)(nil)

type (
	// AuthKeyInfosModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthKeyInfosModel.
	AuthKeyInfosModel interface {
		auth_key_infosModel
		bizAuthKeyInfosModel
		extendAuthKeyInfosModel
	}

	customAuthKeyInfosModel struct {
		*defaultAuthKeyInfosModel
	}
)

// NewAuthKeyInfosModel returns a model for the database table.
func NewAuthKeyInfosModel(db *sqlx.DB) AuthKeyInfosModel {
	return &customAuthKeyInfosModel{
		defaultAuthKeyInfosModel: newAuthKeyInfosModel(db),
	}
}
