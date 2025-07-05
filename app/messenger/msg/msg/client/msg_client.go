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

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg/msgservice"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type MsgClient interface {
	MsgPushUserMessage(ctx context.Context, in *msg.TLMsgPushUserMessage) (*tg.Bool, error)
	MsgReadMessageContents(ctx context.Context, in *msg.TLMsgReadMessageContents) (*tg.MessagesAffectedMessages, error)
	MsgSendMessageV2(ctx context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error)
	MsgEditMessageV2(ctx context.Context, in *msg.TLMsgEditMessageV2) (*tg.Updates, error)
	MsgDeleteMessages(ctx context.Context, in *msg.TLMsgDeleteMessages) (*tg.MessagesAffectedMessages, error)
	MsgDeleteHistory(ctx context.Context, in *msg.TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error)
	MsgDeletePhoneCallHistory(ctx context.Context, in *msg.TLMsgDeletePhoneCallHistory) (*tg.MessagesAffectedFoundMessages, error)
	MsgDeleteChatHistory(ctx context.Context, in *msg.TLMsgDeleteChatHistory) (*tg.Bool, error)
	MsgReadHistory(ctx context.Context, in *msg.TLMsgReadHistory) (*tg.MessagesAffectedMessages, error)
	MsgReadHistoryV2(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error)
	MsgUpdatePinnedMessage(ctx context.Context, in *msg.TLMsgUpdatePinnedMessage) (*tg.Updates, error)
	MsgUnpinAllMessages(ctx context.Context, in *msg.TLMsgUnpinAllMessages) (*tg.MessagesAffectedHistory, error)
}

type defaultMsgClient struct {
	cli client.Client
}

func NewMsgClient(cli client.Client) MsgClient {
	return &defaultMsgClient{
		cli: cli,
	}
}

// MsgPushUserMessage
// msg.pushUserMessage user_id:long auth_key_id:long peer_type:int peer_id:long push_type:int message:OutboxMessage = Bool;
func (m *defaultMsgClient) MsgPushUserMessage(ctx context.Context, in *msg.TLMsgPushUserMessage) (*tg.Bool, error) {
	cli := msgservice.NewRPCMsgClient(m.cli)
	return cli.MsgPushUserMessage(ctx, in)
}

// MsgReadMessageContents
// msg.readMessageContents user_id:long auth_key_id:long peer_type:int peer_id:long id:Vector<ContentMessage> = messages.AffectedMessages;
func (m *defaultMsgClient) MsgReadMessageContents(ctx context.Context, in *msg.TLMsgReadMessageContents) (*tg.MessagesAffectedMessages, error) {
	cli := msgservice.NewRPCMsgClient(m.cli)
	return cli.MsgReadMessageContents(ctx, in)
}

// MsgSendMessageV2
// msg.sendMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
func (m *defaultMsgClient) MsgSendMessageV2(ctx context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
	cli := msgservice.NewRPCMsgClient(m.cli)
	return cli.MsgSendMessageV2(ctx, in)
}

// MsgEditMessageV2
// msg.editMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long edit_type:int new_message:OutboxMessage dst_message:MessageBox = Updates;
func (m *defaultMsgClient) MsgEditMessageV2(ctx context.Context, in *msg.TLMsgEditMessageV2) (*tg.Updates, error) {
	cli := msgservice.NewRPCMsgClient(m.cli)
	return cli.MsgEditMessageV2(ctx, in)
}

// MsgDeleteMessages
// msg.deleteMessages flags:# user_id:long auth_key_id:long peer_type:int peer_id:long revoke:flags.1?true id:Vector<int> = messages.AffectedMessages;
func (m *defaultMsgClient) MsgDeleteMessages(ctx context.Context, in *msg.TLMsgDeleteMessages) (*tg.MessagesAffectedMessages, error) {
	cli := msgservice.NewRPCMsgClient(m.cli)
	return cli.MsgDeleteMessages(ctx, in)
}

// MsgDeleteHistory
// msg.deleteHistory flags:# user_id:long auth_key_id:long peer_type:int peer_id:long just_clear:flags.0?true revoke:flags.1?true max_id:int = messages.AffectedHistory;
func (m *defaultMsgClient) MsgDeleteHistory(ctx context.Context, in *msg.TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error) {
	cli := msgservice.NewRPCMsgClient(m.cli)
	return cli.MsgDeleteHistory(ctx, in)
}

// MsgDeletePhoneCallHistory
// msg.deletePhoneCallHistory flags:# user_id:long auth_key_id:long revoke:flags.1?true = messages.AffectedFoundMessages;
func (m *defaultMsgClient) MsgDeletePhoneCallHistory(ctx context.Context, in *msg.TLMsgDeletePhoneCallHistory) (*tg.MessagesAffectedFoundMessages, error) {
	cli := msgservice.NewRPCMsgClient(m.cli)
	return cli.MsgDeletePhoneCallHistory(ctx, in)
}

// MsgDeleteChatHistory
// msg.deleteChatHistory chat_id:long delete_user_id:long = Bool;
func (m *defaultMsgClient) MsgDeleteChatHistory(ctx context.Context, in *msg.TLMsgDeleteChatHistory) (*tg.Bool, error) {
	cli := msgservice.NewRPCMsgClient(m.cli)
	return cli.MsgDeleteChatHistory(ctx, in)
}

// MsgReadHistory
// msg.readHistory user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
func (m *defaultMsgClient) MsgReadHistory(ctx context.Context, in *msg.TLMsgReadHistory) (*tg.MessagesAffectedMessages, error) {
	cli := msgservice.NewRPCMsgClient(m.cli)
	return cli.MsgReadHistory(ctx, in)
}

// MsgReadHistoryV2
// msg.readHistoryV2 user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
func (m *defaultMsgClient) MsgReadHistoryV2(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error) {
	cli := msgservice.NewRPCMsgClient(m.cli)
	return cli.MsgReadHistoryV2(ctx, in)
}

// MsgUpdatePinnedMessage
// msg.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;
func (m *defaultMsgClient) MsgUpdatePinnedMessage(ctx context.Context, in *msg.TLMsgUpdatePinnedMessage) (*tg.Updates, error) {
	cli := msgservice.NewRPCMsgClient(m.cli)
	return cli.MsgUpdatePinnedMessage(ctx, in)
}

// MsgUnpinAllMessages
// msg.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = messages.AffectedHistory;
func (m *defaultMsgClient) MsgUnpinAllMessages(ctx context.Context, in *msg.TLMsgUnpinAllMessages) (*tg.MessagesAffectedHistory, error) {
	cli := msgservice.NewRPCMsgClient(m.cli)
	return cli.MsgUnpinAllMessages(ctx, in)
}
