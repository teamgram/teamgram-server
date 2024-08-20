// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// AccountUpdateProfile
// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
func (c *UserProfileCore) AccountUpdateProfile(in *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	me, err := c.svcCtx.Dao.UserClient.UserGetImmutableUser(c.ctx, &userpb.TLUserGetImmutableUser{
		Id: c.MD.UserId,
	})

	if in.GetAbout() != nil {
		//// about长度<70并且可以为emtpy
		if len(in.GetAbout().GetValue()) > 70 {
			err = mtproto.ErrAboutTooLong
			c.Logger.Errorf("account.updateProfile - error: %v", err)
			return nil, err
		}

		if in.GetAbout().GetValue() != me.About() {
			if _, err = c.svcCtx.Dao.UserClient.UserUpdateAbout(c.ctx, &userpb.TLUserUpdateAbout{
				UserId: c.MD.UserId,
				About:  in.GetAbout().GetValue(),
			}); err != nil {
				c.Logger.Errorf("account.updateProfile - error: %v", err)
			} else {
				me.SetAbout(in.GetAbout().GetValue())
			}
		}
	} else {
		if in.GetFirstName().GetValue() == "" {
			err = mtproto.ErrFirstnameInvalid
			c.Logger.Errorf("account.updateProfile - error: bad request (%v)", err)
			return nil, err
		}

		if in.GetFirstName().GetValue() != me.FirstName() ||
			in.GetLastName().GetValue() != me.LastName() {
			if _, err = c.svcCtx.Dao.UserClient.UserUpdateFirstAndLastName(c.ctx, &userpb.TLUserUpdateFirstAndLastName{
				UserId:    c.MD.UserId,
				FirstName: in.GetFirstName().GetValue(),
				LastName:  in.GetLastName().GetValue(),
			}); err != nil {
				c.Logger.Errorf("account.updateProfile - error: %v", err)
			} else {
				me.SetFirstName(in.GetFirstName().GetValue())
				me.SetLastName(in.GetLastName().GetValue())
			}

			c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
				UserId:        c.MD.UserId,
				PermAuthKeyId: c.MD.PermAuthKeyId,
				Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateUserName(&mtproto.Update{
					UserId:    c.MD.UserId,
					FirstName: in.GetFirstName().GetValue(),
					LastName:  in.GetLastName().GetValue(),
					Username:  me.Username(),
				}).To_Update()),
			})
		}
	}

	return me.ToSelfUser(), nil
}
