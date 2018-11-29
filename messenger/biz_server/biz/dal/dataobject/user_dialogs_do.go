// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package dataobject

type UserDialogsDO struct {
	Id                  int32  `db:"id"`
	UserId              int32  `db:"user_id"`
	PeerType            int8   `db:"peer_type"`
	PeerId              int32  `db:"peer_id"`
	IsPinned            int8   `db:"is_pinned"`
	TopMessage          int32  `db:"top_message"`
	ReadInboxMaxId      int32  `db:"read_inbox_max_id"`
	ReadOutboxMaxId     int32  `db:"read_outbox_max_id"`
	UnreadCount         int32  `db:"unread_count"`
	UnreadMentionsCount int32  `db:"unread_mentions_count"`
	ShowPreviews        int8   `db:"show_previews"`
	Silent              int8   `db:"silent"`
	MuteUntil           int32  `db:"mute_until"`
	Sound               string `db:"sound"`
	Pts                 int32  `db:"pts"`
	DraftId             int32  `db:"draft_id"`
	DraftType           int8   `db:"draft_type"`
	DraftMessageData    string `db:"draft_message_data"`
	Date2               int32  `db:"date2"`
	Version             int64  `db:"version"`
	CreatedAt           string `db:"created_at"`
	UpdatedAt           string `db:"updated_at"`
}
