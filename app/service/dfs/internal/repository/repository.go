// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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

package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/ffmpeg2"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/imaging2"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/objectstore"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/rpc"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/spool"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/xkv"
	idgenclient "github.com/teamgram/teamgram-server/v2/app/service/idgen/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	kv                  kv.ExtStore
	uploadStateModel    uploadStateModel
	objectStore         minioadapter.ObjectStore
	idgen               rpc.IDGenerator
	imaging             imaging2.Processor
	ffmpeg              ffmpeg2.Processor
	documentsBucket     string
	manifestKeys        objectstore.ManifestKeys
	readLeaseSecret     string
	readLeaseTTLSeconds int64
	localDCID           int32
	uploadNodeMu        sync.RWMutex
	uploadNodeDraining  bool
	uploadDrainReason   string
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config) *Repository {
	kv2 := kv.NewStore(c.Kv)
	var objectStore minioadapter.ObjectStore
	if c.Minio.Endpoint != "" {
		store, err := minioadapter.NewObjectStore(c.Minio)
		if err != nil {
			panic(fmt.Errorf("new dfs minio object store: %w", err))
		}
		objectStore = store
	}
	var idgen rpc.IDGenerator
	if hasRPCClientConfig(c.Idgen) {
		idgen = rpc.NewIDGenClient(idgenclient.NewIdgenClient(idgenclient.MustNewKitexClient(c.Idgen)))
	}
	uploadState := uploadStateModel(xkv.NewUploadStateModel(kv2))
	if c.UploadSpool.RootDir != "" {
		model, err := spool.NewUploadStateModel(c.UploadSpool)
		if err != nil {
			panic(fmt.Errorf("new dfs upload spool: %w", err))
		}
		if err := model.ScanOnStart(context.Background(), time.Now()); err != nil {
			panic(fmt.Errorf("scan dfs upload spool on start: %w", err))
		}
		uploadState = model
	}

	return &Repository{
		kv:                  kv2,
		uploadStateModel:    uploadState,
		objectStore:         objectStore,
		idgen:               idgen,
		imaging:             imaging2.NewProcessor(),
		ffmpeg:              ffmpeg2.NewProcessor(),
		documentsBucket:     c.Minio.DocumentsBucket,
		manifestKeys:        objectstore.ManifestKeys{MetaPrefix: c.FileObject.MetaPrefix},
		readLeaseSecret:     c.ReadLease.Secret,
		readLeaseTTLSeconds: c.ReadLease.TTLSeconds,
		localDCID:           c.FileObject.LocalDCID,
	}
}

func hasRPCClientConfig(c kitex.RpcClientConf) bool {
	if c.DestService == "" {
		return false
	}
	return len(c.Endpoints) > 0 || c.Target != "" || c.HasEtcd()
}
