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

type UserPeerBlocksDO struct {
	Id       int64 `db:"id" json:"id"`
	UserId   int64 `db:"user_id" json:"user_id"`
	PeerType int32 `db:"peer_type" json:"peer_type"`
	PeerId   int64 `db:"peer_id" json:"peer_id"`
	Date     int64 `db:"date" json:"date"`
	Deleted  bool  `db:"deleted" json:"deleted"`
}
