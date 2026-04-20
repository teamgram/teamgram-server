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

var _ UnregisteredContactsModel = (*customUnregisteredContactsModel)(nil)

type (
	// UnregisteredContactsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUnregisteredContactsModel.
	UnregisteredContactsModel interface {
		unregistered_contactsModel
		bizUnregisteredContactsModel
		extendUnregisteredContactsModel
	}

	customUnregisteredContactsModel struct {
		*defaultUnregisteredContactsModel
	}
)

// NewUnregisteredContactsModel returns a model for the database table.
func NewUnregisteredContactsModel(db *sqlx.DB) UnregisteredContactsModel {
	return &customUnregisteredContactsModel{
		defaultUnregisteredContactsModel: newUnregisteredContactsModel(db),
	}
}
