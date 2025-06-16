/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/core"
)

// MessagesGetMessages
// messages.getMessages#63c66506 id:Vector<InputMessage> = messages.Messages;
func (s *Service) MessagesGetMessages(ctx context.Context, request *tg.TLMessagesGetMessages) (*tg.MessagesMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getMessages - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getMessages - reply: {%v}", r)
	return r, err
}

// MessagesGetHistory
// messages.getHistory#4423e6c5 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (s *Service) MessagesGetHistory(ctx context.Context, request *tg.TLMessagesGetHistory) (*tg.MessagesMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getHistory - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getHistory - reply: {%v}", r)
	return r, err
}

// MessagesSearch
// messages.search#29ee847a flags:# peer:InputPeer q:string from_id:flags.0?InputPeer saved_peer_id:flags.2?InputPeer saved_reaction:flags.3?Vector<Reaction> top_msg_id:flags.1?int filter:MessagesFilter min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (s *Service) MessagesSearch(ctx context.Context, request *tg.TLMessagesSearch) (*tg.MessagesMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.search - metadata: {}, request: {%v}", request)

	r, err := c.MessagesSearch(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.search - reply: {%v}", r)
	return r, err
}

// MessagesReadHistory
// messages.readHistory#e306d3a peer:InputPeer max_id:int = messages.AffectedMessages;
func (s *Service) MessagesReadHistory(ctx context.Context, request *tg.TLMessagesReadHistory) (*tg.MessagesAffectedMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.readHistory - metadata: {}, request: {%v}", request)

	r, err := c.MessagesReadHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.readHistory - reply: {%v}", r)
	return r, err
}

// MessagesDeleteHistory
// messages.deleteHistory#b08f922a flags:# just_clear:flags.0?true revoke:flags.1?true peer:InputPeer max_id:int min_date:flags.2?int max_date:flags.3?int = messages.AffectedHistory;
func (s *Service) MessagesDeleteHistory(ctx context.Context, request *tg.TLMessagesDeleteHistory) (*tg.MessagesAffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.deleteHistory - metadata: {}, request: {%v}", request)

	r, err := c.MessagesDeleteHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.deleteHistory - reply: {%v}", r)
	return r, err
}

// MessagesDeleteMessages
// messages.deleteMessages#e58e95d2 flags:# revoke:flags.0?true id:Vector<int> = messages.AffectedMessages;
func (s *Service) MessagesDeleteMessages(ctx context.Context, request *tg.TLMessagesDeleteMessages) (*tg.MessagesAffectedMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.deleteMessages - metadata: {}, request: {%v}", request)

	r, err := c.MessagesDeleteMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.deleteMessages - reply: {%v}", r)
	return r, err
}

// MessagesReceivedMessages
// messages.receivedMessages#5a954c0 max_id:int = Vector<ReceivedNotifyMessage>;
func (s *Service) MessagesReceivedMessages(ctx context.Context, request *tg.TLMessagesReceivedMessages) (*tg.VectorReceivedNotifyMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.receivedMessages - metadata: {}, request: {%v}", request)

	r, err := c.MessagesReceivedMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.receivedMessages - reply: {%v}", r)
	return r, err
}

// MessagesSendMessage
// messages.sendMessage#983f9745 flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true update_stickersets_order:flags.15?true invert_media:flags.16?true allow_paid_floodskip:flags.19?true peer:InputPeer reply_to:flags.0?InputReplyTo message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut effect:flags.18?long = Updates;
func (s *Service) MessagesSendMessage(ctx context.Context, request *tg.TLMessagesSendMessage) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.sendMessage - metadata: {}, request: {%v}", request)

	r, err := c.MessagesSendMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.sendMessage - reply: {%v}", r)
	return r, err
}

// MessagesSendMedia
// messages.sendMedia#7852834e flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true update_stickersets_order:flags.15?true invert_media:flags.16?true allow_paid_floodskip:flags.19?true peer:InputPeer reply_to:flags.0?InputReplyTo media:InputMedia message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut effect:flags.18?long = Updates;
func (s *Service) MessagesSendMedia(ctx context.Context, request *tg.TLMessagesSendMedia) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.sendMedia - metadata: {}, request: {%v}", request)

	r, err := c.MessagesSendMedia(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.sendMedia - reply: {%v}", r)
	return r, err
}

// MessagesForwardMessages
// messages.forwardMessages#d5039208 flags:# silent:flags.5?true background:flags.6?true with_my_score:flags.8?true drop_author:flags.11?true drop_media_captions:flags.12?true noforwards:flags.14?true allow_paid_floodskip:flags.19?true from_peer:InputPeer id:Vector<int> random_id:Vector<long> to_peer:InputPeer top_msg_id:flags.9?int schedule_date:flags.10?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut = Updates;
func (s *Service) MessagesForwardMessages(ctx context.Context, request *tg.TLMessagesForwardMessages) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.forwardMessages - metadata: {}, request: {%v}", request)

	r, err := c.MessagesForwardMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.forwardMessages - reply: {%v}", r)
	return r, err
}

// MessagesReadMessageContents
// messages.readMessageContents#36a73f77 id:Vector<int> = messages.AffectedMessages;
func (s *Service) MessagesReadMessageContents(ctx context.Context, request *tg.TLMessagesReadMessageContents) (*tg.MessagesAffectedMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.readMessageContents - metadata: {}, request: {%v}", request)

	r, err := c.MessagesReadMessageContents(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.readMessageContents - reply: {%v}", r)
	return r, err
}

// MessagesGetMessagesViews
// messages.getMessagesViews#5784d3e1 peer:InputPeer id:Vector<int> increment:Bool = messages.MessageViews;
func (s *Service) MessagesGetMessagesViews(ctx context.Context, request *tg.TLMessagesGetMessagesViews) (*tg.MessagesMessageViews, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getMessagesViews - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetMessagesViews(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getMessagesViews - reply: {%v}", r)
	return r, err
}

// MessagesSearchGlobal
// messages.searchGlobal#4bc6589a flags:# broadcasts_only:flags.1?true folder_id:flags.0?int q:string filter:MessagesFilter min_date:int max_date:int offset_rate:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
func (s *Service) MessagesSearchGlobal(ctx context.Context, request *tg.TLMessagesSearchGlobal) (*tg.MessagesMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.searchGlobal - metadata: {}, request: {%v}", request)

	r, err := c.MessagesSearchGlobal(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.searchGlobal - reply: {%v}", r)
	return r, err
}

// MessagesGetMessageEditData
// messages.getMessageEditData#fda68d36 peer:InputPeer id:int = messages.MessageEditData;
func (s *Service) MessagesGetMessageEditData(ctx context.Context, request *tg.TLMessagesGetMessageEditData) (*tg.MessagesMessageEditData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getMessageEditData - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetMessageEditData(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getMessageEditData - reply: {%v}", r)
	return r, err
}

// MessagesEditMessage
// messages.editMessage#dfd14005 flags:# no_webpage:flags.1?true invert_media:flags.16?true peer:InputPeer id:int message:flags.11?string media:flags.14?InputMedia reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.15?int quick_reply_shortcut_id:flags.17?int = Updates;
func (s *Service) MessagesEditMessage(ctx context.Context, request *tg.TLMessagesEditMessage) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.editMessage - metadata: {}, request: {%v}", request)

	r, err := c.MessagesEditMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.editMessage - reply: {%v}", r)
	return r, err
}

// MessagesGetUnreadMentions
// messages.getUnreadMentions#f107e790 flags:# peer:InputPeer top_msg_id:flags.0?int offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (s *Service) MessagesGetUnreadMentions(ctx context.Context, request *tg.TLMessagesGetUnreadMentions) (*tg.MessagesMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getUnreadMentions - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetUnreadMentions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getUnreadMentions - reply: {%v}", r)
	return r, err
}

// MessagesReadMentions
// messages.readMentions#36e5bf4d flags:# peer:InputPeer top_msg_id:flags.0?int = messages.AffectedHistory;
func (s *Service) MessagesReadMentions(ctx context.Context, request *tg.TLMessagesReadMentions) (*tg.MessagesAffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.readMentions - metadata: {}, request: {%v}", request)

	r, err := c.MessagesReadMentions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.readMentions - reply: {%v}", r)
	return r, err
}

// MessagesGetRecentLocations
// messages.getRecentLocations#702a40e0 peer:InputPeer limit:int hash:long = messages.Messages;
func (s *Service) MessagesGetRecentLocations(ctx context.Context, request *tg.TLMessagesGetRecentLocations) (*tg.MessagesMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getRecentLocations - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetRecentLocations(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getRecentLocations - reply: {%v}", r)
	return r, err
}

// MessagesSendMultiMedia
// messages.sendMultiMedia#37b74355 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true update_stickersets_order:flags.15?true invert_media:flags.16?true allow_paid_floodskip:flags.19?true peer:InputPeer reply_to:flags.0?InputReplyTo multi_media:Vector<InputSingleMedia> schedule_date:flags.10?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut effect:flags.18?long = Updates;
func (s *Service) MessagesSendMultiMedia(ctx context.Context, request *tg.TLMessagesSendMultiMedia) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.sendMultiMedia - metadata: {}, request: {%v}", request)

	r, err := c.MessagesSendMultiMedia(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.sendMultiMedia - reply: {%v}", r)
	return r, err
}

// MessagesUpdatePinnedMessage
// messages.updatePinnedMessage#d2aaf7ec flags:# silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer:InputPeer id:int = Updates;
func (s *Service) MessagesUpdatePinnedMessage(ctx context.Context, request *tg.TLMessagesUpdatePinnedMessage) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.updatePinnedMessage - metadata: {}, request: {%v}", request)

	r, err := c.MessagesUpdatePinnedMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.updatePinnedMessage - reply: {%v}", r)
	return r, err
}

// MessagesGetSearchCounters
// messages.getSearchCounters#1bbcf300 flags:# peer:InputPeer saved_peer_id:flags.2?InputPeer top_msg_id:flags.0?int filters:Vector<MessagesFilter> = Vector<messages.SearchCounter>;
func (s *Service) MessagesGetSearchCounters(ctx context.Context, request *tg.TLMessagesGetSearchCounters) (*tg.VectorMessagesSearchCounter, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getSearchCounters - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetSearchCounters(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getSearchCounters - reply: {%v}", r)
	return r, err
}

// MessagesUnpinAllMessages
// messages.unpinAllMessages#ee22b9a8 flags:# peer:InputPeer top_msg_id:flags.0?int = messages.AffectedHistory;
func (s *Service) MessagesUnpinAllMessages(ctx context.Context, request *tg.TLMessagesUnpinAllMessages) (*tg.MessagesAffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.unpinAllMessages - metadata: {}, request: {%v}", request)

	r, err := c.MessagesUnpinAllMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.unpinAllMessages - reply: {%v}", r)
	return r, err
}

// MessagesGetSearchResultsCalendar
// messages.getSearchResultsCalendar#6aa3f6bd flags:# peer:InputPeer saved_peer_id:flags.2?InputPeer filter:MessagesFilter offset_id:int offset_date:int = messages.SearchResultsCalendar;
func (s *Service) MessagesGetSearchResultsCalendar(ctx context.Context, request *tg.TLMessagesGetSearchResultsCalendar) (*tg.MessagesSearchResultsCalendar, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getSearchResultsCalendar - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetSearchResultsCalendar(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getSearchResultsCalendar - reply: {%v}", r)
	return r, err
}

// MessagesGetSearchResultsPositions
// messages.getSearchResultsPositions#9c7f2f10 flags:# peer:InputPeer saved_peer_id:flags.2?InputPeer filter:MessagesFilter offset_id:int limit:int = messages.SearchResultsPositions;
func (s *Service) MessagesGetSearchResultsPositions(ctx context.Context, request *tg.TLMessagesGetSearchResultsPositions) (*tg.MessagesSearchResultsPositions, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getSearchResultsPositions - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetSearchResultsPositions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getSearchResultsPositions - reply: {%v}", r)
	return r, err
}

// MessagesToggleNoForwards
// messages.toggleNoForwards#b11eafa2 peer:InputPeer enabled:Bool = Updates;
func (s *Service) MessagesToggleNoForwards(ctx context.Context, request *tg.TLMessagesToggleNoForwards) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.toggleNoForwards - metadata: {}, request: {%v}", request)

	r, err := c.MessagesToggleNoForwards(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.toggleNoForwards - reply: {%v}", r)
	return r, err
}

// MessagesSaveDefaultSendAs
// messages.saveDefaultSendAs#ccfddf96 peer:InputPeer send_as:InputPeer = Bool;
func (s *Service) MessagesSaveDefaultSendAs(ctx context.Context, request *tg.TLMessagesSaveDefaultSendAs) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.saveDefaultSendAs - metadata: {}, request: {%v}", request)

	r, err := c.MessagesSaveDefaultSendAs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.saveDefaultSendAs - reply: {%v}", r)
	return r, err
}

// MessagesSearchSentMedia
// messages.searchSentMedia#107e31a0 q:string filter:MessagesFilter limit:int = messages.Messages;
func (s *Service) MessagesSearchSentMedia(ctx context.Context, request *tg.TLMessagesSearchSentMedia) (*tg.MessagesMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.searchSentMedia - metadata: {}, request: {%v}", request)

	r, err := c.MessagesSearchSentMedia(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.searchSentMedia - reply: {%v}", r)
	return r, err
}

// MessagesGetOutboxReadDate
// messages.getOutboxReadDate#8c4bfe5d peer:InputPeer msg_id:int = OutboxReadDate;
func (s *Service) MessagesGetOutboxReadDate(ctx context.Context, request *tg.TLMessagesGetOutboxReadDate) (*tg.OutboxReadDate, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getOutboxReadDate - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetOutboxReadDate(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getOutboxReadDate - reply: {%v}", r)
	return r, err
}

// ChannelsGetSendAs
// channels.getSendAs#dc770ee peer:InputPeer = channels.SendAsPeers;
func (s *Service) ChannelsGetSendAs(ctx context.Context, request *tg.TLChannelsGetSendAs) (*tg.ChannelsSendAsPeers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.getSendAs - metadata: {}, request: {%v}", request)

	r, err := c.ChannelsGetSendAs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.getSendAs - reply: {%v}", r)
	return r, err
}

// ChannelsSearchPosts
// channels.searchPosts#d19f987b hashtag:string offset_rate:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
func (s *Service) ChannelsSearchPosts(ctx context.Context, request *tg.TLChannelsSearchPosts) (*tg.MessagesMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.searchPosts - metadata: {}, request: {%v}", request)

	r, err := c.ChannelsSearchPosts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.searchPosts - reply: {%v}", r)
	return r, err
}
