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
	"google.golang.org/protobuf/types/known/wrapperspb"

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

	mutableUsers, err := c.svcCtx.Dao.UserClient.UserGetMutableUsersV2(c.ctx, &userpb.TLUserGetMutableUsersV2{
		Id:      []int64{c.MD.UserId, peerId},
		Privacy: true,
		HasTo:   true,
		To:      []int64{c.MD.UserId, peerId},
	})
	if err != nil {
		c.Logger.Errorf("users.getFullUser - error: %v", err)
		return nil, err
	} else if len(mutableUsers.GetUsers()) == 0 {
		err = mtproto.ErrInternalServerError
		c.Logger.Errorf("users.getFullUser - error: %v", err)
		return nil, err
	}

	me, _ = mutableUsers.GetImmutableUser(c.MD.UserId)
	user, _ = mutableUsers.GetImmutableUser(peerId)

	if user == nil {
		err = mtproto.ErrInternalServerError
		c.Logger.Errorf("users.getFullUser - error: %v", err)
		return nil, err
	}

	// userFull#3fd81e28 flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true can_pin_message:flags.7?true has_scheduled:flags.12?true video_calls_available:flags.13?true voice_messages_forbidden:flags.20?true translations_disabled:flags.23?true stories_pinned_available:flags.26?true blocked_my_stories_from:flags.27?true wallpaper_overridden:flags.28?true contact_require_premium:flags.29?true read_dates_private:flags.30?true flags2:# sponsored_enabled:flags2.7?true can_view_revenue:flags2.9?true bot_can_manage_emoji_status:flags2.10?true display_gifts_button:flags2.16?true id:long about:flags.1?string settings:PeerSettings personal_photo:flags.21?Photo profile_photo:flags.2?Photo fallback_photo:flags.22?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo pinned_msg_id:flags.6?int common_chats_count:int folder_id:flags.11?int ttl_period:flags.14?int theme_emoticon:flags.15?string private_forward_name:flags.16?string bot_group_admin_rights:flags.17?ChatAdminRights bot_broadcast_admin_rights:flags.18?ChatAdminRights wallpaper:flags.24?WallPaper stories:flags.25?PeerStories business_work_hours:flags2.0?BusinessWorkHours business_location:flags2.1?BusinessLocation business_greeting_message:flags2.2?BusinessGreetingMessage business_away_message:flags2.3?BusinessAwayMessage business_intro:flags2.4?BusinessIntro birthday:flags2.5?Birthday personal_channel_id:flags2.6?long personal_channel_message:flags2.6?int stargifts_count:flags2.8?int starref_program:flags2.11?StarRefProgram bot_verification:flags2.12?BotVerification send_paid_messages_stars:flags2.14?long disallowed_gifts:flags2.15?DisallowedGiftsSettings stars_rating:flags2.17?StarsRating stars_my_pending_rating:flags2.18?StarsRating stars_my_pending_rating_date:flags2.18?int main_tab:flags2.20?ProfileTab saved_music:flags2.21?Document = UserFull;
	// userFull#a02bc13e flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true can_pin_message:flags.7?true has_scheduled:flags.12?true video_calls_available:flags.13?true voice_messages_forbidden:flags.20?true translations_disabled:flags.23?true stories_pinned_available:flags.26?true blocked_my_stories_from:flags.27?true wallpaper_overridden:flags.28?true contact_require_premium:flags.29?true read_dates_private:flags.30?true flags2:# sponsored_enabled:flags2.7?true can_view_revenue:flags2.9?true bot_can_manage_emoji_status:flags2.10?true display_gifts_button:flags2.16?true id:long about:flags.1?string settings:PeerSettings personal_photo:flags.21?Photo profile_photo:flags.2?Photo fallback_photo:flags.22?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo pinned_msg_id:flags.6?int common_chats_count:int folder_id:flags.11?int ttl_period:flags.14?int theme:flags.15?ChatTheme private_forward_name:flags.16?string bot_group_admin_rights:flags.17?ChatAdminRights bot_broadcast_admin_rights:flags.18?ChatAdminRights wallpaper:flags.24?WallPaper stories:flags.25?PeerStories business_work_hours:flags2.0?BusinessWorkHours business_location:flags2.1?BusinessLocation business_greeting_message:flags2.2?BusinessGreetingMessage business_away_message:flags2.3?BusinessAwayMessage business_intro:flags2.4?BusinessIntro birthday:flags2.5?Birthday personal_channel_id:flags2.6?long personal_channel_message:flags2.6?int stargifts_count:flags2.8?int starref_program:flags2.11?StarRefProgram bot_verification:flags2.12?BotVerification send_paid_messages_stars:flags2.14?long disallowed_gifts:flags2.15?DisallowedGiftsSettings stars_rating:flags2.17?StarsRating stars_my_pending_rating:flags2.18?StarsRating stars_my_pending_rating_date:flags2.18?int main_tab:flags2.20?ProfileTab saved_music:flags2.21?Document note:flags2.22?TextWithEntities = UserFull;
	userFull := mtproto.MakeTLUserFull(&mtproto.UserFull{
		Blocked:                  false,
		PhoneCallsAvailable:      c.MD.UserId != peerId,
		PhoneCallsPrivate:        false,
		CanPinMessage:            true,
		HasScheduled:             false,
		VideoCallsAvailable:      c.MD.UserId != peerId,
		VoiceMessagesForbidden:   false,
		TranslationsDisabled:     false,
		StoriesPinnedAvailable:   false,
		BlockedMyStoriesFrom:     false,
		WallpaperOverridden:      false,
		Id:                       peerId,
		About:                    user.GetUser().GetAbout(),
		Settings:                 nil,
		PersonalPhoto:            nil,
		ProfilePhoto:             user.GetUser().GetProfilePhoto(),
		FallbackPhoto:            nil,
		NotifySettings:           nil,
		BotInfo:                  nil,
		PinnedMsgId:              nil,
		CommonChatsCount:         0,
		FolderId:                 nil,
		TtlPeriod:                nil,
		ThemeEmoticon:            nil,
		Theme:                    nil,
		PrivateForwardName:       nil,
		BotGroupAdminRights:      nil,
		BotBroadcastAdminRights:  nil,
		PremiumGifts:             nil,
		Wallpaper:                nil,
		Stories_FLAGPEERSTORIES:  nil,
		Stories_FLAGUSERSTORIES:  nil,
		BusinessWorkHours:        nil,
		BusinessLocation:         nil,
		BusinessGreetingMessage:  nil,
		BusinessAwayMessage:      nil,
		BusinessIntro:            nil,
		Birthday:                 nil,
		PersonalChannelId:        nil,
		PersonalChannelMessage:   nil,
		StargiftsCount:           nil,
		StarrefProgram:           nil,
		BotVerification:          nil,
		SendPaidMessagesStars:    nil,
		DisallowedGifts:          nil,
		StarsMyPendingRating:     nil,
		StarsMyPendingRatingDate: nil,
		MainTab:                  user.GetUser().GetMainTab(),
		SavedMusic:               nil,
		Note:                     nil,
	}).To_UserFull()

	// PremiumGifts
	if user.Premium() {
		// TODO: config able
		userFull.PremiumGifts = []*mtproto.PremiumGiftOption{
			mtproto.MakeTLPremiumGiftOption(&mtproto.PremiumGiftOption{
				Months:       12,
				Currency:     "CNY",
				Amount:       20900,
				BotUrl:       "https://t.me/$premgift448603711_12_5248da16f536f717a2",
				StoreProduct: mtproto.MakeFlagsString("org.telegram.telegramPremium.twelveMonths"),
			}).To_PremiumGiftOption(),
			mtproto.MakeTLPremiumGiftOption(&mtproto.PremiumGiftOption{
				Months:       6,
				Currency:     "CNY",
				Amount:       10900,
				BotUrl:       "https://t.me/$premgift448603711_6_c7aae8edbdae927b72",
				StoreProduct: mtproto.MakeFlagsString("org.telegram.telegramPremium.sixMonths"),
			}).To_PremiumGiftOption(),
			mtproto.MakeTLPremiumGiftOption(&mtproto.PremiumGiftOption{
				Months:       3,
				Currency:     "CNY",
				Amount:       8499,
				BotUrl:       "https://t.me/$premgift448603711_3_051b80db4901b91dd5",
				StoreProduct: mtproto.MakeFlagsString("org.telegram.telegramPremium.threeMonths"),
			}).To_PremiumGiftOption(),
		}
	}
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
				userFull.VideoCallsAvailable = false
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
				if usersChatIdList != nil && len(usersChatIdList.Datas) == 2 {
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
					if dialogExt.WallpaperId != 0 && c.svcCtx.Dao.WallpaperPlugin != nil {
						userFull.Wallpaper = c.svcCtx.Dao.WallpaperPlugin.GetChatWallpaper(c.ctx, c.MD.UserId, dialogExt.WallpaperId)
						userFull.WallpaperOverridden = true
					}
				}
			}
		},
		func() {
			rules, _ := c.svcCtx.Dao.UserClient.UserGetPrivacy(c.ctx, &userpb.TLUserGetPrivacy{
				UserId:  peerId,
				KeyType: mtproto.VOICE_MESSAGES,
			})
			if rules != nil && len(rules.Datas) > 0 {
				allow := mtproto.CheckPrivacyIsAllow(
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
		},
		func() {
			if user.GetUser().GetSavedMusic() != nil {
				rules, _ := c.svcCtx.Dao.UserClient.UserGetPrivacy(c.ctx, &userpb.TLUserGetPrivacy{
					UserId:  peerId,
					KeyType: mtproto.SAVED_MUSIC,
				})
				if rules != nil && len(rules.Datas) > 0 {
					allow := mtproto.CheckPrivacyIsAllow(
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
					if allow {
						userFull.SavedMusic = user.GetUser().GetSavedMusic()
					}
				}
			}
		})

	// TODO: FolderId:    0,

	// TODO: WallPaper

	// TODO: Stories
	if c.svcCtx.Dao.StoryPlugin != nil {
		userFull.StoriesPinnedAvailable = c.svcCtx.Dao.StoryPlugin.GetStoriesPinnedAvailable(c.ctx, peerId, c.MD.UserId)
		userFull.BlockedMyStoriesFrom = c.svcCtx.Dao.StoryPlugin.GetBlockedMyStoriesFrom(c.ctx, peerId, c.MD.UserId)
		// c.Logger.Debugf("getActiveStories: peerId: %s, userId: %s", user, me)
		if peerId == c.MD.UserId {
			// c.Logger.Debugf("getActiveStories(peerId == c.MD.UserId): peerId: %d, userId: %d", peerId, c.MD.UserId)
			userFull.Stories_FLAGPEERSTORIES = c.svcCtx.Dao.StoryPlugin.GetActiveStories(c.ctx, peerId, c.MD.UserId)
		} else if ok, _ := me.CheckReverseContact(peerId); ok {
			// c.Logger.Debugf("getActiveStories(ok, _ := user.CheckContact(c.MD.UserId)): peerId: %d, userId: %d", peerId, c.MD.UserId)
			userFull.Stories_FLAGPEERSTORIES = c.svcCtx.Dao.StoryPlugin.GetActiveStories(c.ctx, peerId, c.MD.UserId)
		}
	}

	chats := make([]*mtproto.Chat, 0)

	if c.svcCtx.Dao.PersonalChannelPlugin != nil {
		personalChannelId := user.GetUser().GetPersonalChannelId()
		if personalChannelId != 0 {
			userFull.PersonalChannelId = mtproto.MakeFlagsInt64(personalChannelId)
			pChannel, topMessageId := c.svcCtx.Dao.PersonalChannelPlugin.GetPersonalChannel(c.ctx, personalChannelId, c.MD.UserId)
			if pChannel != nil {
				userFull.PersonalChannelMessage = &wrapperspb.Int32Value{Value: topMessageId}
				chats = append(chats, pChannel)
			}
		}
	}

	if c.MD.UserId != peerId {
		// TODO
	} else {
		userFull.Birthday = user.Birthday()
	}

	if c.MD.UserId == 777000 {
		userFull.PhoneCallsAvailable = false
		userFull.PhoneCallsPrivate = false
		userFull.VideoCallsAvailable = false
	}

	return mtproto.MakeTLUsersUserFull(&mtproto.Users_UserFull{
		FullUser: userFull,
		Chats:    chats,
		Users:    []*mtproto.User{user.ToUnsafeUser(me)},
	}).To_Users_UserFull(), nil
}
