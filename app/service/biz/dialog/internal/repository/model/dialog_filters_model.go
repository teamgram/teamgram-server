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

var _ DialogFiltersModel = (*customDialogFiltersModel)(nil)

type (
	// DialogFiltersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDialogFiltersModel.
	DialogFiltersModel interface {
		dialogFiltersModel
		bizDialogFiltersModel
		extendDialogFiltersModel
	}

	customDialogFiltersModel struct {
		*defaultDialogFiltersModel
	}
)

// NewDialogFiltersModel returns a model for the database table.
func NewDialogFiltersModel(db *sqlx.DB) DialogFiltersModel {
	return &customDialogFiltersModel{
		defaultDialogFiltersModel: newDialogFiltersModel(db),
	}
}
