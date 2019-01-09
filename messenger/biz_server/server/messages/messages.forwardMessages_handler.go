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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/message"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
	"time"
)

func (s *MessagesServiceImpl) makeForwardMessagesData(selfId int32, idList []int32, peer *base.PeerUtil, ridList []int64) ([]*mtproto.Message, []int64) {
	findRandomIdById := func(id int32) int64 {
		for i := 0; i < len(idList); i++ {
			if id == idList[i] {
				return ridList[i]
			}
		}
		return 0
	}

	// TODO(@benqi): process channel

	// 通过idList找到message
	messages := s.MessageModel.GetUserMessagesByMessageIdList(selfId, idList)
	randomIdList := make([]int64, 0, len(messages))
	for _, m := range messages {
		// TODO(@benqi): rid is 0
		randomIdList = append(randomIdList, findRandomIdById(m.GetData2().GetId()))

		fwdFrom := &mtproto.TLMessageFwdHeader{Data2: &mtproto.MessageFwdHeader_Data{
			FromId: m.GetData2().GetFromId(),
			Date:   m.GetData2().GetDate(),
		}}

		// make message
		m.Data2.ToId = peer.ToPeer()
		m.Data2.FromId = selfId
		m.Data2.FwdFrom = fwdFrom.To_MessageFwdHeader()
		m.Data2.Date = int32(time.Now().Unix())
	}

	// reverse messages
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, randomIdList
}

// messages.forwardMessages#708e0195 flags:# silent:flags.5?true background:flags.6?true with_my_score:flags.8?true grouped:flags.9?true from_peer:InputPeer id:Vector<int> random_id:Vector<long> to_peer:InputPeer = Updates;
func (s *MessagesServiceImpl) MessagesForwardMessages(ctx context.Context, request *mtproto.TLMessagesForwardMessages) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.forwardMessages#708e0195 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//// peer
	var (
		// fromPeer = helper.FromInputPeer2(md.UserId, request.GetFromPeer())
		peer = base.FromInputPeer2(md.UserId, request.GetToPeer())
		// messageOutboxList message2.MessageBoxList
	)

	outboxMessages, randomIdList := s.makeForwardMessagesData(md.UserId, request.GetId(), peer, request.GetRandomId())

	resultCB := func(pts, ptsCount int32, outBoxList []*message.MessageBox2) (*mtproto.Updates, error) {
		resultUpdates := s.makeUpdateNewMessageListUpdates(md.UserId, pts, ptsCount, outBoxList)

		updateList := make([]*mtproto.Update, 0)
		for i := 0; i < len(outBoxList); i++ {
			updateMessageID := &mtproto.TLUpdateMessageID{Data2: &mtproto.Update_Data{
				Id_4:     outBoxList[i].MessageId,
				RandomId: outBoxList[i].RandomId,
			}}
			updateList = append(updateList, updateMessageID.To_Update())
		}
		updateList = append(updateList, resultUpdates.GetUpdates()...)
		resultUpdates.SetUpdates(updateList)

		return resultUpdates.To_Updates(), nil
	}

	syncNotMeCB := func(pts, ptsCount int32, outBoxList []*message.MessageBox2) (int64, *mtproto.Updates, error) {
		syncUpdates := s.makeUpdateNewMessageListUpdates(md.UserId, pts, ptsCount, outBoxList)
		return md.AuthId, syncUpdates.To_Updates(), nil
	}

	pushCB := func(userId, pts, ptsCount int32, inBoxList []*message.MessageBox2) (*mtproto.Updates, error) {
		pushUpdates := s.makeUpdateNewMessageListUpdates(userId, pts, ptsCount, inBoxList)
		return pushUpdates.To_Updates(), nil
	}

	resultUpdates, err := s.MessageModel.SendMultiMessage(
		md.UserId,
		peer,
		randomIdList,
		outboxMessages,
		resultCB,
		syncNotMeCB,
		pushCB)

	glog.Infof("messages.forwardMessages#708e0195 - reply: %s", logger.JsonDebugData(resultUpdates))
	return resultUpdates, err
}
