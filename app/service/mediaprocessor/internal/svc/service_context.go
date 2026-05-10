// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package svc

import (
	"context"
	"fmt"
	"time"

	dfsclient "github.com/teamgram/teamgram-server/v2/app/service/dfs/client"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/internal/processor"
	"github.com/teamgram/teamgram-server/v2/pkg/media/ffmpeg2"
	"github.com/teamgram/teamgram-server/v2/pkg/media/imaging2"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type DfsClient interface {
	DfsGetFileByReadLease(ctx context.Context, in *dfs.TLDfsGetFileByReadLease) (*tg.UploadFile, error)
	DfsPutFile(ctx context.Context, in *dfs.TLDfsPutFile) (*dfs.FileFinalizedObject, error)
}

type ServiceContext struct {
	Config    config.Config
	DfsClient DfsClient
	Processor processor.MediaProcessor
}

func NewServiceContext(c config.Config) *ServiceContext {
	dfsKitexClient := dfsclient.MustNewKitexClient(c.Dfs)
	imageMagickBinary, err := imaging2.ResolveImageMagickBinary(c.ImageMagick.Binary)
	if err != nil {
		panic(fmt.Errorf("resolve ImageMagick binary: %w", err))
	}
	imagingProcessor := imaging2.NewProcessorWithProgressiveEncoder(imaging2.ImageMagickProgressiveEncoder{
		Binary:  imageMagickBinary,
		Timeout: time.Duration(c.ImageMagick.TimeoutSeconds) * time.Second,
		Quality: c.ImageMagick.Quality,
	})
	return &ServiceContext{
		Config:    c,
		DfsClient: dfsclient.NewDfsClient(dfsKitexClient),
		Processor: processor.NewWithDeps(imagingProcessor, ffmpeg2.NewProcessor()),
	}
}

func (s *ServiceContext) Close() error {
	return nil
}
