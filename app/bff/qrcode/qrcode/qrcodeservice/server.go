/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package qrcodeservice

import (
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/server"
)

// Deprecated: prefer pkg/net/kitex.NewServer via the generated internal server bootstrap for TL-aware transport setup.
// NewServer creates a server.Server with the given handler and options.
func NewServer(handler tg.RPCQrCode, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, server.WithCodec(codec.NewZRpcCodec(false)))

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler tg.RPCQrCode, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
