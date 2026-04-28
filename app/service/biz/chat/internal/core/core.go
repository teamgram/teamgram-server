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

	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
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
	GetUsersChatIDList(ctx context.Context, userIDs []int64) ([]*model.ChatParticipants, error)
	GetMyChatList(ctx context.Context, userID int64, isCreator bool) ([]*tg.MutableChat, error)
	Search(ctx context.Context, selfID int64, q string, offset int64, limit int32) ([]*tg.MutableChat, error)
}

type ChatCore struct {
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	readRepo chatReadRepository
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
