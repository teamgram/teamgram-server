package repository

import (
	"errors"
	"testing"

	chatprojection "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chatprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestBuildChatProjectionBundlePreservesDuplicatesAndMissingOrder(t *testing.T) {
	mutable := map[int64]tg.MutableChatClazz{
		3001: tg.MakeTLMutableChat(&tg.TLMutableChat{
			Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{Id: 3001, Creator: 1001, Title: "one", Date: 10}),
		}),
	}

	got, err := buildChatProjectionBundle([]int64{1001, 2002}, []int64{3001, 9999, 3001, 8888, 9999}, mutable)
	if err != nil {
		t.Fatalf("buildChatProjectionBundle error = %v", err)
	}
	if len(got.ViewerChats) != 2 {
		t.Fatalf("viewer count = %d, want 2", len(got.ViewerChats))
	}
	if len(got.ViewerChats[0].Chats) != 2 || chatID(got.ViewerChats[0].Chats[0]) != 3001 || chatID(got.ViewerChats[0].Chats[1]) != 3001 {
		t.Fatalf("viewer chats = %#v, want duplicate projected chat", got.ViewerChats[0].Chats)
	}
	first := got.ViewerChats[0].Chats[0].(*tg.TLChat)
	secondViewer := got.ViewerChats[1].Chats[0].(*tg.TLChat)
	if !first.Creator || secondViewer.Creator {
		t.Fatalf("Creator flags = first:%t second:%t, want true false", first.Creator, secondViewer.Creator)
	}
	if len(got.MissingChatIds) != 2 || got.MissingChatIds[0] != 9999 || got.MissingChatIds[1] != 8888 {
		t.Fatalf("missing = %#v, want [9999 8888]", got.MissingChatIds)
	}
}

func TestBuildChatProjectionBundleReturnsDateOverflow(t *testing.T) {
	mutable := map[int64]tg.MutableChatClazz{
		3001: tg.MakeTLMutableChat(&tg.TLMutableChat{
			Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{Id: 3001, Date: int64(1 << 31)}),
		}),
	}
	_, err := buildChatProjectionBundle([]int64{1001}, []int64{3001}, mutable)
	if !errors.Is(err, chatprojection.ErrChatDateOverflow) {
		t.Fatalf("buildChatProjectionBundle error = %v, want ErrChatDateOverflow", err)
	}
}

func TestNormalizeChatProjectionRequest(t *testing.T) {
	viewers, targets, uniqueTargets := normalizeChatProjectionRequest([]int64{1001, 0, 1001, 2002}, []int64{3001, -1, 3001, 3002})
	if len(viewers) != 2 || viewers[0] != 1001 || viewers[1] != 2002 {
		t.Fatalf("viewers = %#v, want [1001 2002]", viewers)
	}
	if len(targets) != 3 || targets[0] != 3001 || targets[1] != 3001 || targets[2] != 3002 {
		t.Fatalf("targets = %#v, want [3001 3001 3002]", targets)
	}
	if len(uniqueTargets) != 2 || uniqueTargets[0] != 3001 || uniqueTargets[1] != 3002 {
		t.Fatalf("uniqueTargets = %#v, want [3001 3002]", uniqueTargets)
	}
}

func TestChatProjectionEmptyInputsReturnEmptyBundle(t *testing.T) {
	got, err := buildChatProjectionBundle(nil, []int64{3001}, nil)
	if err != nil {
		t.Fatalf("empty viewers error = %v", err)
	}
	if got == nil || len(got.ViewerChats) != 0 || len(got.MissingChatIds) != 0 {
		t.Fatalf("empty viewers bundle = %#v, want empty bundle", got)
	}
	got, err = buildChatProjectionBundle([]int64{1001}, nil, nil)
	if err != nil {
		t.Fatalf("empty targets error = %v", err)
	}
	if got == nil || len(got.ViewerChats) != 0 || len(got.MissingChatIds) != 0 {
		t.Fatalf("empty targets bundle = %#v, want empty bundle", got)
	}
}

func chatID(chat tg.ChatClazz) int64 {
	if x, ok := chat.(*tg.TLChat); ok && x != nil {
		return x.Id
	}
	return 0
}
