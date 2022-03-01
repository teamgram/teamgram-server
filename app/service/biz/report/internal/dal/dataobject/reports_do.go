/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type ReportsDO struct {
	Id                  int64  `db:"id"`
	UserId              int64  `db:"user_id"`
	ReportType          int32  `db:"report_type"`
	PeerType            int32  `db:"peer_type"`
	PeerId              int64  `db:"peer_id"`
	ProfilePhotoId      int64  `db:"profile_photo_id"`
	MessageSenderUserId int64  `db:"message_sender_user_id"`
	MessageId           int32  `db:"message_id"`
	Reason              int32  `db:"reason"`
	Text                string `db:"text"`
}
