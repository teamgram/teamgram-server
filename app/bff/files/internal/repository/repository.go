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

	"github.com/teamgram/teamgram-server/v2/app/bff/files/internal/config"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	dfsclient "github.com/teamgram/teamgram-server/v2/app/service/dfs/client"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	mediaclient "github.com/teamgram/teamgram-server/v2/app/service/media/client"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type DfsFilesClient interface {
	DfsGetFileByReadLease(ctx context.Context, in *dfs.TLDfsGetFileByReadLease) (*tg.UploadFile, error)
	DfsGetFileHashesByReadLease(ctx context.Context, in *dfs.TLDfsGetFileHashesByReadLease) (*dfs.VectorFileHash, error)
	DfsWriteFilePartData(ctx context.Context, in *dfs.TLDfsWriteFilePartData) (*tg.Bool, error)
	DfsDownloadFile(ctx context.Context, in *dfs.TLDfsDownloadFile) (*tg.UploadFile, error)
}

type MediaFilesClient interface {
	MediaUploadPhotoFile(ctx context.Context, in *media.TLMediaUploadPhotoFile) (*tg.Photo, error)
	MediaGetPhoto(ctx context.Context, in *media.TLMediaGetPhoto) (*tg.Photo, error)
	MediaGetPhotoSizeList(ctx context.Context, in *media.TLMediaGetPhotoSizeList) (*media.PhotoSizeList, error)
	MediaUploadedDocumentMedia(ctx context.Context, in *media.TLMediaUploadedDocumentMedia) (*tg.MessageMedia, error)
	MediaGetDocument(ctx context.Context, in *media.TLMediaGetDocument) (*tg.Document, error)
	MediaResolveFileLocation(ctx context.Context, in *media.TLMediaResolveFileLocation) (*media.MediaResolvedFileObject, error)
}

// Repository is the dependency container for repository instances.
type Repository struct {
	DfsClient   DfsFilesClient
	MediaClient MediaFilesClient
	UserClient  userclient.UserClient
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config) *Repository {
	return &Repository{
		DfsClient:   dfsclient.NewDfsClient(dfsclient.MustNewKitexClient(c.DfsClient)),
		MediaClient: mediaclient.NewMediaClient(mediaclient.MustNewKitexClient(c.MediaClient)),
		UserClient:  userclient.NewUserClient(userclient.MustNewKitexClient(c.UserClient)),
	}
}
