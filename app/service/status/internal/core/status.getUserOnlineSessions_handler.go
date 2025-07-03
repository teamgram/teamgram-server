// Copyright 2024 Teamgram Authors
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
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/status/status"

	"github.com/zeromicro/go-zero/core/jsonx"
)

var _ *tg.Bool

// StatusGetUserOnlineSessions
// status.getUserOnlineSessions user_id:long = UserSessionEntryList;
func (c *StatusCore) StatusGetUserOnlineSessions(in *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error) {
	rMap, err := c.svcCtx.Dao.KV.HgetallCtx(c.ctx, getUserKey(in.UserId))
	if err != nil {
		c.Logger.Errorf("status.getUserOnlineSessions(%s) error(%v)", in, err)
		return nil, err
	}

	rValues := &status.TLUserSessionEntryList{
		UserId:       in.UserId,
		UserSessions: make([]*status.SessionEntry, 0, len(rMap)),
	}

	for _, v := range rMap {
		sess := new(status.SessionEntry)
		if err2 := jsonx.UnmarshalFromString(v, sess); err2 == nil {
			rValues.UserSessions = append(rValues.UserSessions, sess)
		}
	}

	return rValues.ToUserSessionEntryList(), nil
}
