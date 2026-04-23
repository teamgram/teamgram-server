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

type Models struct {
	HashTagsModel       HashTagsModel
	UserPtsUpdatesModel UserPtsUpdatesModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		HashTagsModel:       NewHashTagsModel(db),
		UserPtsUpdatesModel: NewUserPtsUpdatesModel(db),
	}
}
