// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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

package netserver

import (
	"encoding/binary"
	"errors"
	"io"

	"github.com/teamgooo/teamgooo-server/app/interface/gnetway/internal/server/netserver/codec"
	"github.com/zeromicro/go-zero/core/logx"
)

func (s *Server) handleTCPConnection(c *connection) {
	defer func() {
		s.wg.Done()
		s.onClose(c, nil)
		s.connMgr.remove(c.id)
	}()

	logx.Debugf("new TCP connection from %s (id=%d)", c.RemoteAddr(), c.id)

	// Connection read loop
	for {
		select {
		case <-s.shutdownCh:
			return
		default:
		}

		if c.IsClosed() {
			return
		}

		shouldClose := s.onTcpData(c)
		if shouldClose {
			return
		}
	}
}

func (s *Server) onTcpData(c *connection) (shouldClose bool) {
	// Reset timeout on each data receive
	c.ResetTimeout()

	if c.codec == nil {
		var err error
		c.codec, err = codec.CreateCodec(&tcpConnAdapter{c})
		if err != nil {
			if errors.Is(err, codec.ErrUnexpectedEOF) {
				return false
			}
			logx.Errorf("conn(%d) create codec error: %v", c.id, err)
			return true
		}
	}

	for {
		needAck, frame, err := c.codec.Decode(&tcpConnAdapter{c})
		if err != nil {
			if errors.Is(err, codec.ErrUnexpectedEOF) {
				return false
			}
			if errors.Is(err, io.EOF) {
				return true
			}
			logx.Errorf("conn(%d) frame decode error: %v", c.id, err)
			return true
		}

		if frame == nil {
			break
		}

		authKeyId := int64(binary.LittleEndian.Uint64(frame))
		shouldClose = s.onMTPRawMessage(c, authKeyId, needAck, frame)
		if shouldClose {
			return true
		}
	}

	return false
}

// tcpConnAdapter adapts connection to codec.CodecReader/CodecWriter interface
type tcpConnAdapter struct {
	*connection
}

func (t *tcpConnAdapter) Peek(n int) ([]byte, error) {
	return t.reader.Peek(n)
}

func (t *tcpConnAdapter) Discard(n int) (int, error) {
	return t.reader.Discard(n)
}

func (t *tcpConnAdapter) Next(n int) ([]byte, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(t.reader, buf)
	return buf, err
}

func (t *tcpConnAdapter) InboundBuffered() int {
	return t.reader.Buffered()
}
