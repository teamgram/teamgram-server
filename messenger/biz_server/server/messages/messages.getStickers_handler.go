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

// messages.getStickers#43d4f2c emoticon:string hash:int = messages.Stickers;
func (s *MessagesServiceImpl) MessagesGetStickers(ctx context.Context, request *mtproto.TLMessagesGetStickers) (*mtproto.Messages_Stickers, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("messages.getStickers#43d4f2c - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    stickers := &mtproto.TLMessagesStickers{Data2: &mtproto.Messages_Stickers_Data{
        Hash:     request.GetHash(),
        Stickers: []*mtproto.Document{},
    }}

    glog.Infof("messages.getStickers#43d4f2c - reply: %s", logger.JsonDebugData(stickers))
    return stickers.To_Messages_Stickers(), nil
}
