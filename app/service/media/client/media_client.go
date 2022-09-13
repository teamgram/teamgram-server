/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package media_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/media/media"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type MediaClient interface {
	MediaUploadPhotoFile(ctx context.Context, in *media.TLMediaUploadPhotoFile) (*mtproto.Photo, error)
	MediaUploadProfilePhotoFile(ctx context.Context, in *media.TLMediaUploadProfilePhotoFile) (*mtproto.Photo, error)
	MediaGetPhoto(ctx context.Context, in *media.TLMediaGetPhoto) (*mtproto.Photo, error)
	MediaGetPhotoSizeList(ctx context.Context, in *media.TLMediaGetPhotoSizeList) (*media.PhotoSizeList, error)
	MediaGetPhotoSizeListList(ctx context.Context, in *media.TLMediaGetPhotoSizeListList) (*media.Vector_PhotoSizeList, error)
	MediaGetVideoSizeList(ctx context.Context, in *media.TLMediaGetVideoSizeList) (*media.VideoSizeList, error)
	MediaUploadedDocumentMedia(ctx context.Context, in *media.TLMediaUploadedDocumentMedia) (*mtproto.MessageMedia, error)
	MediaGetDocument(ctx context.Context, in *media.TLMediaGetDocument) (*mtproto.Document, error)
	MediaGetDocumentList(ctx context.Context, in *media.TLMediaGetDocumentList) (*media.Vector_Document, error)
	MediaUploadEncryptedFile(ctx context.Context, in *media.TLMediaUploadEncryptedFile) (*mtproto.EncryptedFile, error)
	MediaGetEncryptedFile(ctx context.Context, in *media.TLMediaGetEncryptedFile) (*mtproto.EncryptedFile, error)
	MediaUploadWallPaperFile(ctx context.Context, in *media.TLMediaUploadWallPaperFile) (*mtproto.Document, error)
	MediaUploadThemeFile(ctx context.Context, in *media.TLMediaUploadThemeFile) (*mtproto.Document, error)
	MediaUploadStickerFile(ctx context.Context, in *media.TLMediaUploadStickerFile) (*mtproto.Document, error)
	MediaUploadRingtoneFile(ctx context.Context, in *media.TLMediaUploadRingtoneFile) (*mtproto.Document, error)
}

type defaultMediaClient struct {
	cli zrpc.Client
}

func NewMediaClient(cli zrpc.Client) MediaClient {
	return &defaultMediaClient{
		cli: cli,
	}
}

// MediaUploadPhotoFile
// media.uploadPhotoFile flags:# owner_id:long file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = Photo;
func (m *defaultMediaClient) MediaUploadPhotoFile(ctx context.Context, in *media.TLMediaUploadPhotoFile) (*mtproto.Photo, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaUploadPhotoFile(ctx, in)
}

// MediaUploadProfilePhotoFile
// media.uploadProfilePhotoFile flags:# owner_id:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = Photo;
func (m *defaultMediaClient) MediaUploadProfilePhotoFile(ctx context.Context, in *media.TLMediaUploadProfilePhotoFile) (*mtproto.Photo, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaUploadProfilePhotoFile(ctx, in)
}

// MediaGetPhoto
// media.getPhoto photo_id:long = Photo;
func (m *defaultMediaClient) MediaGetPhoto(ctx context.Context, in *media.TLMediaGetPhoto) (*mtproto.Photo, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaGetPhoto(ctx, in)
}

// MediaGetPhotoSizeList
// media.getPhotoSizeList size_id:long = PhotoSizeList;
func (m *defaultMediaClient) MediaGetPhotoSizeList(ctx context.Context, in *media.TLMediaGetPhotoSizeList) (*media.PhotoSizeList, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaGetPhotoSizeList(ctx, in)
}

// MediaGetPhotoSizeListList
// media.getPhotoSizeListList id_list:Vector<long> = Vector<PhotoSizeList>;
func (m *defaultMediaClient) MediaGetPhotoSizeListList(ctx context.Context, in *media.TLMediaGetPhotoSizeListList) (*media.Vector_PhotoSizeList, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaGetPhotoSizeListList(ctx, in)
}

// MediaGetVideoSizeList
// media.getVideoSizeList size_id:long = VideoSizeList;
func (m *defaultMediaClient) MediaGetVideoSizeList(ctx context.Context, in *media.TLMediaGetVideoSizeList) (*media.VideoSizeList, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaGetVideoSizeList(ctx, in)
}

// MediaUploadedDocumentMedia
// media.uploadedDocumentMedia owner_id:long media:InputMedia = MessageMedia;
func (m *defaultMediaClient) MediaUploadedDocumentMedia(ctx context.Context, in *media.TLMediaUploadedDocumentMedia) (*mtproto.MessageMedia, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaUploadedDocumentMedia(ctx, in)
}

// MediaGetDocument
// media.getDocument id:long = Document;
func (m *defaultMediaClient) MediaGetDocument(ctx context.Context, in *media.TLMediaGetDocument) (*mtproto.Document, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaGetDocument(ctx, in)
}

// MediaGetDocumentList
// media.getDocumentList id_list:Vector<long> = Vector<Document>;
func (m *defaultMediaClient) MediaGetDocumentList(ctx context.Context, in *media.TLMediaGetDocumentList) (*media.Vector_Document, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaGetDocumentList(ctx, in)
}

// MediaUploadEncryptedFile
// media.uploadEncryptedFile owner_id:long file:InputEncryptedFile = EncryptedFile;
func (m *defaultMediaClient) MediaUploadEncryptedFile(ctx context.Context, in *media.TLMediaUploadEncryptedFile) (*mtproto.EncryptedFile, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaUploadEncryptedFile(ctx, in)
}

// MediaGetEncryptedFile
// media.getEncryptedFile id:long access_hash:long = EncryptedFile;
func (m *defaultMediaClient) MediaGetEncryptedFile(ctx context.Context, in *media.TLMediaGetEncryptedFile) (*mtproto.EncryptedFile, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaGetEncryptedFile(ctx, in)
}

// MediaUploadWallPaperFile
// media.uploadWallPaperFile owner_id:long file:InputFile mime_type:string admin:Bool = Document;
func (m *defaultMediaClient) MediaUploadWallPaperFile(ctx context.Context, in *media.TLMediaUploadWallPaperFile) (*mtproto.Document, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaUploadWallPaperFile(ctx, in)
}

// MediaUploadThemeFile
// media.uploadThemeFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
func (m *defaultMediaClient) MediaUploadThemeFile(ctx context.Context, in *media.TLMediaUploadThemeFile) (*mtproto.Document, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaUploadThemeFile(ctx, in)
}

// MediaUploadStickerFile
// media.uploadStickerFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string document_attribute_sticker:DocumentAttribute = Document;
func (m *defaultMediaClient) MediaUploadStickerFile(ctx context.Context, in *media.TLMediaUploadStickerFile) (*mtproto.Document, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaUploadStickerFile(ctx, in)
}

// MediaUploadRingtoneFile
// media.uploadRingtoneFile flags:# owner_id:long file:InputFile mime_type:string file_name:string = Document;
func (m *defaultMediaClient) MediaUploadRingtoneFile(ctx context.Context, in *media.TLMediaUploadRingtoneFile) (*mtproto.Document, error) {
	client := media.NewRPCMediaClient(m.cli.Conn())
	return client.MediaUploadRingtoneFile(ctx, in)
}
