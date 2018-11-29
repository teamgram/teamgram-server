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

package messages

import (
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto"
)

// messages.getHistory#afa92846 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (s *MessagesServiceImpl) MessagesGetHistoryLayer51(ctx context.Context, request *mtproto.TLMessagesGetHistoryLayer51) (*mtproto.Messages_Messages, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("messages.getHistory#afa92846 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    request2 := &mtproto.TLMessagesGetHistory{
        Peer:       request.GetPeer(),
        OffsetId:   request.GetOffsetId(),
        OffsetDate: request.GetOffsetDate(),
        AddOffset:  request.GetAddOffset(),
        Limit:      request.GetLimit(),
        MaxId:      request.GetMaxId(),
        MinId:      request.GetMinId(),
        Hash:       0,
    }
    messagesMessages := s.getHistoryMessages(md, request2)

    glog.Infof("messages.getHistory#dcbb8260 - reply: %s", logger.JsonDebugData(messagesMessages))
    return messagesMessages, nil
}
