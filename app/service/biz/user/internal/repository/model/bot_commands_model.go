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

var _ BotCommandsModel = (*customBotCommandsModel)(nil)

type (
	// BotCommandsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBotCommandsModel.
	BotCommandsModel interface {
		botCommandsModel
		bizBotCommandsModel
		extendBotCommandsModel
	}

	customBotCommandsModel struct {
		*defaultBotCommandsModel
	}
)

// NewBotCommandsModel returns a model for the database table.
func NewBotCommandsModel(db *sqlx.DB) BotCommandsModel {
	return &customBotCommandsModel{
		defaultBotCommandsModel: newBotCommandsModel(db),
	}
}
