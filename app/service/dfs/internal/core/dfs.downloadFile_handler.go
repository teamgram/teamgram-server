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
	"fmt"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"
)

// DfsDownloadFile
// dfs.downloadFile location:InputFileLocation offset:int limit:int = upload.File;
func (c *DfsCore) DfsDownloadFile(in *dfs.TLDfsDownloadFile) (*mtproto.Upload_File, error) {
	var (
		bytes []byte
		sType int32

		err error

		location = in.GetLocation()
		offset   = in.GetOffset()
		limit    = in.GetLimit()
	)

	switch location.GetPredicateName() {
	case mtproto.Predicate_inputFileLocation:
		//inputFileLocation#dfdaabe1 volume_id:long local_id:int secret:long file_reference:bytes = InputFileLocation;

		err = mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("dfs.downloadFile - error: %v", err)
		return nil, err

	case mtproto.Predicate_inputEncryptedFileLocation:
		//inputEncryptedFileLocation#f5235d55
		//	id:long
		//	access_hash:long = InputFileLocation;

		bytes, err = c.svcCtx.Dao.GetCacheFile(c.ctx, "encryptedfiles", location.GetId(), offset, limit)
		if err != nil {
			c.Logger.Errorf("download file: %v", err)
			err = nil
			bytes = []byte{}
		}
		sType = int32(mtproto.CRC32_storage_filePartial)
	case mtproto.Predicate_inputDocumentFileLocation:
		//inputDocumentFileLocation#bad07584
		//	id:long
		//	access_hash:long
		//	file_reference:bytes
		//	thumb_size:string = InputFileLocation;

		if location.GetThumbSize() == "" {
			// fileLocation := location.To_InputDocumentFileLocation()
			bytes, err = c.svcCtx.Dao.GetCacheFile(c.ctx, "documents", location.GetId(), offset, limit)
			if err != nil {
				path := fmt.Sprintf("%d.dat", location.GetId())
				bytes, err = c.svcCtx.Dao.GetFile(c.ctx, "documents", path, offset, limit)
				if err != nil {
					c.Logger.Errorf("download file: %v", err)
					err = nil
					bytes = []byte{}
				}
			}
			sType = int32(location.GetAccessHash() >> 32)
		} else {
			path := fmt.Sprintf("%s/%d.dat", location.GetThumbSize(), location.GetId())

			isVideo := mtproto.PhotoSizeIsVideo(location.GetThumbSize())
			c.Logger.Infof("path: %s", path)
			if isVideo {
				bytes, err = c.svcCtx.Dao.GetFile(c.ctx, "videos", path, offset, limit)
				sType = int32(mtproto.CRC32_storage_fileMp4)
			} else {
				bytes, err = c.svcCtx.Dao.GetFile(c.ctx, "photos", path, offset, limit)
				sType = int32(mtproto.CRC32_storage_fileJpeg)
			}
			if err != nil {
				c.Logger.Errorf("download file: %v", err)
				err = nil
				bytes = []byte{}
			}
		}
	case mtproto.Predicate_inputSecureFileLocation:
		//inputSecureFileLocation#cbc7ee28 id:long access_hash:long = InputFileLocation;

		err = mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("dfs.downloadFile - error: %v", err)
		return nil, err
	case mtproto.Predicate_inputTakeoutFileLocation:
		//inputTakeoutFileLocation#29be5899 = InputFileLocation;

		err = mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("dfs.downloadFile - error: %v", err)
		return nil, err
	case mtproto.Predicate_inputPhotoLegacyFileLocation:
		//inputPhotoLegacyFileLocation#d83466f3 id:long access_hash:long file_reference:bytes volume_id:long local_id:int secret:long = InputFileLocation;

		err = mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("dfs.downloadFile - error: %v", err)
		return nil, err
	case mtproto.Predicate_inputPhotoFileLocation:
		//inputPhotoFileLocation#40181ffe id:long access_hash:long file_reference:bytes thumb_size:string = InputFileLocation;

		isVideo := mtproto.PhotoSizeIsVideo(location.GetThumbSize())
		path := fmt.Sprintf("%s/%d.dat", location.GetThumbSize(), location.GetId())
		// log.Debugf("path: %s", path)
		if isVideo {
			bytes, err = c.svcCtx.Dao.GetFile(c.ctx, "videos", path, offset, limit)
			sType = int32(mtproto.CRC32_storage_fileMp4)
		} else {
			bytes, err = c.svcCtx.Dao.GetFile(c.ctx, "photos", path, offset, limit)
			sType = int32(location.GetAccessHash() >> 32)
		}
		if err != nil {
			// log.Warnf("download file: %v", err)
			err = nil
			bytes = []byte{}
		}
	case mtproto.Predicate_inputPeerPhotoFileLocation:
		//inputPeerPhotoFileLocation#37257e99 flags:# big:flags.0?true peer:InputPeer photo_id:long = InputFileLocation;

		var (
			path string
		)

		if location.GetBig() {
			path = fmt.Sprintf("c/%d.dat", location.GetPhotoId())
		} else {
			path = fmt.Sprintf("a/%d.dat", location.GetPhotoId())
		}
		// log.Debugf("path: %s", path)
		bytes, err = c.svcCtx.Dao.GetFile(c.ctx, "photos", path, offset, limit)
		if err != nil {
			c.Logger.Infof("download file: %v", err)
			err = nil
			bytes = []byte{}
		}
		sType = int32(mtproto.CRC32_storage_fileJpeg)
	case mtproto.Predicate_inputStickerSetThumb:
		//inputStickerSetThumb#9d84f3db stickerset:InputStickerSet thumb_version:int = InputFileLocation;

		path := fmt.Sprintf("m/%d.dat", location.GetId())
		c.Logger.Infof("path: %s", path)
		bytes, err = c.svcCtx.Dao.GetFile(c.ctx, "photos", path, offset, limit)
		if err != nil {
			c.Logger.Infof("download file: %v", err)
			err = nil
			bytes = []byte{}
		}
		// sType = int32(location.GetAccessHash() >> 32)
		sType = int32(mtproto.CRC32_storage_fileJpeg)
	case mtproto.Predicate_inputGroupCallStream:
		//inputGroupCallStream#bba51639 call:InputGroupCall time_ms:long scale:int = InputFileLocation;

		err = mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("dfs.downloadFile - error: %v", err)
		return nil, err
	default:
		err = mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("dfs.downloadFile - error: %v", err)
		return nil, err
	}

	uploadFile := mtproto.MakeTLUploadFile(&mtproto.Upload_File{
		Type:  model.MakeStorageFileType(sType),
		Mtime: int32(time.Now().Unix()),
		Bytes: bytes,
	}).To_Upload_File()

	return uploadFile, nil
}
