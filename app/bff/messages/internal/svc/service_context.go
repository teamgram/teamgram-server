// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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

package svc

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/config"
	msgclient "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type MsgSendClient interface {
	MsgSendMessageV2(ctx context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error)
	MsgReadHistory(ctx context.Context, in *msg.TLMsgReadHistory) (*tg.MessagesAffectedMessages, error)
	MsgReadHistoryV2(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error)
	MsgUpdatePinnedMessage(ctx context.Context, in *msg.TLMsgUpdatePinnedMessage) (*tg.Updates, error)
	MsgUnpinAllMessages(ctx context.Context, in *msg.TLMsgUnpinAllMessages) (*tg.MessagesAffectedHistory, error)
	MsgDeleteHistory(ctx context.Context, in *msg.TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error)
}

type ServiceContext struct {
	Config    config.Config
	MsgClient MsgSendClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := &ServiceContext{
		Config: c,
	}
	if hasClient(c.MsgClient) {
		ctx.MsgClient = msgclient.NewMsgClient(msgclient.MustNewKitexClient(c.MsgClient))
	}
	return ctx
}

func hasClient(c kitex.RpcClientConf) bool {
	return c.DestService != "" || c.Target != "" || len(c.Endpoints) > 0 || c.HasEtcd()
}
