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

var _ BotsModel = (*customBotsModel)(nil)

type (
	// BotsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBotsModel.
	BotsModel interface {
		botsModel
		bizBotsModel
		extendBotsModel
	}

	customBotsModel struct {
		*defaultBotsModel
	}
)

// NewBotsModel returns a model for the database table.
func NewBotsModel(db *sqlx.DB) BotsModel {
	return &customBotsModel{
		defaultBotsModel: newBotsModel(db),
	}
}
