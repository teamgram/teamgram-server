/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/users/internal/core"
)

// UsersGetUsers
// users.getUsers#d91a548 id:Vector<InputUser> = Vector<User>;
func (s *Service) UsersGetUsers(ctx context.Context, request *mtproto.TLUsersGetUsers) (*mtproto.Vector_User, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("users.getUsers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UsersGetUsers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("users.getUsers - reply: %s", r.DebugString())
	return r, err
}

// UsersGetFullUser
// users.getFullUser#b60f5918 id:InputUser = users.UserFull;
func (s *Service) UsersGetFullUser(ctx context.Context, request *mtproto.TLUsersGetFullUser) (*mtproto.Users_UserFull, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("users.getFullUserB60F5918 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UsersGetFullUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("users.getFullUserB60F5918 - reply: %s", r.DebugString())
	return r, err
}

// UsersGetMe
// users.getMe id:long token:string = User;
func (s *Service) UsersGetMe(ctx context.Context, request *mtproto.TLUsersGetMe) (*mtproto.User, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("users.getMe - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UsersGetMe(request)
	if err != nil {
		return nil, err
	}

	c.Infof("users.getMe - reply: %s", r.DebugString())
	return r, err
}

// ContactsResolvePhone
// contacts.resolvePhone#8af94344 phone:string = contacts.ResolvedPeer;
func (s *Service) ContactsResolvePhone(ctx context.Context, request *mtproto.TLContactsResolvePhone) (*mtproto.Contacts_ResolvedPeer, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.resolvePhone - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsResolvePhone(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.getLocated - reply: %s", r.DebugString())
	return r, err
}
