package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type failingMediaReader struct {
	err error
}

func (f failingMediaReader) GetChatPhoto(ctx context.Context, photoID int64) (*tg.Photo, error) {
	return nil, f.err
}

func TestGetMutableChatRequiresIntegrationDB(t *testing.T) {
	t.Skip("repository aggregate read requires a MySQL fixture")
}

func TestGetChatListRequiresIntegrationDB(t *testing.T) {
	t.Skip("repository aggregate list read requires a MySQL fixture")
}

func TestSearchRequiresIntegrationDB(t *testing.T) {
	t.Skip("repository search read requires a MySQL fixture")
}

func TestOrderChatRowsByIDsPreservesInputOrderAndSkipsMissing(t *testing.T) {
	rows := []model.Chats{
		{Id: 30, Title: "thirty"},
		{Id: 10, Title: "ten"},
	}
	got := orderChatRowsByIDs([]int64{10, 20, 30, 10}, rows)
	if len(got) != 3 {
		t.Fatalf("len(ordered) = %d, want 3", len(got))
	}
	gotIDs := []int64{got[0].Id, got[1].Id, got[2].Id}
	if !reflect.DeepEqual(gotIDs, []int64{10, 30, 10}) {
		t.Fatalf("ordered ids = %#v, want [10 30 10]", gotIDs)
	}
}

func TestGroupUserChatIDRowsGroupsInFirstSeenUserOrder(t *testing.T) {
	rows := []model.ChatParticipants{
		{UserId: 2, ChatId: 20},
		{UserId: 1, ChatId: 10},
		{UserId: 2, ChatId: 21},
	}
	got := groupUserChatIDRows(rows)
	want := []UserChatIDList{
		{UserID: 2, ChatIDList: []int64{20, 21}},
		{UserID: 1, ChatIDList: []int64{10}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("groupUserChatIDRows = %#v, want %#v", got, want)
	}
}

func TestGetChatParticipantIDListStateFilterIncludesMigrated(t *testing.T) {
	for _, state := range []int32{chatpb.ChatMemberStateNormal, chatpb.ChatMemberStateMigrated} {
		if !isListableChatParticipantState(state) {
			t.Fatalf("state %d should be listable", state)
		}
	}
	for _, state := range []int32{chatpb.ChatMemberStateLeft, chatpb.ChatMemberStateKicked} {
		if isListableChatParticipantState(state) {
			t.Fatalf("state %d should not be listable", state)
		}
	}
}

func TestPhotoFallbackIsNonFatal(t *testing.T) {
	r := &Repository{mediaReader: failingMediaReader{err: errors.New("media unavailable")}}
	got := r.makeMutableChatFromRows(context.Background(), &model.Chats{
		Id:      10,
		PhotoId: 99,
		Title:   "chat",
	}, nil)
	if got == nil || got.Chat == nil || got.Chat.Photo == nil {
		t.Fatalf("mutable chat photo missing: %#v", got)
	}
	photo, ok := got.Chat.Photo.(*tg.TLPhotoEmpty)
	if !ok || photo.Id != 99 {
		t.Fatalf("photo fallback = %#v, want photoEmpty id 99", got.Chat.Photo)
	}
}

func TestMakeMutableChatFromRowsNilSafe(t *testing.T) {
	got := (&Repository{}).makeMutableChatFromRows(context.Background(), nil, nil)
	if got == nil {
		t.Fatal("nil row returned nil mutable chat")
	}
	if got.Chat != nil || len(got.ChatParticipants) != 0 {
		t.Fatalf("nil row mutable chat = %#v", got)
	}
}
