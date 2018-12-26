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

package update

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/sync/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
	"time"
)

/*
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
const (
	PTS_UPDATE_TYPE_UNKNOWN = 0

	// pts
	PTS_UPDATE_NEW_MESSAGE          = 1
	PTS_UPDATE_DELETE_MESSAGES      = 2
	PTS_UPDATE_READ_HISTORY_OUTBOX  = 3
	PTS_UPDATE_READ_HISTORY_INBOX   = 4
	PTS_UPDATE_WEBPAGE              = 5
	PTS_UPDATE_READ_MESSAGE_CONENTS = 6
	PTS_UPDATE_EDIT_MESSAGE         = 7

	// qts
	PTS_UPDATE_NEW_ENCRYPTED_MESSAGE = 8

	// channel pts
	PTS_UPDATE_NEW_CHANNEL_MESSAGE     = 9
	PTS_UPDATE_DELETE_CHANNEL_MESSAGES = 10
	PTS_UPDATE_EDIT_CHANNEL_MESSAGE    = 11
	PTS_UPDATE_EDIT_CHANNEL_WEBPAGE    = 12
)

func (m *UpdateModel) GetUpdatesState(authKeyId int64, userId int32) *mtproto.TLUpdatesState {
	state := mtproto.NewTLUpdatesState()

	// TODO(@benqi): first sign in, state data???
	ptsDO := m.dao.UserPtsUpdatesDAO.SelectLastPts(userId)
	if ptsDO != nil {
		state.SetPts(ptsDO.Pts)
	} else {
		state.SetPts(1)
	}

	qtsDO := m.dao.UserQtsUpdatesDAO.SelectLastQts(userId)
	if qtsDO != nil {
		state.SetQts(qtsDO.Qts)
	} else {
		state.SetQts(0)
	}

	// state.SetSeq(int32(GetSequenceModel().CurrentSeqId(helper.Int64ToString(authKeyId))))
	state.SetSeq(int32(m.CurrentSeqId(userId)))

	//state.SetSeq(0)
	//
	//seqDO := dao.GetAuthSeqUpdatesDAO(dao.DB_SLAVE).SelectLastSeq(authKeyId, userId)
	//if seqDO != nil {
	//	state.SetSeq(seqDO.Seq)
	//} else {
	//	state.SetSeq(0)
	//}

	state.SetDate(int32(time.Now().Unix()))
	// TODO(@benqi): Calc unread
	state.SetUnreadCount(0)

	return state
}

//func AddPtsToUpdatesQueue(userId, pts, peerType, peerId, updateType, messageBoxId, maxMessageBoxId int32, ) int32 {
//	do := &dataobject.UserPtsUpdatesDO{
//		UserId:          userId,
//		PeerType:		 int8(peerType),
//		PeerId:			 peerId,
//		Pts:             pts,
//		UpdateType:      updateType,
//		MessageBoxId:    messageBoxId,
//		MaxMessageBoxId: maxMessageBoxId,
//		Date2:           int32(time.Now().Unix()),
//	}
//
//	return int32(dao.GetUserPtsUpdatesDAO(dao.DB_MASTER).Insert(do))
//}

func (m *UpdateModel) AddQtsToUpdatesQueue(userId, qts, updateType int32, updateData []byte) int32 {
	do := &dataobject.UserQtsUpdatesDO{
		UserId:     userId,
		UpdateType: updateType,
		UpdateData: updateData,
		Date2:      int32(time.Now().Unix()),
		Qts:        qts,
	}

	return int32(m.dao.UserQtsUpdatesDAO.Insert(do))
}

func (m *UpdateModel) AddSeqToUpdatesQueue(authId int64, userId, seq, updateType int32, updateData []byte) int32 {
	do := &dataobject.AuthSeqUpdatesDO{
		AuthId:     authId,
		UserId:     userId,
		UpdateType: updateType,
		UpdateData: updateData,
		Date2:      int32(time.Now().Unix()),
		Seq:        seq,
	}

	return int32(m.dao.AuthSeqUpdatesDAO.Insert(do))
}

//func GetAffectedMessage(userId, messageId int32) *mtproto.TLMessagesAffectedMessages {
//	doList := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectPtsByGTMessageID(userId, messageId)
//	if len(doList) == 0 {
//		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_OTHER2), fmt.Sprintf("GetAffectedMessage(%d, %d) empty", userId, messageId)))
//	}
//
//	affected := &mtproto.TLMessagesAffectedMessages{}
//	affected.Pts = doList[0].Pts
//	affected.PtsCount = int32(len(doList))
//	return affected
//}

//func updateToQueueData(update *mtproto.Update) int8 {
//	switch update.GetConstructor() {
//	case mtproto.TLConstructor_crc32_
//	}
//}

func getUpdateType(update *mtproto.Update) int8 {
	switch update.GetConstructor() {
	case mtproto.TLConstructor_CRC32_updateNewMessage:
		return PTS_UPDATE_NEW_MESSAGE
	case mtproto.TLConstructor_CRC32_updateDeleteMessages:
		return PTS_UPDATE_DELETE_MESSAGES
	case mtproto.TLConstructor_CRC32_updateReadHistoryOutbox:
		return PTS_UPDATE_READ_HISTORY_OUTBOX
	case mtproto.TLConstructor_CRC32_updateReadHistoryInbox:
		return PTS_UPDATE_READ_HISTORY_INBOX
	case mtproto.TLConstructor_CRC32_updateWebPage:
		return PTS_UPDATE_WEBPAGE
	case mtproto.TLConstructor_CRC32_updateReadMessagesContents:
		return PTS_UPDATE_READ_MESSAGE_CONENTS
	case mtproto.TLConstructor_CRC32_updateEditMessage:
		return PTS_UPDATE_EDIT_MESSAGE

	case mtproto.TLConstructor_CRC32_updateNewEncryptedMessage:
		return PTS_UPDATE_NEW_ENCRYPTED_MESSAGE

	case mtproto.TLConstructor_CRC32_updateNewChannelMessage:
		return PTS_UPDATE_NEW_CHANNEL_MESSAGE
	case mtproto.TLConstructor_CRC32_updateDeleteChannelMessages:
		return PTS_UPDATE_DELETE_CHANNEL_MESSAGES
	case mtproto.TLConstructor_CRC32_updateEditChannelMessage:
		return PTS_UPDATE_EDIT_CHANNEL_MESSAGE
	case mtproto.TLConstructor_CRC32_updateChannelWebPage:
		return PTS_UPDATE_EDIT_CHANNEL_WEBPAGE
	}
	return PTS_UPDATE_TYPE_UNKNOWN
}

func (m *UpdateModel) AddToPtsQueue(userId, pts, ptsCount int32, update *mtproto.Update) int32 {
	// TODO(@benqi): check error
	updateData, _ := json.Marshal(update)

	do := &dataobject.UserPtsUpdatesDO{
		UserId:     userId,
		Pts:        pts,
		PtsCount:   ptsCount,
		UpdateType: getUpdateType(update),
		UpdateData: string(updateData),
		Date2:      int32(time.Now().Unix()),
	}

	return int32(m.dao.UserPtsUpdatesDAO.Insert(do))
}

func (m *UpdateModel) AddToChannelPtsQueue(channelId, pts, ptsCount int32, update *mtproto.Update) int32 {
	// TODO(@benqi): check error
	updateData, _ := json.Marshal(update)

	do := &dataobject.ChannelPtsUpdatesDO{
		ChannelId:  channelId,
		Pts:        pts,
		PtsCount:   ptsCount,
		UpdateType: getUpdateType(update),
		UpdateData: string(updateData),
		Date2:      int32(time.Now().Unix()),
	}

	return int32(m.dao.ChannelPtsUpdatesDAO.Insert(do))
}

/*
func GetUpdatesByGtPts(userId, pts int32) (otherUpdates []*mtproto.Update, boxIDList []int32, lastPts int32) {
	lastPts = pts
	doList := dao.GetUserPtsUpdatesDAO(dao.DB_SLAVE).SelectByGtPts(userId, pts)
	if len(doList) == 0 {
		otherUpdates = []*mtproto.Update{}
		boxIDList = []int32{}
	} else {
		boxIDList = make([]int32, 0, len(doList))
		otherUpdates = make([]*mtproto.Update, 0, len(doList))
		for _, do := range doList {
			switch do.UpdateType {
			//  case PTS_UPDATE_SHORT_MESSAGE, PTS_UPDATE_SHORT_CHAT_MESSAGE:
			case PTS_READ_HISTORY_OUTBOX:
				readHistoryOutbox := &mtproto.TLUpdateReadHistoryOutbox{Data2: &mtproto.Update_Data{
					Peer_39:  base2.ToPeerByTypeAndID(do.PeerType, do.PeerId),
					MaxId:    do.MaxMessageBoxId,
					Pts:      do.Pts,
					PtsCount: 0,
				}}
				otherUpdates = append(otherUpdates, readHistoryOutbox.To_Update())
			case PTS_READ_HISTORY_INBOX:
				readHistoryInbox := &mtproto.TLUpdateReadHistoryInbox{Data2: &mtproto.Update_Data{
					Peer_39:  base2.ToPeerByTypeAndID(do.PeerType, do.PeerId),
					MaxId:    do.MaxMessageBoxId,
					Pts:      do.Pts,
					PtsCount: 0,
				}}
				otherUpdates = append(otherUpdates, readHistoryInbox.To_Update())
			//case PTS_MESSAGE_OUTBOX, PTS_MESSAGE_INBOX:
			//	boxIDList = append(boxIDList, do.MessageBoxId)
			}

			if do.Pts > lastPts {
				lastPts = do.Pts
			}
		}
	}
	return
}
*/

func (m *UpdateModel) GetUpdateListByGtPts(userId, pts int32) []*mtproto.Update {
	doList := m.dao.UserPtsUpdatesDAO.SelectByGtPts(userId, pts)
	if len(doList) == 0 {
		return []*mtproto.Update{}
	}

	updates := make([]*mtproto.Update, 0, len(doList))
	for _, do := range doList {
		update := &mtproto.Update{Constructor: mtproto.TLConstructor_CRC32_UNKNOWN, Data2: &mtproto.Update_Data{}}
		err := json.Unmarshal([]byte(do.UpdateData), update)
		if err != nil {
			glog.Errorf("unmarshal pts's update(%d)error: %v", do.Id, err)
			continue
		}
		if getUpdateType(update) != do.UpdateType {
			glog.Errorf("update data error.")
			continue
		}
		updates = append(updates, update)
	}
	return updates
}

func (m *UpdateModel) GetChannelUpdateListByGtPts(channelId, pts int32) []*mtproto.Update {
	doList := m.dao.ChannelPtsUpdatesDAO.SelectByGtPts(channelId, pts)
	if len(doList) == 0 {
		return []*mtproto.Update{}
	}

	updates := make([]*mtproto.Update, 0, len(doList))
	for _, do := range doList {
		update := &mtproto.Update{Constructor: mtproto.TLConstructor_CRC32_UNKNOWN, Data2: &mtproto.Update_Data{}}
		err := json.Unmarshal([]byte(do.UpdateData), update)
		if err != nil {
			glog.Errorf("unmarshal pts's update(%d)error: %v", do.Id, err)
			continue
		}
		if getUpdateType(update) != do.UpdateType {
			glog.Errorf("update data error.")
			continue
		}
		updates = append(updates, update)
	}
	return updates
}

func (m *UpdateModel) CheckAndFixAuthUpdateSeq(authKeyId int64, userId int32) {
	params := map[string]interface{}{
		"auth_key_id": authKeyId,
	}

	if !m.dao.CommonDAO.CheckExists("auth_updates_state", params) {
		do := &dataobject.AuthUpdatesStateDO{
			AuthKeyId: authKeyId,
			UserId:    userId,
			Pts:       0,
			Pts2:      0,
			Qts:       0,
			Qts2:      0,
			Seq:       -1,
			Seq2:      -1,
			Date:      int32(time.Now().Unix()),
		}
		m.dao.AuthUpdatesStateDAO.Insert(do)
	}
}

func (m *UpdateModel) GetUpdatesState2(authKeyId int64, userId int32) *mtproto.TLUpdatesState {
	// TODO(@benqi): insert auth_updates_state in auth.signUp
	m.CheckAndFixAuthUpdateSeq(authKeyId, userId)

	do := m.dao.AuthUpdatesStateDAO.SelectByAuthId(authKeyId)
	state := &mtproto.TLUpdatesState{Data2: &mtproto.Updates_State_Data{
		Pts:  do.Pts,
		Qts:  do.Qts,
		Seq:  do.Seq,
		Date: int32(time.Now().Unix()), // TODO(@benqi): do.Date2???
	}}
	return state
}

func (m *UpdateModel) GetServerUpdatesState(authKeyId int64, userId int32) *mtproto.TLUpdatesState {
	// TODO(@benqi): insert auth_updates_state in auth.signUp
	m.CheckAndFixAuthUpdateSeq(authKeyId, userId)

	do := m.dao.AuthUpdatesStateDAO.SelectByAuthId(authKeyId)
	state := &mtproto.TLUpdatesState{Data2: &mtproto.Updates_State_Data{
		Pts:  do.Pts2,
		Qts:  do.Qts2,
		Seq:  do.Seq2,
		Date: int32(time.Now().Unix()), // TODO(@benqi): do.Date2???
	}}
	return state
}

func (m *UpdateModel) UpdateAuthStateSeq(authKeyId int64, pts, qts int32) {
	m.dao.AuthUpdatesStateDAO.UpdatePtsAndQts(pts, qts, authKeyId)
}

func (m *UpdateModel) UpdateServerAuthStateSeq(authKeyId int64, pts, qts int32) {
	m.dao.AuthUpdatesStateDAO.UpdatePts2AndQts2(pts, qts, authKeyId)
}
