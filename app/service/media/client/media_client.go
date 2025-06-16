/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mediaclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media/mediaservice"

	"github.com/cloudwego/kitex/client"
)

type MediaClient interface {
	MediaUploadPhotoFile(ctx context.Context, in *media.TLMediaUploadPhotoFile) (*tg.Photo, error)
	MediaUploadProfilePhotoFile(ctx context.Context, in *media.TLMediaUploadProfilePhotoFile) (*tg.Photo, error)
	MediaGetPhoto(ctx context.Context, in *media.TLMediaGetPhoto) (*tg.Photo, error)
	MediaGetPhotoSizeList(ctx context.Context, in *media.TLMediaGetPhotoSizeList) (*media.PhotoSizeList, error)
	MediaGetPhotoSizeListList(ctx context.Context, in *media.TLMediaGetPhotoSizeListList) (*media.VectorPhotoSizeList, error)
	MediaGetVideoSizeList(ctx context.Context, in *media.TLMediaGetVideoSizeList) (*media.VideoSizeList, error)
	MediaUploadedDocumentMedia(ctx context.Context, in *media.TLMediaUploadedDocumentMedia) (*tg.MessageMedia, error)
	MediaGetDocument(ctx context.Context, in *media.TLMediaGetDocument) (*tg.Document, error)
	MediaGetDocumentList(ctx context.Context, in *media.TLMediaGetDocumentList) (*media.VectorDocument, error)
	MediaUploadEncryptedFile(ctx context.Context, in *media.TLMediaUploadEncryptedFile) (*tg.EncryptedFile, error)
	MediaGetEncryptedFile(ctx context.Context, in *media.TLMediaGetEncryptedFile) (*tg.EncryptedFile, error)
	MediaUploadWallPaperFile(ctx context.Context, in *media.TLMediaUploadWallPaperFile) (*tg.Document, error)
	MediaUploadThemeFile(ctx context.Context, in *media.TLMediaUploadThemeFile) (*tg.Document, error)
	MediaUploadStickerFile(ctx context.Context, in *media.TLMediaUploadStickerFile) (*tg.Document, error)
	MediaUploadRingtoneFile(ctx context.Context, in *media.TLMediaUploadRingtoneFile) (*tg.Document, error)
	MediaUploadedProfilePhoto(ctx context.Context, in *media.TLMediaUploadedProfilePhoto) (*tg.Photo, error)
}

type defaultMediaClient struct {
	cli client.Client
}

func NewMediaClient(cli client.Client) MediaClient {
	return &defaultMediaClient{
		cli: cli,
	}
}

// MediaUploadPhotoFile
// media.uploadPhotoFile flags:# owner_id:long file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = Photo;
func (m *defaultMediaClient) MediaUploadPhotoFile(ctx context.Context, in *media.TLMediaUploadPhotoFile) (*tg.Photo, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaUploadPhotoFile(ctx, in)
}

// MediaUploadProfilePhotoFile
// media.uploadProfilePhotoFile flags:# owner_id:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = Photo;
func (m *defaultMediaClient) MediaUploadProfilePhotoFile(ctx context.Context, in *media.TLMediaUploadProfilePhotoFile) (*tg.Photo, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaUploadProfilePhotoFile(ctx, in)
}

// MediaGetPhoto
// media.getPhoto photo_id:long = Photo;
func (m *defaultMediaClient) MediaGetPhoto(ctx context.Context, in *media.TLMediaGetPhoto) (*tg.Photo, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaGetPhoto(ctx, in)
}

// MediaGetPhotoSizeList
// media.getPhotoSizeList size_id:long = PhotoSizeList;
func (m *defaultMediaClient) MediaGetPhotoSizeList(ctx context.Context, in *media.TLMediaGetPhotoSizeList) (*media.PhotoSizeList, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaGetPhotoSizeList(ctx, in)
}

// MediaGetPhotoSizeListList
// media.getPhotoSizeListList id_list:Vector<long> = Vector<PhotoSizeList>;
func (m *defaultMediaClient) MediaGetPhotoSizeListList(ctx context.Context, in *media.TLMediaGetPhotoSizeListList) (*media.VectorPhotoSizeList, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaGetPhotoSizeListList(ctx, in)
}

// MediaGetVideoSizeList
// media.getVideoSizeList size_id:long = VideoSizeList;
func (m *defaultMediaClient) MediaGetVideoSizeList(ctx context.Context, in *media.TLMediaGetVideoSizeList) (*media.VideoSizeList, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaGetVideoSizeList(ctx, in)
}

// MediaUploadedDocumentMedia
// media.uploadedDocumentMedia owner_id:long media:InputMedia = MessageMedia;
func (m *defaultMediaClient) MediaUploadedDocumentMedia(ctx context.Context, in *media.TLMediaUploadedDocumentMedia) (*tg.MessageMedia, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaUploadedDocumentMedia(ctx, in)
}

// MediaGetDocument
// media.getDocument id:long = Document;
func (m *defaultMediaClient) MediaGetDocument(ctx context.Context, in *media.TLMediaGetDocument) (*tg.Document, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaGetDocument(ctx, in)
}

// MediaGetDocumentList
// media.getDocumentList id_list:Vector<long> = Vector<Document>;
func (m *defaultMediaClient) MediaGetDocumentList(ctx context.Context, in *media.TLMediaGetDocumentList) (*media.VectorDocument, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaGetDocumentList(ctx, in)
}

// MediaUploadEncryptedFile
// media.uploadEncryptedFile owner_id:long file:InputEncryptedFile = EncryptedFile;
func (m *defaultMediaClient) MediaUploadEncryptedFile(ctx context.Context, in *media.TLMediaUploadEncryptedFile) (*tg.EncryptedFile, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaUploadEncryptedFile(ctx, in)
}

// MediaGetEncryptedFile
// media.getEncryptedFile id:long access_hash:long = EncryptedFile;
func (m *defaultMediaClient) MediaGetEncryptedFile(ctx context.Context, in *media.TLMediaGetEncryptedFile) (*tg.EncryptedFile, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaGetEncryptedFile(ctx, in)
}

// MediaUploadWallPaperFile
// media.uploadWallPaperFile owner_id:long file:InputFile mime_type:string admin:Bool = Document;
func (m *defaultMediaClient) MediaUploadWallPaperFile(ctx context.Context, in *media.TLMediaUploadWallPaperFile) (*tg.Document, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaUploadWallPaperFile(ctx, in)
}

// MediaUploadThemeFile
// media.uploadThemeFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
func (m *defaultMediaClient) MediaUploadThemeFile(ctx context.Context, in *media.TLMediaUploadThemeFile) (*tg.Document, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaUploadThemeFile(ctx, in)
}

// MediaUploadStickerFile
// media.uploadStickerFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string document_attribute_sticker:DocumentAttribute = Document;
func (m *defaultMediaClient) MediaUploadStickerFile(ctx context.Context, in *media.TLMediaUploadStickerFile) (*tg.Document, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaUploadStickerFile(ctx, in)
}

// MediaUploadRingtoneFile
// media.uploadRingtoneFile flags:# owner_id:long file:InputFile mime_type:string file_name:string = Document;
func (m *defaultMediaClient) MediaUploadRingtoneFile(ctx context.Context, in *media.TLMediaUploadRingtoneFile) (*tg.Document, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaUploadRingtoneFile(ctx, in)
}

// MediaUploadedProfilePhoto
// media.uploadedProfilePhoto owner_id:long photo_id:long = Photo;
func (m *defaultMediaClient) MediaUploadedProfilePhoto(ctx context.Context, in *media.TLMediaUploadedProfilePhoto) (*tg.Photo, error) {
	cli := mediaservice.NewRPCMediaClient(m.cli)
	return cli.MediaUploadedProfilePhoto(ctx, in)
}
