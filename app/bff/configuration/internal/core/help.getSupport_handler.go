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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	// TODO(@benqi): add support user.
	supportUserID      = int64(424000)
	supportPhoneNumber = "42400"
	supportName        = "Volunteer Support"
)

var (
	supportUserAccessHash = int64(6599886787491911852)
	supportUserFirstName  = supportName
	supportUser           = tg.MakeTLUser(&tg.TLUser{
		Self:                 false,
		Contact:              false,
		MutualContact:        false,
		Deleted:              false,
		Bot:                  false, // true
		BotChatHistory:       false,
		BotNochats:           false, // true
		Verified:             false,
		Restricted:           false,
		Min:                  false,
		BotInlineGeo:         false,
		Support:              true,
		Scam:                 false,
		Id:                   supportUserID,
		AccessHash:           &supportUserAccessHash,
		FirstName:            &supportUserFirstName,
		LastName:             nil,
		Username:             nil,
		Phone:                nil,
		Photo:                nil,
		Status:               nil,
		BotInfoVersion:       nil,
		RestrictionReason:    nil,
		BotInlinePlaceholder: nil,
		LangCode:             nil,
	})
)

// HelpGetSupport
// help.getSupport#9cdf08cd = help.Support;
func (c *ConfigurationCore) HelpGetSupport(in *tg.TLHelpGetSupport) (*tg.HelpSupport, error) {
	_ = in

	//mUser, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
	//	Id: []int64{c.MD.UserId, supportUserID},
	//})
	//
	//me, _ := mUser.GetUnsafeUser(c.MD.UserId, supportUserID)

	rValue := tg.MakeTLHelpSupport(&tg.TLHelpSupport{
		PhoneNumber: supportPhoneNumber,
		User:        supportUser,
	}).ToHelpSupport()

	return rValue, nil
}
