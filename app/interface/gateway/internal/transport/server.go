package transport

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	gatewaypresence "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/presence"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/push"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/sessionstate"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"

	"github.com/zeromicro/go-zero/core/logx"
)

type Server struct {
	addr              string
	gatewayID         string
	gatewayRPCAddr    string
	gatewayGeneration string
	handshake         *sessionstate.HandshakeManager
	processor         *sessionstate.Processor
	push              *push.LocalWriter
	presence          *gatewaypresence.Registrar
	listener          net.Listener
	wg                sync.WaitGroup
	connsMu           sync.Mutex
	conns             map[net.Conn]struct{}
	stopping          bool
	eventSink         func(transportEvent)
}

func NewServer(addr string, gatewayID string, gatewayRPCAddr string, gatewayGeneration string, handshake *sessionstate.HandshakeManager, processor *sessionstate.Processor, pushWriter *push.LocalWriter, registrar *gatewaypresence.Registrar) *Server {
	return &Server{addr: addr, gatewayID: gatewayID, gatewayRPCAddr: gatewayRPCAddr, gatewayGeneration: gatewayGeneration, handshake: handshake, processor: processor, push: pushWriter, presence: registrar, conns: make(map[net.Conn]struct{})}
}

type transportEvent struct {
	Phase  string
	Remote string
	Err    error
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
			untrack, ok := s.trackConn(conn)
			if !ok {
				_ = conn.Close()
				continue
			}
			s.wg.Add(1)
			go func() {
				defer s.wg.Done()
				_ = s.serveTrackedConn(ctx, conn, untrack)
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
	s.connsMu.Lock()
	s.stopping = true
	s.connsMu.Unlock()
	if s.listener != nil {
		err = s.listener.Close()
	}
	s.closeActiveConns()
	s.wg.Wait()
	return err
}

func (s *Server) ServeConn(ctx context.Context, conn net.Conn) error {
	untrack, ok := s.trackConn(conn)
	if !ok {
		_ = conn.Close()
		return nil
	}
	return s.serveTrackedConn(ctx, conn, untrack)
}

func (s *Server) serveTrackedConn(ctx context.Context, conn net.Conn, untrack func()) error {
	defer untrack()
	handler := newGatewayHandler(s, conn)
	err := newNetDriver(conn, handler).Serve(ctx)
	if err != nil && !handler.wasClosed() {
		phase := "connection"
		if errors.Is(err, ErrUnsupportedTransport) {
			phase = "detect_codec"
		}
		s.reportConnError(conn, phase, err)
	}
	return err
}

type gatewayHandler struct {
	server         *Server
	rawConn        net.Conn
	writers        map[sessionKey]*sessionWriter
	activeSessions map[sessionKey]sessionstate.ActiveSession
	closeMu        sync.Mutex
	closed         bool
	reported       bool
}

func newGatewayHandler(server *Server, rawConn net.Conn) *gatewayHandler {
	return &gatewayHandler{
		server:         server,
		rawConn:        rawConn,
		writers:        make(map[sessionKey]*sessionWriter),
		activeSessions: make(map[sessionKey]sessionstate.ActiveSession),
	}
}

func (h *gatewayHandler) OnOpen(ctx context.Context, conn Connection) {
}

func (h *gatewayHandler) OnFrame(ctx context.Context, conn Connection, frame []byte) error {
	resp, err := h.server.handleFrame(ctx, conn, h.writers, h.activeSessions, frame)
	if err != nil {
		h.server.reportConnError(h.rawConn, "handle_frame", err)
		h.markReported()
		return err
	}
	if resp == nil {
		return nil
	}
	if err := conn.WriteFrame(ctx, resp); err != nil {
		h.server.reportConnError(h.rawConn, "write_frame", err)
		h.markReported()
		return err
	}
	return nil
}

func (h *gatewayHandler) OnClose(ctx context.Context, conn Connection, err error) {
	h.closeMu.Lock()
	h.closed = true
	shouldReport := !h.reported
	h.closeMu.Unlock()
	if shouldReport {
		h.server.reportConnError(h.rawConn, "connection", err)
	}
	for key := range h.writers {
		if h.server.push != nil {
			h.server.push.Unregister(key.authKeyId, key.authKeyType, key.sessionId)
		}
		if h.server.presence != nil {
			h.server.presence.Unregister(context.Background(), key.authKeyId, key.sessionId)
		}
		if h.server.processor != nil {
			h.server.processor.UnregisterSession(h.activeSessions[key])
		}
	}
}

func (h *gatewayHandler) wasClosed() bool {
	h.closeMu.Lock()
	defer h.closeMu.Unlock()
	return h.closed
}

func (h *gatewayHandler) markReported() {
	h.closeMu.Lock()
	defer h.closeMu.Unlock()
	h.reported = true
}

func (s *Server) handleFrame(ctx context.Context, conn Connection, writers map[sessionKey]*sessionWriter, activeSessions map[sessionKey]sessionstate.ActiveSession, frame []byte) ([]byte, error) {
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
	register := func(active sessionstate.ActiveSession, registerPresence bool) sessionstate.SeqNoAllocator {
		if active.AuthKey == nil {
			return nil
		}
		key := sessionKey{authKeyId: active.AuthKeyId, authKeyType: active.AuthKeyType, sessionId: active.SessionId}
		writer := writers[key]
		if s.push != nil && writer == nil {
			writer = &sessionWriter{conn: conn}
		}
		if _, ok := writers[key]; !ok {
			writers[key] = writer
		}
		activeSessions[key] = active
		if s.push != nil && writer != nil {
			writer.Update(active.AuthKey, active.Salt)
			s.push.Register(push.LocalTarget{
				PermAuthKeyId: active.PermAuthKeyId,
				AuthKeyId:     active.AuthKeyId,
				AuthKeyType:   active.AuthKeyType,
				SessionId:     active.SessionId,
				Layer:         active.Layer,
				AuthKey:       active.AuthKey,
				Writer:        writer,
				MainUpdates:   active.MainUpdates,
			})
		}
		if registerPresence && s.presence != nil && active.UserId > 0 {
			s.presence.Register(ctx, gatewaypresence.ActiveSession{
				UserID:        active.UserId,
				PermAuthKeyID: active.PermAuthKeyId,
				AuthKeyID:     active.AuthKeyId,
				AuthKeyType:   active.AuthKeyType,
				SessionID:     active.SessionId,
				Layer:         active.Layer,
				Client:        active.Client,
			})
		}
		if s.push == nil {
			return nil
		}
		return writer
	}
	return s.processor.HandleEncryptedWithSessionRefresh(
		ctx,
		sessionstate.ConnInfo{GatewayId: s.gatewayID, GatewayRpcAddr: s.gatewayRPCAddr, GatewayGeneration: s.gatewayGeneration, ClientAddr: conn.RemoteAddr()},
		frame,
		func(active sessionstate.ActiveSession) sessionstate.SeqNoAllocator {
			return register(active, true)
		},
		func(active sessionstate.ActiveSession) {
			_ = register(active, false)
		},
	)
}

func (s *Server) trackConn(conn net.Conn) (func(), bool) {
	s.connsMu.Lock()
	if s.stopping {
		s.connsMu.Unlock()
		return func() {}, false
	}
	if s.conns == nil {
		s.conns = make(map[net.Conn]struct{})
	}
	s.conns[conn] = struct{}{}
	s.connsMu.Unlock()
	return func() {
		s.connsMu.Lock()
		delete(s.conns, conn)
		s.connsMu.Unlock()
	}, true
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

func (s *Server) reportConnError(conn net.Conn, phase string, err error) {
	if err == nil {
		return
	}
	remote := ""
	if conn != nil && conn.RemoteAddr() != nil {
		remote = conn.RemoteAddr().String()
	}
	event := transportEvent{Phase: phase, Remote: remote, Err: err}
	if s != nil && s.eventSink != nil {
		s.eventSink(event)
	}
	if errors.Is(err, io.EOF) || errors.Is(err, net.ErrClosed) {
		logx.Debugf("gateway transport connection closed: phase=%s remote=%s err=%v", phase, remote, err)
		return
	}
	logx.Errorf("gateway transport connection error: phase=%s remote=%s err=%v", phase, remote, err)
}

type sessionKey struct {
	authKeyId   int64
	authKeyType int32
	sessionId   int64
}

type sessionWriter struct {
	conn Connection

	mu      sync.Mutex
	authKey *crypto.AuthKey
	salt    int64
	seq     int32
}

func (w *sessionWriter) Update(authKey *crypto.AuthKey, salt int64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.authKey = authKey
	w.salt = salt
}

func (w *sessionWriter) NextSeqNo(contentRelated bool) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	seq := w.seq * 2
	if contentRelated {
		seq++
		w.seq++
	}
	return seq
}

func (w *sessionWriter) WriteEncrypted(ctx context.Context, msg gmtproto.EncryptedMessage) error {
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
	return w.conn.WriteFrame(ctx, payload)
}
