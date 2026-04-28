package core

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeReadRepo struct {
	mutableChat       *tg.MutableChat
	mutableChats      []*tg.MutableChat
	participantIDs    []int64
	userChatRows      []repository.UserChatIDList
	err               error
	searchCalls       int
	lastSearchSelfID  int64
	lastSearchQ       string
	lastSearchOffset  int64
	lastSearchLimit   int32
	lastMyListCreator bool
}

func (f *fakeReadRepo) GetMutableChat(ctx context.Context, chatID int64, participantIDs ...int64) (*tg.MutableChat, error) {
	return f.mutableChat, f.err
}

func (f *fakeReadRepo) GetMutableChatByLink(ctx context.Context, link string) (*tg.MutableChat, error) {
	return f.mutableChat, f.err
}

func (f *fakeReadRepo) GetChatBySelfID(ctx context.Context, chatID, selfID int64) (*tg.MutableChat, error) {
	return f.mutableChat, f.err
}

func (f *fakeReadRepo) GetChatListByIDList(ctx context.Context, ids []int64) ([]*tg.MutableChat, error) {
	return f.mutableChats, f.err
}

func (f *fakeReadRepo) GetChatParticipantIDList(ctx context.Context, chatID int64) ([]int64, error) {
	return f.participantIDs, f.err
}

func (f *fakeReadRepo) GetUsersChatIDList(ctx context.Context, userIDs []int64) ([]repository.UserChatIDList, error) {
	return f.userChatRows, f.err
}

func (f *fakeReadRepo) GetMyChatList(ctx context.Context, userID int64, isCreator bool) ([]*tg.MutableChat, error) {
	f.lastMyListCreator = isCreator
	return f.mutableChats, f.err
}

func (f *fakeReadRepo) Search(ctx context.Context, selfID int64, q string, offset int64, limit int32) ([]*tg.MutableChat, error) {
	f.searchCalls++
	f.lastSearchSelfID = selfID
	f.lastSearchQ = q
	f.lastSearchOffset = offset
	f.lastSearchLimit = limit
	return f.mutableChats, f.err
}

func newTestCore(repo *fakeReadRepo) *ChatCore {
	return &ChatCore{
		ctx:      context.Background(),
		readRepo: repo,
	}
}

func testMutableChat(id int64) *tg.MutableChat {
	return tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{Id: id, Title: "chat"}).ToImmutableChat(),
	}).ToMutableChat()
}

func TestChatGetMutableChatWrapsRepositoryResult(t *testing.T) {
	want := testMutableChat(10)
	repo := &fakeReadRepo{mutableChat: want}

	got, err := newTestCore(repo).ChatGetMutableChat(&chat.TLChatGetMutableChat{ChatId: 10})
	if err != nil {
		t.Fatalf("ChatGetMutableChat error: %v", err)
	}
	if got != want {
		t.Fatalf("ChatGetMutableChat = %p, want %p", got, want)
	}

	repo.err = chat.ErrChatNotFound
	if _, err := newTestCore(repo).ChatGetMutableChat(&chat.TLChatGetMutableChat{ChatId: 10}); !errors.Is(err, chat.ErrChatNotFound) {
		t.Fatalf("ChatGetMutableChat error = %v, want ErrChatNotFound", err)
	}
}

func TestChatGetChatListByIdListWrapsVectorInRepositoryOrder(t *testing.T) {
	chats := []*tg.MutableChat{testMutableChat(20), testMutableChat(10)}
	got, err := newTestCore(&fakeReadRepo{mutableChats: chats}).ChatGetChatListByIdList(&chat.TLChatGetChatListByIdList{IdList: []int64{20, 10}})
	if err != nil {
		t.Fatalf("ChatGetChatListByIdList error: %v", err)
	}
	if len(got.Datas) != 2 || got.Datas[0] != chats[0] || got.Datas[1] != chats[1] {
		t.Fatalf("vector datas = %#v, want repository order", got.Datas)
	}
}

func TestChatGetChatListByIdListPropagatesRepositoryError(t *testing.T) {
	repo := &fakeReadRepo{err: chat.ErrChatStorage}
	_, err := newTestCore(repo).ChatGetChatListByIdList(&chat.TLChatGetChatListByIdList{IdList: []int64{10}})
	if !errors.Is(err, chat.ErrChatStorage) {
		t.Fatalf("ChatGetChatListByIdList error = %v, want ErrChatStorage", err)
	}
}

func TestChatSearchNormalizesQueryAndLimit(t *testing.T) {
	repo := &fakeReadRepo{mutableChats: []*tg.MutableChat{testMutableChat(10)}}
	core := newTestCore(repo)

	empty, err := core.ChatSearch(&chat.TLChatSearch{SelfId: 1, Q: "ab", Limit: 10})
	if err != nil {
		t.Fatalf("short ChatSearch error: %v", err)
	}
	if len(empty.Datas) != 0 || repo.searchCalls != 0 {
		t.Fatalf("short query returned %d chats and made %d search calls", len(empty.Datas), repo.searchCalls)
	}

	got, err := core.ChatSearch(&chat.TLChatSearch{SelfId: 1, Q: "team", Offset: 7, Limit: 99})
	if err != nil {
		t.Fatalf("ChatSearch error: %v", err)
	}
	if len(got.Datas) != 1 || got.Datas[0] != repo.mutableChats[0] {
		t.Fatalf("ChatSearch datas = %#v", got.Datas)
	}
	if repo.lastSearchQ != "%team%" || repo.lastSearchLimit != 50 || repo.lastSearchOffset != 7 || repo.lastSearchSelfID != 1 {
		t.Fatalf("search call = q:%q limit:%d offset:%d self:%d", repo.lastSearchQ, repo.lastSearchLimit, repo.lastSearchOffset, repo.lastSearchSelfID)
	}
}

func TestChatSearchPropagatesRepositoryError(t *testing.T) {
	repo := &fakeReadRepo{err: chat.ErrChatStorage}
	_, err := newTestCore(repo).ChatSearch(&chat.TLChatSearch{SelfId: 1, Q: "team", Limit: 10})
	if !errors.Is(err, chat.ErrChatStorage) {
		t.Fatalf("ChatSearch error = %v, want ErrChatStorage", err)
	}
}

func TestChatGetUsersChatIdListGroupsRowsByUser(t *testing.T) {
	rows := []repository.UserChatIDList{
		{UserID: 1, ChatIDList: []int64{10, 11}},
		{UserID: 2, ChatIDList: []int64{20}},
	}
	got, err := newTestCore(&fakeReadRepo{userChatRows: rows}).ChatGetUsersChatIdList(&chat.TLChatGetUsersChatIdList{Id: []int64{1, 2}})
	if err != nil {
		t.Fatalf("ChatGetUsersChatIdList error: %v", err)
	}
	want := map[int64][]int64{1: {10, 11}, 2: {20}}
	if len(got.Datas) != len(want) {
		t.Fatalf("len(datas) = %d, want %d", len(got.Datas), len(want))
	}
	for _, item := range got.Datas {
		if !reflect.DeepEqual(item.ChatIdList, want[item.UserId]) {
			t.Fatalf("user %d chat ids = %#v, want %#v", item.UserId, item.ChatIdList, want[item.UserId])
		}
	}
}

func TestChatReadHandlersDoNotReturnMethodNotImplOrTGErrors(t *testing.T) {
	files := []string{
		"chat.getMutableChat_handler.go",
		"chat.getMutableChatByLink_handler.go",
		"chat.getChatListByIdList_handler.go",
		"chat.getChatBySelfId_handler.go",
		"chat.getChatParticipantIdList_handler.go",
		"chat.getUsersChatIdList_handler.go",
		"chat.getMyChatList_handler.go",
		"chat.search_handler.go",
	}

	for _, name := range files {
		data, err := os.ReadFile(filepath.Join(".", name))
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		s := string(data)
		for _, bad := range []string{"ErrMethodNotImpl", "method Chat", "not impl", "tg.Err"} {
			if strings.Contains(s, bad) {
				t.Fatalf("%s still contains %q", name, bad)
			}
		}
	}
}
