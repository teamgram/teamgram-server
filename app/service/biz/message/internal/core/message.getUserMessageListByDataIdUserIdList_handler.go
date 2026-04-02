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

// MessageGetUserMessageListByDataIdUserIdList
// message.getUserMessageListByDataIdUserIdList id:long user_id_list:Vector<long> = Vector<MessageBox>;
func (c *MessageCore) MessageGetUserMessageListByDataIdUserIdList(in *message.TLMessageGetUserMessageListByDataIdUserIdList) (*message.VectorMessageBox, error) {
	boxes := make([]tg.MessageBoxClazz, 0, len(in.UserIdList))
	for _, userID := range in.UserIdList {
		boxes = append(boxes, makePlaceholderMessageBox(userID, tg.PEER_USER, userID, int32(in.Id)))
	}

	return &message.VectorMessageBox{Datas: boxes}, nil
}
