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
	photo2 "github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/photo"
	"github.com/nebula-chat/chatengine/service/document/client"
	"golang.org/x/net/context"
)

// photos.updateProfilePhoto#f0bb5152 id:InputPhoto = UserProfilePhoto;
func (s *PhotosServiceImpl) PhotosUpdateProfilePhoto(ctx context.Context, request *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.UserProfilePhoto, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("photos.updateProfilePhoto#f0bb5152 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		photo *mtproto.UserProfilePhoto
	)

	if request.GetId().GetConstructor() == mtproto.TLConstructor_CRC32_inputPhotoEmpty {
		photo = mtproto.NewTLUserProfilePhotoEmpty().To_UserProfilePhoto()
	} else {
		id := request.GetId().To_InputPhoto()
		// TODO(@benqi): check inputPhoto.access_hash

		sizes, _ := document_client.GetPhotoSizeList(id.GetId())
		photo = photo2.MakeUserProfilePhoto(id.GetId(), sizes)
	}

	// TODO(@benqi): sync update.

	glog.Infof("photos.uploadProfilePhoto#4f32c098 - reply: %s", logger.JsonDebugData(photo))
	return photo, nil
}
