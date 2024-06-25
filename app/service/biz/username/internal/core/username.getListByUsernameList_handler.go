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

// UsernameGetListByUsernameList
// username.getListByUsernameList names:Vector<string> = Vector<UsernameData>;
func (c *UsernameCore) UsernameGetListByUsernameList(in *username.TLUsernameGetListByUsernameList) (*username.Vector_UsernameData, error) {
	var (
		rValues = &username.Vector_UsernameData{
			Datas: []*username.UsernameData{},
		}
	)

	if _, err := c.svcCtx.Dao.UsernameDAO.SelectListWithCB(c.ctx, in.Names, func(sz, i int, v *dataobject.UsernameDO) {
		var (
			peer *mtproto.Peer
		)

		switch v.PeerType {
		case mtproto.PEER_USER:
			peer = mtproto.MakePeerUser(v.PeerId)
		case mtproto.PEER_CHANNEL:
			peer = mtproto.MakePeerChannel(v.PeerId)
		default:
			return
		}

		rValues.Datas = append(rValues.Datas, username.MakeTLUsernameData(&username.UsernameData{
			Username: v.Username,
			Peer:     peer,
		}).To_UsernameData())
	}); err != nil {
		c.Logger.Errorf("username.getListByUsernameList - error: %v", err)
	}

	return rValues, nil
}
