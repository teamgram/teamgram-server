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

type UserNotifySettingsDO struct {
	Id           int64  `db:"id"`
	UserId       int64  `db:"user_id"`
	PeerType     int32  `db:"peer_type"`
	PeerId       int64  `db:"peer_id"`
	ShowPreviews int32  `db:"show_previews"`
	Silent       int32  `db:"silent"`
	MuteUntil    int32  `db:"mute_until"`
	Sound        string `db:"sound"`
	Deleted      bool   `db:"deleted"`
}
