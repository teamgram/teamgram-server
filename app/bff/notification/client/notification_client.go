/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package notification_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type NotificationClient interface {
	AccountRegisterDevice(ctx context.Context, in *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error)
	AccountUnregisterDevice(ctx context.Context, in *mtproto.TLAccountUnregisterDevice) (*mtproto.Bool, error)
	AccountUpdateNotifySettings(ctx context.Context, in *mtproto.TLAccountUpdateNotifySettings) (*mtproto.Bool, error)
	AccountGetNotifySettings(ctx context.Context, in *mtproto.TLAccountGetNotifySettings) (*mtproto.PeerNotifySettings, error)
	AccountResetNotifySettings(ctx context.Context, in *mtproto.TLAccountResetNotifySettings) (*mtproto.Bool, error)
	AccountUpdateDeviceLocked(ctx context.Context, in *mtproto.TLAccountUpdateDeviceLocked) (*mtproto.Bool, error)
	AccountGetNotifyExceptions(ctx context.Context, in *mtproto.TLAccountGetNotifyExceptions) (*mtproto.Updates, error)
}

type defaultNotificationClient struct {
	cli zrpc.Client
}

func NewNotificationClient(cli zrpc.Client) NotificationClient {
	return &defaultNotificationClient{
		cli: cli,
	}
}

// AccountRegisterDevice
// account.registerDevice#ec86017a flags:# no_muted:flags.0?true token_type:int token:string app_sandbox:Bool secret:bytes other_uids:Vector<long> = Bool;
func (m *defaultNotificationClient) AccountRegisterDevice(ctx context.Context, in *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error) {
	client := mtproto.NewRPCNotificationClient(m.cli.Conn())
	return client.AccountRegisterDevice(ctx, in)
}

// AccountUnregisterDevice
// account.unregisterDevice#6a0d3206 token_type:int token:string other_uids:Vector<long> = Bool;
func (m *defaultNotificationClient) AccountUnregisterDevice(ctx context.Context, in *mtproto.TLAccountUnregisterDevice) (*mtproto.Bool, error) {
	client := mtproto.NewRPCNotificationClient(m.cli.Conn())
	return client.AccountUnregisterDevice(ctx, in)
}

// AccountUpdateNotifySettings
// account.updateNotifySettings#84be5b93 peer:InputNotifyPeer settings:InputPeerNotifySettings = Bool;
func (m *defaultNotificationClient) AccountUpdateNotifySettings(ctx context.Context, in *mtproto.TLAccountUpdateNotifySettings) (*mtproto.Bool, error) {
	client := mtproto.NewRPCNotificationClient(m.cli.Conn())
	return client.AccountUpdateNotifySettings(ctx, in)
}

// AccountGetNotifySettings
// account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;
func (m *defaultNotificationClient) AccountGetNotifySettings(ctx context.Context, in *mtproto.TLAccountGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	client := mtproto.NewRPCNotificationClient(m.cli.Conn())
	return client.AccountGetNotifySettings(ctx, in)
}

// AccountResetNotifySettings
// account.resetNotifySettings#db7e1747 = Bool;
func (m *defaultNotificationClient) AccountResetNotifySettings(ctx context.Context, in *mtproto.TLAccountResetNotifySettings) (*mtproto.Bool, error) {
	client := mtproto.NewRPCNotificationClient(m.cli.Conn())
	return client.AccountResetNotifySettings(ctx, in)
}

// AccountUpdateDeviceLocked
// account.updateDeviceLocked#38df3532 period:int = Bool;
func (m *defaultNotificationClient) AccountUpdateDeviceLocked(ctx context.Context, in *mtproto.TLAccountUpdateDeviceLocked) (*mtproto.Bool, error) {
	client := mtproto.NewRPCNotificationClient(m.cli.Conn())
	return client.AccountUpdateDeviceLocked(ctx, in)
}

// AccountGetNotifyExceptions
// account.getNotifyExceptions#53577479 flags:# compare_sound:flags.1?true peer:flags.0?InputNotifyPeer = Updates;
func (m *defaultNotificationClient) AccountGetNotifyExceptions(ctx context.Context, in *mtproto.TLAccountGetNotifyExceptions) (*mtproto.Updates, error) {
	client := mtproto.NewRPCNotificationClient(m.cli.Conn())
	return client.AccountGetNotifyExceptions(ctx, in)
}
