/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package usernames_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type UsernamesClient interface {
	AccountCheckUsername(ctx context.Context, in *mtproto.TLAccountCheckUsername) (*mtproto.Bool, error)
	AccountUpdateUsername(ctx context.Context, in *mtproto.TLAccountUpdateUsername) (*mtproto.User, error)
	ContactsResolveUsername(ctx context.Context, in *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error)
	ChannelsCheckUsername(ctx context.Context, in *mtproto.TLChannelsCheckUsername) (*mtproto.Bool, error)
	ChannelsUpdateUsername(ctx context.Context, in *mtproto.TLChannelsUpdateUsername) (*mtproto.Bool, error)
}

type defaultUsernamesClient struct {
	cli zrpc.Client
}

func NewUsernamesClient(cli zrpc.Client) UsernamesClient {
	return &defaultUsernamesClient{
		cli: cli,
	}
}

// AccountCheckUsername
// account.checkUsername#2714d86c username:string = Bool;
func (m *defaultUsernamesClient) AccountCheckUsername(ctx context.Context, in *mtproto.TLAccountCheckUsername) (*mtproto.Bool, error) {
	client := mtproto.NewRPCUsernamesClient(m.cli.Conn())
	return client.AccountCheckUsername(ctx, in)
}

// AccountUpdateUsername
// account.updateUsername#3e0bdd7c username:string = User;
func (m *defaultUsernamesClient) AccountUpdateUsername(ctx context.Context, in *mtproto.TLAccountUpdateUsername) (*mtproto.User, error) {
	client := mtproto.NewRPCUsernamesClient(m.cli.Conn())
	return client.AccountUpdateUsername(ctx, in)
}

// ContactsResolveUsername
// contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;
func (m *defaultUsernamesClient) ContactsResolveUsername(ctx context.Context, in *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	client := mtproto.NewRPCUsernamesClient(m.cli.Conn())
	return client.ContactsResolveUsername(ctx, in)
}

// ChannelsCheckUsername
// channels.checkUsername#10e6bd2c channel:InputChannel username:string = Bool;
func (m *defaultUsernamesClient) ChannelsCheckUsername(ctx context.Context, in *mtproto.TLChannelsCheckUsername) (*mtproto.Bool, error) {
	client := mtproto.NewRPCUsernamesClient(m.cli.Conn())
	return client.ChannelsCheckUsername(ctx, in)
}

// ChannelsUpdateUsername
// channels.updateUsername#3514b3de channel:InputChannel username:string = Bool;
func (m *defaultUsernamesClient) ChannelsUpdateUsername(ctx context.Context, in *mtproto.TLChannelsUpdateUsername) (*mtproto.Bool, error) {
	client := mtproto.NewRPCUsernamesClient(m.cli.Conn())
	return client.ChannelsUpdateUsername(ctx, in)
}
