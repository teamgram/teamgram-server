/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dfs_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type DfsClient interface {
	DfsWriteFilePartData(ctx context.Context, in *dfs.TLDfsWriteFilePartData) (*mtproto.Bool, error)
	DfsUploadPhotoFileV2(ctx context.Context, in *dfs.TLDfsUploadPhotoFileV2) (*mtproto.Photo, error)
	DfsUploadProfilePhotoFileV2(ctx context.Context, in *dfs.TLDfsUploadProfilePhotoFileV2) (*mtproto.Photo, error)
	DfsUploadEncryptedFileV2(ctx context.Context, in *dfs.TLDfsUploadEncryptedFileV2) (*mtproto.EncryptedFile, error)
	DfsDownloadFile(ctx context.Context, in *dfs.TLDfsDownloadFile) (*mtproto.Upload_File, error)
	DfsUploadDocumentFileV2(ctx context.Context, in *dfs.TLDfsUploadDocumentFileV2) (*mtproto.Document, error)
	DfsUploadGifDocumentMedia(ctx context.Context, in *dfs.TLDfsUploadGifDocumentMedia) (*mtproto.Document, error)
	DfsUploadMp4DocumentMedia(ctx context.Context, in *dfs.TLDfsUploadMp4DocumentMedia) (*mtproto.Document, error)
	DfsUploadWallPaperFile(ctx context.Context, in *dfs.TLDfsUploadWallPaperFile) (*mtproto.Document, error)
	DfsUploadThemeFile(ctx context.Context, in *dfs.TLDfsUploadThemeFile) (*mtproto.Document, error)
	DfsUploadRingtoneFile(ctx context.Context, in *dfs.TLDfsUploadRingtoneFile) (*mtproto.Document, error)
}

type defaultDfsClient struct {
	cli zrpc.Client
}

func NewDfsClient(cli zrpc.Client) DfsClient {
	return &defaultDfsClient{
		cli: cli,
	}
}

// DfsWriteFilePartData
// dfs.writeFilePartData flags:# creator:long file_id:long file_part:int bytes:bytes big:flags.0?true file_total_parts:flags.1?int = Bool;
func (m *defaultDfsClient) DfsWriteFilePartData(ctx context.Context, in *dfs.TLDfsWriteFilePartData) (*mtproto.Bool, error) {
	client := dfs.NewRPCDfsClient(m.cli.Conn())
	return client.DfsWriteFilePartData(ctx, in)
}

// DfsUploadPhotoFileV2
// dfs.uploadPhotoFileV2 creator:long file:InputFile = Photo;
func (m *defaultDfsClient) DfsUploadPhotoFileV2(ctx context.Context, in *dfs.TLDfsUploadPhotoFileV2) (*mtproto.Photo, error) {
	client := dfs.NewRPCDfsClient(m.cli.Conn())
	return client.DfsUploadPhotoFileV2(ctx, in)
}

// DfsUploadProfilePhotoFileV2
// dfs.uploadProfilePhotoFileV2 flags:# creator:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = Photo;
func (m *defaultDfsClient) DfsUploadProfilePhotoFileV2(ctx context.Context, in *dfs.TLDfsUploadProfilePhotoFileV2) (*mtproto.Photo, error) {
	client := dfs.NewRPCDfsClient(m.cli.Conn())
	return client.DfsUploadProfilePhotoFileV2(ctx, in)
}

// DfsUploadEncryptedFileV2
// dfs.uploadEncryptedFileV2 creator:long file:InputEncryptedFile = EncryptedFile;
func (m *defaultDfsClient) DfsUploadEncryptedFileV2(ctx context.Context, in *dfs.TLDfsUploadEncryptedFileV2) (*mtproto.EncryptedFile, error) {
	client := dfs.NewRPCDfsClient(m.cli.Conn())
	return client.DfsUploadEncryptedFileV2(ctx, in)
}

// DfsDownloadFile
// dfs.downloadFile location:InputFileLocation offset:long limit:int = upload.File;
func (m *defaultDfsClient) DfsDownloadFile(ctx context.Context, in *dfs.TLDfsDownloadFile) (*mtproto.Upload_File, error) {
	client := dfs.NewRPCDfsClient(m.cli.Conn())
	return client.DfsDownloadFile(ctx, in)
}

// DfsUploadDocumentFileV2
// dfs.uploadDocumentFileV2 creator:long media:InputMedia = Document;
func (m *defaultDfsClient) DfsUploadDocumentFileV2(ctx context.Context, in *dfs.TLDfsUploadDocumentFileV2) (*mtproto.Document, error) {
	client := dfs.NewRPCDfsClient(m.cli.Conn())
	return client.DfsUploadDocumentFileV2(ctx, in)
}

// DfsUploadGifDocumentMedia
// dfs.uploadGifDocumentMedia creator:long media:InputMedia = Document;
func (m *defaultDfsClient) DfsUploadGifDocumentMedia(ctx context.Context, in *dfs.TLDfsUploadGifDocumentMedia) (*mtproto.Document, error) {
	client := dfs.NewRPCDfsClient(m.cli.Conn())
	return client.DfsUploadGifDocumentMedia(ctx, in)
}

// DfsUploadMp4DocumentMedia
// dfs.uploadMp4DocumentMedia creator:long media:InputMedia = Document;
func (m *defaultDfsClient) DfsUploadMp4DocumentMedia(ctx context.Context, in *dfs.TLDfsUploadMp4DocumentMedia) (*mtproto.Document, error) {
	client := dfs.NewRPCDfsClient(m.cli.Conn())
	return client.DfsUploadMp4DocumentMedia(ctx, in)
}

// DfsUploadWallPaperFile
// dfs.uploadWallPaperFile creator:long file:InputFile mime_type:string admin:Bool = Document;
func (m *defaultDfsClient) DfsUploadWallPaperFile(ctx context.Context, in *dfs.TLDfsUploadWallPaperFile) (*mtproto.Document, error) {
	client := dfs.NewRPCDfsClient(m.cli.Conn())
	return client.DfsUploadWallPaperFile(ctx, in)
}

// DfsUploadThemeFile
// dfs.uploadThemeFile flags:# creator:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
func (m *defaultDfsClient) DfsUploadThemeFile(ctx context.Context, in *dfs.TLDfsUploadThemeFile) (*mtproto.Document, error) {
	client := dfs.NewRPCDfsClient(m.cli.Conn())
	return client.DfsUploadThemeFile(ctx, in)
}

// DfsUploadRingtoneFile
// dfs.uploadRingtoneFile creator:long file:InputFile mime_type:string file_name:string = Document;
func (m *defaultDfsClient) DfsUploadRingtoneFile(ctx context.Context, in *dfs.TLDfsUploadRingtoneFile) (*mtproto.Document, error) {
	client := dfs.NewRPCDfsClient(m.cli.Conn())
	return client.DfsUploadRingtoneFile(ctx, in)
}
