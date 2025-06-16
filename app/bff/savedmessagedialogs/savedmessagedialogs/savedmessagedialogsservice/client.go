/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package savedmessagedialogsservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MessagesGetSavedDialogs(ctx context.Context, req *tg.TLMessagesGetSavedDialogs, callOptions ...callopt.Option) (r *tg.MessagesSavedDialogs, err error)
	MessagesGetSavedHistory(ctx context.Context, req *tg.TLMessagesGetSavedHistory, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error)
	MessagesDeleteSavedHistory(ctx context.Context, req *tg.TLMessagesDeleteSavedHistory, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error)
	MessagesGetPinnedSavedDialogs(ctx context.Context, req *tg.TLMessagesGetPinnedSavedDialogs, callOptions ...callopt.Option) (r *tg.MessagesSavedDialogs, err error)
	MessagesToggleSavedDialogPin(ctx context.Context, req *tg.TLMessagesToggleSavedDialogPin, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesReorderPinnedSavedDialogs(ctx context.Context, req *tg.TLMessagesReorderPinnedSavedDialogs, callOptions ...callopt.Option) (r *tg.Bool, err error)
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
	return &kSavedMessageDialogsClient{
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

type kSavedMessageDialogsClient struct {
	*kClient
}

func NewRPCSavedMessageDialogsClient(cli client.Client) Client {
	return &kSavedMessageDialogsClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kSavedMessageDialogsClient) MessagesGetSavedDialogs(ctx context.Context, req *tg.TLMessagesGetSavedDialogs, callOptions ...callopt.Option) (r *tg.MessagesSavedDialogs, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetSavedDialogs(ctx, req)
}

func (p *kSavedMessageDialogsClient) MessagesGetSavedHistory(ctx context.Context, req *tg.TLMessagesGetSavedHistory, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetSavedHistory(ctx, req)
}

func (p *kSavedMessageDialogsClient) MessagesDeleteSavedHistory(ctx context.Context, req *tg.TLMessagesDeleteSavedHistory, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesDeleteSavedHistory(ctx, req)
}

func (p *kSavedMessageDialogsClient) MessagesGetPinnedSavedDialogs(ctx context.Context, req *tg.TLMessagesGetPinnedSavedDialogs, callOptions ...callopt.Option) (r *tg.MessagesSavedDialogs, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetPinnedSavedDialogs(ctx, req)
}

func (p *kSavedMessageDialogsClient) MessagesToggleSavedDialogPin(ctx context.Context, req *tg.TLMessagesToggleSavedDialogPin, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesToggleSavedDialogPin(ctx, req)
}

func (p *kSavedMessageDialogsClient) MessagesReorderPinnedSavedDialogs(ctx context.Context, req *tg.TLMessagesReorderPinnedSavedDialogs, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesReorderPinnedSavedDialogs(ctx, req)
}
