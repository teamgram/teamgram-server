/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package usersservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UsersGetUsers(ctx context.Context, req *tg.TLUsersGetUsers, callOptions ...callopt.Option) (r *tg.VectorUser, err error)
	UsersGetFullUser(ctx context.Context, req *tg.TLUsersGetFullUser, callOptions ...callopt.Option) (r *tg.UsersUserFull, err error)
	ContactsResolvePhone(ctx context.Context, req *tg.TLContactsResolvePhone, callOptions ...callopt.Option) (r *tg.ContactsResolvedPeer, err error)
	UsersGetMe(ctx context.Context, req *tg.TLUsersGetMe, callOptions ...callopt.Option) (r *tg.User, err error)
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
	return &kUsersClient{
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

type kUsersClient struct {
	*kClient
}

func NewRPCUsersClient(cli client.Client) Client {
	return &kUsersClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kUsersClient) UsersGetUsers(ctx context.Context, req *tg.TLUsersGetUsers, callOptions ...callopt.Option) (r *tg.VectorUser, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsersGetUsers(ctx, req)
}

func (p *kUsersClient) UsersGetFullUser(ctx context.Context, req *tg.TLUsersGetFullUser, callOptions ...callopt.Option) (r *tg.UsersUserFull, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsersGetFullUser(ctx, req)
}

func (p *kUsersClient) ContactsResolvePhone(ctx context.Context, req *tg.TLContactsResolvePhone, callOptions ...callopt.Option) (r *tg.ContactsResolvedPeer, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsResolvePhone(ctx, req)
}

func (p *kUsersClient) UsersGetMe(ctx context.Context, req *tg.TLUsersGetMe, callOptions ...callopt.Option) (r *tg.User, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsersGetMe(ctx, req)
}
