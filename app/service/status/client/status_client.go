/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package status_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/status/status"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type StatusClient interface {
	StatusSetSessionOnline(ctx context.Context, in *status.TLStatusSetSessionOnline) (*mtproto.Bool, error)
	StatusSetSessionOffline(ctx context.Context, in *status.TLStatusSetSessionOffline) (*mtproto.Bool, error)
	StatusGetUserOnlineSessions(ctx context.Context, in *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error)
	StatusGetUsersOnlineSessionsList(ctx context.Context, in *status.TLStatusGetUsersOnlineSessionsList) (*status.Vector_UserSessionEntryList, error)
	StatusGetChannelOnlineUsers(ctx context.Context, in *status.TLStatusGetChannelOnlineUsers) (*status.Vector_Long, error)
	StatusSetUserChannelsOnline(ctx context.Context, in *status.TLStatusSetUserChannelsOnline) (*mtproto.Bool, error)
	StatusSetUserChannelsOffline(ctx context.Context, in *status.TLStatusSetUserChannelsOffline) (*mtproto.Bool, error)
	StatusSetChannelUserOffline(ctx context.Context, in *status.TLStatusSetChannelUserOffline) (*mtproto.Bool, error)
	StatusSetChannelUsersOnline(ctx context.Context, in *status.TLStatusSetChannelUsersOnline) (*mtproto.Bool, error)
	StatusSetChannelOffline(ctx context.Context, in *status.TLStatusSetChannelOffline) (*mtproto.Bool, error)
}

type defaultStatusClient struct {
	cli zrpc.Client
}

func NewStatusClient(cli zrpc.Client) StatusClient {
	return &defaultStatusClient{
		cli: cli,
	}
}

// StatusSetSessionOnline
// status.setSessionOnline user_id:long session:SessionEntry = Bool;
func (m *defaultStatusClient) StatusSetSessionOnline(ctx context.Context, in *status.TLStatusSetSessionOnline) (*mtproto.Bool, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusSetSessionOnline(ctx, in)
}

// StatusSetSessionOffline
// status.setSessionOffline user_id:long auth_key_id:long = Bool;
func (m *defaultStatusClient) StatusSetSessionOffline(ctx context.Context, in *status.TLStatusSetSessionOffline) (*mtproto.Bool, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusSetSessionOffline(ctx, in)
}

// StatusGetUserOnlineSessions
// status.getUserOnlineSessions user_id:long = UserSessionEntryList;
func (m *defaultStatusClient) StatusGetUserOnlineSessions(ctx context.Context, in *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusGetUserOnlineSessions(ctx, in)
}

// StatusGetUsersOnlineSessionsList
// status.getUsersOnlineSessionsList users:Vector<long> = Vector<UserSessionEntryList>;
func (m *defaultStatusClient) StatusGetUsersOnlineSessionsList(ctx context.Context, in *status.TLStatusGetUsersOnlineSessionsList) (*status.Vector_UserSessionEntryList, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusGetUsersOnlineSessionsList(ctx, in)
}

// StatusGetChannelOnlineUsers
// status.getChannelOnlineUsers channel_id:long = Vector<long>;
func (m *defaultStatusClient) StatusGetChannelOnlineUsers(ctx context.Context, in *status.TLStatusGetChannelOnlineUsers) (*status.Vector_Long, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusGetChannelOnlineUsers(ctx, in)
}

// StatusSetUserChannelsOnline
// status.setUserChannelsOnline user_id:long channels:Vector<long> = Bool;
func (m *defaultStatusClient) StatusSetUserChannelsOnline(ctx context.Context, in *status.TLStatusSetUserChannelsOnline) (*mtproto.Bool, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusSetUserChannelsOnline(ctx, in)
}

// StatusSetUserChannelsOffline
// status.setUserChannelsOffline user_id:long channels:Vector<long> = Bool;
func (m *defaultStatusClient) StatusSetUserChannelsOffline(ctx context.Context, in *status.TLStatusSetUserChannelsOffline) (*mtproto.Bool, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusSetUserChannelsOffline(ctx, in)
}

// StatusSetChannelUserOffline
// status.setChannelUserOffline channel_id:long user_id:long = Bool;
func (m *defaultStatusClient) StatusSetChannelUserOffline(ctx context.Context, in *status.TLStatusSetChannelUserOffline) (*mtproto.Bool, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusSetChannelUserOffline(ctx, in)
}

// StatusSetChannelUsersOnline
// status.setChannelUsersOnline channel_id:long user_id:long = Bool;
func (m *defaultStatusClient) StatusSetChannelUsersOnline(ctx context.Context, in *status.TLStatusSetChannelUsersOnline) (*mtproto.Bool, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusSetChannelUsersOnline(ctx, in)
}

// StatusSetChannelOffline
// status.setChannelOffline channel_id:long = Bool;
func (m *defaultStatusClient) StatusSetChannelOffline(ctx context.Context, in *status.TLStatusSetChannelOffline) (*mtproto.Bool, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusSetChannelOffline(ctx, in)
}
