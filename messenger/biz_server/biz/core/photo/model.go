// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package photo

import (
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/service/document/client"
	"time"
)

type PhotoModel struct {
	// dao *stickersDAO
}

func (m *PhotoModel) InstallModel() {
	// m.dao.StickerPacksDAO = dao.GetStickerPacksDAO(dao.DB_MASTER)
}

func (m *PhotoModel) RegisterCallback(cb interface{}) {
}

func (m *PhotoModel) GetUserProfilePhoto(photoId int64) (photo *mtproto.UserProfilePhoto) {
	if photoId == 0 {
		photo = mtproto.NewTLUserProfilePhotoEmpty().To_UserProfilePhoto()
	} else {
		sizes, _ := document_client.GetPhotoSizeList(photoId)
		photo = MakeUserProfilePhoto(photoId, sizes)
	}

	return
}

func (m *PhotoModel) GetChatPhoto(photoId int64) (photo *mtproto.ChatPhoto) {
	if photoId == 0 {
		photo = mtproto.NewTLChatPhotoEmpty().To_ChatPhoto()
	} else {
		sizes, _ := document_client.GetPhotoSizeList(photoId)
		photo = MakeChatPhoto(sizes)
	}

	return
}

func (m *PhotoModel) GetPhoto(photoId int64) (photo *mtproto.Photo) {
	if photoId == 0 {
		photo = mtproto.NewTLPhotoEmpty().To_Photo()
	} else {
		sizes, _ := document_client.GetPhotoSizeList(photoId)
		if len(sizes) == 0 {
			photo = mtproto.NewTLPhotoEmpty().To_Photo()
		} else {
			photo2 := &mtproto.TLPhoto{Data2: &mtproto.Photo_Data{
				Id:          photoId,
				HasStickers: false,
				AccessHash:  photoId, // photo2.GetFileAccessHash(file.GetData2().GetId(), file.GetData2().GetParts()),
				Date:        int32(time.Now().Unix()),
				Sizes:       sizes,
			}}
			photo = photo2.To_Photo()
		}
	}
	return
}

func MakeUserProfilePhoto(photoId int64, sizes []*mtproto.PhotoSize) *mtproto.UserProfilePhoto {
	if len(sizes) == 0 {
		return mtproto.NewTLUserProfilePhotoEmpty().To_UserProfilePhoto()
	}

	// TODO(@benqi): check PhotoSize is photoSizeEmpty
	photo := &mtproto.TLUserProfilePhoto{Data2: &mtproto.UserProfilePhoto_Data{
		PhotoId:    photoId,
		PhotoSmall: sizes[0].GetData2().GetLocation(),
		PhotoBig:   sizes[len(sizes)-1].GetData2().GetLocation(),
	}}

	return photo.To_UserProfilePhoto()
}

func MakeChatPhoto(sizes []*mtproto.PhotoSize) *mtproto.ChatPhoto {
	if len(sizes) == 0 {
		return mtproto.NewTLChatPhotoEmpty().To_ChatPhoto()
	}

	// TODO(@benqi): check PhotoSize is photoSizeEmpty
	photo := &mtproto.TLChatPhoto{Data2: &mtproto.ChatPhoto_Data{
		PhotoSmall: sizes[0].GetData2().GetLocation(),
		PhotoBig:   sizes[len(sizes)-1].GetData2().GetLocation(),
	}}

	return photo.To_ChatPhoto()
}

func init() {
	core.RegisterCoreModel(&PhotoModel{})
}
