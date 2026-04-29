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

// DfsUploadDocumentFileV2
// dfs.uploadDocumentFileV2 creator:long media:InputMedia = Document;
func (c *DfsCore) DfsUploadDocumentFileV2(in *dfs.TLDfsUploadDocumentFileV2) (*tg.Document, error) {
	if in == nil || in.Media == nil {
		return nil, dfs.ErrDfsInvalidArgument
	}
	media, err := inputMediaUploadedDocument(in.Media)
	if err != nil {
		return nil, err
	}
	file, err := inputFile(media.File)
	if err != nil {
		return nil, err
	}
	if err := checkFileParts(file.parts); err != nil {
		return nil, err
	}
	reader, err := c.uploadSessions().OpenUploadedFile(c.ctx, in.Creator, file.id)
	if err != nil {
		return nil, err
	}
	data, err := readAllSeeker(reader)
	if err != nil {
		return nil, dfs.WrapDfsStorage("read uploaded document", err)
	}
	if err := checkMD5(data, file.md5Checksum); err != nil {
		return nil, err
	}
	info, err := c.uploadSessions().LoadUploadedFileInfo(c.ctx, in.Creator, file.id)
	if err != nil {
		return nil, err
	}
	repo := c.documents()
	if repo == nil {
		return nil, dfs.WrapDfsStorage("upload document", dfs.ErrDfsStorage)
	}
	documentID, err := repo.NextDocumentID(c.ctx)
	if err != nil {
		return nil, err
	}
	if err := c.uploadSessions().SaveObjectCacheRef(c.ctx, documentID, in.Creator, file.id); err != nil {
		return nil, err
	}
	if _, err := repo.SaveDocumentObject(c.ctx, documentID, data); err != nil {
		return nil, err
	}
	date := int32(info.Mtime)
	if date == 0 {
		date = nowUnix()
	}
	size := info.FileSize()
	if size == 0 {
		size = int64(len(data))
	}
	return makeDocument(documentID, fileExt(file.name), date, media.MimeType, size, filterDocumentAttributes(media.MimeType, media.Attributes)), nil
}
