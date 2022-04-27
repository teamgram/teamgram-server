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

type UserPeerSettingsDO struct {
	Id                    int64 `db:"id"`
	UserId                int64 `db:"user_id"`
	PeerType              int32 `db:"peer_type"`
	PeerId                int64 `db:"peer_id"`
	Hide                  bool  `db:"hide"`
	ReportSpam            bool  `db:"report_spam"`
	AddContact            bool  `db:"add_contact"`
	BlockContact          bool  `db:"block_contact"`
	ShareContact          bool  `db:"share_contact"`
	NeedContactsException bool  `db:"need_contacts_exception"`
	ReportGeo             bool  `db:"report_geo"`
	Autoarchived          bool  `db:"autoarchived"`
	InviteMembers         bool  `db:"invite_members"`
	GeoDistance           int32 `db:"geo_distance"`
}
