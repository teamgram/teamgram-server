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

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/users/users/usersservice"

	"github.com/cloudwego/kitex/client"
)

type UsersClient interface {
	UsersGetUsers(ctx context.Context, in *tg.TLUsersGetUsers) (*tg.VectorUser, error)
	UsersGetFullUser(ctx context.Context, in *tg.TLUsersGetFullUser) (*tg.UsersUserFull, error)
	ContactsResolvePhone(ctx context.Context, in *tg.TLContactsResolvePhone) (*tg.ContactsResolvedPeer, error)
	UsersGetMe(ctx context.Context, in *tg.TLUsersGetMe) (*tg.User, error)
}

type defaultUsersClient struct {
	cli client.Client
}

func NewUsersClient(cli client.Client) UsersClient {
	return &defaultUsersClient{
		cli: cli,
	}
}

// UsersGetUsers
// users.getUsers#d91a548 id:Vector<InputUser> = Vector<User>;
func (m *defaultUsersClient) UsersGetUsers(ctx context.Context, in *tg.TLUsersGetUsers) (*tg.VectorUser, error) {
	cli := usersservice.NewRPCUsersClient(m.cli)
	return cli.UsersGetUsers(ctx, in)
}

// UsersGetFullUser
// users.getFullUser#b60f5918 id:InputUser = users.UserFull;
func (m *defaultUsersClient) UsersGetFullUser(ctx context.Context, in *tg.TLUsersGetFullUser) (*tg.UsersUserFull, error) {
	cli := usersservice.NewRPCUsersClient(m.cli)
	return cli.UsersGetFullUser(ctx, in)
}

// ContactsResolvePhone
// contacts.resolvePhone#8af94344 phone:string = contacts.ResolvedPeer;
func (m *defaultUsersClient) ContactsResolvePhone(ctx context.Context, in *tg.TLContactsResolvePhone) (*tg.ContactsResolvedPeer, error) {
	cli := usersservice.NewRPCUsersClient(m.cli)
	return cli.ContactsResolvePhone(ctx, in)
}

// UsersGetMe
// users.getMe id:long token:string = User;
func (m *defaultUsersClient) UsersGetMe(ctx context.Context, in *tg.TLUsersGetMe) (*tg.User, error) {
	cli := usersservice.NewRPCUsersClient(m.cli)
	return cli.UsersGetMe(ctx, in)
}
