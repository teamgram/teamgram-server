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
	"github.com/nebula-chat/chatengine/service/document/client"
	"golang.org/x/net/context"
)

// messages.getStickerSet#2619a90e stickerset:InputStickerSet = messages.StickerSet;
func (s *MessagesServiceImpl) MessagesGetStickerSet(ctx context.Context, request *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getStickerSet#2619a90e - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): check inputStickerSetEmpty
	set := s.StickerModel.GetStickerSet(request.GetStickerset())
	packs, idList := s.StickerModel.GetStickerPackList(set.GetData2().GetId())
	var (
		documents []*mtproto.Document
		err       error
	)

	if len(idList) == 0 {
		documents = []*mtproto.Document{}
	} else {
		documents, err = document_client.GetDocumentByIdList(idList)
		if err != nil {
			glog.Error(err)
			documents = []*mtproto.Document{}
		}
	}

	reply := &mtproto.TLMessagesStickerSet{Data2: &mtproto.Messages_StickerSet_Data{
		Set:       set,
		Packs:     packs,
		Documents: documents,
	}}

	glog.Infof("messages.getStickerSet#2619a90e - reply: %s", logger.JsonDebugData(reply))
	return reply.To_Messages_StickerSet(), nil
}
