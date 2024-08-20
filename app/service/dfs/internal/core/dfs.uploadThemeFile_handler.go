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
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"image"
	"math/rand"
	"time"

	"github.com/teamgram/marmota/pkg/bytes2"
	"github.com/teamgram/marmota/pkg/threading2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/imaging"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"
)

// DfsUploadThemeFile
// dfs.uploadThemeFile flags:# creator:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
func (c *DfsCore) DfsUploadThemeFile(in *dfs.TLDfsUploadThemeFile) (*mtproto.Document, error) {
	var (
		path string
		err  error

		documentId = c.svcCtx.Dao.IDGenClient2.NextId(c.ctx)
		file       = in.GetFile()
		ext        = model.GetFileExtName(file.GetName())
		extType    = model.GetStorageFileTypeConstructor(ext)
		accessHash = int64(extType)<<32 | int64(rand.Uint32())
	)

	var (
		thumbFile     = in.GetThumb()
		thumbSizeList []*mtproto.PhotoSize
	)

	// upload file
	if file == nil {
		c.Logger.Errorf("dfs.uploadThemeFile - ErrInputRequestInvalid")
		return nil, mtproto.ErrWallpaperFileInvalid
	}

	if err = model.CheckFileParts(file.Parts); err != nil {
		c.Logger.Errorf("dfs.uploadThemeFile - %v", err)
		return nil, err
	}

	fileInfo, err := c.svcCtx.Dao.GetFileInfo(c.ctx, in.GetCreator(), file.Id_INT64)
	if err != nil {
		c.Logger.Errorf("dfs.uploadThemeFile - error: %v", err)
		return nil, err
	}
	c.svcCtx.Dao.SetCacheFileInfo(c.ctx, documentId, fileInfo)

	threading2.GoSafeContext(c.ctx, func(ctx context.Context) {
		_, err2 := c.svcCtx.Dao.PutDocumentFile(
			ctx,
			fmt.Sprintf("%d.dat", documentId),
			c.svcCtx.Dao.NewSSDBReader(fileInfo))
		if err2 != nil {
			c.Logger.Errorf("dfs.uploadThemeFile - error: %v", err2)
		}
	})

	// upload thumb file
	if thumbFile != nil {
		var (
			thumbCacheData []byte
			thumb          image.Image
			// photoId        = idgen.GetUUID()
			// ext2           = request.GetThumb().GetName()
			// extType2       = model.GetStorageFileTypeConstructor(ext2)
			// secretId       = int64(extType2)<<32 | int64(rand.Uint32())
			r *dao.SSDBReader
		)
		r, err = c.svcCtx.Dao.OpenFile(c.ctx, in.GetCreator(), thumbFile.Id_INT64, thumbFile.Parts)
		if err != nil {
			c.Logger.Errorf("dfs.uploadThemeFile - %v", err)
			return nil, mtproto.ErrThemeFileInvalid
		}

		thumbCacheData, err = r.ReadAll(c.ctx)
		if err != nil {
			c.Logger.Errorf("dfs.uploadThemeFile - %v", err)
			return nil, mtproto.ErrThemeFileInvalid
		}

		if len(thumbFile.Md5Checksum) > 0 {
			digest := fmt.Sprintf("%x", md5.Sum(thumbCacheData))
			if digest != thumbFile.Md5Checksum {
				c.Logger.Errorf("dfs.uploadThemeFile - (%s, %s) error: %v", digest, thumbFile.Md5Checksum, err)
				return nil, mtproto.ErrCheckSumInvalid
			}
		}

		// build photoStrippedSize
		thumb, err = imaging.Decode(bytes.NewReader(thumbCacheData))
		if err != nil {
			c.Logger.Errorf("dfs.uploadThemeFile - error: %v", err)
			return nil, err

		}
		stripped := bytes2.NewBuffer(make([]byte, 0, 4096))
		if thumb.Bounds().Dx() >= thumb.Bounds().Dy() {
			err = imaging.EncodeStripped(stripped, imaging.Resize(thumb, 40, 0), 30)
		} else {
			err = imaging.EncodeStripped(stripped, imaging.Resize(thumb, 0, 40), 30)
		}
		if err != nil {
			c.Logger.Errorf("dfs.uploadThemeFile - error: %v", err)
			return nil, err
		}

		// upload thumb
		var (
			mThumbData = bytes2.NewBuffer(make([]byte, 0, len(thumbCacheData)))
			mThumb     image.Image
		)
		if thumb.Bounds().Dx() >= thumb.Bounds().Dy() {
			mThumb = imaging.Resize(thumb, 320, 0)
		} else {
			mThumb = imaging.Resize(thumb, 0, 320)
		}

		err = imaging.EncodeJpeg(mThumbData, mThumb)
		if err != nil {
			c.Logger.Errorf("dfs.uploadThemeFile - error: %v", err)
			return nil, err
		}

		// upload thumb
		path = fmt.Sprintf("%s/%d.dat", mtproto.PhotoSZMediumType, documentId)
		// upload
		c.svcCtx.Dao.PutPhotoFile(c.ctx, path, mThumbData.Bytes())

		thumbSizeList = []*mtproto.PhotoSize{
			mtproto.MakeTLPhotoStrippedSize(&mtproto.PhotoSize{
				Type:  mtproto.PhotoSZStrippedType,
				Bytes: stripped.Bytes(),
			}).To_PhotoSize(),
			mtproto.MakeTLPhotoSize(&mtproto.PhotoSize{
				Type:  mtproto.PhotoSZMediumType,
				W:     int32(mThumb.Bounds().Dx()),
				H:     int32(mThumb.Bounds().Dy()),
				Size2: int32(len(mThumbData.Bytes())),
			}).To_PhotoSize(),
		}
	}

	// build document
	document := mtproto.MakeTLDocument(&mtproto.Document{
		Id:            documentId,
		AccessHash:    accessHash,
		FileReference: []byte{}, // TODO(@benqi): gen file_reference
		Date:          int32(time.Now().Unix()),
		MimeType:      in.GetMimeType(),
		Size2:         fileInfo.GetFileSize(),
		Size2_INT32:   int32(fileInfo.GetFileSize()),
		Size2_INT64:   fileInfo.GetFileSize(),
		Thumbs:        thumbSizeList,
		VideoThumbs:   nil,
		DcId:          1,
		Attributes: []*mtproto.DocumentAttribute{
			mtproto.MakeTLDocumentAttributeFilename(&mtproto.DocumentAttribute{
				FileName: in.GetFileName(),
			}).To_DocumentAttribute(),
		},
	}).To_Document()

	return document, nil
}
