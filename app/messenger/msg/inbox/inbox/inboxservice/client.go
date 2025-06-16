/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package inboxservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/inbox/inbox"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	InboxEditUserMessageToInbox(ctx context.Context, req *inbox.TLInboxEditUserMessageToInbox, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxEditChatMessageToInbox(ctx context.Context, req *inbox.TLInboxEditChatMessageToInbox, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxDeleteMessagesToInbox(ctx context.Context, req *inbox.TLInboxDeleteMessagesToInbox, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxDeleteUserHistoryToInbox(ctx context.Context, req *inbox.TLInboxDeleteUserHistoryToInbox, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxDeleteChatHistoryToInbox(ctx context.Context, req *inbox.TLInboxDeleteChatHistoryToInbox, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxReadUserMediaUnreadToInbox(ctx context.Context, req *inbox.TLInboxReadUserMediaUnreadToInbox, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxReadChatMediaUnreadToInbox(ctx context.Context, req *inbox.TLInboxReadChatMediaUnreadToInbox, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxUpdateHistoryReaded(ctx context.Context, req *inbox.TLInboxUpdateHistoryReaded, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxUpdatePinnedMessage(ctx context.Context, req *inbox.TLInboxUpdatePinnedMessage, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxUnpinAllMessages(ctx context.Context, req *inbox.TLInboxUnpinAllMessages, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxSendUserMessageToInboxV2(ctx context.Context, req *inbox.TLInboxSendUserMessageToInboxV2, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxEditMessageToInboxV2(ctx context.Context, req *inbox.TLInboxEditMessageToInboxV2, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxReadInboxHistory(ctx context.Context, req *inbox.TLInboxReadInboxHistory, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxReadOutboxHistory(ctx context.Context, req *inbox.TLInboxReadOutboxHistory, callOptions ...callopt.Option) (r *tg.Void, err error)
	InboxReadMediaUnreadToInboxV2(ctx context.Context, req *inbox.TLInboxReadMediaUnreadToInboxV2, callOptions ...callopt.Option) (r *tg.Void, err error)
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
	return &kInboxClient{
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

type kInboxClient struct {
	*kClient
}

func NewRPCInboxClient(cli client.Client) Client {
	return &kInboxClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kInboxClient) InboxEditUserMessageToInbox(ctx context.Context, req *inbox.TLInboxEditUserMessageToInbox, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxEditUserMessageToInbox(ctx, req)
}

func (p *kInboxClient) InboxEditChatMessageToInbox(ctx context.Context, req *inbox.TLInboxEditChatMessageToInbox, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxEditChatMessageToInbox(ctx, req)
}

func (p *kInboxClient) InboxDeleteMessagesToInbox(ctx context.Context, req *inbox.TLInboxDeleteMessagesToInbox, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxDeleteMessagesToInbox(ctx, req)
}

func (p *kInboxClient) InboxDeleteUserHistoryToInbox(ctx context.Context, req *inbox.TLInboxDeleteUserHistoryToInbox, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxDeleteUserHistoryToInbox(ctx, req)
}

func (p *kInboxClient) InboxDeleteChatHistoryToInbox(ctx context.Context, req *inbox.TLInboxDeleteChatHistoryToInbox, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxDeleteChatHistoryToInbox(ctx, req)
}

func (p *kInboxClient) InboxReadUserMediaUnreadToInbox(ctx context.Context, req *inbox.TLInboxReadUserMediaUnreadToInbox, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxReadUserMediaUnreadToInbox(ctx, req)
}

func (p *kInboxClient) InboxReadChatMediaUnreadToInbox(ctx context.Context, req *inbox.TLInboxReadChatMediaUnreadToInbox, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxReadChatMediaUnreadToInbox(ctx, req)
}

func (p *kInboxClient) InboxUpdateHistoryReaded(ctx context.Context, req *inbox.TLInboxUpdateHistoryReaded, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxUpdateHistoryReaded(ctx, req)
}

func (p *kInboxClient) InboxUpdatePinnedMessage(ctx context.Context, req *inbox.TLInboxUpdatePinnedMessage, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxUpdatePinnedMessage(ctx, req)
}

func (p *kInboxClient) InboxUnpinAllMessages(ctx context.Context, req *inbox.TLInboxUnpinAllMessages, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxUnpinAllMessages(ctx, req)
}

func (p *kInboxClient) InboxSendUserMessageToInboxV2(ctx context.Context, req *inbox.TLInboxSendUserMessageToInboxV2, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxSendUserMessageToInboxV2(ctx, req)
}

func (p *kInboxClient) InboxEditMessageToInboxV2(ctx context.Context, req *inbox.TLInboxEditMessageToInboxV2, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxEditMessageToInboxV2(ctx, req)
}

func (p *kInboxClient) InboxReadInboxHistory(ctx context.Context, req *inbox.TLInboxReadInboxHistory, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxReadInboxHistory(ctx, req)
}

func (p *kInboxClient) InboxReadOutboxHistory(ctx context.Context, req *inbox.TLInboxReadOutboxHistory, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxReadOutboxHistory(ctx, req)
}

func (p *kInboxClient) InboxReadMediaUnreadToInboxV2(ctx context.Context, req *inbox.TLInboxReadMediaUnreadToInboxV2, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InboxReadMediaUnreadToInboxV2(ctx, req)
}
