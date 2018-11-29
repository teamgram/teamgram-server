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
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

/**

##contacts.exportCard

Returns the current user's card that can be later used to contact a Telegram user without knowing his phone number.

```
	---functions---
	contacts.exportCard#84e53737 = Vector<int>;
```

### Parameters
This method does not require any parameters.

### Result
The method returns Vector<int>.

We recommend showing this card as color-separated hex numbers,
e.g: 000623bf:2fe34c70:23f70153:a8a63dc2:62fc8e8f or QR-code representation.

*/

// 客户端未使用
// contacts.exportCard#84e53737 = Vector<int>;
func (s *ContactsServiceImpl) ContactsExportCard(ctx context.Context, request *mtproto.TLContactsExportCard) (*mtproto.VectorInt, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.exportCard#84e53737 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ContactsExportCard logic
	exports := &mtproto.VectorInt{
		Datas: []int32{},
	}

	glog.Info("contacts.exportCard#84e53737 - not impl ContactsExportCard, reply: {}")
	return exports, nil
}
