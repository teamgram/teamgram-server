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
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

// messages.getSavedGifs#83bf3d52 hash:int = messages.SavedGifs;
func (s *MessagesServiceImpl) MessagesGetSavedGifs(ctx context.Context, request *mtproto.TLMessagesGetSavedGifs) (*mtproto.Messages_SavedGifs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getSavedGifs#83bf3d52 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetSavedGifs logic
	stickers := mtproto.TLMessagesSavedGifs{Data2: &mtproto.Messages_SavedGifs_Data{
		Hash: request.Hash,
		Gifs: []*mtproto.Document{},
	}}

	glog.Infof("messages.getSavedGifs#83bf3d52 - reply: %s", logger.JsonDebugData(stickers))
	return stickers.To_Messages_SavedGifs(), nil
}
