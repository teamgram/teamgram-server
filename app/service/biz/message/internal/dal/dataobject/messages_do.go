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

type MessagesDO struct {
	Id                int64  `db:"id"`
	UserId            int64  `db:"user_id"`
	UserMessageBoxId  int32  `db:"user_message_box_id"`
	DialogId1         int64  `db:"dialog_id1"`
	DialogId2         int64  `db:"dialog_id2"`
	DialogMessageId   int64  `db:"dialog_message_id"`
	SenderUserId      int64  `db:"sender_user_id"`
	PeerType          int32  `db:"peer_type"`
	PeerId            int64  `db:"peer_id"`
	RandomId          int64  `db:"random_id"`
	MessageFilterType int32  `db:"message_filter_type"`
	MessageData       string `db:"message_data"`
	Message           string `db:"message"`
	Mentioned         bool   `db:"mentioned"`
	MediaUnread       bool   `db:"media_unread"`
	Pinned            bool   `db:"pinned"`
	Date2             int64  `db:"date2"`
	Deleted           bool   `db:"deleted"`
}
