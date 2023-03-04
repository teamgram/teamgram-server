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
	"context"
	"fmt"
	"math/rand"

	"github.com/teamgram/marmota/pkg/threading2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
)

// DfsUploadEncryptedFileV2
// dfs.uploadEncryptedFileV2 creator:long file:InputEncryptedFile = EncryptedFile;
func (c *DfsCore) DfsUploadEncryptedFileV2(in *dfs.TLDfsUploadEncryptedFileV2) (*mtproto.EncryptedFile, error) {
	var (
		file            = in.GetFile()
		creatorId       = in.GetCreator()
		encryptedFileId = c.svcCtx.Dao.IDGenClient2.NextId(c.ctx)
		accessHash      = int64(mtproto.CRC32_storage_filePartial)<<32 | int64(rand.Uint32())
	)

	fileInfo, err := c.svcCtx.Dao.GetFileInfo(c.ctx, creatorId, file.Id)
	if err != nil {
		c.Logger.Errorf("dfs.uploadDocumentFile - error: %v", err)
		return nil, err
	}
	c.svcCtx.Dao.SetCacheFileInfo(c.ctx, encryptedFileId, fileInfo)
	path := fmt.Sprintf("%d.dat", encryptedFileId)

	threading2.GoSafeContext(c.ctx, func(ctx context.Context) {
		_, err2 := c.svcCtx.Dao.PutEncryptedFile(ctx, path, c.svcCtx.Dao.NewSSDBReader(fileInfo))
		if err2 != nil {
			c.Logger.Errorf("dfs.uploadEncryptedFile - error: %v", err)
		}
	})

	encryptedFile := mtproto.MakeTLEncryptedFile(&mtproto.EncryptedFile{
		Id:             encryptedFileId,
		AccessHash:     accessHash,
		Size2_INT32:    int32(fileInfo.GetFileSize()),
		Size2_INT64:    fileInfo.GetFileSize(),
		DcId:           1,
		KeyFingerprint: 0,
	}).To_EncryptedFile()

	return encryptedFile, nil
}
