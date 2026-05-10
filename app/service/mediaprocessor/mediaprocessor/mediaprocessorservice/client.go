/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mediaprocessorservice

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MediaProcessorProcessPhoto(ctx context.Context, req *mediaprocessor.TLMediaProcessorProcessPhoto, callOptions ...callopt.Option) (r *mediaprocessor.ProcessedPhoto, err error)
	MediaProcessorProcessGif(ctx context.Context, req *mediaprocessor.TLMediaProcessorProcessGif, callOptions ...callopt.Option) (r *mediaprocessor.ProcessedDocument, err error)
	MediaProcessorProcessMp4(ctx context.Context, req *mediaprocessor.TLMediaProcessorProcessMp4, callOptions ...callopt.Option) (r *mediaprocessor.ProcessedDocument, err error)
}

// Deprecated: prefer the generated app client helper or pkg/net/kitex.NewClient for TL-aware transport setup.
// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithCodec(codec.NewZRpcCodec(false)))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kMediaProcessorClient{
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

type kMediaProcessorClient struct {
	*kClient
}

func NewRPCMediaProcessorClient(cli client.Client) Client {
	return &kMediaProcessorClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kMediaProcessorClient) MediaProcessorProcessPhoto(ctx context.Context, req *mediaprocessor.TLMediaProcessorProcessPhoto, callOptions ...callopt.Option) (r *mediaprocessor.ProcessedPhoto, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaProcessorProcessPhoto(ctx, req)
}

func (p *kMediaProcessorClient) MediaProcessorProcessGif(ctx context.Context, req *mediaprocessor.TLMediaProcessorProcessGif, callOptions ...callopt.Option) (r *mediaprocessor.ProcessedDocument, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaProcessorProcessGif(ctx, req)
}

func (p *kMediaProcessorClient) MediaProcessorProcessMp4(ctx context.Context, req *mediaprocessor.TLMediaProcessorProcessMp4, callOptions ...callopt.Option) (r *mediaprocessor.ProcessedDocument, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MediaProcessorProcessMp4(ctx, req)
}
