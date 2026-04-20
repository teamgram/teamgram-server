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

var _ PopularContactsModel = (*customPopularContactsModel)(nil)

type (
	// PopularContactsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPopularContactsModel.
	PopularContactsModel interface {
		popular_contactsModel
		bizPopularContactsModel
		extendPopularContactsModel
	}

	customPopularContactsModel struct {
		*defaultPopularContactsModel
	}
)

// NewPopularContactsModel returns a model for the database table.
func NewPopularContactsModel(db *sqlx.DB) PopularContactsModel {
	return &customPopularContactsModel{
		defaultPopularContactsModel: newPopularContactsModel(db),
	}
}
