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

type HashTagsDO struct {
	Id               int64  `db:"id"`
	UserId           int64  `db:"user_id"`
	PeerType         int32  `db:"peer_type"`
	PeerId           int64  `db:"peer_id"`
	HashTag          string `db:"hash_tag"`
	HashTagMessageId int32  `db:"hash_tag_message_id"`
	Deleted          bool   `db:"deleted"`
}
