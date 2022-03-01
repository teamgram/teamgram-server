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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// MediaUploadEncryptedFile
// media.uploadEncryptedFile owner_id:long file:InputEncryptedFile = EncryptedFile;
func (c *MediaCore) MediaUploadEncryptedFile(in *media.TLMediaUploadEncryptedFile) (*mtproto.EncryptedFile, error) {
	inputFile := in.File
	if inputFile == nil {
		return nil, fmt.Errorf("bad request")
	}

	encryptedFile, err := c.svcCtx.Dao.DfsClient.DfsUploadEncryptedFileV2(c.ctx, &dfs.TLDfsUploadEncryptedFileV2{
		Constructor:          0,
		Creator:              in.OwnerId,
		File:                 inputFile,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	if err != nil {
		c.Logger.Errorf("media.uploadEncryptedFile - error: %v", err)
		return nil, err
	}
	c.svcCtx.Dao.SaveEncryptedFileV2(c.ctx, encryptedFile)

	return encryptedFile, nil
}
