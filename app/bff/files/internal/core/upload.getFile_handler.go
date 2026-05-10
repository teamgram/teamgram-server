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
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UploadGetFile
// upload.getFile#be5335be flags:# precise:flags.0?true cdn_supported:flags.1?true location:InputFileLocation offset:long limit:int = upload.File;
func (c *FilesCore) UploadGetFile(in *tg.TLUploadGetFile) (*tg.UploadFile, error) {
	if in == nil {
		return nil, tg.ErrLocationInvalid
	}
	location := in.Location
	if location == nil {
		return nil, tg.ErrLocationInvalid
	}

	switch location.InputFileLocationClazzName() {
	case tg.ClazzName_inputFileLocation:
		return nil, tg.ErrInputRequestInvalid
	case tg.ClazzName_inputDocumentFileLocation,
		tg.ClazzName_inputPhotoFileLocation,
		tg.ClazzName_inputPeerPhotoFileLocation:
		return c.downloadResolvedFile(location, in.Offset, in.Limit)
	case tg.ClazzName_inputEncryptedFileLocation,
		tg.ClazzName_inputSecureFileLocation,
		tg.ClazzName_inputTakeoutFileLocation:
		if !c.svcCtx.Config.Download.LegacyDownloadFallback {
			return nil, tg.ErrLocationInvalid
		}
		return c.downloadLegacyFile(location, in.Offset, in.Limit)
	case tg.ClazzName_inputStickerSetThumb:
		// TODO: master resolves sticker set thumbs through enterprise plugin hooks.
		return nil, tg.ErrMethodNotImpl
	case tg.ClazzName_inputGroupCallStream:
		// TODO: master resolves group call streams through enterprise plugin hooks.
		return nil, tg.ErrMethodNotImpl
	default:
		return nil, tg.ErrLocationInvalid
	}
}

func (c *FilesCore) downloadResolvedFile(location tg.InputFileLocationClazz, offset int64, limit int32) (*tg.UploadFile, error) {
	resolved, err := c.resolveFileLocation(location)
	if err != nil {
		return nil, err
	}
	uploadFile, err := c.svcCtx.Repo.DfsClient.DfsGetFileByReadLease(c.ctx, &dfs.TLDfsGetFileByReadLease{
		ReadLease: resolved.ReadLease,
		Offset:    offset,
		Limit:     limit,
	})
	if err != nil {
		return nil, err
	}
	return normalizeUploadFile(uploadFile)
}

func (c *FilesCore) resolveFileLocation(location tg.InputFileLocationClazz) (*mediapb.MediaResolvedFileObject, error) {
	viewerID := int64(0)
	if c.MD != nil {
		viewerID = c.MD.UserId
	}
	resolved, err := c.svcCtx.Repo.MediaClient.MediaResolveFileLocation(c.ctx, &mediapb.TLMediaResolveFileLocation{
		Location: location,
		ViewerId: viewerID,
	})
	if err != nil {
		return nil, mapUploadGetFileResolveError(err)
	}
	if resolved == nil || len(resolved.ReadLease) == 0 {
		return nil, tg.ErrLocationInvalid
	}
	return resolved, nil
}

func (c *FilesCore) downloadLegacyFile(location tg.InputFileLocationClazz, offset int64, limit int32) (*tg.UploadFile, error) {
	uploadFile, err := c.svcCtx.Repo.DfsClient.DfsDownloadFile(c.ctx, &dfs.TLDfsDownloadFile{
		Location: location,
		Offset:   offset,
		Limit:    limit,
	})
	if err != nil {
		return nil, err
	}
	return normalizeUploadFile(uploadFile)
}

func normalizeUploadFile(uploadFile *tg.UploadFile) (*tg.UploadFile, error) {
	if uploadFile == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	file, ok := uploadFile.ToUploadFile()
	if !ok || file.Type == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if file.Type.StorageFileTypeClazzName() == tg.ClazzName_storage_fileUnknown {
		file.Type = tg.MakeTLStorageFilePartial(&tg.TLStorageFilePartial{})
	}

	return uploadFile, nil
}

func mapUploadGetFileResolveError(err error) error {
	switch {
	case errors.Is(err, mediapb.ErrFileReferenceEmpty):
		return tg.ErrFileReferenceEmpty
	case errors.Is(err, mediapb.ErrFileReferenceExpired):
		return tg.ErrFileReferenceExpired
	case errors.Is(err, mediapb.ErrFileReferenceInvalid):
		return tg.ErrFileReferenceInvalid
	case errors.Is(err, mediapb.ErrFileLocationInvalid),
		errors.Is(err, mediapb.ErrDocumentNotFound),
		errors.Is(err, mediapb.ErrPhotoNotFound),
		errors.Is(err, mediapb.ErrMediaInvalidArgument):
		return tg.ErrLocationInvalid
	default:
		return err
	}
}
