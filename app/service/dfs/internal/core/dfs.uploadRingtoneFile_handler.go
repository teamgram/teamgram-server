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

// DfsUploadRingtoneFile
// dfs.uploadRingtoneFile creator:long file:InputFile mime_type:string file_name:string = Document;
func (c *DfsCore) DfsUploadRingtoneFile(in *dfs.TLDfsUploadRingtoneFile) (*tg.Document, error) {
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
	metadata, err := repo.GetDocumentVideoMetadata(c.ctx, uploaded.data)
	if err != nil {
		return nil, err
	}
	var duration int32
	if metadata != nil {
		duration = metadata.Duration
	}
	size, err := repo.SaveDocumentObject(c.ctx, documentID, uploaded.data)
	if err != nil {
		return nil, err
	}
	title := in.FileName
	performer := ""
	return makeDocument(documentID, uploaded.ext, uploaded.date, in.MimeType, size, []tg.DocumentAttributeClazz{
		tg.MakeTLDocumentAttributeAudio(&tg.TLDocumentAttributeAudio{
			Duration:  duration,
			Title:     &title,
			Performer: &performer,
		}),
		tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: in.FileName}),
	}), nil
}
