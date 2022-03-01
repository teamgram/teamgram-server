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
	"github.com/teamgram/teamgram-server/app/bff/secretchats/internal/core"
)

// MessagesGetDhConfig
// messages.getDhConfig#26cf8950 version:int random_length:int = messages.DhConfig;
func (s *Service) MessagesGetDhConfig(ctx context.Context, request *mtproto.TLMessagesGetDhConfig) (*mtproto.Messages_DhConfig, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getDhConfig - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetDhConfig(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getDhConfig - reply: %s", r.DebugString())
	return r, err
}

// MessagesRequestEncryption
// messages.requestEncryption#f64daf43 user_id:InputUser random_id:int g_a:bytes = EncryptedChat;
func (s *Service) MessagesRequestEncryption(ctx context.Context, request *mtproto.TLMessagesRequestEncryption) (*mtproto.EncryptedChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.requestEncryption - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesRequestEncryption(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.requestEncryption - reply: %s", r.DebugString())
	return r, err
}

// MessagesAcceptEncryption
// messages.acceptEncryption#3dbc0415 peer:InputEncryptedChat g_b:bytes key_fingerprint:long = EncryptedChat;
func (s *Service) MessagesAcceptEncryption(ctx context.Context, request *mtproto.TLMessagesAcceptEncryption) (*mtproto.EncryptedChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.acceptEncryption - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesAcceptEncryption(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.acceptEncryption - reply: %s", r.DebugString())
	return r, err
}

// MessagesDiscardEncryption
// messages.discardEncryption#f393aea0 flags:# delete_history:flags.0?true chat_id:int = Bool;
func (s *Service) MessagesDiscardEncryption(ctx context.Context, request *mtproto.TLMessagesDiscardEncryption) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.discardEncryption - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesDiscardEncryption(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.discardEncryption - reply: %s", r.DebugString())
	return r, err
}

// MessagesSetEncryptedTyping
// messages.setEncryptedTyping#791451ed peer:InputEncryptedChat typing:Bool = Bool;
func (s *Service) MessagesSetEncryptedTyping(ctx context.Context, request *mtproto.TLMessagesSetEncryptedTyping) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.setEncryptedTyping - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSetEncryptedTyping(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.setEncryptedTyping - reply: %s", r.DebugString())
	return r, err
}

// MessagesReadEncryptedHistory
// messages.readEncryptedHistory#7f4b690a peer:InputEncryptedChat max_date:int = Bool;
func (s *Service) MessagesReadEncryptedHistory(ctx context.Context, request *mtproto.TLMessagesReadEncryptedHistory) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.readEncryptedHistory - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReadEncryptedHistory(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.readEncryptedHistory - reply: %s", r.DebugString())
	return r, err
}

// MessagesSendEncrypted
// messages.sendEncrypted#44fa7a15 flags:# silent:flags.0?true peer:InputEncryptedChat random_id:long data:bytes = messages.SentEncryptedMessage;
func (s *Service) MessagesSendEncrypted(ctx context.Context, request *mtproto.TLMessagesSendEncrypted) (*mtproto.Messages_SentEncryptedMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.sendEncrypted - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSendEncrypted(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.sendEncrypted - reply: %s", r.DebugString())
	return r, err
}

// MessagesSendEncryptedFile
// messages.sendEncryptedFile#5559481d flags:# silent:flags.0?true peer:InputEncryptedChat random_id:long data:bytes file:InputEncryptedFile = messages.SentEncryptedMessage;
func (s *Service) MessagesSendEncryptedFile(ctx context.Context, request *mtproto.TLMessagesSendEncryptedFile) (*mtproto.Messages_SentEncryptedMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.sendEncryptedFile - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSendEncryptedFile(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.sendEncryptedFile - reply: %s", r.DebugString())
	return r, err
}

// MessagesSendEncryptedService
// messages.sendEncryptedService#32d439a4 peer:InputEncryptedChat random_id:long data:bytes = messages.SentEncryptedMessage;
func (s *Service) MessagesSendEncryptedService(ctx context.Context, request *mtproto.TLMessagesSendEncryptedService) (*mtproto.Messages_SentEncryptedMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.sendEncryptedService - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSendEncryptedService(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.sendEncryptedService - reply: %s", r.DebugString())
	return r, err
}

// MessagesReceivedQueue
// messages.receivedQueue#55a5bb66 max_qts:int = Vector<long>;
func (s *Service) MessagesReceivedQueue(ctx context.Context, request *mtproto.TLMessagesReceivedQueue) (*mtproto.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.receivedQueue - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReceivedQueue(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.receivedQueue - reply: %s", r.DebugString())
	return r, err
}
