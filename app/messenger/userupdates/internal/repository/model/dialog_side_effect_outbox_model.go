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

var _ DialogSideEffectOutboxModel = (*customDialogSideEffectOutboxModel)(nil)

type (
	// DialogSideEffectOutboxModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDialogSideEffectOutboxModel.
	DialogSideEffectOutboxModel interface {
		dialogSideEffectOutboxModel
		bizDialogSideEffectOutboxModel
		extendDialogSideEffectOutboxModel
	}

	customDialogSideEffectOutboxModel struct {
		*defaultDialogSideEffectOutboxModel
	}
)

// NewDialogSideEffectOutboxModel returns a model for the database table.
func NewDialogSideEffectOutboxModel(db *sqlx.DB) DialogSideEffectOutboxModel {
	return &customDialogSideEffectOutboxModel{
		defaultDialogSideEffectOutboxModel: newDialogSideEffectOutboxModel(db),
	}
}
