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
	"path"

	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/internal/processor"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
)

// MediaProcessorProcessPhoto
// mediaProcessor.processPhoto owner_id:long object_id:string read_lease:bytes file_name:string profile:Bool = ProcessedPhoto;
func (c *MediaProcessorCore) MediaProcessorProcessPhoto(in *mediaprocessor.TLMediaProcessorProcessPhoto) (*mediaprocessor.ProcessedPhoto, error) {
	if in == nil || !validProcessInput(in.OwnerId, in.ObjectId, in.ReadLease, in.FileName) {
		return nil, mediaprocessor.ErrMediaProcessorInvalidArgument
	}
	original, err := c.readOriginalBytes(in.ReadLease)
	if err != nil {
		return nil, err
	}
	sizes, err := c.svcCtx.Processor.ResizePhoto(c.ctx, original, path.Ext(in.FileName), profileBool(in.Profile))
	if err != nil {
		return nil, err
	}

	out := mediaprocessor.MakeTLProcessedPhoto(&mediaprocessor.TLProcessedPhoto{
		OriginalObjectId: in.ObjectId,
	})
	for _, size := range sizes {
		fileName := size.Type + "_" + in.FileName
		stored, err := putDerivative(c, in.OwnerId, fileName, jpegMimeType, size.Bytes)
		if err != nil {
			return nil, err
		}
		out.Sizes = append(out.Sizes, makeDerivative(
			processor.DerivativePhotoSize,
			stored,
			fileName,
			jpegMimeType,
			int64(len(size.Bytes)),
			size.W,
			size.H,
			size.Bytes,
		))
	}
	return out.ToProcessedPhoto(), nil
}
