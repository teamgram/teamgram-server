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

var _ UserPeerBlocksModel = (*customUserPeerBlocksModel)(nil)

type (
	// UserPeerBlocksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPeerBlocksModel.
	UserPeerBlocksModel interface {
		userPeerBlocksModel
		bizUserPeerBlocksModel
		extendUserPeerBlocksModel
	}

	customUserPeerBlocksModel struct {
		*defaultUserPeerBlocksModel
	}
)

// NewUserPeerBlocksModel returns a model for the database table.
func NewUserPeerBlocksModel(db *sqlx.DB) UserPeerBlocksModel {
	return &customUserPeerBlocksModel{
		defaultUserPeerBlocksModel: newUserPeerBlocksModel(db),
	}
}
