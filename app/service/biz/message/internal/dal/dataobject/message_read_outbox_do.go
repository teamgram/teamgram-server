/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type MessageReadOutboxDO struct {
	Id                int64 `db:"id" json:"id"`
	UserId            int64 `db:"user_id" json:"user_id"`
	PeerDialogId      int64 `db:"peer_dialog_id" json:"peer_dialog_id"`
	ReadUserId        int64 `db:"read_user_id" json:"read_user_id"`
	ReadOutboxMaxId   int32 `db:"read_outbox_max_id" json:"read_outbox_max_id"`
	ReadOutboxMaxDate int64 `db:"read_outbox_max_date" json:"read_outbox_max_date"`
}
