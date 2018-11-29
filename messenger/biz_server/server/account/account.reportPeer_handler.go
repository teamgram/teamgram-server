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

/*
 Android client source code:
	if (ChatObject.isChannel(currentChat) && !currentChat.creator && (!currentChat.megagroup || currentChat.username != null && currentChat.username.length() > 0)) {
		headerItem.addSubItem(report, LocaleController.getString("ReportChat", R.string.ReportChat));
	}
*/
// account.reportPeer#ae189d5f peer:InputPeer reason:ReportReason = Bool;
func (s *AccountServiceImpl) AccountReportPeer(ctx context.Context, request *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.reportPeer#ae189d5f - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Check peer invalid
	peer := request.Peer
	// TODO(@benqi): Check peer access_hash
	if peer.GetConstructor() != mtproto.TLConstructor_CRC32_inputPeerChannel {
		// TODO(@benqi): Add INPUT_PEER_INVALID code
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err)
		return nil, err
	} else {
		// TODO(@benqi): !currentChat.creator && (!currentChat.megagroup || currentChat.username != null && currentChat.username.length() > 0)
	}

	// peer := helper.FromInputPeer(request.GetPeer())
	reason := base.FromReportReason(request.GetReason())

	text := ""
	if reason == base.REASON_OTHER {
		text = request.GetReason().GetData2().GetText()
	}

	s.AccountModel.InsertReportData(md.UserId, base.PEER_CHANNEL, peer.GetData2().GetChannelId(), int32(reason), text)

	glog.Infof("account.reportPeer#ae189d5f - reply: {true}")
	return mtproto.ToBool(true), nil
}
