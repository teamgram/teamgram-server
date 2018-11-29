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

// langpack.getDifference#b2e4d7d from_version:int = LangPackDifference;
func (s *LangpackServiceImpl) LangpackGetDifference(ctx context.Context, request *mtproto.TLLangpackGetDifference) (*mtproto.LangPackDifference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("langpack.getDifference#b2e4d7d - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl LangpackGetDifference logic
	diff := mtproto.NewTLLangPackDifference()
	diff.SetLangCode("en")
	diff.SetVersion(langs.Version)
	diff.SetFromVersion(request.FromVersion)

	if request.FromVersion < langs.Version {
		// TODO(@benqi): 找出不同版本的增量更新数据
	}

	glog.Infof("langpack.getDifference#b2e4d7d - reply: %s", logger.JsonDebugData(diff))
	return diff.To_LangPackDifference(), nil
}
