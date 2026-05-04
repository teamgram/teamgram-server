// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tg

import (
	"time"
)

func MakeEmptyUpdates() *Updates {
	return MakeTLUpdates(&TLUpdates{
		Updates: []UpdateClazz{},
		Users:   []UserClazz{},
		Chats:   []ChatClazz{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).ToUpdates()
}

func MakeUpdatesByUpdates(updates ...UpdateClazz) *Updates {
	return MakeTLUpdates(&TLUpdates{
		Updates: updates,
		Users:   []UserClazz{},
		Chats:   []ChatClazz{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).ToUpdates()
}

func MakeUpdatesByUpdatesUsers(users []UserClazz, updates ...UpdateClazz) *Updates {
	if users == nil {
		users = []UserClazz{}
	}
	return MakeTLUpdates(&TLUpdates{
		Updates: updates,
		Users:   users,
		Chats:   []ChatClazz{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).ToUpdates()
}

func MakeUpdatesByUpdatesChats(chats []ChatClazz, updates ...UpdateClazz) *Updates {
	if chats == nil {
		chats = []ChatClazz{}
	}
	return MakeTLUpdates(&TLUpdates{
		Updates: updates,
		Users:   []UserClazz{},
		Chats:   chats,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).ToUpdates()
}

func MakeUpdatesByUpdatesUsersChats(users []UserClazz, chats []ChatClazz, updates ...UpdateClazz) *Updates {
	if users == nil {
		users = []UserClazz{}
	}
	if chats == nil {
		chats = []ChatClazz{}
	}

	return MakeTLUpdates(&TLUpdates{
		Updates: updates,
		Users:   users,
		Chats:   chats,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).ToUpdates()
}
