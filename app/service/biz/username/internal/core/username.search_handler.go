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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/username/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
)

// UsernameSearch
// username.search q:string excluded_contacts:Vector<long> limit:int = Vector<UsernameData>;
func (c *UsernameCore) UsernameSearch(in *username.TLUsernameSearch) (*username.Vector_UsernameData, error) {
	var (
		rValList = &username.Vector_UsernameData{}
	)

	// Check query string and limit
	if len(in.Q) < 3 || in.Limit <= 0 {
		return rValList, nil
	}

	if in.Limit > 50 {
		in.Limit = 50
	}

	// 构造模糊查询字符串
	q2 := in.Q + "%"
	doList, _ := c.svcCtx.Dao.UsernameDAO.SearchByQueryNotIdListWithCB(
		c.ctx,
		q2,
		in.ExcludedContacts,
		in.Limit,
		func(sz, i int, v *dataobject.UsernameDO) {
			switch v.PeerType {
			case mtproto.PEER_USER:
				rValList.Datas = append(rValList.Datas, username.MakeTLUsernameData(&username.UsernameData{
					Username: v.Username,
					Peer:     mtproto.MakePeerUser(v.PeerId),
				}).To_UsernameData())
			case mtproto.PEER_CHANNEL:
				rValList.Datas = append(rValList.Datas, username.MakeTLUsernameData(&username.UsernameData{
					Username: v.Username,
					Peer:     mtproto.MakePeerChannel(v.PeerId),
				}).To_UsernameData())
			}
		})

	c.Logger.Infof("username.search - doList: %v", doList)
	return rValList, nil
}
