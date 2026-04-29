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

var _ MessagePeerSequencesModel = (*customMessagePeerSequencesModel)(nil)

type (
	// MessagePeerSequencesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessagePeerSequencesModel.
	MessagePeerSequencesModel interface {
		messagePeerSequencesModel
		bizMessagePeerSequencesModel
		extendMessagePeerSequencesModel
	}

	customMessagePeerSequencesModel struct {
		*defaultMessagePeerSequencesModel
	}
)

// NewMessagePeerSequencesModel returns a model for the database table.
func NewMessagePeerSequencesModel(db *sqlx.DB) MessagePeerSequencesModel {
	return &customMessagePeerSequencesModel{
		defaultMessagePeerSequencesModel: newMessagePeerSequencesModel(db),
	}
}
