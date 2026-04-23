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
	DocumentsModel  DocumentsModel
	PhotoSizesModel PhotoSizesModel
	PhotosModel     PhotosModel
	VideoSizesModel VideoSizesModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		DocumentsModel:  NewDocumentsModel(db),
		PhotoSizesModel: NewPhotoSizesModel(db),
		PhotosModel:     NewPhotosModel(db),
		VideoSizesModel: NewVideoSizesModel(db),
	}
}
