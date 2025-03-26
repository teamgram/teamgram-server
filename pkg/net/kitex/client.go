// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package kitex

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/zeromicro/go-zero/core/logx"
)

type (
	// Client is an alias of internal.Client.
	Client = client.Client
	// ClientOption is an alias of internal.ClientOption.
	ClientOption = client.Option

	// A RpcClient is a rpc client.
	RpcClient struct {
		client Client
	}
)

type (
	// NewClientFn defines the method to create a client
	NewClientFn func(opts ...client.Option) (Client, error)
)

// MustNewClient returns a Client, exits on any error.
func MustNewClient(c RpcClientConf, newF NewClientFn) Client {
	cli, err := NewClient(c, newF)
	logx.Must(err)
	return cli
}

// NewClient returns a Client.
func NewClient(c RpcClientConf, newF NewClientFn) (Client, error) {
	var options []client.Option

	options = append(options, client.WithDestService(c.ServiceName))
	if c.Codec == "zrpc" {
		options = append(options, client.WithCodec(codec.NewZRpcCodec(true)))
	}

	// options = append(options, opts...)

	//// middleware
	//options = append(options, client.WithMiddleware(middleware.CommonMiddleware))
	//options = append(options, client.WithInstanceMW(middleware.ClientMiddleware))
	//
	//// mux
	//options = append(options, client.WithMuxConnection(1))
	//
	//// rpc timeout
	//options = append(options, client.WithRPCTimeout(3*time.Second))
	//
	//// conn timeout
	//options = append(options, client.WithConnectTimeout(50*time.Millisecond))
	//
	//// retry
	//options = append(options, client.WithFailureRetry(retry.NewFailurePolicy()))
	//
	//// tracer
	//options = append(options, client.WithSuite(trace.NewDefaultClientSuite()))

	// resolver
	if c.HasEtcd() {
		r, err := etcd.NewEtcdResolver(c.Etcd.Hosts)
		if err != nil {
			panic(err)
		}

		options = append(options, client.WithResolver(r))
	} else {
		options = append(options, client.WithHostPorts(c.Endpoints...))
	}

	cli, err := newF(options...)
	if err != nil {
		return nil, err
	}

	return &RpcClient{
		client: cli,
	}, nil
}

// Call returns the underlying grpc.ClientConn.
func (rc *RpcClient) Call(ctx context.Context, method string, request, response interface{}) error {
	return rc.client.Call(ctx, method, request, response)
}
