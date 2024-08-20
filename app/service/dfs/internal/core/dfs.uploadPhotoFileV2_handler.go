/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"crypto/md5"
	"fmt"
	"math/rand"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/imaging"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"

	"github.com/minio/minio-go/v7"
)

// DfsUploadPhotoFileV2
// dfs.uploadPhotoFileV2 creator:long file:InputFile = Photo;
func (c *DfsCore) DfsUploadPhotoFileV2(in *dfs.TLDfsUploadPhotoFileV2) (*mtproto.Photo, error) {
	var (
		fileSize  minio.UploadInfo
		cacheData []byte
		file      = in.GetFile()
		err       error
		r         *dao.SSDBReader
	)

	if file == nil {
		c.Logger.Errorf("dfs.uploadPhotoFile - ErrInputRequestInvalid")
		return nil, mtproto.ErrInputRequestInvalid
	}

	if err = model.CheckFileParts(file.Parts); err != nil {
		c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
		return nil, err
	}

	r, err = c.svcCtx.Dao.OpenFile(c.ctx, in.GetCreator(), file.Id_INT64, file.Parts)
	if err != nil {
		c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
		return nil, mtproto.ErrMediaInvalid
	}

	cacheData, err = r.ReadAll(c.ctx)
	if err != nil {
		c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
		return nil, mtproto.ErrMediaInvalid
	} else {
		// log.Debugf("cacheData: %s", hex.EncodeToString(cacheData))
	}

	if len(file.Md5Checksum) > 0 {
		digest := fmt.Sprintf("%x", md5.Sum(cacheData))
		if digest != file.Md5Checksum {
			c.Logger.Errorf("dfs.uploadPhotoFile - (%s, %s) error: %v", digest, file.Md5Checksum, err)
			return nil, mtproto.ErrCheckSumInvalid
		}
	}

	var (
		photoId    = c.svcCtx.Dao.IDGenClient2.NextId(c.ctx)
		ext        = model.GetFileExtName(file.GetName())
		sizeList   = make([]*mtproto.PhotoSize, 0, 3)
		extType    = model.GetStorageFileTypeConstructor(ext)
		accessHash = int64(extType)<<32 | int64(rand.Uint32())
	)

	err = imaging.ReSizeImage(cacheData, ext, false, func(szType string, localId int, w, h int32, b []byte) error {
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
		Date:          int32(r.Mtime),
		Sizes:         sizeList,
		VideoSizes:    nil,
		DcId:          1,
	}).To_Photo()

	return photo, nil
}
