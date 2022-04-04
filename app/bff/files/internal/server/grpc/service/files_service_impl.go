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
	"github.com/teamgram/teamgram-server/app/bff/files/internal/core"
)

// MessagesGetDocumentByHash
// messages.getDocumentByHash#338e2464 sha256:bytes size:int mime_type:string = Document;
func (s *Service) MessagesGetDocumentByHash(ctx context.Context, request *mtproto.TLMessagesGetDocumentByHash) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getDocumentByHash - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetDocumentByHash(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getDocumentByHash - reply: %s", r.DebugString())
	return r, err
}

// MessagesUploadMedia
// messages.uploadMedia#519bc2b1 peer:InputPeer media:InputMedia = MessageMedia;
func (s *Service) MessagesUploadMedia(ctx context.Context, request *mtproto.TLMessagesUploadMedia) (*mtproto.MessageMedia, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.uploadMedia - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesUploadMedia(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.uploadMedia - reply: %s", r.DebugString())
	return r, err
}

// MessagesUploadEncryptedFile
// messages.uploadEncryptedFile#5057c497 peer:InputEncryptedChat file:InputEncryptedFile = EncryptedFile;
func (s *Service) MessagesUploadEncryptedFile(ctx context.Context, request *mtproto.TLMessagesUploadEncryptedFile) (*mtproto.EncryptedFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.uploadEncryptedFile - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesUploadEncryptedFile(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.uploadEncryptedFile - reply: %s", r.DebugString())
	return r, err
}

// UploadSaveFilePart
// upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
func (s *Service) UploadSaveFilePart(ctx context.Context, request *mtproto.TLUploadSaveFilePart) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("upload.saveFilePart - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UploadSaveFilePart(request)
	if err != nil {
		return nil, err
	}

	c.Infof("upload.saveFilePart#b304a621 - metadata: %s, request: {file_id: %d, file_part: %d, bytes_len: %d}",
		c.MD.DebugString(),
		request.FileId,
		request.FilePart,
		len(request.Bytes))
	return r, err
}

// UploadGetFile
// upload.getFile#b15a9afc flags:# precise:flags.0?true cdn_supported:flags.1?true location:InputFileLocation offset:int limit:int = upload.File;
func (s *Service) UploadGetFile(ctx context.Context, request *mtproto.TLUploadGetFile) (*mtproto.Upload_File, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("upload.getFile - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UploadGetFile(request)
	if err != nil {
		return nil, err
	}

	c.Infof("upload.getFile - reply: {type: %v, mime: %d, len_bytes: %d}",
		r.GetType(),
		r.GetMtime(),
		len(r.GetBytes()))
	return r, err
}

// UploadSaveBigFilePart
// upload.saveBigFilePart#de7b673d file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;
func (s *Service) UploadSaveBigFilePart(ctx context.Context, request *mtproto.TLUploadSaveBigFilePart) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("upload.saveBigFilePart#de7b673d - metadata: %s, request: {file_id: %d, file_part: %d, bytes_len: %d}",
		c.MD.DebugString(),
		request.FileId,
		request.FilePart,
		len(request.Bytes))

	r, err := c.UploadSaveBigFilePart(request)
	if err != nil {
		return nil, err
	}

	c.Infof("upload.saveBigFilePart - reply: %s", r.DebugString())
	return r, err
}

// UploadGetWebFile
// upload.getWebFile#24e6818d location:InputWebFileLocation offset:int limit:int = upload.WebFile;
func (s *Service) UploadGetWebFile(ctx context.Context, request *mtproto.TLUploadGetWebFile) (*mtproto.Upload_WebFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("upload.getWebFile - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UploadGetWebFile(request)
	if err != nil {
		return nil, err
	}

	c.Infof("upload.getWebFile - reply: {size: %d, mime_type: %, file_type: %s, len_bytes: %d}",
		r.GetSize2(),
		r.GetMimeType(),
		r.GetFileType().DebugString(),
		len(r.GetBytes()))
	return r, err
}

// UploadGetCdnFile
// upload.getCdnFile#2000bcc3 file_token:bytes offset:int limit:int = upload.CdnFile;
func (s *Service) UploadGetCdnFile(ctx context.Context, request *mtproto.TLUploadGetCdnFile) (*mtproto.Upload_CdnFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("upload.getCdnFile - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UploadGetCdnFile(request)
	if err != nil {
		return nil, err
	}

	c.Infof("upload.getCdnFile - reply: %s", r.DebugString())
	return r, err
}

// UploadReuploadCdnFile
// upload.reuploadCdnFile#9b2754a8 file_token:bytes request_token:bytes = Vector<FileHash>;
func (s *Service) UploadReuploadCdnFile(ctx context.Context, request *mtproto.TLUploadReuploadCdnFile) (*mtproto.Vector_FileHash, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("upload.reuploadCdnFile - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UploadReuploadCdnFile(request)
	if err != nil {
		return nil, err
	}

	c.Infof("upload.reuploadCdnFile - reply: %s", r.DebugString())
	return r, err
}

// UploadGetCdnFileHashes
// upload.getCdnFileHashes#4da54231 file_token:bytes offset:int = Vector<FileHash>;
func (s *Service) UploadGetCdnFileHashes(ctx context.Context, request *mtproto.TLUploadGetCdnFileHashes) (*mtproto.Vector_FileHash, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("upload.getCdnFileHashes - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UploadGetCdnFileHashes(request)
	if err != nil {
		return nil, err
	}

	c.Infof("upload.getCdnFileHashes - reply: %s", r.DebugString())
	return r, err
}

// UploadGetFileHashes
// upload.getFileHashes#c7025931 location:InputFileLocation offset:int = Vector<FileHash>;
func (s *Service) UploadGetFileHashes(ctx context.Context, request *mtproto.TLUploadGetFileHashes) (*mtproto.Vector_FileHash, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("upload.getFileHashes - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UploadGetFileHashes(request)
	if err != nil {
		return nil, err
	}

	c.Infof("upload.getFileHashes - reply: %s", r.DebugString())
	return r, err
}

// HelpGetCdnConfig
// help.getCdnConfig#52029342 = CdnConfig;
func (s *Service) HelpGetCdnConfig(ctx context.Context, request *mtproto.TLHelpGetCdnConfig) (*mtproto.CdnConfig, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("help.getCdnConfig - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.HelpGetCdnConfig(request)
	if err != nil {
		return nil, err
	}

	c.Infof("help.getCdnConfig - reply: %s", r.DebugString())
	return r, err
}
