package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/updates/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/updates"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// --- fake updates client ---

type fakeUpdatesClient struct {
	stateReq *updates.TLUpdatesGetStateV2
	stateRes *tg.UpdatesState

	diffReq *updates.TLUpdatesGetDifferenceV2
	diffRes *updates.Difference
}

func (f *fakeUpdatesClient) UpdatesGetStateV2(ctx context.Context, in *updates.TLUpdatesGetStateV2) (*tg.UpdatesState, error) {
	f.stateReq = in
	return f.stateRes, nil
}

func (f *fakeUpdatesClient) UpdatesGetDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetDifferenceV2) (*updates.Difference, error) {
	f.diffReq = in
	return f.diffRes, nil
}

var _ svc.UpdatesStateClient = (*fakeUpdatesClient)(nil)

func newUpdatesWithMD(ctx context.Context, svcCtx *svc.ServiceContext, authKeyId, userId int64) *UpdatesCore {
	c := New(ctx, svcCtx)
	c.MD = &metadata.RpcMetadata{AuthId: authKeyId, UserId: userId}
	return c
}

func TestUpdatesGetStateDelegatesToUpdatesClient(t *testing.T) {
	fake := &fakeUpdatesClient{
		stateRes: tg.MakeTLUpdatesState(&tg.TLUpdatesState{
			Pts:         42,
			Qts:         1,
			Date:        1000,
			Seq:         5,
			UnreadCount: 3,
		}),
	}

	c := newUpdatesWithMD(context.Background(), &svc.ServiceContext{
		UpdatesClient: fake,
	}, 111, 222)

	result, err := c.UpdatesGetState(&tg.TLUpdatesGetState{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if fake.stateReq == nil {
		t.Fatal("expected UpdatesGetStateV2 to be called")
	}
	if fake.stateReq.AuthKeyId != 111 {
		t.Errorf("expected authKeyId=111, got %d", fake.stateReq.AuthKeyId)
	}
	if fake.stateReq.UserId != 222 {
		t.Errorf("expected userId=222, got %d", fake.stateReq.UserId)
	}

	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Pts != 42 || result.Date != 1000 {
		t.Errorf("expected pts=42 date=1000, got pts=%d date=%d", result.Pts, result.Date)
	}
}

func TestUpdatesGetDifferenceDelegatesToUpdatesClient(t *testing.T) {
	state := tg.MakeTLUpdatesState(&tg.TLUpdatesState{
		Pts:  50,
		Qts:  1,
		Date: 2000,
		Seq:  10,
	})

	fake := &fakeUpdatesClient{
		diffRes: updates.MakeTLDifference(&updates.TLDifference{
			NewMessages: []tg.MessageClazz{
				tg.MakeTLMessage(&tg.TLMessage{
					Id:      100,
					Date:    2000,
					Message: "hello",
					PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1}),
				}),
			},
			OtherUpdates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Pts:      50,
					PtsCount: 1,
				}),
			},
			State: state,
		}).ToDifference(),
	}

	c := newUpdatesWithMD(context.Background(), &svc.ServiceContext{
		UpdatesClient: fake,
	}, 111, 222)

	result, err := c.UpdatesGetDifference(&tg.TLUpdatesGetDifference{
		Pts:  40,
		Date: 1000,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if fake.diffReq == nil {
		t.Fatal("expected UpdatesGetDifferenceV2 to be called")
	}
	if fake.diffReq.AuthKeyId != 111 || fake.diffReq.UserId != 222 {
		t.Errorf("expected authKeyId=111 userId=222, got authKeyId=%d userId=%d", fake.diffReq.AuthKeyId, fake.diffReq.UserId)
	}
	if fake.diffReq.Pts != 40 || fake.diffReq.Date != 1000 {
		t.Errorf("expected pts=40 date=1000, got pts=%d date=%d", fake.diffReq.Pts, fake.diffReq.Date)
	}

	diff, ok := result.ToUpdatesDifference()
	if !ok {
		t.Fatalf("expected updates.difference, got %T", result.Clazz)
	}
	if len(diff.NewMessages) != 1 {
		t.Fatalf("expected 1 new message, got %d", len(diff.NewMessages))
	}
	if len(diff.OtherUpdates) != 1 {
		t.Fatalf("expected 1 other update, got %d", len(diff.OtherUpdates))
	}
}

func TestUpdatesGetDifferenceReturnsEmptyForNilDiff(t *testing.T) {
	fake := &fakeUpdatesClient{
		diffRes: nil,
	}

	c := newUpdatesWithMD(context.Background(), &svc.ServiceContext{
		UpdatesClient: fake,
	}, 111, 222)

	result, err := c.UpdatesGetDifference(&tg.TLUpdatesGetDifference{
		Pts:  40,
		Date: 1000,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	_, ok := result.ToUpdatesDifferenceEmpty()
	if !ok {
		t.Fatalf("expected updates.differenceEmpty for nil diff, got %T", result.Clazz)
	}
}

func TestUpdatesGetDifferenceMapsEmptyDifference(t *testing.T) {
	state := tg.MakeTLUpdatesState(&tg.TLUpdatesState{
		Pts:  50,
		Qts:  1,
		Date: 2000,
		Seq:  10,
	})

	fake := &fakeUpdatesClient{
		diffRes: updates.MakeTLDifferenceEmpty(&updates.TLDifferenceEmpty{
			State: state,
		}).ToDifference(),
	}

	c := newUpdatesWithMD(context.Background(), &svc.ServiceContext{
		UpdatesClient: fake,
	}, 111, 222)

	result, err := c.UpdatesGetDifference(&tg.TLUpdatesGetDifference{
		Pts:  40,
		Date: 1000,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	diffEmpty, ok := result.ToUpdatesDifferenceEmpty()
	if !ok {
		t.Fatalf("expected updates.differenceEmpty, got %T", result.Clazz)
	}
	if diffEmpty.Date != 2000 || diffEmpty.Seq != 10 {
		t.Errorf("expected date=2000 seq=10, got date=%d seq=%d", diffEmpty.Date, diffEmpty.Seq)
	}
}
