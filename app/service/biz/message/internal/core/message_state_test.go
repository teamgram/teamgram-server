package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMessageGetUserMessageListReturnsStablePlaceholderBoxes(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetUserMessageList(&message.TLMessageGetUserMessageList{
		UserId: 1,
		IdList: []int32{10, 11},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected message box vector, got nil")
	}
	if len(result.Datas) != 2 {
		t.Fatalf("expected 2 placeholder boxes, got %d items", len(result.Datas))
	}

	first := result.Datas[0]
	if first.MessageId != 10 {
		t.Fatalf("expected first message_id=10, got %d", first.MessageId)
	}
	second := result.Datas[1]
	if second.MessageId != 11 {
		t.Fatalf("expected second message_id=11, got %d", second.MessageId)
	}
}

func TestMessageGetUserMessageReturnsPlaceholderMessageBox(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetUserMessage(&message.TLMessageGetUserMessage{
		UserId: 1,
		Id:     10,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected message box, got nil")
	}
	if result.MessageId != 10 {
		t.Fatalf("expected message_id=10, got %d", result.MessageId)
	}
	if result.UserId != 1 {
		t.Fatalf("expected user_id=1, got %d", result.UserId)
	}
}

func TestMessageGetHistoryMessagesReturnsStablePlaceholderBoxes(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetHistoryMessages(&message.TLMessageGetHistoryMessages{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
		Limit:    2,
		OffsetId: 20,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected history placeholder vector, got nil")
	}
	if len(result.Datas) != 2 {
		t.Fatalf("expected 2 history boxes, got %d items", len(result.Datas))
	}

	first := result.Datas[0]
	if first.MessageId != 20 {
		t.Fatalf("expected first message_id=20, got %d", first.MessageId)
	}
	if first.PeerId != 42 || first.PeerType != 2 {
		t.Fatalf("expected first peer=2/42, got %d/%d", first.PeerType, first.PeerId)
	}

	second := result.Datas[1]
	if second.MessageId != 21 {
		t.Fatalf("expected second message_id=21, got %d", second.MessageId)
	}
}

func TestMessageGetHistoryMessagesReturnsEmptyForZeroLimit(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetHistoryMessages(&message.TLMessageGetHistoryMessages{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
		Limit:    0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected history placeholder vector, got nil")
	}
	if len(result.Datas) != 0 {
		t.Fatalf("expected empty history for zero limit, got %d items", len(result.Datas))
	}
}

func TestMessageGetPeerUserMessageReturnsPeerScopedPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetPeerUserMessage(&message.TLMessageGetPeerUserMessage{
		UserId:     1,
		PeerUserId: 42,
		MsgId:      88,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected peer-scoped message box, got nil")
	}
	if result.MessageId != 88 {
		t.Fatalf("expected message_id=88, got %d", result.MessageId)
	}
	if result.PeerType != 2 || result.PeerId != 42 {
		t.Fatalf("expected peer=2/42, got %d/%d", result.PeerType, result.PeerId)
	}
}

func TestMessageGetPeerUserMessageIdReturnsMsgIDPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetPeerUserMessageId(&message.TLMessageGetPeerUserMessageId{
		UserId:     1,
		PeerUserId: 42,
		MsgId:      88,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected int32 placeholder, got nil")
	}
	if result.V != 88 {
		t.Fatalf("expected msg_id passthrough=88, got %d", result.V)
	}
}

func TestMessageGetUserMessageListByDataIdListReturnsStablePlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetUserMessageListByDataIdList(&message.TLMessageGetUserMessageListByDataIdList{
		UserId: 1,
		IdList: []int64{50, 51},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected message box vector, got nil")
	}
	if len(result.Datas) != 2 {
		t.Fatalf("expected 2 placeholders, got %d", len(result.Datas))
	}
	if result.Datas[0].MessageId != 50 || result.Datas[1].MessageId != 51 {
		t.Fatalf("expected message ids 50/51, got %d/%d", result.Datas[0].MessageId, result.Datas[1].MessageId)
	}
}

func TestMessageGetUserMessageListByDataIdUserIdListReturnsPeerScopedPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetUserMessageListByDataIdUserIdList(&message.TLMessageGetUserMessageListByDataIdUserIdList{
		Id:         60,
		UserIdList: []int64{7, 8},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected message box vector, got nil")
	}
	if len(result.Datas) != 2 {
		t.Fatalf("expected 2 placeholders, got %d", len(result.Datas))
	}
	if result.Datas[0].UserId != 7 || result.Datas[1].UserId != 8 {
		t.Fatalf("expected user ids 7/8, got %d/%d", result.Datas[0].UserId, result.Datas[1].UserId)
	}
	if result.Datas[0].MessageId != 60 || result.Datas[1].MessageId != 60 {
		t.Fatalf("expected shared message id 60, got %d/%d", result.Datas[0].MessageId, result.Datas[1].MessageId)
	}
}

func TestMessageGetHistoryMessagesCountReturnsPlaceholderCount(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetHistoryMessagesCount(&message.TLMessageGetHistoryMessagesCount{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.V != 1 {
		t.Fatalf("expected placeholder count=1, got %#v", result)
	}
}

func TestMessageGetSearchCounterReturnsPlaceholderCount(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetSearchCounter(&message.TLMessageGetSearchCounter{
		UserId:    1,
		PeerType:  2,
		PeerId:    42,
		MediaType: 0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.V != 1 {
		t.Fatalf("expected placeholder count=1, got %#v", result)
	}
}

func TestMessageGetUnreadMentionsCountReturnsPlaceholderCount(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetUnreadMentionsCount(&message.TLMessageGetUnreadMentionsCount{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.V != 1 {
		t.Fatalf("expected placeholder count=1, got %#v", result)
	}
}

func TestMessageSearchReturnsScopedPlaceholderBoxes(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageSearch(&message.TLMessageSearch{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
		Offset:   3,
		Limit:    2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(result.Datas) != 2 {
		t.Fatalf("expected 2 results, got %d", len(result.Datas))
	}
	if result.Datas[0].MessageId != 4 || result.Datas[1].MessageId != 5 {
		t.Fatalf("expected message ids 4/5, got %d/%d", result.Datas[0].MessageId, result.Datas[1].MessageId)
	}
}

func TestMessageSearchV2ReturnsScopedPlaceholderBoxes(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageSearchV2(&message.TLMessageSearchV2{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
		OffsetId: 9,
		Limit:    2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(result.Datas) != 2 || result.Datas[0].MessageId != 9 || result.Datas[1].MessageId != 10 {
		t.Fatalf("expected message ids 9/10, got %#v", result.Datas)
	}
}

func TestMessageSearchGlobalReturnsUserScopedPlaceholderBoxes(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageSearchGlobal(&message.TLMessageSearchGlobal{
		UserId: 1,
		Offset: 1,
		Limit:  2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(result.Datas) != 2 {
		t.Fatalf("expected 2 results, got %d", len(result.Datas))
	}
	if result.Datas[0].PeerId != 1 || result.Datas[1].PeerId != 1 {
		t.Fatalf("expected peer id 1/1, got %d/%d", result.Datas[0].PeerId, result.Datas[1].PeerId)
	}
}

func TestMessageSearchByMediaTypeReturnsScopedPlaceholderBoxes(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageSearchByMediaType(&message.TLMessageSearchByMediaType{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
		Offset:   2,
		Limit:    1,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(result.Datas) != 1 || result.Datas[0].MessageId != 3 {
		t.Fatalf("expected message id 3, got %#v", result.Datas)
	}
}

func TestMessageSearchByPinnedReturnsSingleScopedPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageSearchByPinned(&message.TLMessageSearchByPinned{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(result.Datas) != 1 || result.Datas[0].MessageId != 1 {
		t.Fatalf("expected single message id 1, got %#v", result.Datas)
	}
}

func TestMessageGetUnreadMentionsReturnsScopedPlaceholderBoxes(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetUnreadMentions(&message.TLMessageGetUnreadMentions{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
		OffsetId: 5,
		Limit:    2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(result.Datas) != 2 || result.Datas[0].MessageId != 5 || result.Datas[1].MessageId != 6 {
		t.Fatalf("expected message ids 5/6, got %#v", result.Datas)
	}
}

func TestMessageUpdatePinnedMessageIdReturnsPinnedBool(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageUpdatePinnedMessageId(&message.TLMessageUpdatePinnedMessageId{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
		Id:       7,
		Pinned:   tg.BoolTrueClazz,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result != tg.BoolTrue {
		t.Fatalf("expected BoolTrue, got %v", result)
	}
}

func TestMessageGetPinnedMessageIdListReturnsPlaceholderIDs(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetPinnedMessageIdList(&message.TLMessageGetPinnedMessageIdList{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(result.Datas) != 1 || result.Datas[0] != 1 {
		t.Fatalf("expected pinned ids [1], got %#v", result.Datas)
	}
}

func TestMessageGetLastTwoPinnedMessageIdReturnsPlaceholderIDs(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetLastTwoPinnedMessageId(&message.TLMessageGetLastTwoPinnedMessageId{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(result.Datas) != 2 || result.Datas[0] != 1 || result.Datas[1] != 2 {
		t.Fatalf("expected pinned ids [1 2], got %#v", result.Datas)
	}
}

func TestMessageUnPinAllMessagesReturnsEmptyPlaceholderList(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageUnPinAllMessages(&message.TLMessageUnPinAllMessages{
		UserId:   1,
		PeerType: 2,
		PeerId:   42,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(result.Datas) != 0 {
		t.Fatalf("expected empty list, got %#v", result.Datas)
	}
}
