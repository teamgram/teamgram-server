// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package model

import (
	"path"
	"strings"

	"github.com/teamgram/proto/v2/tg"
	// "github.com/teamgram/teamgram-server/model"
)

func GetFileExtName(filePath string) string {
	var ext = path.Ext(filePath)
	if ext == "" {
		ext = ".partial"
	}
	return strings.ToLower(ext)
}

/*
storage.fileUnknown#aa963b05 = storage.FileType;
storage.filePartial#40bc6f52 = storage.FileType;
storage.fileJpeg#7efe0e = storage.FileType;
storage.fileGif#cae1aadf = storage.FileType;
storage.filePng#a4f63c0 = storage.FileType;
storage.filePdf#ae1e508d = storage.FileType;
storage.fileMp3#528a0677 = storage.FileType;
storage.fileMov#4b09ebbc = storage.FileType;
storage.fileMp4#b3cea0e4 = storage.FileType;
storage.fileWebp#1081464c = storage.FileType;
*/
var (
	storageFileUnknown tg.StorageFileTypeClazz = tg.MakeTLStorageFileUnknown(&tg.TLStorageFileUnknown{})
	storageFilePartial tg.StorageFileTypeClazz = tg.MakeTLStorageFilePartial(&tg.TLStorageFilePartial{})
	storageFileJpeg    tg.StorageFileTypeClazz = tg.MakeTLStorageFileJpeg(&tg.TLStorageFileJpeg{})
	storageFileGif     tg.StorageFileTypeClazz = tg.MakeTLStorageFileGif(&tg.TLStorageFileGif{})
	storageFilePng     tg.StorageFileTypeClazz = tg.MakeTLStorageFilePng(&tg.TLStorageFilePng{})
	storageFilePdf     tg.StorageFileTypeClazz = tg.MakeTLStorageFilePdf(&tg.TLStorageFilePdf{})
	storageFileMp3     tg.StorageFileTypeClazz = tg.MakeTLStorageFileMp3(&tg.TLStorageFileMp3{})
	storageFileMov     tg.StorageFileTypeClazz = tg.MakeTLStorageFileMov(&tg.TLStorageFileMov{})
	storageFileMp4     tg.StorageFileTypeClazz = tg.MakeTLStorageFileMp4(&tg.TLStorageFileMp4{})
	storageFileWebp    tg.StorageFileTypeClazz = tg.MakeTLStorageFileWebp(&tg.TLStorageFileWebp{})
)

var (
	ImageExtensions = [5]string{".jpg", ".jpeg", ".gif", ".bmp", ".png"}
	ImageMimeTypes  = map[string]string{".jpg": "image/jpeg", ".jpeg": "image/jpeg", ".gif": "image/gif", ".bmp": "image/bmp", ".png": "image/png", ".tiff": "image/tiff"}
)

func GetStorageFileTypeConstructor(extName string) int32 {
	var (
		c uint32
	)
	switch extName {
	case ".partial":
		c = tg.ClazzID_storage_filePartial
	case ".jpeg", ".jpg":
		c = tg.ClazzID_storage_fileJpeg
	case ".gif":
		c = tg.ClazzID_storage_fileGif
	case ".png":
		c = tg.ClazzID_storage_filePng
	case ".pdf":
		c = tg.ClazzID_storage_filePdf
	case ".mp3":
		c = tg.ClazzID_storage_fileMp3
	case ".mov":
		c = tg.ClazzID_storage_fileMov
	case ".mp4":
		c = tg.ClazzID_storage_fileMp4
	case ".webp":
		c = tg.ClazzID_storage_fileWebp
	default:
		c = tg.ClazzID_storage_filePartial
	}
	return int32(c)
}

func MakeStorageFileType(c int32) tg.StorageFileTypeClazz {
	switch uint32(c) {
	case tg.ClazzID_storage_filePartial:
		return storageFilePartial
	case tg.ClazzID_storage_fileJpeg:
		return storageFileJpeg
	case tg.ClazzID_storage_fileGif:
		return storageFileGif
	case tg.ClazzID_storage_filePng:
		return storageFilePng
	case tg.ClazzID_storage_filePdf:
		return storageFilePdf
	case tg.ClazzID_storage_fileMp3:
		return storageFileMp3
	case tg.ClazzID_storage_fileMov:
		return storageFileMov
	case tg.ClazzID_storage_fileMp4:
		return storageFileMp4
	case tg.ClazzID_storage_fileWebp:
		return storageFileWebp
	default:
		return storageFileUnknown
	}
}

func IsFileExtImage(ext string) bool {
	ext = strings.ToLower(ext)
	for _, imgExt := range ImageExtensions {
		if ext == imgExt {
			return true
		}
	}
	return false
}

func GetImageMimeType(ext string) string {
	ext = strings.ToLower(ext)
	if len(ImageMimeTypes[ext]) == 0 {
		return "image"
	} else {
		return ImageMimeTypes[ext]
	}
}

func GetMimeTypeByStorageFileTyp(t *tg.StorageFileType) string {
	var ext = "partial"
	switch t.ClazzName() {
	case tg.ClazzName_storage_filePartial:
		ext = "partial"
	case tg.ClazzName_storage_fileJpeg:
		ext = "jpeg"
	case tg.ClazzName_storage_fileGif:
		ext = "gif"
	case tg.ClazzName_storage_filePng:
		ext = "png"
	case tg.ClazzName_storage_filePdf:
		ext = "pdf"
	case tg.ClazzName_storage_fileMp3:
		ext = "mp3"
	case tg.ClazzName_storage_fileMov:
		ext = "mov"
	case tg.ClazzName_storage_fileMp4:
		ext = "mp4"
	case tg.ClazzName_storage_fileWebp:
		ext = "webp"
	case tg.ClazzName_storage_fileUnknown:
		ext = "partial"
	}

	_ = ext
	return "" // mtproto.GuessMimeTypeByFileExtension(ext)
}
