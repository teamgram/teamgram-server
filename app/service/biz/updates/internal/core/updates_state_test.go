package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/updates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeUpdatesMessageClient struct {
	req    *message.TLMessageGetHistoryMessages
	result *message.VectorMessageBox
}

func (f *fakeUpdatesMessageClient) MessageGetHistoryMessages(ctx context.Context, in *message.TLMessageGetHistoryMessages) (*message.VectorMessageBox, error) {
	f.req = in
	return f.result, nil
}

var _ svc.MessageQueryClient = (*fakeUpdatesMessageClient)(nil)

func TestUpdatesGetStateV2ReturnsPlaceholderState(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetStateV2(&updates.TLUpdatesGetStateV2{
		AuthKeyId: 1,
		UserId:    2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected updates state, got nil")
	}
	if result.Pts != 1 {
		t.Fatalf("expected placeholder pts=1, got %d", result.Pts)
	}
	if result.Date != 10 {
		t.Fatalf("expected placeholder date=10, got %d", result.Date)
	}
}

func TestUpdatesGetDifferenceV2ReturnsCatchUpDifferenceForBehindClient(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetDifferenceV2(&updates.TLUpdatesGetDifferenceV2{
		AuthKeyId: 1,
		UserId:    2,
		Pts:       0,
		Date:      0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected difference, got nil")
	}
	diff, ok := result.ToDifference()
	if !ok {
		t.Fatalf("expected difference, got %T", result.Clazz)
	}
	if diff.State == nil {
		t.Fatal("expected placeholder state, got nil")
	}
	if len(diff.NewMessages) != 1 || len(diff.OtherUpdates) != 1 {
		t.Fatalf("expected single catch-up payload, got messages=%d updates=%d", len(diff.NewMessages), len(diff.OtherUpdates))
	}
	if diff.State.Pts != 1 {
		t.Fatalf("expected placeholder pts=1 for zero input, got %d", diff.State.Pts)
	}
	if diff.State.Date != 10 {
		t.Fatalf("expected placeholder date=10 for zero input, got %d", diff.State.Date)
	}
}

func TestUpdatesGetDifferenceV2PreservesForwardProgressFromRequest(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetDifferenceV2(&updates.TLUpdatesGetDifferenceV2{
		AuthKeyId: 1,
		UserId:    2,
		Pts:       7,
		Date:      99,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	diff, ok := result.ToDifferenceEmpty()
	if !ok {
		t.Fatalf("expected differenceEmpty, got %T", result.Clazz)
	}
	if diff.State == nil {
		t.Fatal("expected placeholder state, got nil")
	}
	if diff.State.Pts != 7 {
		t.Fatalf("expected pts=7, got %d", diff.State.Pts)
	}
	if diff.State.Date != 99 {
		t.Fatalf("expected date=99, got %d", diff.State.Date)
	}
}

func TestUpdatesGetDifferenceV2DelegatesCatchUpMessageToMessageService(t *testing.T) {
	fakeMsgCli := &fakeUpdatesMessageClient{
		result: &message.VectorMessageBox{
			Datas: []tg.MessageBoxClazz{
				tg.MakeTLMessageBox(&tg.TLMessageBox{
					UserId:    2,
					MessageId: 1,
					PeerType:  tg.PEER_USER,
					PeerId:    2,
					Message: tg.MakeTLMessage(&tg.TLMessage{
						Id:      1,
						Date:    10,
						Message: "from-message-service",
					}),
				}),
			},
		},
	}
	c := New(context.Background(), &svc.ServiceContext{
		MessageClient: fakeMsgCli,
	})

	result, err := c.UpdatesGetDifferenceV2(&updates.TLUpdatesGetDifferenceV2{
		AuthKeyId: 1,
		UserId:    2,
		Pts:       0,
		Date:      0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if fakeMsgCli.req == nil {
		t.Fatal("expected MessageGetHistoryMessages to be called")
	}
	diff, ok := result.ToDifference()
	if !ok {
		t.Fatalf("expected difference, got %T", result.Clazz)
	}
	msg, ok := diff.NewMessages[0].(*tg.TLMessage)
	if !ok || msg.Message != "from-message-service" {
		t.Fatalf("expected delegated catch-up message, got %#v", diff.NewMessages[0])
	}
}

func TestUpdatesGetChannelDifferenceV2ReturnsPlaceholderChannelDifference(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetChannelDifferenceV2(&updates.TLUpdatesGetChannelDifferenceV2{
		AuthKeyId: 1,
		UserId:    2,
		ChannelId: 3,
		Pts:       4,
		Limit:     100,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected channel difference, got nil")
	}
	channelDiff := result.ToChannelDifference()
	if channelDiff.Pts != 4 {
		t.Fatalf("expected pts=4, got %d", channelDiff.Pts)
	}
	if len(channelDiff.NewMessages) != 0 || len(channelDiff.OtherUpdates) != 0 {
		t.Fatalf("expected empty channel difference payload, got new_messages=%d other_updates=%d",
			len(channelDiff.NewMessages), len(channelDiff.OtherUpdates))
	}
}

func TestUpdatesGetStateV2ReturnsRealStateFromRepository(t *testing.T) {
	repo := repository.NewRepository(config.Config{})
	// Seed a known state for user 2.
	repo.UpdatesState.SetUserUpdatesState(context.Background(), &repository.UpdatesStateDO{
		UserId: 2, Pts: 42, Qts: 1, Date: 1000, Seq: 5, UnreadCount: 3,
	})

	c := New(context.Background(), &svc.ServiceContext{
		Repository: repo,
	})

	result, err := c.UpdatesGetStateV2(&updates.TLUpdatesGetStateV2{
		AuthKeyId: 1,
		UserId:    2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result.Pts != 42 {
		t.Fatalf("expected pts=42, got %d", result.Pts)
	}
	if result.Date != 1000 {
		t.Fatalf("expected date=1000, got %d", result.Date)
	}
	if result.Seq != 5 {
		t.Fatalf("expected seq=5, got %d", result.Seq)
	}
	if result.UnreadCount != 3 {
		t.Fatalf("expected unreadCount=3, got %d", result.UnreadCount)
	}
}

func TestUpdatesGetDifferenceV2ReturnsCatchUpWhenClientBehind(t *testing.T) {
	repo := repository.NewRepository(config.Config{})
	// Seed server state with pts=5.
	repo.UpdatesState.SetUserUpdatesState(context.Background(), &repository.UpdatesStateDO{
		UserId: 2, Pts: 5, Qts: 0, Date: 1000, Seq: 0, UnreadCount: 0,
	})

	c := New(context.Background(), &svc.ServiceContext{
		Repository: repo,
	})

	result, err := c.UpdatesGetDifferenceV2(&updates.TLUpdatesGetDifferenceV2{
		AuthKeyId: 1,
		UserId:    2,
		Pts:       2,
		Date:      500,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	diff, ok := result.ToDifference()
	if !ok {
		t.Fatalf("expected difference, got %T", result.Clazz)
	}
	if diff.State.Pts != 5 {
		t.Fatalf("expected state pts=5 (server), got %d", diff.State.Pts)
	}
}

func TestUpdatesGetDifferenceV2ReturnsEmptyWhenClientCurrent(t *testing.T) {
	repo := repository.NewRepository(config.Config{})
	// Seed server state with pts=5.
	repo.UpdatesState.SetUserUpdatesState(context.Background(), &repository.UpdatesStateDO{
		UserId: 2, Pts: 5, Qts: 0, Date: 1000, Seq: 0, UnreadCount: 0,
	})

	c := New(context.Background(), &svc.ServiceContext{
		Repository: repo,
	})

	result, err := c.UpdatesGetDifferenceV2(&updates.TLUpdatesGetDifferenceV2{
		AuthKeyId: 1,
		UserId:    2,
		Pts:       5,
		Date:      1000,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	_, ok := result.ToDifferenceEmpty()
	if !ok {
		t.Fatalf("expected differenceEmpty, got %T", result.Clazz)
	}
}

func TestUpdatesGetDifferenceV2FreshClientGetsCatchUp(t *testing.T) {
	repo := repository.NewRepository(config.Config{})
	// No pre-existing state for user 2 (fresh user).
	c := New(context.Background(), &svc.ServiceContext{
		Repository: repo,
	})

	result, err := c.UpdatesGetDifferenceV2(&updates.TLUpdatesGetDifferenceV2{
		AuthKeyId: 1,
		UserId:    2,
		Pts:       0,
		Date:      0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	diff, ok := result.ToDifference()
	if !ok {
		t.Fatalf("expected difference for fresh client, got %T", result.Clazz)
	}
	// State should use normalized defaults.
	if diff.State.Pts != 1 {
		t.Fatalf("expected state pts=1 for fresh client, got %d", diff.State.Pts)
	}
}
