// Copyright 2022 Teamgram Authors
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

package dao

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"
)

func (m *Dao) SavePhotoSizeV2(ctx context.Context, szId int64, szList []*mtproto.PhotoSize) error {
	for _, sz := range szList {
		szDO := &dataobject.PhotoSizesDO{
			PhotoSizeId: szId,
			SizeType:    sz.Type,
			FilePath:    "",
		}
		if sz.PredicateName == mtproto.Predicate_photoStrippedSize {
			szDO.HasStripped = true
			szDO.FileSize = int32(len(sz.Bytes))
			szDO.StrippedBytes = base64.RawStdEncoding.EncodeToString(sz.Bytes)
		} else {
			// szDO.VolumeId = sz.GetLocation().GetVolumeId()
			// szDO.LocalId = sz.GetLocation().GetLocalId()
			// szDO.Secret = sz.GetLocation().GetSecret()
			szDO.Width = sz.W
			szDO.Height = sz.H
			szDO.FileSize = sz.Size2
		}

		if _, _, err := m.PhotoSizesDAO.Insert(ctx, szDO); err != nil {
			return err
		}
	}

	return nil
}

func (m *Dao) SavePhotoV2(ctx context.Context, id, accessHash int64, hasStickers, hasVideo bool, fileName string) error {
	_, _, err := m.PhotosDAO.Insert(ctx, &dataobject.PhotosDO{
		PhotoId:       id,
		AccessHash:    accessHash,
		HasStickers:   hasStickers,
		DcId:          1,
		Date2:         time.Now().Unix(),
		HasVideo:      hasVideo,
		InputFileName: fileName,
		Ext:           getFileExtName(fileName),
	})
	return err
}

func (m *Dao) GetPhotoSizeListList(ctx context.Context, idList []int64) (sizes map[int64][]*mtproto.PhotoSize) {
	sizes = make(map[int64][]*mtproto.PhotoSize)
	if len(idList) == 0 {
		return
	}

	sizeDOList, _ := m.PhotoSizesDAO.SelectListByPhotoSizeIdList(ctx, idList)
	for i := 0; i < len(sizeDOList); i++ {
		szList, ok := sizes[sizeDOList[i].PhotoSizeId]
		if !ok {
			szList = []*mtproto.PhotoSize{}
		}

		size := &sizeDOList[i]
		if size.SizeType == "i" {
			bytes, _ := base64.RawStdEncoding.DecodeString(size.StrippedBytes)
			if len(bytes) == 0 {
				continue
			}
			szList = append(szList, mtproto.MakeTLPhotoStrippedSize(&mtproto.PhotoSize{
				Type:  size.SizeType,
				Bytes: bytes,
			}).To_PhotoSize())
		} else {
			szList = append(szList, mtproto.MakeTLPhotoSize(&mtproto.PhotoSize{
				Type:  size.SizeType,
				W:     size.Width,
				H:     size.Height,
				Size2: size.FileSize,
			}).To_PhotoSize())
		}
		sizes[sizeDOList[i].PhotoSizeId] = szList
	}
	return
}

func (m *Dao) GetPhotoSizeListV2(ctx context.Context, sizeId int64) (sizes []*mtproto.PhotoSize) {
	sizeDOList, _ := m.PhotoSizesDAO.SelectListByPhotoSizeId(ctx, sizeId)

	if len(sizeDOList) >= 0 {
		sizes = make([]*mtproto.PhotoSize, 0, len(sizeDOList))
		for i := 0; i < len(sizeDOList); i++ {
			size := &sizeDOList[i]
			if size.SizeType == "i" {
				bytes, _ := base64.RawStdEncoding.DecodeString(size.StrippedBytes)
				if len(bytes) == 0 {
					continue
				}
				sizes = append(sizes, mtproto.MakeTLPhotoStrippedSize(&mtproto.PhotoSize{
					Type:  size.SizeType,
					Bytes: bytes,
				}).To_PhotoSize())
			} else {
				sizes = append(sizes, mtproto.MakeTLPhotoSize(&mtproto.PhotoSize{
					Type:  size.SizeType,
					W:     size.Width,
					H:     size.Height,
					Size2: size.FileSize,
				}).To_PhotoSize())
			}
		}
	}

	return
}

func (m *Dao) GetPhotoV2(ctx context.Context, photoId int64) (*mtproto.Photo, error) {
	var (
		photoSizes []*mtproto.PhotoSize
		videoSizes []*mtproto.VideoSize
	)

	photoDO, err := m.PhotosDAO.SelectByPhotoId(ctx, photoId)
	if err != nil {
		return nil, err
	} else if photoDO == nil {
		return emptyPhoto, nil
	}

	photoSizes = m.GetPhotoSizeListV2(ctx, photoDO.PhotoId)
	if photoSizes == nil {
		// photoSizes = []*mtproto.PhotoSize{}
		return emptyPhoto, nil
	}

	if photoDO.HasVideo {
		videoSizes = m.GetVideoSizeList(ctx, photoDO.VideoSizeId)
	}

	photo := mtproto.MakeTLPhoto(&mtproto.Photo{
		Id:            photoId,
		HasStickers:   photoDO.HasStickers,
		AccessHash:    photoDO.AccessHash,
		FileReference: []byte{},
		Date:          int32(photoDO.Date2),
		Sizes:         photoSizes,
		VideoSizes:    videoSizes,
		DcId:          photoDO.DcId,
	}).To_Photo()

	return photo, nil
}
