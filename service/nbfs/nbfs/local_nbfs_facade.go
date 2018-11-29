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

package nbfs_client

import (
	"crypto/md5"
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/service/idgen/client"
	"github.com/nebula-chat/chatengine/service/nbfs/cachefs"
	"github.com/nebula-chat/chatengine/service/nbfs/proto"
	"hash"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

type localNbfsFacade struct {
	idgen.UUIDGen
}

func localNbfsFacadeInstance() NbfsFacade {
	return &localNbfsFacade{}
}

func NewLocalNbfsFacade(dbName string) (*localNbfsFacade, error) {
	var err error

	facade := &localNbfsFacade{}
	// facade.ContactModel, err = core.InitContactModel(dbName)

	return facade, err
}

func getFileExtName(filePath string) string {
	var ext = path.Ext(filePath)
	return strings.ToLower(ext)
}

func getStorageFileTypeConstructor(extName string) int32 {
	var c mtproto.TLConstructor
	switch extName {
	case ".partial":
		c = mtproto.TLConstructor_CRC32_storage_filePartial
	case ".jpeg", ".jpg":
		c = mtproto.TLConstructor_CRC32_storage_fileJpeg
	case ".gif":
		c = mtproto.TLConstructor_CRC32_storage_fileGif
	case ".png":
		c = mtproto.TLConstructor_CRC32_storage_filePng
	case ".pdf":
		c = mtproto.TLConstructor_CRC32_storage_filePdf
	case ".mp3":
		c = mtproto.TLConstructor_CRC32_storage_fileMp3
	case ".mov":
		c = mtproto.TLConstructor_CRC32_storage_fileMov
	case ".mp4":
		c = mtproto.TLConstructor_CRC32_storage_fileMp4
	case ".webp":
		c = mtproto.TLConstructor_CRC32_storage_fileWebp
	default:
		// fileType.Constructor = mtproto.TLConstructor_CRC32_storage_fileUnknown
		c = mtproto.TLConstructor_CRC32_storage_filePartial
	}
	return int32(c)
}

func makeStorageFileType(c int32) *mtproto.Storage_FileType {
	fileType := &mtproto.Storage_FileType{Data2: &mtproto.Storage_FileType_Data{}}

	switch mtproto.TLConstructor(c) {
	case mtproto.TLConstructor_CRC32_storage_filePartial,
		mtproto.TLConstructor_CRC32_storage_fileJpeg,
		mtproto.TLConstructor_CRC32_storage_fileGif,
		mtproto.TLConstructor_CRC32_storage_filePng,
		mtproto.TLConstructor_CRC32_storage_filePdf,
		mtproto.TLConstructor_CRC32_storage_fileMp3,
		mtproto.TLConstructor_CRC32_storage_fileMov,
		mtproto.TLConstructor_CRC32_storage_fileMp4,
		mtproto.TLConstructor_CRC32_storage_fileWebp:

		fileType.Constructor = mtproto.TLConstructor(c)
	default:
		fileType.Constructor = mtproto.TLConstructor_CRC32_storage_filePartial
	}
	return fileType
}

func (c *localNbfsFacade) Initialize(config string) error {
	glog.Info("localNbfsFacade - Initialize config: ", config)

	var err error
	c.UUIDGen, err = idgen.NewUUIDGen("snowflake", config)
	if err != nil {
		glog.Fatal("uuidgen init error: ", err)
	}

	// dbName := config
	// c.ContactModel, err = core.InitContactModel(dbName)

	return err
}

func (c *localNbfsFacade) UploadPhotoFile(creatorId int64, file *mtproto.InputFile) (fileMDList []*nbfs.PhotoFileMetadata, err error) {
	var (
		inputFile = file.GetData2()
		md5Hash   hash.Hash
		fileSize  = int64(0)
	)

	var cacheData []byte

	if inputFile.GetMd5Checksum() != "" {
		md5Hash = md5.New()
	}
	cacheFile := cachefs.NewCacheFile(creatorId, inputFile.GetId())
	err = cacheFile.ReadFileParts(inputFile.GetParts(), func(i int, bytes []byte) {
		if md5Hash != nil {
			md5Hash.Write(bytes)
		}
		cacheData = append(cacheData, bytes...)
		fileSize += int64(len(bytes))
	})
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	if md5Hash != nil {
		if fmt.Sprintf("%x", md5Hash.Sum(nil)) != inputFile.GetMd5Checksum() {
			err = fmt.Errorf("invalid md5")
			return nil, err
		}
	}

	photoId, _ := c.UUIDGen.GetUUID()
	// accessHash := rand.Int63()
	photoFile2 := cachefs.NewPhotoFile(photoId, 0, 0)

	ext := getFileExtName(inputFile.GetName())
	extType := getStorageFileTypeConstructor(ext)

	err = cachefs.DoUploadedPhotoFile(photoFile2, ext, cacheData, false, func(pi *cachefs.PhotoInfo) {
		// 有点难理解，主要是为了不在这里引入snowflake
		// ext := getFileExtName(inputFile.GetName())
		// extType := getStorageFileTypeConstructor(ext)
		secretId := int64(extType)<<32 | int64(rand.Uint32())

		srcFile := cachefs.NewPhotoFile(photoId, pi.LocalId, 0)
		dstFile := cachefs.NewPhotoFile(photoId, pi.LocalId, secretId)
		os.Rename(srcFile.ToFilePath(), dstFile.ToFilePath())

		fileMD := &nbfs.PhotoFileMetadata{
			FileId:   inputFile.GetId(),
			PhotoId:  photoId,
			DcId:     2,
			VolumeId: photoId,
			LocalId:  pi.LocalId,
			SecretId: secretId,
			Width:    pi.Width,
			Height:   pi.Height,
			FileSize: int32(pi.FileSize),
			FilePath: dstFile.ToFilePath2(),
			Ext:      ext,
		}
		fileMDList = append(fileMDList, fileMD)
	})
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return
}

func (c *localNbfsFacade) UploadProfilePhotoFile(creatorId int64, file *mtproto.InputFile) (fileMDList []*nbfs.PhotoFileMetadata, err error) {
	var (
		inputFile = file.GetData2()
		md5Hash   hash.Hash
		fileSize  = int64(0)
	)

	var cacheData []byte

	if inputFile.GetMd5Checksum() != "" {
		md5Hash = md5.New()
	}
	cacheFile := cachefs.NewCacheFile(creatorId, inputFile.GetId())
	err = cacheFile.ReadFileParts(inputFile.GetParts(), func(i int, bytes []byte) {
		if md5Hash != nil {
			md5Hash.Write(bytes)
		}
		cacheData = append(cacheData, bytes...)
		fileSize += int64(len(bytes))
	})
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	if md5Hash != nil {
		if fmt.Sprintf("%x", md5Hash.Sum(nil)) != inputFile.GetMd5Checksum() {
			err = fmt.Errorf("invalid md5")
			return nil, err
		}
	}

	photoId, _ := c.UUIDGen.GetUUID()
	// accessHash := rand.Int63()
	photoFile2 := cachefs.NewPhotoFile(photoId, 0, 0)

	ext := getFileExtName(inputFile.GetName())
	extType := getStorageFileTypeConstructor(ext)

	err = cachefs.DoUploadedPhotoFile(photoFile2, ext, cacheData, false, func(pi *cachefs.PhotoInfo) {
		// 有点难理解，主要是为了不在这里引入snowflake
		secretId := int64(extType)<<32 | int64(rand.Uint32())

		srcFile := cachefs.NewPhotoFile(photoId, pi.LocalId, 0)
		dstFile := cachefs.NewPhotoFile(photoId, pi.LocalId, secretId)
		os.Rename(srcFile.ToFilePath(), dstFile.ToFilePath())

		fileMD := &nbfs.PhotoFileMetadata{
			FileId:   inputFile.GetId(),
			PhotoId:  photoId,
			DcId:     2,
			VolumeId: photoId,
			LocalId:  pi.LocalId,
			SecretId: secretId,
			Width:    pi.Width,
			Height:   pi.Height,
			FileSize: int32(pi.FileSize),
			FilePath: dstFile.ToFilePath2(),
			Ext:      ext,
		}
		fileMDList = append(fileMDList, fileMD)
	})
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return
}

func (c *localNbfsFacade) UploadDocumentFile(creatorId int64, file *mtproto.InputFile) (fileMD *nbfs.DocumentFileMetadata, err error) {
	var (
		inputFile    = file.GetData2()
		fileSize     = int64(0)
		documentFile *cachefs.DocumentFile
	)

	documentId, _ := c.UUIDGen.GetUUID()

	// 有点难理解，主要是为了不在这里引入snowflake
	ext := getFileExtName(inputFile.GetName())
	extType := getStorageFileTypeConstructor(ext)
	accessHash := int64(extType)<<32 | int64(rand.Uint32())

	documentFile, err = cachefs.CreateDocumentFile(documentId, accessHash)
	if err != nil {
		return nil, err
	}
	defer documentFile.Close()

	// var cacheData []byte
	cacheFile := cachefs.NewCacheFile(creatorId, inputFile.GetId())
	err = cacheFile.ReadFileParts(inputFile.GetParts(), func(i int, bytes []byte) {
		// cacheData = append(cacheData, bytes...)
		fileSize += int64(len(bytes))
		documentFile.Write(bytes)
		documentFile.Sync()
	})
	if err != nil {
		glog.Error(err)
		return
	}

	fileMD = &nbfs.DocumentFileMetadata{
		FileId:           inputFile.GetId(),
		DocumentId:       documentId,
		AccessHash:       accessHash,
		DcId:             2,
		FileSize:         int32(fileSize),
		FilePath:         documentFile.ToFilePath2(),
		UploadedFileName: inputFile.GetName(),
		Ext:              ext,
		// MimeType:         inputFile.GetName(),
	}
	return
}

//func (c *localNbfsFacade) UploadFileParts(creatorId, filePartId int64) (bool, error) {
//	return false, nil
//}
//

func (c *localNbfsFacade) DownloadFile(location *mtproto.InputFileLocation, offset, limit int32) (file *mtproto.Upload_File, err error) {
	var (
		// uploadFile *mtproto.Upload_File
		// err        error
		bytes []byte
		sType int32
	)

	switch location.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputFileLocation:
		fileLocation := location.To_InputFileLocation()
		file := cachefs.NewPhotoFile(fileLocation.GetVolumeId(), fileLocation.GetLocalId(), fileLocation.GetSecret())
		bytes, err = file.ReadData(offset, limit)
		sType = int32(fileLocation.GetSecret() >> 32)
	case mtproto.TLConstructor_CRC32_inputFileLocationLayer86:
		fileLocation := location.To_InputFileLocationLayer86()
		file := cachefs.NewPhotoFile(fileLocation.GetVolumeId(), fileLocation.GetLocalId(), fileLocation.GetSecret())
		bytes, err = file.ReadData(offset, limit)
		sType = int32(fileLocation.GetSecret() >> 32)
	case mtproto.TLConstructor_CRC32_inputEncryptedFileLocation:
	case mtproto.TLConstructor_CRC32_inputDocumentFileLocation:
		fileLocation := location.To_InputDocumentFileLocation()
		file := cachefs.NewDocumentFile(fileLocation.GetId(), fileLocation.GetAccessHash())
		bytes, err = file.ReadData(offset, limit)
		sType = int32(fileLocation.GetAccessHash() >> 32)
	case mtproto.TLConstructor_CRC32_inputDocumentFileLocationLayer11:
		fileLocation := location.To_InputDocumentFileLocation()
		file := cachefs.NewDocumentFile(fileLocation.GetId(), fileLocation.GetAccessHash())
		bytes, err = file.ReadData(offset, limit)
		sType = int32(fileLocation.GetAccessHash() >> 32)
	case mtproto.TLConstructor_CRC32_inputDocumentFileLocationLayer86:
		fileLocation := location.To_InputDocumentFileLocationLayer86()
		file := cachefs.NewDocumentFile(fileLocation.GetId(), fileLocation.GetAccessHash())
		bytes, err = file.ReadData(offset, limit)
		sType = int32(fileLocation.GetAccessHash() >> 32)
	default:
		err = fmt.Errorf("invalid InputFileLocation type: {%v}", location)
		glog.Error(err)
		return nil, err
	}

	uploadFile := &mtproto.TLUploadFile{Data2: &mtproto.Upload_File_Data{
		Type:  makeStorageFileType(sType),
		Mtime: int32(time.Now().Unix()),
		Bytes: bytes,
	}}

	file = uploadFile.To_Upload_File()
	return
}

func init() {
	Register("local", localNbfsFacadeInstance)
}
