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

func TestProjectImmutableUser(t *testing.T) {
	immutable := immutableUserFixture(2002, "Ada", "Lovelace", "ada")

	projected := projectImmutableUser(immutable)
	user, ok := projected.(*tg.TLUser)
	if !ok {
		t.Fatalf("projected user = %T, want *tg.TLUser", projected)
	}
	if user.Id != 2002 {
		t.Fatalf("user id = %d, want 2002", user.Id)
	}
	if user.FirstName == nil || *user.FirstName != "Ada" {
		t.Fatalf("first name = %v, want Ada", user.FirstName)
	}
	if user.LastName == nil || *user.LastName != "Lovelace" {
		t.Fatalf("last name = %v, want Lovelace", user.LastName)
	}
	if user.Username == nil || *user.Username != "ada" {
		t.Fatalf("username = %v, want ada", user.Username)
	}
	if len(user.Usernames) != 1 {
		t.Fatalf("usernames len = %d, want 1", len(user.Usernames))
	}
	username := user.Usernames[0]
	if username.Username != "ada" || !username.Active {
		t.Fatalf("usernames[0] = %+v, want active ada", username)
	}
	if user.Status == nil {
		t.Fatal("status is nil")
	}
}

func TestProjectSelfImmutableUserMarksUsernameEditable(t *testing.T) {
	immutable := immutableUserFixture(2002, "Ada", "Lovelace", "ada")

	projected := projectSelfImmutableUser(immutable)
	user, ok := projected.(*tg.TLUser)
	if !ok {
		t.Fatalf("projected user = %T, want *tg.TLUser", projected)
	}
	if len(user.Usernames) != 1 {
		t.Fatalf("usernames len = %d, want 1", len(user.Usernames))
	}
	username := user.Usernames[0]
	if username.Username != "ada" || !username.Active || !username.Editable {
		t.Fatalf("usernames[0] = %+v, want active editable ada", username)
	}
}

func TestProjectDeletedImmutableUser(t *testing.T) {
	immutable := tg.MakeTLImmutableUser(&tg.TLImmutableUser{
		User: tg.MakeTLUserData(&tg.TLUserData{Id: 2002, Deleted: true}),
	})

	projected := projectImmutableUser(immutable)
	user, ok := projected.(*tg.TLUser)
	if !ok {
		t.Fatalf("projected user = %T, want *tg.TLUser", projected)
	}
	if !user.Deleted || user.Id != 2002 {
		t.Fatalf("deleted user = %+v, want deleted user id 2002", user)
	}
}
