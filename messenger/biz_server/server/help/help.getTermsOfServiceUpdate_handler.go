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

package help

import (
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto"
    "time"
)

//
// help.termsOfServiceUpdateEmpty#e3309f7f expires:int = help.TermsOfServiceUpdate;
// help.termsOfServiceUpdate#28ecf961 expires:int terms_of_service:help.TermsOfService = help.TermsOfServiceUpdate;
//
// help.termsOfService#780a0310 flags:# popup:flags.0?true id:DataJSON text:string entities:Vector<MessageEntity> min_age_confirm:flags.1?int = help.TermsOfService;
//
// help.getTermsOfServiceUpdate#2ca51fd1 = help.TermsOfServiceUpdate;
func (s *HelpServiceImpl) HelpGetTermsOfServiceUpdate(ctx context.Context, request *mtproto.TLHelpGetTermsOfServiceUpdate) (*mtproto.Help_TermsOfServiceUpdate, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("HelpGetTermsOfServiceUpdate - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))


    termsOfServiceUpdate := &mtproto.TLHelpTermsOfServiceUpdateEmpty{Data2: &mtproto.Help_TermsOfServiceUpdate_Data{
        Expires: int32(time.Now().Unix() + 3600),
    }}

    glog.Infof("help.getTermsOfServiceUpdate#2ca51fd1 - reply: %s", logger.JsonDebugData(termsOfServiceUpdate))
    return termsOfServiceUpdate.To_Help_TermsOfServiceUpdate(), nil
}
