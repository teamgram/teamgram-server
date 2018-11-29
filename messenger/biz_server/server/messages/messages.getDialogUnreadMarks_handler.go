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

package messages

import (
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "fmt"
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto"
)


// messages.getDialogUnreadMarks#22e24e22 = Vector<DialogPeer>;
func (s *MessagesServiceImpl) MessagesGetDialogUnreadMarks(ctx context.Context, request *mtproto.TLMessagesGetDialogUnreadMarks) (*mtproto.Vector_DialogPeer, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("MessagesGetDialogUnreadMarks - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // TODO(@benqi): Impl MessagesGetDialogUnreadMarks logic

    return nil, fmt.Errorf("Not impl MessagesGetDialogUnreadMarks")
}
