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

var _ PtsUpdatesNgenModel = (*customPtsUpdatesNgenModel)(nil)

type (
	// PtsUpdatesNgenModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPtsUpdatesNgenModel.
	PtsUpdatesNgenModel interface {
		ptsUpdatesNgenModel
		bizPtsUpdatesNgenModel
		extendPtsUpdatesNgenModel
	}

	customPtsUpdatesNgenModel struct {
		*defaultPtsUpdatesNgenModel
	}
)

// NewPtsUpdatesNgenModel returns a model for the database table.
func NewPtsUpdatesNgenModel(db *sqlx.DB) PtsUpdatesNgenModel {
	return &customPtsUpdatesNgenModel{
		defaultPtsUpdatesNgenModel: newPtsUpdatesNgenModel(db),
	}
}
