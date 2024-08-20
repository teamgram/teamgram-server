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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/imaging"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"
)

// DfsUploadWallPaperFile
// dfs.uploadWallPaperFile creator:long file:InputFile mime_type:string admin:Bool = Document;
func (c *DfsCore) DfsUploadWallPaperFile(in *dfs.TLDfsUploadWallPaperFile) (*mtproto.Document, error) {
	var (
		documentId = c.svcCtx.Dao.IDGenClient2.NextId(c.ctx)
		path       string
		err        error

		file       = in.GetFile()
		cacheData  []byte
		ext        = model.GetFileExtName(file.GetName())
		extType    = model.GetStorageFileTypeConstructor(ext)
		accessHash = int64(extType)<<32 | int64(rand.Uint32())
		r          *dao.SSDBReader
	)

	var (
		thumb image.Image
		// photoId = idgen.GetUUID()
		// ext2     = ".jpg"
		// extType2 = model.GetStorageFileTypeConstructor(ext2)
		// secretId = int64(extType2)<<32 | int64(rand.Uint32())
	)

	if file == nil {
		c.Logger.Errorf("dfs.uploadWallPaperFile - ErrInputRequestInvalid")
		return nil, mtproto.ErrWallpaperFileInvalid
	}

	if err = model.CheckFileParts(file.Parts); err != nil {
		c.Logger.Errorf("dfs.uploadWallPaperFile - %v", err)
		return nil, err
	}

	r, err = c.svcCtx.Dao.OpenFile(c.ctx, in.GetCreator(), file.Id_INT64, file.Parts)
	if err != nil {
		c.Logger.Errorf("dfs.uploadWallPaperFile - %v", err)
		return nil, mtproto.ErrWallpaperFileInvalid
	}

	cacheData, err = r.ReadAll(c.ctx)
	if err != nil {
		c.Logger.Errorf("dfs.uploadWallPaperFile - %v", err)
		return nil, mtproto.ErrWallpaperFileInvalid
	} else {
		// log.Debugf("cacheData: %s", hex.EncodeToString(cacheData))
	}

	if len(file.Md5Checksum) > 0 {
		digest := fmt.Sprintf("%x", md5.Sum(cacheData))
		if digest != file.Md5Checksum {
			c.Logger.Errorf("dfs.uploadWallPaperFile - (%s, %s) error: %v", digest, file.Md5Checksum, err)
			return nil, mtproto.ErrCheckSumInvalid
		}
	}

	// build photoStrippedSize
	thumb, err = imaging.Decode(bytes.NewReader(cacheData))
	if err != nil {
		// ioutil.WriteFile("./t.jpg", cacheData, 0644)
		c.Logger.Errorf("dfs.uploadWallPaperFile - error: %v", err)
		return nil, err

	}
	stripped := bytes2.NewBuffer(make([]byte, 0, 4096))
	if thumb.Bounds().Dx() >= thumb.Bounds().Dy() {
		err = imaging.EncodeStripped(stripped, imaging.Resize(thumb, 40, 0), 30)
	} else {
		err = imaging.EncodeStripped(stripped, imaging.Resize(thumb, 0, 40), 30)
	}
	if err != nil {
		c.Logger.Errorf("dfs.uploadWallPaperFile - error: %v", err)
		return nil, err
	}

	// upload thumb
	var (
		mThumbData = bytes2.NewBuffer(make([]byte, 0, len(cacheData)))
		mThumb     image.Image
	)
	if thumb.Bounds().Dx() >= thumb.Bounds().Dy() {
		mThumb = imaging.Resize(thumb, 320, 0)
		// err = imaging.Encode(mThumbData, mThumb, 80)
	} else {
		mThumb = imaging.Resize(thumb, 0, 320)
		// err = imaging.Encode(mThumbData, imaging.Resize(thumb, 0, 320), 80)
	}

	err = imaging.EncodeJpeg(mThumbData, mThumb)
	if err != nil {
		c.Logger.Errorf("dfs.uploadWallPaperFile - error: %v", err)
		return nil, err
	}

	// upload thumb
	path = fmt.Sprintf("%s/%d.dat", mtproto.PhotoSZMediumType, documentId)
	// upload
	c.svcCtx.Dao.PutPhotoFile(c.ctx, path, mThumbData.Bytes())

	szList := []*mtproto.PhotoSize{
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

	c.svcCtx.Dao.SetCacheFileInfo(c.ctx, documentId, r.DfsFileInfo)

	go func() {
		_, err2 := c.svcCtx.Dao.PutDocumentFile(context.Background(),
			fmt.Sprintf("%d.dat", documentId),
			bytes.NewReader(cacheData))
		if err2 != nil {
			c.Logger.Errorf("dfs.uploadWallPaperFile - error: %v", err2)
		}
	}()

	// build document
	document := mtproto.MakeTLDocument(&mtproto.Document{
		Id:            documentId,
		AccessHash:    accessHash,
		FileReference: []byte{}, // TODO(@benqi): gen file_reference
		Date:          int32(time.Now().Unix()),
		MimeType:      in.GetMimeType(),
		Size2:         int64(len(cacheData)),
		Size2_INT32:   int32(len(cacheData)),
		Size2_INT64:   int64(len(cacheData)),
		Thumbs:        szList,
		VideoThumbs:   nil,
		DcId:          1,
		Attributes: []*mtproto.DocumentAttribute{
			mtproto.MakeTLDocumentAttributeImageSize(&mtproto.DocumentAttribute{
				W: int32(thumb.Bounds().Dx()),
				H: int32(thumb.Bounds().Dy()),
			}).To_DocumentAttribute(),
		},
	}).To_Document()

	if mtproto.FromBool(in.GetAdmin()) {
		document.Attributes = append(document.Attributes,
			mtproto.MakeTLDocumentAttributeFilename(&mtproto.DocumentAttribute{
				FileName: file.GetName(),
			}).To_DocumentAttribute())
	}

	return document, nil
}
