// Copyright 2024 Teamgram Authors
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
	"fmt"
	"math/rand"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/imaging"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"

	"github.com/minio/minio-go/v7"
)

// DfsUploadedProfilePhoto
// dfs.uploadedProfilePhoto creator:long photo_id:long = Photo;
func (c *DfsCore) DfsUploadedProfilePhoto(in *dfs.TLDfsUploadedProfilePhoto) (*mtproto.Photo, error) {
	var (
		fileSize minio.UploadInfo
		fileName = fmt.Sprintf("0/%d.jpeg", in.GetPhotoId())
	)

	cacheData, err := c.svcCtx.Dao.MinioUtil.GetPhotoFileData(c.ctx, fmt.Sprintf("0/%d.dat", in.GetPhotoId()))
	if err != nil {
		c.Logger.Errorf("dfs.uploadedProfilePhoto - %v", err)
		return nil, err
	}

	var (
		photoId    = c.svcCtx.Dao.IDGenClient2.NextId(c.ctx)
		ext        = model.GetFileExtName(fileName)
		sizeList   = make([]*mtproto.PhotoSize, 0, 3)
		extType    = model.GetStorageFileTypeConstructor(ext)
		accessHash = int64(extType)<<32 | int64(rand.Uint32())
	)

	err = imaging.ReSizeImage(cacheData, ext, true, func(szType string, localId int, w, h int32, b []byte) error {
		// secretId := int64(extType)<<32 | int64(rand.Uint32())
		path := fmt.Sprintf("%s/%d.dat", szType, photoId)
		// c.Logger.Debugf("path: %s", path)
		fileSize, err = c.svcCtx.Dao.PutPhotoFile(c.ctx, path, b)
		if err != nil {
			c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
			return err
		}

		sizeList = append(sizeList, mtproto.MakeTLPhotoSize(&mtproto.PhotoSize{
			Type:  szType,
			W:     w,
			H:     h,
			Size2: int32(fileSize.Size),
		}).To_PhotoSize())

		return nil
	})

	if err != nil {
		c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
		err = mtproto.ErrImageProcessFailed
		return nil, err
	} else if len(sizeList) == 0 {
		err = mtproto.ErrImageProcessFailed
		c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
		return nil, err
	}

	// photo#fb197a65 flags:#
	//	has_stickers:flags.0?true
	//	id:long
	//	access_hash:long
	//	file_reference:bytes
	//	date:int
	//	sizes:Vector<PhotoSize>
	//	video_sizes:flags.1?Vector<VideoSize>
	//	dc_id:int = Photo;
	photo := mtproto.MakeTLPhoto(&mtproto.Photo{
		HasStickers:   false,
		Id:            photoId,
		AccessHash:    accessHash,
		FileReference: []byte{},
		Date:          int32(time.Now().Unix()),
		Sizes:         sizeList,
		VideoSizes:    nil,
		DcId:          1,
	}).To_Photo()

	return photo, nil
}
