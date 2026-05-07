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
	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountSetPrivacy
// account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
func (c *PrivacySettingsCore) AccountSetPrivacy(in *tg.TLAccountSetPrivacy) (*tg.AccountPrivacyRules, error) {
	key := tg.FromInputPrivacyKeyType(&tg.InputPrivacyKey{Clazz: in.Key})

	if key == tg.KEY_TYPE_INVALID {
		c.Logger.Errorf("account.setPrivacy - error: invalid privacy key")
		return nil, tg.ErrPrivacyKeyInvalid
	}

	ruleList := tg.ToPrivacyRuleListByInput(c.MD.UserId, in.Rules)

	if _, err := c.svcCtx.Repo.UserClient.UserSetPrivacy(c.ctx, &user.TLUserSetPrivacy{
		UserId:  c.MD.UserId,
		KeyType: int32(key),
		Rules:   ruleList,
	}); err != nil {
		return nil, err
	}

	rVal := tg.MakeTLAccountPrivacyRules(&tg.TLAccountPrivacyRules{
		Rules: ruleList,
		Users: []tg.UserClazz{},
		Chats: []tg.ChatClazz{},
	}).ToAccountPrivacyRules()

	var (
		userIds    []int64
		chatIds    []int64
		channelIds []int64
	)

	for _, r := range ruleList {
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
		users, err := userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, c.MD.UserId, userIds, userprojection.MissingStoredReference)
		if err != nil {
			c.Logger.Errorf("account.setPrivacy - resolve users error: %v", err)
		} else {
			rVal.Users = append(rVal.Users, users...)
		}
	}

	if len(chatIds) > 0 {
		chats, err := c.svcCtx.Repo.ChatClient.ChatGetChatListByIdList(c.ctx,
			&chat.TLChatGetChatListByIdList{
				SelfId: c.MD.UserId,
				IdList: chatIds,
			})
		if err != nil {
			c.Logger.Errorf("account.setPrivacy - resolve chats error: %v", err)
		} else {
			for _, ch := range chats.Datas {
				rVal.Chats = append(rVal.Chats, projectMutableChat(ch, c.MD.UserId))
			}
		}
	}

	_ = channelIds // TODO: handle channels

	// TODO: syncUpdatesNotMe
	// syncUpdates := tg.MakeUpdatesByUpdates(tg.MakeTLUpdatePrivacy(&tg.TLUpdate{
	//     Key:   tg.ToPrivacyKey(key),
	//     Rules: ruleList,
	// }).ToUpdate())
	//
	// for _, u := range rVal.Users {
	//     syncUpdates.PushUser(u)
	// }
	// for _, ch := range rVal.Chats {
	//     syncUpdates.PushChat(ch)
	// }
	//
	// c.svcCtx.Repo.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
	//     UserId:        c.MD.UserId,
	//     PermAuthKeyId: c.MD.PermAuthKeyId,
	//     Updates:       syncUpdates,
	// })

	return rVal, nil
}
