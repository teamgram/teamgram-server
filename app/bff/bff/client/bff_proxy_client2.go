// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package bffproxyclient

import (
	"strings"

	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

type proxyClient struct {
	c kitex.Client
}

//// NewClientByServiceInfoForClient returns a Client.
//func NewClientByServiceInfoForClient(c RpcClientConf, svc *serviceinfo.ServiceInfo) (Client, error) {
//	return NewClient(c, func(opts ...client.Option) (Client, error) {
//		return client.NewClient(svc, opts...)
//	})
//}

type BFFProxyClient2 struct {
	// zrpc.Client
	BFFClients map[string]proxyClient
}

func NewBFFProxyClient2(cList []kitex.RpcClientConf, idMap map[string]string) *BFFProxyClient2 {
	var (
		clients   = make(map[string]proxyClient)
		registers = iface.GetRPCContextRegisters()
	)

	for _, c := range cList {
		cli := kitex.GetCachedKitexClient(c)
		for k, v := range idMap {
			if v == c.Etcd.Key {
				clients[k] = proxyClient{c: cli}
			}
		}
	}

	bizClients := make(map[string]proxyClient)
	for m, ctx := range registers {
		for k, _ := range idMap {
			if strings.HasPrefix(ctx.Method, k) {
				bizClients[m] = clients[k]
				break
			}
		}
	}

	return &BFFProxyClient2{
		BFFClients: bizClients,
	}
}
