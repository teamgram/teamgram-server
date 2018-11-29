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

type ChannelMessageBoxesDO struct {
	Id                  int32  `db:"id"`
	SenderUserId        int32  `db:"sender_user_id"`
	ChannelId           int32  `db:"channel_id"`
	DialogId            int64  `db:"dialog_id"`
	DialogMessageId     int32  `db:"dialog_message_id"`
	ChannelMessageBoxId int32  `db:"channel_message_box_id"`
	MessageId           int64  `db:"message_id"`
	Date                int32  `db:"date"`
	Deleted             int8   `db:"deleted"`
	CreatedAt           string `db:"created_at"`
	UpdatedAt           string `db:"updated_at"`
}
