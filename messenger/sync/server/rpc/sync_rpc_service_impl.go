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
	"fmt"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/sync/biz/core/update"
	"github.com/nebula-chat/chatengine/mtproto"
	zrpc "github.com/nebula-chat/chatengine/mtproto/rpc"
	status_client "github.com/nebula-chat/chatengine/service/status/client"
	status "github.com/nebula-chat/chatengine/service/status/proto"
)

/*
 android client source code:
    private int getUpdateType(TLRPC.Update update) {
        if (update instanceof TLRPC.TL_updateNewMessage || update instanceof TLRPC.TL_updateReadMessagesContents || update instanceof TLRPC.TL_updateReadHistoryInbox ||
                update instanceof TLRPC.TL_updateReadHistoryOutbox || update instanceof TLRPC.TL_updateDeleteMessages || update instanceof TLRPC.TL_updateWebPage ||
                update instanceof TLRPC.TL_updateEditMessage) {
            return 0;
        } else if (update instanceof TLRPC.TL_updateNewEncryptedMessage) {
            return 1;
        } else if (update instanceof TLRPC.TL_updateNewChannelMessage || update instanceof TLRPC.TL_updateDeleteChannelMessages || update instanceof TLRPC.TL_updateEditChannelMessage ||
                update instanceof TLRPC.TL_updateChannelWebPage) {
            return 2;
        } else {
            return 3;
        }
    }
*/

type SyncType int

const (
	syncTypeUser      SyncType = 1 // 该用户所有设备
	syncTypeUserNotMe SyncType = 2 // 该用户除了某个设备
	syncTypeUserMe    SyncType = 3 // 该用户指定某个设备
	syncTypeRpcResult SyncType = 4 // 通过push通道返回rpc
)

// messages.AffectedHistory
// messages.AffectedMessages

type PushData struct {
	cntl       *zrpc.ZRpcController
	rpcMessage proto.Message
}

func MakePushUpdatesData(authKeyId int64, cntl *zrpc.ZRpcController, pts, ptsCount int32, pushRawData []byte) *PushData {
	rpcMessage := &mtproto.TLPushPushUpdatesData{
		AuthKeyId: authKeyId,
		Pts:       pts,
		PtsCount:  ptsCount,
	}

	cntl.SetServiceName("session")
	cntl.SetMethodName(proto.MessageName(rpcMessage))
	cntl.SetAttachment(pushRawData)

	return &PushData{
		cntl:       cntl,
		rpcMessage: rpcMessage,
	}
}

func MakePushRpcResultData(authKeyId, clientMsgId int64, cntl *zrpc.ZRpcController, pushRawData []byte) *PushData {
	rpcMessage := &mtproto.TLPushPushRpcResultData{
		AuthKeyId:      authKeyId,
		ClientReqMsgId: clientMsgId,
	}

	cntl.SetServiceName("session")
	cntl.SetMethodName(proto.MessageName(rpcMessage))
	cntl.SetAttachment(pushRawData)

	return &PushData{
		cntl:       cntl,
		rpcMessage: rpcMessage,
	}
}

type PushDataCallback interface {
	SendToSessionServer(serverId int, cntl *zrpc.ZRpcController, m proto.Message)
}

type SyncServiceImpl struct {
	mu        sync.RWMutex
	pushCB    PushDataCallback
	status    status_client.StatusClient
	closeChan chan int
	pushChan  chan struct {
		int
		*PushData
	}
	*update.UpdateModel
}

func NewSyncService(pushCB PushDataCallback, status status_client.StatusClient, updateModel *update.UpdateModel) *SyncServiceImpl {
	s := &SyncServiceImpl{
		pushCB:    pushCB,
		status:    status,
		closeChan: make(chan int),
		pushChan: make(chan struct {
			int
			*PushData
		}, 1024),
		UpdateModel: updateModel,
	}

	go s.pushUpdatesLoop()
	return s
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func (s *SyncServiceImpl) pushUpdatesLoop() {
	defer func() {
		close(s.pushChan)
	}()

	for {
		select {
		case pair, ok := <-s.pushChan:
			if ok {
				s.pushCB.SendToSessionServer(pair.int, pair.cntl, pair.rpcMessage)
			}
		case <-s.closeChan:
			return
		}
	}
}

func (s *SyncServiceImpl) Destroy() {
	s.closeChan <- 1
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func updateShortMessageToMessage(userId int32, shortMessage *mtproto.TLUpdateShortMessage) *mtproto.Message {
	var (
		fromId, peerId int32
	)
	if shortMessage.GetOut() {
		fromId = userId
		peerId = shortMessage.GetUserId()
	} else {
		fromId = shortMessage.GetUserId()
		peerId = userId
	}

	message := &mtproto.TLMessage{Data2: &mtproto.Message_Data{
		Out:          shortMessage.GetOut(),
		Mentioned:    shortMessage.GetMentioned(),
		MediaUnread:  shortMessage.GetMediaUnread(),
		Silent:       shortMessage.GetSilent(),
		Id:           shortMessage.GetId(),
		FromId:       fromId,
		ToId:         &mtproto.Peer{Constructor: mtproto.TLConstructor_CRC32_peerUser, Data2: &mtproto.Peer_Data{UserId: peerId}},
		Message:      shortMessage.GetMessage(),
		Date:         shortMessage.GetDate(),
		FwdFrom:      shortMessage.GetFwdFrom(),
		ViaBotId:     shortMessage.GetViaBotId(),
		ReplyToMsgId: shortMessage.GetReplyToMsgId(),
		Entities:     shortMessage.GetEntities(),
	}}
	return message.To_Message()
}

func updateShortChatMessageToMessage(shortMessage *mtproto.TLUpdateShortChatMessage) *mtproto.Message {
	message := &mtproto.TLMessage{Data2: &mtproto.Message_Data{
		Out:          shortMessage.GetOut(),
		Mentioned:    shortMessage.GetMentioned(),
		MediaUnread:  shortMessage.GetMediaUnread(),
		Silent:       shortMessage.GetSilent(),
		Id:           shortMessage.GetId(),
		FromId:       shortMessage.GetFromId(),
		ToId:         &mtproto.Peer{Constructor: mtproto.TLConstructor_CRC32_peerChat, Data2: &mtproto.Peer_Data{ChatId: shortMessage.GetChatId()}},
		Message:      shortMessage.GetMessage(),
		Date:         shortMessage.GetDate(),
		FwdFrom:      shortMessage.GetFwdFrom(),
		ViaBotId:     shortMessage.GetViaBotId(),
		ReplyToMsgId: shortMessage.GetReplyToMsgId(),
		Entities:     shortMessage.GetEntities(),
	}}
	return message.To_Message()
}

func updateShortToUpdateNewMessage(userId int32, shortMessage *mtproto.TLUpdateShortMessage) *mtproto.Update {
	updateNew := &mtproto.TLUpdateNewMessage{Data2: &mtproto.Update_Data{
		Message_1: updateShortMessageToMessage(userId, shortMessage),
		Pts:       shortMessage.GetPts(),
		PtsCount:  shortMessage.GetPtsCount(),
	}}
	return updateNew.To_Update()
}

func updateShortChatToUpdateNewMessage(userId int32, shortMessage *mtproto.TLUpdateShortChatMessage) *mtproto.Update {
	updateNew := &mtproto.TLUpdateNewMessage{Data2: &mtproto.Update_Data{
		Message_1: updateShortChatMessageToMessage(shortMessage),
		Pts:       shortMessage.GetPts(),
		PtsCount:  shortMessage.GetPtsCount(),
	}}
	return updateNew.To_Update()
}

func (s *SyncServiceImpl) processUpdatesRequest(userId int32, ups *mtproto.Updates) (int32, int32, error) {
	var pts, ptsCount int32

	switch ups.GetConstructor() {
	case mtproto.TLConstructor_CRC32_updateAccountResetAuthorization:
	case mtproto.TLConstructor_CRC32_updateShortMessage:
		shortMessage := ups.To_UpdateShortMessage()
		s.UpdateModel.AddToPtsQueue(userId, shortMessage.GetPts(), shortMessage.GetPtsCount(), updateShortToUpdateNewMessage(userId, shortMessage))
		pts = shortMessage.GetPts()
		ptsCount = shortMessage.GetPtsCount()
	case mtproto.TLConstructor_CRC32_updateShortChatMessage:
		shortMessage := ups.To_UpdateShortChatMessage()
		s.UpdateModel.AddToPtsQueue(userId, shortMessage.GetPts(), shortMessage.GetPtsCount(), updateShortChatToUpdateNewMessage(userId, shortMessage))
		pts = shortMessage.GetPts()
		ptsCount = shortMessage.GetPtsCount()
	case mtproto.TLConstructor_CRC32_updateShort:
	case mtproto.TLConstructor_CRC32_updates:
		updates2 := ups.To_Updates()
		// totalPtsCount := int32(0)
		for _, update := range updates2.GetUpdates() {
			switch update.GetConstructor() {
			case mtproto.TLConstructor_CRC32_updateNewMessage,
				mtproto.TLConstructor_CRC32_updateReadHistoryOutbox,
				mtproto.TLConstructor_CRC32_updateReadHistoryInbox,
				mtproto.TLConstructor_CRC32_updateWebPage,
				mtproto.TLConstructor_CRC32_updateReadMessagesContents,
				mtproto.TLConstructor_CRC32_updateEditMessage,
				mtproto.TLConstructor_CRC32_updateNewReact:
				s.UpdateModel.AddToPtsQueue(userId, update.Data2.Pts, update.Data2.PtsCount, update)
				if update.Data2.Pts > pts {
					pts = update.Data2.Pts
				}
				ptsCount += update.Data2.Pts
			case mtproto.TLConstructor_CRC32_updateDeleteMessages:
				// deleteMessages := update.To_UpdateDeleteMessages().GetMessages()
				//// TODO(@benqi): NextPtsCountId
				//for i := 0; i < len(deleteMessages); i++ {
				//	pts = int32(s.UpdateModel.NextPtsId(pushUserId))
				//}
				//
				//ptsCount = int32(len(deleteMessages))
				//totalPtsCount += ptsCount
				// @benqi: 以上都有Pts和PtsCount
				// update.Data2.Pts = pts
				// update.Data2.PtsCount = ptsCount
				s.UpdateModel.AddToPtsQueue(userId, update.Data2.Pts, update.Data2.PtsCount, update)
				if updates2.Data2.Pts > pts {
					pts = updates2.Data2.Pts
				}
				if update.Data2.Pts > pts {
					pts = update.Data2.Pts
				}
				ptsCount += update.Data2.Pts
			case mtproto.TLConstructor_CRC32_updateNewChannelMessage:
				s.UpdateModel.AddToPtsQueue(userId, update.Data2.Pts, update.Data2.PtsCount, update)
				if update.Data2.Pts > pts {
					pts = update.Data2.Pts
				}
				ptsCount += update.Data2.Pts
			}
		}

		// 有可能有多个
		// ptsCount = totalPtsCount
		// updates2.SetDate(date)
		// updates2.SetSeq(seq)
	default:
		err := fmt.Errorf("invalid updates data: {%d}", ups.GetConstructor())
		// glog.Error(err)
		return 0, 0, err
	}

	//state := &mtproto.ClientUpdatesState{
	//	Pts:      pts,
	//	PtsCount: ptsCount,
	//	Date:     date,
	//}

	return pts, ptsCount, nil
}

func (s *SyncServiceImpl) processChannelUpdatesRequest(userId, channelId int32, ups *mtproto.Updates) (int32, int32, error) {
	var pts, ptsCount int32
	switch ups.GetConstructor() {
	case mtproto.TLConstructor_CRC32_updates:
		updates2 := ups.To_Updates()
		for _, update := range updates2.GetUpdates() {
			switch update.GetConstructor() {
			case mtproto.TLConstructor_CRC32_updateNewChannelMessage:
				s.UpdateModel.AddToChannelPtsQueue(channelId, update.Data2.Pts, update.Data2.PtsCount, update)
				s.UpdateModel.AddToPtsQueue(userId, update.Data2.Pts, update.Data2.PtsCount, update)
				if updates2.Data2.Pts > pts {
					pts = updates2.Data2.Pts
				}
				ptsCount += updates2.Data2.PtsCount
			case mtproto.TLConstructor_CRC32_updateDeleteChannelMessages:
				s.UpdateModel.AddToChannelPtsQueue(channelId, update.Data2.Pts, update.Data2.PtsCount, update)
				if updates2.Data2.Pts > pts {
					pts = updates2.Data2.Pts
				}
				ptsCount += updates2.Data2.PtsCount
			case mtproto.TLConstructor_CRC32_updateEditChannelMessage:
				s.UpdateModel.AddToChannelPtsQueue(channelId, update.Data2.Pts, update.Data2.PtsCount, update)
				if updates2.Data2.Pts > pts {
					pts = updates2.Data2.Pts
				}
				ptsCount += updates2.Data2.PtsCount
			case mtproto.TLConstructor_CRC32_updateChannelWebPage:
				s.UpdateModel.AddToChannelPtsQueue(channelId, update.Data2.Pts, update.Data2.PtsCount, update)
				if updates2.Data2.Pts > pts {
					pts = updates2.Data2.Pts
				}
				ptsCount += updates2.Data2.PtsCount
			}
		}
	default:
		err := fmt.Errorf("invalid updates data: {%d}", ups.GetConstructor())
		// glog.Error(err)
		return 0, 0, err
	}
	return pts, ptsCount, nil
}

func (s *SyncServiceImpl) pushUpdatesToSession(syncType SyncType, userId int32, authKeyId, clientMsgId int64, cntl *zrpc.ZRpcController, pushData []byte, hasServerId, pts, ptsCount int32) {
	if (syncType == syncTypeUserMe || syncType == syncTypeRpcResult) && hasServerId > 0 {
		glog.Infof("pushUpdatesToSession - pushData: {server_id: %d, auth_key_id: %d}", hasServerId, authKeyId)
		if syncType == syncTypeUserMe {
			s.pushChan <- struct {
				int
				*PushData
			}{int(hasServerId), MakePushUpdatesData(authKeyId, cntl, pts, ptsCount, pushData)}
		} else {
			s.pushChan <- struct {
				int
				*PushData
			}{int(hasServerId), MakePushRpcResultData(authKeyId, clientMsgId, cntl, pushData)}
		}
	} else {
		statusList, _ := s.status.GetUserOnlineSessions(userId)

		sendedAuthKeyIdList := make(map[int64]bool)

		if len(statusList.Sessions) > 0 {
			glog.Info("statusList - ", statusList)

			ss := make(map[int32][]*status.SessionEntry)
			for _, status2 := range statusList.Sessions {
				if _, ok := ss[status2.ServerId]; !ok {
					ss[status2.ServerId] = []*status.SessionEntry{}
				}
				ss[status2.ServerId] = append(ss[status2.ServerId], status2)
			}

			glog.Info(ss)
			for k, ss3 := range ss {
				glog.Info(ss3)
				for _, ss4 := range ss3 {
					sendedAuthKeyIdList[ss4.AuthKeyId] = true
					if syncType == syncTypeUserNotMe && authKeyId == ss4.AuthKeyId {
						continue
					}
					glog.Infof("pushUpdatesToSession - pushData: {server_id: %d, auth_key_id: %d}", k, ss4.AuthKeyId)
					s.pushChan <- struct {
						int
						*PushData
					}{int(k), MakePushUpdatesData(ss4.AuthKeyId, cntl, pts, ptsCount, pushData)}
				}
			}
		}
	}
}
