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
	"github.com/teamgram/teamgram-server/v2/app/bff/files/internal/core"
)

// MessagesGetDocumentByHash
// messages.getDocumentByHash#b1f2061f sha256:bytes size:long mime_type:string = Document;
func (s *Service) MessagesGetDocumentByHash(ctx context.Context, request *tg.TLMessagesGetDocumentByHash) (*tg.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getDocumentByHash - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetDocumentByHash(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getDocumentByHash - reply: {%v}", r)
	return r, err
}

// MessagesUploadMedia
// messages.uploadMedia#14967978 flags:# business_connection_id:flags.0?string peer:InputPeer media:InputMedia = MessageMedia;
func (s *Service) MessagesUploadMedia(ctx context.Context, request *tg.TLMessagesUploadMedia) (*tg.MessageMedia, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.uploadMedia - metadata: {}, request: {%v}", request)

	r, err := c.MessagesUploadMedia(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.uploadMedia - reply: {%v}", r)
	return r, err
}

// MessagesUploadEncryptedFile
// messages.uploadEncryptedFile#5057c497 peer:InputEncryptedChat file:InputEncryptedFile = EncryptedFile;
func (s *Service) MessagesUploadEncryptedFile(ctx context.Context, request *tg.TLMessagesUploadEncryptedFile) (*tg.EncryptedFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.uploadEncryptedFile - metadata: {}, request: {%v}", request)

	r, err := c.MessagesUploadEncryptedFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.uploadEncryptedFile - reply: {%v}", r)
	return r, err
}

// UploadSaveFilePart
// upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
func (s *Service) UploadSaveFilePart(ctx context.Context, request *tg.TLUploadSaveFilePart) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("upload.saveFilePart - metadata: {}, request: {%v}", request)

	r, err := c.UploadSaveFilePart(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("upload.saveFilePart - reply: {%v}", r)
	return r, err
}

// UploadGetFile
// upload.getFile#be5335be flags:# precise:flags.0?true cdn_supported:flags.1?true location:InputFileLocation offset:long limit:int = upload.File;
func (s *Service) UploadGetFile(ctx context.Context, request *tg.TLUploadGetFile) (*tg.UploadFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("upload.getFile - metadata: {}, request: {%v}", request)

	r, err := c.UploadGetFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("upload.getFile - reply: {%v}", r)
	return r, err
}

// UploadSaveBigFilePart
// upload.saveBigFilePart#de7b673d file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;
func (s *Service) UploadSaveBigFilePart(ctx context.Context, request *tg.TLUploadSaveBigFilePart) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("upload.saveBigFilePart - metadata: {}, request: {%v}", request)

	r, err := c.UploadSaveBigFilePart(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("upload.saveBigFilePart - reply: {%v}", r)
	return r, err
}

// UploadGetWebFile
// upload.getWebFile#24e6818d location:InputWebFileLocation offset:int limit:int = upload.WebFile;
func (s *Service) UploadGetWebFile(ctx context.Context, request *tg.TLUploadGetWebFile) (*tg.UploadWebFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("upload.getWebFile - metadata: {}, request: {%v}", request)

	r, err := c.UploadGetWebFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("upload.getWebFile - reply: {%v}", r)
	return r, err
}

// UploadGetCdnFile
// upload.getCdnFile#395f69da file_token:bytes offset:long limit:int = upload.CdnFile;
func (s *Service) UploadGetCdnFile(ctx context.Context, request *tg.TLUploadGetCdnFile) (*tg.UploadCdnFile, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("upload.getCdnFile - metadata: {}, request: {%v}", request)

	r, err := c.UploadGetCdnFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("upload.getCdnFile - reply: {%v}", r)
	return r, err
}

// UploadReuploadCdnFile
// upload.reuploadCdnFile#9b2754a8 file_token:bytes request_token:bytes = Vector<FileHash>;
func (s *Service) UploadReuploadCdnFile(ctx context.Context, request *tg.TLUploadReuploadCdnFile) (*tg.VectorFileHash, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("upload.reuploadCdnFile - metadata: {}, request: {%v}", request)

	r, err := c.UploadReuploadCdnFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("upload.reuploadCdnFile - reply: {%v}", r)
	return r, err
}

// UploadGetCdnFileHashes
// upload.getCdnFileHashes#91dc3f31 file_token:bytes offset:long = Vector<FileHash>;
func (s *Service) UploadGetCdnFileHashes(ctx context.Context, request *tg.TLUploadGetCdnFileHashes) (*tg.VectorFileHash, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("upload.getCdnFileHashes - metadata: {}, request: {%v}", request)

	r, err := c.UploadGetCdnFileHashes(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("upload.getCdnFileHashes - reply: {%v}", r)
	return r, err
}

// UploadGetFileHashes
// upload.getFileHashes#9156982a location:InputFileLocation offset:long = Vector<FileHash>;
func (s *Service) UploadGetFileHashes(ctx context.Context, request *tg.TLUploadGetFileHashes) (*tg.VectorFileHash, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("upload.getFileHashes - metadata: {}, request: {%v}", request)

	r, err := c.UploadGetFileHashes(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("upload.getFileHashes - reply: {%v}", r)
	return r, err
}

// HelpGetCdnConfig
// help.getCdnConfig#52029342 = CdnConfig;
func (s *Service) HelpGetCdnConfig(ctx context.Context, request *tg.TLHelpGetCdnConfig) (*tg.CdnConfig, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getCdnConfig - metadata: {}, request: {%v}", request)

	r, err := c.HelpGetCdnConfig(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getCdnConfig - reply: {%v}", r)
	return r, err
}
