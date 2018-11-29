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

package user

import (
	"github.com/golang/glog"
	base2 "github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
	"time"
)

func dialogDOToDialog(dialogDO *dataobject.UserDialogsDO) *mtproto.TLDialog {
	dialog := mtproto.NewTLDialog()
	// dialogData := dialog.GetData2()
	// draftIdList := make([]int32, 0)

	dialog.SetPinned(dialogDO.IsPinned == 1)
	dialog.SetPeer(base.ToPeerByTypeAndID(dialogDO.PeerType, dialogDO.PeerId))
	if dialogDO.PeerType == base.PEER_CHANNEL {
		// TODO(@benqi): only channel has pts
		// dialog.SetPts(messageBoxsDO.Pts)
		// peerChannlIdList = append(peerChannlIdList, dialogDO.PeerId)
		dialog.SetPts(dialogDO.Pts)
	}

	dialog.SetTopMessage(dialogDO.TopMessage)
	dialog.SetReadInboxMaxId(dialogDO.ReadInboxMaxId)
	dialog.SetReadOutboxMaxId(dialogDO.ReadOutboxMaxId)
	dialog.SetUnreadCount(dialogDO.UnreadCount)
	dialog.SetUnreadMentionsCount(dialogDO.UnreadMentionsCount)

	// TODO(@benqi): draft message.
	dialog.SetDraft(mtproto.NewTLDraftMessageEmpty().To_DraftMessage())
	//if dialogDO.DraftType == 2 {
	//	draft := &mtproto.DraftMessage{}
	//	err := json.Unmarshal([]byte(dialogDO.DraftMessageData), &draft)
	//	if err == nil {
	//		dialog.SetDraft(draft)
	//	}
	//}

	// NotifySettings
	peerNotifySettings := mtproto.NewTLPeerNotifySettings()
	if dialogDO.ShowPreviews == 1 {
		peerNotifySettings.SetShowPreviews(mtproto.ToBool(true))
	}
	if dialogDO.Silent == 1 {
		peerNotifySettings.SetSilent(mtproto.ToBool(true))
	}
	if dialogDO.MuteUntil == 0 {
		peerNotifySettings.SetMuteUntil(1)
	} else {
		peerNotifySettings.SetMuteUntil(dialogDO.MuteUntil)
	}
	if dialogDO.Sound == "" {
		peerNotifySettings.SetSound("default")
	} else {
		peerNotifySettings.SetSound(dialogDO.Sound)
	}
	dialog.SetNotifySettings(peerNotifySettings.To_PeerNotifySettings())
	return dialog
}

func dialogDOListToDialogList(dialogDOList []dataobject.UserDialogsDO) (dialogs []*mtproto.Dialog) {
	draftIdList := make([]int32, 0)
	for _, dialogDO := range dialogDOList {
		if dialogDO.DraftId > 0 {
			draftIdList = append(draftIdList, dialogDO.DraftId)
		}
		dialogs = append(dialogs, dialogDOToDialog(&dialogDO).To_Dialog())
	}

	// TODO(@benqi): fetch draft message list
	return
}

func (m *UserModel) GetDialogsByOffsetId(userId int32, isPinned bool, offsetId int32, limit int32) (dialogs []*mtproto.Dialog) {
	dialogDOList := m.dao.UserDialogsDAO.SelectByPinnedAndOffset(
		userId, base2.BoolToInt8(isPinned), offsetId, limit)
	glog.Infof("GetDialogsByOffsetId - dialogDOList: {%v}, query: {userId: %d, isPinned: %v, offeetId: %d, limit: %d ", dialogDOList, userId, isPinned, offsetId, limit)

	dialogs = dialogDOListToDialogList(dialogDOList)
	return
}

//func (m *dialogModel) GetDialogsByOffsetDate(userId int32, excludePinned bool, offsetData int32, limit int32) (dialogs []*mtproto.TLDialog) {
//	dialogDOList := m.dao.UserDialogsDAO.SelectDialogsByPinnedAndOffsetDate(
//		userId, base2.BoolToInt8(!excludePinned), offsetData, limit)
//	for _, dialogDO := range dialogDOList {
//		dialogs = append(dialogs, dialogDOToDialog(&dialogDO))
//	}
//	return
//}

func (m *UserModel) GetDialogsByUserIDAndType(userId, peerType int32) (dialogs []*mtproto.Dialog) {
	var dialogDOList []dataobject.UserDialogsDO
	if peerType == base.PEER_UNKNOWN || peerType == base.PEER_EMPTY {
		dialogDOList = m.dao.UserDialogsDAO.SelectDialogsByUserID(userId)
		glog.Infof("SelectDialogsByUserID(%d) - {%v}", userId, dialogDOList)
	} else {
		dialogDOList = m.dao.UserDialogsDAO.SelectDialogsByPeerType(userId, int8(peerType))
		glog.Infof("SelectDialogsByPeerType(%d, %d) - {%v}", userId, int32(peerType), dialogDOList)
	}

	dialogs = dialogDOListToDialogList(dialogDOList)
	// glog.Infof("SelectDialogsByPeerType(%d, %d) - {%v}", userId, int32(peerType), dialogs)
	return
}

func (m *UserModel) GetPinnedDialogs(userId int32) (dialogs []*mtproto.Dialog) {
	dialogDOList := m.dao.UserDialogsDAO.SelectPinnedDialogs(userId)
	dialogs = dialogDOListToDialogList(dialogDOList)
	return
}

func (m *UserModel) GetPeersDialogs(selfId int32, peers []*base.PeerUtil) (dialogs []*mtproto.Dialog) {
	for _, peer := range peers {
		// peerUtil := base.FromInputPeer2(selfId, peer)
		dialogDO := m.dao.UserDialogsDAO.SelectByPeer(selfId, int8(peer.PeerType), peer.PeerId)
		if dialogDO != nil {
			dialogs = append(dialogs, dialogDOToDialog(dialogDO).To_Dialog())
		}
	}
	return
}

// 发件箱
func (m *UserModel) CreateOrUpdateByOutbox(userId, peerType int32, peerId int32, topMessage int32, unreadMentions, clearDraft bool) {
	var (
		affectedRows = int64(0)
		date         = int32(time.Now().Unix())
	)

	if clearDraft && unreadMentions {
		affectedRows = m.dao.UserDialogsDAO.UpdateTopMessageAndMentionsAndClearDraft(topMessage, date, userId, int8(peerType), peerId)
	} else if clearDraft && !unreadMentions {
		affectedRows = m.dao.UserDialogsDAO.UpdateTopMessageAndClearDraft(topMessage, date, userId, int8(peerType), peerId)
	} else if !clearDraft && unreadMentions {
		affectedRows = m.dao.UserDialogsDAO.UpdateTopMessageAndMentions(topMessage, date, userId, int8(peerType), peerId)
	} else {
		affectedRows = m.dao.UserDialogsDAO.UpdateTopMessage(topMessage, date, userId, int8(peerType), peerId)
	}

	if affectedRows == 0 {
		// 创建会话
		dialog := &dataobject.UserDialogsDO{}
		dialog.UserId = userId
		dialog.PeerType = int8(peerType)
		dialog.PeerId = peerId
		if unreadMentions {
			dialog.UnreadMentionsCount = 1
		} else {
			dialog.UnreadMentionsCount = 0
		}
		dialog.UnreadCount = 0
		dialog.TopMessage = topMessage
		dialog.CreatedAt = base2.NowFormatYMDHMS()
		dialog.Date2 = date
		m.dao.UserDialogsDAO.Insert(dialog)
	}
	return
}

// 收件箱
func (m *UserModel) CreateOrUpdateByInbox(userId, peerType int32, peerId int32, topMessage int32, unreadMentions bool) {
	var (
		affectedRows = int64(0)
		date         = int32(time.Now().Unix())
	)

	if unreadMentions {
		affectedRows = m.dao.UserDialogsDAO.UpdateTopMessageAndUnreadAndMentions(topMessage, date, userId, int8(peerType), peerId)
	} else {
		affectedRows = m.dao.UserDialogsDAO.UpdateTopMessageAndUnread(topMessage, date, userId, int8(peerType), peerId)
	}

	glog.Info("createOrUpdateByInbox - ", affectedRows)
	if affectedRows == 0 {
		// 创建会话
		dialog := &dataobject.UserDialogsDO{}
		dialog.UserId = userId
		dialog.PeerType = int8(peerType)
		dialog.PeerId = peerId
		if unreadMentions {
			dialog.UnreadMentionsCount = 1
		} else {
			dialog.UnreadMentionsCount = 0
		}
		dialog.UnreadCount = 1
		dialog.TopMessage = topMessage
		dialog.CreatedAt = base2.NowFormatYMDHMS()
		dialog.DraftMessageData = ""
		dialog.Date2 = date
		m.dao.UserDialogsDAO.Insert(dialog)
	} else {

	}
	return
}
