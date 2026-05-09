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

// MediaProcessorProcessMp4
// mediaProcessor.processMp4 owner_id:long object_id:string read_lease:bytes file_name:string attributes:bytes = ProcessedDocument;
func (c *MediaProcessorCore) MediaProcessorProcessMp4(in *mediaprocessor.TLMediaProcessorProcessMp4) (*mediaprocessor.ProcessedDocument, error) {
	if in == nil || !validProcessInput(in.OwnerId, in.ObjectId, in.ReadLease, in.FileName) {
		return nil, mediaprocessor.ErrMediaProcessorInvalidArgument
	}
	original, err := c.readOriginalBytes(in.ReadLease)
	if err != nil {
		return nil, err
	}
	metadata, err := c.svcCtx.Processor.ProbeMP4(c.ctx, original)
	if err != nil {
		return nil, err
	}
	attributes, err := encodeVideoAttributes(metadata, in.FileName, false)
	if err != nil {
		return nil, err
	}

	out := mediaprocessor.MakeTLProcessedDocument(&mediaprocessor.TLProcessedDocument{
		ObjectId:   in.ObjectId,
		MimeType:   mp4MimeType,
		Size2:      int64(len(original)),
		Attributes: attributes,
	})
	cover, err := c.svcCtx.Processor.ExtractMP4Cover(c.ctx, original)
	if err != nil {
		c.Logger.Errorf("mediaProcessor.processMp4 - cover extraction failed: object_id: %s, file_name: %s, err: %v", in.ObjectId, in.FileName, err)
		return out.ToProcessedDocument(), nil
	}
	thumbName := thumbFileName(in.FileName)
	thumbStored, err := putDerivative(c, in.OwnerId, thumbName, jpegMimeType, cover)
	if err != nil {
		return nil, err
	}
	out.Thumbs = append(out.Thumbs, makeThumbDerivative(thumbStored, thumbName, metadata, cover))
	return out.ToProcessedDocument(), nil
}
