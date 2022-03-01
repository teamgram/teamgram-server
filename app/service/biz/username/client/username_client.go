/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package username_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type UsernameClient interface {
	UsernameGetAccountUsername(ctx context.Context, in *username.TLUsernameGetAccountUsername) (*username.UsernameData, error)
	UsernameCheckAccountUsername(ctx context.Context, in *username.TLUsernameCheckAccountUsername) (*username.UsernameExist, error)
	UsernameGetChannelUsername(ctx context.Context, in *username.TLUsernameGetChannelUsername) (*username.UsernameData, error)
	UsernameCheckChannelUsername(ctx context.Context, in *username.TLUsernameCheckChannelUsername) (*username.UsernameExist, error)
	UsernameUpdateUsernameByPeer(ctx context.Context, in *username.TLUsernameUpdateUsernameByPeer) (*mtproto.Bool, error)
	UsernameCheckUsername(ctx context.Context, in *username.TLUsernameCheckUsername) (*username.UsernameExist, error)
	UsernameUpdateUsername(ctx context.Context, in *username.TLUsernameUpdateUsername) (*mtproto.Bool, error)
	UsernameDeleteUsername(ctx context.Context, in *username.TLUsernameDeleteUsername) (*mtproto.Bool, error)
	UsernameResolveUsername(ctx context.Context, in *username.TLUsernameResolveUsername) (*mtproto.Peer, error)
	UsernameGetListByUsernameList(ctx context.Context, in *username.TLUsernameGetListByUsernameList) (*username.Vector_UsernameData, error)
	UsernameDeleteUsernameByPeer(ctx context.Context, in *username.TLUsernameDeleteUsernameByPeer) (*mtproto.Bool, error)
	UsernameSearch(ctx context.Context, in *username.TLUsernameSearch) (*username.Vector_UsernameData, error)
}

type defaultUsernameClient struct {
	cli zrpc.Client
}

func NewUsernameClient(cli zrpc.Client) UsernameClient {
	return &defaultUsernameClient{
		cli: cli,
	}
}

// UsernameGetAccountUsername
// username.getAccountUsername user_id:long = UsernameData;
func (m *defaultUsernameClient) UsernameGetAccountUsername(ctx context.Context, in *username.TLUsernameGetAccountUsername) (*username.UsernameData, error) {
	client := username.NewRPCUsernameClient(m.cli.Conn())
	return client.UsernameGetAccountUsername(ctx, in)
}

// UsernameCheckAccountUsername
// username.checkAccountUsername user_id:long username:string = UsernameExist;
func (m *defaultUsernameClient) UsernameCheckAccountUsername(ctx context.Context, in *username.TLUsernameCheckAccountUsername) (*username.UsernameExist, error) {
	client := username.NewRPCUsernameClient(m.cli.Conn())
	return client.UsernameCheckAccountUsername(ctx, in)
}

// UsernameGetChannelUsername
// username.getChannelUsername channel_id:long = UsernameData;
func (m *defaultUsernameClient) UsernameGetChannelUsername(ctx context.Context, in *username.TLUsernameGetChannelUsername) (*username.UsernameData, error) {
	client := username.NewRPCUsernameClient(m.cli.Conn())
	return client.UsernameGetChannelUsername(ctx, in)
}

// UsernameCheckChannelUsername
// username.checkChannelUsername channel_id:long username:string = UsernameExist;
func (m *defaultUsernameClient) UsernameCheckChannelUsername(ctx context.Context, in *username.TLUsernameCheckChannelUsername) (*username.UsernameExist, error) {
	client := username.NewRPCUsernameClient(m.cli.Conn())
	return client.UsernameCheckChannelUsername(ctx, in)
}

// UsernameUpdateUsernameByPeer
// username.updateUsernameByPeer peer_type:int peer_id:long username:string = Bool;
func (m *defaultUsernameClient) UsernameUpdateUsernameByPeer(ctx context.Context, in *username.TLUsernameUpdateUsernameByPeer) (*mtproto.Bool, error) {
	client := username.NewRPCUsernameClient(m.cli.Conn())
	return client.UsernameUpdateUsernameByPeer(ctx, in)
}

// UsernameCheckUsername
// username.checkUsername username:string = UsernameExist;
func (m *defaultUsernameClient) UsernameCheckUsername(ctx context.Context, in *username.TLUsernameCheckUsername) (*username.UsernameExist, error) {
	client := username.NewRPCUsernameClient(m.cli.Conn())
	return client.UsernameCheckUsername(ctx, in)
}

// UsernameUpdateUsername
// username.updateUsername peer_type:int peer_id:long username:string = Bool;
func (m *defaultUsernameClient) UsernameUpdateUsername(ctx context.Context, in *username.TLUsernameUpdateUsername) (*mtproto.Bool, error) {
	client := username.NewRPCUsernameClient(m.cli.Conn())
	return client.UsernameUpdateUsername(ctx, in)
}

// UsernameDeleteUsername
// username.deleteUsername username:string = Bool;
func (m *defaultUsernameClient) UsernameDeleteUsername(ctx context.Context, in *username.TLUsernameDeleteUsername) (*mtproto.Bool, error) {
	client := username.NewRPCUsernameClient(m.cli.Conn())
	return client.UsernameDeleteUsername(ctx, in)
}

// UsernameResolveUsername
// username.resolveUsername username:string = Peer;
func (m *defaultUsernameClient) UsernameResolveUsername(ctx context.Context, in *username.TLUsernameResolveUsername) (*mtproto.Peer, error) {
	client := username.NewRPCUsernameClient(m.cli.Conn())
	return client.UsernameResolveUsername(ctx, in)
}

// UsernameGetListByUsernameList
// username.getListByUsernameList names:Vector<string> = Vector<UsernameData>;
func (m *defaultUsernameClient) UsernameGetListByUsernameList(ctx context.Context, in *username.TLUsernameGetListByUsernameList) (*username.Vector_UsernameData, error) {
	client := username.NewRPCUsernameClient(m.cli.Conn())
	return client.UsernameGetListByUsernameList(ctx, in)
}

// UsernameDeleteUsernameByPeer
// username.deleteUsernameByPeer peer_type:int peer_id:long = Bool;
func (m *defaultUsernameClient) UsernameDeleteUsernameByPeer(ctx context.Context, in *username.TLUsernameDeleteUsernameByPeer) (*mtproto.Bool, error) {
	client := username.NewRPCUsernameClient(m.cli.Conn())
	return client.UsernameDeleteUsernameByPeer(ctx, in)
}

// UsernameSearch
// username.search q:string excluded_contacts:Vector<long> limit:int = Vector<UsernameData>;
func (m *defaultUsernameClient) UsernameSearch(ctx context.Context, in *username.TLUsernameSearch) (*username.Vector_UsernameData, error) {
	client := username.NewRPCUsernameClient(m.cli.Conn())
	return client.UsernameSearch(ctx, in)
}
