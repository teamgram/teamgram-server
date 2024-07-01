// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package gnet

import (
	"errors"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/gnet/codec"

	"github.com/panjf2000/gnet/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

func (s *Server) onWebsocketData(ctx *connContext, c gnet.Conn) (action gnet.Action) {
	ws := ctx.wsCodec
	if ws.ReadBufferBytes(c) == gnet.Close {
		return gnet.Close
	}
	ok, action := ws.Upgrade(c)
	if !ok {
		return
	}

	if ws.Buf.Len() <= 0 {
		return gnet.None
	}
	messages, err := ws.Decode(c)
	if err != nil {
		return gnet.Close
	}
	if messages == nil {
		return
	}
	for _, message := range messages {
		ws.Conn.Buffer = message.Payload

		if ctx.codec == nil {
			ctx.codec, err = codec.CreateMTProtoCodec(&ctx.wsCodec.Conn)
			if err != nil {
				if errors.Is(err, codec.ErrUnexpectedEOF) {
					return gnet.None
				}
				logx.Errorf("conn(%s) create codec error: %v", c, err)
				return gnet.Close
			}
		}

		frame, err := ctx.codec.Decode(&ws.Conn)
		if err != nil {
			logx.Errorf("conn(%s) frame is error: %v", c, err)
			action = gnet.Close
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

		logx.Infof("conn(%s) recv frame: %s", c, msg2)

		action = s.onMTPRawMessage(ctx, c, msg2)
		if action == gnet.Close {
			return
		}

		_, _ = ws.Conn.InboundBuffer.Write(ws.Conn.Buffer)
		ws.Conn.Buffer = ws.Conn.Buffer[:0]

	}
	return gnet.None
}
