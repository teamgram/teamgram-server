/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package usernameclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/username/username"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/username/username/usernameservice"

	"github.com/cloudwego/kitex/client"
)

type UsernameClient interface {
	UsernameGetAccountUsername(ctx context.Context, in *username.TLUsernameGetAccountUsername) (*username.UsernameData, error)
	UsernameCheckAccountUsername(ctx context.Context, in *username.TLUsernameCheckAccountUsername) (*username.UsernameExist, error)
	UsernameGetChannelUsername(ctx context.Context, in *username.TLUsernameGetChannelUsername) (*username.UsernameData, error)
	UsernameCheckChannelUsername(ctx context.Context, in *username.TLUsernameCheckChannelUsername) (*username.UsernameExist, error)
	UsernameUpdateUsernameByPeer(ctx context.Context, in *username.TLUsernameUpdateUsernameByPeer) (*tg.Bool, error)
	UsernameCheckUsername(ctx context.Context, in *username.TLUsernameCheckUsername) (*username.UsernameExist, error)
	UsernameUpdateUsername(ctx context.Context, in *username.TLUsernameUpdateUsername) (*tg.Bool, error)
	UsernameDeleteUsername(ctx context.Context, in *username.TLUsernameDeleteUsername) (*tg.Bool, error)
	UsernameResolveUsername(ctx context.Context, in *username.TLUsernameResolveUsername) (*tg.Peer, error)
	UsernameGetListByUsernameList(ctx context.Context, in *username.TLUsernameGetListByUsernameList) (*username.VectorUsernameData, error)
	UsernameDeleteUsernameByPeer(ctx context.Context, in *username.TLUsernameDeleteUsernameByPeer) (*tg.Bool, error)
	UsernameSearch(ctx context.Context, in *username.TLUsernameSearch) (*username.VectorUsernameData, error)
}

type defaultUsernameClient struct {
	cli client.Client
}

func NewUsernameClient(cli client.Client) UsernameClient {
	return &defaultUsernameClient{
		cli: cli,
	}
}

// UsernameGetAccountUsername
// username.getAccountUsername user_id:long = UsernameData;
func (m *defaultUsernameClient) UsernameGetAccountUsername(ctx context.Context, in *username.TLUsernameGetAccountUsername) (*username.UsernameData, error) {
	cli := usernameservice.NewRPCUsernameClient(m.cli)
	return cli.UsernameGetAccountUsername(ctx, in)
}

// UsernameCheckAccountUsername
// username.checkAccountUsername user_id:long username:string = UsernameExist;
func (m *defaultUsernameClient) UsernameCheckAccountUsername(ctx context.Context, in *username.TLUsernameCheckAccountUsername) (*username.UsernameExist, error) {
	cli := usernameservice.NewRPCUsernameClient(m.cli)
	return cli.UsernameCheckAccountUsername(ctx, in)
}

// UsernameGetChannelUsername
// username.getChannelUsername channel_id:long = UsernameData;
func (m *defaultUsernameClient) UsernameGetChannelUsername(ctx context.Context, in *username.TLUsernameGetChannelUsername) (*username.UsernameData, error) {
	cli := usernameservice.NewRPCUsernameClient(m.cli)
	return cli.UsernameGetChannelUsername(ctx, in)
}

// UsernameCheckChannelUsername
// username.checkChannelUsername channel_id:long username:string = UsernameExist;
func (m *defaultUsernameClient) UsernameCheckChannelUsername(ctx context.Context, in *username.TLUsernameCheckChannelUsername) (*username.UsernameExist, error) {
	cli := usernameservice.NewRPCUsernameClient(m.cli)
	return cli.UsernameCheckChannelUsername(ctx, in)
}

// UsernameUpdateUsernameByPeer
// username.updateUsernameByPeer peer_type:int peer_id:long username:string = Bool;
func (m *defaultUsernameClient) UsernameUpdateUsernameByPeer(ctx context.Context, in *username.TLUsernameUpdateUsernameByPeer) (*tg.Bool, error) {
	cli := usernameservice.NewRPCUsernameClient(m.cli)
	return cli.UsernameUpdateUsernameByPeer(ctx, in)
}

// UsernameCheckUsername
// username.checkUsername username:string = UsernameExist;
func (m *defaultUsernameClient) UsernameCheckUsername(ctx context.Context, in *username.TLUsernameCheckUsername) (*username.UsernameExist, error) {
	cli := usernameservice.NewRPCUsernameClient(m.cli)
	return cli.UsernameCheckUsername(ctx, in)
}

// UsernameUpdateUsername
// username.updateUsername peer_type:int peer_id:long username:string = Bool;
func (m *defaultUsernameClient) UsernameUpdateUsername(ctx context.Context, in *username.TLUsernameUpdateUsername) (*tg.Bool, error) {
	cli := usernameservice.NewRPCUsernameClient(m.cli)
	return cli.UsernameUpdateUsername(ctx, in)
}

// UsernameDeleteUsername
// username.deleteUsername username:string = Bool;
func (m *defaultUsernameClient) UsernameDeleteUsername(ctx context.Context, in *username.TLUsernameDeleteUsername) (*tg.Bool, error) {
	cli := usernameservice.NewRPCUsernameClient(m.cli)
	return cli.UsernameDeleteUsername(ctx, in)
}

// UsernameResolveUsername
// username.resolveUsername username:string = Peer;
func (m *defaultUsernameClient) UsernameResolveUsername(ctx context.Context, in *username.TLUsernameResolveUsername) (*tg.Peer, error) {
	cli := usernameservice.NewRPCUsernameClient(m.cli)
	return cli.UsernameResolveUsername(ctx, in)
}

// UsernameGetListByUsernameList
// username.getListByUsernameList names:Vector<string> = Vector<UsernameData>;
func (m *defaultUsernameClient) UsernameGetListByUsernameList(ctx context.Context, in *username.TLUsernameGetListByUsernameList) (*username.VectorUsernameData, error) {
	cli := usernameservice.NewRPCUsernameClient(m.cli)
	return cli.UsernameGetListByUsernameList(ctx, in)
}

// UsernameDeleteUsernameByPeer
// username.deleteUsernameByPeer peer_type:int peer_id:long = Bool;
func (m *defaultUsernameClient) UsernameDeleteUsernameByPeer(ctx context.Context, in *username.TLUsernameDeleteUsernameByPeer) (*tg.Bool, error) {
	cli := usernameservice.NewRPCUsernameClient(m.cli)
	return cli.UsernameDeleteUsernameByPeer(ctx, in)
}

// UsernameSearch
// username.search q:string excluded_contacts:Vector<long> limit:int = Vector<UsernameData>;
func (m *defaultUsernameClient) UsernameSearch(ctx context.Context, in *username.TLUsernameSearch) (*username.VectorUsernameData, error) {
	cli := usernameservice.NewRPCUsernameClient(m.cli)
	return cli.UsernameSearch(ctx, in)
}
