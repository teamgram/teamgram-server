/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package sessionservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SessionQueryAuthKey(ctx context.Context, req *session.TLSessionQueryAuthKey, callOptions ...callopt.Option) (r *tg.AuthKeyInfo, err error)
	SessionSetAuthKey(ctx context.Context, req *session.TLSessionSetAuthKey, callOptions ...callopt.Option) (r *tg.Bool, err error)
	SessionCreateSession(ctx context.Context, req *session.TLSessionCreateSession, callOptions ...callopt.Option) (r *tg.Bool, err error)
	SessionSendDataToSession(ctx context.Context, req *session.TLSessionSendDataToSession, callOptions ...callopt.Option) (r *tg.Bool, err error)
	SessionSendHttpDataToSession(ctx context.Context, req *session.TLSessionSendHttpDataToSession, callOptions ...callopt.Option) (r *session.HttpSessionData, err error)
	SessionCloseSession(ctx context.Context, req *session.TLSessionCloseSession, callOptions ...callopt.Option) (r *tg.Bool, err error)
	SessionPushUpdatesData(ctx context.Context, req *session.TLSessionPushUpdatesData, callOptions ...callopt.Option) (r *tg.Bool, err error)
	SessionPushSessionUpdatesData(ctx context.Context, req *session.TLSessionPushSessionUpdatesData, callOptions ...callopt.Option) (r *tg.Bool, err error)
	SessionPushRpcResultData(ctx context.Context, req *session.TLSessionPushRpcResultData, callOptions ...callopt.Option) (r *tg.Bool, err error)
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
	return &kSessionClient{
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

type kSessionClient struct {
	*kClient
}

func NewRPCSessionClient(cli client.Client) Client {
	return &kSessionClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kSessionClient) SessionQueryAuthKey(ctx context.Context, req *session.TLSessionQueryAuthKey, callOptions ...callopt.Option) (r *tg.AuthKeyInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SessionQueryAuthKey(ctx, req)
}

func (p *kSessionClient) SessionSetAuthKey(ctx context.Context, req *session.TLSessionSetAuthKey, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SessionSetAuthKey(ctx, req)
}

func (p *kSessionClient) SessionCreateSession(ctx context.Context, req *session.TLSessionCreateSession, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SessionCreateSession(ctx, req)
}

func (p *kSessionClient) SessionSendDataToSession(ctx context.Context, req *session.TLSessionSendDataToSession, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SessionSendDataToSession(ctx, req)
}

func (p *kSessionClient) SessionSendHttpDataToSession(ctx context.Context, req *session.TLSessionSendHttpDataToSession, callOptions ...callopt.Option) (r *session.HttpSessionData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SessionSendHttpDataToSession(ctx, req)
}

func (p *kSessionClient) SessionCloseSession(ctx context.Context, req *session.TLSessionCloseSession, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SessionCloseSession(ctx, req)
}

func (p *kSessionClient) SessionPushUpdatesData(ctx context.Context, req *session.TLSessionPushUpdatesData, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SessionPushUpdatesData(ctx, req)
}

func (p *kSessionClient) SessionPushSessionUpdatesData(ctx context.Context, req *session.TLSessionPushSessionUpdatesData, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SessionPushSessionUpdatesData(ctx, req)
}

func (p *kSessionClient) SessionPushRpcResultData(ctx context.Context, req *session.TLSessionPushRpcResultData, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SessionPushRpcResultData(ctx, req)
}
