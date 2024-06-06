// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package minio_util

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/logx"
)

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

// endpoint := "127.0.0.1:9000"
// accessKeyID := "Q3AM3UQ867SPQQA43P2F"
// secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
// useSSL := true

func MustNewMinioClient(c *MinioConfig) *minio.Core {
	core, err := minio.NewCore(
		c.Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(c.AccessKeyID, c.SecretAccessKey, ""),
			Secure: c.UseSSL,
		})
	if err != nil {
		logx.Must(err)
	}
	return core
}

type MinioHelper interface {
	GetFileObject(ctx context.Context, bucket, path string) (*minio.Object, error)
	GetFile(ctx context.Context, bucket, path string, offset int64, limit int32) (bytes []byte, err error)
	PutPhotoFile(ctx context.Context, path string, buf []byte) (n minio.UploadInfo, err error)
	PutPhotoFileV2(ctx context.Context, path string, r io.Reader) (n minio.UploadInfo, err error)
	PutVideoFile(ctx context.Context, path string, buf []byte) (n minio.UploadInfo, err error)
	PutDocumentFile(ctx context.Context, path string, r io.Reader) (n minio.UploadInfo, err error)
	FPutDocumentFile(ctx context.Context, path string, r string) (n minio.UploadInfo, err error)
	PutEncryptedFile(ctx context.Context, path string, r io.Reader) (n minio.UploadInfo, err error)
}
