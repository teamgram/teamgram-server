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
	"log"

	"github.com/minio/minio-go"
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
	core, err := minio.NewCore(c.Endpoint, c.AccessKeyID, c.SecretAccessKey, c.UseSSL)
	if err != nil {
		log.Fatal(err)
	}
	return core
}
