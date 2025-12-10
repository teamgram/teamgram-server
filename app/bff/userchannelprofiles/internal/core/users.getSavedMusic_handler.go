// Copyright 2025 Teamgram Authors
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
// limitations under the License.s
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// UsersGetSavedMusic
// users.getSavedMusic#788d7fe3 id:InputUser offset:int limit:int hash:long = users.SavedMusic;
func (c *UserChannelProfilesCore) UsersGetSavedMusic(in *mtproto.TLUsersGetSavedMusic) (*mtproto.Users_SavedMusic, error) {
	peer := mtproto.FromInputUser(c.MD.UserId, in.GetId())

	idList, err := c.svcCtx.Dao.UserClient.UserGetSavedMusicIdList(c.ctx, &user.TLUserGetSavedMusicIdList{
		UserId: peer.PeerId,
	})
	if err != nil {
		c.Logger.Errorf("users.getSavedMusic - error: %v", err)
		return nil, err
	}

	if len(idList.GetDatas()) == 0 {
		return mtproto.MakeTLUsersSavedMusic(&mtproto.Users_SavedMusic{
			Count:     0,
			Documents: []*mtproto.Document{},
		}).To_Users_SavedMusic(), nil
	}

	// TODO: load documents by idList, offset, limit

	dList, err := c.svcCtx.Dao.MediaClient.MediaGetDocumentList(c.ctx, &media.TLMediaGetDocumentList{
		IdList: idList.GetDatas(),
	})
	if err != nil {
		c.Logger.Errorf("users.getSavedMusic - error: %v", err)
		return nil, err
	}

	return mtproto.MakeTLUsersSavedMusic(&mtproto.Users_SavedMusic{
		Count:     int32(len(dList.GetDatas())),
		Documents: dList.GetDatas(),
	}).To_Users_SavedMusic(), nil
}
