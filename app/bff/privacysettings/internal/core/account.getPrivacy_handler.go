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
	"github.com/teamgram/proto/mtproto"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// AccountGetPrivacy
// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (c *PrivacySettingsCore) AccountGetPrivacy(in *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	var (
		key  = mtproto.FromInputPrivacyKeyType(in.Key)
		rVal *mtproto.Account_PrivacyRules
	)

	// TODO(@benqi): Check request valid.
	if key == mtproto.KEY_TYPE_INVALID {
		err := mtproto.ErrPrivacyKeyInvalid
		c.Logger.Errorf("account.getPrivacy - error: %v", err)
		return nil, err
	}

	ruleList, _ := c.svcCtx.Dao.UserClient.UserGetPrivacy(c.ctx, &userpb.TLUserGetPrivacy{
		UserId:  c.MD.UserId,
		KeyType: int32(key),
	})

	if len(ruleList.GetDatas()) == 0 {
		if key == mtproto.PHONE_NUMBER {
			rVal = mtproto.MakeTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
				Rules: []*mtproto.PrivacyRule{mtproto.MakeTLPrivacyValueDisallowAll(nil).To_PrivacyRule()},
				Users: []*mtproto.User{},
				Chats: []*mtproto.Chat{},
			}).To_Account_PrivacyRules()
		} else if key == mtproto.BIRTHDAY {
			// Birthday default privacy rules is allow contacts
			rVal = mtproto.MakeTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
				Rules: []*mtproto.PrivacyRule{mtproto.MakeTLPrivacyValueAllowContacts(nil).To_PrivacyRule()},
				Users: []*mtproto.User{},
				Chats: []*mtproto.Chat{},
			}).To_Account_PrivacyRules()
		} else {
			rVal = mtproto.MakeTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
				Rules: []*mtproto.PrivacyRule{mtproto.MakeTLPrivacyValueAllowAll(nil).To_PrivacyRule()},
				Users: []*mtproto.User{},
				Chats: []*mtproto.Chat{},
			}).To_Account_PrivacyRules()
		}
	} else {
		rVal = mtproto.MakeTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
			Rules: ruleList.GetDatas(),
			Users: []*mtproto.User{}, // TODO
			Chats: []*mtproto.Chat{}, // TODO
		}).To_Account_PrivacyRules()

		idHelper := mtproto.NewIDListHelper(c.MD.UserId)
		idHelper.PickByRules(ruleList.GetDatas())
		idHelper.Visit(
			func(userIdList []int64) {
				users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
					&userpb.TLUserGetMutableUsers{
						Id: userIdList,
					})
				rVal.Users = users.GetUserListByIdList(c.MD.UserId, userIdList...)
			},
			func(chatIdList []int64) {
				chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
					&chatpb.TLChatGetChatListByIdList{
						IdList: chatIdList,
					})
				rVal.Chats = chats.GetChatListByIdList(c.MD.UserId, chatIdList...)
			},

			func(channelIdList []int64) {
				// TODO
			})
	}

	return rVal, nil
}
