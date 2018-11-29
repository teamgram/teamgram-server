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

package model

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"testing"
)

func TestGetLangPacks(t *testing.T) {
	var langPacks LangPacks

	if _, err := toml.DecodeFile("../../lang_pack_en.toml", &langPacks); err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("%v\n", langPacks)
}
