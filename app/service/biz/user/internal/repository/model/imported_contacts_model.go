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

var _ ImportedContactsModel = (*customImportedContactsModel)(nil)

type (
	// ImportedContactsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customImportedContactsModel.
	ImportedContactsModel interface {
		importedContactsModel
		bizImportedContactsModel
		extendImportedContactsModel
	}

	customImportedContactsModel struct {
		*defaultImportedContactsModel
	}
)

// NewImportedContactsModel returns a model for the database table.
func NewImportedContactsModel(db *sqlx.DB) ImportedContactsModel {
	return &customImportedContactsModel{
		defaultImportedContactsModel: newImportedContactsModel(db),
	}
}
