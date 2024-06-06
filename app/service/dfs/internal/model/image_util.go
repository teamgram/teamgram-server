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

	"github.com/teamgram/proto/mtproto"
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
	storageFileUnknown = mtproto.MakeTLStorageFileUnknown(nil).To_Storage_FileType()
	storageFilePartial = mtproto.MakeTLStorageFilePartial(nil).To_Storage_FileType()
	storageFileJpeg    = mtproto.MakeTLStorageFileJpeg(nil).To_Storage_FileType()
	storageFileGif     = mtproto.MakeTLStorageFileGif(nil).To_Storage_FileType()
	storageFilePng     = mtproto.MakeTLStorageFilePng(nil).To_Storage_FileType()
	storageFilePdf     = mtproto.MakeTLStorageFilePdf(nil).To_Storage_FileType()
	storageFileMp3     = mtproto.MakeTLStorageFileMp3(nil).To_Storage_FileType()
	storageFileMov     = mtproto.MakeTLStorageFileMov(nil).To_Storage_FileType()
	storageFileMp4     = mtproto.MakeTLStorageFileMp4(nil).To_Storage_FileType()
	storageFileWebp    = mtproto.MakeTLStorageFileWebp(nil).To_Storage_FileType()
)

var (
	ImageExtensions = [5]string{".jpg", ".jpeg", ".gif", ".bmp", ".png"}
	ImageMimeTypes  = map[string]string{".jpg": "image/jpeg", ".jpeg": "image/jpeg", ".gif": "image/gif", ".bmp": "image/bmp", ".png": "image/png", ".tiff": "image/tiff"}
)

func GetStorageFileTypeConstructor(extName string) int32 {
	var c mtproto.TLConstructor
	switch extName {
	case ".partial":
		c = mtproto.CRC32_storage_filePartial
	case ".jpeg", ".jpg":
		c = mtproto.CRC32_storage_fileJpeg
	case ".gif":
		c = mtproto.CRC32_storage_fileGif
	case ".png":
		c = mtproto.CRC32_storage_filePng
	case ".pdf":
		c = mtproto.CRC32_storage_filePdf
	case ".mp3":
		c = mtproto.CRC32_storage_fileMp3
	case ".mov":
		c = mtproto.CRC32_storage_fileMov
	case ".mp4":
		c = mtproto.CRC32_storage_fileMp4
	case ".webp":
		c = mtproto.CRC32_storage_fileWebp
	default:
		c = mtproto.CRC32_storage_filePartial
	}
	return int32(c)
}

func MakeStorageFileType(c int32) *mtproto.Storage_FileType {
	switch mtproto.TLConstructor(c) {
	case mtproto.CRC32_storage_filePartial:
		return storageFilePartial
	case mtproto.CRC32_storage_fileJpeg:
		return storageFileJpeg
	case mtproto.CRC32_storage_fileGif:
		return storageFileGif
	case mtproto.CRC32_storage_filePng:
		return storageFilePng
	case mtproto.CRC32_storage_filePdf:
		return storageFilePdf
	case mtproto.CRC32_storage_fileMp3:
		return storageFileMp3
	case mtproto.CRC32_storage_fileMov:
		return storageFileMov
	case mtproto.CRC32_storage_fileMp4:
		return storageFileMp4
	case mtproto.CRC32_storage_fileWebp:
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

func GetMimeTypeByStorageFileTyp(t *mtproto.Storage_FileType) string {
	var ext = "partial"
	switch t.GetPredicateName() {
	case mtproto.Predicate_storage_filePartial:
		ext = "partial"
	case mtproto.Predicate_storage_fileJpeg:
		ext = "jpeg"
	case mtproto.Predicate_storage_fileGif:
		ext = "gif"
	case mtproto.Predicate_storage_filePng:
		ext = "png"
	case mtproto.Predicate_storage_filePdf:
		ext = "pdf"
	case mtproto.Predicate_storage_fileMp3:
		ext = "mp3"
	case mtproto.Predicate_storage_fileMov:
		ext = "mov"
	case mtproto.Predicate_storage_fileMp4:
		ext = "mp4"
	case mtproto.Predicate_storage_fileWebp:
		ext = "webp"
	case mtproto.Predicate_storage_fileUnknown:
		ext = "partial"
	}

	return mtproto.GuessMimeTypeByFileExtension(ext)
}
