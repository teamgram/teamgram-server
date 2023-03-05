// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package bff_proxy_client

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/proto/mtproto"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type defaultNotifyMqClient struct {
	cli *kafka.Producer
}

func NewNotifyMqClient(cli *kafka.Producer) NotifyClient {
	return &defaultNotifyMqClient{
		cli: cli,
	}
}

func (m *defaultNotifyMqClient) sendMessage(ctx context.Context, k string, in interface{}) (*mtproto.Bool, error) {
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

	return mtproto.BoolTrue, nil
}

// NotifySendNotifyData
// notify.sendNotifyData notifier:long trace_id:string date:long notify_type:string data:bytes = Bool;
func (m *defaultNotifyMqClient) NotifySendNotifyData(ctx context.Context, in *mtproto.TLNotifySendNotifyData) (*mtproto.Bool, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.Notifier),
		in)
}
