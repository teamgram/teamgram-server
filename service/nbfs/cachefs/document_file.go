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

package cachefs

import (
	"fmt"
	"github.com/golang/glog"
	"os"
)

type DocumentFile struct {
	fileId     int64
	accessHash int64
	*os.File
}

func CreateDocumentFile(fileId, accessHash int64) (d *DocumentFile, err error) {
	d = &DocumentFile{fileId: fileId, accessHash: accessHash}
	d.File, err = os.Create(d.ToFilePath())
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return d, nil
}

func NewDocumentFile(fileId, accessHash int64) *DocumentFile {
	return &DocumentFile{fileId: fileId, accessHash: accessHash}
}

func (f *DocumentFile) ToFilePath() string {
	return fmt.Sprintf("%s/0/%d.%d.dat", rootDataPath, f.fileId, f.accessHash)
}

func (f *DocumentFile) ToFilePath2() string {
	return fmt.Sprintf("/0/%d.%d.dat", f.fileId, f.accessHash)
}

func (f *DocumentFile) Write(b []byte) (int, error) {
	if f.File == nil {
		return 0, fmt.Errorf("file not open")
	}

	return f.File.Write(b)
}

func (f *DocumentFile) Sync() {
	if f.File != nil {
		f.File.Sync()
	}
}

func (f *DocumentFile) Close() {
	if f.File != nil {
		f.File.Close()
	}
}

func (f *DocumentFile) ReadData(offset int32, limit int32) ([]byte, error) {
	return ReadFileOffsetData(f.ToFilePath(), offset, limit)
}
