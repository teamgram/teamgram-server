/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package banned_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/banned/banned"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type BannedClient interface {
	BannedCheckPhoneNumberBanned(ctx context.Context, in *banned.TLBannedCheckPhoneNumberBanned) (*mtproto.Bool, error)
	BannedGetBannedByPhoneList(ctx context.Context, in *banned.TLBannedGetBannedByPhoneList) (*banned.Vector_String, error)
	BannedBan(ctx context.Context, in *banned.TLBannedBan) (*mtproto.Bool, error)
	BannedUnBan(ctx context.Context, in *banned.TLBannedUnBan) (*mtproto.Bool, error)
}

type defaultBannedClient struct {
	cli zrpc.Client
}

func NewBannedClient(cli zrpc.Client) BannedClient {
	return &defaultBannedClient{
		cli: cli,
	}
}

// BannedCheckPhoneNumberBanned
// banned.checkPhoneNumberBanned phone:string = Bool;
func (m *defaultBannedClient) BannedCheckPhoneNumberBanned(ctx context.Context, in *banned.TLBannedCheckPhoneNumberBanned) (*mtproto.Bool, error) {
	client := banned.NewRPCBannedClient(m.cli.Conn())
	return client.BannedCheckPhoneNumberBanned(ctx, in)
}

// BannedGetBannedByPhoneList
// banned.getBannedByPhoneList phones:Vector<string> = Vector<string>;
func (m *defaultBannedClient) BannedGetBannedByPhoneList(ctx context.Context, in *banned.TLBannedGetBannedByPhoneList) (*banned.Vector_String, error) {
	client := banned.NewRPCBannedClient(m.cli.Conn())
	return client.BannedGetBannedByPhoneList(ctx, in)
}

// BannedBan
// banned.ban phone:string expires:int reason:string = Bool;
func (m *defaultBannedClient) BannedBan(ctx context.Context, in *banned.TLBannedBan) (*mtproto.Bool, error) {
	client := banned.NewRPCBannedClient(m.cli.Conn())
	return client.BannedBan(ctx, in)
}

// BannedUnBan
// banned.unBan phone:string = Bool;
func (m *defaultBannedClient) BannedUnBan(ctx context.Context, in *banned.TLBannedUnBan) (*mtproto.Bool, error) {
	client := banned.NewRPCBannedClient(m.cli.Conn())
	return client.BannedUnBan(ctx, in)
}
