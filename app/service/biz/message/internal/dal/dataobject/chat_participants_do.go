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

type ChatParticipantsDO struct {
	Id                             int64  `db:"id" json:"id"`
	ChatId                         int64  `db:"chat_id" json:"chat_id"`
	UserId                         int64  `db:"user_id" json:"user_id"`
	ParticipantType                int32  `db:"participant_type" json:"participant_type"`
	Link                           string `db:"link" json:"link"`
	Usage2                         int32  `db:"usage2" json:"usage2"`
	AdminRights                    int32  `db:"admin_rights" json:"admin_rights"`
	InviterUserId                  int64  `db:"inviter_user_id" json:"inviter_user_id"`
	InvitedAt                      int64  `db:"invited_at" json:"invited_at"`
	KickedAt                       int64  `db:"kicked_at" json:"kicked_at"`
	LeftAt                         int64  `db:"left_at" json:"left_at"`
	GroupcallDefaultJoinAsPeerType int32  `db:"groupcall_default_join_as_peer_type" json:"groupcall_default_join_as_peer_type"`
	GroupcallDefaultJoinAsPeerId   int64  `db:"groupcall_default_join_as_peer_id" json:"groupcall_default_join_as_peer_id"`
	IsBot                          bool   `db:"is_bot" json:"is_bot"`
	State                          int32  `db:"state" json:"state"`
	Date2                          int64  `db:"date2" json:"date2"`
}
