/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026 The Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package model

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ UserPtsUpdatesModel = (*customUserPtsUpdatesModel)(nil)

type (
	// UserPtsUpdatesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPtsUpdatesModel.
	UserPtsUpdatesModel interface {
		user_pts_updatesModel
		bizUserPtsUpdatesModel
		extendUserPtsUpdatesModel
	}

	customUserPtsUpdatesModel struct {
		*defaultUserPtsUpdatesModel
	}
)

// NewUserPtsUpdatesModel returns a model for the database table.
func NewUserPtsUpdatesModel(db *sqlx.DB) UserPtsUpdatesModel {
	return &customUserPtsUpdatesModel{
		defaultUserPtsUpdatesModel: newUserPtsUpdatesModel(db),
	}
}
