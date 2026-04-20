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

var _ PhoneBooksModel = (*customPhoneBooksModel)(nil)

type (
	// PhoneBooksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPhoneBooksModel.
	PhoneBooksModel interface {
		phoneBooksModel
		bizPhoneBooksModel
		extendPhoneBooksModel
	}

	customPhoneBooksModel struct {
		*defaultPhoneBooksModel
	}
)

// NewPhoneBooksModel returns a model for the database table.
func NewPhoneBooksModel(db *sqlx.DB) PhoneBooksModel {
	return &customPhoneBooksModel{
		defaultPhoneBooksModel: newPhoneBooksModel(db),
	}
}
