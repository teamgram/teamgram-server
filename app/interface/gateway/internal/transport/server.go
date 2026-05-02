package transport

import (
	"bufio"
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"sync"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/push"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/sessionstate"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
)

type Server struct {
	addr      string
	gatewayID string
	handshake *sessionstate.HandshakeManager
	processor *sessionstate.Processor
	push      *push.LocalWriter
	listener  net.Listener
	wg        sync.WaitGroup
	connsMu   sync.Mutex
	conns     map[net.Conn]struct{}
}

func NewServer(addr string, gatewayID string, handshake *sessionstate.HandshakeManager, processor *sessionstate.Processor, pushWriter *push.LocalWriter) *Server {
	return &Server{addr: addr, gatewayID: gatewayID, handshake: handshake, processor: processor, push: pushWriter, conns: make(map[net.Conn]struct{})}
}

func (s *Server) Start(ctx context.Context) error {
	if s == nil || s.addr == "" {
		return nil
	}
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	s.listener = ln
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			s.wg.Add(1)
			go func() {
				defer s.wg.Done()
				_ = s.ServeConn(ctx, conn)
			}()
		}
	}()
	return nil
}

func (s *Server) Stop() error {
	if s == nil {
		return nil
	}
	var err error
	if s.listener != nil {
		err = s.listener.Close()
	}
	s.closeActiveConns()
	s.wg.Wait()
	return err
}

func (s *Server) ServeConn(ctx context.Context, conn net.Conn) error {
	untrack := s.trackConn(conn)
	defer untrack()
	defer conn.Close()
	reader := bufio.NewReader(conn)
	codec, err := DetectCodec(reader)
	if err != nil {
		return fmt.Errorf("gateway transport detect codec: %w", err)
	}
	var writeMu sync.Mutex
	writers := make(map[sessionKey]*connSessionWriter)
	defer func() {
		if s.push == nil {
			return
		}
		for key := range writers {
			s.push.Unregister(key.authKeyId, key.sessionId)
		}
	}()
	for {
		frame, err := codec.ReadFrame(reader)
		if err != nil {
			return err
		}
		resp, err := s.handleFrame(ctx, conn, codec, &writeMu, writers, frame)
		if err != nil {
			return err
		}
		if resp == nil {
			continue
		}
		writeMu.Lock()
		if err := codec.WriteFrame(conn, resp); err != nil {
			writeMu.Unlock()
			return err
		}
		writeMu.Unlock()
	}
}

func (s *Server) handleFrame(ctx context.Context, conn net.Conn, codec Codec, writeMu *sync.Mutex, writers map[sessionKey]*connSessionWriter, frame []byte) ([]byte, error) {
	if len(frame) < 8 {
		return nil, fmt.Errorf("gateway transport: frame too short")
	}
	authKeyId := int64(binary.LittleEndian.Uint64(frame[:8]))
	if authKeyId == 0 {
		if s.handshake == nil {
			return nil, fmt.Errorf("gateway transport: handshake manager is nil")
		}
		msg, err := gmtproto.DecodePlainMessage(frame)
		if err != nil {
			return nil, err
		}
		return s.handshake.HandlePlain(ctx, msg)
	}
	if s.processor == nil {
		return nil, fmt.Errorf("gateway transport: session processor is nil")
	}
	return s.processor.HandleEncryptedWithSession(ctx, sessionstate.ConnInfo{GatewayId: s.gatewayID, ClientAddr: conn.RemoteAddr().String()}, frame, func(active sessionstate.ActiveSession) sessionstate.SeqNoAllocator {
		if s.push == nil || active.AuthKey == nil {
			return nil
		}
		key := sessionKey{authKeyId: active.AuthKeyId, sessionId: active.SessionId}
		writer := writers[key]
		if writer == nil {
			writer = &connSessionWriter{
				conn:    conn,
				codec:   codec,
				writeMu: writeMu,
			}
			writers[key] = writer
		}
		writer.Update(active.AuthKey, active.Salt)
		s.push.Register(push.LocalTarget{
			AuthKeyId: active.AuthKeyId,
			SessionId: active.SessionId,
			AuthKey:   active.AuthKey,
			Writer:    writer,
		})
		return writer
	})
}

func (s *Server) trackConn(conn net.Conn) func() {
	s.connsMu.Lock()
	if s.conns == nil {
		s.conns = make(map[net.Conn]struct{})
	}
	s.conns[conn] = struct{}{}
	s.connsMu.Unlock()
	return func() {
		s.connsMu.Lock()
		delete(s.conns, conn)
		s.connsMu.Unlock()
	}
}

func (s *Server) closeActiveConns() {
	s.connsMu.Lock()
	conns := make([]net.Conn, 0, len(s.conns))
	for conn := range s.conns {
		conns = append(conns, conn)
	}
	s.connsMu.Unlock()
	for _, conn := range conns {
		_ = conn.Close()
	}
}

type sessionKey struct {
	authKeyId int64
	sessionId int64
}

type connSessionWriter struct {
	conn    net.Conn
	codec   Codec
	writeMu *sync.Mutex

	mu      sync.Mutex
	authKey *crypto.AuthKey
	salt    int64
	seq     int32
}

func (w *connSessionWriter) Update(authKey *crypto.AuthKey, salt int64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.authKey = authKey
	w.salt = salt
}

func (w *connSessionWriter) NextSeqNo(contentRelated bool) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	seq := w.seq * 2
	if contentRelated {
		seq++
		w.seq++
	}
	return seq
}

func (w *connSessionWriter) WriteEncrypted(ctx context.Context, msg gmtproto.EncryptedMessage) error {
	w.mu.Lock()
	authKey := w.authKey
	if msg.Salt == 0 {
		msg.Salt = w.salt
	}
	w.mu.Unlock()
	if authKey == nil {
		return fmt.Errorf("gateway transport: session auth key is nil")
	}
	payload, err := gmtproto.EncodeEncryptedMessage(msg, authKey)
	if err != nil {
		return err
	}
	w.writeMu.Lock()
	defer w.writeMu.Unlock()
	return w.codec.WriteFrame(w.conn, payload)
}
