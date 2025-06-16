/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package codeservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CodeCreatePhoneCode(ctx context.Context, req *code.TLCodeCreatePhoneCode, callOptions ...callopt.Option) (r *code.PhoneCodeTransaction, err error)
	CodeGetPhoneCode(ctx context.Context, req *code.TLCodeGetPhoneCode, callOptions ...callopt.Option) (r *code.PhoneCodeTransaction, err error)
	CodeDeletePhoneCode(ctx context.Context, req *code.TLCodeDeletePhoneCode, callOptions ...callopt.Option) (r *tg.Bool, err error)
	CodeUpdatePhoneCodeData(ctx context.Context, req *code.TLCodeUpdatePhoneCodeData, callOptions ...callopt.Option) (r *tg.Bool, err error)
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
	return &kCodeClient{
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

type kCodeClient struct {
	*kClient
}

func NewRPCCodeClient(cli client.Client) Client {
	return &kCodeClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kCodeClient) CodeCreatePhoneCode(ctx context.Context, req *code.TLCodeCreatePhoneCode, callOptions ...callopt.Option) (r *code.PhoneCodeTransaction, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CodeCreatePhoneCode(ctx, req)
}

func (p *kCodeClient) CodeGetPhoneCode(ctx context.Context, req *code.TLCodeGetPhoneCode, callOptions ...callopt.Option) (r *code.PhoneCodeTransaction, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CodeGetPhoneCode(ctx, req)
}

func (p *kCodeClient) CodeDeletePhoneCode(ctx context.Context, req *code.TLCodeDeletePhoneCode, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CodeDeletePhoneCode(ctx, req)
}

func (p *kCodeClient) CodeUpdatePhoneCodeData(ctx context.Context, req *code.TLCodeUpdatePhoneCodeData, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CodeUpdatePhoneCodeData(ctx, req)
}
