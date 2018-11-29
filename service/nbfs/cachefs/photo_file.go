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
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/golang/glog"
	"image"
	"io/ioutil"
	"os"
	"strings"
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

type resizeInfo struct {
	isWidth bool
	size    int
}

func makeResizeInfo(img image.Image) resizeInfo {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	if w >= h {
		return resizeInfo{
			isWidth: true,
			size:    w,
		}
	} else {
		return resizeInfo{
			isWidth: false,
			size:    h,
		}
	}
}

type PhotoFile struct {
	volumeId int64
	localId  int32
	secretId int64
}

func NewPhotoFile(volumeId int64, localId int32, secretId int64) *PhotoFile {
	return &PhotoFile{volumeId, localId, secretId}
}

func (f *PhotoFile) ToFilePath() string {
	return fmt.Sprintf("%s/%s/%d.%d.dat", rootDataPath, getSizeType(int(f.localId)), f.volumeId, f.secretId)
}

func (f *PhotoFile) ToFilePath2() string {
	return fmt.Sprintf("/%s/%d.%d.dat", getSizeType(int(f.localId)), f.volumeId, f.secretId)
}

func (f *PhotoFile) WritePhotoFile(b []byte) error {
	return ioutil.WriteFile(f.ToFilePath(), b, 0644)
}

func (f *PhotoFile) ReadData(offset int32, limit int32) ([]byte, error) {
	return ReadFileOffsetData(f.ToFilePath(), offset, limit)
}

type PhotoInfo struct {
	LocalId  int32
	Width    int32
	Height   int32
	FileSize int64
}

// Save saves the image to file with the specified filename.
// The format is determined from the filename extension: "jpg" (or "jpeg"), "png", "gif", "tif" (or "tiff") and "bmp" are supported.
//
// Examples:
//
//	// Save the image as PNG.
//	err := imaging.Save(img, "out.png")
//
//	// Save the image as JPEG with optional quality parameter set to 80.
//	err := imaging.Save(img, "out.jpg", imaging.JPEGQuality(80))
//
func saveImage(img image.Image, filename string, f imaging.Format, opts ...imaging.EncodeOption) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return imaging.Encode(file, img, f, opts...)
}

func getImageFormat(extName string) (int, error) {
	formats := map[string]imaging.Format{
		".jpg":  imaging.JPEG,
		".jpeg": imaging.JPEG,
		".png":  imaging.PNG,
		".tif":  imaging.TIFF,
		".tiff": imaging.TIFF,
		".bmp":  imaging.BMP,
		".gif":  imaging.GIF,
	}

	ext := strings.ToLower(extName)
	f, ok := formats[ext]
	if !ok {
		return -1, imaging.ErrUnsupportedFormat
	}

	return int(f), nil
}

// TODO(@benqi):
// 	我们未来的图片存储系统可能会按facebook的Haystack论文来实现
// 	mtproto协议也定义了一套自己的文件存储方案，fileLocation#53d69076 dc_id:int volume_id:long local_id:int secret:long = FileLocation;
// 	在这里，我们重新定义mtproto的volume_id和local_id，对应Haystack的key和alternate_key，secret对应cookie
//  在当前简单实现里，volume_id由sonwflake生成，local_id对应于图片类型，secret为access_hash
// TODO(@benqi):
//  参数使用mtproto.File
func DoUploadedPhotoFile(src *PhotoFile, extName string, srcData []byte, isABC bool, cb func(pi *PhotoInfo)) error {
	img, err := imaging.Decode(bytes.NewReader(srcData))
	if err != nil {
		glog.Errorf("Decode %s error: {%v}", src, err)
		return err
	}

	pf := &PhotoFile{
		volumeId: src.volumeId,
		localId:  src.localId,
		// secretId: src.secretId,
	}

	imgSz := makeResizeInfo(img)
	for i, sz := range sizeList {
		pf.localId = int32(i)
		if i != 0 {
			if isABC {
				if i <= kPhotoSizeAIndex {
					continue
				}
			} else {
				if i > kPhotoSizeAIndex {
					continue
				}
			}
		}

		pi := &PhotoInfo{
			LocalId: int32(pf.localId),
		}

		if i == 0 {
			err = pf.WritePhotoFile(srcData)
			if err != nil {
				glog.Errorf("encode error: {%v}", err)
				return err
			}
			// pi.LocalId = int32(0)
			pi.Width = int32(img.Bounds().Dx())
			pi.Height = int32(img.Bounds().Dy())
			pi.FileSize = int64(len(srcData))
		} else {
			var dst *image.NRGBA
			if imgSz.isWidth {
				dst = imaging.Resize(img, sz, 0, imaging.Lanczos)
			} else {
				dst = imaging.Resize(img, 0, sz, imaging.Lanczos)
			}

			f, err := getImageFormat(extName)
			if err != nil {
				glog.Error(err)
				return err
			}

			dstFileName := pf.ToFilePath()
			if f == int(imaging.JPEG) {
				err = saveImage(dst, dstFileName, imaging.JPEG, imaging.JPEGQuality(80))
			} else {
				err = saveImage(dst, dstFileName, imaging.Format(f))
			}
			// err = imaging.Save(dst, pf.ToFilePath())
			if err != nil {
				glog.Errorf("encode error: {%v}", err)
				return err
			}

			pi.Width = int32(dst.Bounds().Dx())
			pi.Height = int32(dst.Bounds().Dy())
			pi.FileSize = getFileSize(dstFileName)
		}

		if cb != nil {
			cb(pi)
		}
	}

	return nil
}
