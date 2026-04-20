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

var _ VideoSizesModel = (*customVideoSizesModel)(nil)

type (
	// VideoSizesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoSizesModel.
	VideoSizesModel interface {
		videoSizesModel
		bizVideoSizesModel
		extendVideoSizesModel
	}

	customVideoSizesModel struct {
		*defaultVideoSizesModel
	}
)

// NewVideoSizesModel returns a model for the database table.
func NewVideoSizesModel(db *sqlx.DB) VideoSizesModel {
	return &customVideoSizesModel{
		defaultVideoSizesModel: newVideoSizesModel(db),
	}
}
