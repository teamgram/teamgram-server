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

var _ MessageFanoutManifestsModel = (*customMessageFanoutManifestsModel)(nil)

type (
	// MessageFanoutManifestsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageFanoutManifestsModel.
	MessageFanoutManifestsModel interface {
		messageFanoutManifestsModel
		bizMessageFanoutManifestsModel
		extendMessageFanoutManifestsModel
	}

	customMessageFanoutManifestsModel struct {
		*defaultMessageFanoutManifestsModel
	}
)

// NewMessageFanoutManifestsModel returns a model for the database table.
func NewMessageFanoutManifestsModel(db *sqlx.DB) MessageFanoutManifestsModel {
	return &customMessageFanoutManifestsModel{
		defaultMessageFanoutManifestsModel: newMessageFanoutManifestsModel(db),
	}
}
