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

package photo

import (
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/service/document/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/service/nbfs/proto"
)

const (
	kPhotoSizeOriginalType = "0" // client upload original photo
	kPhotoSizeSmallType    = "s"
	kPhotoSizeMediumType   = "m"
	kPhotoSizeXLargeType   = "x"
	kPhotoSizeYLargeType   = "y"
	kPhotoSizeAType        = "a"
	kPhotoSizeBType        = "b"
	kPhotoSizeCType        = "c"

	kPhotoSizeOriginalSize = 0 // client upload original photo
	kPhotoSizeSmallSize    = 90
	kPhotoSizeMediumSize   = 320
	kPhotoSizeXLargeSize   = 800
	kPhotoSizeYLargeSize   = 1280
	kPhotoSizeASize        = 160
	kPhotoSizeBSize        = 320
	kPhotoSizeCSize        = 640

	kPhotoSizeAIndex = 4
)

var sizeList = []int{
	kPhotoSizeOriginalSize,
	kPhotoSizeSmallSize,
	kPhotoSizeMediumSize,
	kPhotoSizeXLargeSize,
	kPhotoSizeYLargeSize,
	kPhotoSizeASize,
	kPhotoSizeBSize,
	kPhotoSizeCSize,
}

func getSizeType(idx int) string {
	switch idx {
	case 0:
		return kPhotoSizeOriginalType
	case 1:
		return kPhotoSizeSmallType
	case 2:
		return kPhotoSizeMediumType
	case 3:
		return kPhotoSizeXLargeType
	case 4:
		return kPhotoSizeYLargeType
	case 5:
		return kPhotoSizeAType
	case 6:
		return kPhotoSizeBType
	case 7:
		return kPhotoSizeCType
	}

	return ""
}

//storage.fileUnknown#aa963b05 = storage.FileType;
//storage.filePartial#40bc6f52 = storage.FileType;
//storage.fileJpeg#7efe0e = storage.FileType;
//storage.fileGif#cae1aadf = storage.FileType;
//storage.filePng#a4f63c0 = storage.FileType;
//storage.filePdf#ae1e508d = storage.FileType;
//storage.fileMp3#528a0677 = storage.FileType;
//storage.fileMov#4b09ebbc = storage.FileType;
//storage.fileMp4#b3cea0e4 = storage.FileType;
//storage.fileWebp#1081464c = storage.FileType;

//func checkIsABC(idx int) bool {
//	if idx ==
//}

//import "github.com/golang/glog"
//

type resizeInfo struct {
	isWidth bool
	size    int
}

////////////////////////////////////////////////////////////////////////////
func (m *PhotoModel) GetPhotoSizeList(photoId int64) (sizes []*mtproto.PhotoSize) {
	doList := m.dao.PhotoDatasDAO.SelectListByPhotoId(photoId)
	sizes = make([]*mtproto.PhotoSize, 0, len(doList))
	for i := 1; i < len(doList); i++ {
		sizeData := &mtproto.PhotoSize_Data{
			Type: getSizeType(int(doList[i].LocalId)),
			W:    doList[i].Width,
			H:    doList[i].Height,
			Size: doList[i].FileSize,
			Location: &mtproto.FileLocation{
				// Constructor: mtproto.TLConstructor_CRC32_fileLocationLayer86,
				Constructor: mtproto.TLConstructor_CRC32_fileLocation,
				Data2: &mtproto.FileLocation_Data{
					VolumeId:      doList[i].VolumeId,
					LocalId:       int32(doList[i].LocalId),
					Secret:        doList[i].AccessHash,
					DcId:          doList[i].DcId,
					FileReference: []byte("@benqi-not-impl-file-reference"),
				},
			},
		}

		sizes = append(sizes, &mtproto.PhotoSize{
			Constructor: mtproto.TLConstructor_CRC32_photoSize,
			Data2:       sizeData,
		})

		// TODO(@benqi): add photoCachedSize
		//if i == 1 {
		//	var filename = core.NBFS_DATA_PATH + doList[i].FilePath
		//	cacheData, err := ioutil.ReadFile(filename)
		//	if err != nil {
		//		glog.Errorf("read file %s error: %v", filename, err)
		//		sizeData.Bytes = []byte{}
		//	} else {
		//		sizeData.Bytes = cacheData
		//	}
		//	sizes = append(sizes, &mtproto.PhotoSize{
		//		Constructor: mtproto.TLConstructor_CRC32_photoCachedSize,
		//		Data2:       sizeData,
		//	})
		//} else {
		//	sizes = append(sizes, &mtproto.PhotoSize{
		//		Constructor: mtproto.TLConstructor_CRC32_photoSize,
		//		Data2:       sizeData,
		//	})
		//}
	}
	return
}

func (m *PhotoModel) UploadPhotoFile2(fileMDList []*nbfs.PhotoFileMetadata) (photoId, accessHsh int64, sizeList []*mtproto.PhotoSize, err error) {
	sizeList = make([]*mtproto.PhotoSize, 0, 4)

	for i, fileMD := range fileMDList {
		photoDatasDO := &dataobject.PhotoDatasDO{
			PhotoId:    fileMD.PhotoId,
			PhotoType:  fileMD.PhotoType,
			DcId:       fileMD.DcId,
			VolumeId:   fileMD.VolumeId,
			LocalId:    fileMD.LocalId,
			AccessHash: fileMD.SecretId,
			Width:      fileMD.Width,
			Height:     fileMD.Height,
			FileSize:   fileMD.FileSize,
			FilePath:   fileMD.FilePath,
			Ext:        fileMD.Ext,
		}

		// photoDatasDO.Bytes = imgBuf.Bytes()
		m.dao.PhotoDatasDAO.Insert(photoDatasDO)

		photoSizeData := &mtproto.PhotoSize_Data{
			Type: getSizeType(int(fileMD.LocalId)),
			W:    photoDatasDO.Width,
			H:    photoDatasDO.Height,
			// Size: int32(len(photoDatasDO.Bytes)),
			Size: photoDatasDO.FileSize,
			Location: &mtproto.FileLocation{
				// Constructor: mtproto.TLConstructor_CRC32_fileLocationLayer86,
				Constructor: mtproto.TLConstructor_CRC32_fileLocation,
				Data2: &mtproto.FileLocation_Data{
					VolumeId:      photoDatasDO.VolumeId,
					LocalId:       fileMD.LocalId,
					Secret:        photoDatasDO.AccessHash,
					DcId:          photoDatasDO.DcId,
					FileReference: []byte("@benqi-not-impl-file-reference"),
				},
			},
		}

		if i == 0 {
			photoId = fileMD.PhotoId
			accessHsh = fileMD.SecretId
			continue
		} else {
			size := &mtproto.PhotoSize{
				Constructor: mtproto.TLConstructor_CRC32_photoSize,
				Data2:       photoSizeData,
			}
			sizeList = append(sizeList, size)
		}
	}
	return
}
