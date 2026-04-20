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

var _ UserContactsModel = (*customUserContactsModel)(nil)

type (
	// UserContactsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserContactsModel.
	UserContactsModel interface {
		user_contactsModel
		bizUserContactsModel
		extendUserContactsModel
	}

	customUserContactsModel struct {
		*defaultUserContactsModel
	}
)

// NewUserContactsModel returns a model for the database table.
func NewUserContactsModel(db *sqlx.DB) UserContactsModel {
	return &customUserContactsModel{
		defaultUserContactsModel: newUserContactsModel(db),
	}
}
