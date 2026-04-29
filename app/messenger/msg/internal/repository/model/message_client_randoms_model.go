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

var _ MessageClientRandomsModel = (*customMessageClientRandomsModel)(nil)

type (
	// MessageClientRandomsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageClientRandomsModel.
	MessageClientRandomsModel interface {
		messageClientRandomsModel
		bizMessageClientRandomsModel
		extendMessageClientRandomsModel
	}

	customMessageClientRandomsModel struct {
		*defaultMessageClientRandomsModel
	}
)

// NewMessageClientRandomsModel returns a model for the database table.
func NewMessageClientRandomsModel(db *sqlx.DB) MessageClientRandomsModel {
	return &customMessageClientRandomsModel{
		defaultMessageClientRandomsModel: newMessageClientRandomsModel(db),
	}
}
