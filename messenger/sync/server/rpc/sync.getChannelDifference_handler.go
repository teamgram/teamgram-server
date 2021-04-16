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

package rpc

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// sync.getChannelDifference flags:# auth_key_id:long user_id:int force:flags.0?true channel:InputChannel filter:ChannelMessagesFilter pts:int limit:int = updates.ChannelDifference;
func (s *SyncServiceImpl) SyncGetChannelDifference(ctx context.Context, request *mtproto.TLSyncGetChannelDifference) (*mtproto.Updates_ChannelDifference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("sync.getChannelDifference - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))
	var (
		lastPts      = request.GetPts()
		otherUpdates []*mtproto.Update
		messages     []*mtproto.Message
		userList     []*mtproto.User
		chatList     []*mtproto.Chat
	)
	var difference *mtproto.Updates_ChannelDifference
	updateList := s.GetChannelUpdateListByGtPts(request.GetChannel().GetData2().GetChannelId(), lastPts)

	for _, update := range updateList {
		switch update.GetConstructor() {
		case mtproto.TLConstructor_CRC32_updateNewChannelMessage:
			newMessage := update.To_UpdateNewChannelMessage()
			messages = append(messages, newMessage.GetMessage())
			otherUpdates = append(otherUpdates, update)

		case mtproto.TLConstructor_CRC32_updateDeleteChannelMessages:
			readHistoryOutbox := update.To_UpdateReadHistoryOutbox()
			readHistoryOutbox.SetPtsCount(0)
			otherUpdates = append(otherUpdates, readHistoryOutbox.To_Update())
		case mtproto.TLConstructor_CRC32_updateEditChannelMessage:
			readHistoryInbox := update.To_UpdateReadHistoryInbox()
			readHistoryInbox.SetPtsCount(0)
			otherUpdates = append(otherUpdates, readHistoryInbox.To_Update())
		case mtproto.TLConstructor_CRC32_updateChannelWebPage:
		default:
			continue
		}
		if update.Data2.GetPts() > lastPts {
			lastPts = update.Data2.GetPts()
		}
	}
	if lastPts <= request.GetPts() {
		lastPts = 0
	}

	// state := &mtproto.TLUpdatesState{Data2: &mtproto.Updates_State_Data{
	// 	Pts:         lastPts,
	// 	Date:        int32(time.Now().Unix()),
	// 	UnreadCount: 0,
	// 	// Seq:         int32(model.GetSequenceModel().CurrentSeqId(base2.Int32ToString(md.UserId))),
	// 	Seq: 0,
	// }}
	difference2 := &mtproto.TLUpdatesChannelDifference{Data2: &mtproto.Updates_ChannelDifference_Data{
		Pts:          lastPts,
		NewMessages:  messages,
		OtherUpdates: otherUpdates,
		Users:        userList,
		Chats:        chatList,
	}}
	difference = difference2.To_Updates_ChannelDifference()
	return difference, nil
}
