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
	"golang.org/x/net/context"
)

// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
func (s *AccountServiceImpl) AccountSetAccountTTL(ctx context.Context, request *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountSetAccountTTL - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Check ttl
	ttl := request.GetTtl().GetData2().GetDays()
	if ttl <= 0 || ttl > 365 {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("ttl_days error: ", err)
		return nil, err
	}
	s.AccountModel.SetAccountDaysTTL(md.UserId, ttl)

	glog.Infof("account.setAccountTTL#2442485e - reply: {true}")
	return mtproto.ToBool(true), nil
}
