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
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// langpack.getStrings#efea3803 lang_pack:string lang_code:string keys:Vector<string> = Vector<LangPackString>;
func (s *LangpackServiceImpl) LangpackGetStrings(ctx context.Context, request *mtproto.TLLangpackGetStrings) (*mtproto.Vector_LangPackString, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("langpack.getStrings#efea3803 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	langPacks := queryLangPacks(request.GetLangCode())
	langpackStrings := &mtproto.Vector_LangPackString{}
	for _, s := range request.Keys {
		s2 := &mtproto.TLLangPackString{Data2: &mtproto.LangPackString_Data{
			Key:   s,
			Value: langPacks.Query(s),
		}}
		langpackStrings.Datas = append(langpackStrings.Datas, s2.To_LangPackString())
	}

	glog.Infof("langpack.getStrings#efea3803 - reply: %s", logger.JsonDebugData(langpackStrings))
	return langpackStrings, nil
}
