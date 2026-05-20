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

var _ AuthUpdatePayloadsModel = (*customAuthUpdatePayloadsModel)(nil)

type (
	// AuthUpdatePayloadsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthUpdatePayloadsModel.
	AuthUpdatePayloadsModel interface {
		authUpdatePayloadsModel
		bizAuthUpdatePayloadsModel
		extendAuthUpdatePayloadsModel
	}

	customAuthUpdatePayloadsModel struct {
		*defaultAuthUpdatePayloadsModel
	}
)

// NewAuthUpdatePayloadsModel returns a model for the database table.
func NewAuthUpdatePayloadsModel(db *sqlx.DB) AuthUpdatePayloadsModel {
	return &customAuthUpdatePayloadsModel{
		defaultAuthUpdatePayloadsModel: newAuthUpdatePayloadsModel(db),
	}
}
