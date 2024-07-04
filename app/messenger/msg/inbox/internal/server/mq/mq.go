// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package mq

import (
	"context"
	"encoding/json"
	"fmt"

	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/internal/core"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// New new a grpc server.
func New(svcCtx *svc.ServiceContext, conf kafka.KafkaConsumerConf) *kafka.ConsumerGroup {
	s := kafka.MustKafkaConsumer(&conf)
	s.RegisterHandlers(
		conf.Topics[0],
		func(ctx context.Context, method, key string, value []byte) {
			logx.WithContext(ctx).Debugf("method: %s, key: %s, value: %s", method, key, value)

			switch protoreflect.FullName(method) {
			case proto.MessageName((*inbox.TLInboxEditUserMessageToInbox)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxEditUserMessageToInbox)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.editUserMessageToInbox - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.editUserMessageToInbox - request: %s", r)

					c.InboxEditUserMessageToInbox(r)
				})
			case proto.MessageName((*inbox.TLInboxEditChatMessageToInbox)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxEditChatMessageToInbox)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.editChatMessageToInbox - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.editChatMessageToInbox - request: %s", r)

					c.InboxEditChatMessageToInbox(r)
				})
			case proto.MessageName((*inbox.TLInboxDeleteMessagesToInbox)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxDeleteMessagesToInbox)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.deleteMessagesToInbox - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.deleteMessagesToInbox - request: %s", r)

					c.InboxDeleteMessagesToInbox(r)
				})
			case proto.MessageName((*inbox.TLInboxDeleteUserHistoryToInbox)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxDeleteUserHistoryToInbox)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.deleteUserHistoryToInbox - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.deleteUserHistoryToInbox - request: %s", r)

					c.InboxDeleteUserHistoryToInbox(r)
				})
			case proto.MessageName((*inbox.TLInboxDeleteChatHistoryToInbox)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxDeleteChatHistoryToInbox)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.deleteChatHistoryToInbox - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.deleteChatHistoryToInbox - request: %s", r)

					c.InboxDeleteChatHistoryToInbox(r)
				})
			case proto.MessageName((*inbox.TLInboxReadUserMediaUnreadToInbox)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxReadUserMediaUnreadToInbox)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.readUserMediaUnreadToInbox - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.readUserMediaUnreadToInbox - request: %s", r)

					c.InboxReadUserMediaUnreadToInbox(r)
				})
			case proto.MessageName((*inbox.TLInboxReadChatMediaUnreadToInbox)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxReadChatMediaUnreadToInbox)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.readChatMediaUnreadToInbox - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.readChatMediaUnreadToInbox - request: %s", r)

					c.InboxReadChatMediaUnreadToInbox(r)
				})
			case proto.MessageName((*inbox.TLInboxUpdateHistoryReaded)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxUpdateHistoryReaded)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.updateHistoryReaded - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.updateHistoryReaded - request: %s", r)

					c.InboxUpdateHistoryReaded(r)
				})
			case proto.MessageName((*inbox.TLInboxUpdatePinnedMessage)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxUpdatePinnedMessage)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.updatePinnedMessage - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.updatePinnedMessage - request: %s", r)

					c.InboxUpdatePinnedMessage(r)
				})
			case proto.MessageName((*inbox.TLInboxUnpinAllMessages)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxUnpinAllMessages)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.unpinAllMessages - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.unpinAllMessages - request: %s", r)

					c.InboxUnpinAllMessages(r)
				})
			case proto.MessageName((*inbox.TLInboxSendUserMessageToInboxV2)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxSendUserMessageToInboxV2)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.sendUserMessageToInboxV2 - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.sendUserMessageToInboxV2 - request: %s", r)

					c.InboxSendUserMessageToInboxV2(r)
				})
			case proto.MessageName((*inbox.TLInboxEditMessageToInboxV2)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxEditMessageToInboxV2)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.editMessageToInboxV2 - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.editMessageToInboxV2 - request: %s", r)

					c.InboxEditMessageToInboxV2(r)
				})
			case proto.MessageName((*inbox.TLInboxReadInboxHistory)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxReadInboxHistory)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.readInboxHistory - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.readInboxHistory - request: %s", r)

					c.InboxReadInboxHistory(r)
				})
			case proto.MessageName((*inbox.TLInboxReadOutboxHistory)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxReadOutboxHistory)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.readOutboxHistory - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.readOutboxHistory - request: %s", r)

					c.InboxReadOutboxHistory(r)
				})
			case proto.MessageName((*inbox.TLInboxReadMediaUnreadToInboxV2)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(inbox.TLInboxReadMediaUnreadToInboxV2)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Errorf("inbox.readMediaUnreadToInboxV2 - error: %v", err)
						return
					}
					c.Logger.Debugf("inbox.readMediaUnreadToInboxV2 - request: %s", r)

					c.InboxReadMediaUnreadToInboxV2(r)
				})
			default:
				err := fmt.Errorf("invalid key: %s", key)
				logx.Error(err.Error())
			}
		})
	return s
}
