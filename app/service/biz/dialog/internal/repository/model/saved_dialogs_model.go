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

var _ SavedDialogsModel = (*customSavedDialogsModel)(nil)

type (
	// SavedDialogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSavedDialogsModel.
	SavedDialogsModel interface {
		saved_dialogsModel
		bizSavedDialogsModel
		extendSavedDialogsModel
	}

	customSavedDialogsModel struct {
		*defaultSavedDialogsModel
	}
)

// NewSavedDialogsModel returns a model for the database table.
func NewSavedDialogsModel(db *sqlx.DB) SavedDialogsModel {
	return &customSavedDialogsModel{
		defaultSavedDialogsModel: newSavedDialogsModel(db),
	}
}
