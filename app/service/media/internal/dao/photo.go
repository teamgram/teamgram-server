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
	"fmt"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/media/media"

	"github.com/zeromicro/go-zero/core/jsonx"
)

var (
	cachePhotoPrefix = "photo"
	GetPhotoSize     = getPhotoSize
)

func genCachePhotoKey(id int64) string {
	return fmt.Sprintf("%s#%d", cachePhotoPrefix, id)
}

type CachePhotoData struct {
	Id            int64                      `json:"id"`
	Photo         *dataobject.PhotosDO       `json:"photo"`
	SizeList      []*dataobject.PhotoSizesDO `json:"size_list"`
	VideoSizeList []*dataobject.VideoSizesDO `json:"video_size_list"`
}

func (c *CachePhotoData) ToPhoto() *mtproto.Photo {
	if c == nil {
		return makePhotoEmpty(0)
	}

	if c.Photo == nil {
		return makePhotoEmpty(c.Id)
	}

	if len(c.SizeList) == 0 {
		return makePhotoEmpty(c.Id)
	}

	return mtproto.MakeTLPhoto(&mtproto.Photo{
		Id:            c.Id,
		HasStickers:   c.Photo.HasStickers,
		AccessHash:    c.Photo.AccessHash,
		FileReference: []byte{},
		Date:          int32(c.Photo.Date2),
		Sizes:         c.ToSizes(),
		VideoSizes:    c.ToVideoSizes(),
		DcId:          c.Photo.DcId,
	}).To_Photo()
}

func (c *CachePhotoData) ToSizes() []*mtproto.PhotoSize {
	szList := make([]*mtproto.PhotoSize, 0, len(c.SizeList))

	for _, sz := range c.SizeList {
		szList = append(szList, getPhotoSize(sz))
	}

	return szList
}

func (c *CachePhotoData) ToPhotoSizeList() *media.PhotoSizeList {
	return &media.PhotoSizeList{
		SizeId: c.Id,
		Sizes:  c.ToSizes(),
		DcId:   1,
	}
}

func (c *CachePhotoData) ToVideoSizes() []*mtproto.VideoSize {
	if !c.Photo.HasVideo || len(c.VideoSizeList) == 0 {
		return nil
	}

	szList := make([]*mtproto.VideoSize, 0, len(c.VideoSizeList))

	for _, sz := range c.VideoSizeList {
		szList = append(szList, getVideoSize(sz))
	}

	return szList
}

func (c *CachePhotoData) ToVideoSizeList() *media.VideoSizeList {
	return &media.VideoSizeList{
		SizeId: c.Id,
		Sizes:  c.ToVideoSizes(),
		DcId:   1,
	}
}

func makePhotoSizesDO(szId int64, sz *mtproto.PhotoSize) *dataobject.PhotoSizesDO {
	szDO := &dataobject.PhotoSizesDO{
		PhotoSizeId: szId,
		SizeType:    sz.Type,
		Width:       sz.W,
		Height:      sz.H,
		FileSize:    sz.Size2,
		FilePath:    fmt.Sprintf("%s/%d.dat", sz.Type, szId),
		CachedType:  0,
		CachedBytes: "",
	}

	switch sz.GetPredicateName() {
	case mtproto.Predicate_photoPathSize:
		szDO.CachedType = CachedTypePathSize
		szDO.CachedBytes = base64.RawStdEncoding.EncodeToString(sz.Bytes)
		szDO.FileSize = int32(len(sz.Bytes))
	case mtproto.Predicate_photoStrippedSize:
		szDO.CachedType = CachedTypeStrippedSize
		szDO.CachedBytes = base64.RawStdEncoding.EncodeToString(sz.Bytes)
		szDO.FileSize = int32(len(sz.Bytes))
	case mtproto.Predicate_photoCachedSize:
		szDO.CachedType = CachedTypeSizeProgressive
		szDO.CachedBytes = base64.RawStdEncoding.EncodeToString(sz.Bytes)
		szDO.FileSize = int32(len(sz.Bytes))
	case mtproto.Predicate_photoSizeProgressive:
		szDO.CachedType = CachedTypeSizeProgressive
		cachedBytes, _ := jsonx.Marshal(sz.Sizes)
		if cachedBytes != nil {
			szDO.CachedBytes = string(cachedBytes)
		}
		szDO.FileSize = sz.Size2
	case mtproto.Predicate_photoSize:
		szDO.CachedType = CachedTypeSize
		szDO.FileSize = sz.Size2
	default:
		// TODO: log
		return nil
	}

	return szDO
}

func getPhotoSize(szDO *dataobject.PhotoSizesDO) *mtproto.PhotoSize {
	switch szDO.SizeType {
	case "j":
		bytes, _ := base64.RawStdEncoding.DecodeString(szDO.CachedBytes)
		if len(bytes) == 0 {
			return nil
		}
		return mtproto.MakeTLPhotoPathSize(&mtproto.PhotoSize{
			Type:  szDO.SizeType,
			Bytes: bytes,
		}).To_PhotoSize()
	case "i":
		bytes, _ := base64.RawStdEncoding.DecodeString(szDO.CachedBytes)
		if len(bytes) == 0 {
			return nil
		}
		return mtproto.MakeTLPhotoStrippedSize(&mtproto.PhotoSize{
			Type:  szDO.SizeType,
			Bytes: bytes,
		}).To_PhotoSize()
	default:
		switch szDO.CachedType {
		case CachedTypeCachedSize:
			bytes, _ := base64.RawStdEncoding.DecodeString(szDO.CachedBytes)
			if len(bytes) == 0 {
				return nil
			}
			return mtproto.MakeTLPhotoCachedSize(&mtproto.PhotoSize{
				Type:  szDO.SizeType,
				W:     szDO.Width,
				H:     szDO.Height,
				Bytes: bytes,
			}).To_PhotoSize()
		case CachedTypeSizeProgressive:
			if len(szDO.CachedBytes) == 0 {
				return nil
			}
			var (
				sizes []int32
			)
			err := jsonx.UnmarshalFromString(szDO.CachedBytes, sizes)
			if err != nil {
				return nil
			}
			return mtproto.MakeTLPhotoSizeProgressive(&mtproto.PhotoSize{
				Type:  szDO.SizeType,
				W:     szDO.Width,
				H:     szDO.Height,
				Sizes: sizes,
			}).To_PhotoSize()
		case CachedTypeSize:
			return mtproto.MakeTLPhotoSize(&mtproto.PhotoSize{
				Type:  szDO.SizeType,
				W:     szDO.Width,
				H:     szDO.Height,
				Size2: szDO.FileSize,
			}).To_PhotoSize()
		default:
			return nil
		}
	}
}

func (m *Dao) SavePhotoSizeV2(ctx context.Context, szId int64, szList []*mtproto.PhotoSize) error {
	for _, sz := range szList {
		szDO := makePhotoSizesDO(szId, sz)
		if szDO == nil {
			continue
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

		sz := getPhotoSize(&sizeDOList[i])
		if sz != nil {
			szList = append(szList, sz)
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
			sz := getPhotoSize(&sizeDOList[i])
			if sz != nil {
				sizes = append(sizes, sz)
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
		videoSizes = m.GetVideoSizeList(ctx, photoDO.PhotoId)
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

func (m *Dao) GetCachePhotoData(ctx context.Context, photoId int64) (*CachePhotoData, error) {
	cacheData := &CachePhotoData{
		Id:            photoId,
		Photo:         nil,
		SizeList:      nil,
		VideoSizeList: nil,
	}

	err := m.CachedConn.QueryRow(
		ctx,
		cacheData,
		genCachePhotoKey(photoId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			photoDO, err := m.PhotosDAO.SelectByPhotoId(ctx, photoId)
			if err != nil {
				return err
			}
			//if photoDO == nil {
			//	return sqlc.ErrNotFound
			//}

			cData := v.(*CachePhotoData)
			cData.Photo = photoDO

			_, err = m.PhotoSizesDAO.SelectListByPhotoSizeIdWithCB(
				ctx,
				photoId,
				func(sz, i int, v *dataobject.PhotoSizesDO) {
					cData.SizeList = append(cData.SizeList, v)
				})
			if err != nil {
				return err
			}

			if photoDO != nil && photoDO.HasVideo {
				m.VideoSizesDAO.SelectListByVideoSizeIdWithCB(
					ctx,
					photoId,
					func(sz, i int, v *dataobject.VideoSizesDO) {
						cData.VideoSizeList = append(cData.VideoSizeList, v)
					})
			}

			return nil
		})

	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}

	return cacheData, nil
}
