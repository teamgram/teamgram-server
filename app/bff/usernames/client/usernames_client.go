/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package usernamesclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/usernames/usernames/usernamesservice"

	"github.com/cloudwego/kitex/client"
)

type UsernamesClient interface {
	AccountCheckUsername(ctx context.Context, in *tg.TLAccountCheckUsername) (*tg.Bool, error)
	AccountUpdateUsername(ctx context.Context, in *tg.TLAccountUpdateUsername) (*tg.User, error)
	ContactsResolveUsername(ctx context.Context, in *tg.TLContactsResolveUsername) (*tg.ContactsResolvedPeer, error)
	ChannelsCheckUsername(ctx context.Context, in *tg.TLChannelsCheckUsername) (*tg.Bool, error)
	ChannelsUpdateUsername(ctx context.Context, in *tg.TLChannelsUpdateUsername) (*tg.Bool, error)
}

type defaultUsernamesClient struct {
	cli client.Client
}

func NewUsernamesClient(cli client.Client) UsernamesClient {
	return &defaultUsernamesClient{
		cli: cli,
	}
}

// AccountCheckUsername
// account.checkUsername#2714d86c username:string = Bool;
func (m *defaultUsernamesClient) AccountCheckUsername(ctx context.Context, in *tg.TLAccountCheckUsername) (*tg.Bool, error) {
	cli := usernamesservice.NewRPCUsernamesClient(m.cli)
	return cli.AccountCheckUsername(ctx, in)
}

// AccountUpdateUsername
// account.updateUsername#3e0bdd7c username:string = User;
func (m *defaultUsernamesClient) AccountUpdateUsername(ctx context.Context, in *tg.TLAccountUpdateUsername) (*tg.User, error) {
	cli := usernamesservice.NewRPCUsernamesClient(m.cli)
	return cli.AccountUpdateUsername(ctx, in)
}

// ContactsResolveUsername
// contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;
func (m *defaultUsernamesClient) ContactsResolveUsername(ctx context.Context, in *tg.TLContactsResolveUsername) (*tg.ContactsResolvedPeer, error) {
	cli := usernamesservice.NewRPCUsernamesClient(m.cli)
	return cli.ContactsResolveUsername(ctx, in)
}

// ChannelsCheckUsername
// channels.checkUsername#10e6bd2c channel:InputChannel username:string = Bool;
func (m *defaultUsernamesClient) ChannelsCheckUsername(ctx context.Context, in *tg.TLChannelsCheckUsername) (*tg.Bool, error) {
	cli := usernamesservice.NewRPCUsernamesClient(m.cli)
	return cli.ChannelsCheckUsername(ctx, in)
}

// ChannelsUpdateUsername
// channels.updateUsername#3514b3de channel:InputChannel username:string = Bool;
func (m *defaultUsernamesClient) ChannelsUpdateUsername(ctx context.Context, in *tg.TLChannelsUpdateUsername) (*tg.Bool, error) {
	cli := usernamesservice.NewRPCUsernamesClient(m.cli)
	return cli.ChannelsUpdateUsername(ctx, in)
}
