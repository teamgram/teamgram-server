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
	"github.com/teamgram/teamgram-server/app/service/status/status"

	"github.com/zeromicro/go-zero/core/jsonx"
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
			sessions = status.MakeTLUserSessionEntryList(&status.UserSessionEntryList{
				UserId:       id,
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

	return rValues, nil
}
