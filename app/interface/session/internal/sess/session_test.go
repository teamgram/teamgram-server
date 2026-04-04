package sess

import (
	"context"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
)

const testMTProtoLayer int32 = 223

func mustEncodeMTObject(t *testing.T, obj interface{ Encode(*bin.Encoder, int32) error }) []byte {
	t.Helper()

	x := bin.NewEncoder()
	defer x.End()

	if err := obj.Encode(x, testMTProtoLayer); err != nil {
		t.Fatalf("encode mt object: %v", err)
	}
	return append([]byte(nil), x.Bytes()...)
}

func mustDecodeMTObject(t *testing.T, body []byte) iface.TLObject {
	t.Helper()

	obj, err := iface.DecodeObject(bin.NewDecoder(body))
	if err != nil {
		t.Fatalf("decode mt object: %v", err)
	}
	return obj
}

func TestSessionConnNewRefreshesGatewayIDWhileOnline(t *testing.T) {
	s := newSession(1, &SessionList{})

	s.onSessionConnNew(context.Background(), "gateway-a")
	if got := s.getGatewayId(); got != "gateway-a" {
		t.Fatalf("expected initial gateway id gateway-a, got %q", got)
	}
	if s.connState != kStateOnline {
		t.Fatalf("expected session to be online after first connection, got %d", s.connState)
	}

	s.onSessionConnNew(context.Background(), "gateway-b")
	if got := s.getGatewayId(); got != "gateway-b" {
		t.Fatalf("expected gateway id to refresh to gateway-b, got %q", got)
	}
}

func TestSessionCloseIgnoresStaleGatewayAfterSwitch(t *testing.T) {
	s := newSession(1, &SessionList{})

	s.onSessionConnNew(context.Background(), "gateway-a")
	s.onSessionConnNew(context.Background(), "gateway-b")
	s.onSessionConnClose(context.Background(), "gateway-a")

	if s.connState != kStateOnline {
		t.Fatalf("expected stale gateway close to keep session online, got %d", s.connState)
	}
	if got := s.getGatewayId(); got != "gateway-b" {
		t.Fatalf("expected active gateway to remain gateway-b, got %q", got)
	}
}

func TestOnSyncRpcResultDataRemovesPendingAndQueuesResult(t *testing.T) {
	s := newSession(1, &SessionList{})

	const reqMsgID int64 = 1001
	s.pendingQueue.Add(reqMsgID)
	s.onSyncRpcResultData(context.Background(), reqMsgID, []byte{1, 2, 3})

	if got := s.pendingQueue.q.Len(); got != 0 {
		t.Fatalf("expected pending queue to be empty, got %d", got)
	}

	oMsg := s.outQueue.Lookup(reqMsgID)
	if oMsg == nil {
		t.Fatal("expected rpc result to be queued")
	}
	if oMsg.msg == nil {
		t.Fatal("expected raw message payload, got nil")
	}
	if oMsg.msg.Bytes != 3 {
		t.Fatalf("expected raw message bytes=3, got %d", oMsg.msg.Bytes)
	}
	if string(oMsg.msg.Body) != string([]byte{1, 2, 3}) {
		t.Fatalf("expected rpc result body to match input")
	}
	if _, ok := any(oMsg.msg).(*mt.TLMessageRawData); !ok {
		t.Fatal("expected queued message to be TLMessageRawData")
	}
}

func TestOnSyncRpcResultDataDoesNotDuplicateQueuedResultOnRetry(t *testing.T) {
	s := newSession(1, &SessionList{})

	const reqMsgID int64 = 2002
	s.pendingQueue.Add(reqMsgID)
	s.onSyncRpcResultData(context.Background(), reqMsgID, []byte{4, 5, 6})
	s.onSyncRpcResultData(context.Background(), reqMsgID, []byte{7, 8, 9})

	if got := s.outQueue.oMsgs.Len(); got != 1 {
		t.Fatalf("expected exactly one queued rpc result after retry, got %d", got)
	}

	oMsg := s.outQueue.Lookup(reqMsgID)
	if oMsg == nil {
		t.Fatal("expected queued rpc result after retry")
	}
	if string(oMsg.msg.Body) != string([]byte{4, 5, 6}) {
		t.Fatalf("expected first queued payload to be retained on retry")
	}
}

func TestOnSyncRpcResultDataIgnoresUnknownPendingRequest(t *testing.T) {
	s := newSession(1, &SessionList{})

	const reqMsgID int64 = 3003
	s.onSyncRpcResultData(context.Background(), reqMsgID, []byte{9, 8, 7})

	if got := s.pendingQueue.q.Len(); got != 0 {
		t.Fatalf("expected pending queue to stay empty, got %d", got)
	}
	if got := s.outQueue.oMsgs.Len(); got != 0 {
		t.Fatalf("expected unknown rpc result not to be queued, got %d queued messages", got)
	}
}

func TestOnMsgsAllInfoMarksRequestedMessagesForResend(t *testing.T) {
	s := newSession(1, &SessionList{})

	resendMsg := s.outQueue.AddNotifyMsg(4004, true, &mt.TLMessageRawData{
		MsgId: 5005,
		Body:  []byte{1},
		Bytes: 1,
	})
	resendMsg.sent = time.Now().Unix()

	ack := newInboxMsg(6006)
	s.onMsgsAllInfo(context.Background(), "", ack, &mt.TLMsgsAllInfo{
		MsgIds: []int64{4004},
		Info:   string([]byte{NOT_RECEIVED}),
	})

	if resendMsg.sent != 0 {
		t.Fatalf("expected NOT_RECEIVED message to be marked for resend, got sent=%d", resendMsg.sent)
	}
	if ack.state != RECEIVED|NEED_NO_ACK {
		t.Fatalf("expected msgs_all_info request to become no-ack receipt, got state=%d", ack.state)
	}
}

func TestOnMsgsStateInfoMarksRequestedMessagesForResend(t *testing.T) {
	s := newSession(1, &SessionList{})

	const (
		reqMsgID    int64 = 7007
		targetMsgID int64 = 7008
	)

	stateReqBody := mustEncodeMTObject(t, mt.MakeTLMsgsStateReq(&mt.TLMsgsStateReq{
		MsgIds: []int64{targetMsgID},
	}))
	reqStateMsg := s.outQueue.AddNotifyMsg(reqMsgID, true, &mt.TLMessageRawData{
		MsgId: reqMsgID,
		Body:  stateReqBody,
		Bytes: int32(len(stateReqBody)),
	})
	reqStateMsg.sent = time.Now().Unix()

	resendMsg := s.outQueue.AddNotifyMsg(targetMsgID, true, &mt.TLMessageRawData{
		MsgId: 8008,
		Body:  []byte{1},
		Bytes: 1,
	})
	resendMsg.sent = time.Now().Unix()

	ack := newInboxMsg(9009)
	s.onMsgsStateInfo(context.Background(), "", ack, &mt.TLMsgsStateInfo{
		ReqMsgId: reqMsgID,
		Info:     string([]byte{NOT_RECEIVED}),
	})

	if s.outQueue.Lookup(reqMsgID) != nil {
		t.Fatalf("expected original msgs_state_req request to be acked and removed")
	}
	if resendMsg.sent != 0 {
		t.Fatalf("expected NOT_RECEIVED state info to mark target for resend, got sent=%d", resendMsg.sent)
	}
	if ack.state != RECEIVED|NEED_NO_ACK {
		t.Fatalf("expected msgs_state_info request to become no-ack receipt, got state=%d", ack.state)
	}
}

func TestSessionStringHandlesBareSession(t *testing.T) {
	s := newSession(42, &SessionList{})

	got := s.String()
	if got == "" {
		t.Fatal("expected session String output for bare session")
	}
}

func TestOnMsgResendReqMarksKnownMessagesForResend(t *testing.T) {
	s := newSession(1, &SessionList{})

	const targetMsgID int64 = 10010
	s.inQueue.AddMsgId(targetMsgID)

	resendMsg := s.outQueue.AddNotifyMsg(targetMsgID, true, &mt.TLMessageRawData{
		MsgId: targetMsgID,
		Body:  []byte{1},
		Bytes: 1,
	})
	resendMsg.sent = time.Now().Unix()

	ack := newInboxMsg(10011)
	s.onMsgResendReq(context.Background(), "", ack, &mt.TLMsgResendReq{
		MsgIds: []int64{targetMsgID},
	})

	if resendMsg.sent != 0 {
		t.Fatalf("expected known resend request to mark message for resend, got sent=%d", resendMsg.sent)
	}
	if ack.state != RECEIVED|NEED_NO_ACK {
		t.Fatalf("expected msg_resend_req request to become no-ack receipt, got state=%d", ack.state)
	}
}

func TestOnMsgResendReqQueuesMsgsStateInfoForUnknownMessages(t *testing.T) {
	s := newSession(1, &SessionList{})

	ack := newInboxMsg(11011)
	s.onMsgResendReq(context.Background(), "", ack, &mt.TLMsgResendReq{
		MsgIds: []int64{11010},
	})

	queued := s.outQueue.Lookup(ack.msgId)
	if queued == nil {
		t.Fatal("expected msg_resend_req for unknown message to queue msgs_state_info")
	}

	msgsStateInfo, ok := mustDecodeMTObject(t, queued.msg.Body).(*mt.TLMsgsStateInfo)
	if !ok {
		t.Fatalf("expected queued object to be TLMsgsStateInfo")
	}
	if msgsStateInfo.ReqMsgId != ack.msgId {
		t.Fatalf("expected msgs_state_info req_msg_id=%d, got %d", ack.msgId, msgsStateInfo.ReqMsgId)
	}
	if []byte(msgsStateInfo.Info)[0] != NOT_RECEIVED {
		t.Fatalf("expected unknown resend target to be reported as NOT_RECEIVED, got %d", []byte(msgsStateInfo.Info)[0])
	}
	if ack.state != RECEIVED|NEED_NO_ACK {
		t.Fatalf("expected msg_resend_req request to become no-ack receipt, got state=%d", ack.state)
	}
}

func TestOnMsgsStateReqQueuesMsgsStateInfo(t *testing.T) {
	s := newSession(1, &SessionList{})

	const knownMsgID int64 = 12010
	known := s.inQueue.AddMsgId(knownMsgID)
	known.state = RECEIVED | ACKNOWLEDGED

	ack := newInboxMsg(12011)
	s.onMsgsStateReq(context.Background(), "", ack, &mt.TLMsgsStateReq{
		MsgIds: []int64{knownMsgID, knownMsgID + 1},
	})

	queued := s.outQueue.Lookup(ack.msgId)
	if queued == nil {
		t.Fatal("expected msgs_state_req to queue msgs_state_info response")
	}

	msgsStateInfo, ok := mustDecodeMTObject(t, queued.msg.Body).(*mt.TLMsgsStateInfo)
	if !ok {
		t.Fatalf("expected queued object to be TLMsgsStateInfo")
	}
	if msgsStateInfo.ReqMsgId != ack.msgId {
		t.Fatalf("expected msgs_state_info req_msg_id=%d, got %d", ack.msgId, msgsStateInfo.ReqMsgId)
	}

	info := []byte(msgsStateInfo.Info)
	if len(info) != 2 {
		t.Fatalf("expected 2 state bytes, got %d", len(info))
	}
	if info[0] != known.state {
		t.Fatalf("expected known message state byte=%d, got %d", known.state, info[0])
	}
	if info[1] != NOT_RECEIVED {
		t.Fatalf("expected future message state to be NOT_RECEIVED, got %d", info[1])
	}
	if ack.state != RECEIVED|NEED_NO_ACK {
		t.Fatalf("expected msgs_state_req request to become no-ack receipt, got state=%d", ack.state)
	}
}
