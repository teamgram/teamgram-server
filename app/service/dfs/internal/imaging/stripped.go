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
	"image"
	"io"

	"github.com/teamgram/teamgram-server/app/service/dfs/internal/imaging/jpeg"
)

func EncodeStripped(w io.Writer, img image.Image, quality int) error {
	var (
		rgba *image.RGBA
		err  error
	)

	if nrgba, ok := img.(*image.NRGBA); ok {
		if nrgba.Opaque() {
			rgba = &image.RGBA{
				Pix:    nrgba.Pix,
				Stride: nrgba.Stride,
				Rect:   nrgba.Rect,
			}
		}
	}

	if rgba != nil {
		err = jpeg.EncodeStripped(w, rgba, &jpeg.Options{Quality: quality})
	} else {
		err = jpeg.EncodeStripped(w, img, &jpeg.Options{Quality: quality})
	}

	return err
}
