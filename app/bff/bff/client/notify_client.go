// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package bff_proxy_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

type NotifyClient interface {
	NotifySendNotifyData(ctx context.Context, in *mtproto.TLNotifySendNotifyData) (*mtproto.Bool, error)
}

type defaultNotifyClient struct {
	cli zrpc.Client
}

func NewNotifyClient(cli zrpc.Client) NotifyClient {
	return &defaultNotifyClient{
		cli: cli,
	}
}

// NotifySendNotifyData
// notify.sendNotifyData notifier:long trace_id:string date:long notify_type:string data:bytes = Bool;
func (m *defaultNotifyClient) NotifySendNotifyData(ctx context.Context, in *mtproto.TLNotifySendNotifyData) (*mtproto.Bool, error) {
	client := mtproto.NewRPCNotifyClient(m.cli.Conn())
	return client.NotifySendNotifyData(ctx, in)
}
