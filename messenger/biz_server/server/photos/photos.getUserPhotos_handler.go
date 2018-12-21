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

/*
 rpc_requst:
	body: { photos_getUserPhotos
	  user_id: { inputUserSelf },
	  offset: 1 [INT],
	  max_id: 0 [LONG],
	  limit: 5 [INT],
	},

 rpc_result:
  body: { rpc_result
    req_msg_id: 6537205080566771468 [LONG],
    result: { photos_photosSlice
      count: 1 [INT],
      photos: [ vector<0x0> ],
      users: [ vector<0x0> ],
    },
  },
*/

// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (s *PhotosServiceImpl) PhotosGetUserPhotos(ctx context.Context, request *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("photos.getUserPhotos#91cd32a8 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl PhotosGetUserPhotos logic
	var userId int32 = 0
	switch request.GetUserId().GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputUserSelf:
		userId = md.UserId
	case mtproto.TLConstructor_CRC32_inputUser:
		userId = request.GetUserId().GetData2().GetUserId()
	default:
		// TODO(@benqi): bad request
	}

	photos := mtproto.NewTLPhotosPhotos()
	photoIdList := s.UserModel.GetUserPhotoIDList(userId)
	// idList := []int32{}
	for _, photoId := range photoIdList {
		sizes, _ := document_client.GetPhotoSizeList(photoId)
		// photo2 := photo2.MakeUserProfilePhoto(photoId, sizes)
		photo := &mtproto.TLPhotoLayer86{Data2: &mtproto.Photo_Data{
			Id:          photoId,
			HasStickers: false,
			AccessHash:  photoId, // photo2.GetFileAccessHash(file.GetData2().GetId(), file.GetData2().GetParts()),
			Date:        int32(time.Now().Unix()),
			Sizes:       sizes,
		}}
		photos.Data2.Photos = append(photos.Data2.Photos, photo.To_Photo())
	}
	// if idList

	glog.Infof("photos.getUserPhotos#91cd32a8 - reply: %s", logger.JsonDebugData(photos))
	return photos.To_Photos_Photos(), nil
}
