// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
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

package gnet

import (
	"errors"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/gnet/codec"

	"github.com/panjf2000/gnet/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

func (s *Server) onTcpData(ctx *connContext, c gnet.Conn) (action gnet.Action) {
	if ctx.codec == nil {
		var (
			err error
		)
		ctx.codec, err = codec.CreateMTProtoCodec(c)
		if err != nil {
			if errors.Is(err, codec.ErrUnexpectedEOF) {
				return gnet.None
			} else {
				logx.Errorf("conn(%s) create codec error: %v", c, err)
				return gnet.Close
			}
		}
	}

	for {
		frame, err := ctx.codec.Decode(c)
		if err != nil {
			if errors.Is(err, codec.ErrUnexpectedEOF) {
				return gnet.None
			} else {
				logx.Errorf("conn(%s) frame is error: %v", c, err)
				action = gnet.Close
			}
			return
		} else if frame == nil {
			logx.Debugf("conn(%s) frame is nil", c)
			return
		}

		msg2, ok := frame.(*mtproto.MTPRawMessage)
		if !ok {
			logx.Errorf("conn(%s) recv error: msg2 not codec.MTPRawMessage type", c)
			action = gnet.Close
			return
		}

		logx.Infof("conn(%s) recv frame: %s", c, msg2.String())

		action = s.onMTPRawMessage(ctx, c, msg2)
		if action == gnet.Close {
			return
		}
	}

	return gnet.None
}
