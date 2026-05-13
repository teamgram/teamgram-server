package projection

import (
	"encoding/json"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/eventtypes"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestProjectUserEventChatNewMessageKeepsPeerChatAndFromUser(t *testing.T) {
	body, err := json.Marshal(payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 7001,
		MessageID:          101,
		PeerSeq:            1,
		PeerType:           payload.PeerTypeChat,
		PeerID:             55,
		FromUserID:         1001,
		ToUserID:           1002,
		Date:               1778584910,
		MessageText:        "hello chat",
		Out:                false,
	})
	if err != nil {
		t.Fatal(err)
	}
	result, err := ProjectUserEvent(eventtypes.UserEvent{
		UserID:             1002,
		EventType:          eventtypes.EventTypeNewMessage,
		PeerType:           payload.PeerTypeChat,
		PeerID:             55,
		Pts:                11,
		PtsCount:           1,
		EventSchemaVersion: payload.MessageEventSchemaVersionV3,
		EventCodec:         eventtypes.PayloadCodecJSON,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes(body),
	}, ModeDifference)
	if err != nil {
		t.Fatalf("ProjectUserEvent() error = %v", err)
	}
	msg, ok := result.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("projected message = %T, want *tg.TLMessage", result.Message)
	}
	peer, ok := msg.PeerId.(*tg.TLPeerChat)
	if !ok || peer.ChatId != 55 {
		t.Fatalf("peer_id = %#v, want peerChat 55", msg.PeerId)
	}
	from, ok := msg.FromId.(*tg.TLPeerUser)
	if !ok || from.UserId != 1001 {
		t.Fatalf("from_id = %#v, want peerUser 1001", msg.FromId)
	}
	if msg.Out {
		t.Fatalf("out = true, want false for receiver view")
	}
}
