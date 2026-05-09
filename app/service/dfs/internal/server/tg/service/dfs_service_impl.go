/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/core"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// DfsCommitUpload
// dfs.commitUpload upload_session_id:string owner_id:long file:InputFile purpose:string = FileFinalizedObject;
func (s *Service) DfsCommitUpload(ctx context.Context, request *dfs.TLDfsCommitUpload) (*dfs.FileFinalizedObject, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.commitUpload - request: %s", request)

	r, err := c.DfsCommitUpload(request)
	if err != nil {
		c.Logger.Errorf("dfs.commitUpload - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.commitUpload - reply: %s", r)
	return r, err
}

// DfsPutFile
// dfs.putFile owner_id:long purpose:string file_name:string mime_type:string bytes:bytes = FileFinalizedObject;
func (s *Service) DfsPutFile(ctx context.Context, request *dfs.TLDfsPutFile) (*dfs.FileFinalizedObject, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.putFile - request: %s", request)

	r, err := c.DfsPutFile(request)
	if err != nil {
		c.Logger.Errorf("dfs.putFile - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.putFile - reply: %s", r)
	return r, err
}

// DfsGetFileByReadLease
// dfs.getFileByReadLease read_lease:bytes offset:long limit:int = upload.File;
func (s *Service) DfsGetFileByReadLease(ctx context.Context, request *dfs.TLDfsGetFileByReadLease) (*tg.UploadFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.getFileByReadLease - request: %s", request)

	r, err := c.DfsGetFileByReadLease(request)
	if err != nil {
		c.Logger.Errorf("dfs.getFileByReadLease - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.getFileByReadLease - reply: %s", r)
	return r, err
}

// DfsGetFileHashesByReadLease
// dfs.getFileHashesByReadLease read_lease:bytes offset:long limit:int = Vector<FileHash>;
func (s *Service) DfsGetFileHashesByReadLease(ctx context.Context, request *dfs.TLDfsGetFileHashesByReadLease) (*dfs.VectorFileHash, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.getFileHashesByReadLease - request: %s", request)

	r, err := c.DfsGetFileHashesByReadLease(request)
	if err != nil {
		c.Logger.Errorf("dfs.getFileHashesByReadLease - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.getFileHashesByReadLease - reply: %s", r)
	return r, err
}

// DfsWriteFilePartData
// dfs.writeFilePartData flags:# creator:long file_id:long file_part:int bytes:bytes big:flags.0?true file_total_parts:flags.1?int = Bool;
func (s *Service) DfsWriteFilePartData(ctx context.Context, request *dfs.TLDfsWriteFilePartData) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.writeFilePartData - request: {creator: %d, file_id: %d, file_part: %d, bytes_len: %d}",
		request.Creator,
		request.FileId,
		request.FilePart,
		len(request.Bytes))

	r, err := c.DfsWriteFilePartData(request)
	if err != nil {
		c.Logger.Errorf("dfs.writeFilePartData - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.writeFilePartData - reply: %s", r)
	return r, err
}

// DfsUploadPhotoFileV2
// dfs.uploadPhotoFileV2 creator:long file:InputFile = Photo;
func (s *Service) DfsUploadPhotoFileV2(ctx context.Context, request *dfs.TLDfsUploadPhotoFileV2) (*tg.Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadPhotoFileV2 - request: %s", request)

	r, err := c.DfsUploadPhotoFileV2(request)
	if err != nil {
		c.Logger.Errorf("dfs.uploadPhotoFileV2 - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.uploadPhotoFileV2 - reply: %s", r)
	return r, err
}

// DfsUploadProfilePhotoFileV2
// dfs.uploadProfilePhotoFileV2 flags:# creator:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.4?VideoSize = Photo;
func (s *Service) DfsUploadProfilePhotoFileV2(ctx context.Context, request *dfs.TLDfsUploadProfilePhotoFileV2) (*tg.Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadProfilePhotoFileV2 - request: %s", request)

	r, err := c.DfsUploadProfilePhotoFileV2(request)
	if err != nil {
		c.Logger.Errorf("dfs.uploadProfilePhotoFileV2 - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.uploadProfilePhotoFileV2 - reply: %s", r)
	return r, err
}

// DfsUploadEncryptedFileV2
// dfs.uploadEncryptedFileV2 creator:long file:InputEncryptedFile = EncryptedFile;
func (s *Service) DfsUploadEncryptedFileV2(ctx context.Context, request *dfs.TLDfsUploadEncryptedFileV2) (*tg.EncryptedFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadEncryptedFileV2 - request: %s", request)

	r, err := c.DfsUploadEncryptedFileV2(request)
	if err != nil {
		c.Logger.Errorf("dfs.uploadEncryptedFileV2 - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.uploadEncryptedFileV2 - reply: %s", r)
	return r, err
}

// DfsDownloadFile
// dfs.downloadFile location:InputFileLocation offset:long limit:int = upload.File;
func (s *Service) DfsDownloadFile(ctx context.Context, request *dfs.TLDfsDownloadFile) (*tg.UploadFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.downloadFile - request: %s", request)

	r, err := c.DfsDownloadFile(request)
	if err != nil {
		c.Logger.Errorf("dfs.downloadFile - error: request: %s, err: %v", request, err)
		return nil, err
	}

	switch r2 := r.Clazz.(type) {
	case *tg.TLUploadFile:
		c.Logger.Debugf("upload.getFile - reply: {type: %v, mime: %d, len_bytes: %d}",
			r2.Type,
			r2.Mtime,
			len(r2.Bytes))
	default:
		c.Logger.Debugf("upload.getFile - reply: %s", r)
	}

	return r, err
}

// DfsUploadDocumentFileV2
// dfs.uploadDocumentFileV2 creator:long media:InputMedia = Document;
func (s *Service) DfsUploadDocumentFileV2(ctx context.Context, request *dfs.TLDfsUploadDocumentFileV2) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadDocumentFileV2 - request: %s", request)

	r, err := c.DfsUploadDocumentFileV2(request)
	if err != nil {
		c.Logger.Errorf("dfs.uploadDocumentFileV2 - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.uploadDocumentFileV2 - reply: %s", r)
	return r, err
}

// DfsUploadGifDocumentMedia
// dfs.uploadGifDocumentMedia creator:long media:InputMedia = Document;
func (s *Service) DfsUploadGifDocumentMedia(ctx context.Context, request *dfs.TLDfsUploadGifDocumentMedia) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadGifDocumentMedia - request: %s", request)

	r, err := c.DfsUploadGifDocumentMedia(request)
	if err != nil {
		c.Logger.Errorf("dfs.uploadGifDocumentMedia - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.uploadGifDocumentMedia - reply: %s", r)
	return r, err
}

// DfsUploadMp4DocumentMedia
// dfs.uploadMp4DocumentMedia creator:long media:InputMedia = Document;
func (s *Service) DfsUploadMp4DocumentMedia(ctx context.Context, request *dfs.TLDfsUploadMp4DocumentMedia) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadMp4DocumentMedia - request: %s", request)

	r, err := c.DfsUploadMp4DocumentMedia(request)
	if err != nil {
		c.Logger.Errorf("dfs.uploadMp4DocumentMedia - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.uploadMp4DocumentMedia - reply: %s", r)
	return r, err
}

// DfsUploadWallPaperFile
// dfs.uploadWallPaperFile creator:long file:InputFile mime_type:string admin:Bool = Document;
func (s *Service) DfsUploadWallPaperFile(ctx context.Context, request *dfs.TLDfsUploadWallPaperFile) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadWallPaperFile - request: %s", request)

	r, err := c.DfsUploadWallPaperFile(request)
	if err != nil {
		c.Logger.Errorf("dfs.uploadWallPaperFile - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.uploadWallPaperFile - reply: %s", r)
	return r, err
}

// DfsUploadThemeFile
// dfs.uploadThemeFile flags:# creator:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
func (s *Service) DfsUploadThemeFile(ctx context.Context, request *dfs.TLDfsUploadThemeFile) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadThemeFile - request: %s", request)

	r, err := c.DfsUploadThemeFile(request)
	if err != nil {
		c.Logger.Errorf("dfs.uploadThemeFile - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.uploadThemeFile - reply: %s", r)
	return r, err
}

// DfsUploadRingtoneFile
// dfs.uploadRingtoneFile creator:long file:InputFile mime_type:string file_name:string = Document;
func (s *Service) DfsUploadRingtoneFile(ctx context.Context, request *dfs.TLDfsUploadRingtoneFile) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadRingtoneFile - request: %s", request)

	r, err := c.DfsUploadRingtoneFile(request)
	if err != nil {
		c.Logger.Errorf("dfs.uploadRingtoneFile - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.uploadRingtoneFile - reply: %s", r)
	return r, err
}

// DfsUploadedProfilePhoto
// dfs.uploadedProfilePhoto creator:long photo_id:long = Photo;
func (s *Service) DfsUploadedProfilePhoto(ctx context.Context, request *dfs.TLDfsUploadedProfilePhoto) (*tg.Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadedProfilePhoto - request: %s", request)

	r, err := c.DfsUploadedProfilePhoto(request)
	if err != nil {
		c.Logger.Errorf("dfs.uploadedProfilePhoto - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("dfs.uploadedProfilePhoto - reply: %s", r)
	return r, err
}
