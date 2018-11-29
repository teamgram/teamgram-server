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

package server

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

// upload.getFile#e3a6cfb5 location:InputFileLocation offset:int limit:int = upload.File;
func (s *UploadServiceImpl) UploadGetFile(ctx context.Context, request *mtproto.TLUploadGetFile) (*mtproto.Upload_File, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("upload.getFile#e3a6cfb5 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))


	uploadFile, err := s.NbfsFacade.DownloadFile(request.GetLocation(), request.GetOffset(), request.GetLimit())
	if err != nil {
		glog.Error(err)
	}

	// type:storage.FileType mtime:int bytes:bytes
	glog.Infof("upload.getFile#e3a6cfb5 - reply: {type: %v, mime: %d, len_bytes: %d}",
		uploadFile.GetData2().GetType(),
		uploadFile.GetData2().GetMtime(),
		len(uploadFile.GetData2().GetBytes()))
	return uploadFile, nil
}
