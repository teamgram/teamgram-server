/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package mediaservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MediaUploadPhotoFile(ctx context.Context, req *media.TLMediaUploadPhotoFile, callOptions ...callopt.Option) (r *tg.Photo, err error)
	MediaUploadProfilePhotoFile(ctx context.Context, req *media.TLMediaUploadProfilePhotoFile, callOptions ...callopt.Option) (r *tg.Photo, err error)
	MediaGetPhoto(ctx context.Context, req *media.TLMediaGetPhoto, callOptions ...callopt.Option) (r *tg.Photo, err error)
	MediaGetPhotoSizeList(ctx context.Context, req *media.TLMediaGetPhotoSizeList, callOptions ...callopt.Option) (r *media.PhotoSizeList, err error)
	MediaGetPhotoSizeListList(ctx context.Context, req *media.TLMediaGetPhotoSizeListList, callOptions ...callopt.Option) (r *media.VectorPhotoSizeList, err error)
	MediaGetVideoSizeList(ctx context.Context, req *media.TLMediaGetVideoSizeList, callOptions ...callopt.Option) (r *media.VideoSizeList, err error)
	MediaUploadedDocumentMedia(ctx context.Context, req *media.TLMediaUploadedDocumentMedia, callOptions ...callopt.Option) (r *tg.MessageMedia, err error)
	MediaGetDocument(ctx context.Context, req *media.TLMediaGetDocument, callOptions ...callopt.Option) (r *tg.Document, err error)
	MediaGetDocumentList(ctx context.Context, req *media.TLMediaGetDocumentList, callOptions ...callopt.Option) (r *media.VectorDocument, err error)
	MediaUploadEncryptedFile(ctx context.Context, req *media.TLMediaUploadEncryptedFile, callOptions ...callopt.Option) (r *tg.EncryptedFile, err error)
	MediaGetEncryptedFile(ctx context.Context, req *media.TLMediaGetEncryptedFile, callOptions ...callopt.Option) (r *tg.EncryptedFile, err error)
	MediaUploadWallPaperFile(ctx context.Context, req *media.TLMediaUploadWallPaperFile, callOptions ...callopt.Option) (r *tg.Document, err error)
	MediaUploadThemeFile(ctx context.Context, req *media.TLMediaUploadThemeFile, callOptions ...callopt.Option) (r *tg.Document, err error)
	MediaUploadStickerFile(ctx context.Context, req *media.TLMediaUploadStickerFile, callOptions ...callopt.Option) (r *tg.Document, err error)
	MediaUploadRingtoneFile(ctx context.Context, req *media.TLMediaUploadRingtoneFile, callOptions ...callopt.Option) (r *tg.Document, err error)
	MediaUploadedProfilePhoto(ctx context.Context, req *media.TLMediaUploadedProfilePhoto, callOptions ...callopt.Option) (r *tg.Photo, err error)
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
	return &kMediaClient{
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

type kMediaClient struct {
	*kClient
}

func NewRPCMediaClient(cli client.Client) Client {
	return &kMediaClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kMediaClient) MediaUploadPhotoFile(ctx context.Context, req *media.TLMediaUploadPhotoFile, callOptions ...callopt.Option) (r *tg.Photo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaUploadPhotoFile(ctx, req)
}

func (p *kMediaClient) MediaUploadProfilePhotoFile(ctx context.Context, req *media.TLMediaUploadProfilePhotoFile, callOptions ...callopt.Option) (r *tg.Photo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaUploadProfilePhotoFile(ctx, req)
}

func (p *kMediaClient) MediaGetPhoto(ctx context.Context, req *media.TLMediaGetPhoto, callOptions ...callopt.Option) (r *tg.Photo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaGetPhoto(ctx, req)
}

func (p *kMediaClient) MediaGetPhotoSizeList(ctx context.Context, req *media.TLMediaGetPhotoSizeList, callOptions ...callopt.Option) (r *media.PhotoSizeList, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaGetPhotoSizeList(ctx, req)
}

func (p *kMediaClient) MediaGetPhotoSizeListList(ctx context.Context, req *media.TLMediaGetPhotoSizeListList, callOptions ...callopt.Option) (r *media.VectorPhotoSizeList, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaGetPhotoSizeListList(ctx, req)
}

func (p *kMediaClient) MediaGetVideoSizeList(ctx context.Context, req *media.TLMediaGetVideoSizeList, callOptions ...callopt.Option) (r *media.VideoSizeList, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaGetVideoSizeList(ctx, req)
}

func (p *kMediaClient) MediaUploadedDocumentMedia(ctx context.Context, req *media.TLMediaUploadedDocumentMedia, callOptions ...callopt.Option) (r *tg.MessageMedia, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaUploadedDocumentMedia(ctx, req)
}

func (p *kMediaClient) MediaGetDocument(ctx context.Context, req *media.TLMediaGetDocument, callOptions ...callopt.Option) (r *tg.Document, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaGetDocument(ctx, req)
}

func (p *kMediaClient) MediaGetDocumentList(ctx context.Context, req *media.TLMediaGetDocumentList, callOptions ...callopt.Option) (r *media.VectorDocument, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaGetDocumentList(ctx, req)
}

func (p *kMediaClient) MediaUploadEncryptedFile(ctx context.Context, req *media.TLMediaUploadEncryptedFile, callOptions ...callopt.Option) (r *tg.EncryptedFile, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaUploadEncryptedFile(ctx, req)
}

func (p *kMediaClient) MediaGetEncryptedFile(ctx context.Context, req *media.TLMediaGetEncryptedFile, callOptions ...callopt.Option) (r *tg.EncryptedFile, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaGetEncryptedFile(ctx, req)
}

func (p *kMediaClient) MediaUploadWallPaperFile(ctx context.Context, req *media.TLMediaUploadWallPaperFile, callOptions ...callopt.Option) (r *tg.Document, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaUploadWallPaperFile(ctx, req)
}

func (p *kMediaClient) MediaUploadThemeFile(ctx context.Context, req *media.TLMediaUploadThemeFile, callOptions ...callopt.Option) (r *tg.Document, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaUploadThemeFile(ctx, req)
}

func (p *kMediaClient) MediaUploadStickerFile(ctx context.Context, req *media.TLMediaUploadStickerFile, callOptions ...callopt.Option) (r *tg.Document, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaUploadStickerFile(ctx, req)
}

func (p *kMediaClient) MediaUploadRingtoneFile(ctx context.Context, req *media.TLMediaUploadRingtoneFile, callOptions ...callopt.Option) (r *tg.Document, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaUploadRingtoneFile(ctx, req)
}

func (p *kMediaClient) MediaUploadedProfilePhoto(ctx context.Context, req *media.TLMediaUploadedProfilePhoto, callOptions ...callopt.Option) (r *tg.Photo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaUploadedProfilePhoto(ctx, req)
}
