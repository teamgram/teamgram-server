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

type DialogsDO struct {
	Id                   int64  `db:"id" json:"id"`
	UserId               int64  `db:"user_id" json:"user_id"`
	PeerType             int32  `db:"peer_type" json:"peer_type"`
	PeerId               int64  `db:"peer_id" json:"peer_id"`
	PeerDialogId         int64  `db:"peer_dialog_id" json:"peer_dialog_id"`
	Pinned               int64  `db:"pinned" json:"pinned"`
	TopMessage           int32  `db:"top_message" json:"top_message"`
	PinnedMsgId          int32  `db:"pinned_msg_id" json:"pinned_msg_id"`
	ReadInboxMaxId       int32  `db:"read_inbox_max_id" json:"read_inbox_max_id"`
	ReadOutboxMaxId      int32  `db:"read_outbox_max_id" json:"read_outbox_max_id"`
	UnreadCount          int32  `db:"unread_count" json:"unread_count"`
	UnreadMentionsCount  int32  `db:"unread_mentions_count" json:"unread_mentions_count"`
	UnreadReactionsCount int32  `db:"unread_reactions_count" json:"unread_reactions_count"`
	UnreadMark           bool   `db:"unread_mark" json:"unread_mark"`
	DraftType            int32  `db:"draft_type" json:"draft_type"`
	DraftMessageData     string `db:"draft_message_data" json:"draft_message_data"`
	FolderId             int32  `db:"folder_id" json:"folder_id"`
	FolderPinned         int64  `db:"folder_pinned" json:"folder_pinned"`
	HasScheduled         bool   `db:"has_scheduled" json:"has_scheduled"`
	TtlPeriod            int32  `db:"ttl_period" json:"ttl_period"`
	ThemeEmoticon        string `db:"theme_emoticon" json:"theme_emoticon"`
	Date2                int64  `db:"date2" json:"date2"`
	Deleted              bool   `db:"deleted" json:"deleted"`
}
