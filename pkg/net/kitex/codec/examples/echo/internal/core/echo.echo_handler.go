// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/echo"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
)

var _ *tg.Bool

// EchoEcho
// echo.echo message:string = Echo;
func (c *EchoCore) EchoEcho(in *echo.TLEchoEcho) (resp *echo.Echo, err error) {
	// do something here
	/*
		temp, ok1 := metainfo.GetValue(c.ctx, "temp")
		logid, ok2 := metainfo.GetPersistentValue(c.ctx, "logid")

		if !(ok1 && ok2) {
			// panic("It looks like the protocol does not support transmitting meta information")
		}
		klog.Debug(temp)  // "temp-value"
		klog.Debug(logid) // "12345"
		klog.Debug("echo called")
	*/

	md := metadata.RpcMetadataFromIncoming(c.ctx)
	c.Logger.Debugf("md message: %s", md)
	return echo.MakeEcho(&echo.TLEcho{
		ClazzID: echo.ClazzID_echo,
		Message: in.Message,
	}), nil
}
