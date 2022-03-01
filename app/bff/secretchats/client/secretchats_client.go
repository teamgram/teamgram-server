/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package secretchats_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type SecretChatsClient interface {
	MessagesGetDhConfig(ctx context.Context, in *mtproto.TLMessagesGetDhConfig) (*mtproto.Messages_DhConfig, error)
	MessagesRequestEncryption(ctx context.Context, in *mtproto.TLMessagesRequestEncryption) (*mtproto.EncryptedChat, error)
	MessagesAcceptEncryption(ctx context.Context, in *mtproto.TLMessagesAcceptEncryption) (*mtproto.EncryptedChat, error)
	MessagesDiscardEncryption(ctx context.Context, in *mtproto.TLMessagesDiscardEncryption) (*mtproto.Bool, error)
	MessagesSetEncryptedTyping(ctx context.Context, in *mtproto.TLMessagesSetEncryptedTyping) (*mtproto.Bool, error)
	MessagesReadEncryptedHistory(ctx context.Context, in *mtproto.TLMessagesReadEncryptedHistory) (*mtproto.Bool, error)
	MessagesSendEncrypted(ctx context.Context, in *mtproto.TLMessagesSendEncrypted) (*mtproto.Messages_SentEncryptedMessage, error)
	MessagesSendEncryptedFile(ctx context.Context, in *mtproto.TLMessagesSendEncryptedFile) (*mtproto.Messages_SentEncryptedMessage, error)
	MessagesSendEncryptedService(ctx context.Context, in *mtproto.TLMessagesSendEncryptedService) (*mtproto.Messages_SentEncryptedMessage, error)
	MessagesReceivedQueue(ctx context.Context, in *mtproto.TLMessagesReceivedQueue) (*mtproto.Vector_Long, error)
}

type defaultSecretChatsClient struct {
	cli zrpc.Client
}

func NewSecretChatsClient(cli zrpc.Client) SecretChatsClient {
	return &defaultSecretChatsClient{
		cli: cli,
	}
}

// MessagesGetDhConfig
// messages.getDhConfig#26cf8950 version:int random_length:int = messages.DhConfig;
func (m *defaultSecretChatsClient) MessagesGetDhConfig(ctx context.Context, in *mtproto.TLMessagesGetDhConfig) (*mtproto.Messages_DhConfig, error) {
	client := mtproto.NewRPCSecretChatsClient(m.cli.Conn())
	return client.MessagesGetDhConfig(ctx, in)
}

// MessagesRequestEncryption
// messages.requestEncryption#f64daf43 user_id:InputUser random_id:int g_a:bytes = EncryptedChat;
func (m *defaultSecretChatsClient) MessagesRequestEncryption(ctx context.Context, in *mtproto.TLMessagesRequestEncryption) (*mtproto.EncryptedChat, error) {
	client := mtproto.NewRPCSecretChatsClient(m.cli.Conn())
	return client.MessagesRequestEncryption(ctx, in)
}

// MessagesAcceptEncryption
// messages.acceptEncryption#3dbc0415 peer:InputEncryptedChat g_b:bytes key_fingerprint:long = EncryptedChat;
func (m *defaultSecretChatsClient) MessagesAcceptEncryption(ctx context.Context, in *mtproto.TLMessagesAcceptEncryption) (*mtproto.EncryptedChat, error) {
	client := mtproto.NewRPCSecretChatsClient(m.cli.Conn())
	return client.MessagesAcceptEncryption(ctx, in)
}

// MessagesDiscardEncryption
// messages.discardEncryption#f393aea0 flags:# delete_history:flags.0?true chat_id:int = Bool;
func (m *defaultSecretChatsClient) MessagesDiscardEncryption(ctx context.Context, in *mtproto.TLMessagesDiscardEncryption) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSecretChatsClient(m.cli.Conn())
	return client.MessagesDiscardEncryption(ctx, in)
}

// MessagesSetEncryptedTyping
// messages.setEncryptedTyping#791451ed peer:InputEncryptedChat typing:Bool = Bool;
func (m *defaultSecretChatsClient) MessagesSetEncryptedTyping(ctx context.Context, in *mtproto.TLMessagesSetEncryptedTyping) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSecretChatsClient(m.cli.Conn())
	return client.MessagesSetEncryptedTyping(ctx, in)
}

// MessagesReadEncryptedHistory
// messages.readEncryptedHistory#7f4b690a peer:InputEncryptedChat max_date:int = Bool;
func (m *defaultSecretChatsClient) MessagesReadEncryptedHistory(ctx context.Context, in *mtproto.TLMessagesReadEncryptedHistory) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSecretChatsClient(m.cli.Conn())
	return client.MessagesReadEncryptedHistory(ctx, in)
}

// MessagesSendEncrypted
// messages.sendEncrypted#44fa7a15 flags:# silent:flags.0?true peer:InputEncryptedChat random_id:long data:bytes = messages.SentEncryptedMessage;
func (m *defaultSecretChatsClient) MessagesSendEncrypted(ctx context.Context, in *mtproto.TLMessagesSendEncrypted) (*mtproto.Messages_SentEncryptedMessage, error) {
	client := mtproto.NewRPCSecretChatsClient(m.cli.Conn())
	return client.MessagesSendEncrypted(ctx, in)
}

// MessagesSendEncryptedFile
// messages.sendEncryptedFile#5559481d flags:# silent:flags.0?true peer:InputEncryptedChat random_id:long data:bytes file:InputEncryptedFile = messages.SentEncryptedMessage;
func (m *defaultSecretChatsClient) MessagesSendEncryptedFile(ctx context.Context, in *mtproto.TLMessagesSendEncryptedFile) (*mtproto.Messages_SentEncryptedMessage, error) {
	client := mtproto.NewRPCSecretChatsClient(m.cli.Conn())
	return client.MessagesSendEncryptedFile(ctx, in)
}

// MessagesSendEncryptedService
// messages.sendEncryptedService#32d439a4 peer:InputEncryptedChat random_id:long data:bytes = messages.SentEncryptedMessage;
func (m *defaultSecretChatsClient) MessagesSendEncryptedService(ctx context.Context, in *mtproto.TLMessagesSendEncryptedService) (*mtproto.Messages_SentEncryptedMessage, error) {
	client := mtproto.NewRPCSecretChatsClient(m.cli.Conn())
	return client.MessagesSendEncryptedService(ctx, in)
}

// MessagesReceivedQueue
// messages.receivedQueue#55a5bb66 max_qts:int = Vector<long>;
func (m *defaultSecretChatsClient) MessagesReceivedQueue(ctx context.Context, in *mtproto.TLMessagesReceivedQueue) (*mtproto.Vector_Long, error) {
	client := mtproto.NewRPCSecretChatsClient(m.cli.Conn())
	return client.MessagesReceivedQueue(ctx, in)
}
