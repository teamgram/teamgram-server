// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountGetPrivacy
// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (c *PrivacySettingsCore) AccountGetPrivacy(in *tg.TLAccountGetPrivacy) (*tg.AccountPrivacyRules, error) {
	key := tg.FromInputPrivacyKeyType(&tg.InputPrivacyKey{Clazz: in.Key})

	if key == tg.KEY_TYPE_INVALID {
		c.Logger.Errorf("account.getPrivacy - error: invalid privacy key")
		return nil, tg.ErrPrivacyKeyInvalid
	}

	ruleList, err := c.svcCtx.Repo.UserClient.UserGetPrivacy(c.ctx, &user.TLUserGetPrivacy{
		UserId:  c.MD.UserId,
		KeyType: int32(key),
	})
	if err != nil {
		c.Logger.Errorf("account.getPrivacy - UserGetPrivacy error: %v", err)
		return nil, err
	}

	var rVal *tg.AccountPrivacyRules

	if len(ruleList.Datas) == 0 {
		if key == tg.PHONE_NUMBER {
			rVal = tg.MakeTLAccountPrivacyRules(&tg.TLAccountPrivacyRules{
				Rules: []tg.PrivacyRuleClazz{&tg.TLPrivacyValueDisallowAll{}},
				Users: []tg.UserClazz{},
				Chats: []tg.ChatClazz{},
			}).ToAccountPrivacyRules()
		} else if key == tg.BIRTHDAY {
			rVal = tg.MakeTLAccountPrivacyRules(&tg.TLAccountPrivacyRules{
				Rules: []tg.PrivacyRuleClazz{&tg.TLPrivacyValueAllowContacts{}},
				Users: []tg.UserClazz{},
				Chats: []tg.ChatClazz{},
			}).ToAccountPrivacyRules()
		} else {
			rVal = tg.MakeTLAccountPrivacyRules(&tg.TLAccountPrivacyRules{
				Rules: []tg.PrivacyRuleClazz{&tg.TLPrivacyValueAllowAll{}},
				Users: []tg.UserClazz{},
				Chats: []tg.ChatClazz{},
			}).ToAccountPrivacyRules()
		}
	} else {
		rVal = tg.MakeTLAccountPrivacyRules(&tg.TLAccountPrivacyRules{
			Rules: ruleList.Datas,
			Users: []tg.UserClazz{},
			Chats: []tg.ChatClazz{},
		}).ToAccountPrivacyRules()

		var (
			userIds    []int64
			chatIds    []int64
			channelIds []int64
		)

		for _, r := range ruleList.Datas {
			switch rv := r.(type) {
			case *tg.TLPrivacyValueAllowUsers:
				userIds = append(userIds, rv.Users...)
			case *tg.TLPrivacyValueDisallowUsers:
				userIds = append(userIds, rv.Users...)
			case *tg.TLPrivacyValueAllowChatParticipants:
				for _, id := range rv.Chats {
					if id >= tg.MinNebulaChatChannelID {
						channelIds = append(channelIds, id)
					} else {
						chatIds = append(chatIds, id)
					}
				}
			case *tg.TLPrivacyValueDisallowChatParticipants:
				for _, id := range rv.Chats {
					if id >= tg.MinNebulaChatChannelID {
						channelIds = append(channelIds, id)
					} else {
						chatIds = append(chatIds, id)
					}
				}
			}
		}

		if len(userIds) > 0 {
			users, err := c.svcCtx.Repo.UserClient.UserGetMutableUsers(c.ctx,
				&user.TLUserGetMutableUsers{
					Id: userIds,
				})
			if err != nil {
				c.Logger.Errorf("account.getPrivacy - get users error: %v", err)
			} else {
				for _, u := range users.Datas {
					rVal.Users = append(rVal.Users, projectImmutableUser(u))
				}
			}
		}

		if len(chatIds) > 0 {
			chats, err := c.svcCtx.Repo.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chat.TLChatGetChatListByIdList{
					SelfId: c.MD.UserId,
					IdList: chatIds,
				})
			if err != nil {
				c.Logger.Errorf("account.getPrivacy - get chats error: %v", err)
			} else {
				for _, ch := range chats.Datas {
					rVal.Chats = append(rVal.Chats, projectMutableChat(ch, c.MD.UserId))
				}
			}
		}

		_ = channelIds // TODO: handle channels
	}

	return rVal, nil
}
