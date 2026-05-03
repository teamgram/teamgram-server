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

// PhotosUpdateProfilePhoto
// photos.updateProfilePhoto#9e82039 flags:# fallback:flags.0?true bot:flags.1?InputUser id:InputPhoto = photos.Photo;
func (c *UserChannelProfilesCore) PhotosUpdateProfilePhoto(in *tg.TLPhotosUpdateProfilePhoto) (*tg.PhotosPhoto, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	photoID, err := photoIDFromInputPhoto(in.Id)
	if err != nil {
		return nil, err
	}
	if err := requireUserClient(c); err != nil {
		return nil, err
	}
	if err := requireMediaClient(c); err != nil {
		return nil, err
	}

	updatedPhotoID, err := c.svcCtx.Repo.UserClient.UserUpdateProfilePhoto(c.ctx, &userpb.TLUserUpdateProfilePhoto{
		UserId: selfID,
		Id:     photoID,
	})
	if err != nil {
		return nil, err
	}

	var photo tg.PhotoClazz
	if updatedPhotoID != nil && updatedPhotoID.V > 0 {
		gotPhoto, err := c.svcCtx.Repo.MediaClient.MediaGetPhoto(c.ctx, &mediapb.TLMediaGetPhoto{
			PhotoId: updatedPhotoID.V,
		})
		if err != nil {
			return nil, err
		}
		if gotPhoto != nil {
			photo = gotPhoto.Clazz
		}
	}
	if photo == nil {
		photo = tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{Id: photoID})
	}
	// TODO(v2 userchannelprofiles): sync delivery is intentionally not migrated here; route profile photo updates through userupdates/gateway when the V2 delivery contract is defined.

	return tg.MakeTLPhotosPhoto(&tg.TLPhotosPhoto{
		Photo: photo,
		Users: []tg.UserClazz{},
	}).ToPhotosPhoto(), nil
}
