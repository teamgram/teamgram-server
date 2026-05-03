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
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UsersGetSavedMusicByID
// users.getSavedMusicByID#7573a4e9 id:InputUser documents:Vector<InputDocument> = users.SavedMusic;
func (c *UserChannelProfilesCore) UsersGetSavedMusicByID(in *tg.TLUsersGetSavedMusicByID) (*tg.UsersSavedMusic, error) {
	if _, err := requireSelfID(c); err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if err := requireMediaClient(c); err != nil {
		return nil, err
	}

	ids := make([]int64, 0, len(in.Documents))
	for _, inputDocument := range in.Documents {
		id, err := documentIDFromInputDocument(inputDocument)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	documents, err := c.svcCtx.Repo.MediaClient.MediaGetDocumentList(c.ctx, &mediapb.TLMediaGetDocumentList{
		IdList: ids,
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
