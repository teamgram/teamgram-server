/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/internal/core"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// MediaProcessorProcessPhoto
// mediaProcessor.processPhoto owner_id:long object_id:string read_lease:bytes file_name:string profile:Bool = ProcessedPhoto;
func (s *Service) MediaProcessorProcessPhoto(ctx context.Context, request *mediaprocessor.TLMediaProcessorProcessPhoto) (*mediaprocessor.ProcessedPhoto, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("mediaProcessor.processPhoto - request: %s", request)

	r, err := c.MediaProcessorProcessPhoto(request)
	if err != nil {
		c.Logger.Errorf("mediaProcessor.processPhoto - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("mediaProcessor.processPhoto - reply: %s", r)
	return r, err
}

// MediaProcessorProcessGif
// mediaProcessor.processGif owner_id:long object_id:string read_lease:bytes file_name:string thumb_object_id:string thumb_read_lease:bytes = ProcessedDocument;
func (s *Service) MediaProcessorProcessGif(ctx context.Context, request *mediaprocessor.TLMediaProcessorProcessGif) (*mediaprocessor.ProcessedDocument, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("mediaProcessor.processGif - request: %s", request)

	r, err := c.MediaProcessorProcessGif(request)
	if err != nil {
		c.Logger.Errorf("mediaProcessor.processGif - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("mediaProcessor.processGif - reply: %s", r)
	return r, err
}

// MediaProcessorProcessMp4
// mediaProcessor.processMp4 owner_id:long object_id:string read_lease:bytes file_name:string attributes:bytes = ProcessedDocument;
func (s *Service) MediaProcessorProcessMp4(ctx context.Context, request *mediaprocessor.TLMediaProcessorProcessMp4) (*mediaprocessor.ProcessedDocument, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("mediaProcessor.processMp4 - request: %s", request)

	r, err := c.MediaProcessorProcessMp4(request)
	if err != nil {
		c.Logger.Errorf("mediaProcessor.processMp4 - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("mediaProcessor.processMp4 - reply: %s", r)
	return r, err
}
