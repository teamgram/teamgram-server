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
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/imaging"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"

	"github.com/minio/minio-go/v7"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// DfsUploadProfilePhotoFileV2
// dfs.uploadProfilePhotoFileV2 flags:# creator:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = Photo;
func (c *DfsCore) DfsUploadProfilePhotoFileV2(in *dfs.TLDfsUploadProfilePhotoFileV2) (*mtproto.Photo, error) {
	if in.GetFile() == nil && in.GetVideo() == nil {
		c.Logger.Errorf("dfs.uploadPhotoFile - ErrInputRequestInvalid")
		return nil, mtproto.ErrInputRequestInvalid
	}

	var (
		file    *mtproto.InputFile
		err     error
		photoV2 *mtproto.Photo
	)

	if in.GetFile() != nil {
		file = in.GetFile()
		//if file == nil {
		//	c.Logger.Errorf("dfs.uploadPhotoFile - ErrInputRequestInvalid")
		//	return nil, mtproto.ErrInputRequestInvalid
		//}

		if err = model.CheckFileParts(file.Parts); err != nil {
			c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
			return nil, err
		}

		photoV2, err = c.uploadPhotoSizeListV2(in.GetCreator(), in.GetFile(), true)
		if err != nil {
			c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
			return nil, err
		}
	} else {
		file = in.GetVideo()

		if err = model.CheckFileParts(file.Parts); err != nil {
			c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
			return nil, err
		}

		photoV2, err = c.uploadVideoSizeListV2(in.GetCreator(), file, in.GetVideoStartTs().GetValue())
		if err != nil {
			c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
			return nil, err
		}
	}

	return photoV2, nil
}

func (c *DfsCore) uploadVideoSizeListV2(creatorId int64, video *mtproto.InputFile, videoStartTs float64) (photo *mtproto.Photo, err error) {
	var (
		videoMp4Data []byte
		photoData    []byte
		// photoImg     image.Image
		r *dao.SSDBReader
	)

	r, err = c.svcCtx.Dao.OpenFile(c.ctx, creatorId, video.Id_INT64, video.Parts)
	if err != nil {
		c.Logger.Errorf("dfs.uploadVideoSizeList - %v", err)
		return nil, mtproto.ErrMediaInvalid
	}

	videoMp4Data, err = r.ReadAll(c.ctx)
	if err != nil {
		c.Logger.Errorf("dfs.uploadVideoSizeList - %v", err)
		return nil, mtproto.ErrMediaInvalid
	}

	if len(video.Md5Checksum) > 0 {
		digest := fmt.Sprintf("%x", md5.Sum(videoMp4Data))
		if digest != video.Md5Checksum {
			c.Logger.Errorf("dfs.uploadPhotoFile - (%s, %s) error: %v", digest, video.Md5Checksum, err)
			return nil, mtproto.ErrCheckSumInvalid
		}
	}

	var (
		// videoId = idgen.GetUUID()
		// ext         = model.GetFileExtName(video.GetName())
		// extType     = model.GetStorageFileTypeConstructor(ext)
		tmpFileName = fmt.Sprintf("http://127.0.0.1:11701/dfs/file/%d_%d.mp4", creatorId, video.GetId_INT64())
	)

	videoMp4Data, _, err = c.svcCtx.FFmpegUtil.ConvertToMp4ByPipe(tmpFileName, 800, 800)
	if err != nil {
		c.Logger.Errorf("uploadVideoSizeList - error: %v", err)
		return nil, err
	}

	// getFirstFrame
	photoData, err = c.svcCtx.FFmpegUtil.GetFirstFrameByPipe(videoMp4Data)
	if err != nil {
		c.Logger.Errorf("getFirstFrameByPipe - error: %v", err)
		return nil, err
	}

	var (
		photoId  = c.svcCtx.Dao.IDGenClient2.NextId(c.ctx)
		sizeList = make([]*mtproto.PhotoSize, 0, 3)
		// extType  = model.GetStorageFileTypeConstructor(ext)
		fileSize minio.UploadInfo
		//ext         = model.GetFileExtName(video.GetName())
		ext        = ".jpg"
		extType    = model.GetStorageFileTypeConstructor(ext)
		accessHash = int64(extType)<<32 | int64(rand.Uint32())
	)

	err = imaging.ReSizeImage(photoData, ext, true, func(szType string, localId int, w, h int32, b []byte) error {
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

	if len(sizeList) == 0 {
		err = mtproto.ErrImageProcessFailed
		c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
		return nil, err
	}

	// ext = model.GetFileExtName(video.GetName())
	// extType = model.GetStorageFileTypeConstructor(ext)
	// secretId := int64(extType)<<32 | int64(rand.Uint32())
	path := fmt.Sprintf("%s/%d.dat", mtproto.VideoSZVType, photoId)
	_, err = c.svcCtx.Dao.PutVideoFile(c.ctx, path, videoMp4Data)
	if err != nil {
		c.Logger.Errorf("uploadVideoSizeList - error: %v", err)
		return nil, err
	}

	videoSize := mtproto.MakeTLVideoSize(&mtproto.VideoSize{
		Type:         mtproto.VideoSZVType,
		W:            800,
		H:            800,
		Size2:        int32(len(videoMp4Data)),
		VideoStartTs: nil,
	}).To_VideoSize()

	if videoStartTs > 0 {
		videoSize.VideoStartTs = &wrapperspb.DoubleValue{Value: videoStartTs}
	}

	return mtproto.MakeTLPhoto(&mtproto.Photo{
		Id:            photoId,
		HasStickers:   false,
		AccessHash:    accessHash,
		FileReference: []byte{},
		Date:          int32(time.Now().Unix()),
		Sizes:         sizeList,
		VideoSizes:    []*mtproto.VideoSize{videoSize},
		DcId:          1,
	}).To_Photo(), nil
}

func (c *DfsCore) uploadPhotoSizeListV2(creatorId int64, file *mtproto.InputFile, isABC bool) (photo *mtproto.Photo, err error) {
	var (
		fileSize  minio.UploadInfo
		cacheData []byte
		r         *dao.SSDBReader
	)

	r, err = c.svcCtx.Dao.OpenFile(c.ctx, creatorId, file.Id_INT64, file.Parts)
	if err != nil {
		c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
		return nil, mtproto.ErrMediaInvalid
	}

	cacheData, err = r.ReadAll(c.ctx)
	if err != nil {
		c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
		return nil, mtproto.ErrMediaInvalid
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

	err = imaging.ReSizeImage(cacheData, ext, isABC, func(szType string, localId int, w, h int32, b []byte) error {
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

	if len(sizeList) == 0 {
		err = mtproto.ErrImageProcessFailed
		c.Logger.Errorf("dfs.uploadPhotoFile - %v", err)
		return nil, err
	}

	return mtproto.MakeTLPhoto(&mtproto.Photo{
		Id:            photoId,
		HasStickers:   false,
		AccessHash:    accessHash,
		FileReference: []byte{},
		Date:          int32(time.Now().Unix()),
		Sizes:         sizeList,
		VideoSizes:    nil,
		DcId:          1,
	}).To_Photo(), nil
}
