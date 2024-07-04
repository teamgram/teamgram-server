/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package msgclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type MsgClient interface {
	MsgPushUserMessage(ctx context.Context, in *msg.TLMsgPushUserMessage) (*mtproto.Bool, error)
	MsgReadMessageContents(ctx context.Context, in *msg.TLMsgReadMessageContents) (*mtproto.Messages_AffectedMessages, error)
	MsgSendMessageV2(ctx context.Context, in *msg.TLMsgSendMessageV2) (*mtproto.Updates, error)
	MsgEditMessage(ctx context.Context, in *msg.TLMsgEditMessage) (*mtproto.Updates, error)
	MsgEditMessageV2(ctx context.Context, in *msg.TLMsgEditMessageV2) (*mtproto.Updates, error)
	MsgDeleteMessages(ctx context.Context, in *msg.TLMsgDeleteMessages) (*mtproto.Messages_AffectedMessages, error)
	MsgDeleteHistory(ctx context.Context, in *msg.TLMsgDeleteHistory) (*mtproto.Messages_AffectedHistory, error)
	MsgDeletePhoneCallHistory(ctx context.Context, in *msg.TLMsgDeletePhoneCallHistory) (*mtproto.Messages_AffectedFoundMessages, error)
	MsgDeleteChatHistory(ctx context.Context, in *msg.TLMsgDeleteChatHistory) (*mtproto.Bool, error)
	MsgReadHistory(ctx context.Context, in *msg.TLMsgReadHistory) (*mtproto.Messages_AffectedMessages, error)
	MsgReadHistoryV2(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*mtproto.Messages_AffectedMessages, error)
	MsgUpdatePinnedMessage(ctx context.Context, in *msg.TLMsgUpdatePinnedMessage) (*mtproto.Updates, error)
	MsgUnpinAllMessages(ctx context.Context, in *msg.TLMsgUnpinAllMessages) (*mtproto.Messages_AffectedHistory, error)
}

type defaultMsgClient struct {
	cli zrpc.Client
}

func NewMsgClient(cli zrpc.Client) MsgClient {
	return &defaultMsgClient{
		cli: cli,
	}
}

// MsgPushUserMessage
// msg.pushUserMessage user_id:long auth_key_id:long peer_type:int peer_id:long push_type:int message:OutboxMessage = Bool;
func (m *defaultMsgClient) MsgPushUserMessage(ctx context.Context, in *msg.TLMsgPushUserMessage) (*mtproto.Bool, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgPushUserMessage(ctx, in)
}

// MsgReadMessageContents
// msg.readMessageContents user_id:long auth_key_id:long peer_type:int peer_id:long id:Vector<ContentMessage> = messages.AffectedMessages;
func (m *defaultMsgClient) MsgReadMessageContents(ctx context.Context, in *msg.TLMsgReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgReadMessageContents(ctx, in)
}

// MsgSendMessageV2
// msg.sendMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
func (m *defaultMsgClient) MsgSendMessageV2(ctx context.Context, in *msg.TLMsgSendMessageV2) (*mtproto.Updates, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgSendMessageV2(ctx, in)
}

// MsgEditMessage
// msg.editMessage user_id:long auth_key_id:long peer_type:int peer_id:long edit_type:int message:OutboxMessage = Updates;
func (m *defaultMsgClient) MsgEditMessage(ctx context.Context, in *msg.TLMsgEditMessage) (*mtproto.Updates, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgEditMessage(ctx, in)
}

// MsgEditMessageV2
// msg.editMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long edit_type:int new_message:OutboxMessage dst_message:MessageBox = Updates;
func (m *defaultMsgClient) MsgEditMessageV2(ctx context.Context, in *msg.TLMsgEditMessageV2) (*mtproto.Updates, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgEditMessageV2(ctx, in)
}

// MsgDeleteMessages
// msg.deleteMessages flags:# user_id:long auth_key_id:long peer_type:int peer_id:long revoke:flags.1?true id:Vector<int> = messages.AffectedMessages;
func (m *defaultMsgClient) MsgDeleteMessages(ctx context.Context, in *msg.TLMsgDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgDeleteMessages(ctx, in)
}

// MsgDeleteHistory
// msg.deleteHistory flags:# user_id:long auth_key_id:long peer_type:int peer_id:long just_clear:flags.0?true revoke:flags.1?true max_id:int = messages.AffectedHistory;
func (m *defaultMsgClient) MsgDeleteHistory(ctx context.Context, in *msg.TLMsgDeleteHistory) (*mtproto.Messages_AffectedHistory, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgDeleteHistory(ctx, in)
}

// MsgDeletePhoneCallHistory
// msg.deletePhoneCallHistory flags:# user_id:long auth_key_id:long revoke:flags.1?true = messages.AffectedFoundMessages;
func (m *defaultMsgClient) MsgDeletePhoneCallHistory(ctx context.Context, in *msg.TLMsgDeletePhoneCallHistory) (*mtproto.Messages_AffectedFoundMessages, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgDeletePhoneCallHistory(ctx, in)
}

// MsgDeleteChatHistory
// msg.deleteChatHistory chat_id:long delete_user_id:long = Bool;
func (m *defaultMsgClient) MsgDeleteChatHistory(ctx context.Context, in *msg.TLMsgDeleteChatHistory) (*mtproto.Bool, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgDeleteChatHistory(ctx, in)
}

// MsgReadHistory
// msg.readHistory user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
func (m *defaultMsgClient) MsgReadHistory(ctx context.Context, in *msg.TLMsgReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgReadHistory(ctx, in)
}

// MsgReadHistoryV2
// msg.readHistoryV2 user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
func (m *defaultMsgClient) MsgReadHistoryV2(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*mtproto.Messages_AffectedMessages, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgReadHistoryV2(ctx, in)
}

// MsgUpdatePinnedMessage
// msg.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;
func (m *defaultMsgClient) MsgUpdatePinnedMessage(ctx context.Context, in *msg.TLMsgUpdatePinnedMessage) (*mtproto.Updates, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgUpdatePinnedMessage(ctx, in)
}

// MsgUnpinAllMessages
// msg.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = messages.AffectedHistory;
func (m *defaultMsgClient) MsgUnpinAllMessages(ctx context.Context, in *msg.TLMsgUnpinAllMessages) (*mtproto.Messages_AffectedHistory, error) {
	client := msg.NewRPCMsgClient(m.cli.Conn())
	return client.MsgUnpinAllMessages(ctx, in)
}
