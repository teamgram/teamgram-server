/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package filesclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/files/files/filesservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

type FilesClient interface {
	MessagesGetDocumentByHash(ctx context.Context, in *tg.TLMessagesGetDocumentByHash) (*tg.Document, error)
	MessagesUploadMedia(ctx context.Context, in *tg.TLMessagesUploadMedia) (*tg.MessageMedia, error)
	MessagesUploadEncryptedFile(ctx context.Context, in *tg.TLMessagesUploadEncryptedFile) (*tg.EncryptedFile, error)
	UploadSaveFilePart(ctx context.Context, in *tg.TLUploadSaveFilePart) (*tg.Bool, error)
	UploadGetFile(ctx context.Context, in *tg.TLUploadGetFile) (*tg.UploadFile, error)
	UploadSaveBigFilePart(ctx context.Context, in *tg.TLUploadSaveBigFilePart) (*tg.Bool, error)
	UploadGetWebFile(ctx context.Context, in *tg.TLUploadGetWebFile) (*tg.UploadWebFile, error)
	UploadGetCdnFile(ctx context.Context, in *tg.TLUploadGetCdnFile) (*tg.UploadCdnFile, error)
	UploadReuploadCdnFile(ctx context.Context, in *tg.TLUploadReuploadCdnFile) (*tg.VectorFileHash, error)
	UploadGetCdnFileHashes(ctx context.Context, in *tg.TLUploadGetCdnFileHashes) (*tg.VectorFileHash, error)
	UploadGetFileHashes(ctx context.Context, in *tg.TLUploadGetFileHashes) (*tg.VectorFileHash, error)
	HelpGetCdnConfig(ctx context.Context, in *tg.TLHelpGetCdnConfig) (*tg.CdnConfig, error)
}

type defaultFilesClient struct {
	cli client.Client
	rpc filesservice.Client
}

func NewFilesClient(cli client.Client) FilesClient {
	return &defaultFilesClient{
		cli: cli,
		rpc: filesservice.NewRPCFilesClient(cli),
	}
}

// MessagesGetDocumentByHash
// messages.getDocumentByHash#b1f2061f sha256:bytes size:long mime_type:string = Document;
func (m *defaultFilesClient) MessagesGetDocumentByHash(ctx context.Context, in *tg.TLMessagesGetDocumentByHash) (*tg.Document, error) {
	return m.rpc.MessagesGetDocumentByHash(ctx, in)
}

// MessagesUploadMedia
// messages.uploadMedia#14967978 flags:# business_connection_id:flags.0?string peer:InputPeer media:InputMedia = MessageMedia;
func (m *defaultFilesClient) MessagesUploadMedia(ctx context.Context, in *tg.TLMessagesUploadMedia) (*tg.MessageMedia, error) {
	return m.rpc.MessagesUploadMedia(ctx, in)
}

// MessagesUploadEncryptedFile
// messages.uploadEncryptedFile#5057c497 peer:InputEncryptedChat file:InputEncryptedFile = EncryptedFile;
func (m *defaultFilesClient) MessagesUploadEncryptedFile(ctx context.Context, in *tg.TLMessagesUploadEncryptedFile) (*tg.EncryptedFile, error) {
	return m.rpc.MessagesUploadEncryptedFile(ctx, in)
}

// UploadSaveFilePart
// upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
func (m *defaultFilesClient) UploadSaveFilePart(ctx context.Context, in *tg.TLUploadSaveFilePart) (*tg.Bool, error) {
	return m.rpc.UploadSaveFilePart(ctx, in)
}

// UploadGetFile
// upload.getFile#be5335be flags:# precise:flags.0?true cdn_supported:flags.1?true location:InputFileLocation offset:long limit:int = upload.File;
func (m *defaultFilesClient) UploadGetFile(ctx context.Context, in *tg.TLUploadGetFile) (*tg.UploadFile, error) {
	return m.rpc.UploadGetFile(ctx, in)
}

// UploadSaveBigFilePart
// upload.saveBigFilePart#de7b673d file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;
func (m *defaultFilesClient) UploadSaveBigFilePart(ctx context.Context, in *tg.TLUploadSaveBigFilePart) (*tg.Bool, error) {
	return m.rpc.UploadSaveBigFilePart(ctx, in)
}

// UploadGetWebFile
// upload.getWebFile#24e6818d location:InputWebFileLocation offset:int limit:int = upload.WebFile;
func (m *defaultFilesClient) UploadGetWebFile(ctx context.Context, in *tg.TLUploadGetWebFile) (*tg.UploadWebFile, error) {
	return m.rpc.UploadGetWebFile(ctx, in)
}

// UploadGetCdnFile
// upload.getCdnFile#395f69da file_token:bytes offset:long limit:int = upload.CdnFile;
func (m *defaultFilesClient) UploadGetCdnFile(ctx context.Context, in *tg.TLUploadGetCdnFile) (*tg.UploadCdnFile, error) {
	return m.rpc.UploadGetCdnFile(ctx, in)
}

// UploadReuploadCdnFile
// upload.reuploadCdnFile#9b2754a8 file_token:bytes request_token:bytes = Vector<FileHash>;
func (m *defaultFilesClient) UploadReuploadCdnFile(ctx context.Context, in *tg.TLUploadReuploadCdnFile) (*tg.VectorFileHash, error) {
	return m.rpc.UploadReuploadCdnFile(ctx, in)
}

// UploadGetCdnFileHashes
// upload.getCdnFileHashes#91dc3f31 file_token:bytes offset:long = Vector<FileHash>;
func (m *defaultFilesClient) UploadGetCdnFileHashes(ctx context.Context, in *tg.TLUploadGetCdnFileHashes) (*tg.VectorFileHash, error) {
	return m.rpc.UploadGetCdnFileHashes(ctx, in)
}

// UploadGetFileHashes
// upload.getFileHashes#9156982a location:InputFileLocation offset:long = Vector<FileHash>;
func (m *defaultFilesClient) UploadGetFileHashes(ctx context.Context, in *tg.TLUploadGetFileHashes) (*tg.VectorFileHash, error) {
	return m.rpc.UploadGetFileHashes(ctx, in)
}

// HelpGetCdnConfig
// help.getCdnConfig#52029342 = CdnConfig;
func (m *defaultFilesClient) HelpGetCdnConfig(ctx context.Context, in *tg.TLHelpGetCdnConfig) (*tg.CdnConfig, error) {
	return m.rpc.HelpGetCdnConfig(ctx, in)
}
