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

package contacts

import (
	// "fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

/**

## contacts.importCard
Returns general information on a user using his previously exported card as input.
The app may use it to open a conversation without knowing the user's phone number.

*/

// 客户端未使用
// contacts.importCard#4fe196fe export_card:Vector<int> = User;
func (s *ContactsServiceImpl) ContactsImportCard(ctx context.Context, request *mtproto.TLContactsImportCard) (*mtproto.User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.importCard#4fe196fe - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ContactsImportCard logic
	user := mtproto.NewTLUserEmpty()

	// userEmpty#200250ba id:int = User;
	glog.Infof("contacts.importCard#4fe196fe - reply: %s", logger.JsonDebugData(user))
	return user.To_User(), nil
}
