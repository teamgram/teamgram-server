/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package authsessionservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AuthsessionGetAuthorizations(ctx context.Context, req *authsession.TLAuthsessionGetAuthorizations, callOptions ...callopt.Option) (r *tg.AccountAuthorizations, err error)
	AuthsessionResetAuthorization(ctx context.Context, req *authsession.TLAuthsessionResetAuthorization, callOptions ...callopt.Option) (r *authsession.VectorLong, err error)
	AuthsessionGetLayer(ctx context.Context, req *authsession.TLAuthsessionGetLayer, callOptions ...callopt.Option) (r *tg.Int32, err error)
	AuthsessionGetLangPack(ctx context.Context, req *authsession.TLAuthsessionGetLangPack, callOptions ...callopt.Option) (r *tg.String, err error)
	AuthsessionGetClient(ctx context.Context, req *authsession.TLAuthsessionGetClient, callOptions ...callopt.Option) (r *tg.String, err error)
	AuthsessionGetLangCode(ctx context.Context, req *authsession.TLAuthsessionGetLangCode, callOptions ...callopt.Option) (r *tg.String, err error)
	AuthsessionGetUserId(ctx context.Context, req *authsession.TLAuthsessionGetUserId, callOptions ...callopt.Option) (r *tg.Int64, err error)
	AuthsessionGetPushSessionId(ctx context.Context, req *authsession.TLAuthsessionGetPushSessionId, callOptions ...callopt.Option) (r *tg.Int64, err error)
	AuthsessionGetFutureSalts(ctx context.Context, req *authsession.TLAuthsessionGetFutureSalts, callOptions ...callopt.Option) (r *tg.FutureSalts, err error)
	AuthsessionQueryAuthKey(ctx context.Context, req *authsession.TLAuthsessionQueryAuthKey, callOptions ...callopt.Option) (r *tg.AuthKeyInfo, err error)
	AuthsessionSetAuthKey(ctx context.Context, req *authsession.TLAuthsessionSetAuthKey, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthsessionBindAuthKeyUser(ctx context.Context, req *authsession.TLAuthsessionBindAuthKeyUser, callOptions ...callopt.Option) (r *tg.Int64, err error)
	AuthsessionUnbindAuthKeyUser(ctx context.Context, req *authsession.TLAuthsessionUnbindAuthKeyUser, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthsessionGetPermAuthKeyId(ctx context.Context, req *authsession.TLAuthsessionGetPermAuthKeyId, callOptions ...callopt.Option) (r *tg.Int64, err error)
	AuthsessionBindTempAuthKey(ctx context.Context, req *authsession.TLAuthsessionBindTempAuthKey, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthsessionSetClientSessionInfo(ctx context.Context, req *authsession.TLAuthsessionSetClientSessionInfo, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthsessionGetAuthorization(ctx context.Context, req *authsession.TLAuthsessionGetAuthorization, callOptions ...callopt.Option) (r *tg.Authorization, err error)
	AuthsessionGetAuthStateData(ctx context.Context, req *authsession.TLAuthsessionGetAuthStateData, callOptions ...callopt.Option) (r *authsession.AuthKeyStateData, err error)
	AuthsessionSetLayer(ctx context.Context, req *authsession.TLAuthsessionSetLayer, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthsessionSetInitConnection(ctx context.Context, req *authsession.TLAuthsessionSetInitConnection, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthsessionSetAndroidPushSessionId(ctx context.Context, req *authsession.TLAuthsessionSetAndroidPushSessionId, callOptions ...callopt.Option) (r *tg.Bool, err error)
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
	return &kAuthsessionClient{
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

type kAuthsessionClient struct {
	*kClient
}

func NewRPCAuthsessionClient(cli client.Client) Client {
	return &kAuthsessionClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kAuthsessionClient) AuthsessionGetAuthorizations(ctx context.Context, req *authsession.TLAuthsessionGetAuthorizations, callOptions ...callopt.Option) (r *tg.AccountAuthorizations, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionGetAuthorizations(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionResetAuthorization(ctx context.Context, req *authsession.TLAuthsessionResetAuthorization, callOptions ...callopt.Option) (r *authsession.VectorLong, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionResetAuthorization(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionGetLayer(ctx context.Context, req *authsession.TLAuthsessionGetLayer, callOptions ...callopt.Option) (r *tg.Int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionGetLayer(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionGetLangPack(ctx context.Context, req *authsession.TLAuthsessionGetLangPack, callOptions ...callopt.Option) (r *tg.String, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionGetLangPack(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionGetClient(ctx context.Context, req *authsession.TLAuthsessionGetClient, callOptions ...callopt.Option) (r *tg.String, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionGetClient(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionGetLangCode(ctx context.Context, req *authsession.TLAuthsessionGetLangCode, callOptions ...callopt.Option) (r *tg.String, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionGetLangCode(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionGetUserId(ctx context.Context, req *authsession.TLAuthsessionGetUserId, callOptions ...callopt.Option) (r *tg.Int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionGetUserId(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionGetPushSessionId(ctx context.Context, req *authsession.TLAuthsessionGetPushSessionId, callOptions ...callopt.Option) (r *tg.Int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionGetPushSessionId(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionGetFutureSalts(ctx context.Context, req *authsession.TLAuthsessionGetFutureSalts, callOptions ...callopt.Option) (r *tg.FutureSalts, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionGetFutureSalts(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionQueryAuthKey(ctx context.Context, req *authsession.TLAuthsessionQueryAuthKey, callOptions ...callopt.Option) (r *tg.AuthKeyInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionQueryAuthKey(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionSetAuthKey(ctx context.Context, req *authsession.TLAuthsessionSetAuthKey, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionSetAuthKey(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionBindAuthKeyUser(ctx context.Context, req *authsession.TLAuthsessionBindAuthKeyUser, callOptions ...callopt.Option) (r *tg.Int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionBindAuthKeyUser(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionUnbindAuthKeyUser(ctx context.Context, req *authsession.TLAuthsessionUnbindAuthKeyUser, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionUnbindAuthKeyUser(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionGetPermAuthKeyId(ctx context.Context, req *authsession.TLAuthsessionGetPermAuthKeyId, callOptions ...callopt.Option) (r *tg.Int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionGetPermAuthKeyId(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionBindTempAuthKey(ctx context.Context, req *authsession.TLAuthsessionBindTempAuthKey, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionBindTempAuthKey(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionSetClientSessionInfo(ctx context.Context, req *authsession.TLAuthsessionSetClientSessionInfo, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionSetClientSessionInfo(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionGetAuthorization(ctx context.Context, req *authsession.TLAuthsessionGetAuthorization, callOptions ...callopt.Option) (r *tg.Authorization, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionGetAuthorization(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionGetAuthStateData(ctx context.Context, req *authsession.TLAuthsessionGetAuthStateData, callOptions ...callopt.Option) (r *authsession.AuthKeyStateData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionGetAuthStateData(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionSetLayer(ctx context.Context, req *authsession.TLAuthsessionSetLayer, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionSetLayer(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionSetInitConnection(ctx context.Context, req *authsession.TLAuthsessionSetInitConnection, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionSetInitConnection(ctx, req)
}

func (p *kAuthsessionClient) AuthsessionSetAndroidPushSessionId(ctx context.Context, req *authsession.TLAuthsessionSetAndroidPushSessionId, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthsessionSetAndroidPushSessionId(ctx, req)
}
