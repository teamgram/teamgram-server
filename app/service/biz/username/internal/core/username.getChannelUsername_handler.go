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

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
)

// UsernameGetChannelUsername
// username.getChannelUsername channel_id:int = UsernameData;
func (c *UsernameCore) UsernameGetChannelUsername(in *username.TLUsernameGetChannelUsername) (*username.UsernameData, error) {
	do, _ := c.svcCtx.Dao.UsernameDAO.SelectByPeer(c.ctx, mtproto.PEER_CHANNEL, in.ChannelId)
	if do == nil {
		err := fmt.Errorf("channel_id's username empty")
		c.Logger.Errorf("username.getChannelUsername - error: %v", err)
		return nil, err
	}

	return username.MakeTLUsernameData(&username.UsernameData{
		Username: do.Username,
		Peer:     nil,
	}).To_UsernameData(), nil
}
