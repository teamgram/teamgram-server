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

package jpeg

import (
	"bufio"
	"errors"
	"image"
	"io"
)

func EncodeStripped(w io.Writer, m image.Image, o *Options) error {
	b := m.Bounds()
	if b.Dx() >= 1<<16 || b.Dy() >= 1<<16 {
		return errors.New("jpeg: image is too large to encode")
	}
	var e encoder
	e.w = bufio.NewWriter(w)
	//if ww, ok := w.(writer); ok {
	//	e.w = ww
	//} else {
	//	e.w = bufio.NewWriter(w)
	//}
	// Clip quality to [1, 100].
	quality := DefaultQuality
	if o != nil {
		quality = o.Quality
		if quality < 1 {
			quality = 1
		} else if quality > 100 {
			quality = 100
		}
	}
	// Convert from a quality rating to a scaling factor.
	var scale int
	if quality < 50 {
		scale = 5000 / quality
	} else {
		scale = 200 - quality*2
	}
	// Initialize the quantization tables.
	for i := range e.quant {
		for j := range e.quant[i] {
			x := int(unscaledQuant[i][j])
			x = (x*scale + 50) / 100
			if x < 1 {
				x = 1
			} else if x > 255 {
				x = 255
			}
			e.quant[i][j] = uint8(x)
		}
	}
	//// Compute number of components based on input image type.
	//nComponent := 3
	//switch m.(type) {
	//// TODO(wathiede): switch on m.ColorModel() instead of type.
	//case *image.Gray:
	//	nComponent = 1
	//}
	//// Write the Start Of Image marker.
	//e.buf[0] = 0xff
	//e.buf[1] = 0xd8
	//e.write(e.buf[:2])
	//// Write the quantization tables.
	//e.writeDQT()
	//// Write the image dimensions.
	//e.writeSOF0(b.Size(), nComponent)
	//// Write the Huffman tables.
	//e.writeDHT(nComponent)
	// Write the image data.

	e.buf[0] = 0x01
	e.buf[1] = byte(b.Dy())
	e.buf[2] = byte(b.Dx())
	e.write(e.buf[:3])
	e.writeSOSIgnoreHeader(m)
	// Write the End Of Image marker.
	//e.buf[0] = 0xff
	//e.buf[1] = 0xd9
	//e.write(e.buf[:2])
	e.flush()
	return e.err
}

// writeSOS writes the StartOfScan marker.
func (e *encoder) writeSOSIgnoreHeader(m image.Image) {
	//switch m.(type) {
	//case *image.Gray:
	//	e.write(sosHeaderY)
	//default:
	//	e.write(sosHeaderYCbCr)
	//}
	var (
		// Scratch buffers to hold the YCbCr values.
		// The blocks are in natural (not zig-zag) order.
		b      block
		cb, cr [4]block
		// DC components are delta-encoded.
		prevDCY, prevDCCb, prevDCCr int32
	)
	bounds := m.Bounds()
	switch m := m.(type) {
	// TODO(wathiede): switch on m.ColorModel() instead of type.
	case *image.Gray:
		for y := bounds.Min.Y; y < bounds.Max.Y; y += 8 {
			for x := bounds.Min.X; x < bounds.Max.X; x += 8 {
				p := image.Pt(x, y)
				grayToY(m, p, &b)
				prevDCY = e.writeBlock(&b, 0, prevDCY)
			}
		}
	default:
		rgba, _ := m.(*image.RGBA)
		ycbcr, _ := m.(*image.YCbCr)
		for y := bounds.Min.Y; y < bounds.Max.Y; y += 16 {
			for x := bounds.Min.X; x < bounds.Max.X; x += 16 {
				for i := 0; i < 4; i++ {
					xOff := (i & 1) * 8
					yOff := (i & 2) * 4
					p := image.Pt(x+xOff, y+yOff)
					if rgba != nil {
						rgbaToYCbCr(rgba, p, &b, &cb[i], &cr[i])
					} else if ycbcr != nil {
						yCbCrToYCbCr(ycbcr, p, &b, &cb[i], &cr[i])
					} else {
						toYCbCr(m, p, &b, &cb[i], &cr[i])
					}
					prevDCY = e.writeBlock(&b, 0, prevDCY)
				}
				scale(&b, &cb)
				prevDCCb = e.writeBlock(&b, 1, prevDCCb)
				scale(&b, &cr)
				prevDCCr = e.writeBlock(&b, 1, prevDCCr)
			}
		}
	}
	// Pad the last byte with 1's.
	e.emit(0x7f, 7)
}
