/*
 *  Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpc

import (
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto"
)

// session.getLangCode auth_key_id:long = String;
func (s *SessionServiceImpl) SessionGetLangCode(ctx context.Context, request *mtproto.TLSessionGetLangCode) (*mtproto.String, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("session.getLangCode - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    langCode := s.AuthSessionModel.GetLangCode(request.GetAuthKeyId())
    reply := &mtproto.TLString{Data2: &mtproto.String_Data{
        V: langCode,
    }}

    glog.Infof("session.getLangCode - reply: {%s}", logger.JsonDebugData(reply))
    return reply.To_String(), nil
}
