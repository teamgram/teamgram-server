/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"math"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
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
	loadTypeBackward           = 0
	loadTypeForward            = 1
	loadTypeFirstUnread        = 2
	loadTypeFirstAroundMessage = 3
	loadTypeFirstAroundDate    = 4
	loadTypeLimit1             = 16
)

// MessageGetHistoryMessages
// message.getHistoryMessages user_id:long peer_type:int peer_id:long offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = Vector<MessageBox>;
func (c *MessageCore) MessageGetHistoryMessages(in *message.TLMessageGetHistoryMessages) (*message.Vector_MessageBox, error) {
	var (
		selfUserId = in.UserId
		peer       = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
		addOffset  = in.AddOffset
		limit      = in.Limit
		offsetId   = in.OffsetId
		minId      = in.MinId
		maxId      = in.MaxId
		hash       = in.Hash
		boxList    []*mtproto.MessageBox
	)

	loadType := loadTypeBackward
	if addOffset >= 0 {
		loadType = loadTypeBackward
	} else if addOffset+limit > 0 {
		loadType = loadTypeFirstAroundDate
	} else {
		loadType = loadTypeForward
	}

	if offsetId == 0 {
		offsetId = math.MaxInt32
	}
	//if offsetDate > 0 {
	//	switch loadType {
	//	case loadTypeBackward:
	//	case loadTypeFirstAroundDate:
	//	case loadTypeForward:
	//	}
	//} else
	{
		switch loadType {
		case loadTypeBackward:
			if offsetId == 0 {
				offsetId = math.MaxInt32
			}
			// c.svcCtx.Dao.MessageClient.MessageGet
			boxList = c.svcCtx.Dao.GetOffsetIdBackwardHistoryMessages(c.ctx, selfUserId, peer, offsetId, minId, maxId, addOffset+limit, hash)
		case loadTypeFirstAroundDate:
			boxList1 := c.svcCtx.GetOffsetIdForwardHistoryMessages(c.ctx, selfUserId, peer, offsetId, minId, maxId, -addOffset, hash)
			for i, j := 0, len(boxList1)-1; i < j; i, j = i+1, j-1 {
				boxList1[i], boxList1[j] = boxList1[j], boxList1[i]
			}
			boxList = append(boxList, boxList1...)
			// 降序
			boxList2 := c.svcCtx.Dao.GetOffsetIdBackwardHistoryMessages(c.ctx, selfUserId, peer, offsetId, minId, maxId, limit+addOffset, hash)
			// log.Infof("%v", messages2)
			boxList = append(boxList, boxList2...)
		case loadTypeForward:
			boxList = c.svcCtx.Dao.GetOffsetIdForwardHistoryMessages(c.ctx, selfUserId, peer, offsetId, minId, maxId, -addOffset, hash)
			for i, j := 0, len(boxList)-1; i < j; i, j = i+1, j-1 {
				boxList[i], boxList[j] = boxList[j], boxList[i]
			}
		}
	}
	return &message.Vector_MessageBox{
		Datas: boxList,
	}, nil
}
