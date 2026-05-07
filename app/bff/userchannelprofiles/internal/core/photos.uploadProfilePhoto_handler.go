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
	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// PhotosUploadProfilePhoto
// photos.uploadProfilePhoto#388a3b5 flags:# fallback:flags.3?true bot:flags.5?InputUser file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.4?VideoSize = photos.Photo;
func (c *UserChannelProfilesCore) PhotosUploadProfilePhoto(in *tg.TLPhotosUploadProfilePhoto) (*tg.PhotosPhoto, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil || (in.File == nil && in.Video == nil) {
		return nil, tg.ErrInputRequestInvalid
	}
	if err := requireUserClient(c); err != nil {
		return nil, err
	}
	if err := requireMediaClient(c); err != nil {
		return nil, err
	}

	photo, err := c.svcCtx.Repo.MediaClient.MediaUploadProfilePhotoFile(c.ctx, &mediapb.TLMediaUploadProfilePhotoFile{
		OwnerId:          c.MD.PermAuthKeyId,
		File:             in.File,
		Video:            in.Video,
		VideoStartTs:     in.VideoStartTs,
		VideoEmojiMarkup: in.VideoEmojiMarkup,
	})
	if err != nil {
		return nil, err
	}
	if photo == nil || photo.Clazz == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	uploaded, ok := photo.Clazz.(*tg.TLPhoto)
	if !ok || uploaded.Id <= 0 {
		return nil, tg.ErrInputRequestInvalid
	}

	if _, err = c.svcCtx.Repo.UserClient.UserUpdateProfilePhoto(c.ctx, &userpb.TLUserUpdateProfilePhoto{
		UserId: selfID,
		Id:     uploaded.Id,
	}); err != nil {
		return nil, err
	}
	// TODO(v2 userchannelprofiles): sync delivery is intentionally not migrated here; route profile photo updates through userupdates/gateway when the V2 delivery contract is defined.
	users, err := userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, selfID, []int64{selfID}, userprojection.MissingExplicitInput)
	if err != nil {
		return nil, err
	}

	return tg.MakeTLPhotosPhoto(&tg.TLPhotosPhoto{
		Photo: photo.Clazz,
		Users: users,
	}).ToPhotosPhoto(), nil
}
