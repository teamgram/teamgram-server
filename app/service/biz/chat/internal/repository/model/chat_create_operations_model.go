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

var _ ChatCreateOperationsModel = (*customChatCreateOperationsModel)(nil)

type (
	// ChatCreateOperationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatCreateOperationsModel.
	ChatCreateOperationsModel interface {
		chatCreateOperationsModel
		bizChatCreateOperationsModel
		extendChatCreateOperationsModel
	}

	customChatCreateOperationsModel struct {
		*defaultChatCreateOperationsModel
	}
)

// NewChatCreateOperationsModel returns a model for the database table.
func NewChatCreateOperationsModel(db *sqlx.DB) ChatCreateOperationsModel {
	return &customChatCreateOperationsModel{
		defaultChatCreateOperationsModel: newChatCreateOperationsModel(db),
	}
}
