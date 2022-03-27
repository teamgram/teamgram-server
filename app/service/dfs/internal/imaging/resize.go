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

package imaging

import (
	"bytes"
	"image"
	"strings"

	"github.com/teamgram/marmota/pkg/bytes2"
	"github.com/teamgram/proto/mtproto"

	"github.com/disintegration/imaging"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	JPEG imaging.Format = iota
	PNG
	GIF
	TIFF
	BMP
	WEBP
)

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

func getImageFormat(extName string) (int, error) {
	formats := map[string]imaging.Format{
		".jpg":  JPEG,
		".jpeg": JPEG,
		".png":  PNG,
		".tif":  TIFF,
		".tiff": TIFF,
		".bmp":  BMP,
		".gif":  GIF,
		// ".webp": WEBP,
	}

	ext := strings.ToLower(extName)
	f, ok := formats[ext]
	if !ok {
		return -1, imaging.ErrUnsupportedFormat
	}

	return int(f), nil
}

func ReSizeImage(rb []byte, extName string, isABC bool, cb func(szType string, localId int, w, h int32, b []byte) error) (err error) {
	var (
		img image.Image
		f   int
	)

	img, err = imaging.Decode(bytes.NewReader(rb))
	if err != nil {
		logx.Errorf("decode r(%d) error: %v", len(rb), err)
		return
	}
	imgSz := makeResizeInfo(img)

	var (
		szList    []mtproto.ReSizeInfo
		willBreak = false
		rsz       int
	)

	if isABC {
		szList = mtproto.ReSizeInfoABCList
	} else {
		szList = mtproto.ReSizeInfoPhotoList
	}

	for _, sz := range szList {
		rsz = sz.Size
		if !isABC {
			if rsz >= imgSz.size {
				rsz = imgSz.size
				willBreak = true
			}
		}

		// TODO(@benqi): FIXME
		var dst *image.NRGBA
		if imgSz.isWidth {
			dst = imaging.Resize(img, rsz, 0, imaging.Lanczos)
		} else {
			dst = imaging.Resize(img, 0, rsz, imaging.Lanczos)
		}

		f, err = getImageFormat(extName)
		if err != nil {
			logx.Error(err.Error())
			return
		}

		o := bytes2.NewBuffer(make([]byte, 0, len(rb)))
		if f == int(imaging.JPEG) {
			// err = imaging.Encode(o, dst, imaging.JPEG, imaging.JPEGQuality(95))
			err = imaging.Encode(o, dst, imaging.JPEG)
		} else {
			err = imaging.Encode(o, dst, imaging.Format(f))
		}

		if err != nil {
			logx.Error(err.Error())
			return
		}
		err = cb(sz.Type, sz.LocalId, int32(dst.Bounds().Dx()), int32(dst.Bounds().Dy()), o.Bytes())
		if err != nil {
			return
		}

		if willBreak {
			break
		}
	}

	return
}
