// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package userclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type UserClientHelper struct {
	cli UserClient
}

func NewUserClientHelper(cli zrpc.Client) *UserClientHelper {
	return &UserClientHelper{
		cli: NewUserClient(cli),
	}
}

func (m *UserClientHelper) Client() UserClient {
	return m.cli
}

func (m *UserClientHelper) GetUserSelf(ctx context.Context, id int64) (*mtproto.User, error) {
	user, err := m.cli.UserGetImmutableUser(ctx, &user.TLUserGetImmutableUser{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return user.ToSelfUser(), nil
}

func (m *UserClientHelper) GetUserById(ctx context.Context, selfId, id int64) (*mtproto.User, error) {
	if selfId == id {
		return m.GetUserSelf(ctx, selfId)
	}

	users := m.GetUserListByIdList(ctx, selfId, id)
	if len(users) == 0 {
		return nil, mtproto.ErrUserIdInvalid
	}

	return users[0], nil
}

func (m *UserClientHelper) GetUserListByIdList(ctx context.Context, selfId int64, id ...int64) []*mtproto.User {
	users, err := m.cli.UserGetMutableUsers(ctx, &user.TLUserGetMutableUsers{
		Id: append(id, selfId),
		To: []int64{selfId},
	})
	if err != nil {
		return []*mtproto.User{}
	}

	return users.GetUserListByIdList(selfId, id...)
}
