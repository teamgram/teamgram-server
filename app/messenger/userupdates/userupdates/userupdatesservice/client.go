/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package userupdatesservice

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UserupdatesProcessUserOperation(ctx context.Context, req *userupdates.TLUserupdatesProcessUserOperation, callOptions ...callopt.Option) (r *userupdates.UserOperationResult, err error)
	UserupdatesGetOperationResult(ctx context.Context, req *userupdates.TLUserupdatesGetOperationResult, callOptions ...callopt.Option) (r *userupdates.UserOperationResult, err error)
	UserupdatesGetState(ctx context.Context, req *userupdates.TLUserupdatesGetState, callOptions ...callopt.Option) (r *userupdates.UserState, err error)
	UserupdatesGetDifference(ctx context.Context, req *userupdates.TLUserupdatesGetDifference, callOptions ...callopt.Option) (r *userupdates.UserDifference, err error)
	UserupdatesListDialogs(ctx context.Context, req *userupdates.TLUserupdatesListDialogs, callOptions ...callopt.Option) (r *userupdates.DialogProjectionList, err error)
	UserupdatesGetDialogsByPeers(ctx context.Context, req *userupdates.TLUserupdatesGetDialogsByPeers, callOptions ...callopt.Option) (r *userupdates.VectorDialogProjection, err error)
	UserupdatesGetDialogCount(ctx context.Context, req *userupdates.TLUserupdatesGetDialogCount, callOptions ...callopt.Option) (r *tg.Int32, err error)
	UserupdatesGetMessageViewsByPeerSeqs(ctx context.Context, req *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs, callOptions ...callopt.Option) (r *userupdates.MessageViewList, err error)
	UserupdatesGetOutboxReadDate(ctx context.Context, req *userupdates.TLUserupdatesGetOutboxReadDate, callOptions ...callopt.Option) (r *tg.OutboxReadDate, err error)
	UserupdatesAppendDialogAuthSeqSideEffect(ctx context.Context, req *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect, callOptions ...callopt.Option) (r *userupdates.UserAuthSeqAppendResult, err error)
	UserupdatesAppendDialogPtsSideEffect(ctx context.Context, req *userupdates.TLUserupdatesAppendDialogPtsSideEffect, callOptions ...callopt.Option) (r *userupdates.UserPtsAppendResult, err error)
}

// Deprecated: prefer the generated app client helper or pkg/net/kitex.NewClient for TL-aware transport setup.
// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithCodec(codec.NewZRpcCodec(false)))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserupdatesClient{
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

type kUserupdatesClient struct {
	*kClient
}

func NewRPCUserupdatesClient(cli client.Client) Client {
	return &kUserupdatesClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kUserupdatesClient) UserupdatesProcessUserOperation(ctx context.Context, req *userupdates.TLUserupdatesProcessUserOperation, callOptions ...callopt.Option) (r *userupdates.UserOperationResult, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserupdatesProcessUserOperation(ctx, req)
}

func (p *kUserupdatesClient) UserupdatesGetOperationResult(ctx context.Context, req *userupdates.TLUserupdatesGetOperationResult, callOptions ...callopt.Option) (r *userupdates.UserOperationResult, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserupdatesGetOperationResult(ctx, req)
}

func (p *kUserupdatesClient) UserupdatesGetState(ctx context.Context, req *userupdates.TLUserupdatesGetState, callOptions ...callopt.Option) (r *userupdates.UserState, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserupdatesGetState(ctx, req)
}

func (p *kUserupdatesClient) UserupdatesGetDifference(ctx context.Context, req *userupdates.TLUserupdatesGetDifference, callOptions ...callopt.Option) (r *userupdates.UserDifference, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserupdatesGetDifference(ctx, req)
}

func (p *kUserupdatesClient) UserupdatesListDialogs(ctx context.Context, req *userupdates.TLUserupdatesListDialogs, callOptions ...callopt.Option) (r *userupdates.DialogProjectionList, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserupdatesListDialogs(ctx, req)
}

func (p *kUserupdatesClient) UserupdatesGetDialogsByPeers(ctx context.Context, req *userupdates.TLUserupdatesGetDialogsByPeers, callOptions ...callopt.Option) (r *userupdates.VectorDialogProjection, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserupdatesGetDialogsByPeers(ctx, req)
}

func (p *kUserupdatesClient) UserupdatesGetDialogCount(ctx context.Context, req *userupdates.TLUserupdatesGetDialogCount, callOptions ...callopt.Option) (r *tg.Int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserupdatesGetDialogCount(ctx, req)
}

func (p *kUserupdatesClient) UserupdatesGetMessageViewsByPeerSeqs(ctx context.Context, req *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs, callOptions ...callopt.Option) (r *userupdates.MessageViewList, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserupdatesGetMessageViewsByPeerSeqs(ctx, req)
}

func (p *kUserupdatesClient) UserupdatesGetOutboxReadDate(ctx context.Context, req *userupdates.TLUserupdatesGetOutboxReadDate, callOptions ...callopt.Option) (r *tg.OutboxReadDate, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserupdatesGetOutboxReadDate(ctx, req)
}

func (p *kUserupdatesClient) UserupdatesAppendDialogAuthSeqSideEffect(ctx context.Context, req *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect, callOptions ...callopt.Option) (r *userupdates.UserAuthSeqAppendResult, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserupdatesAppendDialogAuthSeqSideEffect(ctx, req)
}

func (p *kUserupdatesClient) UserupdatesAppendDialogPtsSideEffect(ctx context.Context, req *userupdates.TLUserupdatesAppendDialogPtsSideEffect, callOptions ...callopt.Option) (r *userupdates.UserPtsAppendResult, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserupdatesAppendDialogPtsSideEffect(ctx, req)
}
