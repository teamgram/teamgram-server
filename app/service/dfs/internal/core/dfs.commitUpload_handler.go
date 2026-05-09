// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
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
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DfsCommitUpload
// dfs.commitUpload upload_session_id:string owner_id:long file:InputFile purpose:string = FileFinalizedObject;
func (c *DfsCore) DfsCommitUpload(in *dfs.TLDfsCommitUpload) (*dfs.FileFinalizedObject, error) {
	if in == nil || in.UploadSessionId == "" || in.OwnerId == 0 || in.File == nil || in.Purpose == "" {
		return nil, dfs.ErrDfsInvalidArgument
	}
	out, err := c.fileObjects().CommitUpload(c.ctx, in.UploadSessionId, in.OwnerId, in.File, in.Purpose)
	if err != nil {
		var missing *dfs.MissingUploadPartError
		if errors.As(err, &missing) {
			return nil, tg.NewFilePartXMissing(missing.Part)
		}
		return nil, err
	}
	return out, nil
}
