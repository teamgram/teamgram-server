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
	CanonicalMessagesModel      CanonicalMessagesModel
	HashTagsModel               HashTagsModel
	MessageClientRandomsModel   MessageClientRandomsModel
	MessageFanoutManifestsModel MessageFanoutManifestsModel
	MessageFanoutReceiversModel MessageFanoutReceiversModel
	MessagePeerSequencesModel   MessagePeerSequencesModel
	MessageSendStatesModel      MessageSendStatesModel
	UserPtsUpdatesModel         UserPtsUpdatesModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		CanonicalMessagesModel:      NewCanonicalMessagesModel(db),
		HashTagsModel:               NewHashTagsModel(db),
		MessageClientRandomsModel:   NewMessageClientRandomsModel(db),
		MessageFanoutManifestsModel: NewMessageFanoutManifestsModel(db),
		MessageFanoutReceiversModel: NewMessageFanoutReceiversModel(db),
		MessagePeerSequencesModel:   NewMessagePeerSequencesModel(db),
		MessageSendStatesModel:      NewMessageSendStatesModel(db),
		UserPtsUpdatesModel:         NewUserPtsUpdatesModel(db),
	}
}
