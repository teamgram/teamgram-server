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

// StatusGetUserOnlineSessions
// status.getUserOnlineSessions user_id:long = UserSessionEntryList;
func (c *StatusCore) StatusGetUserOnlineSessions(in *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error) {
	rMap, err := c.svcCtx.Dao.KV.HgetallCtx(c.ctx, getUserKey(in.GetUserId()))
	if err != nil {
		c.Logger.Errorf("status.getUserOnlineSessions(%s) error(%v)", in, err)
		return nil, err
	}

	var (
		rValues = status.MakeTLUserSessionEntryList(&status.UserSessionEntryList{
			UserId:       in.UserId,
			UserSessions: make([]*status.SessionEntry, 0, len(rMap)),
		}).To_UserSessionEntryList()
	)

	for _, v := range rMap {
		sess := new(status.SessionEntry)
		if err2 := jsonx.UnmarshalFromString(v, sess); err2 == nil {
			rValues.UserSessions = append(rValues.UserSessions, sess)
		}
	}

	return rValues, nil
}
