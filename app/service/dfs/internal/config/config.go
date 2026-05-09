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

package config

import (
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	kitex.RpcServerConf
	MiniHttp    rest.RestConf
	Kv          kv.KvConf
	Minio       MinioConf
	UploadSpool UploadSpoolConf
	FileObject  FileObjectConf
	ReadLease   ReadLeaseConf
	InternalAPI InternalAPIConf
	Idgen       kitex.RpcClientConf
}

type MinioConf struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	PhotosBucket    string
	VideosBucket    string
	DocumentsBucket string
	EncryptedBucket string
}

type UploadSpoolConf struct {
	RootDir         string
	NodeIDFile      string
	SegmentSize     int64
	LocalShardCount int
	PartTTLSeconds  int64
}

type FileObjectConf struct {
	MetaPrefix string
	LocalDCID  int32
}

type ReadLeaseConf struct {
	Secret     string
	TTLSeconds int64
}

type InternalAPIConf struct {
	Secret     string
	TTLSeconds int64
}
