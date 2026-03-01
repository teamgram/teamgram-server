// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/config"
	"github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	streamSendBufSize       = 8192
	maxReconnectBackoff     = 10 * time.Second
	initialReconnectBackoff = 100 * time.Millisecond
	queryAuthKeyTimeout     = 5 * time.Second
	streamMaxNodeFailures   = 3
)

type sessionNodeStream struct {
	nodeAddr string
	conn     *grpc.ClientConn
	stream   grpc.BidiStreamingClient[session.SessionStreamRequest, session.SessionStreamResponse]
	sendCh   chan *session.SessionStreamRequest
	closed   atomic.Int32
	cancel   context.CancelFunc
}

type StreamingSessionDispatcher struct {
	mu           sync.RWMutex
	dispatcher   *hash.ConsistentHash
	streams      map[string]*sessionNodeStream
	failCounters map[string]int
	pendingReqs  sync.Map // requestId → chan *session.SessionStreamResponse
	reqIdCounter atomic.Int64
	sessionConf  config.Config
}

func NewStreamingSessionDispatcher(c config.Config) *StreamingSessionDispatcher {
	d := &StreamingSessionDispatcher{
		dispatcher:   hash.NewConsistentHash(),
		streams:      make(map[string]*sessionNodeStream),
		failCounters: make(map[string]int),
		sessionConf:  c,
	}
	d.watch(c.Session)
	return d
}

func (d *StreamingSessionDispatcher) watch(c zrpc.RpcClientConf) {
	sub, err := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key)
	if err != nil {
		logx.Errorf("StreamingSessionDispatcher watch NewSubscriber(%+v) error: %v", c.Etcd, err)
		return
	}

	update := func() {
		var (
			addNodes    []string
			removeNodes []string
		)

		d.mu.Lock()
		defer d.mu.Unlock()

		values := sub.Values()
		newStreams := map[string]*sessionNodeStream{}
		for _, v := range values {
			if old, ok := d.streams[v]; ok {
				newStreams[v] = old
				continue
			}
			addNodes = append(addNodes, v)
		}

		for key := range d.streams {
			if !stringx.Contains(values, key) {
				removeNodes = append(removeNodes, key)
			}
		}

		for _, n := range addNodes {
			ns, err := d.createNodeStream(n)
			if err != nil {
				logx.Errorf("StreamingSessionDispatcher createNodeStream(%s) error: %v", n, err)
				continue
			}
			newStreams[n] = ns
			d.dispatcher.Add(n)
		}

		for _, n := range removeNodes {
			if old, ok := d.streams[n]; ok {
				old.close()
			}
			d.dispatcher.Remove(n)
			delete(d.failCounters, n)
		}

		d.streams = newStreams
	}

	sub.AddListener(update)
	update()
}

func (d *StreamingSessionDispatcher) createNodeStream(addr string) (*sessionNodeStream, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithReadBufferSize(16*1024*1024),
		grpc.WithWriteBufferSize(16*1024*1024),
	)
	if err != nil {
		return nil, fmt.Errorf("dial %s: %w", addr, err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	client := session.NewRPCSessionStreamClient(conn)
	stream, err := client.SessionDataStream(ctx)
	if err != nil {
		cancel()
		conn.Close()
		return nil, fmt.Errorf("open stream to %s: %w", addr, err)
	}

	ns := &sessionNodeStream{
		nodeAddr: addr,
		conn:     conn,
		stream:   stream,
		sendCh:   make(chan *session.SessionStreamRequest, streamSendBufSize),
		cancel:   cancel,
	}

	go d.sendLoop(ns)
	go d.recvLoop(ns)

	logx.Infof("StreamingSessionDispatcher: connected stream to session node %s", addr)
	return ns, nil
}

func (ns *sessionNodeStream) close() {
	if ns.closed.CompareAndSwap(0, 1) {
		ns.cancel()
		close(ns.sendCh)
		ns.conn.Close()
	}
}

func (d *StreamingSessionDispatcher) sendLoop(ns *sessionNodeStream) {
	for req := range ns.sendCh {
		if ns.closed.Load() != 0 {
			return
		}
		if err := ns.stream.Send(req); err != nil {
			logx.Errorf("StreamingSessionDispatcher sendLoop(%s) Send error: %v", ns.nodeAddr, err)
			d.handleStreamError(ns)
			return
		}
	}
}

func (d *StreamingSessionDispatcher) recvLoop(ns *sessionNodeStream) {
	for {
		resp, err := ns.stream.Recv()
		if err != nil {
			if ns.closed.Load() == 0 {
				logx.Errorf("StreamingSessionDispatcher recvLoop(%s) Recv error: %v", ns.nodeAddr, err)
				d.handleStreamError(ns)
			}
			return
		}

		reqId := resp.GetRequestId()
		if reqId == "" {
			continue
		}

		switch p := resp.GetPayload().(type) {
		case *session.SessionStreamResponse_Ack:
			// fire-and-forget ack, resolve pending if any
			if ch, ok := d.pendingReqs.LoadAndDelete(reqId); ok {
				select {
				case ch.(chan *session.SessionStreamResponse) <- resp:
				default:
				}
			}
		case *session.SessionStreamResponse_Error:
			if p.Error.GetCode() == 700 {
				// redirect: parse target from message "REDIRECT_TO_<addr>"
				msg := p.Error.GetMessage()
				const prefix = "REDIRECT_TO_"
				if strings.HasPrefix(msg, prefix) {
					target := msg[len(prefix):]
					d.handleRedirect(reqId, target)
				}
			} else {
				if ch, ok := d.pendingReqs.LoadAndDelete(reqId); ok {
					select {
					case ch.(chan *session.SessionStreamResponse) <- resp:
					default:
					}
				}
			}
		case *session.SessionStreamResponse_AuthKey:
			if ch, ok := d.pendingReqs.LoadAndDelete(reqId); ok {
				select {
				case ch.(chan *session.SessionStreamResponse) <- resp:
				default:
				}
			}
		}
	}
}

func (d *StreamingSessionDispatcher) handleRedirect(reqId, target string) {
	// Redirect is for fire-and-forget requests stored temporarily.
	// We look up the original request from pendingReqs if available.
	// For fire-and-forget (SendData/CloseSession), we just log and ignore
	// since the request was already sent and we can't replay it easily.
	logx.Infof("StreamingSessionDispatcher: redirect reqId=%s to %s (fire-and-forget, ignoring)", reqId, target)
	d.pendingReqs.Delete(reqId)
}

func (d *StreamingSessionDispatcher) handleStreamError(ns *sessionNodeStream) {
	addr := ns.nodeAddr
	ns.close()

	d.mu.Lock()
	d.failCounters[addr]++
	failCount := d.failCounters[addr]
	if failCount >= streamMaxNodeFailures {
		logx.Errorf("StreamingSessionDispatcher: node %s unreachable (%d failures), removing", addr, failCount)
		d.dispatcher.Remove(addr)
		delete(d.streams, addr)
		delete(d.failCounters, addr)
		d.mu.Unlock()
		return
	}
	d.mu.Unlock()

	// try reconnect with backoff
	go d.reconnect(addr)
}

func (d *StreamingSessionDispatcher) reconnect(addr string) {
	backoff := initialReconnectBackoff
	for {
		time.Sleep(backoff)

		d.mu.Lock()
		// check if node was removed while we were sleeping
		if _, ok := d.streams[addr]; !ok {
			// check if it's still in the hash ring
			if _, ok := d.dispatcher.Get(addr); !ok {
				d.mu.Unlock()
				logx.Infof("StreamingSessionDispatcher: node %s removed, stopping reconnect", addr)
				return
			}
		}
		d.mu.Unlock()

		ns, err := d.createNodeStream(addr)
		if err != nil {
			logx.Errorf("StreamingSessionDispatcher reconnect(%s) error: %v, backoff: %v", addr, err, backoff)
			backoff = backoff * 2
			if backoff > maxReconnectBackoff {
				backoff = maxReconnectBackoff
			}
			continue
		}

		d.mu.Lock()
		d.streams[addr] = ns
		delete(d.failCounters, addr)
		d.mu.Unlock()

		logx.Infof("StreamingSessionDispatcher: reconnected to %s", addr)
		return
	}
}

func (d *StreamingSessionDispatcher) nextRequestId() string {
	return strconv.FormatInt(d.reqIdCounter.Add(1), 10)
}

func (d *StreamingSessionDispatcher) getNodeStream(permAuthKeyId int64) (*sessionNodeStream, error) {
	key := strconv.FormatInt(permAuthKeyId, 10)

	d.mu.RLock()
	val, ok := d.dispatcher.Get(key)
	if !ok {
		d.mu.RUnlock()
		return nil, ErrSessionNotFound
	}
	node := val.(string)
	ns, ok := d.streams[node]
	d.mu.RUnlock()

	if !ok {
		return nil, ErrSessionNotFound
	}
	return ns, nil
}

func (d *StreamingSessionDispatcher) SendData(ctx context.Context, permAuthKeyId int64, in *session.TLSessionSendDataToSession) error {
	ns, err := d.getNodeStream(permAuthKeyId)
	if err != nil {
		return err
	}

	req := &session.SessionStreamRequest{
		RequestId: d.nextRequestId(),
		Payload:   &session.SessionStreamRequest_SendData{SendData: in},
	}

	select {
	case ns.sendCh <- req:
		return nil
	default:
		return fmt.Errorf("StreamingSessionDispatcher: sendCh full for node %s", ns.nodeAddr)
	}
}

func (d *StreamingSessionDispatcher) CloseSession(ctx context.Context, permAuthKeyId int64, in *session.TLSessionCloseSession) error {
	ns, err := d.getNodeStream(permAuthKeyId)
	if err != nil {
		return err
	}

	req := &session.SessionStreamRequest{
		RequestId: d.nextRequestId(),
		Payload:   &session.SessionStreamRequest_CloseSession{CloseSession: in},
	}

	select {
	case ns.sendCh <- req:
		return nil
	default:
		return fmt.Errorf("StreamingSessionDispatcher: sendCh full for node %s", ns.nodeAddr)
	}
}

func (d *StreamingSessionDispatcher) QueryAuthKey(ctx context.Context, authKeyId int64, in *session.TLSessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	ns, err := d.getNodeStream(authKeyId)
	if err != nil {
		return nil, err
	}

	reqId := d.nextRequestId()
	respCh := make(chan *session.SessionStreamResponse, 1)
	d.pendingReqs.Store(reqId, respCh)
	defer d.pendingReqs.Delete(reqId)

	req := &session.SessionStreamRequest{
		RequestId: reqId,
		Payload:   &session.SessionStreamRequest_QueryAuthKey{QueryAuthKey: in},
	}

	select {
	case ns.sendCh <- req:
	default:
		return nil, fmt.Errorf("StreamingSessionDispatcher: sendCh full for node %s", ns.nodeAddr)
	}

	select {
	case resp := <-respCh:
		switch p := resp.GetPayload().(type) {
		case *session.SessionStreamResponse_AuthKey:
			return p.AuthKey, nil
		case *session.SessionStreamResponse_Error:
			return nil, fmt.Errorf("session stream error: code=%d, msg=%s", p.Error.GetCode(), p.Error.GetMessage())
		default:
			return nil, fmt.Errorf("unexpected response type for QueryAuthKey")
		}
	case <-time.After(queryAuthKeyTimeout):
		return nil, fmt.Errorf("QueryAuthKey timeout for authKeyId=%d", authKeyId)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (d *StreamingSessionDispatcher) Close() {
	d.mu.Lock()
	defer d.mu.Unlock()

	for _, ns := range d.streams {
		ns.close()
	}
	d.streams = make(map[string]*sessionNodeStream)
}
