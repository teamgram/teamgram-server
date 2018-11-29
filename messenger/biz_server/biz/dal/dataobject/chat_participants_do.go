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

type ChatParticipantsDO struct {
	Id              int32  `db:"id"`
	ChatId          int32  `db:"chat_id"`
	UserId          int32  `db:"user_id"`
	ParticipantType int8   `db:"participant_type"`
	InviterUserId   int32  `db:"inviter_user_id"`
	InvitedAt       int32  `db:"invited_at"`
	KickedAt        int32  `db:"kicked_at"`
	LeftAt          int32  `db:"left_at"`
	State           int8   `db:"state"`
	CreatedAt       string `db:"created_at"`
	UpdatedAt       string `db:"updated_at"`
}
