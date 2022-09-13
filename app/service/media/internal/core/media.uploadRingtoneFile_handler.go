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
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// MediaUploadRingtoneFile
// media.uploadRingtoneFile flags:# owner_id:long file:InputFile mime_type:string file_name:string = Document;
func (c *MediaCore) MediaUploadRingtoneFile(in *media.TLMediaUploadRingtoneFile) (*mtproto.Document, error) {
	var (
		err      error
		document *mtproto.Document
		file     = in.GetFile()
	)

	if file == nil {
		c.Logger.Errorf("media.uploadRingtoneFile - error: file is nil")
		return nil, mtproto.ErrInputRequestInvalid
	}

	document, err = c.svcCtx.Dao.DfsClient.DfsUploadRingtoneFile(c.ctx, &dfs.TLDfsUploadRingtoneFile{
		Creator:  in.OwnerId,
		File:     file,
		FileName: in.GetFileName(),
		MimeType: in.GetMimeType(),
	})
	if err != nil {
		c.Logger.Errorf("media.uploadRingtoneFile - error: %v", err)
		// err = mtproto.ErrThemeFileInvalid
		return nil, err
	}

	if len(document.GetThumbs()) > 0 {
		c.svcCtx.Dao.SavePhotoSizeV2(c.ctx, document.GetId(), document.GetThumbs())
	}
	c.svcCtx.Dao.SaveDocumentV2(c.ctx, file.GetName(), document)

	return document, nil
}
