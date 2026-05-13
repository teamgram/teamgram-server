/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package msgclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msgservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type MsgClient interface {
	MsgPushUserMessage(ctx context.Context, in *msg.TLMsgPushUserMessage) (*tg.Bool, error)
	MsgReadMessageContents(ctx context.Context, in *msg.TLMsgReadMessageContents) (*tg.MessagesAffectedMessages, error)
	MsgSendMessage(ctx context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error)
	MsgEditMessage(ctx context.Context, in *msg.TLMsgEditMessage) (*tg.Updates, error)
	MsgDeleteMessages(ctx context.Context, in *msg.TLMsgDeleteMessages) (*tg.MessagesAffectedMessages, error)
	MsgDeleteHistory(ctx context.Context, in *msg.TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error)
	MsgDeletePhoneCallHistory(ctx context.Context, in *msg.TLMsgDeletePhoneCallHistory) (*tg.MessagesAffectedFoundMessages, error)
	MsgDeleteChatHistory(ctx context.Context, in *msg.TLMsgDeleteChatHistory) (*tg.Bool, error)
	MsgReadHistory(ctx context.Context, in *msg.TLMsgReadHistory) (*tg.MessagesAffectedMessages, error)
	MsgReadHistoryV2(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error)
	MsgGetHistory(ctx context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error)
	MsgGetUserMessage(ctx context.Context, in *msg.TLMsgGetUserMessage) (*tg.MessageBox, error)
	MsgGetUserMessageList(ctx context.Context, in *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error)
	MsgSearchHashtag(ctx context.Context, in *msg.TLMsgSearchHashtag) (*tg.MessagesMessages, error)
	MsgResolveDialogCursorTopMessage(ctx context.Context, in *msg.TLMsgResolveDialogCursorTopMessage) (*msg.ResolvedDialogCursor, error)
	MsgUpdatePinnedMessage(ctx context.Context, in *msg.TLMsgUpdatePinnedMessage) (*tg.Updates, error)
	MsgUnpinAllMessages(ctx context.Context, in *msg.TLMsgUnpinAllMessages) (*tg.MessagesAffectedHistory, error)
}

type defaultMsgClient struct {
	cli client.Client
	rpc msgservice.Client
}

func NewMsgClient(cli client.Client) MsgClient {
	return &defaultMsgClient{
		cli: cli,
		rpc: msgservice.NewRPCMsgClient(cli),
	}
}

// MsgPushUserMessage
// msg.pushUserMessage user_id:long auth_key_id:long peer_type:int peer_id:long push_type:int message:OutboxMessage = Bool;
func (m *defaultMsgClient) MsgPushUserMessage(ctx context.Context, in *msg.TLMsgPushUserMessage) (*tg.Bool, error) {
	return m.rpc.MsgPushUserMessage(ctx, in)
}

// MsgReadMessageContents
// msg.readMessageContents user_id:long auth_key_id:long peer_type:int peer_id:long id:Vector<ContentMessage> = messages.AffectedMessages;
func (m *defaultMsgClient) MsgReadMessageContents(ctx context.Context, in *msg.TLMsgReadMessageContents) (*tg.MessagesAffectedMessages, error) {
	return m.rpc.MsgReadMessageContents(ctx, in)
}

// MsgSendMessage
// msg.sendMessage flags:# clear_draft:flags.0?true user_id:long auth_key_id:long source_perm_auth_key_id:flags.1?long clear_draft_before_date:flags.2?int attach_facts:flags.3?Vector<UpdateFact> peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
func (m *defaultMsgClient) MsgSendMessage(ctx context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
	return m.rpc.MsgSendMessage(ctx, in)
}

// MsgEditMessage
// msg.editMessage user_id:long auth_key_id:long peer_type:int peer_id:long edit_type:int new_message:OutboxMessage dst_message:MessageBox = Updates;
func (m *defaultMsgClient) MsgEditMessage(ctx context.Context, in *msg.TLMsgEditMessage) (*tg.Updates, error) {
	return m.rpc.MsgEditMessage(ctx, in)
}

// MsgDeleteMessages
// msg.deleteMessages flags:# user_id:long auth_key_id:long peer_type:int peer_id:long revoke:flags.1?true id:Vector<int> = messages.AffectedMessages;
func (m *defaultMsgClient) MsgDeleteMessages(ctx context.Context, in *msg.TLMsgDeleteMessages) (*tg.MessagesAffectedMessages, error) {
	return m.rpc.MsgDeleteMessages(ctx, in)
}

// MsgDeleteHistory
// msg.deleteHistory flags:# user_id:long auth_key_id:long peer_type:int peer_id:long just_clear:flags.0?true revoke:flags.1?true max_id:int = messages.AffectedHistory;
func (m *defaultMsgClient) MsgDeleteHistory(ctx context.Context, in *msg.TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error) {
	return m.rpc.MsgDeleteHistory(ctx, in)
}

// MsgDeletePhoneCallHistory
// msg.deletePhoneCallHistory flags:# user_id:long auth_key_id:long revoke:flags.1?true = messages.AffectedFoundMessages;
func (m *defaultMsgClient) MsgDeletePhoneCallHistory(ctx context.Context, in *msg.TLMsgDeletePhoneCallHistory) (*tg.MessagesAffectedFoundMessages, error) {
	return m.rpc.MsgDeletePhoneCallHistory(ctx, in)
}

// MsgDeleteChatHistory
// msg.deleteChatHistory chat_id:long delete_user_id:long = Bool;
func (m *defaultMsgClient) MsgDeleteChatHistory(ctx context.Context, in *msg.TLMsgDeleteChatHistory) (*tg.Bool, error) {
	return m.rpc.MsgDeleteChatHistory(ctx, in)
}

// MsgReadHistory
// msg.readHistory user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
func (m *defaultMsgClient) MsgReadHistory(ctx context.Context, in *msg.TLMsgReadHistory) (*tg.MessagesAffectedMessages, error) {
	return m.rpc.MsgReadHistory(ctx, in)
}

// MsgReadHistoryV2
// msg.readHistoryV2 user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
func (m *defaultMsgClient) MsgReadHistoryV2(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error) {
	return m.rpc.MsgReadHistoryV2(ctx, in)
}

// MsgGetHistory
// msg.getHistory user_id:long auth_key_id:long peer_type:int peer_id:long offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (m *defaultMsgClient) MsgGetHistory(ctx context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error) {
	return m.rpc.MsgGetHistory(ctx, in)
}

// MsgGetUserMessage
// msg.getUserMessage user_id:long id:int = MessageBox;
func (m *defaultMsgClient) MsgGetUserMessage(ctx context.Context, in *msg.TLMsgGetUserMessage) (*tg.MessageBox, error) {
	return m.rpc.MsgGetUserMessage(ctx, in)
}

// MsgGetUserMessageList
// msg.getUserMessageList user_id:long id_list:Vector<int> = Vector<MessageBox>;
func (m *defaultMsgClient) MsgGetUserMessageList(ctx context.Context, in *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
	return m.rpc.MsgGetUserMessageList(ctx, in)
}

// MsgSearchHashtag
// msg.searchHashtag user_id:long auth_key_id:long peer_type:int peer_id:long hash_tag:string offset_id:int limit:int = messages.Messages;
func (m *defaultMsgClient) MsgSearchHashtag(ctx context.Context, in *msg.TLMsgSearchHashtag) (*tg.MessagesMessages, error) {
	return m.rpc.MsgSearchHashtag(ctx, in)
}

// MsgResolveDialogCursorTopMessage
// msg.resolveDialogCursorTopMessage user_id:long top_message_id:int = ResolvedDialogCursor;
func (m *defaultMsgClient) MsgResolveDialogCursorTopMessage(ctx context.Context, in *msg.TLMsgResolveDialogCursorTopMessage) (*msg.ResolvedDialogCursor, error) {
	return m.rpc.MsgResolveDialogCursorTopMessage(ctx, in)
}

// MsgUpdatePinnedMessage
// msg.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;
func (m *defaultMsgClient) MsgUpdatePinnedMessage(ctx context.Context, in *msg.TLMsgUpdatePinnedMessage) (*tg.Updates, error) {
	return m.rpc.MsgUpdatePinnedMessage(ctx, in)
}

// MsgUnpinAllMessages
// msg.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = messages.AffectedHistory;
func (m *defaultMsgClient) MsgUnpinAllMessages(ctx context.Context, in *msg.TLMsgUnpinAllMessages) (*tg.MessagesAffectedHistory, error) {
	return m.rpc.MsgUnpinAllMessages(ctx, in)
}
