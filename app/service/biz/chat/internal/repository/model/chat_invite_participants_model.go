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

var _ ChatInviteParticipantsModel = (*customChatInviteParticipantsModel)(nil)

type (
	// ChatInviteParticipantsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatInviteParticipantsModel.
	ChatInviteParticipantsModel interface {
		chatInviteParticipantsModel
		bizChatInviteParticipantsModel
		extendChatInviteParticipantsModel
	}

	customChatInviteParticipantsModel struct {
		*defaultChatInviteParticipantsModel
	}
)

// NewChatInviteParticipantsModel returns a model for the database table.
func NewChatInviteParticipantsModel(db *sqlx.DB) ChatInviteParticipantsModel {
	return &customChatInviteParticipantsModel{
		defaultChatInviteParticipantsModel: newChatInviteParticipantsModel(db),
	}
}
