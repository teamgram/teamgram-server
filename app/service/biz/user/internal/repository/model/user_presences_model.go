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

var _ UserPresencesModel = (*customUserPresencesModel)(nil)

type (
	// UserPresencesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPresencesModel.
	UserPresencesModel interface {
		userPresencesModel
		bizUserPresencesModel
		extendUserPresencesModel
	}

	customUserPresencesModel struct {
		*defaultUserPresencesModel
	}
)

// NewUserPresencesModel returns a model for the database table.
func NewUserPresencesModel(db *sqlx.DB) UserPresencesModel {
	return &customUserPresencesModel{
		defaultUserPresencesModel: newUserPresencesModel(db),
	}
}
