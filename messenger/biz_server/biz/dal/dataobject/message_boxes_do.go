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

type MessageBoxesDO struct {
	Id               int32  `db:"id"`
	UserId           int32  `db:"user_id"`
	UserMessageBoxId int32  `db:"user_message_box_id"`
	DialogId         int64  `db:"dialog_id"`
	DialogMessageId  int32  `db:"dialog_message_id"`
	MessageDataId    int64  `db:"message_data_id"`
	MessageBoxType   int8   `db:"message_box_type"`
	ReplyToMsgId     int32  `db:"reply_to_msg_id"`
	Mentioned        int8   `db:"mentioned"`
	MediaUnread      int8   `db:"media_unread"`
	Date2            int32  `db:"date2"`
	Deleted          int8   `db:"deleted"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}
