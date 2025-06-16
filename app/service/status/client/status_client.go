/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package statusclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
	"github.com/teamgram/teamgram-server/v2/app/service/status/status/statusservice"

	"github.com/cloudwego/kitex/client"
)

type StatusClient interface {
	StatusSetSessionOnline(ctx context.Context, in *status.TLStatusSetSessionOnline) (*tg.Bool, error)
	StatusSetSessionOffline(ctx context.Context, in *status.TLStatusSetSessionOffline) (*tg.Bool, error)
	StatusGetUserOnlineSessions(ctx context.Context, in *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error)
	StatusGetUsersOnlineSessionsList(ctx context.Context, in *status.TLStatusGetUsersOnlineSessionsList) (*status.VectorUserSessionEntryList, error)
	StatusGetChannelOnlineUsers(ctx context.Context, in *status.TLStatusGetChannelOnlineUsers) (*status.VectorLong, error)
	StatusSetUserChannelsOnline(ctx context.Context, in *status.TLStatusSetUserChannelsOnline) (*tg.Bool, error)
	StatusSetUserChannelsOffline(ctx context.Context, in *status.TLStatusSetUserChannelsOffline) (*tg.Bool, error)
	StatusSetChannelUserOffline(ctx context.Context, in *status.TLStatusSetChannelUserOffline) (*tg.Bool, error)
	StatusSetChannelUsersOnline(ctx context.Context, in *status.TLStatusSetChannelUsersOnline) (*tg.Bool, error)
	StatusSetChannelOffline(ctx context.Context, in *status.TLStatusSetChannelOffline) (*tg.Bool, error)
}

type defaultStatusClient struct {
	cli client.Client
}

func NewStatusClient(cli client.Client) StatusClient {
	return &defaultStatusClient{
		cli: cli,
	}
}

// StatusSetSessionOnline
// status.setSessionOnline user_id:long session:SessionEntry = Bool;
func (m *defaultStatusClient) StatusSetSessionOnline(ctx context.Context, in *status.TLStatusSetSessionOnline) (*tg.Bool, error) {
	cli := statusservice.NewRPCStatusClient(m.cli)
	return cli.StatusSetSessionOnline(ctx, in)
}

// StatusSetSessionOffline
// status.setSessionOffline user_id:long auth_key_id:long = Bool;
func (m *defaultStatusClient) StatusSetSessionOffline(ctx context.Context, in *status.TLStatusSetSessionOffline) (*tg.Bool, error) {
	cli := statusservice.NewRPCStatusClient(m.cli)
	return cli.StatusSetSessionOffline(ctx, in)
}

// StatusGetUserOnlineSessions
// status.getUserOnlineSessions user_id:long = UserSessionEntryList;
func (m *defaultStatusClient) StatusGetUserOnlineSessions(ctx context.Context, in *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error) {
	cli := statusservice.NewRPCStatusClient(m.cli)
	return cli.StatusGetUserOnlineSessions(ctx, in)
}

// StatusGetUsersOnlineSessionsList
// status.getUsersOnlineSessionsList users:Vector<long> = Vector<UserSessionEntryList>;
func (m *defaultStatusClient) StatusGetUsersOnlineSessionsList(ctx context.Context, in *status.TLStatusGetUsersOnlineSessionsList) (*status.VectorUserSessionEntryList, error) {
	cli := statusservice.NewRPCStatusClient(m.cli)
	return cli.StatusGetUsersOnlineSessionsList(ctx, in)
}

// StatusGetChannelOnlineUsers
// status.getChannelOnlineUsers channel_id:long = Vector<long>;
func (m *defaultStatusClient) StatusGetChannelOnlineUsers(ctx context.Context, in *status.TLStatusGetChannelOnlineUsers) (*status.VectorLong, error) {
	cli := statusservice.NewRPCStatusClient(m.cli)
	return cli.StatusGetChannelOnlineUsers(ctx, in)
}

// StatusSetUserChannelsOnline
// status.setUserChannelsOnline user_id:long channels:Vector<long> = Bool;
func (m *defaultStatusClient) StatusSetUserChannelsOnline(ctx context.Context, in *status.TLStatusSetUserChannelsOnline) (*tg.Bool, error) {
	cli := statusservice.NewRPCStatusClient(m.cli)
	return cli.StatusSetUserChannelsOnline(ctx, in)
}

// StatusSetUserChannelsOffline
// status.setUserChannelsOffline user_id:long channels:Vector<long> = Bool;
func (m *defaultStatusClient) StatusSetUserChannelsOffline(ctx context.Context, in *status.TLStatusSetUserChannelsOffline) (*tg.Bool, error) {
	cli := statusservice.NewRPCStatusClient(m.cli)
	return cli.StatusSetUserChannelsOffline(ctx, in)
}

// StatusSetChannelUserOffline
// status.setChannelUserOffline channel_id:long user_id:long = Bool;
func (m *defaultStatusClient) StatusSetChannelUserOffline(ctx context.Context, in *status.TLStatusSetChannelUserOffline) (*tg.Bool, error) {
	cli := statusservice.NewRPCStatusClient(m.cli)
	return cli.StatusSetChannelUserOffline(ctx, in)
}

// StatusSetChannelUsersOnline
// status.setChannelUsersOnline channel_id:long id:Vector<long> = Bool;
func (m *defaultStatusClient) StatusSetChannelUsersOnline(ctx context.Context, in *status.TLStatusSetChannelUsersOnline) (*tg.Bool, error) {
	cli := statusservice.NewRPCStatusClient(m.cli)
	return cli.StatusSetChannelUsersOnline(ctx, in)
}

// StatusSetChannelOffline
// status.setChannelOffline channel_id:long = Bool;
func (m *defaultStatusClient) StatusSetChannelOffline(ctx context.Context, in *status.TLStatusSetChannelOffline) (*tg.Bool, error) {
	cli := statusservice.NewRPCStatusClient(m.cli)
	return cli.StatusSetChannelOffline(ctx, in)
}
