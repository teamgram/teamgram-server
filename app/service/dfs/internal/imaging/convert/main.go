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
	"flag"
	"fmt"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io/ioutil"
	"path/filepath"

	"github.com/minio/minio-go/v7"

	"github.com/teamgram/teamgram-server/app/service/dfs/internal/imaging"
)

var imageFile = flag.String("image", "", "convert image file")
var isABC = flag.Bool("abc", true, "output isABC")

var (
	minioCore *minio.Core
)

func main() {
	flag.Parse()

	if *imageFile == "" {
		flag.Usage()
		return
	}

	ext := filepath.Ext(*imageFile)
	rb, err := ioutil.ReadFile(*imageFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	imaging.ReSizeImage(rb, ext, *isABC, func(szType string, localId int, w, h int32, b []byte) error {
		return ioutil.WriteFile(fmt.Sprintf("%s.%s%s", *imageFile, szType, ext), b, 0644)
	})
}

func init() {
	var err error
	minioCore, err = minio.NewCore(
		"127.0.0.1:9000",
		&minio.Options{
			Creds:  credentials.NewStaticV4("TLXH0OZVP0AKOJAZ8DIT", "9Sw+Xbhc3aWvxQ78rRgUkTQQLLZ24SyelA+B6Rwe", ""),
			Secure: false,
		})
	if err != nil {
		panic("new minio error")
	}
}
