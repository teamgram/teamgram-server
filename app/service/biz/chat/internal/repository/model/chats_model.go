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

var _ ChatsModel = (*customChatsModel)(nil)

type (
	// ChatsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatsModel.
	ChatsModel interface {
		chatsModel
		bizChatsModel
		extendChatsModel
	}

	customChatsModel struct {
		*defaultChatsModel
	}
)

// NewChatsModel returns a model for the database table.
func NewChatsModel(db *sqlx.DB) ChatsModel {
	return &customChatsModel{
		defaultChatsModel: newChatsModel(db),
	}
}
