// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package document_client

import (
	"fmt"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/service/nbfs/proto"
)

type DocumentFacade interface {
	Initialize(config string) error
	UploadPhotoFile(creatorId int64, file *mtproto.InputFile) ([]*nbfs.PhotoFileMetadata, error)
	UploadProfilePhotoFile(creatorId int64, file *mtproto.InputFile) ([]*nbfs.PhotoFileMetadata, error)
	UploadDocumentFile(creatorId int64, file *mtproto.InputFile) (*nbfs.DocumentFileMetadata, error)
	// UploadFileParts(creatorId, filePartId int64) (bool, error)
	DownloadFile(location *mtproto.InputFileLocation, offset, limit int32) (*mtproto.Upload_File, error)
}

type Instance func() DocumentFacade

var instances = make(map[string]Instance)

func Register(name string, inst Instance) {
	if inst == nil {
		panic("register instance is nil")
	}
	if _, ok := instances[name]; ok {
		panic("register called twice for instance " + name)
	}
	instances[name] = inst
}

func NewDocumentFacade(name, config string) (inst DocumentFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown instance name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	err = inst.Initialize(config)
	if err != nil {
		inst = nil
	}
	return
}
