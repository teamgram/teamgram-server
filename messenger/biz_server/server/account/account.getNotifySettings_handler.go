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

package account

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"golang.org/x/net/context"
)

// account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;
func (s *AccountServiceImpl) AccountGetNotifySettings(ctx context.Context, request *mtproto.TLAccountGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getNotifySettings#12b3ad31 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		settings *mtproto.PeerNotifySettings
	)

	peer := base.FromInputNotifyPeer(md.UserId, request.GetPeer())
	if peer.PeerType == base.PEER_EMPTY {
		glog.Error("peer is empty")
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PEER_ID_INVALID)
		return nil, err
	}
	settings = s.AccountModel.GetNotifySettings(md.UserId, peer)
	glog.Infof("account.getNotifySettings#12b3ad31 - reply: %s", logger.JsonDebugData(settings))
	return settings, nil
}
