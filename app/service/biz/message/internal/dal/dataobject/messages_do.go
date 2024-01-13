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

type MessagesDO struct {
	Id                int64  `db:"id" json:"id"`
	UserId            int64  `db:"user_id" json:"user_id"`
	UserMessageBoxId  int32  `db:"user_message_box_id" json:"user_message_box_id"`
	DialogId1         int64  `db:"dialog_id1" json:"dialog_id1"`
	DialogId2         int64  `db:"dialog_id2" json:"dialog_id2"`
	DialogMessageId   int64  `db:"dialog_message_id" json:"dialog_message_id"`
	SenderUserId      int64  `db:"sender_user_id" json:"sender_user_id"`
	PeerType          int32  `db:"peer_type" json:"peer_type"`
	PeerId            int64  `db:"peer_id" json:"peer_id"`
	RandomId          int64  `db:"random_id" json:"random_id"`
	MessageFilterType int32  `db:"message_filter_type" json:"message_filter_type"`
	MessageData       string `db:"message_data" json:"message_data"`
	Message           string `db:"message" json:"message"`
	Mentioned         bool   `db:"mentioned" json:"mentioned"`
	MediaUnread       bool   `db:"media_unread" json:"media_unread"`
	Pinned            bool   `db:"pinned" json:"pinned"`
	HasReaction       bool   `db:"has_reaction" json:"has_reaction"`
	Reaction          string `db:"reaction" json:"reaction"`
	ReactionDate      int64  `db:"reaction_date" json:"reaction_date"`
	ReactionUnread    bool   `db:"reaction_unread" json:"reaction_unread"`
	Date2             int64  `db:"date2" json:"date2"`
	TtlPeriod         int32  `db:"ttl_period" json:"ttl_period"`
	SavedPeerType     int32  `db:"saved_peer_type" json:"saved_peer_type"`
	SavedPeerId       int64  `db:"saved_peer_id" json:"saved_peer_id"`
	Deleted           bool   `db:"deleted" json:"deleted"`
}
