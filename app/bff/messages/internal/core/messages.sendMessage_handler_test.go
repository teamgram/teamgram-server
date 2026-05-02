package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/svc"
	msgclient "github.com/teamgram/teamgram-server/v2/app/messenger/msg/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type messagesFakeMsgClient struct {
	msgclient.MsgClient
	sendMessageV2 func(ctx context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error)
	getHistory    func(ctx context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error)
	readHistoryV2 func(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error)
}

func (f *messagesFakeMsgClient) MsgSendMessageV2(ctx context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
	return f.sendMessageV2(ctx, in)
}

func (f *messagesFakeMsgClient) MsgGetHistory(ctx context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error) {
	return f.getHistory(ctx, in)
}

func (f *messagesFakeMsgClient) MsgReadHistoryV2(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error) {
	return f.readHistoryV2(ctx, in)
}

func newSendMsgCore(client msgclient.MsgClient, selfID, authKeyID int64) *MessagesCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{MsgClient: client},
	})
	c.MD = &metadata.RpcMetadata{
		UserId:        selfID,
		PermAuthKeyId: authKeyID,
	}
	return c
}

func testUpdates() *tg.Updates {
	return &tg.Updates{
		Clazz: tg.MakeTLUpdates(&tg.TLUpdates{
			Updates: []tg.UpdateClazz{},
			Users:   []tg.UserClazz{},
			Chats:   []tg.ChatClazz{},
			Date:    1000000,
			Seq:     0,
		}),
	}
}

func inputPeerUser(userID int64) *tg.TLInputPeerUser {
	return &tg.TLInputPeerUser{UserId: userID}
}

func inputPeerSelf() *tg.TLInputPeerSelf {
	return &tg.TLInputPeerSelf{}
}

func inputPeerChat(chatID int64) *tg.TLInputPeerChat {
	return &tg.TLInputPeerChat{ChatId: chatID}
}

// --- Success ---

func TestMessagesSendMessage_Success(t *testing.T) {
	var got *msg.TLMsgSendMessageV2
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	r, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if got == nil {
		t.Fatal("msg service was not called")
	}
	if got.UserId != 100 {
		t.Fatalf("UserId = %d, want 100", got.UserId)
	}
	if got.AuthKeyId != 200 {
		t.Fatalf("AuthKeyId = %d, want 200", got.AuthKeyId)
	}
	if got.PeerType != payload.PeerTypeUser {
		t.Fatalf("PeerType = %d, want %d", got.PeerType, payload.PeerTypeUser)
	}
	if got.PeerId != 300 {
		t.Fatalf("PeerId = %d, want 300", got.PeerId)
	}
	if len(got.Message) == 0 || got.Message[0] == nil {
		t.Fatal("OutboxMessage is nil or empty")
	}
	if r == nil {
		t.Fatal("result is nil")
	}
}

func TestMessagesGetHistory_UserPeerSuccess(t *testing.T) {
	var got *msg.TLMsgGetHistory
	reply := tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: []tg.MessageClazz{
			tg.MakeTLMessage(&tg.TLMessage{Id: 5, Message: "hello"}),
		},
		Chats: []tg.ChatClazz{},
		Users: []tg.UserClazz{},
	}).ToMessagesMessages()
	c := newSendMsgCore(&messagesFakeMsgClient{
		getHistory: func(_ context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error) {
			got = in
			return reply, nil
		},
	}, 100, 200)

	r, err := c.MessagesGetHistory(&tg.TLMessagesGetHistory{
		Peer:       inputPeerUser(300),
		OffsetId:   7,
		OffsetDate: 8,
		AddOffset:  9,
		Limit:      10,
		MaxId:      11,
		MinId:      12,
		Hash:       13,
	})
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if r != reply {
		t.Fatalf("reply mismatch: got %p want %p", r, reply)
	}
	if got == nil {
		t.Fatal("msg service was not called")
	}
	if got.UserId != 100 || got.AuthKeyId != 200 || got.PeerType != payload.PeerTypeUser || got.PeerId != 300 {
		t.Fatalf("unexpected service identity/peer: %+v", got)
	}
	if got.OffsetId != 7 || got.OffsetDate != 8 || got.AddOffset != 9 || got.Limit != 10 || got.MaxId != 11 || got.MinId != 12 || got.Hash != 13 {
		t.Fatalf("unexpected paging input: %+v", got)
	}
}

func TestMessagesGetHistory_InputPeerSelfTargetsCurrentUser(t *testing.T) {
	var got *msg.TLMsgGetHistory
	reply := tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: []tg.MessageClazz{},
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
	}).ToMessagesMessages()
	c := newSendMsgCore(&messagesFakeMsgClient{
		getHistory: func(_ context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error) {
			got = in
			return reply, nil
		},
	}, 100, 200)

	if _, err := c.MessagesGetHistory(&tg.TLMessagesGetHistory{
		Peer:  inputPeerSelf(),
		Limit: 30,
	}); err != nil {
		t.Fatalf("error = %v", err)
	}
	if got == nil || got.UserId != 100 || got.PeerId != 100 || got.Limit != 30 {
		t.Fatalf("unexpected history request: %+v", got)
	}
}

func TestMessagesReadHistory_InputPeerSelfSuccess(t *testing.T) {
	var got *msg.TLMsgReadHistoryV2
	reply := tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{
		Pts:      3,
		PtsCount: 0,
	}).ToMessagesAffectedMessages()
	c := newSendMsgCore(&messagesFakeMsgClient{
		readHistoryV2: func(_ context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error) {
			got = in
			return reply, nil
		},
	}, 100, 200)

	r, err := c.MessagesReadHistory(&tg.TLMessagesReadHistory{
		Peer:  inputPeerSelf(),
		MaxId: 2,
	})
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if r != reply {
		t.Fatalf("reply mismatch: got %p want %p", r, reply)
	}
	if got == nil || got.UserId != 100 || got.AuthKeyId != 200 || got.PeerType != payload.PeerTypeUser || got.PeerId != 100 || got.MaxId != 2 {
		t.Fatalf("unexpected read history request: %+v", got)
	}
}

func TestMessagesReadHistory_NonUserPeerRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		readHistoryV2: func(context.Context, *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesReadHistory(&tg.TLMessagesReadHistory{
		Peer:  inputPeerChat(300),
		MaxId: 2,
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

// --- Input validation (must NOT call msg) ---

func TestMessagesSendMessage_InputPeerSelfTargetsCurrentUser(t *testing.T) {
	var got *msg.TLMsgSendMessageV2
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	if _, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerSelf(),
		Message:  "hello",
		RandomId: 42,
	}); err != nil {
		t.Fatalf("error = %v", err)
	}
	if got == nil || got.UserId != 100 || got.PeerId != 100 {
		t.Fatalf("unexpected msg request: %+v", got)
	}
}

func TestMessagesSendMessage_NilPeerRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_NonUserPeerRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerChat(42),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_EmptyMessageRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "",
		RandomId: 42,
	})
	if err != tg.ErrMessageEmpty {
		t.Fatalf("error = %v, want %v", err, tg.ErrMessageEmpty)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_WhitespaceMessageRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "   ",
		RandomId: 42,
	})
	if err != tg.ErrMessageEmpty {
		t.Fatalf("error = %v, want %v", err, tg.ErrMessageEmpty)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_MessageTooLongRejected(t *testing.T) {
	text := ""
	for i := 0; i < 4097; i++ {
		text += "a"
	}
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  text,
		RandomId: 42,
	})
	if err != tg.ErrMessageTooLong {
		t.Fatalf("error = %v, want %v", err, tg.ErrMessageTooLong)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_Message4096CodeUnitsAccepted(t *testing.T) {
	text := ""
	for i := 0; i < 4096; i++ {
		text += "a"
	}
	var got *msg.TLMsgSendMessageV2
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  text,
		RandomId: 42,
	})
	if err != nil {
		t.Fatalf("error = %v, want nil", err)
	}
	if got == nil {
		t.Fatal("msg service was not called")
	}
}

func TestMessagesSendMessage_RandomIdZeroRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 0,
	})
	if err != tg.ErrRandomIdEmpty {
		t.Fatalf("error = %v, want %v", err, tg.ErrRandomIdEmpty)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

// --- Unsupported field rejection ---

func TestMessagesSendMessage_EntitiesRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
		Entities: []tg.MessageEntityClazz{tg.MakeTLMessageEntityBold(&tg.TLMessageEntityBold{})},
	})
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrInputRequestInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_SilentTrueRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
		Silent:   true,
	})
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrInputRequestInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_NoforwardsTrueRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:       inputPeerUser(300),
		Message:    "hello",
		RandomId:   42,
		Noforwards: true,
	})
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrInputRequestInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_ScheduleDateRejected(t *testing.T) {
	called := false
	sched := int32(2000000)
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:         inputPeerUser(300),
		Message:      "hello",
		RandomId:     42,
		ScheduleDate: &sched,
	})
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrInputRequestInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_ReplyMarkupRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:        inputPeerUser(300),
		Message:     "hello",
		RandomId:    42,
		ReplyMarkup: tg.MakeTLReplyKeyboardMarkup(&tg.TLReplyKeyboardMarkup{}),
	})
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrInputRequestInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

// --- Default/empty fields accepted ---

func TestMessagesSendMessage_EmptyEntitiesAccepted(t *testing.T) {
	var got *msg.TLMsgSendMessageV2
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
		Entities: []tg.MessageEntityClazz{},
	})
	if err != nil {
		t.Fatalf("error = %v, want nil", err)
	}
	if got == nil {
		t.Fatal("msg service was not called")
	}
}

func TestMessagesSendMessage_SilentFalseAccepted(t *testing.T) {
	var got *msg.TLMsgSendMessageV2
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
		Silent:   false,
	})
	if err != nil {
		t.Fatalf("error = %v, want nil", err)
	}
	if got == nil {
		t.Fatal("msg service was not called")
	}
}

// --- Error mapping ---

func TestMessagesSendMessage_RandomIdConflictMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			return nil, msg.ErrRandomIdConflict
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrRandomIdDuplicate {
		t.Fatalf("error = %v, want %v", err, tg.ErrRandomIdDuplicate)
	}
}

func TestMessagesSendMessage_ReceiverBackpressureMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			return nil, msg.ErrReceiverBackpressure
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesSendMessage_SenderSyncFailedMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			return nil, msg.ErrSenderSyncFailed
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesSendMessage_MsgStorageMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			return nil, msg.ErrMsgStorage
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesSendMessage_SendStateConflictMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			return nil, msg.ErrSendStateConflict
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesSendMessage_ContextDeadlineMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			return nil, context.DeadlineExceeded
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrTimeout {
		t.Fatalf("error = %v, want %v", err, tg.ErrTimeout)
	}
}

func TestMessagesSendMessage_UnknownErrorMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			return nil, errors.New("some unknown transport error")
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesSendMessage_TgErrorPassThrough(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			return nil, tg.ErrChatIdInvalid
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrChatIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrChatIdInvalid)
	}
}

// --- Metadata validation ---

func TestMessagesSendMessage_MissingMetadataRejected(t *testing.T) {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{MsgClient: &messagesFakeMsgClient{}},
	})

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrUserIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrUserIdInvalid)
	}
}

func TestMessagesSendMessage_UserIdZeroRejected(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{}, 0, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrUserIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrUserIdInvalid)
	}
}
