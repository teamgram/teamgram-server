// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"github.com/teamgram/proto/mtproto"
)

type idxId struct {
	idx int
	id  int64
}

type kv struct {
	k string
	v interface{}
}

func removeAllNil(participants []*mtproto.ImmutableChatParticipant) []*mtproto.ImmutableChatParticipant {
	for i := 0; i < len(participants); {
		if participants[i] != nil {
			i++
			continue
		}

		if i < len(participants)-1 {
			copy(participants[i:], participants[i+1:])
		}

		participants[len(participants)-1] = nil
		participants = participants[:len(participants)-1]
	}

	return participants
}
