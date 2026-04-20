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

var _ UserPrivaciesModel = (*customUserPrivaciesModel)(nil)

type (
	// UserPrivaciesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPrivaciesModel.
	UserPrivaciesModel interface {
		user_privaciesModel
		bizUserPrivaciesModel
		extendUserPrivaciesModel
	}

	customUserPrivaciesModel struct {
		*defaultUserPrivaciesModel
	}
)

// NewUserPrivaciesModel returns a model for the database table.
func NewUserPrivaciesModel(db *sqlx.DB) UserPrivaciesModel {
	return &customUserPrivaciesModel{
		defaultUserPrivaciesModel: newUserPrivaciesModel(db),
	}
}
