package repository

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestProjectUserSelfSetsMasterCompatibleFlags(t *testing.T) {
	facts := projectionFacts{
		Users: map[int64]*projectionUserFact{
			1001: {User: tg.MakeTLUserData(&tg.TLUserData{Id: 1001, AccessHash: 11, FirstName: "Self", LastName: "User", Phone: "100", Username: "self"})},
		},
	}
	got := projectUserForViewer(1001, 1001, facts)
	user, ok := got.(*tg.TLUser)
	if !ok {
		t.Fatalf("got %T, want *tg.TLUser", got)
	}
	if !user.Self || !user.Contact || !user.MutualContact || user.Phone == nil || *user.Phone != "100" {
		t.Fatalf("self projection = %+v", user)
	}
}

func TestProjectUserContactUsesViewerLocalName(t *testing.T) {
	facts := projectionFacts{
		Users: map[int64]*projectionUserFact{
			1002: {User: tg.MakeTLUserData(&tg.TLUserData{Id: 1002, AccessHash: 22, FirstName: "Server", LastName: "Name", Phone: "200"})},
		},
		Contacts: map[contactKey]*projectionContactFact{
			{OwnerUserId: 1001, ContactUserId: 1002}: {FirstName: "Local", LastName: "Friend", Phone: "555", Mutual: true},
		},
	}
	got := projectUserForViewer(1001, 1002, facts)
	user := got.(*tg.TLUser)
	if !user.Contact || !user.MutualContact || user.FirstName == nil || *user.FirstName != "Local" || user.Phone == nil || *user.Phone != "555" {
		t.Fatalf("contact projection = %+v", user)
	}
}

func TestProjectDeletedUserDoesNotLeakFields(t *testing.T) {
	facts := projectionFacts{
		Users: map[int64]*projectionUserFact{
			1002: {User: tg.MakeTLUserData(&tg.TLUserData{Id: 1002, AccessHash: 22, FirstName: "Deleted", LastName: "User", Phone: "200", Deleted: true})},
		},
	}
	got := projectUserForViewer(1001, 1002, facts)
	user := got.(*tg.TLUser)
	if !user.Deleted || user.Phone != nil || user.FirstName != nil || user.LastName != nil || user.Status != nil || user.Photo != nil {
		t.Fatalf("deleted projection = %+v", user)
	}
}

func TestProjectUserNonContactDoesNotLeakPhoneByDefault(t *testing.T) {
	facts := projectionFacts{
		Users: map[int64]*projectionUserFact{
			1002: {User: tg.MakeTLUserData(&tg.TLUserData{Id: 1002, AccessHash: 22, FirstName: "Server", LastName: "Name", Phone: "200"})},
		},
	}
	got := projectUserForViewer(1001, 1002, facts)
	user := got.(*tg.TLUser)
	if user.Phone != nil {
		t.Fatalf("non-contact projection leaked phone: %+v", user)
	}
}
