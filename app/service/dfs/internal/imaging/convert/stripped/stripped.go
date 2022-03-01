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

package main

import (
	"fmt"
	"os"

	"github.com/teamgram/marmota/pkg/bytes2"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/imaging"
)

func main() {
	img, err := imaging.Open("../test002.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(img.Bounds())
	w := 40
	h := 40
	if img.Bounds().Dx() > img.Bounds().Dy() {
		h = 0
	} else {
		w = 0
	}
	dst := imaging.Resize(img, w, h)
	fmt.Println(dst.Bounds())

	sw := bytes2.NewBuffer(make([]byte, 0, 4096))
	err = imaging.EncodeStripped(sw, dst, 30)
	if err != nil {
		fmt.Println(err)
		return
	}
	bb := sw.Bytes()
	fmt.Println(len(bb))

	file, err := os.Create("./test002.jpeg")
	if err != nil {
		// return err
	}
	defer file.Close()

	head := imaging.JpegHeader
	head[164] = bb[1]
	head[166] = bb[2]
	file.Write(head)
	file.Write(bb[3:])
	file.Write(imaging.JpegFooter)
}
