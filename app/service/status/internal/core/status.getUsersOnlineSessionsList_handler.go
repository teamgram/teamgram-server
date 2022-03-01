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
	"strconv"

	"github.com/teamgram/teamgram-server/app/service/status/status"
)

// StatusGetUsersOnlineSessionsList
// status.getUsersOnlineSessionsList Vector<long>:users = Vector<UserSessionEntryList>;
func (c *StatusCore) StatusGetUsersOnlineSessionsList(in *status.TLStatusGetUsersOnlineSessionsList) (*status.Vector_UserSessionEntryList, error) {
	rValues := &status.Vector_UserSessionEntryList{
		Datas: make([]*status.UserSessionEntryList, 0, len(in.GetUsers())),
	}

	for _, id := range in.GetUsers() {
		rMap, err := c.svcCtx.KV.Hgetall(getUserKey(id))
		if err != nil {
			c.Logger.Errorf("status.getUsersOnlineSessionsList(%s) error(%v)", in.DebugString(), err)
			return nil, err
		}

		var (
			// i        = 0
			sessions = status.MakeTLUserSessionEntryList(&status.UserSessionEntryList{
				UserId:       id,
				UserSessions: make([]*status.SessionEntry, len(rMap)),
			}).To_UserSessionEntryList()
		)

		sessions.UserSessions = make([]*status.SessionEntry, 0, len(rMap))
		for k, v := range rMap {
			keyId, _ := strconv.ParseInt(k, 10, 64)

			sessions.UserSessions = append(sessions.UserSessions, &status.SessionEntry{
				UserId:    id,
				AuthKeyId: keyId,
				Gateway:   v,
				Expired:   0,
				Layer:     0,
			})
		}

		rValues.Datas = append(rValues.Datas, sessions)
	}

	return rValues, nil
}
