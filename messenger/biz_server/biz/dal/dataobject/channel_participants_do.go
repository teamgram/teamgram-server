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

type ChannelParticipantsDO struct {
	Id              int64  `db:"id"`
	ChannelId       int32  `db:"channel_id"`
	UserId          int32  `db:"user_id"`
	IsCreator       int32  `db:"is_creator"`
	ParticipantType int8   `db:"participant_type"`
	InviterUserId   int32  `db:"inviter_user_id"`
	InvitedAt       int32  `db:"invited_at"`
	JoinedAt        int32  `db:"joined_at"`
	PromotedBy      int32  `db:"promoted_by"`
	AdminRights     int32  `db:"admin_rights"`
	PromotedAt      int32  `db:"promoted_at"`
	IsLeft          int8   `db:"is_left"`
	LeftAt          int32  `db:"left_at"`
	IsKicked        int8   `db:"is_kicked"`
	KickedBy        int32  `db:"kicked_by"`
	KickedAt        int32  `db:"kicked_at"`
	BannedRights    int32  `db:"banned_rights"`
	BannedUntilDate int32  `db:"banned_until_date"`
	BannedAt        int32  `db:"banned_at"`
	ReadInboxMaxId  int32  `db:"read_inbox_max_id"`
	ReadOutboxMaxId int32  `db:"read_outbox_max_id"`
	Date            int32  `db:"date"`
	State           int8   `db:"state"`
	CreatedAt       string `db:"created_at"`
	UpdatedAt       string `db:"updated_at"`
}
