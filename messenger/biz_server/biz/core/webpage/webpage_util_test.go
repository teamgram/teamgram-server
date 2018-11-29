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

package webpage

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetWebpageOgList(t *testing.T) {
	ogContents := GetWebpageOgList("https://github.com/nebula-chat/chatengine", []string{"image", "site_name", "title", "description"})
	fmt.Println(ogContents)
}

func TestUrlParser(t *testing.T) {
	var (
		u   *url.URL
		err error
	)

	u, err = url.Parse("aaaa")
	fmt.Println(u, err)
	u, err = url.Parse("https://github.com/nebula-chat/chatengine")
	fmt.Println(u, err)
}
