// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package kitex

import (
	"github.com/teamgram/proto/v2/iface"
	"io"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/syncx"
)

var (
	clientManager = syncx.NewResourceManager()
)

type client2 struct {
	Client
}

func (c *client2) Close() error {
	return nil
}

/*
	cli := kitex.MustNewClient(
		"interface.session",
		c,
		func(destService string, opts ...client.Option) (client.Client, error) {
			return client.NewClient(sessionservice.NewServiceInfoForClient(), opts...)
		},
		client.WithCodec(codec.NewZRpcCodec(true)))
*/

func GetCachedKitexClient(c RpcClientConf) Client {
	var (
		val io.Closer
		err error
	)
	if c.Etcd.Key == "" && len(c.Endpoints) == 0 {
		panic(c)
	}
	logx.Infof("client: %v", c)
	if len(c.Endpoints) > 0 {
		val, err = clientManager.GetResource(strings.Join(c.Endpoints, "/"), func() (io.Closer, error) {
			cli, _ := NewClientWithServiceInfo(c, iface.GetKitexServiceInfoForClient(c.ServiceName))
			return &client2{
				Client: cli,
			}, nil
		})
		if err != nil {
			panic(err)
		}
	} else {
		val, err = clientManager.GetResource(c.Etcd.Key, func() (io.Closer, error) {
			cli, _ := NewClientWithServiceInfo(c, iface.GetKitexServiceInfoForClient(c.ServiceName))
			return &client2{
				Client: cli,
			}, nil
		})
		if err != nil {
			panic(err)
		}
	}

	return val.(Client)
}
