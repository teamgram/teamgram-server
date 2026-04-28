// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/zeromicro/go-zero/core/logx"
)

type chatReadRepository interface {
	GetMutableChat(ctx context.Context, chatID int64, participantIDs ...int64) (*tg.MutableChat, error)
	GetMutableChatByLink(ctx context.Context, link string) (*tg.MutableChat, error)
	GetChatBySelfID(ctx context.Context, chatID, selfID int64) (*tg.MutableChat, error)
	GetChatListByIDList(ctx context.Context, ids []int64) ([]*tg.MutableChat, error)
	GetChatParticipantIDList(ctx context.Context, chatID int64) ([]int64, error)
	GetUsersChatIDList(ctx context.Context, userIDs []int64) ([]repository.UserChatIDList, error)
	GetMyChatList(ctx context.Context, userID int64, isCreator bool) ([]*tg.MutableChat, error)
	Search(ctx context.Context, selfID int64, q string, offset int64, limit int32) ([]*tg.MutableChat, error)
}

type chatWriteRepository interface {
	CreateChat(ctx context.Context, arg repository.CreateChatArg) (*tg.MutableChat, error)
	DeleteChat(ctx context.Context, chatID int64) error
	AddChatUser(ctx context.Context, arg repository.AddChatUserArg) (*tg.ImmutableChatParticipant, error)
	DeleteChatUser(ctx context.Context, arg repository.DeleteChatUserArg) (int64, error)
	MigratedToChannel(ctx context.Context, arg repository.MigratedToChannelArg) error
	UpdateChatTitle(ctx context.Context, chatID int64, title string) (int64, error)
	UpdateChatAbout(ctx context.Context, chatID int64, about string) (int64, error)
	UpdateChatPhoto(ctx context.Context, chatID int64, photoID int64) (int64, error)
	UpdateChatAdmin(ctx context.Context, arg repository.UpdateChatAdminArg) (*tg.ImmutableChatParticipant, int64, error)
	UpdateChatDefaultBannedRights(ctx context.Context, chatID int64, rights tg.ChatBannedRightsClazz) (int64, error)
	UpdateChatNoForwards(ctx context.Context, chatID int64, noforwards bool) (int64, error)
	UpdateChatTTLPeriod(ctx context.Context, chatID int64, ttlPeriod int32) (int64, error)
	UpdateChatAvailableReactions(ctx context.Context, chatID int64, kind int32, reactions []string) (int64, error)
}

type ChatCore struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	readRepo  chatReadRepository
	writeRepo chatWriteRepository
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *ChatCore {
	return &ChatCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func (c *ChatCore) repo() chatReadRepository {
	if c.readRepo != nil {
		return c.readRepo
	}
	return c.svcCtx.Repo
}

func (c *ChatCore) writeRepository() chatWriteRepository {
	if c.writeRepo != nil {
		return c.writeRepo
	}
	return c.svcCtx.Repo
}
