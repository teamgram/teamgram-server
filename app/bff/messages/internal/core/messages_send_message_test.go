package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// --- fake msg client ---

type fakeMsgSendClient struct {
	req *msg.TLMsgSendMessageV2
}

func (f *fakeMsgSendClient) MsgSendMessageV2(ctx context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
	f.req = in
	return tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{
		Out:      true,
		Id:       42,
		Pts:      1,
		PtsCount: 1,
		Date:     99,
	}).ToUpdates(), nil
}

func (f *fakeMsgSendClient) MsgReadHistory(ctx context.Context, in *msg.TLMsgReadHistory) (*tg.MessagesAffectedMessages, error) {
	return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{
		Pts:      in.MaxId,
		PtsCount: 1,
	}).ToMessagesAffectedMessages(), nil
}

func (f *fakeMsgSendClient) MsgReadHistoryV2(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error) {
	return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{
		Pts:      in.MaxId,
		PtsCount: 1,
	}).ToMessagesAffectedMessages(), nil
}

func (f *fakeMsgSendClient) MsgUpdatePinnedMessage(ctx context.Context, in *msg.TLMsgUpdatePinnedMessage) (*tg.Updates, error) {
	return tg.MakeTLUpdateShort(&tg.TLUpdateShort{
		Update: tg.MakeTLUpdatePinnedMessages(&tg.TLUpdatePinnedMessages{
			Pinned:   !in.Unpin,
			Messages: []int32{in.Id},
			Pts:      in.Id,
			PtsCount: 1,
		}),
		Date: 99,
	}).ToUpdates(), nil
}

func (f *fakeMsgSendClient) MsgUnpinAllMessages(ctx context.Context, in *msg.TLMsgUnpinAllMessages) (*tg.MessagesAffectedHistory, error) {
	return tg.MakeTLMessagesAffectedHistory(&tg.TLMessagesAffectedHistory{
		Pts:      1,
		PtsCount: 1,
	}).ToMessagesAffectedHistory(), nil
}

func (f *fakeMsgSendClient) MsgDeleteHistory(ctx context.Context, in *msg.TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error) {
	return tg.MakeTLMessagesAffectedHistory(&tg.TLMessagesAffectedHistory{
		Pts:      in.MaxId,
		PtsCount: 1,
	}).ToMessagesAffectedHistory(), nil
}

func (f *fakeMsgSendClient) MsgDeleteMessages(ctx context.Context, in *msg.TLMsgDeleteMessages) (*tg.MessagesAffectedMessages, error) {
	pts := int32(1)
	ptsCount := int32(1)
	if len(in.Id) > 0 {
		ptsCount = int32(len(in.Id))
		for _, id := range in.Id {
			if id > pts {
				pts = id
			}
		}
	}
	return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{
		Pts:      pts,
		PtsCount: ptsCount,
	}).ToMessagesAffectedMessages(), nil
}

func (f *fakeMsgSendClient) MsgReadMessageContents(ctx context.Context, in *msg.TLMsgReadMessageContents) (*tg.MessagesAffectedMessages, error) {
	pts := int32(1)
	ptsCount := int32(1)
	if len(in.Id) > 0 {
		ptsCount = int32(len(in.Id))
		for _, contentMsg := range in.Id {
			if contentMsg != nil && contentMsg.Id > pts {
				pts = contentMsg.Id
			}
		}
	}
	return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{
		Pts:      pts,
		PtsCount: ptsCount,
	}).ToMessagesAffectedMessages(), nil
}

func (f *fakeMsgSendClient) MsgEditMessageV2(ctx context.Context, in *msg.TLMsgEditMessageV2) (*tg.Updates, error) {
	return tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{
		Out:      true,
		Id:       in.DstMessage.MessageId,
		Pts:      1,
		PtsCount: 1,
		Date:     99,
	}).ToUpdates(), nil
}

var _ svc.MsgSendClient = (*fakeMsgSendClient)(nil)

// --- basic tests ---

func TestMessagesSendMessageRejectsEmptyMessage(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2, AccessHash: 0}),
		Message:  "",
		RandomId: 1,
	})
	if err != tg.ErrMessageEmpty {
		t.Fatalf("expected ErrMessageEmpty, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}

func TestMessagesSendMessageReturnsShortSentMessagePlaceholderForUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2, AccessHash: 0}),
		Message:  "hello",
		RandomId: 22,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected updates result, got nil")
	}

	updates, ok := result.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected updateShortSentMessage placeholder, got %T", result.Clazz)
	}
	if !updates.Out {
		t.Fatal("expected out=true")
	}
	if updates.Id != 22 {
		t.Fatalf("expected placeholder id=22, got %d", updates.Id)
	}
	if updates.Pts != 1 || updates.PtsCount != 1 {
		t.Fatalf("expected pts/pts_count to be 1/1, got %d/%d", updates.Pts, updates.PtsCount)
	}
	if updates.Date == 0 {
		t.Fatal("expected non-zero date")
	}
}

func TestMessagesSendMessageRejectsChannelPeerPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerChannel(&tg.TLInputPeerChannel{ChannelId: 3, AccessHash: 0}),
		Message:  "hello",
		RandomId: 3,
	})
	if err != tg.ErrEnterpriseIsBlocked {
		t.Fatalf("expected ErrEnterpriseIsBlocked, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}

func TestMessagesSendMessageRejectsEmptyPeerPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerEmpty(&tg.TLInputPeerEmpty{}),
		Message:  "hello",
		RandomId: 4,
	})
	if err != tg.ErrPeerIdInvalid {
		t.Fatalf("expected ErrPeerIdInvalid, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}

func TestMessagesSendMessageReusesPlaceholderIDForSameRandomID(t *testing.T) {
	c := New(context.Background(), nil)

	req := &tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2, AccessHash: 0}),
		Message:  "hello",
		RandomId: 77,
	}

	first, err := c.MessagesSendMessage(req)
	if err != nil {
		t.Fatalf("first send: expected nil error, got %v", err)
	}
	second, err := c.MessagesSendMessage(req)
	if err != nil {
		t.Fatalf("second send: expected nil error, got %v", err)
	}

	firstShort, ok := first.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected first result to be updateShortSentMessage, got %T", first.Clazz)
	}
	secondShort, ok := second.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected second result to be updateShortSentMessage, got %T", second.Clazz)
	}
	if firstShort.Id != secondShort.Id {
		t.Fatalf("expected same placeholder id for repeated random_id, got %d vs %d", firstShort.Id, secondShort.Id)
	}
}

// --- MsgClient delegation test ---

func TestMessagesSendMessageDelegatesToMsgClient(t *testing.T) {
	fakeMsgCli := &fakeMsgSendClient{}
	c := New(context.Background(), &svc.ServiceContext{
		MsgClient: fakeMsgCli,
	})

	result, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 30, AccessHash: 0}),
		Message:  "delegated",
		RandomId: 888,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	if fakeMsgCli.req == nil {
		t.Fatal("expected MsgSendMessageV2 to be called")
	}
	if fakeMsgCli.req.PeerId != 30 {
		t.Errorf("expected peerId=30, got %d", fakeMsgCli.req.PeerId)
	}
	if len(fakeMsgCli.req.Message) != 1 {
		t.Fatalf("expected 1 outbox message, got %d", len(fakeMsgCli.req.Message))
	}
	if fakeMsgCli.req.Message[0].RandomId != 888 {
		t.Errorf("expected randomId=888, got %d", fakeMsgCli.req.Message[0].RandomId)
	}

	updates, ok := result.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected updateShortSentMessage, got %T", result.Clazz)
	}
	if updates.Id != 42 {
		t.Errorf("expected id=42 from fake, got %d", updates.Id)
	}
}
