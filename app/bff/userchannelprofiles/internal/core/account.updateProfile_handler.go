// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountUpdateProfile
// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
func (c *UserChannelProfilesCore) AccountUpdateProfile(in *tg.TLAccountUpdateProfile) (*tg.User, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if err := requireUserClient(c); err != nil {
		return nil, err
	}

	me, err := c.svcCtx.Repo.UserClient.UserGetImmutableUser(c.ctx, &userpb.TLUserGetImmutableUser{
		Id: selfID,
	})
	if err != nil {
		return nil, err
	}
	if me == nil || me.User == nil {
		return nil, tg.ErrUserIdInvalid
	}

	if in.About != nil {
		if len(*in.About) > 128 {
			return nil, tg.ErrAboutTooLong
		}
		currentAbout := ""
		if me.User.About != nil {
			currentAbout = *me.User.About
		}
		if *in.About != currentAbout {
			if _, err = c.svcCtx.Repo.UserClient.UserUpdateAbout(c.ctx, &userpb.TLUserUpdateAbout{
				UserId: selfID,
				About:  *in.About,
			}); err != nil {
				return nil, err
			}
			me.User.About = in.About
		}
	}

	firstName := ""
	lastName := ""
	if in.FirstName != nil {
		firstName = *in.FirstName
	}
	if in.LastName != nil {
		lastName = *in.LastName
	}
	if firstName != me.User.FirstName || lastName != me.User.LastName {
		if _, err = c.svcCtx.Repo.UserClient.UserUpdateFirstAndLastName(c.ctx, &userpb.TLUserUpdateFirstAndLastName{
			UserId:    selfID,
			FirstName: firstName,
			LastName:  lastName,
		}); err != nil {
			return nil, err
		}
		me.User.FirstName = firstName
		me.User.LastName = lastName
		// TODO(v2 userchannelprofiles): sync delivery is intentionally not migrated here; route profile updates through userupdates/gateway when the V2 delivery contract is defined.
	}

	users, err := userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, selfID, []int64{selfID}, userprojection.MissingExplicitInput)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, tg.ErrUserIdInvalid
	}

	return &tg.User{Clazz: users[0]}, nil
}
