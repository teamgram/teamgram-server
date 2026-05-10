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

var _ FileReferencesModel = (*customFileReferencesModel)(nil)

type (
	// FileReferencesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFileReferencesModel.
	FileReferencesModel interface {
		fileReferencesModel
		bizFileReferencesModel
		extendFileReferencesModel
	}

	customFileReferencesModel struct {
		*defaultFileReferencesModel
	}
)

// NewFileReferencesModel returns a model for the database table.
func NewFileReferencesModel(db *sqlx.DB) FileReferencesModel {
	return &customFileReferencesModel{
		defaultFileReferencesModel: newFileReferencesModel(db),
	}
}
