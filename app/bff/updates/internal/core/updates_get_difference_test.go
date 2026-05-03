package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/updates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/updates/internal/svc"
	userupdatesclient "github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeUserupdatesClient struct {
	userupdatesclient.UserupdatesClient
	gotState      *userupdates.TLUserupdatesGetState
	state         *userupdates.UserState
	gotDifference *userupdates.TLUserupdatesGetDifference
	difference    *userupdates.UserDifference
}

func (f *fakeUserupdatesClient) UserupdatesGetState(ctx context.Context, in *userupdates.TLUserupdatesGetState) (*userupdates.UserState, error) {
	f.gotState = in
	return f.state, nil
}

func (f *fakeUserupdatesClient) UserupdatesGetDifference(ctx context.Context, in *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error) {
	f.gotDifference = in
	return f.difference, nil
}

func newUpdatesCore(client userupdatesclient.UserupdatesClient) *UpdatesCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{UserupdatesClient: client},
	})
	c.MD = &metadata.RpcMetadata{UserId: 1001, PermAuthKeyId: 2002}
	return c
}

func TestUpdatesGetStateReturnsUserupdatesState(t *testing.T) {
	client := &fakeUserupdatesClient{state: userupdates.MakeTLUserState(&userupdates.TLUserState{
		Pts:         88,
		Qts:         -1,
		Date:        123,
		Seq:         2,
		UnreadCount: 3,
	}).ToUserState()}
	core := newUpdatesCore(client)

	got, err := core.UpdatesGetState(&tg.TLUpdatesGetState{})
	if err != nil {
		t.Fatalf("UpdatesGetState() error = %v", err)
	}
	if client.gotState == nil || client.gotState.UserId != 1001 || client.gotState.AuthKeyId != 2002 {
		t.Fatalf("userupdates request = %#v", client.gotState)
	}
	if got.Pts != 88 || got.Qts != -1 || got.Date != 123 || got.Seq != 2 || got.UnreadCount != 3 {
		t.Fatalf("updates state = %#v", got)
	}
}

func TestUpdatesGetDifferenceReturnsNonEmptyDifference(t *testing.T) {
	client := &fakeUserupdatesClient{difference: userupdates.MakeTLUserDifference(&userupdates.TLUserDifference{
		NewMessages: []tg.MessageClazz{
			tg.MakeTLMessage(&tg.TLMessage{Id: 9, Message: "hello"}),
		},
		OtherUpdates: []tg.UpdateClazz{
			tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{Pts: 18, PtsCount: 1}),
		},
		State: userupdates.MakeTLUserState(&userupdates.TLUserState{Pts: 18, Qts: -1, Date: 123, Seq: 0, UnreadCount: 0}),
	}).ToUserDifference()}
	core := newUpdatesCore(client)

	got, err := core.UpdatesGetDifference(&tg.TLUpdatesGetDifference{Pts: 17, PtsTotalLimit: int32Ptr(50), Date: 1, Qts: -1})
	if err != nil {
		t.Fatalf("UpdatesGetDifference() error = %v", err)
	}
	if client.gotDifference == nil || client.gotDifference.UserId != 1001 || client.gotDifference.AuthKeyId != 2002 || client.gotDifference.Pts != 17 {
		t.Fatalf("userupdates request = %#v", client.gotDifference)
	}
	diff, ok := got.ToUpdatesDifference()
	if !ok {
		t.Fatalf("got %s, want updates.difference", got.ClazzName())
	}
	if len(diff.NewMessages) != 1 || len(diff.OtherUpdates) != 1 || diff.State == nil || diff.State.Pts != 18 || diff.State.Qts != -1 {
		t.Fatalf("difference = %#v", diff)
	}
}

func TestUpdatesGetDifferenceReturnsEmptyDifference(t *testing.T) {
	client := &fakeUserupdatesClient{difference: userupdates.MakeTLUserDifferenceEmpty(&userupdates.TLUserDifferenceEmpty{
		State: userupdates.MakeTLUserState(&userupdates.TLUserState{Pts: 17, Qts: -1, Date: 123, Seq: 0}),
	}).ToUserDifference()}
	core := newUpdatesCore(client)

	got, err := core.UpdatesGetDifference(&tg.TLUpdatesGetDifference{Pts: 17, Date: 1, Qts: -1})
	if err != nil {
		t.Fatalf("UpdatesGetDifference() error = %v", err)
	}
	empty, ok := got.ToUpdatesDifferenceEmpty()
	if !ok {
		t.Fatalf("got %s, want updates.differenceEmpty", got.ClazzName())
	}
	if empty.Date != 123 || empty.Seq != 0 {
		t.Fatalf("empty difference = %#v", empty)
	}
}

func int32Ptr(v int32) *int32 {
	return &v
}
