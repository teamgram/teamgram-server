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

// MediaUploadThemeFile
// media.uploadThemeFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
func (c *MediaCore) MediaUploadThemeFile(in *media.TLMediaUploadThemeFile) (*mtproto.Document, error) {
	var (
		err      error
		document *mtproto.Document
		file     = in.GetFile()
		thumb    = in.GetThumb()
	)

	if file == nil {
		c.Logger.Errorf("media.uploadThemeFile - error: file is nil")
		return nil, mtproto.ErrThemeFileInvalid
	}

	document, err = c.svcCtx.Dao.DfsClient.DfsUploadThemeFile(c.ctx, &dfs.TLDfsUploadThemeFile{
		Creator:  in.OwnerId,
		File:     file,
		Thumb:    thumb,
		MimeType: in.GetFileName(),
		FileName: in.GetMimeType(),
	})
	if err != nil {
		c.Logger.Errorf("media.uploadThemeFile - error: %v", err)
		err = mtproto.ErrThemeFileInvalid
		return nil, err
	}

	if len(document.GetThumbs()) > 0 {
		c.svcCtx.Dao.SavePhotoSizeV2(c.ctx, document.GetId(), document.GetThumbs())
	}
	c.svcCtx.Dao.SaveDocumentV2(c.ctx, file.GetName(), document)

	return document, nil
}
