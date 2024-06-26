/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/teamgram-server/app/service/status/status"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// StatusGetUsersOnlineSessionsList
// status.getUsersOnlineSessionsList Vector<long>:users = Vector<UserSessionEntryList>;
func (c *StatusCore) StatusGetUsersOnlineSessionsList(in *status.TLStatusGetUsersOnlineSessionsList) (*status.Vector_UserSessionEntryList, error) {
	var (
		// pipeline
		pipes = make(map[interface{}][]string)

		rValues = &status.Vector_UserSessionEntryList{
			Datas: make([]*status.UserSessionEntryList, 0, len(in.GetUsers())),
		}
	)

	for _, id := range in.GetUsers() {
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
				sessions = status.MakeTLUserSessionEntryList(&status.UserSessionEntryList{
					UserId:       getIdByUserKey(kList[i]),
					UserSessions: make([]*status.SessionEntry, len(rMap)),
				}).To_UserSessionEntryList()
			)

			sessions.UserSessions = make([]*status.SessionEntry, 0, len(rMap))
			for _, v := range rMap {
				// keyId, _ := strconv.ParseInt(k, 10, 64)
				sess := new(status.SessionEntry)
				if err2 := jsonx.UnmarshalFromString(v, sess); err2 == nil {
					sessions.UserSessions = append(sessions.UserSessions, sess)
				}
			}

			rValues.Datas = append(rValues.Datas, sessions)
		}
	}

	return rValues, nil
}
