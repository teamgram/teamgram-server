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
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DfsUploadProfilePhotoFileV2
// dfs.uploadProfilePhotoFileV2 flags:# creator:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.4?VideoSize = Photo;
func (c *DfsCore) DfsUploadProfilePhotoFileV2(in *dfs.TLDfsUploadProfilePhotoFileV2) (*tg.Photo, error) {
	if in == nil || (in.File == nil && in.Video == nil) {
		return nil, dfs.ErrDfsInvalidArgument
	}
	if in.Video != nil {
		video, err := inputFile(in.Video)
		if err != nil {
			return nil, err
		}
		return c.buildVideoProfilePhoto(in.Creator, video, in.VideoStartTs)
	}
	file, err := inputFile(in.File)
	if err != nil {
		return nil, err
	}
	return c.buildPhotoFromUpload(in.Creator, file, true, nowUnix(), nil)
}

func (c *DfsCore) buildVideoProfilePhoto(creator int64, file *uploadedPhotoFile, videoStartTs *float64) (*tg.Photo, error) {
	if file == nil {
		return nil, dfs.ErrDfsInvalidArgument
	}
	if err := checkFileParts(file.parts); err != nil {
		return nil, err
	}
	reader, err := c.uploadSessions().OpenUploadedFile(c.ctx, creator, file.id)
	if err != nil {
		return nil, err
	}
	videoData, err := readAllSeeker(reader)
	if err != nil {
		return nil, dfs.WrapDfsStorage("read uploaded profile video", err)
	}
	if err := checkMD5(videoData, file.md5Checksum); err != nil {
		return nil, err
	}
	repo := c.photos()
	if repo == nil {
		return nil, dfs.WrapDfsStorage("profile video", errors.New("photo repository unavailable"))
	}
	mp4Data, err := repo.ConvertToMP4(c.ctx, videoData)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", dfs.ErrDfsVideoProcessFailed, err)
	}
	frameData, err := repo.ExtractFirstFrame(c.ctx, mp4Data)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", dfs.ErrDfsVideoProcessFailed, err)
	}
	w, h := int32(800), int32(800)
	videoSizes := []tg.VideoSizeClazz{
		tg.MakeTLVideoSize(&tg.TLVideoSize{
			Type:         "v",
			W:            w,
			H:            h,
			Size2:        int32(len(mp4Data)),
			VideoStartTs: videoStartTs,
		}),
	}
	photo, err := c.buildPhotoFromBytes(frameData, ".jpg", true, nowUnix(), videoSizes, false)
	if err != nil {
		return nil, err
	}
	if p, ok := photo.ToPhoto(); ok {
		if _, err := repo.SaveProfileVideoObject(c.ctx, p.Id, mp4Data); err != nil {
			return nil, err
		}
	}
	return photo, nil
}
