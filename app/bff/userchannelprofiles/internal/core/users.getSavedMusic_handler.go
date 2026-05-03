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
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UsersGetSavedMusic
// users.getSavedMusic#788d7fe3 id:InputUser offset:int limit:int hash:long = users.SavedMusic;
func (c *UserChannelProfilesCore) UsersGetSavedMusic(in *tg.TLUsersGetSavedMusic) (*tg.UsersSavedMusic, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	targetID, err := userIDFromInputUser(selfID, in.Id)
	if err != nil {
		return nil, err
	}
	if err := requireUserClient(c); err != nil {
		return nil, err
	}

	idList, err := c.svcCtx.Repo.UserClient.UserGetSavedMusicIdList(c.ctx, &userpb.TLUserGetSavedMusicIdList{
		UserId: targetID,
	})
	if err != nil {
		return nil, err
	}
	if idList == nil || len(idList.Datas) == 0 {
		return tg.MakeTLUsersSavedMusic(&tg.TLUsersSavedMusic{
			Count:     0,
			Documents: []tg.DocumentClazz{},
		}).ToUsersSavedMusic(), nil
	}
	if err := requireMediaClient(c); err != nil {
		return nil, err
	}

	documents, err := c.svcCtx.Repo.MediaClient.MediaGetDocumentList(c.ctx, &mediapb.TLMediaGetDocumentList{
		IdList: idList.Datas,
	})
	if err != nil {
		return nil, err
	}
	datas := []tg.DocumentClazz{}
	if documents != nil {
		datas = documents.Datas
	}
	return tg.MakeTLUsersSavedMusic(&tg.TLUsersSavedMusic{
		Count:     int32(len(datas)),
		Documents: datas,
	}).ToUsersSavedMusic(), nil
}
