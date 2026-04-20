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

var _ UserProfilePhotosModel = (*customUserProfilePhotosModel)(nil)

type (
	// UserProfilePhotosModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserProfilePhotosModel.
	UserProfilePhotosModel interface {
		user_profile_photosModel
		bizUserProfilePhotosModel
		extendUserProfilePhotosModel
	}

	customUserProfilePhotosModel struct {
		*defaultUserProfilePhotosModel
	}
)

// NewUserProfilePhotosModel returns a model for the database table.
func NewUserProfilePhotosModel(db *sqlx.DB) UserProfilePhotosModel {
	return &customUserProfilePhotosModel{
		defaultUserProfilePhotosModel: newUserProfilePhotosModel(db),
	}
}
