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

// StatusGetUserOnlineSessions
// status.getUserOnlineSessions user_id:long = UserSessionEntryList;
func (c *StatusCore) StatusGetUserOnlineSessions(in *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error) {
	rMap, err := c.svcCtx.KV.Hgetall(getUserKey(in.GetUserId()))
	if err != nil {
		c.Logger.Errorf("status.getUserOnlineSessions(%s) error(%v)", in.DebugString(), err)
		return nil, err
	}

	var (
		rValues = status.MakeTLUserSessionEntryList(&status.UserSessionEntryList{
			UserId:       in.UserId,
			UserSessions: make([]*status.SessionEntry, 0, len(rMap)),
		}).To_UserSessionEntryList()
	)

	for k, v := range rMap {
		keyId, _ := strconv.ParseInt(k, 10, 64)
		rValues.UserSessions = append(rValues.UserSessions, &status.SessionEntry{
			UserId:    in.GetUserId(),
			AuthKeyId: keyId,
			Gateway:   v,
			Expired:   0,
			Layer:     0,
		})
	}

	return rValues, nil
}
