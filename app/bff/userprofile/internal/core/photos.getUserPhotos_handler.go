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
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	mediapb "github.com/teamgram/teamgram-server/app/service/media/media"
)

// PhotosGetUserPhotos
// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (c *UserProfileCore) PhotosGetUserPhotos(in *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error) {
	userId := mtproto.FromInputUser(c.MD.UserId, in.UserId)
	switch userId.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER:
	default:
		err := mtproto.ErrUserIdInvalid
		c.Logger.Errorf("photos.getUserPhotos - error: %v", err)
		return nil, err
	}

	cachePhotos, err := c.svcCtx.Dao.UserClient.UserGetProfilePhotos(c.ctx, &userpb.TLUserGetProfilePhotos{
		UserId: userId.PeerId,
	})
	if err != nil {
		c.Logger.Errorf("photos.getUserPhotos - error: %v", err)
		return nil, err
	}

	photos := mtproto.MakeTLPhotosPhotos(&mtproto.Photos_Photos{
		Photos: make([]*mtproto.Photo, 0, len(cachePhotos.GetDatas())),
		Users:  []*mtproto.User{},
	}).To_Photos_Photos()

	for _, id := range cachePhotos.GetDatas() {
		if photo, err := c.svcCtx.Dao.MediaClient.MediaGetPhoto(c.ctx,
			&mediapb.TLMediaGetPhoto{
				PhotoId: id,
			}); err != nil {
			c.Logger.Errorf("photos.getUserPhotos - error: %v", err)
		} else if photo != nil {
			photos.Photos = append(photos.Photos, photo)
		}
	}

	return photos, nil
}
