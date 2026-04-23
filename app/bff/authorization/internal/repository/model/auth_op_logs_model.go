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

var _ AuthOpLogsModel = (*customAuthOpLogsModel)(nil)

type (
	// AuthOpLogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthOpLogsModel.
	AuthOpLogsModel interface {
		authOpLogsModel
		bizAuthOpLogsModel
		extendAuthOpLogsModel
	}

	customAuthOpLogsModel struct {
		*defaultAuthOpLogsModel
	}
)

// NewAuthOpLogsModel returns a model for the database table.
func NewAuthOpLogsModel(db *sqlx.DB) AuthOpLogsModel {
	return &customAuthOpLogsModel{
		defaultAuthOpLogsModel: newAuthOpLogsModel(db),
	}
}
