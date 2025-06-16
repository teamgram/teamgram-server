/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package idgenservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	IdgenNextId(ctx context.Context, req *idgen.TLIdgenNextId, callOptions ...callopt.Option) (r *tg.Int64, err error)
	IdgenNextIds(ctx context.Context, req *idgen.TLIdgenNextIds, callOptions ...callopt.Option) (r *idgen.VectorLong, err error)
	IdgenGetCurrentSeqId(ctx context.Context, req *idgen.TLIdgenGetCurrentSeqId, callOptions ...callopt.Option) (r *tg.Int64, err error)
	IdgenSetCurrentSeqId(ctx context.Context, req *idgen.TLIdgenSetCurrentSeqId, callOptions ...callopt.Option) (r *tg.Bool, err error)
	IdgenGetNextSeqId(ctx context.Context, req *idgen.TLIdgenGetNextSeqId, callOptions ...callopt.Option) (r *tg.Int64, err error)
	IdgenGetNextNSeqId(ctx context.Context, req *idgen.TLIdgenGetNextNSeqId, callOptions ...callopt.Option) (r *tg.Int64, err error)
	IdgenGetNextIdValList(ctx context.Context, req *idgen.TLIdgenGetNextIdValList, callOptions ...callopt.Option) (r *idgen.VectorIdVal, err error)
	IdgenGetCurrentSeqIdList(ctx context.Context, req *idgen.TLIdgenGetCurrentSeqIdList, callOptions ...callopt.Option) (r *idgen.VectorIdVal, err error)
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
	return &kIdgenClient{
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

type kIdgenClient struct {
	*kClient
}

func NewRPCIdgenClient(cli client.Client) Client {
	return &kIdgenClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kIdgenClient) IdgenNextId(ctx context.Context, req *idgen.TLIdgenNextId, callOptions ...callopt.Option) (r *tg.Int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IdgenNextId(ctx, req)
}

func (p *kIdgenClient) IdgenNextIds(ctx context.Context, req *idgen.TLIdgenNextIds, callOptions ...callopt.Option) (r *idgen.VectorLong, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IdgenNextIds(ctx, req)
}

func (p *kIdgenClient) IdgenGetCurrentSeqId(ctx context.Context, req *idgen.TLIdgenGetCurrentSeqId, callOptions ...callopt.Option) (r *tg.Int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IdgenGetCurrentSeqId(ctx, req)
}

func (p *kIdgenClient) IdgenSetCurrentSeqId(ctx context.Context, req *idgen.TLIdgenSetCurrentSeqId, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IdgenSetCurrentSeqId(ctx, req)
}

func (p *kIdgenClient) IdgenGetNextSeqId(ctx context.Context, req *idgen.TLIdgenGetNextSeqId, callOptions ...callopt.Option) (r *tg.Int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IdgenGetNextSeqId(ctx, req)
}

func (p *kIdgenClient) IdgenGetNextNSeqId(ctx context.Context, req *idgen.TLIdgenGetNextNSeqId, callOptions ...callopt.Option) (r *tg.Int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IdgenGetNextNSeqId(ctx, req)
}

func (p *kIdgenClient) IdgenGetNextIdValList(ctx context.Context, req *idgen.TLIdgenGetNextIdValList, callOptions ...callopt.Option) (r *idgen.VectorIdVal, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IdgenGetNextIdValList(ctx, req)
}

func (p *kIdgenClient) IdgenGetCurrentSeqIdList(ctx context.Context, req *idgen.TLIdgenGetCurrentSeqIdList, callOptions ...callopt.Option) (r *idgen.VectorIdVal, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IdgenGetCurrentSeqIdList(ctx, req)
}
