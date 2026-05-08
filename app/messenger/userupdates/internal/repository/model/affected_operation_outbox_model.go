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

var _ AffectedOperationOutboxModel = (*customAffectedOperationOutboxModel)(nil)

type (
	// AffectedOperationOutboxModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAffectedOperationOutboxModel.
	AffectedOperationOutboxModel interface {
		affectedOperationOutboxModel
		bizAffectedOperationOutboxModel
		extendAffectedOperationOutboxModel
	}

	customAffectedOperationOutboxModel struct {
		*defaultAffectedOperationOutboxModel
	}
)

// NewAffectedOperationOutboxModel returns a model for the database table.
func NewAffectedOperationOutboxModel(db *sqlx.DB) AffectedOperationOutboxModel {
	return &customAffectedOperationOutboxModel{
		defaultAffectedOperationOutboxModel: newAffectedOperationOutboxModel(db),
	}
}
