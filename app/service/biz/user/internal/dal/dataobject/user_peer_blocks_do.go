/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type UserPeerBlocksDO struct {
	Id       int64 `db:"id"`
	UserId   int64 `db:"user_id"`
	PeerType int32 `db:"peer_type"`
	PeerId   int64 `db:"peer_id"`
	Date     int64 `db:"date"`
	Deleted  bool  `db:"deleted"`
}
