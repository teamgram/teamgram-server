package core

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestInputHelpers(t *testing.T) {
	selfID := int64(1001)

	if got, err := userIDFromInputUser(selfID, tg.MakeTLInputUserSelf(&tg.TLInputUserSelf{})); err != nil || got != selfID {
		t.Fatalf("self input = %d, %v; want %d, nil", got, err, selfID)
	}
	if got, err := userIDFromInputUser(selfID, tg.MakeTLInputUser(&tg.TLInputUser{UserId: 2002})); err != nil || got != 2002 {
		t.Fatalf("user input = %d, %v; want 2002, nil", got, err)
	}
	if _, err := userIDFromInputUser(selfID, tg.MakeTLInputUserEmpty(&tg.TLInputUserEmpty{})); err != tg.ErrUserIdInvalid {
		t.Fatalf("inputUserEmpty error = %v, want USER_ID_INVALID", err)
	}

	if got, err := photoIDFromInputPhoto(tg.MakeTLInputPhoto(&tg.TLInputPhoto{Id: 3003})); err != nil || got != 3003 {
		t.Fatalf("photo id = %d, %v; want 3003, nil", got, err)
	}
	if _, err := photoIDFromInputPhoto(tg.MakeTLInputPhotoEmpty(&tg.TLInputPhotoEmpty{})); err != tg.ErrInputRequestInvalid {
		t.Fatalf("inputPhotoEmpty error = %v, want INPUT_REQUEST_INVALID", err)
	}

	if got, err := documentIDFromInputDocument(tg.MakeTLInputDocument(&tg.TLInputDocument{Id: 4004})); err != nil || got != 4004 {
		t.Fatalf("document id = %d, %v; want 4004, nil", got, err)
	}
	if got := optionalDocumentID(tg.MakeTLInputDocument(&tg.TLInputDocument{Id: 5005})); got == nil || *got != 5005 {
		t.Fatalf("optional document id = %v, want 5005", got)
	}
	if got := optionalDocumentID(nil); got != nil {
		t.Fatalf("nil optional document id = %v, want nil", got)
	}

	if got, err := channelIDFromInputChannel(tg.MakeTLInputChannel(&tg.TLInputChannel{ChannelId: 6006})); err != nil || got != 6006 {
		t.Fatalf("channel id = %d, %v; want 6006, nil", got, err)
	}
	if _, err := channelIDFromInputChannel(tg.MakeTLInputChannelEmpty(&tg.TLInputChannelEmpty{})); err != tg.ErrInputRequestInvalid {
		t.Fatalf("inputChannelEmpty error = %v, want INPUT_REQUEST_INVALID", err)
	}
}
