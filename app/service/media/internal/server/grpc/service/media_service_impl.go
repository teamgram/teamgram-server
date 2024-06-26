/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/media/internal/core"
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// MediaUploadPhotoFile
// media.uploadPhotoFile flags:# owner_id:long file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = Photo;
func (s *Service) MediaUploadPhotoFile(ctx context.Context, request *media.TLMediaUploadPhotoFile) (*mtproto.Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.uploadPhotoFile - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaUploadPhotoFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.uploadPhotoFile - reply: %s", r)
	return r, err
}

// MediaUploadProfilePhotoFile
// media.uploadProfilePhotoFile flags:# owner_id:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = Photo;
func (s *Service) MediaUploadProfilePhotoFile(ctx context.Context, request *media.TLMediaUploadProfilePhotoFile) (*mtproto.Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.uploadProfilePhotoFile - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaUploadProfilePhotoFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.uploadProfilePhotoFile - reply: %s", r)
	return r, err
}

// MediaGetPhoto
// media.getPhoto photo_id:long = Photo;
func (s *Service) MediaGetPhoto(ctx context.Context, request *media.TLMediaGetPhoto) (*mtproto.Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.getPhoto - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaGetPhoto(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.getPhoto - reply: %s", r)
	return r, err
}

// MediaGetPhotoSizeList
// media.getPhotoSizeList size_id:long = PhotoSizeList;
func (s *Service) MediaGetPhotoSizeList(ctx context.Context, request *media.TLMediaGetPhotoSizeList) (*media.PhotoSizeList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.getPhotoSizeList - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaGetPhotoSizeList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.getPhotoSizeList - reply: %s", r)
	return r, err
}

// MediaGetPhotoSizeListList
// media.getPhotoSizeListList id_list:Vector<long> = Vector<PhotoSizeList>;
func (s *Service) MediaGetPhotoSizeListList(ctx context.Context, request *media.TLMediaGetPhotoSizeListList) (*media.Vector_PhotoSizeList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.getPhotoSizeListList - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaGetPhotoSizeListList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.getPhotoSizeListList - reply: %s", r)
	return r, err
}

// MediaGetVideoSizeList
// media.getVideoSizeList size_id:long = VideoSizeList;
func (s *Service) MediaGetVideoSizeList(ctx context.Context, request *media.TLMediaGetVideoSizeList) (*media.VideoSizeList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.getVideoSizeList - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaGetVideoSizeList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.getVideoSizeList - reply: %s", r)
	return r, nil
}

// MediaUploadedDocumentMedia
// media.uploadedDocumentMedia owner_id:long media:InputMedia = MessageMedia;
func (s *Service) MediaUploadedDocumentMedia(ctx context.Context, request *media.TLMediaUploadedDocumentMedia) (*mtproto.MessageMedia, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.uploadedDocumentMedia - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaUploadedDocumentMedia(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.uploadedDocumentMedia - reply: %s", r)
	return r, err
}

// MediaGetDocument
// media.getDocument id:long = Document;
func (s *Service) MediaGetDocument(ctx context.Context, request *media.TLMediaGetDocument) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.getDocument - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaGetDocument(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.getDocument - reply: %s", r)
	return r, err
}

// MediaGetDocumentList
// media.getDocumentList id_list:Vector<long> = Vector<Document>;
func (s *Service) MediaGetDocumentList(ctx context.Context, request *media.TLMediaGetDocumentList) (*media.Vector_Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.getDocumentList - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaGetDocumentList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.getDocumentList - reply: %s", r)
	return r, err
}

// MediaUploadEncryptedFile
// media.uploadEncryptedFile owner_id:long file:InputEncryptedFile = EncryptedFile;
func (s *Service) MediaUploadEncryptedFile(ctx context.Context, request *media.TLMediaUploadEncryptedFile) (*mtproto.EncryptedFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.uploadEncryptedFile - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaUploadEncryptedFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.uploadEncryptedFile - reply: %s", r)
	return r, err
}

// MediaGetEncryptedFile
// media.getEncryptedFile id:long access_hash:long = EncryptedFile;
func (s *Service) MediaGetEncryptedFile(ctx context.Context, request *media.TLMediaGetEncryptedFile) (*mtproto.EncryptedFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.getEncryptedFile - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaGetEncryptedFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.getEncryptedFile - reply: %s", r)
	return r, err
}

// MediaUploadWallPaperFile
// media.uploadWallPaperFile owner_id:long file:InputFile mime_type:string admin:Bool = Document;
func (s *Service) MediaUploadWallPaperFile(ctx context.Context, request *media.TLMediaUploadWallPaperFile) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.uploadWallPaperFile - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaUploadWallPaperFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.uploadWallPaperFile - reply: %s", r)
	return r, err
}

// MediaUploadThemeFile
// media.uploadThemeFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
func (s *Service) MediaUploadThemeFile(ctx context.Context, request *media.TLMediaUploadThemeFile) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.uploadThemeFile - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaUploadThemeFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.uploadThemeFile - reply: %s", r)
	return r, err
}

// MediaUploadStickerFile
// media.uploadStickerFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string document_attribute_sticker:DocumentAttribute = Document;
func (s *Service) MediaUploadStickerFile(ctx context.Context, request *media.TLMediaUploadStickerFile) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.uploadStickerFile - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaUploadStickerFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.uploadStickerFile - reply: %s", r)
	return r, err
}

// MediaUploadRingtoneFile
// media.uploadRingtoneFile flags:# owner_id:long file:InputFile mime_type:string file_name:string = Document;
func (s *Service) MediaUploadRingtoneFile(ctx context.Context, request *media.TLMediaUploadRingtoneFile) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("media.uploadRingtoneFile - metadata: %s, request: %s", c.MD, request)

	r, err := c.MediaUploadRingtoneFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("media.uploadRingtoneFile - reply: %s", r)
	return r, err
}
