/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/core"
)

// DfsWriteFilePartData
// dfs.writeFilePartData flags:# creator:long file_id:long bytes:bytes big:flags.0?true file_total_parts:flags.1?int = Bool;
func (s *Service) DfsWriteFilePartData(ctx context.Context, request *dfs.TLDfsWriteFilePartData) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dfs.writeFilePartData - request: {creator: %d, file_id: %d, file_part: %d, file_total_parts: %d, bytes: %d}",
		request.Creator,
		request.FileId,
		request.FilePart,
		request.FileTotalParts,
		len(request.Bytes))

	r, err := c.DfsWriteFilePartData(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dfs.writeFilePartData - reply: %s", r.DebugString())
	return r, err
}

// DfsUploadPhotoFileV2
// dfs.uploadPhotoFileV2 creator:long file:InputFile = Photo;
func (s *Service) DfsUploadPhotoFileV2(ctx context.Context, request *dfs.TLDfsUploadPhotoFileV2) (*mtproto.Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dfs.uploadPhotoFileV2 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DfsUploadPhotoFileV2(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dfs.uploadPhotoFileV2 - reply: %s", r.DebugString())
	return r, err
}

// DfsUploadProfilePhotoFileV2
// dfs.uploadProfilePhotoFileV2 flags:# creator:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = Photo;
func (s *Service) DfsUploadProfilePhotoFileV2(ctx context.Context, request *dfs.TLDfsUploadProfilePhotoFileV2) (*mtproto.Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dfs.uploadProfilePhotoFileV2 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DfsUploadProfilePhotoFileV2(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dfs.uploadProfilePhotoFileV2 - reply: %s", r.DebugString())
	return r, err
}

// DfsUploadEncryptedFileV2
// dfs.uploadEncryptedFileV2 creator:long file:InputEncryptedFile = EncryptedFile;
func (s *Service) DfsUploadEncryptedFileV2(ctx context.Context, request *dfs.TLDfsUploadEncryptedFileV2) (*mtproto.EncryptedFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dfs.uploadEncryptedFileV2 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DfsUploadEncryptedFileV2(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dfs.uploadEncryptedFileV2 - reply: %s", r.DebugString())
	return r, err
}

// DfsDownloadFile
// dfs.downloadFile location:InputFileLocation offset:int limit:int = upload.File;
func (s *Service) DfsDownloadFile(ctx context.Context, request *dfs.TLDfsDownloadFile) (*mtproto.Upload_File, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dfs.downloadFile - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DfsDownloadFile(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dfs.downloadFile - reply: {type: %s, mtime: %d, bytes: %d}",
		r.Type.DebugString(),
		r.Mtime,
		len(r.Bytes))

	return r, err
}

// DfsUploadDocumentFileV2
// dfs.uploadDocumentFileV2 creator:long media:InputMedia = Document;
func (s *Service) DfsUploadDocumentFileV2(ctx context.Context, request *dfs.TLDfsUploadDocumentFileV2) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dfs.uploadDocumentFileV2 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DfsUploadDocumentFileV2(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dfs.uploadDocumentFileV2 - reply: %s", r.DebugString())
	return r, err
}

// DfsUploadGifDocumentMedia
// dfs.uploadGifDocumentMedia creator:long media:InputMedia = Document;
func (s *Service) DfsUploadGifDocumentMedia(ctx context.Context, request *dfs.TLDfsUploadGifDocumentMedia) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dfs.uploadGifDocumentMedia - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DfsUploadGifDocumentMedia(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dfs.uploadGifDocumentMedia - reply: %s", r.DebugString())
	return r, err
}

// DfsUploadMp4DocumentMedia
// dfs.uploadMp4DocumentMedia creator:long media:InputMedia = Document;
func (s *Service) DfsUploadMp4DocumentMedia(ctx context.Context, request *dfs.TLDfsUploadMp4DocumentMedia) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dfs.uploadMp4DocumentMedia - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DfsUploadMp4DocumentMedia(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dfs.uploadMp4DocumentMedia - reply: %s", r.DebugString())
	return r, err
}

// DfsUploadWallPaperFile
// dfs.uploadWallPaperFile creator:long file:InputFile mime_type:string admin:Bool = Document;
func (s *Service) DfsUploadWallPaperFile(ctx context.Context, request *dfs.TLDfsUploadWallPaperFile) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dfs.uploadWallPaperFile - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DfsUploadWallPaperFile(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dfs.uploadWallPaperFile - reply: %s", r.DebugString())
	return r, err
}

// DfsUploadThemeFile
// dfs.uploadThemeFile flags:# creator:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
func (s *Service) DfsUploadThemeFile(ctx context.Context, request *dfs.TLDfsUploadThemeFile) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dfs.uploadThemeFile - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DfsUploadThemeFile(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dfs.uploadThemeFile - reply: %s", r.DebugString())
	return r, err
}
