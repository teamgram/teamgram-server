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

var _ UserPtsEventsModel = (*customUserPtsEventsModel)(nil)

type (
	// UserPtsEventsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPtsEventsModel.
	UserPtsEventsModel interface {
		userPtsEventsModel
		bizUserPtsEventsModel
		extendUserPtsEventsModel
	}

	customUserPtsEventsModel struct {
		*defaultUserPtsEventsModel
	}
)

// NewUserPtsEventsModel returns a model for the database table.
func NewUserPtsEventsModel(db *sqlx.DB) UserPtsEventsModel {
	return &customUserPtsEventsModel{
		defaultUserPtsEventsModel: newUserPtsEventsModel(db),
	}
}
