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
	Editable bool   `db:"editable" json:"editable"`
	Active   bool   `db:"active" json:"active"`
	Order2   int64  `db:"order2" json:"order2"`
	Deleted  bool   `db:"deleted" json:"deleted"`
}
