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

package langpack

import (
	"github.com/BurntSushi/toml"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	model2 "github.com/nebula-chat/chatengine/messenger/biz_server/server/langpack/model"
	"github.com/nebula-chat/chatengine/pkg/util"
)

const (
	LANG_PACK_EN_FILE = "/lang_pack_en.toml"
	LANG_PACK_RU_FILE = "/lang_pack_ru.toml"
)

var (
	langPacksEn model2.LangPacks
	langPacksRu model2.LangPacks
)

func init() {
	if _, err := toml.DecodeFile(util.GetWorkingDirectory() + "/" + LANG_PACK_EN_FILE, &langPacksEn); err != nil {
		panic(err)
	}
	if _, err := toml.DecodeFile(util.GetWorkingDirectory() + "/" + LANG_PACK_RU_FILE, &langPacksRu); err != nil {
		panic(err)
	}
}

type LangpackServiceImpl struct {
}

func NewLangpackServiceImpl(models []core.CoreModel) *LangpackServiceImpl {
	impl := &LangpackServiceImpl{}

	for _, m := range models {
		switch m.(type) {
		}
	}

	return impl
}

var langPackVersion = int32(77)

func queryLangPacks(langCode string) *model2.LangPacks {
	var langPacks *model2.LangPacks

	switch langCode {
	case "en":
		langPacks = &langPacksEn
	case "ru":
		langPacks = &langPacksRu
	default:
		langPacks = &langPacksEn
	}

	return langPacks
}