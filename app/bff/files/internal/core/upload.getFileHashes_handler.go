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
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UploadGetFileHashes
// upload.getFileHashes#9156982a location:InputFileLocation offset:long = Vector<FileHash>;
func (c *FilesCore) UploadGetFileHashes(in *tg.TLUploadGetFileHashes) (*tg.VectorFileHash, error) {
	if in == nil || in.Location == nil {
		return nil, tg.ErrLocationInvalid
	}

	switch in.Location.InputFileLocationClazzName() {
	case tg.ClazzName_inputDocumentFileLocation,
		tg.ClazzName_inputPhotoFileLocation,
		tg.ClazzName_inputPeerPhotoFileLocation:
		resolved, err := c.resolveFileLocation(in.Location)
		if err != nil {
			return nil, err
		}
		hashes, err := c.svcCtx.Repo.DfsClient.DfsGetFileHashesByReadLease(c.ctx, &dfs.TLDfsGetFileHashesByReadLease{
			ReadLease: resolved.ReadLease,
			Offset:    in.Offset,
			Limit:     0,
		})
		if err != nil {
			return nil, err
		}
		if hashes == nil {
			return nil, tg.ErrInputRequestInvalid
		}
		return &tg.VectorFileHash{Datas: hashes.Datas}, nil
	default:
		return nil, tg.ErrLocationInvalid
	}
}
