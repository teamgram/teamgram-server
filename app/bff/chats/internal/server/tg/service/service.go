/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgooo Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"
	"github.com/teamgram/teamgram-server/v2/app/bff/chats/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type Service struct {
	svcCtx *svc.ServiceContext
}

// MessagesEditChatCreator implements [tg.RPCChats].
func (s *Service) MessagesEditChatCreator(ctx context.Context, in *tg.TLMessagesEditChatCreator) (*tg.Updates, error) {
	panic("unimplemented")
}

// MessagesEditChatParticipantRank implements [tg.RPCChats].
func (s *Service) MessagesEditChatParticipantRank(ctx context.Context, in *tg.TLMessagesEditChatParticipantRank) (*tg.Updates, error) {
	panic("unimplemented")
}

// MessagesGetFutureChatCreatorAfterLeave implements [tg.RPCChats].
func (s *Service) MessagesGetFutureChatCreatorAfterLeave(ctx context.Context, in *tg.TLMessagesGetFutureChatCreatorAfterLeave) (*tg.User, error) {
	panic("unimplemented")
}

func (s *Service) GetServiceContext() *svc.ServiceContext {
	return s.svcCtx
}

func New(ctx *svc.ServiceContext) *Service {
	return &Service{
		svcCtx: ctx,
	}
}
