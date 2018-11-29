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

type ChannelMessagesDO struct {
	Id               int32  `db:"id"`
	ChannelId        int32  `db:"channel_id"`
	ChannelMessageId int32  `db:"channel_message_id"`
	SenderUserId     int32  `db:"sender_user_id"`
	RandomId         int64  `db:"random_id"`
	MessageDataId    int64  `db:"message_data_id"`
	MessageType      int8   `db:"message_type"`
	MessageData      string `db:"message_data"`
	HasMediaUnread   int8   `db:"has_media_unread"`
	EditMessage      string `db:"edit_message"`
	EditDate         int32  `db:"edit_date"`
	Views            int32  `db:"views"`
	Date             int32  `db:"date"`
	Deleted          int8   `db:"deleted"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}
