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

	"github.com/gogo/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

// New new a grpc server.
func New(svcCtx *svc.ServiceContext, conf kafka.KafkaConsumerConf) *kafka.ConsumerGroup {
	s := kafka.MustKafkaConsumer(&conf)
	s.RegisterHandlers(
		conf.Topics[0],
		func(ctx context.Context, key string, value []byte) {
			switch key {
			case proto.MessageName((*inbox.TLInboxSendUserMessageToInbox)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxSendUserMessageToInbox)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.sendUserMessageToInbox - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.sendUserMessageToInbox - request: %s", r.DebugString())

				c.InboxSendUserMessageToInbox(r)
			case proto.MessageName((*inbox.TLInboxSendChatMessageToInbox)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxSendChatMessageToInbox)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.sendChatMessageToInbox - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.sendChatMessageToInbox - request: %s", r.DebugString())

				c.InboxSendChatMessageToInbox(r)
			case proto.MessageName((*inbox.TLInboxSendUserMultiMessageToInbox)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxSendUserMultiMessageToInbox)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.sendUserMultiMessageToInbox - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.sendUserMultiMessageToInbox - request: %s", r.DebugString())

				c.InboxSendUserMultiMessageToInbox(r)
			case proto.MessageName((*inbox.TLInboxSendChatMultiMessageToInbox)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxSendChatMultiMessageToInbox)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.sendChatMultiMessageToInbox - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.sendChatMultiMessageToInbox - request: %s", r.DebugString())

				c.InboxSendChatMultiMessageToInbox(r)
			case proto.MessageName((*inbox.TLInboxEditUserMessageToInbox)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxEditUserMessageToInbox)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.editUserMessageToInbox - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.editUserMessageToInbox - request: %s", r.DebugString())

				c.InboxEditUserMessageToInbox(r)
			case proto.MessageName((*inbox.TLInboxEditChatMessageToInbox)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxEditChatMessageToInbox)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.editChatMessageToInbox - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.editChatMessageToInbox - request: %s", r.DebugString())

				c.InboxEditChatMessageToInbox(r)
			case proto.MessageName((*inbox.TLInboxDeleteMessagesToInbox)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxDeleteMessagesToInbox)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.deleteMessagesToInbox - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.deleteMessagesToInbox - request: %s", r.DebugString())

				c.InboxDeleteMessagesToInbox(r)
			case proto.MessageName((*inbox.TLInboxDeleteUserHistoryToInbox)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxDeleteUserHistoryToInbox)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.deleteUserHistoryToInbox - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.deleteUserHistoryToInbox - request: %s", r.DebugString())

				c.InboxDeleteUserHistoryToInbox(r)
			case proto.MessageName((*inbox.TLInboxDeleteChatHistoryToInbox)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxDeleteChatHistoryToInbox)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.deleteChatHistoryToInbox - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.deleteChatHistoryToInbox - request: %s", r.DebugString())

				c.InboxDeleteChatHistoryToInbox(r)
			case proto.MessageName((*inbox.TLInboxReadUserMediaUnreadToInbox)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxReadUserMediaUnreadToInbox)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.readUserMediaUnreadToInbox - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.readUserMediaUnreadToInbox - request: %s", r.DebugString())

				c.InboxReadUserMediaUnreadToInbox(r)
			case proto.MessageName((*inbox.TLInboxReadChatMediaUnreadToInbox)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxReadChatMediaUnreadToInbox)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.readChatMediaUnreadToInbox - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.readChatMediaUnreadToInbox - request: %s", r.DebugString())

				c.InboxReadChatMediaUnreadToInbox(r)
			case proto.MessageName((*inbox.TLInboxUpdateHistoryReaded)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxUpdateHistoryReaded)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.updateHistoryReaded - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.updateHistoryReaded - request: %s", r.DebugString())

				c.InboxUpdateHistoryReaded(r)
			case proto.MessageName((*inbox.TLInboxUpdatePinnedMessage)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxUpdatePinnedMessage)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.updatePinnedMessage - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.updatePinnedMessage - request: %s", r.DebugString())

				c.InboxUpdatePinnedMessage(r)
			case proto.MessageName((*inbox.TLInboxUnpinAllMessages)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(inbox.TLInboxUnpinAllMessages)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Errorf("inbox.unpinAllMessages - error: %v", err)
					return
				}
				c.Logger.Infof("inbox.unpinAllMessages - request: %s", r.DebugString())

				c.InboxUnpinAllMessages(r)
			default:
				err := fmt.Errorf("invalid key: %s", key)
				logx.Error(err.Error())
			}
		})
	return s
}
