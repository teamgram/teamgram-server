/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dialogsservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MessagesGetDialogs(ctx context.Context, req *tg.TLMessagesGetDialogs, callOptions ...callopt.Option) (r *tg.MessagesDialogs, err error)
	MessagesSetTyping(ctx context.Context, req *tg.TLMessagesSetTyping, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesGetPeerSettings(ctx context.Context, req *tg.TLMessagesGetPeerSettings, callOptions ...callopt.Option) (r *tg.MessagesPeerSettings, err error)
	MessagesGetPeerDialogs(ctx context.Context, req *tg.TLMessagesGetPeerDialogs, callOptions ...callopt.Option) (r *tg.MessagesPeerDialogs, err error)
	MessagesToggleDialogPin(ctx context.Context, req *tg.TLMessagesToggleDialogPin, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesReorderPinnedDialogs(ctx context.Context, req *tg.TLMessagesReorderPinnedDialogs, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesGetPinnedDialogs(ctx context.Context, req *tg.TLMessagesGetPinnedDialogs, callOptions ...callopt.Option) (r *tg.MessagesPeerDialogs, err error)
	MessagesSendScreenshotNotification(ctx context.Context, req *tg.TLMessagesSendScreenshotNotification, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesMarkDialogUnread(ctx context.Context, req *tg.TLMessagesMarkDialogUnread, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesGetDialogUnreadMarks(ctx context.Context, req *tg.TLMessagesGetDialogUnreadMarks, callOptions ...callopt.Option) (r *tg.VectorDialogPeer, err error)
	MessagesGetOnlines(ctx context.Context, req *tg.TLMessagesGetOnlines, callOptions ...callopt.Option) (r *tg.ChatOnlines, err error)
	MessagesHidePeerSettingsBar(ctx context.Context, req *tg.TLMessagesHidePeerSettingsBar, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesSetHistoryTTL(ctx context.Context, req *tg.TLMessagesSetHistoryTTL, callOptions ...callopt.Option) (r *tg.Updates, err error)
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
	return &kDialogsClient{
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

type kDialogsClient struct {
	*kClient
}

func NewRPCDialogsClient(cli client.Client) Client {
	return &kDialogsClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kDialogsClient) MessagesGetDialogs(ctx context.Context, req *tg.TLMessagesGetDialogs, callOptions ...callopt.Option) (r *tg.MessagesDialogs, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetDialogs(ctx, req)
}

func (p *kDialogsClient) MessagesSetTyping(ctx context.Context, req *tg.TLMessagesSetTyping, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesSetTyping(ctx, req)
}

func (p *kDialogsClient) MessagesGetPeerSettings(ctx context.Context, req *tg.TLMessagesGetPeerSettings, callOptions ...callopt.Option) (r *tg.MessagesPeerSettings, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetPeerSettings(ctx, req)
}

func (p *kDialogsClient) MessagesGetPeerDialogs(ctx context.Context, req *tg.TLMessagesGetPeerDialogs, callOptions ...callopt.Option) (r *tg.MessagesPeerDialogs, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetPeerDialogs(ctx, req)
}

func (p *kDialogsClient) MessagesToggleDialogPin(ctx context.Context, req *tg.TLMessagesToggleDialogPin, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesToggleDialogPin(ctx, req)
}

func (p *kDialogsClient) MessagesReorderPinnedDialogs(ctx context.Context, req *tg.TLMessagesReorderPinnedDialogs, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesReorderPinnedDialogs(ctx, req)
}

func (p *kDialogsClient) MessagesGetPinnedDialogs(ctx context.Context, req *tg.TLMessagesGetPinnedDialogs, callOptions ...callopt.Option) (r *tg.MessagesPeerDialogs, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetPinnedDialogs(ctx, req)
}

func (p *kDialogsClient) MessagesSendScreenshotNotification(ctx context.Context, req *tg.TLMessagesSendScreenshotNotification, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesSendScreenshotNotification(ctx, req)
}

func (p *kDialogsClient) MessagesMarkDialogUnread(ctx context.Context, req *tg.TLMessagesMarkDialogUnread, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesMarkDialogUnread(ctx, req)
}

func (p *kDialogsClient) MessagesGetDialogUnreadMarks(ctx context.Context, req *tg.TLMessagesGetDialogUnreadMarks, callOptions ...callopt.Option) (r *tg.VectorDialogPeer, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetDialogUnreadMarks(ctx, req)
}

func (p *kDialogsClient) MessagesGetOnlines(ctx context.Context, req *tg.TLMessagesGetOnlines, callOptions ...callopt.Option) (r *tg.ChatOnlines, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetOnlines(ctx, req)
}

func (p *kDialogsClient) MessagesHidePeerSettingsBar(ctx context.Context, req *tg.TLMessagesHidePeerSettingsBar, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesHidePeerSettingsBar(ctx, req)
}

func (p *kDialogsClient) MessagesSetHistoryTTL(ctx context.Context, req *tg.TLMessagesSetHistoryTTL, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesSetHistoryTTL(ctx, req)
}
