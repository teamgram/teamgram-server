/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
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
	c.Logger.Debugf("users.getUsers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.UsersGetUsers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("users.getUsers - reply: {%s}", r)
	return r, err
}

// UsersGetFullUser
// users.getFullUser#b60f5918 id:InputUser = users.UserFull;
func (s *Service) UsersGetFullUser(ctx context.Context, request *mtproto.TLUsersGetFullUser) (*mtproto.Users_UserFull, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("users.getFullUser - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.UsersGetFullUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("users.getFullUser - reply: {%s}", r)
	return r, err
}

// ContactsResolvePhone
// contacts.resolvePhone#8af94344 phone:string = contacts.ResolvedPeer;
func (s *Service) ContactsResolvePhone(ctx context.Context, request *mtproto.TLContactsResolvePhone) (*mtproto.Contacts_ResolvedPeer, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.resolvePhone - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsResolvePhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.resolvePhone - reply: {%s}", r)
	return r, err
}

// UsersGetMe
// users.getMe id:long token:string = User;
func (s *Service) UsersGetMe(ctx context.Context, request *mtproto.TLUsersGetMe) (*mtproto.User, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("users.getMe - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.UsersGetMe(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("users.getMe - reply: {%s}", r)
	return r, err
}
