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

package crypto

import (
	"crypto/md5"
	"fmt"
	"github.com/golang/glog"
	"io"
	"os"
)

// TODO(@benqi): remove to baselib
func CalcMd5File(filename string) (string, error) {
	// fileName := core.NBFS_DATA_PATH + m.data.FilePath
	f, err := os.Open(filename)
	if err != nil {
		glog.Error(err)
		return "", err
	}

	defer f.Close()

	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, f); err != nil {
		// fmt.Println("Copy", err)
		glog.Error("Copy - ", err)
		return "", err
	}

	return fmt.Sprintf("%x", md5Hash.Sum(nil)), nil
}
