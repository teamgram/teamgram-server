/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package inbox_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type InboxClient interface {
	InboxSendUserMessageToInbox(ctx context.Context, in *inbox.TLInboxSendUserMessageToInbox) (*mtproto.Void, error)
	InboxSendChatMessageToInbox(ctx context.Context, in *inbox.TLInboxSendChatMessageToInbox) (*mtproto.Void, error)
	InboxSendUserMultiMessageToInbox(ctx context.Context, in *inbox.TLInboxSendUserMultiMessageToInbox) (*mtproto.Void, error)
	InboxSendChatMultiMessageToInbox(ctx context.Context, in *inbox.TLInboxSendChatMultiMessageToInbox) (*mtproto.Void, error)
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
}

type defaultInboxClient struct {
	cli zrpc.Client
}

func NewInboxClient(cli zrpc.Client) InboxClient {
	return &defaultInboxClient{
		cli: cli,
	}
}

// InboxSendUserMessageToInbox
// inbox.sendUserMessageToInbox from_id:long peer_user_id:long message:InboxMessageData = Void;
func (m *defaultInboxClient) InboxSendUserMessageToInbox(ctx context.Context, in *inbox.TLInboxSendUserMessageToInbox) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxSendUserMessageToInbox(ctx, in)
}

// InboxSendChatMessageToInbox
// inbox.sendChatMessageToInbox from_id:long peer_chat_id:long message:InboxMessageData = Void;
func (m *defaultInboxClient) InboxSendChatMessageToInbox(ctx context.Context, in *inbox.TLInboxSendChatMessageToInbox) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxSendChatMessageToInbox(ctx, in)
}

// InboxSendUserMultiMessageToInbox
// inbox.sendUserMultiMessageToInbox from_id:long peer_user_id:long message:Vector<InboxMessageData> = Void;
func (m *defaultInboxClient) InboxSendUserMultiMessageToInbox(ctx context.Context, in *inbox.TLInboxSendUserMultiMessageToInbox) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxSendUserMultiMessageToInbox(ctx, in)
}

// InboxSendChatMultiMessageToInbox
// inbox.sendChatMultiMessageToInbox from_id:long peer_chat_id:long message:Vector<InboxMessageData> = Void;
func (m *defaultInboxClient) InboxSendChatMultiMessageToInbox(ctx context.Context, in *inbox.TLInboxSendChatMultiMessageToInbox) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxSendChatMultiMessageToInbox(ctx, in)
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
// inbox.deleteMessagesToInbox from_id:long id:Vector<long> = Void;
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
// inbox.readUserMediaUnreadToInbox from_id:long id:Vector<int> = Void;
func (m *defaultInboxClient) InboxReadUserMediaUnreadToInbox(ctx context.Context, in *inbox.TLInboxReadUserMediaUnreadToInbox) (*mtproto.Void, error) {
	client := inbox.NewRPCInboxClient(m.cli.Conn())
	return client.InboxReadUserMediaUnreadToInbox(ctx, in)
}

// InboxReadChatMediaUnreadToInbox
// inbox.readChatMediaUnreadToInbox from_id:long peer_chat_id:long id:Vector<int> = Void;
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
