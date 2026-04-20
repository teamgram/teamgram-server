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

var _ ChatInvitesModel = (*customChatInvitesModel)(nil)

type (
	// ChatInvitesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatInvitesModel.
	ChatInvitesModel interface {
		chat_invitesModel
		bizChatInvitesModel
		extendChatInvitesModel
	}

	customChatInvitesModel struct {
		*defaultChatInvitesModel
	}
)

// NewChatInvitesModel returns a model for the database table.
func NewChatInvitesModel(db *sqlx.DB) ChatInvitesModel {
	return &customChatInvitesModel{
		defaultChatInvitesModel: newChatInvitesModel(db),
	}
}
