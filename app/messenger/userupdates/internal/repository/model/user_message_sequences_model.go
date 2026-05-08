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

var _ UserMessageSequencesModel = (*customUserMessageSequencesModel)(nil)

type (
	// UserMessageSequencesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserMessageSequencesModel.
	UserMessageSequencesModel interface {
		userMessageSequencesModel
		bizUserMessageSequencesModel
		extendUserMessageSequencesModel
	}

	customUserMessageSequencesModel struct {
		*defaultUserMessageSequencesModel
	}
)

// NewUserMessageSequencesModel returns a model for the database table.
func NewUserMessageSequencesModel(db *sqlx.DB) UserMessageSequencesModel {
	return &customUserMessageSequencesModel{
		defaultUserMessageSequencesModel: newUserMessageSequencesModel(db),
	}
}
