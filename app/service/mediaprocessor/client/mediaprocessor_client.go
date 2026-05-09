/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mediaprocessorclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor/mediaprocessorservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type MediaProcessorClient interface {
	MediaProcessorProcessPhoto(ctx context.Context, in *mediaprocessor.TLMediaProcessorProcessPhoto) (*mediaprocessor.ProcessedPhoto, error)
	MediaProcessorProcessGif(ctx context.Context, in *mediaprocessor.TLMediaProcessorProcessGif) (*mediaprocessor.ProcessedDocument, error)
	MediaProcessorProcessMp4(ctx context.Context, in *mediaprocessor.TLMediaProcessorProcessMp4) (*mediaprocessor.ProcessedDocument, error)
}

type defaultMediaProcessorClient struct {
	cli client.Client
	rpc mediaprocessorservice.Client
}

func NewMediaProcessorClient(cli client.Client) MediaProcessorClient {
	return &defaultMediaProcessorClient{
		cli: cli,
		rpc: mediaprocessorservice.NewRPCMediaProcessorClient(cli),
	}
}

// MediaProcessorProcessPhoto
// mediaProcessor.processPhoto owner_id:long object_id:string read_lease:bytes file_name:string profile:Bool = ProcessedPhoto;
func (m *defaultMediaProcessorClient) MediaProcessorProcessPhoto(ctx context.Context, in *mediaprocessor.TLMediaProcessorProcessPhoto) (*mediaprocessor.ProcessedPhoto, error) {
	return m.rpc.MediaProcessorProcessPhoto(ctx, in)
}

// MediaProcessorProcessGif
// mediaProcessor.processGif owner_id:long object_id:string read_lease:bytes file_name:string thumb_object_id:string thumb_read_lease:bytes = ProcessedDocument;
func (m *defaultMediaProcessorClient) MediaProcessorProcessGif(ctx context.Context, in *mediaprocessor.TLMediaProcessorProcessGif) (*mediaprocessor.ProcessedDocument, error) {
	return m.rpc.MediaProcessorProcessGif(ctx, in)
}

// MediaProcessorProcessMp4
// mediaProcessor.processMp4 owner_id:long object_id:string read_lease:bytes file_name:string attributes:bytes = ProcessedDocument;
func (m *defaultMediaProcessorClient) MediaProcessorProcessMp4(ctx context.Context, in *mediaprocessor.TLMediaProcessorProcessMp4) (*mediaprocessor.ProcessedDocument, error) {
	return m.rpc.MediaProcessorProcessMp4(ctx, in)
}
