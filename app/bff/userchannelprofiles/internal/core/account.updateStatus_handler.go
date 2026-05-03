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
	"time"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountUpdateStatus
// account.updateStatus#6628562c offline:Bool = Bool;
func (c *UserChannelProfilesCore) AccountUpdateStatus(in *tg.TLAccountUpdateStatus) (*tg.Bool, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil || in.Offline == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if err := requireUserClient(c); err != nil {
		return nil, err
	}

	expires := int32(0)
	if _, offline := in.Offline.(*tg.TLBoolTrue); !offline {
		expires = 300
	}
	if _, err = c.svcCtx.Repo.UserClient.UserUpdateLastSeen(c.ctx, &userpb.TLUserUpdateLastSeen{
		Id:         selfID,
		LastSeenAt: time.Now().Unix(),
		Expires:    expires,
	}); err != nil {
		return nil, err
	}
	// TODO(v2 userchannelprofiles): sync delivery is intentionally not migrated here; route status updates through userupdates/gateway when the V2 delivery contract is defined.
	return tg.BoolTrue, nil
}
