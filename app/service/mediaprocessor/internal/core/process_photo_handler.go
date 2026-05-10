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
	"fmt"
	"path"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
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
		return nil, fmt.Errorf("read original photo bytes: %w", err)
	}
	derivatives, err := c.svcCtx.Processor.BuildPhotoDerivatives(c.ctx, original, path.Ext(in.FileName), profileBool(in.Profile))
	if err != nil {
		return nil, fmt.Errorf("build photo derivatives: %w", err)
	}

	out := mediaprocessor.MakeTLProcessedPhoto(&mediaprocessor.TLProcessedPhoto{
		OriginalObjectId: in.ObjectId,
	})
	for _, derivative := range derivatives {
		fileName := derivative.Type + "_" + in.FileName
		kind := processor.DerivativePhotoSize
		var stored *dfs.FileFinalizedObject
		if derivative.Stripped {
			kind = processor.DerivativePhotoStripped
		} else {
			stored, err = putDerivative(c, in.OwnerId, fileName, jpegMimeType, derivative.Bytes)
			if err != nil {
				return nil, fmt.Errorf("put photo derivative %s: %w", derivative.Type, err)
			}
		}
		out.Sizes = append(out.Sizes, makeDerivative(
			kind,
			stored,
			fileName,
			jpegMimeType,
			int64(len(derivative.Bytes)),
			derivative.W,
			derivative.H,
			derivative.Bytes,
			derivative.ProgressiveSizes,
		))
	}
	return out.ToProcessedPhoto(), nil
}
