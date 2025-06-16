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

package sess

import (
	"time"

	"github.com/teamgram/marmota/pkg/sync2"
)

var msgIdSeq = sync2.NewAtomicInt64(0)

func nextMessageId(isRpc bool) int64 {
	unixNano := time.Now().UnixNano()
	ts := unixNano / 1e9
	ms := (unixNano % 1e9) / 1e6
	sid := msgIdSeq.Add(1) & 0x1ffff
	msgIdSeq.CompareAndSwap(0x1ffff, 0)
	last := 1
	if !isRpc {
		last = 3
	}
	msgId := ts<<32 | int64(ms)<<21 | sid<<3 | int64(last)
	return msgId
}
