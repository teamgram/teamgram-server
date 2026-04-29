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

// DfsUploadThemeFile
// dfs.uploadThemeFile flags:# creator:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
func (c *DfsCore) DfsUploadThemeFile(in *dfs.TLDfsUploadThemeFile) (*tg.Document, error) {
	file, err := inputFile(in.File)
	if err != nil {
		return nil, err
	}
	uploaded, err := c.readUploadedDocumentData(in.Creator, file)
	if err != nil {
		return nil, err
	}
	repo := c.documents()
	documentID, err := repo.NextDocumentID(c.ctx)
	if err != nil {
		return nil, err
	}
	if err := c.uploadSessions().SaveObjectCacheRef(c.ctx, documentID, in.Creator, file.id); err != nil {
		return nil, err
	}
	var thumbs []tg.PhotoSizeClazz
	if in.Thumb != nil {
		thumbFile, err := inputFile(in.Thumb)
		if err != nil {
			return nil, err
		}
		thumb, err := c.readUploadedDocumentData(in.Creator, thumbFile)
		if err != nil {
			return nil, err
		}
		stored, err := repo.SaveDocumentThumbs(c.ctx, documentID, thumb.data, thumb.ext)
		if err != nil {
			return nil, err
		}
		thumbs = photoSizesFromStored(stored)
	}
	size, err := repo.SaveDocumentObject(c.ctx, documentID, uploaded.data)
	if err != nil {
		return nil, err
	}
	return makeDocumentWithThumbs(documentID, uploaded.ext, uploaded.date, in.MimeType, size, thumbs, nil, []tg.DocumentAttributeClazz{
		tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: in.FileName}),
	}), nil
}
