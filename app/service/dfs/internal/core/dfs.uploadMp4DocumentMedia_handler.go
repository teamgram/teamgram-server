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
	"fmt"
	"image"
	"math/rand"
	"time"

	"github.com/teamgram/marmota/pkg/bytes2"
	"github.com/teamgram/marmota/pkg/threading2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/imaging"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"

	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/threading"
)

// DfsUploadMp4DocumentMedia
// dfs.uploadMp4DocumentMedia creator:long media:InputMedia = Document;
func (c *DfsCore) DfsUploadMp4DocumentMedia(in *dfs.TLDfsUploadMp4DocumentMedia) (*mtproto.Document, error) {
	var (
		documentId = c.svcCtx.Dao.IDGenClient2.NextId(c.ctx)
		path       string
		err        error

		file      = in.GetMedia().GetFile()
		creatorId = in.GetCreator()
		media     = in.GetMedia()

		ext        = model.GetFileExtName(in.GetMedia().GetFile().GetName())
		extType    = model.GetStorageFileTypeConstructor(ext)
		accessHash = int64(extType)<<32 | int64(rand.Uint32())
	)

	var (
		thumbData []byte
		thumb     image.Image
		// photoId   = idgen.GetUUID()
		// ext2      = ".jpg"
		// extType2  = model.GetStorageFileTypeConstructor(ext2)
		// secretId  = int64(extType2)<<32 | int64(rand.Uint32())
	)

	// getFirstFrame
	tmpFileName := fmt.Sprintf("http://127.0.0.1:11701/dfs/file/%d_%d.mp4", creatorId, file.GetId_INT64())
	thumbData, err = c.svcCtx.FFmpegUtil.GetFirstFrame(tmpFileName)
	if thumbData == nil || err != nil {
		// upload mp4 file
		fileInfo, err := c.svcCtx.Dao.GetFileInfo(c.ctx, creatorId, file.Id_INT64)
		if err != nil {
			c.Logger.Errorf("dfs.uploadDocumentFile - error: %v", err)
			return nil, err
		}
		c.svcCtx.Dao.SetCacheFileInfo(c.ctx, documentId, fileInfo)
		path = fmt.Sprintf("%d.dat", documentId)

		threading.RunSafe(func() {
			_, err2 := c.svcCtx.Dao.PutDocumentFile(
				contextx.ValueOnlyFrom(c.ctx),
				path,
				c.svcCtx.Dao.NewSSDBReader(fileInfo))
			if err2 != nil {
				c.Logger.Errorf("dfs.PutDocumentFile - error: %v", err)
			}
		})
		//c.Logger.Errorf("getFirstFrameByPipe - error: %v", err)
		//return nil, err

		attributes := make([]*mtproto.DocumentAttribute, 0, 2)
		attrVideo := mtproto.GetDocumentAttribute(media.GetAttributes(), mtproto.Predicate_documentAttributeVideo)
		if attrVideo != nil {
			attrVideo.SupportsStreaming = true
			attributes = append(attributes, attrVideo)
		}

		attrFileName := mtproto.GetDocumentAttribute(media.GetAttributes(), mtproto.Predicate_documentAttributeFilename)
		if attrFileName != nil {
			attributes = append(attributes, attrFileName)
		}

		// build document
		document := mtproto.MakeTLDocument(&mtproto.Document{
			Id:            documentId,
			AccessHash:    accessHash,
			FileReference: []byte{}, // TODO(@benqi): gen file_reference
			Date:          int32(time.Now().Unix()),
			MimeType:      "video/mp4",
			Size2:         fileInfo.GetFileSize(),
			Size2_INT32:   int32(fileInfo.GetFileSize()),
			Size2_INT64:   fileInfo.GetFileSize(),
			Thumbs:        nil,
			VideoThumbs:   nil,
			DcId:          1,
			Attributes:    attributes,
		}).To_Document()
		return document, nil
	} else {
		// 1. getFirstFrame
		// build photoStrippedSize
		thumb, err = imaging.Decode(bytes.NewReader(thumbData))
		if err != nil {
			return nil, err

		}
		stripped := bytes2.NewBuffer(make([]byte, 0, 4096))
		if thumb.Bounds().Dx() >= thumb.Bounds().Dy() {
			err = imaging.EncodeStripped(stripped, imaging.Resize(thumb, 40, 0), 30)
		} else {
			err = imaging.EncodeStripped(stripped, imaging.Resize(thumb, 0, 40), 30)
		}
		if err != nil {
			return nil, err
		}

		// upload thumb
		var (
			mThumbData = bytes2.NewBuffer(make([]byte, 0, len(thumbData)))
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

		// upload mp4 file
		fileInfo, err := c.svcCtx.Dao.GetFileInfo(c.ctx, creatorId, file.Id_INT64)
		if err != nil {
			c.Logger.Errorf("dfs.uploadDocumentFile - error: %v", err)
			return nil, err
		}
		c.svcCtx.Dao.SetCacheFileInfo(c.ctx, documentId, fileInfo)
		path = fmt.Sprintf("%d.dat", documentId)

		threading2.GoSafeContext(c.ctx, func(ctx context.Context) {
			_, err2 := c.svcCtx.Dao.PutDocumentFile(
				ctx,
				path,
				c.svcCtx.Dao.NewSSDBReader(fileInfo))
			if err2 != nil {
				c.Logger.Errorf("dfs.PutDocumentFile - error: %v", err)
			}
		})

		attributes := make([]*mtproto.DocumentAttribute, 0, 2)
		attrVideo := mtproto.GetDocumentAttribute(media.GetAttributes(), mtproto.Predicate_documentAttributeVideo)
		if attrVideo != nil {
			attrVideo.SupportsStreaming = true
			attributes = append(attributes, attrVideo)
		}

		attrFileName := mtproto.GetDocumentAttribute(media.GetAttributes(), mtproto.Predicate_documentAttributeFilename)
		if attrFileName != nil {
			attributes = append(attributes, attrFileName)
		}

		// build document
		document := mtproto.MakeTLDocument(&mtproto.Document{
			Id:            documentId,
			AccessHash:    accessHash,
			FileReference: []byte{}, // TODO(@benqi): gen file_reference
			Date:          int32(time.Now().Unix()),
			MimeType:      "video/mp4",
			Size2:         fileInfo.GetFileSize(),
			Size2_INT32:   int32(fileInfo.GetFileSize()),
			Size2_INT64:   fileInfo.GetFileSize(),
			Thumbs:        szList,
			VideoThumbs:   nil,
			DcId:          1,
			Attributes:    attributes,
		}).To_Document()

		return document, nil
	}
}
