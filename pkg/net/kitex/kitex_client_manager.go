// Copyright 2024 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package kitex

import (
	"fmt"
	"io"
	"strings"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/syncx"
)

var (
	clientManager            = syncx.NewResourceManager()
	newClientWithServiceInfo = NewClientWithServiceInfo
)

type client2 struct {
	Client
}

func (c *client2) Close() error {
	if closer, ok := c.Client.(interface{ Close() error }); ok {
		return closer.Close()
	}
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

func newCachedClient(c RpcClientConf) (io.Closer, error) {
	cli, err := newClientWithServiceInfo(c, iface.GetKitexServiceInfoForClient(c.ServiceName))
	if err != nil {
		return nil, err
	}
	if cli == nil {
		return nil, fmt.Errorf("nil client created for %s", c.ServiceName)
	}

	return &client2{Client: cli}, nil
}

func GetCachedKitexClient(c RpcClientConf) Client {
	if c.Etcd.Key == "" && len(c.Endpoints) == 0 {
		panic(c)
	}
	logx.Infof("kitex client cache lookup: service=%s endpoints=%v etcd_key=%s", c.ServiceName, c.Endpoints, c.Etcd.Key)

	var (
		val io.Closer
		err error
	)
	if len(c.Endpoints) > 0 {
		val, err = clientManager.GetResource(c.ServiceName+"@"+strings.Join(c.Endpoints, "/"), func() (io.Closer, error) {
			return newCachedClient(c)
		})
		if err != nil {
			panic(err)
		}
	} else {
		val, err = clientManager.GetResource(c.ServiceName+"@"+c.Etcd.Key, func() (io.Closer, error) {
			return newCachedClient(c)
		})
		if err != nil {
			panic(err)
		}
	}

	return val.(Client)
}
