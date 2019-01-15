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
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// messages.getDocumentByHash#338e2464 sha256:bytes size:int mime_type:string = Document;
func (s *MessagesServiceImpl) MessagesGetDocumentByHash(ctx context.Context, request *mtproto.TLMessagesGetDocumentByHash) (*mtproto.Document, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getDocumentByHash#338e2464 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetDocumentByHash logic
	document := mtproto.NewTLDocumentEmpty().To_Document()

	glog.Infof("messages.getDocumentByHash#338e2464 - reply: %s", logger.JsonDebugData(document))
	return document, nil
}
