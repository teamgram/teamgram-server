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
	"bytes"
	"encoding/binary"
	"errors"
	"io"

	"github.com/lxzan/gws"
	"github.com/teamgooo/teamgooo-server/app/interface/gnetway/internal/server/gnet/codec"
	"github.com/zeromicro/go-zero/core/logx"
)

// WebSocket event handler for gws
type websocketHandler struct {
	server *Server
	conn   *connection
}

func (h *websocketHandler) OnOpen(socket *gws.Conn) {
	logx.Debugf("websocket connection opened from %s (id=%d)", h.conn.RemoteAddr(), h.conn.id)
}

func (h *websocketHandler) OnClose(socket *gws.Conn, err error) {
	logx.Debugf("websocket connection closed from %s (id=%d), err: %v", h.conn.RemoteAddr(), h.conn.id, err)
}

func (h *websocketHandler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.WritePong(payload)
}

func (h *websocketHandler) OnPong(socket *gws.Conn, payload []byte) {
	// Do nothing
}

func (h *websocketHandler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	// Reset timeout
	h.conn.ResetTimeout()

	// Only process binary messages
	if message.Opcode != gws.OpcodeBinary {
		return
	}

	payload := message.Bytes()

	// Create codec if not exists
	if h.conn.codec == nil {
		adapter := &wsConnAdapter{
			conn:   h.conn,
			reader: bytes.NewReader(payload),
		}

		var err error
		h.conn.codec, err = codec.CreateCodec(adapter)
		if err != nil {
			if !errors.Is(err, codec.ErrUnexpectedEOF) {
				logx.Errorf("conn(%d) create codec error: %v", h.conn.id, err)
				_ = socket.WriteClose(1002, []byte("codec error"))
			}
			return
		}
	}

	adapter := &wsConnAdapter{
		conn:   h.conn,
		reader: bytes.NewReader(payload),
	}

	needAck, frame, err := h.conn.codec.Decode(adapter)
	if err != nil {
		if !errors.Is(err, codec.ErrUnexpectedEOF) {
			logx.Errorf("conn(%d) frame decode error: %v", h.conn.id, err)
			_ = socket.WriteClose(1002, []byte("decode error"))
		}
		return
	}

	if frame == nil {
		return
	}

	authKeyId := int64(binary.LittleEndian.Uint64(frame))
	shouldClose := h.server.onMTPRawMessage(h.conn, authKeyId, needAck, frame)
	if shouldClose {
		_ = socket.WriteClose(1000, []byte("session closed"))
	}
}

func (s *Server) handleWebSocketConnection(c *connection) {
	defer func() {
		s.wg.Done()
		s.onClose(c, nil)
		s.connMgr.remove(c.id)
	}()

	logx.Debugf("new WebSocket connection from %s (id=%d)", c.RemoteAddr(), c.id)

	upgrader := gws.NewUpgrader(&websocketHandler{
		server: s,
		conn:   c,
	}, &gws.ServerOption{
		ReadBufferSize:   65536,
		WriteBufferSize:  65536,
		CompressEnabled:  false,
		CheckUtf8Enabled: false,
	})

	conn, err := upgrader.Upgrade(c.conn)
	if err != nil {
		logx.Errorf("websocket upgrade error: %v", err)
		return
	}

	c.gwsConn = conn

	// Run read loop - this blocks until connection closes
	conn.ReadLoop()
}

// wsConnAdapter adapts connection to codec.CodecReader/CodecWriter interface
type wsConnAdapter struct {
	conn   *connection
	reader *bytes.Reader
}

func (w *wsConnAdapter) Peek(n int) ([]byte, error) {
	if w.reader.Len() < n {
		return nil, io.EOF
	}

	buf := make([]byte, n)
	_, err := w.reader.Read(buf)
	if err != nil {
		return nil, err
	}

	// Reset reader position
	_, _ = w.reader.Seek(-int64(n), io.SeekCurrent)

	return buf, nil
}

func (w *wsConnAdapter) Discard(n int) (int, error) {
	if w.reader.Len() < n {
		return 0, io.EOF
	}

	_, err := w.reader.Seek(int64(n), io.SeekCurrent)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (w *wsConnAdapter) Next(n int) ([]byte, error) {
	buf := make([]byte, n)
	readN, err := io.ReadFull(w.reader, buf)
	return buf[:readN], err
}

func (w *wsConnAdapter) InboundBuffered() int {
	return w.reader.Len()
}

func (w *wsConnAdapter) Write(data []byte) (int, error) {
	return w.conn.Write(data)
}
