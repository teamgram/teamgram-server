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
)

// PhotosDeletePhotos
// photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
func (c *UserProfileCore) PhotosDeletePhotos(in *mtproto.TLPhotosDeletePhotos) (*mtproto.Vector_Long, error) {
	var (
		photo        *mtproto.Photo
		deleteIdList = make([]int64, 0, len(in.GetId()))
	)

	for _, id := range in.GetId() {
		deleteIdList = append(deleteIdList, id.GetId())
	}

	// TODO: ALBUM_PHOTOS_TOO_MANY
	mainId, err := c.svcCtx.Dao.UserClient.UserDeleteProfilePhotos(c.ctx, &userpb.TLUserDeleteProfilePhotos{
		UserId: c.MD.UserId,
		Id:     deleteIdList,
	})
	if err != nil {
		c.Logger.Errorf("photos.deletePhotos - error: %v", err)
		return nil, err
	}

	if mainId.V != 0 {
		photo, err = c.svcCtx.Dao.MediaClient.MediaGetPhoto(c.ctx, &mediapb.TLMediaGetPhoto{
			PhotoId: mainId.V,
		})
		if err != nil {
			c.Logger.Errorf("photos.deletePhotos - error: %v", err)
			return nil, err
		}
	}

	if photo == nil {
		photo = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
	}

	me, err := c.svcCtx.Dao.UserClient.UserGetImmutableUser(
		c.ctx,
		&userpb.TLUserGetImmutableUser{
			Id:       c.MD.UserId,
			Privacy:  false,
			Contacts: nil,
		})
	if err != nil {
		c.Logger.Errorf("photos.updateProfilePhoto - error: %v", err)
		return nil, err
	}

	c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
		UserId: c.MD.UserId,
		Updates: mtproto.MakeUpdatesByUpdatesUsers(
			[]*mtproto.User{me.ToSelfUser()},
			mtproto.MakeTLUpdateUser(&mtproto.Update{
				UserId: c.MD.UserId,
			}).To_Update()),
	})

	return &mtproto.Vector_Long{
		Datas: deleteIdList,
	}, nil
}
