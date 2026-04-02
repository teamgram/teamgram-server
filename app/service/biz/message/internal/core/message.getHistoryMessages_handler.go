// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// MessageGetHistoryMessages
// message.getHistoryMessages user_id:long peer_type:int peer_id:long offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = Vector<MessageBox>;
func (c *MessageCore) MessageGetHistoryMessages(in *message.TLMessageGetHistoryMessages) (*message.VectorMessageBox, error) {
	if in == nil || in.Limit <= 0 {
		return &message.VectorMessageBox{Datas: []tg.MessageBoxClazz{}}, nil
	}

	startID := int32(10)
	if in.OffsetId > 0 {
		startID = in.OffsetId
	}
	if in.MaxId > 0 {
		startID = in.MaxId
	}
	if in.MinId > 0 && startID < in.MinId {
		startID = in.MinId
	}

	limit := in.Limit
	if limit > 3 {
		limit = 3
	}

	boxes := make([]tg.MessageBoxClazz, 0, limit)
	for i := int32(0); i < limit; i++ {
		boxes = append(boxes, makePlaceholderMessageBox(in.UserId, in.PeerType, in.PeerId, startID+i))
	}

	return &message.VectorMessageBox{Datas: boxes}, nil
}
