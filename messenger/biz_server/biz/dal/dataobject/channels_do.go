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

type ChannelsDO struct {
	Id               int32  `db:"id"`
	CreatorUserId    int32  `db:"creator_user_id"`
	AccessHash       int64  `db:"access_hash"`
	RandomId         int64  `db:"random_id"`
	TopMessage       int32  `db:"top_message"`
	ParticipantCount int32  `db:"participant_count"`
	Title            string `db:"title"`
	About            string `db:"about"`
	PhotoId          int64  `db:"photo_id"`
	Public           int8   `db:"public"`
	Link             string `db:"link"`
	Broadcast        int8   `db:"broadcast"`
	Verified         int8   `db:"verified"`
	Megagroup        int8   `db:"megagroup"`
	Democracy        int8   `db:"democracy"`
	Signatures       int8   `db:"signatures"`
	AdminsEnabled    int8   `db:"admins_enabled"`
	Deactivated      int8   `db:"deactivated"`
	Version          int32  `db:"version"`
	Date             int32  `db:"date"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}
