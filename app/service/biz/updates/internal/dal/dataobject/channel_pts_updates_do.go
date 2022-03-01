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

type ChannelPtsUpdatesDO struct {
	Id           int64  `db:"id"`
	ChannelId    int64  `db:"channel_id"`
	Pts          int32  `db:"pts"`
	PtsCount     int32  `db:"pts_count"`
	UpdateType   int32  `db:"update_type"`
	NewMessageId int32  `db:"new_message_id"`
	UpdateData   string `db:"update_data"`
	Date2        int64  `db:"date2"`
}
