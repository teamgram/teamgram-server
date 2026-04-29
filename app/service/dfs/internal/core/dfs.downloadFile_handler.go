// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DfsDownloadFile
// dfs.downloadFile location:InputFileLocation offset:long limit:int = upload.File;
func (c *DfsCore) DfsDownloadFile(in *dfs.TLDfsDownloadFile) (*tg.UploadFile, error) {
	repo := c.downloads()
	offset := in.Offset
	limit := in.Limit

	switch location := in.Location.(type) {
	case *tg.TLInputEncryptedFileLocation:
		data, err := c.downloadCachedObject(location.Id, offset, limit)
		if err != nil {
			data, err = repo.GetEncryptedObject(c.ctx, minioadapterObjectPath(location.Id), offset, limit)
			data, err = c.compatibleDownloadBytes("download encrypted object", data, err)
			if err != nil {
				return nil, err
			}
		}
		return makeUploadFile(storageConstructor(tg.ClazzID_storage_filePartial), data), nil
	case *tg.TLInputDocumentFileLocation:
		if location.ThumbSize == "" {
			data, err := c.downloadCachedObject(location.Id, offset, limit)
			if err != nil {
				data, err = repo.GetDocumentObject(c.ctx, minioadapterObjectPath(location.Id), offset, limit)
				data, err = c.compatibleDownloadBytes("download document object", data, err)
				if err != nil {
					return nil, err
				}
			}
			return makeUploadFile(storageTypeFromAccessHash(location.AccessHash), data), nil
		}
		data, storageType, err := c.downloadDocumentThumbObject(location.Id, location.ThumbSize, offset, limit)
		if err != nil {
			return nil, err
		}
		return makeUploadFile(storageType, data), nil
	case *tg.TLInputPhotoFileLocation:
		data, storageType, err := c.downloadPhotoLikeObject(location.Id, location.AccessHash, location.ThumbSize, offset, limit)
		if err != nil {
			return nil, err
		}
		return makeUploadFile(storageType, data), nil
	case *tg.TLInputPeerPhotoFileLocation:
		size := "a"
		if location.Big {
			size = "c"
		}
		data, err := repo.GetPhotoObject(c.ctx, objectPath(size, location.PhotoId), offset, limit)
		data, err = c.compatibleDownloadBytes("download peer photo object", data, err)
		if err != nil {
			return nil, err
		}
		return makeUploadFile(storageConstructor(tg.ClazzID_storage_fileJpeg), data), nil
	case *tg.TLInputStickerSetThumb:
		id, ok := stickerSetID(location.Stickerset)
		if !ok {
			return nil, dfs.ErrDfsInvalidArgument
		}
		data, err := repo.GetPhotoObject(c.ctx, objectPath("m", id), offset, limit)
		data, err = c.compatibleDownloadBytes("download sticker set thumb", data, err)
		if err != nil {
			return nil, err
		}
		return makeUploadFile(storageConstructor(tg.ClazzID_storage_fileJpeg), data), nil
	default:
		return nil, dfs.ErrDfsInvalidArgument
	}
}

func (c *DfsCore) downloadCachedObject(objectID int64, offset int64, limit int32) ([]byte, error) {
	info, err := c.uploadSessions().LoadObjectCacheRef(c.ctx, objectID)
	if err != nil {
		return nil, err
	}
	return c.uploadSessions().ReadUploadedFileRange(c.ctx, info.Creator, info.FileID, offset, int64(limit))
}

func (c *DfsCore) downloadPhotoLikeObject(id int64, accessHash int64, thumbSize string, offset int64, limit int32) ([]byte, int32, error) {
	if thumbSize == "" {
		return nil, 0, dfs.ErrDfsInvalidArgument
	}
	repo := c.downloads()
	path := objectPath(thumbSize, id)
	if photoSizeIsVideo(thumbSize) {
		data, err := repo.GetVideoObject(c.ctx, path, offset, limit)
		data, err = c.compatibleDownloadBytes("download video object", data, err)
		return data, storageConstructor(tg.ClazzID_storage_fileMp4), err
	}
	data, err := repo.GetPhotoObject(c.ctx, path, offset, limit)
	data, err = c.compatibleDownloadBytes("download photo object", data, err)
	return data, storageTypeFromAccessHash(accessHash), err
}

func (c *DfsCore) downloadDocumentThumbObject(id int64, thumbSize string, offset int64, limit int32) ([]byte, int32, error) {
	if thumbSize == "" {
		return nil, 0, dfs.ErrDfsInvalidArgument
	}
	repo := c.downloads()
	path := objectPath(thumbSize, id)
	if photoSizeIsVideo(thumbSize) {
		data, err := repo.GetVideoObject(c.ctx, path, offset, limit)
		data, err = c.compatibleDownloadBytes("download document video thumb", data, err)
		return data, storageConstructor(tg.ClazzID_storage_fileMp4), err
	}
	data, err := repo.GetPhotoObject(c.ctx, path, offset, limit)
	data, err = c.compatibleDownloadBytes("download document photo thumb", data, err)
	return data, storageConstructor(tg.ClazzID_storage_fileJpeg), err
}

func minioadapterObjectPath(id int64) string {
	return minioadapter.ObjectPath(id)
}
