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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"golang.org/x/net/context"
	"time"
)

func makeDraftMessageBySaveDraft(request *mtproto.TLMessagesSaveDraft) *mtproto.TLDraftMessage {
	return &mtproto.TLDraftMessage{Data2: &mtproto.DraftMessage_Data{
		NoWebpage:    request.GetNoWebpage(),
		ReplyToMsgId: request.GetReplyToMsgId(),
		Message:      request.GetMessage(),
		Entities:     request.GetEntities(),
		Date:         int32(time.Now().Unix()),
	}}
}

// messages.saveDraft#bc39e14b flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> = Bool;
func (s *MessagesServiceImpl) MessagesSaveDraft(ctx context.Context, request *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.saveDraft#bc39e14b - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		peer *base.PeerUtil
	)

	if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerSelf {
		peer = &base.PeerUtil{PeerType: base.PEER_USER, PeerId: md.UserId}
	} else {
		peer = base.FromInputPeer(request.GetPeer())
	}

	draft := makeDraftMessageBySaveDraft(request)

	// TODO(@benqi): 会话未存在如何处理？
	s.DialogModel.SaveDraftMessage(md.UserId, peer.PeerType, peer.PeerId, draft.To_DraftMessage())


	// TODO(@benqi): sync other client

	reply := mtproto.ToBool(true)

	glog.Infof("messages.saveDraft#bc39e14b - reply: {%v}", reply)
	return reply, nil
}
