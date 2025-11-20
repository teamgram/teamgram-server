/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2025 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package msgtransferclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msgtransfer/msgtransfer"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type MsgtransferClient interface {
	MsgtransferSendMessageToOutbox(ctx context.Context, in *msgtransfer.TLMsgtransferSendMessageToOutbox) (*mtproto.MessageBoxList, error)
	MsgtransferSendMessageToInbox(ctx context.Context, in *msgtransfer.TLMsgtransferSendMessageToInbox) (*mtproto.Void, error)
}

type defaultMsgtransferClient struct {
	cli zrpc.Client
}

func NewMsgtransferClient(cli zrpc.Client) MsgtransferClient {
	return &defaultMsgtransferClient{
		cli: cli,
	}
}

// MsgtransferSendMessageToOutbox
// msgtransfer.sendMessageToOutbox flags:# user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> users:flags.1?Vector<User> chats:flags.2?Vector<Chat> = MessageBoxList;
func (m *defaultMsgtransferClient) MsgtransferSendMessageToOutbox(ctx context.Context, in *msgtransfer.TLMsgtransferSendMessageToOutbox) (*mtproto.MessageBoxList, error) {
	md := metadata.RpcMetadataFromIncoming(ctx)
	if md != nil {
		ctx, _ = metadata.RpcMetadataToOutgoing(ctx, md)
	}
	client := msgtransfer.NewRPCMsgtransferClient(m.cli.Conn())
	return client.MsgtransferSendMessageToOutbox(ctx, in)
}

// MsgtransferSendMessageToInbox
// msgtransfer.sendMessageToInbox flags:# user_id:long from_id:long from_auth_keyId:long peer_type:int peer_id:long box_list:Vector<MessageBox> users:flags.1?Vector<User> chats:flags.2?Vector<Chat> = Void;
func (m *defaultMsgtransferClient) MsgtransferSendMessageToInbox(ctx context.Context, in *msgtransfer.TLMsgtransferSendMessageToInbox) (*mtproto.Void, error) {
	md := metadata.RpcMetadataFromIncoming(ctx)
	if md != nil {
		ctx, _ = metadata.RpcMetadataToOutgoing(ctx, md)
	}
	client := msgtransfer.NewRPCMsgtransferClient(m.cli.Conn())
	return client.MsgtransferSendMessageToInbox(ctx, in)
}
