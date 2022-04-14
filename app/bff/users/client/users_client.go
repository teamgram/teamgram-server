/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package users_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type UsersClient interface {
	UsersGetUsers(ctx context.Context, in *mtproto.TLUsersGetUsers) (*mtproto.Vector_User, error)
	UsersGetFullUserB60F5918(ctx context.Context, in *mtproto.TLUsersGetFullUserB60F5918) (*mtproto.Users_UserFull, error)
	UsersGetFullUserCA30A5B1(ctx context.Context, in *mtproto.TLUsersGetFullUserCA30A5B1) (*mtproto.UserFull, error)
	UsersGetMe(ctx context.Context, in *mtproto.TLUsersGetMe) (*mtproto.User, error)
	ContactsResolvePhone(ctx context.Context, in *mtproto.TLContactsResolvePhone) (*mtproto.Contacts_ResolvedPeer, error)
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

// UsersGetFullUserB60F5918
// users.getFullUser#b60f5918 id:InputUser = users.UserFull;
func (m *defaultUsersClient) UsersGetFullUserB60F5918(ctx context.Context, in *mtproto.TLUsersGetFullUserB60F5918) (*mtproto.Users_UserFull, error) {
	client := mtproto.NewRPCUsersClient(m.cli.Conn())
	return client.UsersGetFullUserB60F5918(ctx, in)
}

// UsersGetFullUserCA30A5B1
// users.getFullUser#ca30a5b1 id:InputUser = UserFull;
func (m *defaultUsersClient) UsersGetFullUserCA30A5B1(ctx context.Context, in *mtproto.TLUsersGetFullUserCA30A5B1) (*mtproto.UserFull, error) {
	client := mtproto.NewRPCUsersClient(m.cli.Conn())
	return client.UsersGetFullUserCA30A5B1(ctx, in)
}

// UsersGetMe
// users.getMe id:long token:string = User;
func (m *defaultUsersClient) UsersGetMe(ctx context.Context, in *mtproto.TLUsersGetMe) (*mtproto.User, error) {
	client := mtproto.NewRPCUsersClient(m.cli.Conn())
	return client.UsersGetMe(ctx, in)
}

// ContactsResolvePhone
// contacts.resolvePhone#8af94344 phone:string = contacts.ResolvedPeer;
func (m *defaultUsersClient) ContactsResolvePhone(ctx context.Context, in *mtproto.TLContactsResolvePhone) (*mtproto.Contacts_ResolvedPeer, error) {
	client := mtproto.NewRPCUsersClient(m.cli.Conn())
	return client.ContactsResolvePhone(ctx, in)
}
