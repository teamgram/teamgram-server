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

var _ MessageReadOutboxModel = (*customMessageReadOutboxModel)(nil)

type (
	// MessageReadOutboxModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageReadOutboxModel.
	MessageReadOutboxModel interface {
		messageReadOutboxModel
		bizMessageReadOutboxModel
		extendMessageReadOutboxModel
	}

	customMessageReadOutboxModel struct {
		*defaultMessageReadOutboxModel
	}
)

// NewMessageReadOutboxModel returns a model for the database table.
func NewMessageReadOutboxModel(db *sqlx.DB) MessageReadOutboxModel {
	return &customMessageReadOutboxModel{
		defaultMessageReadOutboxModel: newMessageReadOutboxModel(db),
	}
}
