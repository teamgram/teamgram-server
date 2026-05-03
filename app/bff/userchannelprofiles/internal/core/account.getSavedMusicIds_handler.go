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
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountGetSavedMusicIds
// account.getSavedMusicIds#e09d5faf hash:long = account.SavedMusicIds;
func (c *UserChannelProfilesCore) AccountGetSavedMusicIds(in *tg.TLAccountGetSavedMusicIds) (*tg.AccountSavedMusicIds, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if err := requireUserClient(c); err != nil {
		return nil, err
	}

	idList, err := c.svcCtx.Repo.UserClient.UserGetSavedMusicIdList(c.ctx, &userpb.TLUserGetSavedMusicIdList{
		UserId: selfID,
	})
	if err != nil {
		return nil, err
	}
	ids := []int64{}
	if idList != nil {
		ids = idList.Datas
	}
	return tg.MakeTLAccountSavedMusicIds(&tg.TLAccountSavedMusicIds{
		Ids: ids,
	}).ToAccountSavedMusicIds(), nil
}
