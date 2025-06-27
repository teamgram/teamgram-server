/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package echo1client

import (
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1/echo1/echo1service"

	"github.com/cloudwego/kitex/client"
)

func MustNewKitexClient(c kitex.RpcClientConf) client.Client {
	return kitex.MustNewClient(
		c,
		func(opts ...client.Option) (client.Client, error) {
			return client.NewClient(echo1service.NewServiceInfoForClient(), opts...)
		})
}

func NewKitexClient(c kitex.RpcClientConf) (client.Client, error) {
	return kitex.NewClient(
		c,
		func(opts ...client.Option) (client.Client, error) {
			return client.NewClient(echo1service.NewServiceInfoForClient(), opts...)
		})
}
