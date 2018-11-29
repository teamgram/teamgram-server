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
    "golang.org/x/net/context"
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto"
    "github.com/BurntSushi/toml"
)

// langpack.getLangPack#f2f2330a lang_pack:string lang_code:string = LangPackDifference;
func (s *LangpackServiceImpl) LangpackGetLangPack(ctx context.Context, request *mtproto.TLLangpackGetLangPack) (*mtproto.LangPackDifference, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("langpack.getLangPack#f2f2330a - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    if _, err := toml.DecodeFile(LANG_PACK_EN_FILE, &langs); err != nil {
        glog.Errorf("langpack.getLangPack#f2f2330a - decode file %s error: %v", LANG_PACK_EN_FILE, err)
        return nil, err
    }

    diff := mtproto.NewTLLangPackDifference()
    diff.SetLangCode(request.LangCode)
    diff.SetVersion(langs.Version)
    diff.SetFromVersion(0)

    diffStrings := make([]*mtproto.LangPackString, 0)
    for _, strings := range langs.Strings {
        diffStrings = append(diffStrings, &mtproto.LangPackString{
            Constructor: mtproto.TLConstructor_CRC32_langPackString,
            Data2:       strings,
        })
    }

    for _, stringPluralizeds := range langs.StringPluralizeds {
        diffStrings = append(diffStrings, &mtproto.LangPackString{
            Constructor: mtproto.TLConstructor_CRC32_langPackStringPluralized,
            Data2:       stringPluralizeds,
        })
    }

    diff.SetStrings(diffStrings)
    // reply := mtproto.MakeLangPackDifference(diff)
    glog.Infof("langpack.getLangPack#f2f2330a - reply: %s", logger.JsonDebugData(diff))
    return diff.To_LangPackDifference(), nil
}
