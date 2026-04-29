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

var _ CanonicalMessagesModel = (*customCanonicalMessagesModel)(nil)

type (
	// CanonicalMessagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCanonicalMessagesModel.
	CanonicalMessagesModel interface {
		canonicalMessagesModel
		bizCanonicalMessagesModel
		extendCanonicalMessagesModel
	}

	customCanonicalMessagesModel struct {
		*defaultCanonicalMessagesModel
	}
)

// NewCanonicalMessagesModel returns a model for the database table.
func NewCanonicalMessagesModel(db *sqlx.DB) CanonicalMessagesModel {
	return &customCanonicalMessagesModel{
		defaultCanonicalMessagesModel: newCanonicalMessagesModel(db),
	}
}
