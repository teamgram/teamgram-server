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

package inbox_client

import (
	"context"

	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"

	"github.com/gogo/protobuf/proto"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type defaultInboxMqClient struct {
	cli *kafka.Producer
}

func NewInboxMqClient(cli *kafka.Producer) InboxClient {
	return &defaultInboxMqClient{
		cli: cli,
	}
}

func (m *defaultInboxMqClient) sendMessage(ctx context.Context, k string, in interface{}) (*mtproto.Void, error) {
	var (
		b   []byte
		err error
	)

	b, err = jsonx.Marshal(in)
	if err != nil {
		return nil, err
	}

	_, _, err = m.cli.SendMessage(ctx, k, b)
	if err != nil {
		return nil, err
	}

	return mtproto.EmptyVoid, nil
}

// InboxSendUserMessageToInbox
// inbox.sendUserMessageToInbox from_id:long peer_user_id:long message:InboxMessageData = Void;
func (m *defaultInboxMqClient) InboxSendUserMessageToInbox(ctx context.Context, in *inbox.TLInboxSendUserMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxSendChatMessageToInbox
// inbox.sendChatMessageToInbox from_id:long peer_chat_id:long message:InboxMessageData = Void;
func (m *defaultInboxMqClient) InboxSendChatMessageToInbox(ctx context.Context, in *inbox.TLInboxSendChatMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxSendUserMultiMessageToInbox
// inbox.sendUserMultiMessageToInbox from_id:long peer_user_id:long message:Vector<InboxMessageData> = Void;
func (m *defaultInboxMqClient) InboxSendUserMultiMessageToInbox(ctx context.Context, in *inbox.TLInboxSendUserMultiMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxSendChatMultiMessageToInbox
// inbox.sendChatMultiMessageToInbox from_id:long peer_chat_id:long message:Vector<InboxMessageData> = Void;
func (m *defaultInboxMqClient) InboxSendChatMultiMessageToInbox(ctx context.Context, in *inbox.TLInboxSendChatMultiMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxEditUserMessageToInbox
// inbox.editUserMessageToInbox from_id:long peer_user_id:long message:Message = Void;
func (m *defaultInboxMqClient) InboxEditUserMessageToInbox(ctx context.Context, in *inbox.TLInboxEditUserMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxEditChatMessageToInbox
// inbox.editChatMessageToInbox from_id:long peer_chat_id:long message:Message = Void;
func (m *defaultInboxMqClient) InboxEditChatMessageToInbox(ctx context.Context, in *inbox.TLInboxEditChatMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxDeleteMessagesToInbox
// inbox.deleteMessagesToInbox from_id:long id:Vector<int> = Void;
func (m *defaultInboxMqClient) InboxDeleteMessagesToInbox(ctx context.Context, in *inbox.TLInboxDeleteMessagesToInbox) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxDeleteUserHistoryToInbox
// inbox.deleteUserHistoryToInbox flags:# from_id:long peer_user_id:long just_clear:flags.1?true max_id:int = Void;
func (m *defaultInboxMqClient) InboxDeleteUserHistoryToInbox(ctx context.Context, in *inbox.TLInboxDeleteUserHistoryToInbox) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxDeleteChatHistoryToInbox
// inbox.deleteChatHistoryToInbox from_id:long peer_chat_id:long max_id:int = Void;
func (m *defaultInboxMqClient) InboxDeleteChatHistoryToInbox(ctx context.Context, in *inbox.TLInboxDeleteChatHistoryToInbox) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxReadUserMediaUnreadToInbox
// inbox.readUserMediaUnreadToInbox from_id:long id:Vector<int> = Void;
func (m *defaultInboxMqClient) InboxReadUserMediaUnreadToInbox(ctx context.Context, in *inbox.TLInboxReadUserMediaUnreadToInbox) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxReadChatMediaUnreadToInbox
// inbox.readChatMediaUnreadToInbox from_id:long peer_chat_id:long id:Vector<int> = Void;
func (m *defaultInboxMqClient) InboxReadChatMediaUnreadToInbox(ctx context.Context, in *inbox.TLInboxReadChatMediaUnreadToInbox) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxUpdateHistoryReaded
// inbox.updateHistoryReaded from_id:long peer_type:int peer_id:long max_id:int = Void;
func (m *defaultInboxMqClient) InboxUpdateHistoryReaded(ctx context.Context, in *inbox.TLInboxUpdateHistoryReaded) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxUpdatePinnedMessage
// inbox.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;
func (m *defaultInboxMqClient) InboxUpdatePinnedMessage(ctx context.Context, in *inbox.TLInboxUpdatePinnedMessage) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// InboxUnpinAllMessages
// inbox.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = messages.AffectedHistory;
func (m *defaultInboxMqClient) InboxUnpinAllMessages(ctx context.Context, in *inbox.TLInboxUnpinAllMessages) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}
