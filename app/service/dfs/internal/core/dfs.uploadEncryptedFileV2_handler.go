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

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DfsUploadEncryptedFileV2
// dfs.uploadEncryptedFileV2 creator:long file:InputEncryptedFile = EncryptedFile;
func (c *DfsCore) DfsUploadEncryptedFileV2(in *dfs.TLDfsUploadEncryptedFileV2) (*tg.EncryptedFile, error) {
	if in == nil || in.File == nil {
		return nil, dfs.ErrDfsInvalidArgument
	}
	file, err := inputEncryptedFile(in.File)
	if err != nil {
		return nil, err
	}
	if err := checkFileParts(file.parts); err != nil {
		return nil, err
	}
	reader, err := c.uploadSessions().OpenUploadedFile(c.ctx, in.Creator, file.id)
	if err != nil {
		return nil, err
	}
	data, err := readAllSeeker(reader)
	if err != nil {
		return nil, dfs.WrapDfsStorage("read uploaded encrypted file", err)
	}
	if err := checkMD5(data, file.md5Checksum); err != nil {
		return nil, err
	}
	info, err := c.uploadSessions().LoadUploadedFileInfo(c.ctx, in.Creator, file.id)
	if err != nil {
		return nil, err
	}
	repo := c.documents()
	if repo == nil {
		return nil, dfs.WrapDfsStorage("upload encrypted file", dfs.ErrDfsStorage)
	}
	encryptedID, err := repo.NextEncryptedFileID(c.ctx)
	if err != nil {
		return nil, err
	}
	if err := c.uploadSessions().SaveObjectCacheRef(c.ctx, encryptedID, in.Creator, file.id); err != nil {
		return nil, err
	}
	if _, err := repo.SaveEncryptedObject(c.ctx, encryptedID, data); err != nil {
		return nil, err
	}
	size := info.FileSize()
	if size == 0 {
		size = int64(len(data))
	}
	return makeEncryptedFile(
		encryptedID,
		minioadapter.MakeAccessHash(storageFileTypeConstructor(".partial"), rand32()),
		size,
		0,
	), nil
}
