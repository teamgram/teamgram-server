package core

import (
	"context"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/mention"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestRemakeMessageTextEntitiesDetectsGeneratedEntitiesWithUTF16Offsets(t *testing.T) {
	text := "😀 mail a@example.com visit https://example.com/a #go @alice"

	entities := remakeMessageTextEntities(text, nil, 100, false)

	assertEntityRange(t, entities, 0, tg.ClazzName_messageEntityEmail, text, "a@example.com")
	assertEntityRange(t, entities, 1, tg.ClazzName_messageEntityUrl, text, "https://example.com/a")
	assertEntityRange(t, entities, 2, tg.ClazzName_messageEntityHashtag, text, "#go")
	assertEntityRange(t, entities, 3, tg.ClazzName_messageEntityMention, text, "@alice")
}

func TestRemakeMessageTextEntitiesConvertsInputMentionName(t *testing.T) {
	inputSelf := tg.MakeTLInputMessageEntityMentionName(&tg.TLInputMessageEntityMentionName{
		Offset: 0,
		Length: 4,
		UserId: tg.MakeTLInputUserSelf(&tg.TLInputUserSelf{}),
	})
	inputUser := tg.MakeTLInputMessageEntityMentionName(&tg.TLInputMessageEntityMentionName{
		Offset: 5,
		Length: 4,
		UserId: tg.MakeTLInputUser(&tg.TLInputUser{UserId: 200, AccessHash: 1}),
	})

	entities := remakeMessageTextEntities("self user", []tg.MessageEntityClazz{inputSelf, inputUser}, 100, false)

	if len(entities) != 2 {
		t.Fatalf("entities len = %d, want 2", len(entities))
	}
	first, ok := entities[0].(*tg.TLMessageEntityMentionName)
	if !ok {
		t.Fatalf("entities[0] = %T, want *tg.TLMessageEntityMentionName", entities[0])
	}
	if first.UserId != 100 || first.Offset != 0 || first.Length != 4 {
		t.Fatalf("first mention name = %#v, want self user id 100", first)
	}
	second, ok := entities[1].(*tg.TLMessageEntityMentionName)
	if !ok {
		t.Fatalf("entities[1] = %T, want *tg.TLMessageEntityMentionName", entities[1])
	}
	if second.UserId != 200 || second.Offset != 5 || second.Length != 4 {
		t.Fatalf("second mention name = %#v, want input user id 200", second)
	}
}

func TestMessagesSendMessageGeneratesMessageTextEntities(t *testing.T) {
	var gotEntities []tg.MessageEntityClazz
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			message, ok := in.Message[0].Message.(*tg.TLMessage)
			if !ok {
				t.Fatalf("message = %T, want *tg.TLMessage", in.Message[0].Message)
			}
			gotEntities = message.Entities
			return testUpdates(), nil
		},
	}, 100, 200)

	text := "see https://teamgram.io #teamgram"
	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  text,
		RandomId: 42,
	})
	if err != nil {
		t.Fatalf("MessagesSendMessage() error = %v", err)
	}

	assertEntityRange(t, gotEntities, 0, tg.ClazzName_messageEntityUrl, text, "https://teamgram.io")
	assertEntityRange(t, gotEntities, 1, tg.ClazzName_messageEntityHashtag, text, "#teamgram")
}

func TestBuildMessageMediaOutboxGeneratesCaptionEntities(t *testing.T) {
	text := "caption https://teamgram.io #photo"

	outbox := buildMessageMediaOutbox(mediaOutboxInput{
		RandomId: 11,
		FromId:   100,
		PeerId:   200,
		Message:  text,
		Media:    tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{}),
	})

	message, ok := outbox.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("outbox message = %T, want *tg.TLMessage", outbox.Message)
	}
	assertEntityRange(t, message.Entities, 0, tg.ClazzName_messageEntityUrl, text, "https://teamgram.io")
	assertEntityRange(t, message.Entities, 1, tg.ClazzName_messageEntityHashtag, text, "#photo")
}

func assertEntityRange(t *testing.T, entities []tg.MessageEntityClazz, index int, wantClazz string, text string, part string) {
	t.Helper()
	if len(entities) <= index {
		t.Fatalf("entities len = %d, want index %d", len(entities), index)
	}
	if got := entities[index].MessageEntityClazzName(); got != wantClazz {
		t.Fatalf("entities[%d] clazz = %s, want %s", index, got, wantClazz)
	}
	offset, length := entityRange(entities[index])
	wantOffset, wantLength := utf16Range(text, part)
	if offset != wantOffset || length != wantLength {
		t.Fatalf("entities[%d] range = (%d,%d), want (%d,%d)", index, offset, length, wantOffset, wantLength)
	}
}

func utf16Range(text string, part string) (int32, int32) {
	start := strings.Index(text, part)
	if start < 0 {
		panic("test substring not found")
	}
	idx := mention.EncodeStringToUTF16Index(text)
	end := start + len(part)
	return int32(idx[start]), int32(idx[end] - idx[start])
}
