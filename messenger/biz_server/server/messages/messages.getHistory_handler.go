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
	"math"
)

// From android client
//
// load_type == 0 ? backward loading
// load_type == 1 ? forward loading
// load_type == 2 ? load from first unread
// load_type == 3 ? load around message
// load_type == 4 ? load around date
/*
  // @benqi: 这什么鬼规则啊？？？
  1. getHistory, ps: max_id:int min_id:int未使用
	TLRPC.TL_messages_getHistory req = new TLRPC.TL_messages_getHistory();
	req.peer = getInputPeer(lower_part);
	if (load_type == 4) {
		req.add_offset = -count + 5;
	} else if (load_type == 3) {
		req.add_offset = -count / 2;
	} else if (load_type == 1) {
		req.add_offset = -count - 1;
	} else if (load_type == 2 && max_id != 0) {
		req.add_offset = -count + 6;
	} else {
		if (lower_part < 0 && max_id != 0) {
			TLRPC.Chat chat = getChat(-lower_part);
			if (ChatObject.isChannel(chat)) {
				req.add_offset = -1;
				req.limit += 1;
			}
		}
	}
	req.limit = count;
	req.offset_id = max_id;
	req.offset_date = offset_date;

  2. Load dialog last message, ps: limit = 1
	TLRPC.TL_messages_getHistory req = new TLRPC.TL_messages_getHistory();
	req.peer = peer == null ? getInputPeer(lower_id) : peer;
	if (req.peer == null) {
		return;
	}
	req.limit = 1;
*/

// From tdesktop client
/*
  1. void MainWindow::sendServiceHistoryRequest() {
	auto offsetId = 0;
	auto offsetDate = 0;
	auto addOffset = 0;
	auto limit = 1;
	auto maxId = 0;
	auto minId = 0;
	auto historyHash = 0;
	_serviceHistoryRequest = MTP::send(

  2. void MainWidget::checkPeerHistory(PeerData *peer) {
	auto offsetId = 0;
	auto offsetDate = 0;
	auto addOffset = 0;
	auto limit = 1;
	auto maxId = 0;
	auto minId = 0;
	auto historyHash = 0;
	MTP::send(
		MTPmessages_GetHistory(

  3. void ApiWrap::requestMessageAfterDate(
	// API returns a message with date <= offset_date.
	// So we request a message with offset_date = desired_date - 1 and add_offset = -1.
	// This should give us the first message with date >= desired_date.
	auto offsetId = 0;
	auto offsetDate = static_cast<int>(QDateTime(date).toTime_t()) - 1;
	auto addOffset = -1;
	auto limit = 1;
	auto maxId = 0;
	auto minId = 0;
	auto historyHash = 0;
	request(MTPmessages_GetHistory(

  4. void HistoryWidget::delayedShowAt(MsgId showAtMsgId) {
	_delayedShowAtMsgId = showAtMsgId;

	auto from = _peer;
	auto offsetId = 0;
	auto offset = 0;
	auto loadCount = kMessagesPerPage;
	if (_delayedShowAtMsgId == ShowAtUnreadMsgId) {
		if (_migrated && _migrated->unreadCount()) {
			from = _migrated->peer;
			offset = -loadCount / 2;
			offsetId = _migrated->inboxReadBefore;
		} else if (_history->unreadCount()) {
			offset = -loadCount / 2;
			offsetId = _history->inboxReadBefore;
		} else {
			loadCount = kMessagesPerPageFirst;
		}
	} else if (_delayedShowAtMsgId == ShowAtTheEndMsgId) {
		loadCount = kMessagesPerPageFirst;
	} else if (_delayedShowAtMsgId > 0) {
		offset = -loadCount / 2;
		offsetId = _delayedShowAtMsgId;
	} else if (_delayedShowAtMsgId < 0 && _history->isChannel()) {
		if (_delayedShowAtMsgId < 0 && -_delayedShowAtMsgId < ServerMaxMsgId && _migrated) {
			from = _migrated->peer;
			offset = -loadCount / 2;
			offsetId = -_delayedShowAtMsgId;
		}
	}
	auto offsetDate = 0;
	auto maxId = 0;
	auto minId = 0;
	auto historyHash = 0;

	_delayedShowAtRequest = MTP::send(
		MTPmessages_GetHistory(

  5. void HistoryWidget::loadMessagesDown() {
	if (!_history || _preloadDownRequest) return;

	if (_history->isEmpty() && _migrated && _migrated->isEmpty()) {
		return firstLoadMessages();
	}

	auto loadMigrated = _migrated && !(_migrated->isEmpty() || _migrated->loadedAtBottom() || (!_history->isEmpty() && !_history->loadedAtTop()));
	auto from = loadMigrated ? _migrated : _history;
	if (from->loadedAtBottom()) {
		return;
	}

	auto loadCount = kMessagesPerPage;
	auto addOffset = -loadCount;
	auto offsetId = from->maxMsgId();
	if (!offsetId) {
		if (loadMigrated || !_migrated) return;
		++offsetId;
		++addOffset;
	}
	auto offsetDate = 0;
	auto maxId = 0;
	auto minId = 0;
	auto historyHash = 0;

	_debug_preloadDownOffsetId = offsetId + 1;
	_debug_preloadDownAddOffset = addOffset;
	_debug_preloadDownLoadCount = loadCount;
	_debug_preloadDownPeer = from->peer->id;
	_preloadDownRequest = MTP::send(
		MTPmessages_GetHistory(

  6. void HistoryWidget::loadMessages() {
	if (!_history || _preloadRequest) return;

	if (_history->isEmpty() && _migrated && _migrated->isEmpty()) {
		return firstLoadMessages();
	}

	auto loadMigrated = _migrated && (_history->isEmpty() || _history->loadedAtTop() || (!_migrated->isEmpty() && !_migrated->loadedAtBottom()));
	auto from = loadMigrated ? _migrated : _history;
	if (from->loadedAtTop()) {
		return;
	}

	auto offsetId = from->minMsgId();
	auto addOffset = 0;
	auto loadCount = offsetId
		? kMessagesPerPage
		: kMessagesPerPageFirst;
	auto offsetDate = 0;
	auto maxId = 0;
	auto minId = 0;
	auto historyHash = 0;

	_debug_preloadOffsetId = offsetId + 1;
	_debug_preloadAddOffset = addOffset;
	_debug_preloadLoadCount = loadCount;
	_debug_preloadPeer = from->peer->id;
	_preloadRequest = MTP::send(
		MTPmessages_GetHistory(

  7. void HistoryWidget::firstLoadMessages() {
	if (!_history || _firstLoadRequest) return;

	auto from = _peer;
	auto offsetId = 0;
	auto offset = 0;
	auto loadCount = kMessagesPerPage;
	if (_showAtMsgId == ShowAtUnreadMsgId) {
		if (_migrated && _migrated->unreadCount()) {
			_history->getReadyFor(_showAtMsgId);
			from = _migrated->peer;
			offset = -loadCount / 2;
			offsetId = _migrated->inboxReadBefore;
		} else if (_history->unreadCount()) {
			_history->getReadyFor(_showAtMsgId);
			offset = -loadCount / 2;
			offsetId = _history->inboxReadBefore;
		} else {
			_history->getReadyFor(ShowAtTheEndMsgId);
		}
	} else if (_showAtMsgId == ShowAtTheEndMsgId) {
		_history->getReadyFor(_showAtMsgId);
		loadCount = kMessagesPerPageFirst;
	} else if (_showAtMsgId > 0) {
		_history->getReadyFor(_showAtMsgId);
		offset = -loadCount / 2;
		offsetId = _showAtMsgId;
	} else if (_showAtMsgId < 0 && _history->isChannel()) {
		if (_showAtMsgId < 0 && -_showAtMsgId < ServerMaxMsgId && _migrated) {
			_history->getReadyFor(_showAtMsgId);
			from = _migrated->peer;
			offset = -loadCount / 2;
			offsetId = -_showAtMsgId;
		} else if (_showAtMsgId == SwitchAtTopMsgId) {
			_history->getReadyFor(_showAtMsgId);
		}
	}

	auto offsetDate = 0;
	auto maxId = 0;
	auto minId = 0;
	auto historyHash = 0;

	_firstLoadRequest = MTP::send(
		MTPmessages_GetHistory(
*/

const (
	kLoadTypeBackward           = 0
	kLoadTypeForward            = 1
	kLoadTypeFirstUnread        = 2
	kLoadTypeFirstAroundMessage = 3
	kLoadTypeFirstAroundDate    = 4

	kLoadTypeLimit1 = 16
)

// TODO(@benqi): only android client
// limit = count
// offset_id = max_id
func calcLoadHistoryType(isChannel bool, offsetId, offsetDate, addOffset, limit, maxId, minId int32) int {
	if limit == 1 {
		return kLoadTypeLimit1
	}

	// check isChannel??
	if isChannel && addOffset == -1 && maxId != 0 {
		return kLoadTypeBackward
	}

	if addOffset == 0 {
		return kLoadTypeBackward
	} else if addOffset == -1 {
		return kLoadTypeBackward
	} else if addOffset == -limit+5 {
		return kLoadTypeFirstAroundDate
	} else if addOffset == -limit/2 {
		return kLoadTypeFirstAroundMessage
	} else if addOffset == -limit-1 {
		return kLoadTypeForward
	} else if addOffset == -limit+6 {
		if maxId != 0 {
			return kLoadTypeFirstUnread
		}
	}
	return kLoadTypeForward
}

func (s *MessagesServiceImpl) loadHistoryMessage(loadType int, selfUserId int32, peer *base.PeerUtil, offsetId, offsetDate, addOffset, limit, maxId, minId int32) []*mtproto.Message {
	messages := []*mtproto.Message{}

	switch loadType {
	case kLoadTypeLimit1:
		// 1. Load dialog last messag
		offsetId = math.MaxInt32
		messages = s.MessageModel.LoadBackwardHistoryMessages(selfUserId, peer.PeerType, peer.PeerId, offsetId, limit)
	case kLoadTypeBackward:
		if addOffset == 0 || addOffset == -1 && offsetId == 0 {
			offsetId = math.MaxInt32
		}
		messages = s.MessageModel.LoadBackwardHistoryMessages(selfUserId, peer.PeerType, peer.PeerId, offsetId, addOffset+limit)
	case kLoadTypeFirstAroundDate:
	case kLoadTypeFirstAroundMessage:
		// LOAD_HISTORY_TYPE_FORWARD and LOAD_HISTORY_TYPE_BACKWARD
		// 按升序排
		messages1 := s.MessageModel.LoadForwardHistoryMessages(selfUserId, peer.PeerType, peer.PeerId, offsetId, -addOffset)
		for i, j := 0, len(messages1)-1; i < j; i, j = i+1, j-1 {
			messages1[i], messages1[j] = messages1[j], messages1[i]
		}
		messages = append(messages, messages1...)
		// 降序
		messages2 := s.MessageModel.LoadBackwardHistoryMessages(selfUserId, peer.PeerType, peer.PeerId, offsetId, limit+addOffset)
		glog.Info(messages2)
		messages = append(messages, messages2...)
	case kLoadTypeForward:
		messages = s.MessageModel.LoadForwardHistoryMessages(selfUserId, peer.PeerType, peer.PeerId, offsetId, -addOffset)
		for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
			messages[i], messages[j] = messages[j], messages[i]
		}
	case kLoadTypeFirstUnread:
		messages = s.MessageModel.LoadBackwardHistoryMessages(selfUserId, peer.PeerType, peer.PeerId, offsetId, addOffset+limit)
	}

	return messages
}

func (s *MessagesServiceImpl) getHistoryMessages(md *grpc_util.RpcMetadata, request *mtproto.TLMessagesGetHistory) (messagesMessages *mtproto.Messages_Messages) {
	peer := base.FromInputPeer(request.GetPeer())
	if peer.PeerType == base.PEER_SELF {
		peer.PeerType = base.PEER_USER
		peer.PeerId = md.UserId
	}

	offsetId := request.GetOffsetId()
	addOffset := request.GetAddOffset()
	limit := request.GetLimit()

	var (
		isChannel = peer.PeerType == base.PEER_CHANNEL
		users     []*mtproto.User
		chats     []*mtproto.Chat
	)

	loadType := calcLoadHistoryType(isChannel, offsetId, request.GetOffsetDate(), addOffset, limit, request.GetMaxId(), request.GetMinId())
	messages := s.loadHistoryMessage(loadType, md.UserId, peer, offsetId, request.GetOffsetDate(), addOffset, limit, request.GetMaxId(), request.GetMinId())

	// messagesMessages.SetMessages(messages)
	userIdList, chatIdList, _ := message.PickAllIDListByMessages(messages)
	if len(userIdList) > 0 {
		users = s.UserModel.GetUserListByIdList(md.UserId, userIdList)
		// messagesMessages.Data2.Users = users
	} else {
		users = []*mtproto.User{}
	}

	if len(chatIdList) > 0 {
		chats = s.ChatModel.GetChatListBySelfAndIDList(md.UserId, chatIdList)
	} else {
		chats = []*mtproto.Chat{}
	}

	// TODO(@benqi): Add channel's pts
	messagesSlice := &mtproto.TLMessagesMessages{Data2: &mtproto.Messages_Messages_Data{
		Messages: messages,
		Chats:    chats,
		Users:    users,
	}}
	messagesMessages = messagesSlice.To_Messages_Messages()

	return
}

// request: {"peer":{"constructor":2072935910,"data2":{"user_id":2,"access_hash":5166926832673632934}},"offset_id":2834,"limit":50}
// request: {"peer":{"constructor":2072935910,"data2":{"user_id":5,"access_hash":1006843769775067136}},"offset_id":1,"add_offset":-25,"limit":50}
// request: {"peer":{"constructor":2072935910,"data2":{"user_id":4,"access_hash":405858233924775823}},"offset_id":2147483647,"offset_date":2147483647,"limit":1,"max_id":2147483647,"min_id":1}
// request: {"peer":{"constructor":2072935910,"data2":{"user_id":4,"access_hash":405858233924775823}},"offset_id":2147483647,"offset_date":2147483647,"limit":1,"max_id":2147483647,"min_id":1}
// messages.getHistory#dcbb8260 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:int = messages.Messages;
func (s *MessagesServiceImpl) MessagesGetHistory(ctx context.Context, request *mtproto.TLMessagesGetHistory) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getHistory#dcbb8260 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	messagesMessages := s.getHistoryMessages(md, request)

	glog.Infof("messages.getHistory#dcbb8260 - reply: %s", logger.JsonDebugData(messagesMessages))
	return messagesMessages, nil
}
