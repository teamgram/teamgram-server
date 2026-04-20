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

var _ DocumentsModel = (*customDocumentsModel)(nil)

type (
	// DocumentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDocumentsModel.
	DocumentsModel interface {
		documentsModel
		bizDocumentsModel
		extendDocumentsModel
	}

	customDocumentsModel struct {
		*defaultDocumentsModel
	}
)

// NewDocumentsModel returns a model for the database table.
func NewDocumentsModel(db *sqlx.DB) DocumentsModel {
	return &customDocumentsModel{
		defaultDocumentsModel: newDocumentsModel(db),
	}
}
