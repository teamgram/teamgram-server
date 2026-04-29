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

var _ MessageFanoutReceiversModel = (*customMessageFanoutReceiversModel)(nil)

type (
	// MessageFanoutReceiversModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageFanoutReceiversModel.
	MessageFanoutReceiversModel interface {
		messageFanoutReceiversModel
		bizMessageFanoutReceiversModel
		extendMessageFanoutReceiversModel
	}

	customMessageFanoutReceiversModel struct {
		*defaultMessageFanoutReceiversModel
	}
)

// NewMessageFanoutReceiversModel returns a model for the database table.
func NewMessageFanoutReceiversModel(db *sqlx.DB) MessageFanoutReceiversModel {
	return &customMessageFanoutReceiversModel{
		defaultMessageFanoutReceiversModel: newMessageFanoutReceiversModel(db),
	}
}
