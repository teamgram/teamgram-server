/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/messages/internal/core"
)

// MessagesGetMessages
// messages.getMessages#63c66506 id:Vector<InputMessage> = messages.Messages;
func (s *Service) MessagesGetMessages(ctx context.Context, request *mtproto.TLMessagesGetMessages) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getMessages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetMessages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getMessages - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetHistory
// messages.getHistory#4423e6c5 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (s *Service) MessagesGetHistory(ctx context.Context, request *mtproto.TLMessagesGetHistory) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getHistory - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetHistory(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getHistory - reply: %s", r.DebugString())
	return r, err
}

// MessagesSearch
// messages.search#a0fda762 flags:# peer:InputPeer q:string from_id:flags.0?InputPeer top_msg_id:flags.1?int filter:MessagesFilter min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (s *Service) MessagesSearch(ctx context.Context, request *mtproto.TLMessagesSearch) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.search - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSearch(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.search - reply: %s", r.DebugString())
	return r, err
}

// MessagesReadHistory
// messages.readHistory#e306d3a peer:InputPeer max_id:int = messages.AffectedMessages;
func (s *Service) MessagesReadHistory(ctx context.Context, request *mtproto.TLMessagesReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.readHistory - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReadHistory(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.readHistory - reply: %s", r.DebugString())
	return r, err
}

// MessagesDeleteHistory
// messages.deleteHistory#b08f922a flags:# just_clear:flags.0?true revoke:flags.1?true peer:InputPeer max_id:int min_date:flags.2?int max_date:flags.3?int = messages.AffectedHistory;
func (s *Service) MessagesDeleteHistory(ctx context.Context, request *mtproto.TLMessagesDeleteHistory) (*mtproto.Messages_AffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.deleteHistory - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesDeleteHistory(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.deleteHistory - reply: %s", r.DebugString())
	return r, err
}

// MessagesDeleteMessages
// messages.deleteMessages#e58e95d2 flags:# revoke:flags.0?true id:Vector<int> = messages.AffectedMessages;
func (s *Service) MessagesDeleteMessages(ctx context.Context, request *mtproto.TLMessagesDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.deleteMessages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesDeleteMessages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.deleteMessages - reply: %s", r.DebugString())
	return r, err
}

// MessagesReceivedMessages
// messages.receivedMessages#5a954c0 max_id:int = Vector<ReceivedNotifyMessage>;
func (s *Service) MessagesReceivedMessages(ctx context.Context, request *mtproto.TLMessagesReceivedMessages) (*mtproto.Vector_ReceivedNotifyMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.receivedMessages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReceivedMessages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.receivedMessages - reply: %s", r.DebugString())
	return r, err
}

// MessagesSendMessage
// messages.sendMessage#d9d75a4 flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true peer:InputPeer reply_to_msg_id:flags.0?int message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int send_as:flags.13?InputPeer = Updates;
func (s *Service) MessagesSendMessage(ctx context.Context, request *mtproto.TLMessagesSendMessage) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.sendMessage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSendMessage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.sendMessage - reply: %s", r.DebugString())
	return r, err
}

// MessagesSendMedia
// messages.sendMedia#e25ff8e0 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true peer:InputPeer reply_to_msg_id:flags.0?int media:InputMedia message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int send_as:flags.13?InputPeer = Updates;
func (s *Service) MessagesSendMedia(ctx context.Context, request *mtproto.TLMessagesSendMedia) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.sendMedia - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSendMedia(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.sendMedia - reply: %s", r.DebugString())
	return r, err
}

// MessagesForwardMessages
// messages.forwardMessages#cc30290b flags:# silent:flags.5?true background:flags.6?true with_my_score:flags.8?true drop_author:flags.11?true drop_media_captions:flags.12?true noforwards:flags.14?true from_peer:InputPeer id:Vector<int> random_id:Vector<long> to_peer:InputPeer schedule_date:flags.10?int send_as:flags.13?InputPeer = Updates;
func (s *Service) MessagesForwardMessages(ctx context.Context, request *mtproto.TLMessagesForwardMessages) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.forwardMessages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesForwardMessages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.forwardMessages - reply: %s", r.DebugString())
	return r, err
}

// MessagesReadMessageContents
// messages.readMessageContents#36a73f77 id:Vector<int> = messages.AffectedMessages;
func (s *Service) MessagesReadMessageContents(ctx context.Context, request *mtproto.TLMessagesReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.readMessageContents - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReadMessageContents(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.readMessageContents - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetMessagesViews
// messages.getMessagesViews#5784d3e1 peer:InputPeer id:Vector<int> increment:Bool = messages.MessageViews;
func (s *Service) MessagesGetMessagesViews(ctx context.Context, request *mtproto.TLMessagesGetMessagesViews) (*mtproto.Messages_MessageViews, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getMessagesViews - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetMessagesViews(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getMessagesViews - reply: %s", r.DebugString())
	return r, err
}

// MessagesSearchGlobal
// messages.searchGlobal#4bc6589a flags:# folder_id:flags.0?int q:string filter:MessagesFilter min_date:int max_date:int offset_rate:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
func (s *Service) MessagesSearchGlobal(ctx context.Context, request *mtproto.TLMessagesSearchGlobal) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.searchGlobal - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSearchGlobal(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.searchGlobal - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetMessageEditData
// messages.getMessageEditData#fda68d36 peer:InputPeer id:int = messages.MessageEditData;
func (s *Service) MessagesGetMessageEditData(ctx context.Context, request *mtproto.TLMessagesGetMessageEditData) (*mtproto.Messages_MessageEditData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getMessageEditData - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetMessageEditData(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getMessageEditData - reply: %s", r.DebugString())
	return r, err
}

// MessagesEditMessage
// messages.editMessage#48f71778 flags:# no_webpage:flags.1?true peer:InputPeer id:int message:flags.11?string media:flags.14?InputMedia reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.15?int = Updates;
func (s *Service) MessagesEditMessage(ctx context.Context, request *mtproto.TLMessagesEditMessage) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.editMessage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesEditMessage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.editMessage - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetUnreadMentions
// messages.getUnreadMentions#46578472 peer:InputPeer offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (s *Service) MessagesGetUnreadMentions(ctx context.Context, request *mtproto.TLMessagesGetUnreadMentions) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getUnreadMentions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetUnreadMentions(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getUnreadMentions - reply: %s", r.DebugString())
	return r, err
}

// MessagesReadMentions
// messages.readMentions#f0189d3 peer:InputPeer = messages.AffectedHistory;
func (s *Service) MessagesReadMentions(ctx context.Context, request *mtproto.TLMessagesReadMentions) (*mtproto.Messages_AffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.readMentions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReadMentions(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.readMentions - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetRecentLocations
// messages.getRecentLocations#702a40e0 peer:InputPeer limit:int hash:long = messages.Messages;
func (s *Service) MessagesGetRecentLocations(ctx context.Context, request *mtproto.TLMessagesGetRecentLocations) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getRecentLocations - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetRecentLocations(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getRecentLocations - reply: %s", r.DebugString())
	return r, err
}

// MessagesSendMultiMedia
// messages.sendMultiMedia#f803138f flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true peer:InputPeer reply_to_msg_id:flags.0?int multi_media:Vector<InputSingleMedia> schedule_date:flags.10?int send_as:flags.13?InputPeer = Updates;
func (s *Service) MessagesSendMultiMedia(ctx context.Context, request *mtproto.TLMessagesSendMultiMedia) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.sendMultiMedia - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSendMultiMedia(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.sendMultiMedia - reply: %s", r.DebugString())
	return r, err
}

// MessagesUpdatePinnedMessage
// messages.updatePinnedMessage#d2aaf7ec flags:# silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer:InputPeer id:int = Updates;
func (s *Service) MessagesUpdatePinnedMessage(ctx context.Context, request *mtproto.TLMessagesUpdatePinnedMessage) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.updatePinnedMessage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesUpdatePinnedMessage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.updatePinnedMessage - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetSearchCounters
// messages.getSearchCounters#732eef00 peer:InputPeer filters:Vector<MessagesFilter> = Vector<messages.SearchCounter>;
func (s *Service) MessagesGetSearchCounters(ctx context.Context, request *mtproto.TLMessagesGetSearchCounters) (*mtproto.Vector_Messages_SearchCounter, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getSearchCounters - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetSearchCounters(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getSearchCounters - reply: %s", r.DebugString())
	return r, err
}

// MessagesUnpinAllMessages
// messages.unpinAllMessages#f025bc8b peer:InputPeer = messages.AffectedHistory;
func (s *Service) MessagesUnpinAllMessages(ctx context.Context, request *mtproto.TLMessagesUnpinAllMessages) (*mtproto.Messages_AffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.unpinAllMessages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesUnpinAllMessages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.unpinAllMessages - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetSearchResultsCalendar
// messages.getSearchResultsCalendar#49f0bde9 peer:InputPeer filter:MessagesFilter offset_id:int offset_date:int = messages.SearchResultsCalendar;
func (s *Service) MessagesGetSearchResultsCalendar(ctx context.Context, request *mtproto.TLMessagesGetSearchResultsCalendar) (*mtproto.Messages_SearchResultsCalendar, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getSearchResultsCalendar - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetSearchResultsCalendar(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getSearchResultsCalendar - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetSearchResultsPositions
// messages.getSearchResultsPositions#6e9583a3 peer:InputPeer filter:MessagesFilter offset_id:int limit:int = messages.SearchResultsPositions;
func (s *Service) MessagesGetSearchResultsPositions(ctx context.Context, request *mtproto.TLMessagesGetSearchResultsPositions) (*mtproto.Messages_SearchResultsPositions, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getSearchResultsPositions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetSearchResultsPositions(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getSearchResultsPositions - reply: %s", r.DebugString())
	return r, err
}

// MessagesToggleNoForwards
// messages.toggleNoForwards#b11eafa2 peer:InputPeer enabled:Bool = Updates;
func (s *Service) MessagesToggleNoForwards(ctx context.Context, request *mtproto.TLMessagesToggleNoForwards) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.toggleNoForwards - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesToggleNoForwards(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.toggleNoForwards - reply: %s", r.DebugString())
	return r, err
}

// MessagesSaveDefaultSendAs
// messages.saveDefaultSendAs#ccfddf96 peer:InputPeer send_as:InputPeer = Bool;
func (s *Service) MessagesSaveDefaultSendAs(ctx context.Context, request *mtproto.TLMessagesSaveDefaultSendAs) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.saveDefaultSendAs - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSaveDefaultSendAs(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.saveDefaultSendAs - reply: %s", r.DebugString())
	return r, err
}

// ChannelsGetSendAs
// channels.getSendAs#dc770ee peer:InputPeer = channels.SendAsPeers;
func (s *Service) ChannelsGetSendAs(ctx context.Context, request *mtproto.TLChannelsGetSendAs) (*mtproto.Channels_SendAsPeers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("channels.getSendAs - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsGetSendAs(request)
	if err != nil {
		return nil, err
	}

	c.Infof("channels.getSendAs - reply: %s", r.DebugString())
	return r, err
}

// MessagesTranslateText
// messages.translateText#24ce6dee flags:# peer:flags.0?InputPeer msg_id:flags.0?int text:flags.1?string from_lang:flags.2?string to_lang:string = messages.TranslatedText;
func (s *Service) MessagesTranslateText(ctx context.Context, request *mtproto.TLMessagesTranslateText) (*mtproto.Messages_TranslatedText, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.translateText - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesTranslateText(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.translateText - reply: %s", r.DebugString())
	return r, err
}

// MessagesSearchSentMedia
// messages.searchSentMedia#107e31a0 q:string filter:MessagesFilter limit:int = messages.Messages;
func (s *Service) MessagesSearchSentMedia(ctx context.Context, request *mtproto.TLMessagesSearchSentMedia) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.searchSentMedia - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSearchSentMedia(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.translateText - reply: %s", r.DebugString())
	return r, err
}
