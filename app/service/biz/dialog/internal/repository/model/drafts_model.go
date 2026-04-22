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

var _ DraftsModel = (*customDraftsModel)(nil)

type (
	// DraftsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDraftsModel.
	DraftsModel interface {
		draftsModel
		bizDraftsModel
		extendDraftsModel
	}

	customDraftsModel struct {
		*defaultDraftsModel
	}
)

// NewDraftsModel returns a model for the database table.
func NewDraftsModel(db *sqlx.DB) DraftsModel {
	return &customDraftsModel{
		defaultDraftsModel: newDraftsModel(db),
	}
}
