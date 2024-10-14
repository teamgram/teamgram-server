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
	"bytes"
	"context"
	"io"
	"path/filepath"

	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/encrypt"
	"github.com/zeromicro/go-zero/core/logx"
)

// BucketConfig
// endpoint := "127.0.0.1:9000"
// accessKeyID := "Q3AM3UQ867SPQQA43P2F"
// secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
// useSSL := true
type BucketConfig struct {
	Documents      string `json:",default=documents"`
	Photos         string `json:",default=photos"`
	Videos         string `json:",default=videos"`
	EncryptedFiles string `json:",default=encryptedfiles"`
}

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool

	Bucket BucketConfig
}

type MinioUtil struct {
	c     *MinioConfig
	minio *minio.Core
}

func MustNewMinioClient(c *MinioConfig) *MinioUtil {
	core, err := minio.NewCore(
		c.Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(c.AccessKeyID, c.SecretAccessKey, ""),
			Secure: c.UseSSL,
		})
	if err != nil {
		logx.Must(err)
	}

	return &MinioUtil{
		c:     c,
		minio: core,
	}
}

//func (d *Dao) Read() {
//
//}

func s3PutOptions(encrypted bool, contentType string) minio.PutObjectOptions {
	options := minio.PutObjectOptions{}
	if encrypted {
		options.ServerSideEncryption = encrypt.NewSSE()
	}
	options.ContentType = contentType

	return options
}

func (m *MinioUtil) GetFileObject(ctx context.Context, bucket, path string) (*minio.Object, error) {
	object, err := m.minio.Client.GetObject(ctx, bucket, path, minio.GetObjectOptions{})
	if err != nil {
		logx.WithContext(ctx).Errorf("GetFileObject error: %v")
		return nil, err
	}

	return object, nil
}

func (m *MinioUtil) GetFile(ctx context.Context, bucket, path string, offset int64, limit int32) (bytes []byte, err error) {
	var (
		object *minio.Object
		n      int
	)

	object, err = m.minio.Client.GetObject(ctx, bucket, path, minio.GetObjectOptions{})
	if err != nil {
		logx.WithContext(ctx).Errorf("GetFile error: %v")
		return
	}
	defer object.Close()

	bytes = make([]byte, limit)
	n, err = object.ReadAt(bytes, offset)
	//if err != nil {
	//	// return
	//}
	bytes = bytes[:n]
	if n > 0 {
		err = nil
	} else {
		logx.WithContext(ctx).Errorf("GetFile (%s) error: %v", path, err)
	}
	return
}

func (m *MinioUtil) GetPhotoFile(ctx context.Context, path string, offset int64, limit int32) (bytes []byte, err error) {
	return m.GetFile(ctx, m.c.Bucket.Photos, path, offset, limit)
}

func (m *MinioUtil) PutPhotoFile(ctx context.Context, path string, buf []byte) (n minio.UploadInfo, err error) {
	var (
		contentType string
	)

	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = m.minio.Client.PutObject(ctx, m.c.Bucket.Photos, path, bytes.NewReader(buf), int64(len(buf)), options)
	if err != nil {
		logx.WithContext(ctx).Errorf("PutPhotoFile (%s) error: %v", path, err)
	}
	return
}

func (m *MinioUtil) PutPhotoFileV2(ctx context.Context, path string, r io.Reader) (n minio.UploadInfo, err error) {
	var (
		contentType string
	)

	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = m.minio.Client.PutObject(ctx, m.c.Bucket.Photos, path, r, -1, options)
	if err != nil {
		logx.Errorf("PutPhotoFile (%s) error: %v", path, err)
	}
	return
}

func (m *MinioUtil) GetVideoFile(ctx context.Context, path string, offset int64, limit int32) (bytes []byte, err error) {
	return m.GetFile(ctx, m.c.Bucket.Videos, path, offset, limit)
}

func (m *MinioUtil) PutVideoFile(ctx context.Context, path string, buf []byte) (n minio.UploadInfo, err error) {
	var (
		contentType string
	)

	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = m.minio.Client.PutObject(ctx, m.c.Bucket.Videos, path, bytes.NewReader(buf), int64(len(buf)), options)
	if err != nil {
		logx.WithContext(ctx).Errorf("PutVideoFile (%s) error: %v", path, err)
	}

	return
}

func (m *MinioUtil) PutVideoFileV2(ctx context.Context, path string, r io.Reader) (n minio.UploadInfo, err error) {
	var (
		contentType string
	)

	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = m.minio.Client.PutObject(ctx, m.c.Bucket.Videos, path, r, -1, options)
	if err != nil {
		logx.Errorf("PutVideoFileV2 (%s) error: %v", path, err)
	}
	return
}

func (m *MinioUtil) GetDocumentFile(ctx context.Context, path string, offset int64, limit int32) (bytes []byte, err error) {
	return m.GetFile(ctx, m.c.Bucket.Documents, path, offset, limit)
}

func (m *MinioUtil) PutDocumentFile(ctx context.Context, path string, r io.Reader) (n minio.UploadInfo, err error) {
	var (
		contentType string
	)

	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = m.minio.Client.PutObject(ctx, m.c.Bucket.Documents, path, r, -1, options)
	if err != nil {
		logx.WithContext(ctx).Errorf("PutDocumentFile (%s) error: %v", path, err)
	}

	return
}

func (m *MinioUtil) FPutDocumentFile(ctx context.Context, path string, r string) (n minio.UploadInfo, err error) {
	var (
		contentType string
	)

	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = m.minio.Client.FPutObject(ctx, m.c.Bucket.Documents, path, r, options)
	if err != nil {
		logx.WithContext(ctx).Errorf("PutDocumentFile (%s) error: %v", path, err)
	}

	return
}

func (m *MinioUtil) GetEncryptedFile(ctx context.Context, path string, offset int64, limit int32) (bytes []byte, err error) {
	return m.GetFile(ctx, m.c.Bucket.EncryptedFiles, path, offset, limit)
}

func (m *MinioUtil) PutEncryptedFile(ctx context.Context, path string, r io.Reader) (n minio.UploadInfo, err error) {
	options := s3PutOptions(false, "binary/octet-stream")
	n, err = m.minio.Client.PutObject(ctx, m.c.Bucket.EncryptedFiles, path, r, -1, options)
	if err != nil {
		logx.WithContext(ctx).Errorf("PutEncryptedFile (%s) error: %v", path, err)
	}

	return
}
