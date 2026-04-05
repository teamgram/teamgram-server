package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/updates/internal/svc"
	bizupdates "github.com/teamgram/teamgram-server/v2/app/service/biz/updates/updates"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeUpdatesStateClient struct {
	stateReq *bizupdates.TLUpdatesGetStateV2
	diffReq  *bizupdates.TLUpdatesGetDifferenceV2
	state    *tg.UpdatesState
	diff     *bizupdates.Difference
}

func (f *fakeUpdatesStateClient) UpdatesGetStateV2(ctx context.Context, in *bizupdates.TLUpdatesGetStateV2) (*tg.UpdatesState, error) {
	f.stateReq = in
	return f.state, nil
}

func (f *fakeUpdatesStateClient) UpdatesGetDifferenceV2(ctx context.Context, in *bizupdates.TLUpdatesGetDifferenceV2) (*bizupdates.Difference, error) {
	f.diffReq = in
	return f.diff, nil
}

var _ svc.UpdatesStateClient = (*fakeUpdatesStateClient)(nil)

func TestUpdatesGetStateReturnsPlaceholderState(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetState(&tg.TLUpdatesGetState{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.Pts != 1 || result.Date != 10 {
		t.Fatalf("expected placeholder state pts=1 date=10, got %#v", result)
	}
}

func TestUpdatesGetStateDelegatesToUpdatesService(t *testing.T) {
	fakeCli := &fakeUpdatesStateClient{
		state: tg.MakeTLUpdatesState(&tg.TLUpdatesState{
			Pts:  9,
			Date: 99,
		}).ToUpdatesState(),
	}
	c := New(context.Background(), &svc.ServiceContext{
		UpdatesClient: fakeCli,
	})
	c.MD = &metadata.RpcMetadata{
		AuthId: 1001,
		UserId: 42,
	}

	result, err := c.UpdatesGetState(&tg.TLUpdatesGetState{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if fakeCli.stateReq == nil {
		t.Fatal("expected UpdatesGetStateV2 to be called")
	}
	if fakeCli.stateReq.AuthKeyId != 1001 || fakeCli.stateReq.UserId != 42 {
		t.Fatalf("expected propagated metadata, got %#v", fakeCli.stateReq)
	}
	if result.Pts != 9 || result.Date != 99 {
		t.Fatalf("expected delegated state pts=9 date=99, got %#v", result)
	}
}

func TestUpdatesGetDifferenceReturnsCatchUpPayloadForBehindClient(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetDifference(&tg.TLUpdatesGetDifference{
		Pts:  0,
		Date: 0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	diff, ok := result.ToUpdatesDifference()
	if !ok {
		t.Fatalf("expected updates.difference, got %T", result.Clazz)
	}
	if diff.State == nil || diff.State.Pts != 1 || diff.State.Date != 10 {
		t.Fatalf("expected placeholder state pts=1 date=10, got %#v", diff.State)
	}
	if len(diff.NewMessages) != 1 || len(diff.OtherUpdates) != 1 {
		t.Fatalf("expected single catch-up payload, got messages=%d updates=%d", len(diff.NewMessages), len(diff.OtherUpdates))
	}
}

func TestUpdatesGetDifferenceReturnsDifferenceEmptyForForwardClient(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetDifference(&tg.TLUpdatesGetDifference{
		Pts:  7,
		Date: 99,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	diff, ok := result.ToUpdatesDifferenceEmpty()
	if !ok {
		t.Fatalf("expected updates.differenceEmpty, got %T", result.Clazz)
	}
	if diff.Date != 99 || diff.Seq != 0 {
		t.Fatalf("expected placeholder differenceEmpty date=99 seq=0, got %#v", diff)
	}
}

func TestUpdatesGetDifferenceDelegatesToUpdatesService(t *testing.T) {
	fakeCli := &fakeUpdatesStateClient{
		diff: bizupdates.MakeTLDifference(&bizupdates.TLDifference{
			NewMessages: []tg.MessageClazz{
				tg.MakeTLMessage(&tg.TLMessage{
					Id:      5,
					Date:    55,
					Message: "from-updates-service",
				}),
			},
			OtherUpdates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: tg.MakeTLMessage(&tg.TLMessage{
						Id:      5,
						Date:    55,
						Message: "from-updates-service",
					}),
					Pts:      5,
					PtsCount: 1,
				}),
			},
			State: tg.MakeTLUpdatesState(&tg.TLUpdatesState{
				Pts:  5,
				Date: 55,
				Seq:  0,
			}).ToUpdatesState(),
		}).ToDifference(),
	}
	c := New(context.Background(), &svc.ServiceContext{
		UpdatesClient: fakeCli,
	})
	c.MD = &metadata.RpcMetadata{
		AuthId: 2002,
		UserId: 77,
	}

	result, err := c.UpdatesGetDifference(&tg.TLUpdatesGetDifference{
		Pts:  0,
		Date: 0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if fakeCli.diffReq == nil {
		t.Fatal("expected UpdatesGetDifferenceV2 to be called")
	}
	if fakeCli.diffReq.AuthKeyId != 2002 || fakeCli.diffReq.UserId != 77 {
		t.Fatalf("expected propagated metadata, got %#v", fakeCli.diffReq)
	}
	diff, ok := result.ToUpdatesDifference()
	if !ok {
		t.Fatalf("expected updates.difference, got %T", result.Clazz)
	}
	msg, ok := diff.NewMessages[0].(*tg.TLMessage)
	if !ok || msg.Message != "from-updates-service" {
		t.Fatalf("expected delegated difference payload, got %#v", diff.NewMessages[0])
	}
}

func TestUpdatesGetChannelDifferenceReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetChannelDifference(&tg.TLUpdatesGetChannelDifference{
		Pts: 4,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	diff, ok := result.ToUpdatesChannelDifferenceEmpty()
	if !ok {
		t.Fatalf("expected updates.channelDifferenceEmpty, got %T", result.Clazz)
	}
	if !diff.Final || diff.Pts != 4 {
		t.Fatalf("expected final placeholder pts=4, got %#v", diff)
	}
}
