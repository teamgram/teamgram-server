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

type DraftsDO struct {
	Id               int32  `db:"id" json:"id"`
	UserId           int32  `db:"user_id" json:"user_id"`
	PeerDialogId     int64  `db:"peer_dialog_id" json:"peer_dialog_id"`
	DraftType        int32  `db:"draft_type" json:"draft_type"`
	DraftMessageData string `db:"draft_message_data" json:"draft_message_data"`
	Date2            int64  `db:"date2" json:"date2"`
}
