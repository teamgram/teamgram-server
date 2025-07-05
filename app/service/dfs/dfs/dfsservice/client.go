/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dfsservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	DfsWriteFilePartData(ctx context.Context, req *dfs.TLDfsWriteFilePartData, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DfsUploadPhotoFileV2(ctx context.Context, req *dfs.TLDfsUploadPhotoFileV2, callOptions ...callopt.Option) (r *tg.Photo, err error)
	DfsUploadProfilePhotoFileV2(ctx context.Context, req *dfs.TLDfsUploadProfilePhotoFileV2, callOptions ...callopt.Option) (r *tg.Photo, err error)
	DfsUploadEncryptedFileV2(ctx context.Context, req *dfs.TLDfsUploadEncryptedFileV2, callOptions ...callopt.Option) (r *tg.EncryptedFile, err error)
	DfsDownloadFile(ctx context.Context, req *dfs.TLDfsDownloadFile, callOptions ...callopt.Option) (r *tg.UploadFile, err error)
	DfsUploadDocumentFileV2(ctx context.Context, req *dfs.TLDfsUploadDocumentFileV2, callOptions ...callopt.Option) (r *tg.Document, err error)
	DfsUploadGifDocumentMedia(ctx context.Context, req *dfs.TLDfsUploadGifDocumentMedia, callOptions ...callopt.Option) (r *tg.Document, err error)
	DfsUploadMp4DocumentMedia(ctx context.Context, req *dfs.TLDfsUploadMp4DocumentMedia, callOptions ...callopt.Option) (r *tg.Document, err error)
	DfsUploadWallPaperFile(ctx context.Context, req *dfs.TLDfsUploadWallPaperFile, callOptions ...callopt.Option) (r *tg.Document, err error)
	DfsUploadThemeFile(ctx context.Context, req *dfs.TLDfsUploadThemeFile, callOptions ...callopt.Option) (r *tg.Document, err error)
	DfsUploadRingtoneFile(ctx context.Context, req *dfs.TLDfsUploadRingtoneFile, callOptions ...callopt.Option) (r *tg.Document, err error)
	DfsUploadedProfilePhoto(ctx context.Context, req *dfs.TLDfsUploadedProfilePhoto, callOptions ...callopt.Option) (r *tg.Photo, err error)
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
	return &kDfsClient{
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

type kDfsClient struct {
	*kClient
}

func NewRPCDfsClient(cli client.Client) Client {
	return &kDfsClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kDfsClient) DfsWriteFilePartData(ctx context.Context, req *dfs.TLDfsWriteFilePartData, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DfsWriteFilePartData(ctx, req)
}

func (p *kDfsClient) DfsUploadPhotoFileV2(ctx context.Context, req *dfs.TLDfsUploadPhotoFileV2, callOptions ...callopt.Option) (r *tg.Photo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DfsUploadPhotoFileV2(ctx, req)
}

func (p *kDfsClient) DfsUploadProfilePhotoFileV2(ctx context.Context, req *dfs.TLDfsUploadProfilePhotoFileV2, callOptions ...callopt.Option) (r *tg.Photo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DfsUploadProfilePhotoFileV2(ctx, req)
}

func (p *kDfsClient) DfsUploadEncryptedFileV2(ctx context.Context, req *dfs.TLDfsUploadEncryptedFileV2, callOptions ...callopt.Option) (r *tg.EncryptedFile, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DfsUploadEncryptedFileV2(ctx, req)
}

func (p *kDfsClient) DfsDownloadFile(ctx context.Context, req *dfs.TLDfsDownloadFile, callOptions ...callopt.Option) (r *tg.UploadFile, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DfsDownloadFile(ctx, req)
}

func (p *kDfsClient) DfsUploadDocumentFileV2(ctx context.Context, req *dfs.TLDfsUploadDocumentFileV2, callOptions ...callopt.Option) (r *tg.Document, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DfsUploadDocumentFileV2(ctx, req)
}

func (p *kDfsClient) DfsUploadGifDocumentMedia(ctx context.Context, req *dfs.TLDfsUploadGifDocumentMedia, callOptions ...callopt.Option) (r *tg.Document, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DfsUploadGifDocumentMedia(ctx, req)
}

func (p *kDfsClient) DfsUploadMp4DocumentMedia(ctx context.Context, req *dfs.TLDfsUploadMp4DocumentMedia, callOptions ...callopt.Option) (r *tg.Document, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DfsUploadMp4DocumentMedia(ctx, req)
}

func (p *kDfsClient) DfsUploadWallPaperFile(ctx context.Context, req *dfs.TLDfsUploadWallPaperFile, callOptions ...callopt.Option) (r *tg.Document, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DfsUploadWallPaperFile(ctx, req)
}

func (p *kDfsClient) DfsUploadThemeFile(ctx context.Context, req *dfs.TLDfsUploadThemeFile, callOptions ...callopt.Option) (r *tg.Document, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DfsUploadThemeFile(ctx, req)
}

func (p *kDfsClient) DfsUploadRingtoneFile(ctx context.Context, req *dfs.TLDfsUploadRingtoneFile, callOptions ...callopt.Option) (r *tg.Document, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DfsUploadRingtoneFile(ctx, req)
}

func (p *kDfsClient) DfsUploadedProfilePhoto(ctx context.Context, req *dfs.TLDfsUploadedProfilePhoto, callOptions ...callopt.Option) (r *tg.Photo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DfsUploadedProfilePhoto(ctx, req)
}
