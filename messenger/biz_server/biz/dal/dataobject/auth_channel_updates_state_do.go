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

type AuthChannelUpdatesStateDO struct {
	Id        int32  `db:"id"`
	AuthKeyId int64  `db:"auth_key_id"`
	UserId    int32  `db:"user_id"`
	ChannelId int32  `db:"channel_id"`
	Pts       int32  `db:"pts"`
	Pts2      int32  `db:"pts2"`
	Date      int32  `db:"date"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	DeletedAt string `db:"deleted_at"`
}
