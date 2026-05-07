package core

import (
	"context"
	"math"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/updates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/updates/internal/svc"
	userupdatesclient "github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
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

type fakeDifferenceUserClient struct {
	userclient.UserClient
	in    *userpb.TLUserGetUserProjectionBundle
	out   *userpb.UserProjectionBundle
	calls int
}

func (f *fakeDifferenceUserClient) UserGetUserProjectionBundle(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
	f.calls++
	f.in = in
	return f.out, nil
}

func newUpdatesCore(client userupdatesclient.UserupdatesClient) *UpdatesCore {
	return newUpdatesCoreWithUser(client, nil)
}

func newUpdatesCoreWithUser(client userupdatesclient.UserupdatesClient, userClient userclient.UserClient) *UpdatesCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{UserupdatesClient: client, UserClient: userClient},
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

func TestUpdatesGetStateAcceptsNegativePermAuthKeyID(t *testing.T) {
	const permAuthKeyID = int64(-4149588253508792542)
	client := &fakeUserupdatesClient{state: userupdates.MakeTLUserState(&userupdates.TLUserState{
		Pts: 1,
		Qts: -1,
	}).ToUserState()}
	core := newUpdatesCore(client)
	core.MD.PermAuthKeyId = permAuthKeyID

	_, err := core.UpdatesGetState(&tg.TLUpdatesGetState{})
	if err != nil {
		t.Fatalf("UpdatesGetState() error = %v", err)
	}
	if client.gotState == nil || client.gotState.AuthKeyId != permAuthKeyID {
		t.Fatalf("userupdates auth key id = %#v, want %d", client.gotState, permAuthKeyID)
	}
}

func TestUpdatesGetDifferenceReturnsNonEmptyDifference(t *testing.T) {
	client := &fakeUserupdatesClient{difference: userupdates.MakeTLUserDifference(&userupdates.TLUserDifference{
		NewMessages: []tg.MessageClazz{
			tg.MakeTLMessage(&tg.TLMessage{Id: 9, Message: "hello"}),
		},
		OtherUpdates: []tg.UpdateClazz{
			tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
				Message:  tg.MakeTLMessage(&tg.TLMessage{Id: 10, Message: "from update"}),
				Pts:      18,
				PtsCount: 1,
			}),
			tg.MakeTLUpdateReadHistoryInbox(&tg.TLUpdateReadHistoryInbox{
				Peer:             tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1002}),
				MaxId:            10,
				StillUnreadCount: 0,
				Pts:              18,
				PtsCount:         1,
			}),
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
	if len(diff.NewMessages) != 2 || len(diff.OtherUpdates) != 1 || diff.State == nil || diff.State.Pts != 18 || diff.State.Qts != -1 {
		t.Fatalf("difference = %#v", diff)
	}
	if diff.NewMessages[1].(*tg.TLMessage).Id != 10 {
		t.Fatalf("merged updateNewMessage message = %#v", diff.NewMessages[1])
	}
	if _, ok := diff.OtherUpdates[0].(*tg.TLUpdateReadHistoryInbox); !ok {
		t.Fatalf("other update = %T, want TLUpdateReadHistoryInbox", diff.OtherUpdates[0])
	}
}

func TestUpdatesGetDifferenceProjectsUsers(t *testing.T) {
	client := &fakeUserupdatesClient{difference: userupdates.MakeTLUserDifference(&userupdates.TLUserDifference{
		NewMessages: []tg.MessageClazz{
			tg.MakeTLMessage(&tg.TLMessage{
				Id:      9,
				FromId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1001}),
				PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1002}),
				Message: "hello",
			}),
		},
		State: userupdates.MakeTLUserState(&userupdates.TLUserState{Pts: 18, Qts: -1, Date: 123, Seq: 0}),
	}).ToUserDifference()}
	userClient := &fakeDifferenceUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1001, Self: true}),
				tg.MakeTLUser(&tg.TLUser{Id: 1002, Contact: true}),
			}}),
		},
	}).ToUserProjectionBundle()}
	core := newUpdatesCoreWithUser(client, userClient)

	got, err := core.UpdatesGetDifference(&tg.TLUpdatesGetDifference{Pts: 17, Date: 1, Qts: -1})
	if err != nil {
		t.Fatalf("UpdatesGetDifference() error = %v", err)
	}
	diff, ok := got.ToUpdatesDifference()
	if !ok {
		t.Fatalf("got %s, want updates.difference", got.ClazzName())
	}
	if userClient.in == nil || len(userClient.in.ViewerUserIds) != 1 || userClient.in.ViewerUserIds[0] != 1001 {
		t.Fatalf("projection viewer request = %#v", userClient.in)
	}
	if len(userClient.in.TargetUserIds) != 2 || userClient.in.TargetUserIds[0] != 1001 || userClient.in.TargetUserIds[1] != 1002 {
		t.Fatalf("projection target request = %#v", userClient.in)
	}
	if len(diff.Users) != 2 {
		t.Fatalf("users = %#v", diff.Users)
	}
	self, ok := diff.Users[0].(*tg.TLUser)
	if !ok || self.Id != 1001 || !self.Self {
		t.Fatalf("self user = %#v", diff.Users[0])
	}
	contact, ok := diff.Users[1].(*tg.TLUser)
	if !ok || contact.Id != 1002 || !contact.Contact {
		t.Fatalf("contact user = %#v", diff.Users[1])
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

func TestUpdatesGetDifferenceEmptyDoesNotProjectUsers(t *testing.T) {
	client := &fakeUserupdatesClient{difference: userupdates.MakeTLUserDifferenceEmpty(&userupdates.TLUserDifferenceEmpty{
		State: userupdates.MakeTLUserState(&userupdates.TLUserState{Pts: 17, Qts: -1, Date: 123, Seq: 0}),
	}).ToUserDifference()}
	userClient := &fakeDifferenceUserClient{}
	core := newUpdatesCoreWithUser(client, userClient)

	got, err := core.UpdatesGetDifference(&tg.TLUpdatesGetDifference{Pts: 17, Date: 1, Qts: -1})
	if err != nil {
		t.Fatalf("UpdatesGetDifference() error = %v", err)
	}
	if _, ok := got.ToUpdatesDifferenceEmpty(); !ok {
		t.Fatalf("got %s, want updates.differenceEmpty", got.ClazzName())
	}
	if userClient.calls != 0 {
		t.Fatalf("user projection calls = %d, want 0", userClient.calls)
	}
}

func TestUpdatesGetDifferenceReturnsSlice(t *testing.T) {
	client := &fakeUserupdatesClient{difference: userupdates.MakeTLUserDifferenceSlice(&userupdates.TLUserDifferenceSlice{
		NewMessages: []tg.MessageClazz{
			tg.MakeTLMessage(&tg.TLMessage{Id: 9, Message: "hello"}),
		},
		OtherUpdates: []tg.UpdateClazz{
			tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
				Message:  tg.MakeTLMessage(&tg.TLMessage{Id: 10, Message: "from slice update"}),
				Pts:      18,
				PtsCount: 1,
			}),
		},
		IntermediateState: userupdates.MakeTLUserState(&userupdates.TLUserState{
			Pts:         18,
			Qts:         0,
			Date:        123,
			Seq:         2,
			UnreadCount: 0,
		}),
	}).ToUserDifference()}
	core := newUpdatesCore(client)

	got, err := core.UpdatesGetDifference(&tg.TLUpdatesGetDifference{Pts: 17, Date: 100, Qts: 0})
	if err != nil {
		t.Fatalf("UpdatesGetDifference() error = %v", err)
	}
	if client.gotDifference == nil || client.gotDifference.Date == nil || *client.gotDifference.Date != 100 {
		t.Fatalf("userupdates date = %#v, want 100", client.gotDifference)
	}
	slice, ok := got.ToUpdatesDifferenceSlice()
	if !ok {
		t.Fatalf("got %s, want updates.differenceSlice", got.ClazzName())
	}
	if len(slice.NewMessages) != 2 || len(slice.OtherUpdates) != 0 {
		t.Fatalf("slice payload = %#v", slice)
	}
	if slice.NewMessages[1].(*tg.TLMessage).Id != 10 {
		t.Fatalf("merged slice updateNewMessage message = %#v", slice.NewMessages[1])
	}
	if slice.IntermediateState == nil || slice.IntermediateState.Pts != 18 || slice.IntermediateState.Date != 123 || slice.IntermediateState.Seq != 2 {
		t.Fatalf("intermediate state = %#v", slice.IntermediateState)
	}
	if slice.NewEncryptedMessages == nil || slice.Chats == nil || slice.Users == nil {
		t.Fatalf("public vectors must be initialized: %#v", slice)
	}
}

func TestUpdatesGetDifferenceReturnsTooLong(t *testing.T) {
	client := &fakeUserupdatesClient{difference: userupdates.MakeTLUserDifferenceTooLong(&userupdates.TLUserDifferenceTooLong{
		Pts: 123,
	}).ToUserDifference()}
	core := newUpdatesCore(client)

	got, err := core.UpdatesGetDifference(&tg.TLUpdatesGetDifference{Pts: 1, Date: 100, Qts: 0})
	if err != nil {
		t.Fatalf("UpdatesGetDifference() error = %v", err)
	}
	tooLong, ok := got.ToUpdatesDifferenceTooLong()
	if !ok {
		t.Fatalf("got %s, want updates.differenceTooLong", got.ClazzName())
	}
	if tooLong.Pts != 123 {
		t.Fatalf("tooLong pts = %d, want 123", tooLong.Pts)
	}
}

func TestUpdatesGetStateRejectsOutOfRangePublicPts(t *testing.T) {
	client := &fakeUserupdatesClient{state: userupdates.MakeTLUserState(&userupdates.TLUserState{
		Pts: int64(math.MaxInt32) + 1,
	}).ToUserState()}
	core := newUpdatesCore(client)

	_, err := core.UpdatesGetState(&tg.TLUpdatesGetState{})
	if err == nil || !strings.Contains(err.Error(), "updates.state.pts out of int32 range") {
		t.Fatalf("UpdatesGetState() error = %v, want checked pts overflow", err)
	}
}

func TestUpdatesGetDifferenceRejectsOutOfRangeTooLongPts(t *testing.T) {
	client := &fakeUserupdatesClient{difference: userupdates.MakeTLUserDifferenceTooLong(&userupdates.TLUserDifferenceTooLong{
		Pts: int64(math.MaxInt32) + 1,
	}).ToUserDifference()}
	core := newUpdatesCore(client)

	_, err := core.UpdatesGetDifference(&tg.TLUpdatesGetDifference{Pts: 1, Date: 100, Qts: 0})
	if err == nil || !strings.Contains(err.Error(), "updates.differenceTooLong.pts out of int32 range") {
		t.Fatalf("UpdatesGetDifference() error = %v, want checked tooLong pts overflow", err)
	}
}

func TestUpdatesGetStateRejectsMissingPermAuthKeyID(t *testing.T) {
	client := &fakeUserupdatesClient{}
	core := newUpdatesCore(client)
	core.MD.PermAuthKeyId = 0

	_, err := core.UpdatesGetState(&tg.TLUpdatesGetState{})
	if err != tg.ErrAuthKeyPermEmpty {
		t.Fatalf("UpdatesGetState() error = %v, want ErrAuthKeyPermEmpty", err)
	}
	if client.gotState != nil {
		t.Fatalf("userupdates must not be called when perm auth key is missing: %#v", client.gotState)
	}
}

func TestUpdatesGetDifferenceRejectsMissingPermAuthKeyID(t *testing.T) {
	client := &fakeUserupdatesClient{}
	core := newUpdatesCore(client)
	core.MD.PermAuthKeyId = 0

	_, err := core.UpdatesGetDifference(&tg.TLUpdatesGetDifference{Pts: 1, Date: 100, Qts: 0})
	if err != tg.ErrAuthKeyPermEmpty {
		t.Fatalf("UpdatesGetDifference() error = %v, want ErrAuthKeyPermEmpty", err)
	}
	if client.gotDifference != nil {
		t.Fatalf("userupdates must not be called when perm auth key is missing: %#v", client.gotDifference)
	}
}

func int32Ptr(v int32) *int32 {
	return &v
}
