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

package photos

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/service/document/client"
	"golang.org/x/net/context"
	"time"
)

// photos.uploadProfilePhoto#4f32c098
// Updates current user profile photo.
// file: File saved in parts by means of upload.saveFilePart method
//
// photos.uploadProfilePhoto#4f32c098 file:InputFile = photos.Photo;
func (s *PhotosServiceImpl) PhotosUploadProfilePhoto(ctx context.Context, request *mtproto.TLPhotosUploadProfilePhoto) (*mtproto.Photos_Photo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("photos.uploadProfilePhoto#4f32c098 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	file := request.GetFile()
	// uuid := helper.NextSnowflakeId()

	result, err := document_client.UploadPhotoFile(md.AuthId, file) // uuid, file.GetData2().GetId(), file.GetData2().GetParts(), file.GetData2().GetName(), file.GetData2().GetMd5Checksum())
	if err != nil {
		glog.Errorf("UploadPhoto error: %v", err)
		return nil, err
	}

	s.UserModel.SetUserPhotoID(md.UserId, result.PhotoId)

	// TODO(@benqi): sync update userProfilePhoto

	// fileData := mediaData.GetFile().GetData2()
	photo := &mtproto.TLPhoto{Data2: &mtproto.Photo_Data{
		Id:          result.PhotoId,
		HasStickers: false,
		AccessHash:  result.AccessHash, //photo2.GetFileAccessHash(file.GetData2().GetId(), file.GetData2().GetParts()),
		Date:        int32(time.Now().Unix()),
		Sizes:       result.SizeList,
	}}

	photos := &mtproto.TLPhotosPhoto{Data2: &mtproto.Photos_Photo_Data{
		Photo: photo.To_Photo(),
		Users: []*mtproto.User{},
	}}

	//updateUserPhoto := &mtproto.TLUpdateUserPhoto{Data2: &mtproto.Update_Data{
	//	UserId:   md.UserId,
	//	Date:     int32(time.Now().Unix()),
	//	Photo:    photo2.MakeUserProfilePhoto(result.PhotoId, result.SizeList),
	//	Previous: mtproto.ToBool(false),
	//}}
	// sync_client.GetSyncClient().PushToUserUpdateShortData(md.UserId, updateUserPhoto.To_Update())

	glog.Infof("photos.uploadProfilePhoto#4f32c098 - reply: %s", logger.JsonDebugData(photos))
	return photos.To_Photos_Photo(), nil
}
