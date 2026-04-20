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

var _ HashTagsModel = (*customHashTagsModel)(nil)

type (
	// HashTagsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHashTagsModel.
	HashTagsModel interface {
		hashTagsModel
		bizHashTagsModel
		extendHashTagsModel
	}

	customHashTagsModel struct {
		*defaultHashTagsModel
	}
)

// NewHashTagsModel returns a model for the database table.
func NewHashTagsModel(db *sqlx.DB) HashTagsModel {
	return &customHashTagsModel{
		defaultHashTagsModel: newHashTagsModel(db),
	}
}
