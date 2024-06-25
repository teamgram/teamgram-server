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

type UsernameDO struct {
	Id       int64  `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	PeerType int32  `db:"peer_type" json:"peer_type"`
	PeerId   int64  `db:"peer_id" json:"peer_id"`
	Deleted  bool   `db:"deleted" json:"deleted"`
}
