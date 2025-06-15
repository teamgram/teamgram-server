// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package kitex

import (
	"net"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/zeromicro/go-zero/core/logx"
)

type (
	ServerOption = server.Option
)

type (
	// RegisterFn defines the method to register a server.
	RegisterFn func(server.Server) error
)

// A RpcServer is a rpc server.
type RpcServer struct {
	server.Server
	// register RegisterFn
}

// MustNewServer returns a RpcSever, exits on any error.
func MustNewServer(c RpcServerConf, register RegisterFn) *RpcServer {
	server2, err := NewServer(c, register)
	logx.Must(err)
	return server2
}

// NewServer returns a RpcServer.
func NewServer(c RpcServerConf, register RegisterFn) (*RpcServer, error) {
	var (
		err error
	)

	if err = c.Validate(); err != nil {
		return nil, err
	}

	var (
		options []server.Option
	)

	// c.ServiceConf.Name
	options = append(options, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: c.ServiceConf.Name,
	}))

	// codec
	if c.Codec == "zrpc" {
		options = append(options, server.WithCodec(codec.NewZRpcCodec(true)))
	}

	// c.ListenOn
	addr, err := net.ResolveTCPAddr("tcp", c.ListenOn)
	if err != nil {
		return nil, err
	}
	options = append(options, server.WithServiceAddr(addr))

	//// middleware
	//options = append(options, server.WithMiddleware(middleware.CommonMiddleware))
	//options = append(options, server.WithMiddleware(middleware.ServerMiddleware))
	//
	//// limit
	//options = append(options, server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}))
	//
	//// Multiplex
	//options = append(options, server.WithMuxTransport())
	//
	//// BoundHandler
	//options = append(options, server.WithBoundHandler(bound.NewCpuLimitHandler()))

	var (
		server2 server.Server
	)

	// registry
	if c.HasEtcd() {
		if c.Etcd.GoZero {
			server2, err = NewRpcPubServer(c.Etcd.EtcdConf, c.ListenOn, options...)
		} else {
			// etcd
			r, err := etcd.NewEtcdRegistry(c.Etcd.Hosts) // r should not be reused.
			if err != nil {
				panic(err)
			}

			options = append(options, server.WithRegistry(r))

			server2 = server.NewServer(options...)
		}
	} else {
		server2 = server.NewServer(options...)
	}

	// server.WithLogger()
	if err = register(server2); err != nil {
		return nil, err
	}

	rpcServer := &RpcServer{
		Server: server2,
		// register: register,
	}

	return rpcServer, nil
}
