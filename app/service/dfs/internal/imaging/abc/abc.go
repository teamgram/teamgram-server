// Copyright © 2024 Teamgram Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	
	"github.com/teamgram/marmota/pkg/bytes2"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/imaging"
)

var imageFile = flag.String("image", "", "convert image file")

func main() {
	flag.Parse()

	if *imageFile == "" {
		flag.Usage()
		return
	}

	// fill
	{
		img, _ := imaging.Open(*imageFile)

		dstImg := imaging.Fill(img, 800, 800)
		b := bytes2.NewBuffer(make([]byte, 0, 1024*1024))
		_ = imaging.EncodeJpeg(b, dstImg)
		_ = os.WriteFile(*imageFile+".800x800.jpg", b.Bytes(), 0644)
	}

	// resize
	{
		ext := filepath.Ext(*imageFile)
		rb, err := os.ReadFile(*imageFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		_ = imaging.ReSizeImage(rb, ext, true, func(szType string, localId int, w, h int32, b []byte) error {
			return os.WriteFile(fmt.Sprintf("%s.%s%s", *imageFile, szType, ext), b, 0644)
		})
	}
}
