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

// PhotosGetUserPhotos
// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (c *UserChannelProfilesCore) PhotosGetUserPhotos(in *tg.TLPhotosGetUserPhotos) (*tg.PhotosPhotos, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	targetID, err := userIDFromInputUser(selfID, in.UserId)
	if err != nil {
		return nil, err
	}
	if err := requireUserClient(c); err != nil {
		return nil, err
	}
	if err := requireMediaClient(c); err != nil {
		return nil, err
	}

	idList, err := c.svcCtx.Repo.UserClient.UserGetProfilePhotos(c.ctx, &userpb.TLUserGetProfilePhotos{
		UserId: targetID,
	})
	if err != nil {
		return nil, err
	}

	photos := []tg.PhotoClazz{}
	if idList != nil {
		for _, id := range idList.Datas {
			photo, err := c.svcCtx.Repo.MediaClient.MediaGetPhoto(c.ctx, &mediapb.TLMediaGetPhoto{
				PhotoId: id,
			})
			if err != nil || photo == nil || photo.Clazz == nil {
				continue
			}
			photos = append(photos, photo.Clazz)
		}
	}

	return tg.MakeTLPhotosPhotos(&tg.TLPhotosPhotos{
		Photos: photos,
		Users:  []tg.UserClazz{},
	}).ToPhotosPhotos(), nil
}
