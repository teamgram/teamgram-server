// Copyright 2024 Teamgram Authors
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
	"fmt"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// MediaUploadedProfilePhoto
// media.uploadedProfilePhoto owner_id:long photo_id:long = Photo;
func (c *MediaCore) MediaUploadedProfilePhoto(in *media.TLMediaUploadedProfilePhoto) (*mtproto.Photo, error) {
	photo, err := c.svcCtx.Dao.DfsClient.DfsUploadedProfilePhoto(c.ctx, &dfs.TLDfsUploadedProfilePhoto{
		Creator: in.OwnerId,
		PhotoId: in.PhotoId,
	})
	if err != nil {
		c.Logger.Error("media.uploadedProfilePhoto - error: %v", err.Error())
		return nil, err
	}

	if err = c.svcCtx.Dao.SavePhotoSizeV2(c.ctx, photo.GetId(), photo.GetSizes()); err != nil {
		c.Logger.Error("media.uploadedProfilePhoto - error: %v", err.Error())
		return nil, err
	}

	_ = c.svcCtx.SavePhotoV2(c.ctx,
		photo.GetId(),
		photo.GetAccessHash(),
		photo.GetHasStickers(),
		false,
		fmt.Sprintf("0/%d.jpeg", in.GetPhotoId()))

	return photo, nil
}
