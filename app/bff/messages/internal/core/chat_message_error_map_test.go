package core

import (
	"fmt"
	"testing"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMapChatMessageErrorUsesMessagePeerSemantics(t *testing.T) {
	for _, err := range []error{chatpb.ErrChatNotFound, chatpb.ErrParticipantInvalid, chatpb.ErrUserNotParticipant} {
		if got := mapChatMessageError(fmt.Errorf("wrapped: %w", err)); got != tg.Err400PeerIdInvalid {
			t.Fatalf("mapChatMessageError(%v) = %v, want PEER_ID_INVALID", err, got)
		}
	}
	if got := mapChatMessageError(chatpb.ErrChatWriteForbidden); got != tg.ErrChatWriteForbidden {
		t.Fatalf("write forbidden maps to %v, want CHAT_WRITE_FORBIDDEN", got)
	}
}
