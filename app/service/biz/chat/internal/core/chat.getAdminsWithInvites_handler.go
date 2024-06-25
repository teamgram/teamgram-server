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
	"github.com/teamgram/marmota/pkg/container2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
)

// ChatGetAdminsWithInvites
// chat.getAdminsWithInvites self_id:long chat_id:long = Vector<ChatAdminWithInvites>;
func (c *ChatCore) ChatGetAdminsWithInvites(in *chat.TLChatGetAdminsWithInvites) (*chat.Vector_ChatAdminWithInvites, error) {
	var (
		rAdmins        []*mtproto.ChatAdminWithInvites
		canInviteUsers []int64
	)

	chat2, err := c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId)
	if err != nil {
		c.Logger.Errorf("chat.getAdminsWithInvites - error: %v", err)
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}
	me, _ := chat2.GetImmutableChatParticipant(in.SelfId)
	if me == nil {
		c.Logger.Errorf("chat.getAdminsWithInvites - error: not existed chat")
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}
	if !me.CanInviteUsers() {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.getAdminsWithInvites - error: %v", err)
		return nil, err
	}

	chat2.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
		if participant.CanInviteUsers() {
			canInviteUsers = append(canInviteUsers, participant.UserId)
		}
		return nil
	})

	c.svcCtx.Dao.ChatInvitesDAO.SelectListByChatIdWithCB(
		c.ctx,
		in.ChatId,
		func(sz, i int, v *dataobject.ChatInvitesDO) {
			var (
				admin *mtproto.ChatAdminWithInvites
			)

			if ok := container2.ContainsInt64(canInviteUsers, v.AdminId); !ok {
				return
			}

			for _, a := range rAdmins {
				if a.AdminId == v.AdminId {
					admin = a
					break
				}
			}
			if admin == nil {
				admin = mtproto.MakeTLChatAdminWithInvites(&mtproto.ChatAdminWithInvites{
					AdminId:             v.AdminId,
					InvitesCount:        0,
					RevokedInvitesCount: 0,
				}).To_ChatAdminWithInvites()
				rAdmins = append(rAdmins, admin)
			}
			if v.Revoked {
				admin.RevokedInvitesCount++
			} else {
				admin.InvitesCount++
			}
		})

	if rAdmins == nil {
		rAdmins = []*mtproto.ChatAdminWithInvites{}
	}

	return &chat.Vector_ChatAdminWithInvites{
		Datas: rAdmins,
	}, nil
}
