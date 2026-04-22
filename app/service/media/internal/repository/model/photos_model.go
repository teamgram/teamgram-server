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

var _ PhotosModel = (*customPhotosModel)(nil)

type (
	// PhotosModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPhotosModel.
	PhotosModel interface {
		photosModel
		bizPhotosModel
		extendPhotosModel
	}

	customPhotosModel struct {
		*defaultPhotosModel
	}
)

// NewPhotosModel returns a model for the database table.
func NewPhotosModel(db *sqlx.DB) PhotosModel {
	return &customPhotosModel{
		defaultPhotosModel: newPhotosModel(db),
	}
}
