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

const gifNeedConvertSize = 10240

// DfsUploadGifDocumentMedia
// dfs.uploadGifDocumentMedia creator:long media:InputMedia = Document;
func (c *DfsCore) DfsUploadGifDocumentMedia(in *dfs.TLDfsUploadGifDocumentMedia) (*tg.Document, error) {
	media, err := inputMediaUploadedDocument(in.Media)
	if err != nil {
		return nil, err
	}
	file, err := inputFile(media.File)
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

	data := uploaded.data
	mimeType := "image/gif"
	var attrs []tg.DocumentAttributeClazz
	frame := data
	thumbExt := uploaded.ext
	if media.Thumb != nil {
		thumbFile, err := inputFile(media.Thumb)
		if err != nil {
			return nil, err
		}
		thumb, err := c.readUploadedDocumentData(in.Creator, thumbFile)
		if err != nil {
			return nil, err
		}
		frame = thumb.data
		thumbExt = thumb.ext
	}
	if uploaded.size >= gifNeedConvertSize {
		mp4Data, err := repo.ConvertDocumentToMP4(c.ctx, uploaded.data)
		if err != nil {
			return nil, err
		}
		data = mp4Data
		mimeType = "video/mp4"
		metadata, err := repo.GetDocumentVideoMetadata(c.ctx, mp4Data)
		if err != nil {
			return nil, err
		}
		attrs = videoAttributesFromMetadata(metadata, file.name+".mp4")
		if media.Thumb == nil {
			extracted, err := repo.ExtractDocumentFrame(c.ctx, data)
			if err != nil {
				return nil, err
			}
			frame = extracted
			thumbExt = ".jpg"
		}
	}
	thumbs, err := repo.SaveDocumentThumbs(c.ctx, documentID, frame, thumbExt)
	if err != nil {
		return nil, err
	}
	if attrs == nil {
		attrs = []tg.DocumentAttributeClazz{
			imageSizeAttributeFromThumbs(thumbs),
			tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: file.name}),
		}
	}
	size, err := repo.SaveDocumentObject(c.ctx, documentID, data)
	if err != nil {
		return nil, err
	}
	return makeDocumentWithThumbs(documentID, uploaded.ext, uploaded.date, mimeType, size, photoSizesFromStored(thumbs), nil, attrs), nil
}
