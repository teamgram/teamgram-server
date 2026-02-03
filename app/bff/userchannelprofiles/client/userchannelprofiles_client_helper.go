/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgooo Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package userchannelprofilesclient

import (
	"github.com/teamgooo/teamgooo-server/app/bff/userchannelprofiles/userchannelprofiles/userchannelprofilesservice"
	"github.com/teamgooo/teamgooo-server/pkg/net/kitex"

	"github.com/cloudwego/kitex/client"
)

func MustNewKitexClient(c kitex.RpcClientConf) client.Client {
	return kitex.MustNewClient(
		c,
		func(opts ...client.Option) (client.Client, error) {
			return client.NewClient(userchannelprofilesservice.NewServiceInfoForClient(), opts...)
		})
}

func NewKitexClient(c kitex.RpcClientConf) (client.Client, error) {
	return kitex.NewClient(
		c,
		func(opts ...client.Option) (client.Client, error) {
			return client.NewClient(userchannelprofilesservice.NewServiceInfoForClient(), opts...)
		})
}
