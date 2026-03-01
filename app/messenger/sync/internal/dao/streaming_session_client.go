// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	syncStreamSendBufSize       = 8192
	syncInitialReconnectBackoff = 100 * time.Millisecond
	syncMaxReconnectBackoff     = 10 * time.Second
)

// StreamingSession implements the same push interface as Session but uses
// bidirectional gRPC streaming instead of unary RPCs.
type StreamingSession struct {
	serverId    string
	conn        *grpc.ClientConn
	stream      grpc.BidiStreamingClient[session.SessionStreamRequest, session.SessionStreamResponse]
	sendCh      chan *session.SessionStreamRequest
	ctx         context.Context
	cancel      context.CancelFunc
	closed      atomic.Int32
	unavailable atomic.Bool
	reqCounter  atomic.Int64
}

func NewStreamingSession(addr string) (*StreamingSession, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithReadBufferSize(16*1024*1024),
		grpc.WithWriteBufferSize(16*1024*1024),
	)
	if err != nil {
		return nil, fmt.Errorf("dial session %s: %w", addr, err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	client := session.NewRPCSessionStreamClient(conn)
	stream, err := client.SessionDataStream(ctx)
	if err != nil {
		cancel()
		conn.Close()
		return nil, fmt.Errorf("open stream to session %s: %w", addr, err)
	}

	ss := &StreamingSession{
		serverId: addr,
		conn:     conn,
		stream:   stream,
		sendCh:   make(chan *session.SessionStreamRequest, syncStreamSendBufSize),
		ctx:      ctx,
		cancel:   cancel,
	}

	go ss.sendLoop()
	go ss.recvLoop()

	logx.Infof("StreamingSession: connected stream to session node %s", addr)
	return ss, nil
}

func (ss *StreamingSession) sendLoop() {
	for req := range ss.sendCh {
		if ss.closed.Load() != 0 {
			return
		}
		if err := ss.stream.Send(req); err != nil {
			logx.Errorf("StreamingSession sendLoop(%s) error: %v", ss.serverId, err)
			ss.unavailable.Store(true)
			return
		}
	}
}

func (ss *StreamingSession) recvLoop() {
	for {
		_, err := ss.stream.Recv()
		if err != nil {
			if ss.closed.Load() == 0 {
				logx.Errorf("StreamingSession recvLoop(%s) error: %v", ss.serverId, err)
				ss.unavailable.Store(true)
			}
			return
		}
		// fire-and-forget: acks are ignored
	}
}

func (ss *StreamingSession) nextRequestId() string {
	return fmt.Sprintf("%d", ss.reqCounter.Add(1))
}

func (ss *StreamingSession) PushUpdates(ctx context.Context, msg *session.TLSessionPushUpdatesData) error {
	if ss.unavailable.Load() {
		return fmt.Errorf("streaming session(%s) unavailable", ss.serverId)
	}

	req := &session.SessionStreamRequest{
		RequestId: ss.nextRequestId(),
		Payload:   &session.SessionStreamRequest_PushUpdates{PushUpdates: msg},
	}

	select {
	case ss.sendCh <- req:
		return nil
	default:
		return fmt.Errorf("streaming session(%s) sendCh full", ss.serverId)
	}
}

func (ss *StreamingSession) PushSessionUpdates(ctx context.Context, msg *session.TLSessionPushSessionUpdatesData) error {
	if ss.unavailable.Load() {
		return fmt.Errorf("streaming session(%s) unavailable", ss.serverId)
	}

	req := &session.SessionStreamRequest{
		RequestId: ss.nextRequestId(),
		Payload:   &session.SessionStreamRequest_PushSessionUpdates{PushSessionUpdates: msg},
	}

	select {
	case ss.sendCh <- req:
		return nil
	default:
		return fmt.Errorf("streaming session(%s) sendCh full", ss.serverId)
	}
}

func (ss *StreamingSession) PushRpcResult(ctx context.Context, msg *session.TLSessionPushRpcResultData) error {
	if ss.unavailable.Load() {
		return fmt.Errorf("streaming session(%s) unavailable", ss.serverId)
	}

	req := &session.SessionStreamRequest{
		RequestId: ss.nextRequestId(),
		Payload:   &session.SessionStreamRequest_PushRpcResult{PushRpcResult: msg},
	}

	select {
	case ss.sendCh <- req:
		return nil
	default:
		return fmt.Errorf("streaming session(%s) sendCh full", ss.serverId)
	}
}

func (ss *StreamingSession) Close() error {
	if ss.closed.CompareAndSwap(0, 1) {
		ss.cancel()
		close(ss.sendCh)
		ss.conn.Close()
	}
	return nil
}
