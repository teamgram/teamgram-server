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
	"time"

	"github.com/teamgram/proto/mtproto"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/threading"
)

// AccountUpdateStatus
// account.updateStatus#6628562c offline:Bool = Bool;
func (c *UserProfileCore) AccountUpdateStatus(in *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	var (
		offline     = mtproto.FromBool(in.GetOffline())
		now         = time.Now().Unix()
		expires     = int32(0)
		pushUpdates = mtproto.MakeTLUpdateShort(&mtproto.Updates{
			Update: mtproto.MakeTLUpdateUserStatus(&mtproto.Update{
				UserId:            c.MD.UserId,
				Status_USERSTATUS: nil,
			}).To_Update(),
			Date: int32(now) - 1,
		}).To_Updates()
	)

	if !offline {
		expires = 300
		pushUpdates.Update.Status_USERSTATUS = mtproto.MakeTLUserStatusOnline(&mtproto.UserStatus{
			Expires: int32(now) + expires,
		}).To_UserStatus()

		// online
		// s.UserFacade.UpdateUserStatus(ctx, md.UserId, now)
		// c.svcCtx.Dao.UserClient.UserUpdateLastSeen()
	} else {
		pushUpdates.Update.Status_USERSTATUS = mtproto.MakeTLUserStatusOffline(&mtproto.UserStatus{
			WasOnline: int32(now),
		}).To_UserStatus()
	}

	// threading.GoSafe()
	threading.GoSafe(func() {
		c.svcCtx.Dao.UserClient.UserUpdateLastSeen(
			contextx.ValueOnlyFrom(c.ctx),
			&userpb.TLUserUpdateLastSeen{
				Id:         c.MD.UserId,
				LastSeenAt: now,
				Expires:    expires,
			})
	})

	// TODO: push
	//// log.Debugf("account.updateStatus - reply: {true}")
	//return model.WrapperGoFunc(mtproto.BoolTrue, func() {
	//	// log.Debugf("ready push to other contacts...")
	//	// push to other contacts.
	//	contactIdList := s.UserFacade.GetContactUserIdList(context.Background(), md.UserId)
	//	blockedIdList := s.UserFacade.CheckBlockUserList(context.Background(), md.UserId, contactIdList)
	//
	//	// TODO(@benqi): push updateUserStatus规则
	//	for _, id := range contactIdList {
	//		if md.UserId == id {
	//			// why??
	//			continue
	//		}
	//
	//		if blocked, _ := container2.Contains(id, blockedIdList); blocked {
	//			continue
	//		}
	//
	//		// log.Debugf("check blocked...")
	//		blocked := s.UserFacade.IsBlockedByUser(context.Background(), md.UserId, id)
	//		if blocked {
	//			continue
	//		}
	//
	//		sync_client.PushUpdates(context.Background(), id, pushUpdates)
	//	}
	//
	//}).(*mtproto.Bool), nil

	return mtproto.BoolTrue, nil
}
