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
	ChatParticipantsModel  ChatParticipantsModel
	ChatsModel             ChatsModel
	HashTagsModel          HashTagsModel
	MessageReadOutboxModel MessageReadOutboxModel
}

type TxModels struct {
	ChatParticipantsModel  ChatParticipantsTxModel
	ChatsModel             ChatsTxModel
	HashTagsModel          HashTagsTxModel
	MessageReadOutboxModel MessageReadOutboxTxModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		ChatParticipantsModel:  NewChatParticipantsModel(db),
		ChatsModel:             NewChatsModel(db),
		HashTagsModel:          NewHashTagsModel(db),
		MessageReadOutboxModel: NewMessageReadOutboxModel(db),
	}
}

func (m *Models) WithTx(tx *sqlx.Tx) *TxModels {
	return &TxModels{
		ChatParticipantsModel:  NewChatParticipantsTxModel(tx),
		ChatsModel:             NewChatsTxModel(tx),
		HashTagsModel:          NewHashTagsTxModel(tx),
		MessageReadOutboxModel: NewMessageReadOutboxTxModel(tx),
	}
}
