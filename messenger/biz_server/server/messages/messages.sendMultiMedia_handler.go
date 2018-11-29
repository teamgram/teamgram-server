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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"golang.org/x/net/context"
	"time"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/message"
)

func (s *MessagesServiceImpl) makeOutboxMessageBySendMultiMedia(authKeyId int64, fromId int32, peer *base.PeerUtil, request *mtproto.TLMessagesSendMultiMedia) ([]*mtproto.Message, []int64) {
	multi_media := request.GetMultiMedia()
	messages := make([]*mtproto.Message, 0, len(multi_media))
	randomIdList := make([]int64, 0, len(multi_media))
	groupedId := core.GetUUID()
	for _, media := range multi_media {
		message := &mtproto.TLMessage{Data2: &mtproto.Message_Data{
			Out:          true,
			Silent:       request.GetSilent(),
			FromId:       fromId,
			ToId:         peer.ToPeer(),
			ReplyToMsgId: request.GetReplyToMsgId(),
			Media:        s.makeMediaByInputMedia(authKeyId, media.GetData2().GetMedia()),
			// Entities:    media.GetData2()
			// ReplyMarkup: media.GetData2().GetReplyMarkup(),
			Date:         int32(time.Now().Unix()),
			GroupedId:    groupedId,
		}}

		messages = append(messages, message.To_Message())
		randomIdList = append(randomIdList, media.GetData2().GetRandomId())
	}

	return messages, randomIdList
}

func (s *MessagesServiceImpl) makeUpdateNewMessageListUpdates(selfUserId, pts, ptsCount int32, boxList []*message.MessageBox2) *mtproto.TLUpdates {
	var (
		messages = make([]*mtproto.Message, 0, len(boxList))
	)

	pts = pts - ptsCount + 1
	for _, box := range boxList {
		messages = append(messages, box.ToMessage(selfUserId))
	}

	userIdList, _, _ := message.PickAllIDListByMessages(messages)
	userList := s.UserModel.GetUserListByIdList(selfUserId, userIdList)
	updateNewList := make([]*mtproto.Update, 0, len(messages))
	for _, m := range messages {
		// pts += 1
		updateNewMessage := &mtproto.TLUpdateNewMessage{Data2: &mtproto.Update_Data{
			Message_1: m,
			Pts:       pts,
			PtsCount:  1,
		}}
		pts += 1
		updateNewList = append(updateNewList, updateNewMessage.To_Update())
	}

	return &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: updateNewList,
		Users:   userList,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}}
}

// messages.sendMultiMedia#2095512f flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int multi_media:Vector<InputSingleMedia> = Updates;
func (s *MessagesServiceImpl) MessagesSendMultiMedia(ctx context.Context, request *mtproto.TLMessagesSendMultiMedia) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.sendMultiMedia#2095512f - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//md := grpc_util.RpcMetadataFromIncoming(ctx)
	//glog.Infof("messages.sendMedia#c8f16791 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): ???
	// request.NoWebpage
	// request.Background

	// peer
	var (
		peer *base.PeerUtil
		err  error
	)

	if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerEmpty {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("messages.sendMedia#c8f16791 - invalid peer", err)
		return nil, err
	}
	// TODO(@benqi): check user or channels's access_hash

	// peer = helper.FromInputPeer2(md.UserId, request.GetPeer())
	if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerSelf {
		peer = &base.PeerUtil{
			PeerType: base.PEER_USER,
			PeerId:   md.UserId,
		}
	} else {
		peer = base.FromInputPeer(request.GetPeer())
	}

	// 1. draft
	if request.GetClearDraft() {
		s.DoClearDraft(md.UserId, md.AuthId, peer)
	}

	///////////////////////////////////////////////////////////////////////////////////////
	//// 发件箱
	outboxMessages, randomIdList := s.makeOutboxMessageBySendMultiMedia(md.AuthId, md.UserId, peer, request)

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

	glog.Infof("messages.sendMultiMedia#2095512f - reply: %s", logger.JsonDebugData(resultUpdates))
	return resultUpdates, nil
}
