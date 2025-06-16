/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package notificationclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/notification/notification/notificationservice"

	"github.com/cloudwego/kitex/client"
)

type NotificationClient interface {
	AccountRegisterDevice(ctx context.Context, in *tg.TLAccountRegisterDevice) (*tg.Bool, error)
	AccountUnregisterDevice(ctx context.Context, in *tg.TLAccountUnregisterDevice) (*tg.Bool, error)
	AccountUpdateNotifySettings(ctx context.Context, in *tg.TLAccountUpdateNotifySettings) (*tg.Bool, error)
	AccountGetNotifySettings(ctx context.Context, in *tg.TLAccountGetNotifySettings) (*tg.PeerNotifySettings, error)
	AccountResetNotifySettings(ctx context.Context, in *tg.TLAccountResetNotifySettings) (*tg.Bool, error)
	AccountUpdateDeviceLocked(ctx context.Context, in *tg.TLAccountUpdateDeviceLocked) (*tg.Bool, error)
	AccountGetNotifyExceptions(ctx context.Context, in *tg.TLAccountGetNotifyExceptions) (*tg.Updates, error)
}

type defaultNotificationClient struct {
	cli client.Client
}

func NewNotificationClient(cli client.Client) NotificationClient {
	return &defaultNotificationClient{
		cli: cli,
	}
}

// AccountRegisterDevice
// account.registerDevice#ec86017a flags:# no_muted:flags.0?true token_type:int token:string app_sandbox:Bool secret:bytes other_uids:Vector<long> = Bool;
func (m *defaultNotificationClient) AccountRegisterDevice(ctx context.Context, in *tg.TLAccountRegisterDevice) (*tg.Bool, error) {
	cli := notificationservice.NewRPCNotificationClient(m.cli)
	return cli.AccountRegisterDevice(ctx, in)
}

// AccountUnregisterDevice
// account.unregisterDevice#6a0d3206 token_type:int token:string other_uids:Vector<long> = Bool;
func (m *defaultNotificationClient) AccountUnregisterDevice(ctx context.Context, in *tg.TLAccountUnregisterDevice) (*tg.Bool, error) {
	cli := notificationservice.NewRPCNotificationClient(m.cli)
	return cli.AccountUnregisterDevice(ctx, in)
}

// AccountUpdateNotifySettings
// account.updateNotifySettings#84be5b93 peer:InputNotifyPeer settings:InputPeerNotifySettings = Bool;
func (m *defaultNotificationClient) AccountUpdateNotifySettings(ctx context.Context, in *tg.TLAccountUpdateNotifySettings) (*tg.Bool, error) {
	cli := notificationservice.NewRPCNotificationClient(m.cli)
	return cli.AccountUpdateNotifySettings(ctx, in)
}

// AccountGetNotifySettings
// account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;
func (m *defaultNotificationClient) AccountGetNotifySettings(ctx context.Context, in *tg.TLAccountGetNotifySettings) (*tg.PeerNotifySettings, error) {
	cli := notificationservice.NewRPCNotificationClient(m.cli)
	return cli.AccountGetNotifySettings(ctx, in)
}

// AccountResetNotifySettings
// account.resetNotifySettings#db7e1747 = Bool;
func (m *defaultNotificationClient) AccountResetNotifySettings(ctx context.Context, in *tg.TLAccountResetNotifySettings) (*tg.Bool, error) {
	cli := notificationservice.NewRPCNotificationClient(m.cli)
	return cli.AccountResetNotifySettings(ctx, in)
}

// AccountUpdateDeviceLocked
// account.updateDeviceLocked#38df3532 period:int = Bool;
func (m *defaultNotificationClient) AccountUpdateDeviceLocked(ctx context.Context, in *tg.TLAccountUpdateDeviceLocked) (*tg.Bool, error) {
	cli := notificationservice.NewRPCNotificationClient(m.cli)
	return cli.AccountUpdateDeviceLocked(ctx, in)
}

// AccountGetNotifyExceptions
// account.getNotifyExceptions#53577479 flags:# compare_sound:flags.1?true compare_stories:flags.2?true peer:flags.0?InputNotifyPeer = Updates;
func (m *defaultNotificationClient) AccountGetNotifyExceptions(ctx context.Context, in *tg.TLAccountGetNotifyExceptions) (*tg.Updates, error) {
	cli := notificationservice.NewRPCNotificationClient(m.cli)
	return cli.AccountGetNotifyExceptions(ctx, in)
}
