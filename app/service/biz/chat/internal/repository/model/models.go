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
	ChatInviteParticipantsModel ChatInviteParticipantsModel
	ChatInvitesModel            ChatInvitesModel
	ChatParticipantsModel       ChatParticipantsModel
	ChatsModel                  ChatsModel
}

type TxModels struct {
	ChatInviteParticipantsModel ChatInviteParticipantsTxModel
	ChatInvitesModel            ChatInvitesTxModel
	ChatParticipantsModel       ChatParticipantsTxModel
	ChatsModel                  ChatsTxModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		ChatInviteParticipantsModel: NewChatInviteParticipantsModel(db),
		ChatInvitesModel:            NewChatInvitesModel(db),
		ChatParticipantsModel:       NewChatParticipantsModel(db),
		ChatsModel:                  NewChatsModel(db),
	}
}

func (m *Models) WithTx(tx *sqlx.Tx) *TxModels {
	return &TxModels{
		ChatInviteParticipantsModel: NewChatInviteParticipantsTxModel(tx),
		ChatInvitesModel:            NewChatInvitesTxModel(tx),
		ChatParticipantsModel:       NewChatParticipantsTxModel(tx),
		ChatsModel:                  NewChatsTxModel(tx),
	}
}
