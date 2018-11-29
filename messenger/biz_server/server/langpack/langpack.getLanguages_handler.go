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
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

// languages config
/**************************

[[LangPackLanguages]]
Name = "English"
NativeName = "English"
LangCode = "en"

[[LangPackLanguages]]
Name = "German"
NativeName = "Deutsch"
LangCode = "de"

[[LangPackLanguages]]
Name = "Dutch"
NativeName = "Nederlands"
LangCode = "nl"

[[LangPackLanguages]]
Name = "Spanish"
NativeName = "Español"
LangCode = "es"

[[LangPackLanguages]]
Name = "Italian"
NativeName = "Italiano"
LangCode = "it"

[[LangPackLanguages]]
Name = "Portuguese (Brazil)"
NativeName = "Português (Brasil)"
LangCode = "pt-br"

[[LangPackLanguages]]
Name = "Korean"
NativeName = "한국어"
LangCode = "ko"

[[LangPackLanguages]]
Name = "Malay"
NativeName = "Bahasa Melayu"
LangCode = "ms"

[[LangPackLanguages]]
Name = "Russian"
NativeName = "Русский"
LangCode = "ru"

[[LangPackLanguages]]
Name = "French"
NativeName = "Français"
LangCode = "fr"

[[LangPackLanguages]]
Name = "Ukrainian"
NativeName = "Українська"
LangCode = "uk"

**************************************************/

// langpack.getLanguages#42c6978f lang_pack:string = Vector<LangPackLanguage>;
func (s *LangpackServiceImpl) LangpackGetLanguages(ctx context.Context, request *mtproto.TLLangpackGetLanguages) (*mtproto.Vector_LangPackLanguage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("langpack.getLanguages#42c6978f - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Add other language
	languages := &mtproto.Vector_LangPackLanguage{}
	language := &mtproto.TLLangPackLanguage{Data2: &mtproto.LangPackLanguage_Data{
		Name:       "English",
		NativeName: "English",
		LangCode:   "en",
	}}
	languages.Datas = append(languages.Datas, language.To_LangPackLanguage())

	languageRu := &mtproto.TLLangPackLanguage{Data2: &mtproto.LangPackLanguage_Data{
		Name:       "Russian",
		NativeName: "Русский",
		LangCode:   "ru",
	}}
	languages.Datas = append(languages.Datas, languageRu.To_LangPackLanguage())

	glog.Infof("langpack.getLanguages#42c6978f - reply: %s", logger.JsonDebugData(languages))
	return languages, nil
}
