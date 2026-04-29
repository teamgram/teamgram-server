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

var _ UserPtsStateModel = (*customUserPtsStateModel)(nil)

type (
	// UserPtsStateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPtsStateModel.
	UserPtsStateModel interface {
		userPtsStateModel
		bizUserPtsStateModel
		extendUserPtsStateModel
	}

	customUserPtsStateModel struct {
		*defaultUserPtsStateModel
	}
)

// NewUserPtsStateModel returns a model for the database table.
func NewUserPtsStateModel(db *sqlx.DB) UserPtsStateModel {
	return &customUserPtsStateModel{
		defaultUserPtsStateModel: newUserPtsStateModel(db),
	}
}
