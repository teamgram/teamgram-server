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
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
)

// MediaProcessorProcessGif
// mediaProcessor.processGif owner_id:long object_id:string read_lease:bytes file_name:string thumb_object_id:string thumb_read_lease:bytes = ProcessedDocument;
func (c *MediaProcessorCore) MediaProcessorProcessGif(in *mediaprocessor.TLMediaProcessorProcessGif) (*mediaprocessor.ProcessedDocument, error) {
	if in == nil || !validProcessInput(in.OwnerId, in.ObjectId, in.ReadLease, in.FileName) {
		return nil, mediaprocessor.ErrMediaProcessorInvalidArgument
	}
	original, err := c.readOriginalBytes(in.ReadLease)
	if err != nil {
		return nil, err
	}
	if len(original) < minGifBytes {
		return nil, mediaprocessor.ErrMediaProcessorInvalidArgument
	}

	mp4Bytes, err := c.svcCtx.Processor.ConvertGIFToMP4(c.ctx, original)
	if err != nil {
		return nil, err
	}
	metadata, err := c.svcCtx.Processor.ProbeMP4(c.ctx, mp4Bytes)
	if err != nil {
		return nil, err
	}

	mp4Name := mp4FileName(in.FileName)
	stored, err := putDerivative(c, in.OwnerId, mp4Name, mp4MimeType, mp4Bytes)
	if err != nil {
		return nil, err
	}
	size := int64(len(mp4Bytes))
	objectID := ""
	if stored != nil {
		objectID = stored.ObjectId
		if stored.Size2 > 0 {
			size = stored.Size2
		}
	}
	if objectID == "" {
		return nil, mediaprocessor.ErrMediaProcessorInvalidArgument
	}

	attributes, err := encodeVideoAttributes(metadata, mp4Name, true)
	if err != nil {
		return nil, err
	}
	out := mediaprocessor.MakeTLProcessedDocument(&mediaprocessor.TLProcessedDocument{
		ObjectId:   objectID,
		MimeType:   mp4MimeType,
		Size2:      size,
		Attributes: attributes,
	})

	if len(in.ThumbReadLease) != 0 {
		if in.ThumbObjectId == "" {
			return nil, mediaprocessor.ErrMediaProcessorInvalidArgument
		}
		thumb, err := c.readOriginalBytes(in.ThumbReadLease)
		if err != nil {
			return nil, err
		}
		thumbName := thumbFileName(in.FileName)
		thumbStored, err := putDerivative(c, in.OwnerId, thumbName, jpegMimeType, thumb)
		if err != nil {
			return nil, err
		}
		out.Thumbs = append(out.Thumbs, makeThumbDerivative(thumbStored, thumbName, metadata, thumb))
		return out.ToProcessedDocument(), nil
	}

	cover, err := c.svcCtx.Processor.ExtractMP4Cover(c.ctx, mp4Bytes)
	if err != nil {
		c.Logger.Errorf("mediaProcessor.processGif - cover extraction failed: object_id: %s, file_name: %s, err: %v", in.ObjectId, in.FileName, err)
		return out.ToProcessedDocument(), nil
	}
	thumbName := thumbFileName(mp4Name)
	thumbStored, err := putDerivative(c, in.OwnerId, thumbName, jpegMimeType, cover)
	if err != nil {
		return nil, err
	}
	out.Thumbs = append(out.Thumbs, makeThumbDerivative(thumbStored, thumbName, metadata, cover))
	return out.ToProcessedDocument(), nil
}
