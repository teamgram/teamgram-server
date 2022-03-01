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
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	mediapb "github.com/teamgram/teamgram-server/app/service/media/media"
	"time"
)

// PhotosUpdateProfilePhoto
// photos.updateProfilePhoto#72d4742c id:InputPhoto = photos.Photo;
func (c *PhotosCore) PhotosUpdateProfilePhoto(in *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.Photos_Photo, error) {
	var (
		photo *mtproto.Photo
	)

	// TODO: ALBUM_PHOTOS_TOO_MANY
	updatedPhotoId, err := c.svcCtx.Dao.UserClient.UserUpdateProfilePhoto(c.ctx, &userpb.TLUserUpdateProfilePhoto{
		UserId: c.MD.UserId,
		Id:     in.GetId().GetId(),
	})
	if err != nil {
		c.Logger.Errorf("photos.uploadProfilePhoto - error: %v", err)
		return nil, err
	}

	if updatedPhotoId.V != 0 {
		photo, err = c.svcCtx.Dao.MediaClient.MediaGetPhoto(c.ctx, &mediapb.TLMediaGetPhoto{
			PhotoId: updatedPhotoId.V,
		})
		if err != nil {
			c.Logger.Errorf("photos.uploadProfilePhoto - error: %v", err)
			return nil, err
		}
	}

	if photo == nil {
		photo = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
	}

	c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
		UserId: c.MD.UserId,
		Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateUserPhoto(&mtproto.Update{
			UserId:   c.MD.UserId,
			Date:     int32(time.Now().Unix()),
			Photo:    mtproto.MakeUserProfilePhotoByPhoto(photo),
			Previous: mtproto.BoolFalse,
		}).To_Update()),
	})

	return mtproto.MakeTLPhotosPhoto(&mtproto.Photos_Photo{
		Photo: photo,
		Users: []*mtproto.User{},
	}).To_Photos_Photo(), nil
}
