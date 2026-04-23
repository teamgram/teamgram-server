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
	"github.com/teamgram/marmota/pkg/stores/cache"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ AuthKeysModel = (*customAuthKeysModel)(nil)

type (
	// AuthKeysModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthKeysModel.
	AuthKeysModel interface {
		authKeysModel
		bizAuthKeysModel
		extendAuthKeysModel
	}

	customAuthKeysModel struct {
		*defaultAuthKeysModel
	}
)

// NewAuthKeysModel returns a model for the database table.
func NewAuthKeysModel(db *sqlx.DB, c cache.CacheConf) AuthKeysModel {
	return &customAuthKeysModel{
		defaultAuthKeysModel: newAuthKeysModel(db, c),
	}
}
