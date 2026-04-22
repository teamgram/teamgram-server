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

var _ PhotoSizesModel = (*customPhotoSizesModel)(nil)

type (
	// PhotoSizesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPhotoSizesModel.
	PhotoSizesModel interface {
		photoSizesModel
		bizPhotoSizesModel
		extendPhotoSizesModel
	}

	customPhotoSizesModel struct {
		*defaultPhotoSizesModel
	}
)

// NewPhotoSizesModel returns a model for the database table.
func NewPhotoSizesModel(db *sqlx.DB) PhotoSizesModel {
	return &customPhotoSizesModel{
		defaultPhotoSizesModel: newPhotoSizesModel(db),
	}
}
