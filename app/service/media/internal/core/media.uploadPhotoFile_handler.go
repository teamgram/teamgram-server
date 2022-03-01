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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// MediaUploadPhotoFile
// media.uploadPhotoFile flags:# owner_id:long file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = Photo;
func (c *MediaCore) MediaUploadPhotoFile(in *media.TLMediaUploadPhotoFile) (*mtproto.Photo, error) {
	var (
		inputFile = in.GetFile()
		// fileMDList []*dfspb.PhotoFileMetadata
	)

	if in.GetFile() == nil {
		c.Logger.Errorf("media.uploadPhotoFile - error: file is nil")
		return nil, mtproto.ErrMediaInvalid
	}

	photo, err := c.svcCtx.Dao.DfsClient.DfsUploadPhotoFileV2(c.ctx, &dfs.TLDfsUploadPhotoFileV2{
		Creator: in.OwnerId,
		File:    inputFile,
	})
	if err != nil {
		c.Logger.Errorf("media.uploadPhotoFile - error: %v", err)
		return nil, err
	}

	if err = c.svcCtx.Dao.SavePhotoSizeV2(c.ctx, photo.GetId(), photo.GetSizes()); err != nil {
		c.Logger.Errorf("media.uploadPhotoFile - error: %v", err)
		return nil, err
	}

	c.svcCtx.Dao.SavePhotoV2(c.ctx,
		photo.GetId(),
		photo.GetAccessHash(),
		photo.GetHasStickers(),
		false,
		inputFile.GetName())

	return photo, nil
}
