package repository

import (
	"encoding/hex"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestAuthSeqPayloadIDFromHash(t *testing.T) {
	hash := payload.HashBytes([]byte("payload"))
	got := AuthSeqPayloadID(hash)
	want := hex.EncodeToString(hash)
	if got != want {
		t.Fatalf("AuthSeqPayloadID() = %q, want %q", got, want)
	}
}

func TestAuthSeqPayloadIDRejectsEmptyHash(t *testing.T) {
	if got := AuthSeqPayloadID(nil); got != "" {
		t.Fatalf("AuthSeqPayloadID(nil) = %q, want empty", got)
	}
}
