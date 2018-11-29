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

package payments

import (
    "fmt"
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto"
)

// payments.validateRequestedInfo#770a8e74 flags:# save:flags.0?true msg_id:int info:PaymentRequestedInfo = payments.ValidatedRequestedInfo;
func (s *PaymentsServiceImpl) PaymentsValidateRequestedInfo(ctx context.Context, request *mtproto.TLPaymentsValidateRequestedInfo) (*mtproto.Payments_ValidatedRequestedInfo, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("payments.validateRequestedInfo - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // Sorry: not impl PaymentsValidateRequestedInfo logic
    glog.Warning("payments.validateRequestedInfo blocked, License key from https://nebula.chat required to unlock enterprise features.")

    return nil, fmt.Errorf("not imp PaymentsValidateRequestedInfo")
}
