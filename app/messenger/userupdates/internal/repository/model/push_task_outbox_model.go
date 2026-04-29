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

var _ PushTaskOutboxModel = (*customPushTaskOutboxModel)(nil)

type (
	// PushTaskOutboxModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPushTaskOutboxModel.
	PushTaskOutboxModel interface {
		pushTaskOutboxModel
		bizPushTaskOutboxModel
		extendPushTaskOutboxModel
	}

	customPushTaskOutboxModel struct {
		*defaultPushTaskOutboxModel
	}
)

// NewPushTaskOutboxModel returns a model for the database table.
func NewPushTaskOutboxModel(db *sqlx.DB) PushTaskOutboxModel {
	return &customPushTaskOutboxModel{
		defaultPushTaskOutboxModel: newPushTaskOutboxModel(db),
	}
}
