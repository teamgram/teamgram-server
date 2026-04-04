package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeInboxClient struct {
	req *inbox.TLInboxSendUserMessageToInboxV2
}

func (f *fakeInboxClient) InboxSendUserMessageToInboxV2(ctx context.Context, in *inbox.TLInboxSendUserMessageToInboxV2) (*tg.Void, error) {
	f.req = in
	return tg.MakeTLVoid(&tg.TLVoid{}).ToVoid(), nil
}

var _ svc.InboxPushClient = (*fakeInboxClient)(nil)

func TestMsgPushUserMessageRejectsNonUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgPushUserMessage(&msg.TLMsgPushUserMessage{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_CHAT,
		PeerId:    3,
		PushType:  1,
		Message: msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
			Message: tg.MakeTLMessage(&tg.TLMessage{Message: "x"}),
		}),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result != tg.BoolFalse {
		t.Fatalf("expected BoolFalse, got %v", result)
	}
}

func TestMsgPushUserMessageAcceptsUserPeerWithNilContext(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgPushUserMessage(&msg.TLMsgPushUserMessage{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_USER,
		PeerId:    3,
		PushType:  1,
		Message: msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
			Message: tg.MakeTLMessage(&tg.TLMessage{Message: "x"}),
		}),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result != tg.BoolTrue {
		t.Fatalf("expected BoolTrue, got %v", result)
	}
}

func TestMsgPushUserMessageForwardsToInboxClient(t *testing.T) {
	fakeInbox := &fakeInboxClient{}
	c := New(context.Background(), &svc.ServiceContext{
		InboxClient: fakeInbox,
	})

	result, err := c.MsgPushUserMessage(&msg.TLMsgPushUserMessage{
		UserId:    10,
		AuthKeyId: 20,
		PeerType:  tg.PEER_USER,
		PeerId:    30,
		PushType:  1,
		Message: msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
			Message: tg.MakeTLMessage(&tg.TLMessage{Message: "hello"}),
		}),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result != tg.BoolTrue {
		t.Fatalf("expected BoolTrue, got %v", result)
	}
	if fakeInbox.req == nil {
		t.Fatal("expected InboxSendUserMessageToInboxV2 to be called")
	}
	if fakeInbox.req.UserId != 30 {
		t.Errorf("expected inbox userId=30, got %d", fakeInbox.req.UserId)
	}
	if fakeInbox.req.FromId != 10 {
		t.Errorf("expected inbox fromId=10, got %d", fakeInbox.req.FromId)
	}
	if fakeInbox.req.FromAuthKeyId != 20 {
		t.Errorf("expected inbox fromAuthKeyId=20, got %d", fakeInbox.req.FromAuthKeyId)
	}
	if len(fakeInbox.req.BoxList) != 1 {
		t.Fatalf("expected 1 message box, got %d", len(fakeInbox.req.BoxList))
	}
}

func TestMsgPushUserMessageSelfPeer(t *testing.T) {
	fakeInbox := &fakeInboxClient{}
	c := New(context.Background(), &svc.ServiceContext{
		InboxClient: fakeInbox,
	})

	result, err := c.MsgPushUserMessage(&msg.TLMsgPushUserMessage{
		UserId:    10,
		AuthKeyId: 20,
		PeerType:  tg.PEER_SELF,
		PeerId:    10,
		PushType:  0,
		Message: msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
			Message: tg.MakeTLMessage(&tg.TLMessage{Message: "self"}),
		}),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result != tg.BoolTrue {
		t.Fatalf("expected BoolTrue, got %v", result)
	}
	if fakeInbox.req == nil {
		t.Fatal("expected inbox call for self peer")
	}
}
