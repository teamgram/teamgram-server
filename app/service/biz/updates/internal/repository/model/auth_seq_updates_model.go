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

var _ AuthSeqUpdatesModel = (*customAuthSeqUpdatesModel)(nil)

type (
	// AuthSeqUpdatesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthSeqUpdatesModel.
	AuthSeqUpdatesModel interface {
		authSeqUpdatesModel
		bizAuthSeqUpdatesModel
		extendAuthSeqUpdatesModel
	}

	customAuthSeqUpdatesModel struct {
		*defaultAuthSeqUpdatesModel
	}
)

// NewAuthSeqUpdatesModel returns a model for the database table.
func NewAuthSeqUpdatesModel(db *sqlx.DB) AuthSeqUpdatesModel {
	return &customAuthSeqUpdatesModel{
		defaultAuthSeqUpdatesModel: newAuthSeqUpdatesModel(db),
	}
}
