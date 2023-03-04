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
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"
)

// DfsWriteFilePartData
// dfs.writeFilePartData flags:# creator:long file_id:long bytes:bytes big:flags.0?true file_total_parts:flags.1?int = Bool;
func (c *DfsCore) DfsWriteFilePartData(in *dfs.TLDfsWriteFilePartData) (*mtproto.Bool, error) {
	var (
		err error
	)

	if len(in.GetBytes()) == 0 {
		err = mtproto.ErrFilePartEmpty
		c.Logger.Errorf("dfs.writeFilePartData - error: %v", err)
		return nil, err
	}

	err = c.svcCtx.Dao.WriteFilePartData(c.ctx, in.Creator, in.FileId, in.FilePart, in.Bytes)
	if err != nil {
		c.Logger.Errorf("dfs.writeFilePartData - error: %v", err)
		return nil, err
	}

	if in.FilePart == 0 {
		err = c.svcCtx.Dao.SetFileInfo(c.ctx, &model.DfsFileInfo{
			Creator:           in.Creator,
			FileId:            in.FileId,
			Big:               in.Big,
			FileName:          "",
			FileTotalParts:    int(in.GetFileTotalParts().GetValue()),
			FirstFilePartSize: len(in.Bytes),
			FilePartSize:      0,
			LastFilePartSize:  0,
			MimeType:          "",
			Mtime:             time.Now().Unix(),
		})

		// TODO(@benqi): error
		if err != nil {
			c.Logger.Errorf("dfs.writeFilePartData - error: %v", err)
			return nil, mtproto.ErrFilePartInvalid
		}
	} else if in.FilePart == 1 {
		err = c.svcCtx.Dao.SetFileInfo(c.ctx, &model.DfsFileInfo{
			Creator:      in.Creator,
			FileId:       in.FileId,
			FilePartSize: len(in.Bytes),
		})

		// TODO(@benqi): error
		if err != nil {
			c.Logger.Errorf("dfs.writeFilePartData - error: %v", err)
			return nil, mtproto.ErrFilePartInvalid
		}
	}

	if in.GetFileTotalParts() != nil {
		if in.GetFileTotalParts().GetValue() == in.FilePart+1 {
			err = c.svcCtx.Dao.SetFileInfo(
				c.ctx,
				&model.DfsFileInfo{
					Creator:          in.Creator,
					FileId:           in.FileId,
					LastFilePartSize: len(in.Bytes),
				})

			// TODO(@benqi): error
			if err != nil {
				c.Logger.Errorf("dfs.writeFilePartData - error: %v", err)
				return nil, mtproto.ErrFilePartInvalid
			}
		}
	}

	return mtproto.BoolTrue, nil
}
