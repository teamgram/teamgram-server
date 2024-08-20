/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package usersclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type UsersClient interface {
	UsersGetUsers(ctx context.Context, in *mtproto.TLUsersGetUsers) (*mtproto.Vector_User, error)
	UsersGetFullUser(ctx context.Context, in *mtproto.TLUsersGetFullUser) (*mtproto.Users_UserFull, error)
	ContactsResolvePhone(ctx context.Context, in *mtproto.TLContactsResolvePhone) (*mtproto.Contacts_ResolvedPeer, error)
	UsersGetMe(ctx context.Context, in *mtproto.TLUsersGetMe) (*mtproto.User, error)
}

type defaultUsersClient struct {
	cli zrpc.Client
}

func NewUsersClient(cli zrpc.Client) UsersClient {
	return &defaultUsersClient{
		cli: cli,
	}
}

// UsersGetUsers
// users.getUsers#d91a548 id:Vector<InputUser> = Vector<User>;
func (m *defaultUsersClient) UsersGetUsers(ctx context.Context, in *mtproto.TLUsersGetUsers) (*mtproto.Vector_User, error) {
	client := mtproto.NewRPCUsersClient(m.cli.Conn())
	return client.UsersGetUsers(ctx, in)
}

// UsersGetFullUser
// users.getFullUser#b60f5918 id:InputUser = users.UserFull;
func (m *defaultUsersClient) UsersGetFullUser(ctx context.Context, in *mtproto.TLUsersGetFullUser) (*mtproto.Users_UserFull, error) {
	client := mtproto.NewRPCUsersClient(m.cli.Conn())
	return client.UsersGetFullUser(ctx, in)
}

// ContactsResolvePhone
// contacts.resolvePhone#8af94344 phone:string = contacts.ResolvedPeer;
func (m *defaultUsersClient) ContactsResolvePhone(ctx context.Context, in *mtproto.TLContactsResolvePhone) (*mtproto.Contacts_ResolvedPeer, error) {
	client := mtproto.NewRPCUsersClient(m.cli.Conn())
	return client.ContactsResolvePhone(ctx, in)
}

// UsersGetMe
// users.getMe id:long token:string = User;
func (m *defaultUsersClient) UsersGetMe(ctx context.Context, in *mtproto.TLUsersGetMe) (*mtproto.User, error) {
	client := mtproto.NewRPCUsersClient(m.cli.Conn())
	return client.UsersGetMe(ctx, in)
}
