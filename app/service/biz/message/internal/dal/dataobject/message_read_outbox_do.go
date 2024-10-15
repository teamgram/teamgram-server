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
	Id         int64 `db:"id" json:"id"`
	UserId     int64 `db:"user_id" json:"user_id"`
	MessageId  int32 `db:"message_id" json:"message_id"`
	PeerType   int32 `db:"peer_type" json:"peer_type"`
	PeerId     int32 `db:"peer_id" json:"peer_id"`
	ReadUserId int64 `db:"read_user_id" json:"read_user_id"`
	ReadDate   int64 `db:"read_date" json:"read_date"`
}
