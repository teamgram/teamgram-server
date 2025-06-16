/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dfsclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs/dfsservice"

	"github.com/cloudwego/kitex/client"
)

type DfsClient interface {
	DfsWriteFilePartData(ctx context.Context, in *dfs.TLDfsWriteFilePartData) (*tg.Bool, error)
	DfsUploadPhotoFileV2(ctx context.Context, in *dfs.TLDfsUploadPhotoFileV2) (*tg.Photo, error)
	DfsUploadProfilePhotoFileV2(ctx context.Context, in *dfs.TLDfsUploadProfilePhotoFileV2) (*tg.Photo, error)
	DfsUploadEncryptedFileV2(ctx context.Context, in *dfs.TLDfsUploadEncryptedFileV2) (*tg.EncryptedFile, error)
	DfsDownloadFile(ctx context.Context, in *dfs.TLDfsDownloadFile) (*tg.UploadFile, error)
	DfsUploadDocumentFileV2(ctx context.Context, in *dfs.TLDfsUploadDocumentFileV2) (*tg.Document, error)
	DfsUploadGifDocumentMedia(ctx context.Context, in *dfs.TLDfsUploadGifDocumentMedia) (*tg.Document, error)
	DfsUploadMp4DocumentMedia(ctx context.Context, in *dfs.TLDfsUploadMp4DocumentMedia) (*tg.Document, error)
	DfsUploadWallPaperFile(ctx context.Context, in *dfs.TLDfsUploadWallPaperFile) (*tg.Document, error)
	DfsUploadThemeFile(ctx context.Context, in *dfs.TLDfsUploadThemeFile) (*tg.Document, error)
	DfsUploadRingtoneFile(ctx context.Context, in *dfs.TLDfsUploadRingtoneFile) (*tg.Document, error)
	DfsUploadedProfilePhoto(ctx context.Context, in *dfs.TLDfsUploadedProfilePhoto) (*tg.Photo, error)
}

type defaultDfsClient struct {
	cli client.Client
}

func NewDfsClient(cli client.Client) DfsClient {
	return &defaultDfsClient{
		cli: cli,
	}
}

// DfsWriteFilePartData
// dfs.writeFilePartData flags:# creator:long file_id:long file_part:int bytes:bytes big:flags.0?true file_total_parts:flags.1?int = Bool;
func (m *defaultDfsClient) DfsWriteFilePartData(ctx context.Context, in *dfs.TLDfsWriteFilePartData) (*tg.Bool, error) {
	cli := dfsservice.NewRPCDfsClient(m.cli)
	return cli.DfsWriteFilePartData(ctx, in)
}

// DfsUploadPhotoFileV2
// dfs.uploadPhotoFileV2 creator:long file:InputFile = Photo;
func (m *defaultDfsClient) DfsUploadPhotoFileV2(ctx context.Context, in *dfs.TLDfsUploadPhotoFileV2) (*tg.Photo, error) {
	cli := dfsservice.NewRPCDfsClient(m.cli)
	return cli.DfsUploadPhotoFileV2(ctx, in)
}

// DfsUploadProfilePhotoFileV2
// dfs.uploadProfilePhotoFileV2 flags:# creator:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = Photo;
func (m *defaultDfsClient) DfsUploadProfilePhotoFileV2(ctx context.Context, in *dfs.TLDfsUploadProfilePhotoFileV2) (*tg.Photo, error) {
	cli := dfsservice.NewRPCDfsClient(m.cli)
	return cli.DfsUploadProfilePhotoFileV2(ctx, in)
}

// DfsUploadEncryptedFileV2
// dfs.uploadEncryptedFileV2 creator:long file:InputEncryptedFile = EncryptedFile;
func (m *defaultDfsClient) DfsUploadEncryptedFileV2(ctx context.Context, in *dfs.TLDfsUploadEncryptedFileV2) (*tg.EncryptedFile, error) {
	cli := dfsservice.NewRPCDfsClient(m.cli)
	return cli.DfsUploadEncryptedFileV2(ctx, in)
}

// DfsDownloadFile
// dfs.downloadFile location:InputFileLocation offset:long limit:int = upload.File;
func (m *defaultDfsClient) DfsDownloadFile(ctx context.Context, in *dfs.TLDfsDownloadFile) (*tg.UploadFile, error) {
	cli := dfsservice.NewRPCDfsClient(m.cli)
	return cli.DfsDownloadFile(ctx, in)
}

// DfsUploadDocumentFileV2
// dfs.uploadDocumentFileV2 creator:long media:InputMedia = Document;
func (m *defaultDfsClient) DfsUploadDocumentFileV2(ctx context.Context, in *dfs.TLDfsUploadDocumentFileV2) (*tg.Document, error) {
	cli := dfsservice.NewRPCDfsClient(m.cli)
	return cli.DfsUploadDocumentFileV2(ctx, in)
}

// DfsUploadGifDocumentMedia
// dfs.uploadGifDocumentMedia creator:long media:InputMedia = Document;
func (m *defaultDfsClient) DfsUploadGifDocumentMedia(ctx context.Context, in *dfs.TLDfsUploadGifDocumentMedia) (*tg.Document, error) {
	cli := dfsservice.NewRPCDfsClient(m.cli)
	return cli.DfsUploadGifDocumentMedia(ctx, in)
}

// DfsUploadMp4DocumentMedia
// dfs.uploadMp4DocumentMedia creator:long media:InputMedia = Document;
func (m *defaultDfsClient) DfsUploadMp4DocumentMedia(ctx context.Context, in *dfs.TLDfsUploadMp4DocumentMedia) (*tg.Document, error) {
	cli := dfsservice.NewRPCDfsClient(m.cli)
	return cli.DfsUploadMp4DocumentMedia(ctx, in)
}

// DfsUploadWallPaperFile
// dfs.uploadWallPaperFile creator:long file:InputFile mime_type:string admin:Bool = Document;
func (m *defaultDfsClient) DfsUploadWallPaperFile(ctx context.Context, in *dfs.TLDfsUploadWallPaperFile) (*tg.Document, error) {
	cli := dfsservice.NewRPCDfsClient(m.cli)
	return cli.DfsUploadWallPaperFile(ctx, in)
}

// DfsUploadThemeFile
// dfs.uploadThemeFile flags:# creator:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
func (m *defaultDfsClient) DfsUploadThemeFile(ctx context.Context, in *dfs.TLDfsUploadThemeFile) (*tg.Document, error) {
	cli := dfsservice.NewRPCDfsClient(m.cli)
	return cli.DfsUploadThemeFile(ctx, in)
}

// DfsUploadRingtoneFile
// dfs.uploadRingtoneFile creator:long file:InputFile mime_type:string file_name:string = Document;
func (m *defaultDfsClient) DfsUploadRingtoneFile(ctx context.Context, in *dfs.TLDfsUploadRingtoneFile) (*tg.Document, error) {
	cli := dfsservice.NewRPCDfsClient(m.cli)
	return cli.DfsUploadRingtoneFile(ctx, in)
}

// DfsUploadedProfilePhoto
// dfs.uploadedProfilePhoto creator:long photo_id:long = Photo;
func (m *defaultDfsClient) DfsUploadedProfilePhoto(ctx context.Context, in *dfs.TLDfsUploadedProfilePhoto) (*tg.Photo, error) {
	cli := dfsservice.NewRPCDfsClient(m.cli)
	return cli.DfsUploadedProfilePhoto(ctx, in)
}
