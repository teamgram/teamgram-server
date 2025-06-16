/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package qrcodeservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AuthExportLoginToken(ctx context.Context, req *tg.TLAuthExportLoginToken, callOptions ...callopt.Option) (r *tg.AuthLoginToken, err error)
	AuthImportLoginToken(ctx context.Context, req *tg.TLAuthImportLoginToken, callOptions ...callopt.Option) (r *tg.AuthLoginToken, err error)
	AuthAcceptLoginToken(ctx context.Context, req *tg.TLAuthAcceptLoginToken, callOptions ...callopt.Option) (r *tg.Authorization, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kQrCodeClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kQrCodeClient struct {
	*kClient
}

func NewRPCQrCodeClient(cli client.Client) Client {
	return &kQrCodeClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kQrCodeClient) AuthExportLoginToken(ctx context.Context, req *tg.TLAuthExportLoginToken, callOptions ...callopt.Option) (r *tg.AuthLoginToken, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthExportLoginToken(ctx, req)
}

func (p *kQrCodeClient) AuthImportLoginToken(ctx context.Context, req *tg.TLAuthImportLoginToken, callOptions ...callopt.Option) (r *tg.AuthLoginToken, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthImportLoginToken(ctx, req)
}

func (p *kQrCodeClient) AuthAcceptLoginToken(ctx context.Context, req *tg.TLAuthAcceptLoginToken, callOptions ...callopt.Option) (r *tg.Authorization, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthAcceptLoginToken(ctx, req)
}
