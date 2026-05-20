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

var _ AuthSeqStateModel = (*customAuthSeqStateModel)(nil)

type (
	// AuthSeqStateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthSeqStateModel.
	AuthSeqStateModel interface {
		authSeqStateModel
		bizAuthSeqStateModel
		extendAuthSeqStateModel
	}

	customAuthSeqStateModel struct {
		*defaultAuthSeqStateModel
	}
)

// NewAuthSeqStateModel returns a model for the database table.
func NewAuthSeqStateModel(db *sqlx.DB) AuthSeqStateModel {
	return &customAuthSeqStateModel{
		defaultAuthSeqStateModel: newAuthSeqStateModel(db),
	}
}
