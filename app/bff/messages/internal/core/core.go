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

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	idgenpb "github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/zeromicro/go-zero/core/logx"
)

type sendMessageClient interface {
	MsgSendMessageV2(ctx context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error)
	MsgGetUserMessage(ctx context.Context, in *msg.TLMsgGetUserMessage) (*tg.MessageBox, error)
	MsgGetUserMessageList(ctx context.Context, in *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error)
}

type idgenClient interface {
	IdgenNextId(ctx context.Context, in *idgenpb.TLIdgenNextId) (*tg.Int64, error)
}

type resolveMediaClient interface {
	MediaUploadPhotoFile(ctx context.Context, in *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error)
	MediaGetPhoto(ctx context.Context, in *mediapb.TLMediaGetPhoto) (*tg.Photo, error)
	MediaGetPhotoSizeList(ctx context.Context, in *mediapb.TLMediaGetPhotoSizeList) (*mediapb.PhotoSizeList, error)
	MediaUploadedDocumentMedia(ctx context.Context, in *mediapb.TLMediaUploadedDocumentMedia) (*tg.MessageMedia, error)
	MediaGetDocument(ctx context.Context, in *mediapb.TLMediaGetDocument) (*tg.Document, error)
}

type userLookupClient interface {
	UserGetUserIdByPhone(ctx context.Context, in *userpb.TLUserGetUserIdByPhone) (*tg.Int64, error)
}

type getHistoryClient interface {
	MsgGetHistory(ctx context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error)
}

type readHistoryClient interface {
	MsgReadHistoryV2(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error)
}

type updatePinnedMessageClient interface {
	MsgUpdatePinnedMessage(ctx context.Context, in *msg.TLMsgUpdatePinnedMessage) (*tg.Updates, error)
}

type deleteMessagesClient interface {
	MsgDeleteMessages(ctx context.Context, in *msg.TLMsgDeleteMessages) (*tg.MessagesAffectedMessages, error)
}

type deleteHistoryClient interface {
	MsgDeleteHistory(ctx context.Context, in *msg.TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error)
}

type editMessageClient interface {
	MsgEditMessageV2(ctx context.Context, in *msg.TLMsgEditMessageV2) (*tg.Updates, error)
}

type MessagesCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	Logger logx.Logger
	MD     *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *MessagesCore {
	return &MessagesCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func (c *MessagesCore) Debugf(format string, v ...any) {
	c.Logger.Debugf(format, v...)
}

func (c *MessagesCore) Errorf(format string, v ...any) {
	c.Logger.Errorf(format, v...)
}
