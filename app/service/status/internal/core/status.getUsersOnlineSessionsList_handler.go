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
	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/status/status"

	"github.com/zeromicro/go-zero/core/jsonx"
)

var _ *tg.Bool

// StatusGetUsersOnlineSessionsList
// status.getUsersOnlineSessionsList users:Vector<long> = Vector<UserSessionEntryList>;
func (c *StatusCore) StatusGetUsersOnlineSessionsList(in *status.TLStatusGetUsersOnlineSessionsList) (*status.VectorUserSessionEntryList, error) {
	var (
		// pipeline
		pipes = make(map[interface{}][]string)

		rValues = &status.VectorUserSessionEntryList{
			Datas: make([]*status.UserSessionEntryList, 0, len(in.Users)),
		}
	)

	for _, id := range in.Users {
		k := getUserKey(id)
		pipe, err := c.svcCtx.Dao.KV.GetPipeline(k)
		if err != nil {
			c.Logger.Errorf("status.getUsersOnlineSessionsList(%s) error(%v)", in, err)
			continue
		}
		pipes[pipe] = append(pipes[pipe], k)
	}

	for pipe, kList := range pipes {
		var (
			cmds = make([]*kv.MapStringStringCmd, len(kList))
		)

		pipe.(kv.Pipeline).PipelinedCtx(
			c.ctx,
			func(pipeliner kv.Pipeliner) error {
				for i, k := range kList {
					cmds[i] = pipeliner.HGetAll(c.ctx, k)
				}

				return nil
			})

		for i := 0; i < len(kList); i++ {
			rMap, err := cmds[i].Result()
			if err != nil {
				c.Logger.Errorf("status.getUsersOnlineSessionsList(%s) error(%v)", in, err)
				return nil, err
			}

			var (
				sessions = &status.TLUserSessionEntryList{
					UserId:       getIdByUserKey(kList[i]),
					UserSessions: make([]*status.SessionEntry, len(rMap)),
				}
			)

			sessions.UserSessions = make([]*status.SessionEntry, 0, len(rMap))
			for _, v := range rMap {
				// keyId, _ := strconv.ParseInt(k, 10, 64)
				sess := new(status.SessionEntry)
				if err2 := jsonx.UnmarshalFromString(v, sess); err2 == nil {
					sessions.UserSessions = append(sessions.UserSessions, sess)
				}
			}

			rValues.Datas = append(rValues.Datas, sessions.ToUserSessionEntryList())
		}
	}

	return rValues, nil
}
