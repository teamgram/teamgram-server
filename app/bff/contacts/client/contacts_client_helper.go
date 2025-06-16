/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package contactsclient

import (
	"github.com/teamgram/teamgram-server/v2/app/bff/contacts/contacts/contactsservice"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"

	"github.com/cloudwego/kitex/client"
)

func MustNewKitexClient(c kitex.RpcClientConf) client.Client {
	return kitex.MustNewClient(
		c,
		func(opts ...client.Option) (client.Client, error) {
			return client.NewClient(contactsservice.NewServiceInfoForClient(), opts...)
		})
}

func NewKitexClient(c kitex.RpcClientConf) (client.Client, error) {
	return kitex.NewClient(
		c,
		func(opts ...client.Option) (client.Client, error) {
			return client.NewClient(contactsservice.NewServiceInfoForClient(), opts...)
		})
}
