/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package messagesclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type MessagesClient interface {
	MessagesGetMessages(ctx context.Context, in *mtproto.TLMessagesGetMessages) (*mtproto.Messages_Messages, error)
	MessagesGetHistory(ctx context.Context, in *mtproto.TLMessagesGetHistory) (*mtproto.Messages_Messages, error)
	MessagesSearch(ctx context.Context, in *mtproto.TLMessagesSearch) (*mtproto.Messages_Messages, error)
	MessagesReadHistory(ctx context.Context, in *mtproto.TLMessagesReadHistory) (*mtproto.Messages_AffectedMessages, error)
	MessagesDeleteHistory(ctx context.Context, in *mtproto.TLMessagesDeleteHistory) (*mtproto.Messages_AffectedHistory, error)
	MessagesDeleteMessages(ctx context.Context, in *mtproto.TLMessagesDeleteMessages) (*mtproto.Messages_AffectedMessages, error)
	MessagesReceivedMessages(ctx context.Context, in *mtproto.TLMessagesReceivedMessages) (*mtproto.Vector_ReceivedNotifyMessage, error)
	MessagesSendMessage(ctx context.Context, in *mtproto.TLMessagesSendMessage) (*mtproto.Updates, error)
	MessagesSendMedia(ctx context.Context, in *mtproto.TLMessagesSendMedia) (*mtproto.Updates, error)
	MessagesForwardMessages(ctx context.Context, in *mtproto.TLMessagesForwardMessages) (*mtproto.Updates, error)
	MessagesReadMessageContents(ctx context.Context, in *mtproto.TLMessagesReadMessageContents) (*mtproto.Messages_AffectedMessages, error)
	MessagesGetMessagesViews(ctx context.Context, in *mtproto.TLMessagesGetMessagesViews) (*mtproto.Messages_MessageViews, error)
	MessagesSearchGlobal(ctx context.Context, in *mtproto.TLMessagesSearchGlobal) (*mtproto.Messages_Messages, error)
	MessagesGetMessageEditData(ctx context.Context, in *mtproto.TLMessagesGetMessageEditData) (*mtproto.Messages_MessageEditData, error)
	MessagesEditMessage(ctx context.Context, in *mtproto.TLMessagesEditMessage) (*mtproto.Updates, error)
	MessagesGetUnreadMentions(ctx context.Context, in *mtproto.TLMessagesGetUnreadMentions) (*mtproto.Messages_Messages, error)
	MessagesReadMentions(ctx context.Context, in *mtproto.TLMessagesReadMentions) (*mtproto.Messages_AffectedHistory, error)
	MessagesGetRecentLocations(ctx context.Context, in *mtproto.TLMessagesGetRecentLocations) (*mtproto.Messages_Messages, error)
	MessagesSendMultiMedia(ctx context.Context, in *mtproto.TLMessagesSendMultiMedia) (*mtproto.Updates, error)
	MessagesUpdatePinnedMessage(ctx context.Context, in *mtproto.TLMessagesUpdatePinnedMessage) (*mtproto.Updates, error)
	MessagesGetSearchCounters(ctx context.Context, in *mtproto.TLMessagesGetSearchCounters) (*mtproto.Vector_Messages_SearchCounter, error)
	MessagesUnpinAllMessages(ctx context.Context, in *mtproto.TLMessagesUnpinAllMessages) (*mtproto.Messages_AffectedHistory, error)
	MessagesGetSearchResultsCalendar(ctx context.Context, in *mtproto.TLMessagesGetSearchResultsCalendar) (*mtproto.Messages_SearchResultsCalendar, error)
	MessagesGetSearchResultsPositions(ctx context.Context, in *mtproto.TLMessagesGetSearchResultsPositions) (*mtproto.Messages_SearchResultsPositions, error)
	MessagesToggleNoForwards(ctx context.Context, in *mtproto.TLMessagesToggleNoForwards) (*mtproto.Updates, error)
	MessagesSaveDefaultSendAs(ctx context.Context, in *mtproto.TLMessagesSaveDefaultSendAs) (*mtproto.Bool, error)
	MessagesSearchSentMedia(ctx context.Context, in *mtproto.TLMessagesSearchSentMedia) (*mtproto.Messages_Messages, error)
	MessagesGetOutboxReadDate(ctx context.Context, in *mtproto.TLMessagesGetOutboxReadDate) (*mtproto.OutboxReadDate, error)
	ChannelsGetSendAs(ctx context.Context, in *mtproto.TLChannelsGetSendAs) (*mtproto.Channels_SendAsPeers, error)
	ChannelsSearchPosts(ctx context.Context, in *mtproto.TLChannelsSearchPosts) (*mtproto.Messages_Messages, error)
}

type defaultMessagesClient struct {
	cli zrpc.Client
}

func NewMessagesClient(cli zrpc.Client) MessagesClient {
	return &defaultMessagesClient{
		cli: cli,
	}
}

// MessagesGetMessages
// messages.getMessages#63c66506 id:Vector<InputMessage> = messages.Messages;
func (m *defaultMessagesClient) MessagesGetMessages(ctx context.Context, in *mtproto.TLMessagesGetMessages) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesGetMessages(ctx, in)
}

// MessagesGetHistory
// messages.getHistory#4423e6c5 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (m *defaultMessagesClient) MessagesGetHistory(ctx context.Context, in *mtproto.TLMessagesGetHistory) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesGetHistory(ctx, in)
}

// MessagesSearch
// messages.search#29ee847a flags:# peer:InputPeer q:string from_id:flags.0?InputPeer saved_peer_id:flags.2?InputPeer saved_reaction:flags.3?Vector<Reaction> top_msg_id:flags.1?int filter:MessagesFilter min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (m *defaultMessagesClient) MessagesSearch(ctx context.Context, in *mtproto.TLMessagesSearch) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesSearch(ctx, in)
}

// MessagesReadHistory
// messages.readHistory#e306d3a peer:InputPeer max_id:int = messages.AffectedMessages;
func (m *defaultMessagesClient) MessagesReadHistory(ctx context.Context, in *mtproto.TLMessagesReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesReadHistory(ctx, in)
}

// MessagesDeleteHistory
// messages.deleteHistory#b08f922a flags:# just_clear:flags.0?true revoke:flags.1?true peer:InputPeer max_id:int min_date:flags.2?int max_date:flags.3?int = messages.AffectedHistory;
func (m *defaultMessagesClient) MessagesDeleteHistory(ctx context.Context, in *mtproto.TLMessagesDeleteHistory) (*mtproto.Messages_AffectedHistory, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesDeleteHistory(ctx, in)
}

// MessagesDeleteMessages
// messages.deleteMessages#e58e95d2 flags:# revoke:flags.0?true id:Vector<int> = messages.AffectedMessages;
func (m *defaultMessagesClient) MessagesDeleteMessages(ctx context.Context, in *mtproto.TLMessagesDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesDeleteMessages(ctx, in)
}

// MessagesReceivedMessages
// messages.receivedMessages#5a954c0 max_id:int = Vector<ReceivedNotifyMessage>;
func (m *defaultMessagesClient) MessagesReceivedMessages(ctx context.Context, in *mtproto.TLMessagesReceivedMessages) (*mtproto.Vector_ReceivedNotifyMessage, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesReceivedMessages(ctx, in)
}

// MessagesSendMessage
// messages.sendMessage#983f9745 flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true update_stickersets_order:flags.15?true invert_media:flags.16?true peer:InputPeer reply_to:flags.0?InputReplyTo message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut effect:flags.18?long = Updates;
func (m *defaultMessagesClient) MessagesSendMessage(ctx context.Context, in *mtproto.TLMessagesSendMessage) (*mtproto.Updates, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesSendMessage(ctx, in)
}

// MessagesSendMedia
// messages.sendMedia#7852834e flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true update_stickersets_order:flags.15?true invert_media:flags.16?true peer:InputPeer reply_to:flags.0?InputReplyTo media:InputMedia message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut effect:flags.18?long = Updates;
func (m *defaultMessagesClient) MessagesSendMedia(ctx context.Context, in *mtproto.TLMessagesSendMedia) (*mtproto.Updates, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesSendMedia(ctx, in)
}

// MessagesForwardMessages
// messages.forwardMessages#d5039208 flags:# silent:flags.5?true background:flags.6?true with_my_score:flags.8?true drop_author:flags.11?true drop_media_captions:flags.12?true noforwards:flags.14?true from_peer:InputPeer id:Vector<int> random_id:Vector<long> to_peer:InputPeer top_msg_id:flags.9?int schedule_date:flags.10?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut = Updates;
func (m *defaultMessagesClient) MessagesForwardMessages(ctx context.Context, in *mtproto.TLMessagesForwardMessages) (*mtproto.Updates, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesForwardMessages(ctx, in)
}

// MessagesReadMessageContents
// messages.readMessageContents#36a73f77 id:Vector<int> = messages.AffectedMessages;
func (m *defaultMessagesClient) MessagesReadMessageContents(ctx context.Context, in *mtproto.TLMessagesReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesReadMessageContents(ctx, in)
}

// MessagesGetMessagesViews
// messages.getMessagesViews#5784d3e1 peer:InputPeer id:Vector<int> increment:Bool = messages.MessageViews;
func (m *defaultMessagesClient) MessagesGetMessagesViews(ctx context.Context, in *mtproto.TLMessagesGetMessagesViews) (*mtproto.Messages_MessageViews, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesGetMessagesViews(ctx, in)
}

// MessagesSearchGlobal
// messages.searchGlobal#4bc6589a flags:# broadcasts_only:flags.1?true folder_id:flags.0?int q:string filter:MessagesFilter min_date:int max_date:int offset_rate:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
func (m *defaultMessagesClient) MessagesSearchGlobal(ctx context.Context, in *mtproto.TLMessagesSearchGlobal) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesSearchGlobal(ctx, in)
}

// MessagesGetMessageEditData
// messages.getMessageEditData#fda68d36 peer:InputPeer id:int = messages.MessageEditData;
func (m *defaultMessagesClient) MessagesGetMessageEditData(ctx context.Context, in *mtproto.TLMessagesGetMessageEditData) (*mtproto.Messages_MessageEditData, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesGetMessageEditData(ctx, in)
}

// MessagesEditMessage
// messages.editMessage#dfd14005 flags:# no_webpage:flags.1?true invert_media:flags.16?true peer:InputPeer id:int message:flags.11?string media:flags.14?InputMedia reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.15?int quick_reply_shortcut_id:flags.17?int = Updates;
func (m *defaultMessagesClient) MessagesEditMessage(ctx context.Context, in *mtproto.TLMessagesEditMessage) (*mtproto.Updates, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesEditMessage(ctx, in)
}

// MessagesGetUnreadMentions
// messages.getUnreadMentions#f107e790 flags:# peer:InputPeer top_msg_id:flags.0?int offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (m *defaultMessagesClient) MessagesGetUnreadMentions(ctx context.Context, in *mtproto.TLMessagesGetUnreadMentions) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesGetUnreadMentions(ctx, in)
}

// MessagesReadMentions
// messages.readMentions#36e5bf4d flags:# peer:InputPeer top_msg_id:flags.0?int = messages.AffectedHistory;
func (m *defaultMessagesClient) MessagesReadMentions(ctx context.Context, in *mtproto.TLMessagesReadMentions) (*mtproto.Messages_AffectedHistory, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesReadMentions(ctx, in)
}

// MessagesGetRecentLocations
// messages.getRecentLocations#702a40e0 peer:InputPeer limit:int hash:long = messages.Messages;
func (m *defaultMessagesClient) MessagesGetRecentLocations(ctx context.Context, in *mtproto.TLMessagesGetRecentLocations) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesGetRecentLocations(ctx, in)
}

// MessagesSendMultiMedia
// messages.sendMultiMedia#37b74355 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true update_stickersets_order:flags.15?true invert_media:flags.16?true peer:InputPeer reply_to:flags.0?InputReplyTo multi_media:Vector<InputSingleMedia> schedule_date:flags.10?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut effect:flags.18?long = Updates;
func (m *defaultMessagesClient) MessagesSendMultiMedia(ctx context.Context, in *mtproto.TLMessagesSendMultiMedia) (*mtproto.Updates, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesSendMultiMedia(ctx, in)
}

// MessagesUpdatePinnedMessage
// messages.updatePinnedMessage#d2aaf7ec flags:# silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer:InputPeer id:int = Updates;
func (m *defaultMessagesClient) MessagesUpdatePinnedMessage(ctx context.Context, in *mtproto.TLMessagesUpdatePinnedMessage) (*mtproto.Updates, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesUpdatePinnedMessage(ctx, in)
}

// MessagesGetSearchCounters
// messages.getSearchCounters#1bbcf300 flags:# peer:InputPeer saved_peer_id:flags.2?InputPeer top_msg_id:flags.0?int filters:Vector<MessagesFilter> = Vector<messages.SearchCounter>;
func (m *defaultMessagesClient) MessagesGetSearchCounters(ctx context.Context, in *mtproto.TLMessagesGetSearchCounters) (*mtproto.Vector_Messages_SearchCounter, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesGetSearchCounters(ctx, in)
}

// MessagesUnpinAllMessages
// messages.unpinAllMessages#ee22b9a8 flags:# peer:InputPeer top_msg_id:flags.0?int = messages.AffectedHistory;
func (m *defaultMessagesClient) MessagesUnpinAllMessages(ctx context.Context, in *mtproto.TLMessagesUnpinAllMessages) (*mtproto.Messages_AffectedHistory, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesUnpinAllMessages(ctx, in)
}

// MessagesGetSearchResultsCalendar
// messages.getSearchResultsCalendar#6aa3f6bd flags:# peer:InputPeer saved_peer_id:flags.2?InputPeer filter:MessagesFilter offset_id:int offset_date:int = messages.SearchResultsCalendar;
func (m *defaultMessagesClient) MessagesGetSearchResultsCalendar(ctx context.Context, in *mtproto.TLMessagesGetSearchResultsCalendar) (*mtproto.Messages_SearchResultsCalendar, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesGetSearchResultsCalendar(ctx, in)
}

// MessagesGetSearchResultsPositions
// messages.getSearchResultsPositions#9c7f2f10 flags:# peer:InputPeer saved_peer_id:flags.2?InputPeer filter:MessagesFilter offset_id:int limit:int = messages.SearchResultsPositions;
func (m *defaultMessagesClient) MessagesGetSearchResultsPositions(ctx context.Context, in *mtproto.TLMessagesGetSearchResultsPositions) (*mtproto.Messages_SearchResultsPositions, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesGetSearchResultsPositions(ctx, in)
}

// MessagesToggleNoForwards
// messages.toggleNoForwards#b11eafa2 peer:InputPeer enabled:Bool = Updates;
func (m *defaultMessagesClient) MessagesToggleNoForwards(ctx context.Context, in *mtproto.TLMessagesToggleNoForwards) (*mtproto.Updates, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesToggleNoForwards(ctx, in)
}

// MessagesSaveDefaultSendAs
// messages.saveDefaultSendAs#ccfddf96 peer:InputPeer send_as:InputPeer = Bool;
func (m *defaultMessagesClient) MessagesSaveDefaultSendAs(ctx context.Context, in *mtproto.TLMessagesSaveDefaultSendAs) (*mtproto.Bool, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesSaveDefaultSendAs(ctx, in)
}

// MessagesSearchSentMedia
// messages.searchSentMedia#107e31a0 q:string filter:MessagesFilter limit:int = messages.Messages;
func (m *defaultMessagesClient) MessagesSearchSentMedia(ctx context.Context, in *mtproto.TLMessagesSearchSentMedia) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesSearchSentMedia(ctx, in)
}

// MessagesGetOutboxReadDate
// messages.getOutboxReadDate#8c4bfe5d peer:InputPeer msg_id:int = OutboxReadDate;
func (m *defaultMessagesClient) MessagesGetOutboxReadDate(ctx context.Context, in *mtproto.TLMessagesGetOutboxReadDate) (*mtproto.OutboxReadDate, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.MessagesGetOutboxReadDate(ctx, in)
}

// ChannelsGetSendAs
// channels.getSendAs#dc770ee peer:InputPeer = channels.SendAsPeers;
func (m *defaultMessagesClient) ChannelsGetSendAs(ctx context.Context, in *mtproto.TLChannelsGetSendAs) (*mtproto.Channels_SendAsPeers, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.ChannelsGetSendAs(ctx, in)
}

// ChannelsSearchPosts
// channels.searchPosts#d19f987b hashtag:string offset_rate:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
func (m *defaultMessagesClient) ChannelsSearchPosts(ctx context.Context, in *mtproto.TLChannelsSearchPosts) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCMessagesClient(m.cli.Conn())
	return client.ChannelsSearchPosts(ctx, in)
}
