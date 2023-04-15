// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/marmota/pkg/utils"
	"github.com/teamgram/proto/mtproto"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/zeromicro/go-zero/core/mr"
)

// UsersGetFullUser
// users.getFullUser#b60f5918 id:InputUser = users.UserFull;
func (c *UsersCore) UsersGetFullUser(in *mtproto.TLUsersGetFullUser) (*mtproto.Users_UserFull, error) {
	var (
		peerId int64
		id     = mtproto.FromInputUser(c.MD.UserId, in.Id)
		me     *mtproto.ImmutableUser
		user   *mtproto.ImmutableUser
	)

	switch id.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER:
		peerId = id.PeerId
	default:
		err := mtproto.ErrUserIdInvalid
		c.Logger.Errorf("users.getFullUser - error: %v", err)
		return nil, err
	}

	mutableUsers, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, peerId},
	})
	if err != nil {
		c.Logger.Errorf("users.getFullUser - error: %v", err)
		return nil, err
	} else if len(mutableUsers.Datas) == 0 {
		err = mtproto.ErrInternelServerError
		c.Logger.Errorf("users.getFullUser - error: %v", err)
		return nil, err
	}

	me, _ = mutableUsers.GetImmutableUser(c.MD.UserId)
	user, _ = mutableUsers.GetImmutableUser(peerId)

	if user == nil {
		err = mtproto.ErrInternelServerError
		c.Logger.Errorf("users.getFullUser - error: %v", err)
		return nil, err
	}

	userFull := mtproto.MakeTLUserFull(&mtproto.UserFull{
		Blocked:                 false,
		PhoneCallsAvailable:     true,
		PhoneCallsPrivate:       false,
		CanPinMessage:           true,
		HasScheduled:            false,
		VideoCallsAvailable:     true,
		VoiceMessagesForbidden:  false,
		TranslationsDisabled:    false,
		Id:                      peerId,
		About:                   user.GetUser().GetAbout(),
		Settings:                nil,
		PersonalPhoto:           nil,
		ProfilePhoto:            user.GetUser().GetProfilePhoto(),
		FallbackPhoto:           nil,
		NotifySettings:          nil,
		BotInfo:                 nil,
		PinnedMsgId:             nil,
		CommonChatsCount:        0,
		FolderId:                nil,
		TtlPeriod:               nil,
		ThemeEmoticon:           nil,
		PrivateForwardName:      nil,
		BotGroupAdminRights:     nil,
		BotBroadcastAdminRights: nil,
		PremiumGifts:            nil,
	}).To_UserFull()

	mr.FinishVoid(
		func() {
			// blocked
			if c.MD.UserId != peerId {
				blocked, _ := c.svcCtx.Dao.UserClient.UserBlockedByUser(
					c.ctx,
					&userpb.TLUserBlockedByUser{
						UserId:     c.MD.UserId,
						PeerUserId: peerId,
					})
				userFull.Blocked = mtproto.FromBool(blocked)
			}
		},
		func() {
			userFull.Settings, _ = c.svcCtx.Dao.UserClient.UserGetPeerSettings(c.ctx, &userpb.TLUserGetPeerSettings{
				UserId:   c.MD.UserId,
				PeerType: mtproto.PEER_USER,
				PeerId:   peerId,
			})
		},
		func() {
			userFull.NotifySettings, _ = c.svcCtx.Dao.UserClient.UserGetNotifySettings(c.ctx, &userpb.TLUserGetNotifySettings{
				UserId:   c.MD.UserId,
				PeerType: mtproto.PEER_USER,
				PeerId:   peerId,
			})
		},
		func() {
			if user.GetUser().GetBot() != nil {
				userFull.PhoneCallsAvailable = false
				userFull.PhoneCallsPrivate = false
				userFull.BotInfo, _ = c.svcCtx.Dao.UserClient.UserGetBotInfo(c.ctx, &userpb.TLUserGetBotInfo{
					BotId: peerId,
				})
			}
		},
		func() {
			// TODO: PinnedMsgId:         nil,
			if c.MD.UserId != peerId {
				usersChatIdList, _ := c.svcCtx.Dao.ChatClient.ChatGetUsersChatIdList(c.ctx, &chatpb.TLChatGetUsersChatIdList{
					Id: []int64{c.MD.UserId, peerId},
				})
				if len(usersChatIdList.Datas) == 2 {
					commonChats := utils.Int64Intersect(
						usersChatIdList.Datas[0].ChatIdList,
						usersChatIdList.Datas[1].ChatIdList)
					userFull.CommonChatsCount = int32(len(commonChats))
				}

				// TODO: Fetch CommonChannelsCount
			}
		},
		func() {
			if peerId != c.MD.UserId {
				// theme_emoticon
				dialogExt, _ := c.svcCtx.Dao.DialogClient.DialogGetDialogById(c.ctx, &dialog.TLDialogGetDialogById{
					UserId:   c.MD.UserId,
					PeerType: mtproto.PEER_USER,
					PeerId:   peerId,
				})
				if dialogExt != nil {
					userFull.ThemeEmoticon = mtproto.MakeFlagsString(dialogExt.ThemeEmoticon)
					userFull.TtlPeriod = mtproto.MakeFlagsInt32(dialogExt.TtlPeriod)
				}
			}
		},
		func() {
			rules, _ := c.svcCtx.Dao.UserClient.UserGetPrivacy(c.ctx, &userpb.TLUserGetPrivacy{
				UserId:  peerId,
				KeyType: userpb.VOICE_MESSAGES,
			})
			if len(rules.Datas) > 0 {
				allow := userpb.CheckPrivacyIsAllow(
					peerId,
					rules.Datas,
					c.MD.UserId,
					func(id, checkId int64) bool {
						contact, _ := user.CheckContact(checkId)
						return contact
					},
					func(checkId int64, idList []int64) bool {
						// TODO
						chatIdList, _ := mtproto.SplitChatAndChannelIdList(idList)
						_ = chatIdList
						// return c.svcCtx.Dao.ChatClient.CheckParticipantIsExist(c.ctx, checkId, chatIdList)
						return false
					})
				userFull.VoiceMessagesForbidden = !allow
			}
		})

	// TODO: FolderId:    0,

	return mtproto.MakeTLUsersUserFull(&mtproto.Users_UserFull{
		FullUser: userFull,
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{user.ToUnsafeUser(me)},
	}).To_Users_UserFull(), nil
}
