package transport

import (
	"bufio"
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"sync"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/sessionstate"
)

type Server struct {
	addr      string
	handshake *sessionstate.HandshakeManager
	processor *sessionstate.Processor
	listener  net.Listener
	wg        sync.WaitGroup
}

func NewServer(addr string, handshake *sessionstate.HandshakeManager, processor *sessionstate.Processor) *Server {
	return &Server{addr: addr, handshake: handshake, processor: processor}
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
	s.wg.Wait()
	return err
}

func (s *Server) ServeConn(ctx context.Context, conn net.Conn) error {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	codec, err := DetectCodec(reader)
	if err != nil {
		return fmt.Errorf("gateway transport detect codec: %w", err)
	}
	for {
		frame, err := codec.ReadFrame(reader)
		if err != nil {
			return err
		}
		resp, err := s.handleFrame(ctx, conn, frame)
		if err != nil {
			return err
		}
		if resp == nil {
			continue
		}
		if err := codec.WriteFrame(conn, resp); err != nil {
			return err
		}
	}
}

func (s *Server) handleFrame(ctx context.Context, conn net.Conn, frame []byte) ([]byte, error) {
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
	return s.processor.HandleEncrypted(ctx, sessionstate.ConnInfo{ClientAddr: conn.RemoteAddr().String()}, frame)
}
