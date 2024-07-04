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

package inboxclient

import (
	"context"
	"strconv"

	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"

	"github.com/zeromicro/go-zero/core/jsonx"
	"google.golang.org/protobuf/proto"
)

type defaultInboxMqClient struct {
	cli *kafka.Producer
}

func NewInboxMqClient(cli *kafka.Producer) InboxClient {
	return &defaultInboxMqClient{
		cli: cli,
	}
}

func (m *defaultInboxMqClient) sendMessage(ctx context.Context, method, k string, in interface{}) (*mtproto.Void, error) {
	var (
		b   []byte
		err error
	)

	b, err = jsonx.Marshal(in)
	if err != nil {
		return nil, err
	}

	_, _, err = m.cli.SendMessageV2(ctx, method, k, b)
	if err != nil {
		return nil, err
	}

	return mtproto.EmptyVoid, nil
}

// InboxEditUserMessageToInbox
// inbox.editUserMessageToInbox from_id:long peer_user_id:long message:Message = Void;
func (m *defaultInboxMqClient) InboxEditUserMessageToInbox(ctx context.Context, in *inbox.TLInboxEditUserMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetPeerUserId(), 10),
		in)
}

// InboxEditChatMessageToInbox
// inbox.editChatMessageToInbox from_id:long peer_chat_id:long message:Message = Void;
func (m *defaultInboxMqClient) InboxEditChatMessageToInbox(ctx context.Context, in *inbox.TLInboxEditChatMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetPeerChatId(), 10),
		in)
}

// InboxDeleteMessagesToInbox
// inbox.deleteMessagesToInbox from_id:long id:Vector<int> = Void;
func (m *defaultInboxMqClient) InboxDeleteMessagesToInbox(ctx context.Context, in *inbox.TLInboxDeleteMessagesToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetPeerId(), 10),
		in)
}

// InboxDeleteUserHistoryToInbox
// inbox.deleteUserHistoryToInbox flags:# from_id:long peer_user_id:long just_clear:flags.1?true max_id:int = Void;
func (m *defaultInboxMqClient) InboxDeleteUserHistoryToInbox(ctx context.Context, in *inbox.TLInboxDeleteUserHistoryToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetPeerUserId(), 10),
		in)
}

// InboxDeleteChatHistoryToInbox
// inbox.deleteChatHistoryToInbox from_id:long peer_chat_id:long max_id:int = Void;
func (m *defaultInboxMqClient) InboxDeleteChatHistoryToInbox(ctx context.Context, in *inbox.TLInboxDeleteChatHistoryToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetPeerChatId(), 10),
		in)
}

// InboxReadUserMediaUnreadToInbox
// inbox.readUserMediaUnreadToInbox from_id:long id:Vector<int> = Void;
func (m *defaultInboxMqClient) InboxReadUserMediaUnreadToInbox(ctx context.Context, in *inbox.TLInboxReadUserMediaUnreadToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetPeerUserId(), 10),
		in)
}

// InboxReadChatMediaUnreadToInbox
// inbox.readChatMediaUnreadToInbox from_id:long peer_chat_id:long id:Vector<int> = Void;
func (m *defaultInboxMqClient) InboxReadChatMediaUnreadToInbox(ctx context.Context, in *inbox.TLInboxReadChatMediaUnreadToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetPeerChatId(), 10),
		in)
}

// InboxUpdateHistoryReaded
// inbox.updateHistoryReaded from_id:long peer_type:int peer_id:long max_id:int = Void;
func (m *defaultInboxMqClient) InboxUpdateHistoryReaded(ctx context.Context, in *inbox.TLInboxUpdateHistoryReaded) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetPeerId(), 10),
		in)
}

// InboxUpdatePinnedMessage
// inbox.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;
func (m *defaultInboxMqClient) InboxUpdatePinnedMessage(ctx context.Context, in *inbox.TLInboxUpdatePinnedMessage) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetPeerId(), 10),
		in)
}

// InboxUnpinAllMessages
// inbox.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = messages.AffectedHistory;
func (m *defaultInboxMqClient) InboxUnpinAllMessages(ctx context.Context, in *inbox.TLInboxUnpinAllMessages) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetPeerId(), 10),
		in)
}

// InboxSendUserMessageToInboxV2
// inbox.sendUserMessageToInboxV2 flags:# user_id:long out:flags.0?true from_id:long peer_user_id:long inbox:MessageBox users:flags.1?Vector<ImmutableUser> = Void;
func (m *defaultInboxMqClient) InboxSendUserMessageToInboxV2(ctx context.Context, in *inbox.TLInboxSendUserMessageToInboxV2) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetUserId(), 10),
		in)
}

// InboxEditMessageToInboxV2
// inbox.editMessageToInboxV2 flags:# user_id:long out:flags.0?true from_id:long fromAuthKeyId:long peer_type:int peer_id:long box:MessageBox users:flags.1?Vector<User> chats:flags.2?Vector<Chat> = Void;
func (m *defaultInboxMqClient) InboxEditMessageToInboxV2(ctx context.Context, in *inbox.TLInboxEditMessageToInboxV2) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetUserId(), 10),
		in)
}

// InboxReadInboxHistory
// inbox.readInboxHistory user_id:long auth_key_id:long peer_type:int peer_id:long pts:int pts_count:int unread_count:int read_inbox_max_id:int max_id:int = Void;
func (m *defaultInboxMqClient) InboxReadInboxHistory(ctx context.Context, in *inbox.TLInboxReadInboxHistory) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetUserId(), 10),
		in)
}

// InboxReadOutboxHistory
// inbox.readOutboxHistory user_id:long peer_type:int peer_id:long max_dialog_message_id:long = Void;
func (m *defaultInboxMqClient) InboxReadOutboxHistory(ctx context.Context, in *inbox.TLInboxReadOutboxHistory) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetUserId(), 10),
		in)
}

// InboxReadMediaUnreadToInboxV2
// inbox.readMediaUnreadToInboxV2 user_id:long peer_type:int peer_id:long id:Vector<InboxMessageId> = Void;
func (m *defaultInboxMqClient) InboxReadMediaUnreadToInboxV2(ctx context.Context, in *inbox.TLInboxReadMediaUnreadToInboxV2) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		string(proto.MessageName(in)),
		strconv.FormatInt(in.GetUserId(), 10),
		in)
}
