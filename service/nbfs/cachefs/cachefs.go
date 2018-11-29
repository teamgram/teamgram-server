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
	"crypto/md5"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

//type CacheFileIndex struct {
//	FileId        int64
//	FileName      string
//	FileSize      int64
//	FileTotalPart int32
//	FileParts     []int32
//}
//

// serverId.creatorId.fileId.tmp

var rootDataPath = "/opt/nbfs"
var subPaths = []string{"0", "a", "b", "c", "s", "m", "x", "y"}

// var uuidgen idgen.UUIDGen

func InitCacheFS(dataPath string) error {
	if dataPath != "" {
		rootDataPath = dataPath
	}
	for _, p := range subPaths {
		err := os.MkdirAll(rootDataPath+"/"+p, 0755)
		if err != nil {
			glog.Fatal("init cache fs error: ", err)
			return err
		}
	}
	return nil
}

// 判断文件是否存在
func pathExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//获取单个文件的大小
func getFileSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

func getFileExtName(filePath string) string {
	var ext = path.Ext(filePath)
	return strings.ToLower(ext)
}

type cacheFile struct {
	creatorId int64
	fileId    int64
}

func NewCacheFile(creatorId, fileId int64) *cacheFile {
	return &cacheFile{creatorId, fileId}
}

func (f *cacheFile) WriteFilePartData(filePart int32, bytes []byte) error {
	filePath := fmt.Sprintf("%s/0/%d.%d.parts", rootDataPath, f.creatorId, f.fileId)

	exist, err := pathExists(filePath)
	if err != nil {
		glog.Errorf("pathExists error![%v]", err)
		return err
	}

	if !exist {
		err := os.Mkdir(filePath, 0755)
		if err != nil {
			glog.Errorf("mkdir failed![%v]\n", err)
			return err
		}
	}

	fileName := fmt.Sprintf("%s/%d.part", filePath, filePart)
	// exist, _ = pathExists(fileName)
	// 直接覆盖
	err = ioutil.WriteFile(fileName, bytes, 0644)
	if err != nil {
		glog.Error(err)
		return err
	}

	return nil
}

func (f *cacheFile) CheckFileParts(fileParts int32) bool {
	for i := 0; i < int(fileParts); i++ {
		filePath := fmt.Sprintf("%s/0/%d.%d.parts/%d.part", rootDataPath, f.creatorId, f.fileId, i)
		filePartInfo, err := os.Stat(filePath)
		if err != nil {
			glog.Error(err)
			return false
		}

		if filePartInfo.IsDir() {
			err = fmt.Errorf("exist dir - %s", filePath)
			return false
		}
	}
	return true
}

func (f *cacheFile) Md5Checksum(fileParts int32) (string, error) {
	md5Hash := md5.New()
	for i := 0; i < int(fileParts); i++ {
		filePath := fmt.Sprintf("%s/0/%d.%d.parts/%d.part", rootDataPath, f.creatorId, f.fileId, i)
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			err = fmt.Errorf("path not exists: %s", filePath)
			return "", err
		}
		md5Hash.Write(b)
	}
	return fmt.Sprintf("%x", md5Hash.Sum(nil)), nil
}

func (f *cacheFile) ReadFileParts(fileParts int32, cb func(int, []byte)) (err error) {
	if cb == nil {
		return nil
	}

	for i := 0; i < int(fileParts); i++ {
		filePath := fmt.Sprintf("%s/0/%d.%d.parts/%d.part", rootDataPath, f.creatorId, f.fileId, i)
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			err = fmt.Errorf("read %s error: %v", filePath, err)
			return err
		}
		cb(i, b)
	}
	return nil
}

func ReadFileOffsetData(filePath string, offset int32, limit int32) ([]byte, error) {
	fileSize := getFileSize(filePath)

	if int64(offset) > fileSize {
		limit = 0
	} else if int64(offset+limit) > fileSize {
		limit = int32(fileSize - int64(offset))
	}

	f2, err := os.Open(filePath)
	if err != nil {
		glog.Error("open ", filePath, " error: ", err)
		return nil, err
	}
	defer f2.Close()

	bytes := make([]byte, limit)
	_, err = f2.ReadAt(bytes, int64(offset))
	if err != nil {
		glog.Error("read file ", filePath, " error: ", err)
		return nil, err
	}
	return bytes, nil
}
