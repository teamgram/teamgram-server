/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/core"
)

// DfsWriteFilePartData
// dfs.writeFilePartData flags:# creator:long file_id:long file_part:int bytes:bytes big:flags.0?true file_total_parts:flags.1?int = Bool;
func (s *Service) DfsWriteFilePartData(ctx context.Context, request *dfs.TLDfsWriteFilePartData) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.writeFilePartData - metadata: {}, request: %v", request)

	r, err := c.DfsWriteFilePartData(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// DfsUploadPhotoFileV2
// dfs.uploadPhotoFileV2 creator:long file:InputFile = Photo;
func (s *Service) DfsUploadPhotoFileV2(ctx context.Context, request *dfs.TLDfsUploadPhotoFileV2) (*tg.Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadPhotoFileV2 - metadata: {}, request: %v", request)

	r, err := c.DfsUploadPhotoFileV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// DfsUploadProfilePhotoFileV2
// dfs.uploadProfilePhotoFileV2 flags:# creator:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = Photo;
func (s *Service) DfsUploadProfilePhotoFileV2(ctx context.Context, request *dfs.TLDfsUploadProfilePhotoFileV2) (*tg.Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadProfilePhotoFileV2 - metadata: {}, request: %v", request)

	r, err := c.DfsUploadProfilePhotoFileV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// DfsUploadEncryptedFileV2
// dfs.uploadEncryptedFileV2 creator:long file:InputEncryptedFile = EncryptedFile;
func (s *Service) DfsUploadEncryptedFileV2(ctx context.Context, request *dfs.TLDfsUploadEncryptedFileV2) (*tg.EncryptedFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadEncryptedFileV2 - metadata: {}, request: %v", request)

	r, err := c.DfsUploadEncryptedFileV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// DfsDownloadFile
// dfs.downloadFile location:InputFileLocation offset:long limit:int = upload.File;
func (s *Service) DfsDownloadFile(ctx context.Context, request *dfs.TLDfsDownloadFile) (*tg.UploadFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.downloadFile - metadata: {}, request: %v", request)

	r, err := c.DfsDownloadFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// DfsUploadDocumentFileV2
// dfs.uploadDocumentFileV2 creator:long media:InputMedia = Document;
func (s *Service) DfsUploadDocumentFileV2(ctx context.Context, request *dfs.TLDfsUploadDocumentFileV2) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadDocumentFileV2 - metadata: {}, request: %v", request)

	r, err := c.DfsUploadDocumentFileV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// DfsUploadGifDocumentMedia
// dfs.uploadGifDocumentMedia creator:long media:InputMedia = Document;
func (s *Service) DfsUploadGifDocumentMedia(ctx context.Context, request *dfs.TLDfsUploadGifDocumentMedia) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadGifDocumentMedia - metadata: {}, request: %v", request)

	r, err := c.DfsUploadGifDocumentMedia(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// DfsUploadMp4DocumentMedia
// dfs.uploadMp4DocumentMedia creator:long media:InputMedia = Document;
func (s *Service) DfsUploadMp4DocumentMedia(ctx context.Context, request *dfs.TLDfsUploadMp4DocumentMedia) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadMp4DocumentMedia - metadata: {}, request: %v", request)

	r, err := c.DfsUploadMp4DocumentMedia(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// DfsUploadWallPaperFile
// dfs.uploadWallPaperFile creator:long file:InputFile mime_type:string admin:Bool = Document;
func (s *Service) DfsUploadWallPaperFile(ctx context.Context, request *dfs.TLDfsUploadWallPaperFile) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadWallPaperFile - metadata: {}, request: %v", request)

	r, err := c.DfsUploadWallPaperFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// DfsUploadThemeFile
// dfs.uploadThemeFile flags:# creator:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
func (s *Service) DfsUploadThemeFile(ctx context.Context, request *dfs.TLDfsUploadThemeFile) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadThemeFile - metadata: {}, request: %v", request)

	r, err := c.DfsUploadThemeFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// DfsUploadRingtoneFile
// dfs.uploadRingtoneFile creator:long file:InputFile mime_type:string file_name:string = Document;
func (s *Service) DfsUploadRingtoneFile(ctx context.Context, request *dfs.TLDfsUploadRingtoneFile) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadRingtoneFile - metadata: {}, request: %v", request)

	r, err := c.DfsUploadRingtoneFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// DfsUploadedProfilePhoto
// dfs.uploadedProfilePhoto creator:long photo_id:long = Photo;
func (s *Service) DfsUploadedProfilePhoto(ctx context.Context, request *dfs.TLDfsUploadedProfilePhoto) (*tg.Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dfs.uploadedProfilePhoto - metadata: {}, request: %v", request)

	r, err := c.DfsUploadedProfilePhoto(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}
