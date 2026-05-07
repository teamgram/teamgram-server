package core

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestUserIDFromInputUser(t *testing.T) {
	selfID := int64(1001)

	got, err := userIDFromInputUser(selfID, tg.MakeTLInputUserSelf(&tg.TLInputUserSelf{}))
	if err != nil {
		t.Fatalf("self input returned error: %v", err)
	}
	if got != selfID {
		t.Fatalf("self input id = %d, want %d", got, selfID)
	}

	got, err = userIDFromInputUser(selfID, tg.MakeTLInputUser(&tg.TLInputUser{UserId: 2002, AccessHash: 33}))
	if err != nil {
		t.Fatalf("user input returned error: %v", err)
	}
	if got != 2002 {
		t.Fatalf("user input id = %d, want 2002", got)
	}

	if _, err = userIDFromInputUser(selfID, tg.MakeTLInputUserEmpty(&tg.TLInputUserEmpty{})); err != tg.ErrUserIdInvalid {
		t.Fatalf("inputUserEmpty error = %v, want USER_ID_INVALID", err)
	}
}
