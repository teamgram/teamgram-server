/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package inboxclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type InboxClient interface {
	InboxEditUserMessageToInbox(ctx context.Context, in *inbox.TLInboxEditUserMessageToInbox) (*mtproto.Void, error)
	InboxEditChatMessageToInbox(ctx context.Context, in *inbox.TLInboxEditChatMessageToInbox) (*mtproto.Void, error)
	InboxDeleteMessagesToInbox(ctx context.Context, in *inbox.TLInboxDeleteMessagesToInbox) (*mtproto.Void, error)
	InboxDeleteUserHistoryToInbox(ctx context.Context, in *inbox.TLInboxDeleteUserHistoryToInbox) (*mtproto.Void, error)
	InboxDeleteChatHistoryToInbox(ctx context.Context, in *inbox.TLInboxDeleteChatHistoryToInbox) (*mtproto.Void, error)
	InboxReadUserMediaUnreadToInbox(ctx context.Context, in *inbox.TLInboxReadUserMediaUnreadToInbox) (*mtproto.Void, error)
	InboxReadChatMediaUnreadToInbox(ctx context.Context, in *inbox.TLInboxReadChatMediaUnreadToInbox) (*mtproto.Void, error)
	InboxUpdateHistoryReaded(ctx context.Context, in *inbox.TLInboxUpdateHistoryReaded) (*mtproto.Void, error)
	InboxUpdatePinnedMessage(ctx context.Context, in *inbox.TLInboxUpdatePinnedMessage) (*mtproto.Void, error)
	InboxUnpinAllMessages(ctx context.Context, in *inbox.TLInboxUnpinAllMessages) (*mtproto.Void, error)
	InboxSendUserMessageToInboxV2(ctx context.Context, in *inbox.TLInboxSendUserMessageToInboxV2) (*mtproto.Void, error)
	InboxEditMessageToInboxV2(ctx context.Context, in *inbox.TLInboxEditMessageToInboxV2) (*mtproto.Void, error)
	InboxReadInboxHistory(ctx context.Context, in *inbox.TLInboxReadInboxHistory) (*mtproto.Void, error)
	InboxReadOutboxHistory(ctx context.Context, in *inbox.TLInboxReadOutboxHistory) (*mtproto.Void, error)
	InboxReadMediaUnreadToInboxV2(ctx context.Context, in *inbox.TLInboxReadMediaUnreadToInboxV2) (*mtproto.Void, error)
}

type defaultInboxClient struct {
	cli zrpc.Client
}

func NewInboxClient(cli zrpc.Client) InboxClient {
	return &defaultInboxClient{
		cli: cli,
	}
}

// InboxEditUserMessageToInbox
// inbox.editUserMessageToInbox from_id:long peer_user_id:long message:Message = Void;
func (m *defaultInboxClient) InboxEditUserMessageToInbox(ctx context.Context, in *inbox.TLInboxEditUserMessageToInbox) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxEditUserMessageToInbox(ctx, in)
}

// InboxEditChatMessageToInbox
// inbox.editChatMessageToInbox from_id:long peer_chat_id:long message:Message = Void;
func (m *defaultInboxClient) InboxEditChatMessageToInbox(ctx context.Context, in *inbox.TLInboxEditChatMessageToInbox) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxEditChatMessageToInbox(ctx, in)
}

// InboxDeleteMessagesToInbox
// inbox.deleteMessagesToInbox from_id:long peer_type:int peer_id:long id:Vector<long> = Void;
func (m *defaultInboxClient) InboxDeleteMessagesToInbox(ctx context.Context, in *inbox.TLInboxDeleteMessagesToInbox) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxDeleteMessagesToInbox(ctx, in)
}

// InboxDeleteUserHistoryToInbox
// inbox.deleteUserHistoryToInbox flags:# from_id:long peer_user_id:long just_clear:flags.1?true max_id:int = Void;
func (m *defaultInboxClient) InboxDeleteUserHistoryToInbox(ctx context.Context, in *inbox.TLInboxDeleteUserHistoryToInbox) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxDeleteUserHistoryToInbox(ctx, in)
}

// InboxDeleteChatHistoryToInbox
// inbox.deleteChatHistoryToInbox from_id:long peer_chat_id:long max_id:int = Void;
func (m *defaultInboxClient) InboxDeleteChatHistoryToInbox(ctx context.Context, in *inbox.TLInboxDeleteChatHistoryToInbox) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxDeleteChatHistoryToInbox(ctx, in)
}

// InboxReadUserMediaUnreadToInbox
// inbox.readUserMediaUnreadToInbox from_id:long peer_user_id:long id:Vector<InboxMessageId> = Void;
func (m *defaultInboxClient) InboxReadUserMediaUnreadToInbox(ctx context.Context, in *inbox.TLInboxReadUserMediaUnreadToInbox) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxReadUserMediaUnreadToInbox(ctx, in)
}

// InboxReadChatMediaUnreadToInbox
// inbox.readChatMediaUnreadToInbox from_id:long peer_chat_id:long id:Vector<InboxMessageId> = Void;
func (m *defaultInboxClient) InboxReadChatMediaUnreadToInbox(ctx context.Context, in *inbox.TLInboxReadChatMediaUnreadToInbox) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxReadChatMediaUnreadToInbox(ctx, in)
}

// InboxUpdateHistoryReaded
// inbox.updateHistoryReaded from_id:long peer_type:int peer_id:long max_id:int sender:long = Void;
func (m *defaultInboxClient) InboxUpdateHistoryReaded(ctx context.Context, in *inbox.TLInboxUpdateHistoryReaded) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxUpdateHistoryReaded(ctx, in)
}

// InboxUpdatePinnedMessage
// inbox.updatePinnedMessage flags:# user_id:long unpin:flags.1?true peer_type:int peer_id:long id:int dialog_message_id:long = Void;
func (m *defaultInboxClient) InboxUpdatePinnedMessage(ctx context.Context, in *inbox.TLInboxUpdatePinnedMessage) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxUpdatePinnedMessage(ctx, in)
}

// InboxUnpinAllMessages
// inbox.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = Void;
func (m *defaultInboxClient) InboxUnpinAllMessages(ctx context.Context, in *inbox.TLInboxUnpinAllMessages) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxUnpinAllMessages(ctx, in)
}

// InboxSendUserMessageToInboxV2
// inbox.sendUserMessageToInboxV2 flags:# user_id:long out:flags.0?true from_id:long from_auth_keyId:long peer_type:int peer_id:long box_list:Vector<MessageBox> users:flags.1?Vector<User> chats:flags.2?Vector<Chat> = Void;
func (m *defaultInboxClient) InboxSendUserMessageToInboxV2(ctx context.Context, in *inbox.TLInboxSendUserMessageToInboxV2) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxSendUserMessageToInboxV2(ctx, in)
}

// InboxEditMessageToInboxV2
// inbox.editMessageToInboxV2 flags:# user_id:long out:flags.0?true from_id:long from_auth_keyId:long peer_type:int peer_id:long new_message:MessageBox dst_message:flags.1?MessageBox users:flags.2?Vector<User> chats:flags.3?Vector<Chat> = Void;
func (m *defaultInboxClient) InboxEditMessageToInboxV2(ctx context.Context, in *inbox.TLInboxEditMessageToInboxV2) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxEditMessageToInboxV2(ctx, in)
}

// InboxReadInboxHistory
// inbox.readInboxHistory user_id:long auth_key_id:long peer_type:int peer_id:long pts:int pts_count:int unread_count:int read_inbox_max_id:int max_id:int = Void;
func (m *defaultInboxClient) InboxReadInboxHistory(ctx context.Context, in *inbox.TLInboxReadInboxHistory) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxReadInboxHistory(ctx, in)
}

// InboxReadOutboxHistory
// inbox.readOutboxHistory user_id:long peer_type:int peer_id:long max_dialog_message_id:long = Void;
func (m *defaultInboxClient) InboxReadOutboxHistory(ctx context.Context, in *inbox.TLInboxReadOutboxHistory) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxReadOutboxHistory(ctx, in)
}

// InboxReadMediaUnreadToInboxV2
// inbox.readMediaUnreadToInboxV2 user_id:long peer_type:int peer_id:long id:Vector<InboxMessageId> = Void;
func (m *defaultInboxClient) InboxReadMediaUnreadToInboxV2(ctx context.Context, in *inbox.TLInboxReadMediaUnreadToInboxV2) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxReadMediaUnreadToInboxV2(ctx, in)
}
