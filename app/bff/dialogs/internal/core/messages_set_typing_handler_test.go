package core

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/svc"
	syncclient "github.com/teamgram/teamgram-server/v2/app/messenger/sync/client"
	syncpb "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeSyncClient struct {
	syncclient.SyncClient
	err              error
	pushUpdatesCount int
	userIDs          []int64
	updates          []tg.UpdatesClazz
}

func (f *fakeSyncClient) SyncPushUpdates(ctx context.Context, in *syncpb.TLSyncPushUpdates) (*tg.Void, error) {
	f.pushUpdatesCount++
	f.userIDs = append(f.userIDs, in.UserId)
	f.updates = append(f.updates, in.Updates)
	if f.err != nil {
		return nil, f.err
	}
	return &tg.Void{}, nil
}

type fakeChatClient struct {
	chatclient.ChatClient
	participantIDs []int64
	err            error
}

func (f *fakeChatClient) ChatGetChatParticipantIdList(context.Context, *chatpb.TLChatGetChatParticipantIdList) (*chatpb.VectorLong, error) {
	if f.err != nil {
		return nil, f.err
	}
	return &chatpb.VectorLong{Datas: f.participantIDs}, nil
}

func TestMessagesSetTypingPushesUpdateUserTypingToPeer(t *testing.T) {
	syncClient := &fakeSyncClient{}
	c := newTestDialogsCore(syncClient, newTypingLimiter(5*time.Second), metadata.RpcMetadata{UserId: 1001})
	r, err := c.MessagesSetTyping(&tg.TLMessagesSetTyping{
		Peer:   tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 1002}),
		Action: tg.MakeTLSendMessageTypingAction(&tg.TLSendMessageTypingAction{}),
	})
	if err != nil {
		t.Fatalf("MessagesSetTyping() error = %v", err)
	}
	if !isBoolTrue(r) {
		t.Fatalf("reply = %v, want BoolTrue", r)
	}
	if syncClient.pushUpdatesCount != 1 {
		t.Fatalf("pushUpdatesCount = %d, want 1", syncClient.pushUpdatesCount)
	}
	if len(syncClient.userIDs) != 1 || syncClient.userIDs[0] != 1002 {
		t.Fatalf("pushed user ids = %v, want [1002]", syncClient.userIDs)
	}
	short, ok := syncClient.updates[0].(*tg.TLUpdateShort)
	if !ok {
		t.Fatalf("updates = %T, want updateShort", syncClient.updates[0])
	}
	typing, ok := short.Update.(*tg.TLUpdateUserTyping)
	if !ok {
		t.Fatalf("update = %T, want updateUserTyping", short.Update)
	}
	if typing.UserId != 1001 {
		t.Fatalf("typing.UserId = %d, want sender user 1001", typing.UserId)
	}
	if typing.Action == nil {
		t.Fatal("typing.Action is nil")
	}
}

func TestMessagesSetTypingPushesUpdateChatUserTypingToChatMembers(t *testing.T) {
	syncClient := &fakeSyncClient{}
	chatClient := &fakeChatClient{participantIDs: []int64{1001, 1002, 1003}}
	c := newTestDialogsCoreWithChat(syncClient, chatClient, newTypingLimiter(5*time.Second), metadata.RpcMetadata{UserId: 1001})

	r, err := c.MessagesSetTyping(&tg.TLMessagesSetTyping{
		Peer:   tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 6}),
		Action: tg.MakeTLSendMessageTypingAction(&tg.TLSendMessageTypingAction{}),
	})
	if err != nil {
		t.Fatalf("MessagesSetTyping() error = %v", err)
	}
	if !isBoolTrue(r) {
		t.Fatalf("reply = %v, want BoolTrue", r)
	}
	if got, want := syncClient.userIDs, []int64{1002, 1003}; len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Fatalf("pushed user ids = %v, want %v", got, want)
	}
	for i, updates := range syncClient.updates {
		short, ok := updates.(*tg.TLUpdateShort)
		if !ok {
			t.Fatalf("updates[%d] = %T, want updateShort", i, updates)
		}
		typing, ok := short.Update.(*tg.TLUpdateChatUserTyping)
		if !ok {
			t.Fatalf("update[%d] = %T, want updateChatUserTyping", i, short.Update)
		}
		if typing.ChatId != 6 {
			t.Fatalf("typing.ChatId = %d, want 6", typing.ChatId)
		}
		from, ok := typing.FromId.(*tg.TLPeerUser)
		if !ok || from.UserId != 1001 {
			t.Fatalf("typing.FromId = %#v, want peerUser 1001", typing.FromId)
		}
		if typing.Action == nil {
			t.Fatal("typing.Action is nil")
		}
	}
}

func TestMessagesSetTypingSwallowsSyncError(t *testing.T) {
	syncClient := &fakeSyncClient{err: errors.New("sync down")}
	c := newTestDialogsCore(syncClient, newTypingLimiter(5*time.Second), metadata.RpcMetadata{UserId: 1001})
	r, err := c.MessagesSetTyping(validTypingRequest(1002))
	if err != nil {
		t.Fatalf("MessagesSetTyping() error = %v, want nil", err)
	}
	if !isBoolTrue(r) {
		t.Fatalf("reply = %v, want BoolTrue", r)
	}
}

func TestMessagesSetTypingSwallowsMissingRealtimeDependencies(t *testing.T) {
	tests := []struct {
		name   string
		svcCtx *svc.ServiceContext
	}{
		{
			name:   "nil service context",
			svcCtx: nil,
		},
		{
			name: "nil repository",
			svcCtx: &svc.ServiceContext{
				Repo: nil,
			},
		},
		{
			name: "nil sync client",
			svcCtx: &svc.ServiceContext{
				Repo: &repository.Repository{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(context.Background(), tt.svcCtx)
			c.MD = &metadata.RpcMetadata{UserId: 1001}

			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("MessagesSetTyping() panicked: %v", r)
				}
			}()

			r, err := c.MessagesSetTyping(validTypingRequest(1002))
			if err != nil {
				t.Fatalf("MessagesSetTyping() error = %v, want nil", err)
			}
			if !isBoolTrue(r) {
				t.Fatalf("reply = %v, want BoolTrue", r)
			}
		})
	}
}

func TestMessagesSetTypingLimiterSkipsSyncCall(t *testing.T) {
	syncClient := &fakeSyncClient{}
	limiter := newTypingLimiter(time.Hour)
	c := newTestDialogsCore(syncClient, limiter, metadata.RpcMetadata{UserId: 1001})
	_, _ = c.MessagesSetTyping(validTypingRequest(1002))
	_, err := c.MessagesSetTyping(validTypingRequest(1002))
	if err != nil {
		t.Fatalf("MessagesSetTyping() error = %v", err)
	}
	if syncClient.pushUpdatesCount != 1 {
		t.Fatalf("pushUpdatesCount = %d, want only first call", syncClient.pushUpdatesCount)
	}
}

func TestMessagesSetTypingRejectsChannelPeer(t *testing.T) {
	c := newTestDialogsCore(&fakeSyncClient{}, newTypingLimiter(5*time.Second), metadata.RpcMetadata{UserId: 1001})
	_, err := c.MessagesSetTyping(&tg.TLMessagesSetTyping{
		Peer:   tg.MakeTLInputPeerChannel(&tg.TLInputPeerChannel{ChannelId: 2001}),
		Action: tg.MakeTLSendMessageTypingAction(&tg.TLSendMessageTypingAction{}),
	})
	if !errors.Is(err, tg.Err400PeerIdInvalid) && !errors.Is(err, tg.ErrPeerIdNotSupported) {
		t.Fatalf("error = %v, want peer invalid/not supported", err)
	}
}

func validTypingRequest(peerUserID int64) *tg.TLMessagesSetTyping {
	return &tg.TLMessagesSetTyping{
		Peer:   tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: peerUserID}),
		Action: tg.MakeTLSendMessageTypingAction(&tg.TLSendMessageTypingAction{}),
	}
}

func isBoolTrue(v *tg.Bool) bool {
	_, ok := v.ToBoolTrue()
	return ok
}

func newTestDialogsCore(syncClient syncclient.SyncClient, limiter svc.TypingLimiter, md metadata.RpcMetadata) *DialogsCore {
	return newTestDialogsCoreWithChat(syncClient, nil, limiter, md)
}

func newTestDialogsCoreWithChat(syncClient syncclient.SyncClient, chatClient chatclient.ChatClient, limiter svc.TypingLimiter, md metadata.RpcMetadata) *DialogsCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{
			SyncClient: syncClient,
			ChatClient: chatClient,
		},
		TypingLimiter: limiter,
	})
	c.MD = &md
	return c
}
