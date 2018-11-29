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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/account"
	"golang.org/x/net/context"
)

// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (s *AccountServiceImpl) AccountGetPrivacy(ctx context.Context, request *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getPrivacy#dadbc950 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	privacyLogic := s.AccountModel.MakePrivacyLogic(md.UserId)
	rulesData := privacyLogic.GetPrivacy(account.FromInputPrivacyKey(request.Key))

	var rules *mtproto.TLAccountPrivacyRules
	if rulesData == nil {
		// TODO(@benqi): return nil or empty
		// rules = mtproto.NewTLAccountPrivacyRules()
		rules = &mtproto.TLAccountPrivacyRules{Data2: &mtproto.Account_PrivacyRules_Data{
			Rules: []*mtproto.PrivacyRule{mtproto.NewTLPrivacyValueAllowAll().To_PrivacyRule()},
		}}
	} else {
		idList := rulesData.PickAllUserIdList()
		if len(idList) == 0 {
			rules = &mtproto.TLAccountPrivacyRules{Data2: &mtproto.Account_PrivacyRules_Data{
				Rules: rulesData.ToPrivacyRuleList(),
			}}
		} else {
			rules = &mtproto.TLAccountPrivacyRules{Data2: &mtproto.Account_PrivacyRules_Data{
				Rules: rulesData.ToPrivacyRuleList(),
				Users: s.UserModel.GetUserListByIdList(md.UserId, idList),
			}}
		}
	}

	glog.Infof("account.getPrivacy#dadbc950 - reply: %s", logger.JsonDebugData(rules))
	return rules.To_Account_PrivacyRules(), nil
}
