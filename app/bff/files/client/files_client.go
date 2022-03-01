/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package files_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type FilesClient interface {
	MessagesGetDocumentByHash(ctx context.Context, in *mtproto.TLMessagesGetDocumentByHash) (*mtproto.Document, error)
	MessagesUploadMedia(ctx context.Context, in *mtproto.TLMessagesUploadMedia) (*mtproto.MessageMedia, error)
	MessagesUploadEncryptedFile(ctx context.Context, in *mtproto.TLMessagesUploadEncryptedFile) (*mtproto.EncryptedFile, error)
	UploadSaveFilePart(ctx context.Context, in *mtproto.TLUploadSaveFilePart) (*mtproto.Bool, error)
	UploadGetFile(ctx context.Context, in *mtproto.TLUploadGetFile) (*mtproto.Upload_File, error)
	UploadSaveBigFilePart(ctx context.Context, in *mtproto.TLUploadSaveBigFilePart) (*mtproto.Bool, error)
	UploadGetWebFile(ctx context.Context, in *mtproto.TLUploadGetWebFile) (*mtproto.Upload_WebFile, error)
	UploadGetCdnFile(ctx context.Context, in *mtproto.TLUploadGetCdnFile) (*mtproto.Upload_CdnFile, error)
	UploadReuploadCdnFile(ctx context.Context, in *mtproto.TLUploadReuploadCdnFile) (*mtproto.Vector_FileHash, error)
	UploadGetCdnFileHashes(ctx context.Context, in *mtproto.TLUploadGetCdnFileHashes) (*mtproto.Vector_FileHash, error)
	UploadGetFileHashes(ctx context.Context, in *mtproto.TLUploadGetFileHashes) (*mtproto.Vector_FileHash, error)
	HelpGetCdnConfig(ctx context.Context, in *mtproto.TLHelpGetCdnConfig) (*mtproto.CdnConfig, error)
}

type defaultFilesClient struct {
	cli zrpc.Client
}

func NewFilesClient(cli zrpc.Client) FilesClient {
	return &defaultFilesClient{
		cli: cli,
	}
}

// MessagesGetDocumentByHash
// messages.getDocumentByHash#338e2464 sha256:bytes size:int mime_type:string = Document;
func (m *defaultFilesClient) MessagesGetDocumentByHash(ctx context.Context, in *mtproto.TLMessagesGetDocumentByHash) (*mtproto.Document, error) {
	client := mtproto.NewRPCFilesClient(m.cli.Conn())
	return client.MessagesGetDocumentByHash(ctx, in)
}

// MessagesUploadMedia
// messages.uploadMedia#519bc2b1 peer:InputPeer media:InputMedia = MessageMedia;
func (m *defaultFilesClient) MessagesUploadMedia(ctx context.Context, in *mtproto.TLMessagesUploadMedia) (*mtproto.MessageMedia, error) {
	client := mtproto.NewRPCFilesClient(m.cli.Conn())
	return client.MessagesUploadMedia(ctx, in)
}

// MessagesUploadEncryptedFile
// messages.uploadEncryptedFile#5057c497 peer:InputEncryptedChat file:InputEncryptedFile = EncryptedFile;
func (m *defaultFilesClient) MessagesUploadEncryptedFile(ctx context.Context, in *mtproto.TLMessagesUploadEncryptedFile) (*mtproto.EncryptedFile, error) {
	client := mtproto.NewRPCFilesClient(m.cli.Conn())
	return client.MessagesUploadEncryptedFile(ctx, in)
}

// UploadSaveFilePart
// upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
func (m *defaultFilesClient) UploadSaveFilePart(ctx context.Context, in *mtproto.TLUploadSaveFilePart) (*mtproto.Bool, error) {
	client := mtproto.NewRPCFilesClient(m.cli.Conn())
	return client.UploadSaveFilePart(ctx, in)
}

// UploadGetFile
// upload.getFile#b15a9afc flags:# precise:flags.0?true cdn_supported:flags.1?true location:InputFileLocation offset:int limit:int = upload.File;
func (m *defaultFilesClient) UploadGetFile(ctx context.Context, in *mtproto.TLUploadGetFile) (*mtproto.Upload_File, error) {
	client := mtproto.NewRPCFilesClient(m.cli.Conn())
	return client.UploadGetFile(ctx, in)
}

// UploadSaveBigFilePart
// upload.saveBigFilePart#de7b673d file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;
func (m *defaultFilesClient) UploadSaveBigFilePart(ctx context.Context, in *mtproto.TLUploadSaveBigFilePart) (*mtproto.Bool, error) {
	client := mtproto.NewRPCFilesClient(m.cli.Conn())
	return client.UploadSaveBigFilePart(ctx, in)
}

// UploadGetWebFile
// upload.getWebFile#24e6818d location:InputWebFileLocation offset:int limit:int = upload.WebFile;
func (m *defaultFilesClient) UploadGetWebFile(ctx context.Context, in *mtproto.TLUploadGetWebFile) (*mtproto.Upload_WebFile, error) {
	client := mtproto.NewRPCFilesClient(m.cli.Conn())
	return client.UploadGetWebFile(ctx, in)
}

// UploadGetCdnFile
// upload.getCdnFile#2000bcc3 file_token:bytes offset:int limit:int = upload.CdnFile;
func (m *defaultFilesClient) UploadGetCdnFile(ctx context.Context, in *mtproto.TLUploadGetCdnFile) (*mtproto.Upload_CdnFile, error) {
	client := mtproto.NewRPCFilesClient(m.cli.Conn())
	return client.UploadGetCdnFile(ctx, in)
}

// UploadReuploadCdnFile
// upload.reuploadCdnFile#9b2754a8 file_token:bytes request_token:bytes = Vector<FileHash>;
func (m *defaultFilesClient) UploadReuploadCdnFile(ctx context.Context, in *mtproto.TLUploadReuploadCdnFile) (*mtproto.Vector_FileHash, error) {
	client := mtproto.NewRPCFilesClient(m.cli.Conn())
	return client.UploadReuploadCdnFile(ctx, in)
}

// UploadGetCdnFileHashes
// upload.getCdnFileHashes#4da54231 file_token:bytes offset:int = Vector<FileHash>;
func (m *defaultFilesClient) UploadGetCdnFileHashes(ctx context.Context, in *mtproto.TLUploadGetCdnFileHashes) (*mtproto.Vector_FileHash, error) {
	client := mtproto.NewRPCFilesClient(m.cli.Conn())
	return client.UploadGetCdnFileHashes(ctx, in)
}

// UploadGetFileHashes
// upload.getFileHashes#c7025931 location:InputFileLocation offset:int = Vector<FileHash>;
func (m *defaultFilesClient) UploadGetFileHashes(ctx context.Context, in *mtproto.TLUploadGetFileHashes) (*mtproto.Vector_FileHash, error) {
	client := mtproto.NewRPCFilesClient(m.cli.Conn())
	return client.UploadGetFileHashes(ctx, in)
}

// HelpGetCdnConfig
// help.getCdnConfig#52029342 = CdnConfig;
func (m *defaultFilesClient) HelpGetCdnConfig(ctx context.Context, in *mtproto.TLHelpGetCdnConfig) (*mtproto.CdnConfig, error) {
	client := mtproto.NewRPCFilesClient(m.cli.Conn())
	return client.HelpGetCdnConfig(ctx, in)
}
