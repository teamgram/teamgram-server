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

var _ PhotoSizesModel = (*customPhotoSizesModel)(nil)

type (
	// PhotoSizesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPhotoSizesModel.
	PhotoSizesModel interface {
		photo_sizesModel
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
