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
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/teamgram-server/app/service/status/status"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// StatusGetUsersOnlineSessionsList
// status.getUsersOnlineSessionsList Vector<long>:users = Vector<UserSessionEntryList>;
func (c *StatusCore) StatusGetUsersOnlineSessionsList(in *status.TLStatusGetUsersOnlineSessionsList) (*status.Vector_UserSessionEntryList, error) {
	type keyInfo struct {
		userId int64
		key    string
	}
	type pipeGroup struct {
		pipe kv.Pipeline
		keys []keyInfo
	}

	users := in.GetUsers()
	if len(users) > maxBatchUsers {
		return nil, fmt.Errorf("status.getUsersOnlineSessionsList - too many users: %d, max %d", len(users), maxBatchUsers)
	}

	rValues := &status.Vector_UserSessionEntryList{
		Datas: make([]*status.UserSessionEntryList, 0, len(users)),
	}

	groups := make(map[kv.Pipeline]*pipeGroup)
	var failedUsers []int64

	for _, id := range users {
		k := getUserKey(id)
		rawPipe, err := c.svcCtx.Dao.KV.GetPipeline(k)
		if err != nil {
			c.Logger.Errorf("status.getUsersOnlineSessionsList - GetPipeline(userId=%d) error: %v", id, err)
			failedUsers = append(failedUsers, id)
			continue
		}
		p, ok := rawPipe.(kv.Pipeline)
		if !ok {
			c.Logger.Errorf("status.getUsersOnlineSessionsList - unexpected pipeline type: %T", rawPipe)
			failedUsers = append(failedUsers, id)
			continue
		}
		g, exists := groups[p]
		if !exists {
			g = &pipeGroup{pipe: p}
			groups[p] = g
		}
		g.keys = append(g.keys, keyInfo{userId: id, key: k})
	}

	// fill empty results for failed users
	for _, uid := range failedUsers {
		rValues.Datas = append(rValues.Datas, status.MakeTLUserSessionEntryList(&status.UserSessionEntryList{
			UserId: uid,
		}).To_UserSessionEntryList())
	}

	for _, g := range groups {
		cmds := make([]*kv.MapStringStringCmd, len(g.keys))

		err := g.pipe.PipelinedCtx(c.ctx, func(pipeliner kv.Pipeliner) error {
			for i, ki := range g.keys {
				cmds[i] = pipeliner.HGetAll(c.ctx, ki.key)
			}
			return nil
		})
		if err != nil {
			c.Logger.Errorf("status.getUsersOnlineSessionsList - pipeline exec error: %v", err)
			return nil, fmt.Errorf("status.getUsersOnlineSessionsList - pipeline exec: %w", err)
		}

		for i, ki := range g.keys {
			rMap, err := cmds[i].Result()
			if err != nil {
				c.Logger.Errorf("status.getUsersOnlineSessionsList - cmd result(userId=%d) error: %v", ki.userId, err)
				return nil, fmt.Errorf("status.getUsersOnlineSessionsList - cmd result userId=%d: %w", ki.userId, err)
			}

			sessions := make([]*status.SessionEntry, 0, len(rMap))
			for _, v := range rMap {
				sess := new(status.SessionEntry)
				if err2 := jsonx.UnmarshalFromString(v, sess); err2 != nil {
					c.Logger.Infof("status.getUsersOnlineSessionsList - unmarshal(userId=%d) error: %v", ki.userId, err2)
					continue
				}
				sessions = append(sessions, sess)
			}

			rValues.Datas = append(rValues.Datas, status.MakeTLUserSessionEntryList(&status.UserSessionEntryList{
				UserId:       ki.userId,
				UserSessions: sessions,
			}).To_UserSessionEntryList())
		}
	}

	return rValues, nil
}
