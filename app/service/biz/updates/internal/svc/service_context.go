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

	messageclient "github.com/teamgram/teamgram-server/v2/app/service/biz/message/client"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/internal/config"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

type MessageQueryClient interface {
	MessageGetHistoryMessages(ctx context.Context, in *message.TLMessageGetHistoryMessages) (*message.VectorMessageBox, error)
}

type ServiceContext struct {
	Config        config.Config
	MessageClient MessageQueryClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := &ServiceContext{
		Config: c,
	}
	if hasClient(c.MessageClient) {
		ctx.MessageClient = messageclient.NewMessageClient(messageclient.MustNewKitexClient(c.MessageClient))
	}
	return ctx
}

func hasClient(c kitex.RpcClientConf) bool {
	return c.DestService != "" || c.Target != "" || len(c.Endpoints) > 0 || c.HasEtcd()
}
