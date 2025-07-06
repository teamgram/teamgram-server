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

type UserPeerSettingsDO struct {
	Id                    int64 `db:"id" json:"id"`
	UserId                int64 `db:"user_id" json:"user_id"`
	PeerType              int32 `db:"peer_type" json:"peer_type"`
	PeerId                int64 `db:"peer_id" json:"peer_id"`
	Hide                  bool  `db:"hide" json:"hide"`
	ReportSpam            bool  `db:"report_spam" json:"report_spam"`
	AddContact            bool  `db:"add_contact" json:"add_contact"`
	BlockContact          bool  `db:"block_contact" json:"block_contact"`
	ShareContact          bool  `db:"share_contact" json:"share_contact"`
	NeedContactsException bool  `db:"need_contacts_exception" json:"need_contacts_exception"`
	ReportGeo             bool  `db:"report_geo" json:"report_geo"`
	Autoarchived          bool  `db:"autoarchived" json:"autoarchived"`
	InviteMembers         bool  `db:"invite_members" json:"invite_members"`
	GeoDistance           int32 `db:"geo_distance" json:"geo_distance"`
}
