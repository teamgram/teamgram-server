/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package filesservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MessagesGetDocumentByHash(ctx context.Context, req *tg.TLMessagesGetDocumentByHash, callOptions ...callopt.Option) (r *tg.Document, err error)
	MessagesUploadMedia(ctx context.Context, req *tg.TLMessagesUploadMedia, callOptions ...callopt.Option) (r *tg.MessageMedia, err error)
	MessagesUploadEncryptedFile(ctx context.Context, req *tg.TLMessagesUploadEncryptedFile, callOptions ...callopt.Option) (r *tg.EncryptedFile, err error)
	UploadSaveFilePart(ctx context.Context, req *tg.TLUploadSaveFilePart, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UploadGetFile(ctx context.Context, req *tg.TLUploadGetFile, callOptions ...callopt.Option) (r *tg.UploadFile, err error)
	UploadSaveBigFilePart(ctx context.Context, req *tg.TLUploadSaveBigFilePart, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UploadGetWebFile(ctx context.Context, req *tg.TLUploadGetWebFile, callOptions ...callopt.Option) (r *tg.UploadWebFile, err error)
	UploadGetCdnFile(ctx context.Context, req *tg.TLUploadGetCdnFile, callOptions ...callopt.Option) (r *tg.UploadCdnFile, err error)
	UploadReuploadCdnFile(ctx context.Context, req *tg.TLUploadReuploadCdnFile, callOptions ...callopt.Option) (r *tg.VectorFileHash, err error)
	UploadGetCdnFileHashes(ctx context.Context, req *tg.TLUploadGetCdnFileHashes, callOptions ...callopt.Option) (r *tg.VectorFileHash, err error)
	UploadGetFileHashes(ctx context.Context, req *tg.TLUploadGetFileHashes, callOptions ...callopt.Option) (r *tg.VectorFileHash, err error)
	HelpGetCdnConfig(ctx context.Context, req *tg.TLHelpGetCdnConfig, callOptions ...callopt.Option) (r *tg.CdnConfig, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kFilesClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFilesClient struct {
	*kClient
}

func NewRPCFilesClient(cli client.Client) Client {
	return &kFilesClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kFilesClient) MessagesGetDocumentByHash(ctx context.Context, req *tg.TLMessagesGetDocumentByHash, callOptions ...callopt.Option) (r *tg.Document, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetDocumentByHash(ctx, req)
}

func (p *kFilesClient) MessagesUploadMedia(ctx context.Context, req *tg.TLMessagesUploadMedia, callOptions ...callopt.Option) (r *tg.MessageMedia, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesUploadMedia(ctx, req)
}

func (p *kFilesClient) MessagesUploadEncryptedFile(ctx context.Context, req *tg.TLMessagesUploadEncryptedFile, callOptions ...callopt.Option) (r *tg.EncryptedFile, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesUploadEncryptedFile(ctx, req)
}

func (p *kFilesClient) UploadSaveFilePart(ctx context.Context, req *tg.TLUploadSaveFilePart, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadSaveFilePart(ctx, req)
}

func (p *kFilesClient) UploadGetFile(ctx context.Context, req *tg.TLUploadGetFile, callOptions ...callopt.Option) (r *tg.UploadFile, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadGetFile(ctx, req)
}

func (p *kFilesClient) UploadSaveBigFilePart(ctx context.Context, req *tg.TLUploadSaveBigFilePart, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadSaveBigFilePart(ctx, req)
}

func (p *kFilesClient) UploadGetWebFile(ctx context.Context, req *tg.TLUploadGetWebFile, callOptions ...callopt.Option) (r *tg.UploadWebFile, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadGetWebFile(ctx, req)
}

func (p *kFilesClient) UploadGetCdnFile(ctx context.Context, req *tg.TLUploadGetCdnFile, callOptions ...callopt.Option) (r *tg.UploadCdnFile, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadGetCdnFile(ctx, req)
}

func (p *kFilesClient) UploadReuploadCdnFile(ctx context.Context, req *tg.TLUploadReuploadCdnFile, callOptions ...callopt.Option) (r *tg.VectorFileHash, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadReuploadCdnFile(ctx, req)
}

func (p *kFilesClient) UploadGetCdnFileHashes(ctx context.Context, req *tg.TLUploadGetCdnFileHashes, callOptions ...callopt.Option) (r *tg.VectorFileHash, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadGetCdnFileHashes(ctx, req)
}

func (p *kFilesClient) UploadGetFileHashes(ctx context.Context, req *tg.TLUploadGetFileHashes, callOptions ...callopt.Option) (r *tg.VectorFileHash, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadGetFileHashes(ctx, req)
}

func (p *kFilesClient) HelpGetCdnConfig(ctx context.Context, req *tg.TLHelpGetCdnConfig, callOptions ...callopt.Option) (r *tg.CdnConfig, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelpGetCdnConfig(ctx, req)
}
