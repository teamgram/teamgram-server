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

var _ UserSavedMusicModel = (*customUserSavedMusicModel)(nil)

type (
	// UserSavedMusicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSavedMusicModel.
	UserSavedMusicModel interface {
		userSavedMusicModel
		bizUserSavedMusicModel
		extendUserSavedMusicModel
	}

	customUserSavedMusicModel struct {
		*defaultUserSavedMusicModel
	}
)

// NewUserSavedMusicModel returns a model for the database table.
func NewUserSavedMusicModel(db *sqlx.DB) UserSavedMusicModel {
	return &customUserSavedMusicModel{
		defaultUserSavedMusicModel: newUserSavedMusicModel(db),
	}
}
