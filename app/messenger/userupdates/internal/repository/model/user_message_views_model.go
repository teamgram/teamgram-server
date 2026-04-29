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

var _ UserMessageViewsModel = (*customUserMessageViewsModel)(nil)

type (
	// UserMessageViewsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserMessageViewsModel.
	UserMessageViewsModel interface {
		userMessageViewsModel
		bizUserMessageViewsModel
		extendUserMessageViewsModel
	}

	customUserMessageViewsModel struct {
		*defaultUserMessageViewsModel
	}
)

// NewUserMessageViewsModel returns a model for the database table.
func NewUserMessageViewsModel(db *sqlx.DB) UserMessageViewsModel {
	return &customUserMessageViewsModel{
		defaultUserMessageViewsModel: newUserMessageViewsModel(db),
	}
}
