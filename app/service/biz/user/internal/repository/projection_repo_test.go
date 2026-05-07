package repository

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestProjectionBundleOrdersViewerAndMissingIds(t *testing.T) {
	facts := projectionFacts{
		Users: map[int64]*projectionUserFact{
			1: {User: testUserData(1, "One")},
			2: {User: testUserData(2, "Two")},
		},
	}
	got := buildProjectionBundle(normalizedProjectionRequest{
		ViewerUserIds:  []int64{1},
		TargetUserIds:  []int64{2, 3, 1},
		HydrateUserIds: []int64{1, 2, 3},
		WithFacts:      true,
	}, facts)
	if len(got.ViewerUsers) != 1 || got.ViewerUsers[0].ViewerUserId != 1 {
		t.Fatalf("viewer users = %#v", got.ViewerUsers)
	}
	if len(got.ViewerUsers[0].Users) != 2 {
		t.Fatalf("users = %#v", got.ViewerUsers[0].Users)
	}
	firstUser, ok := got.ViewerUsers[0].Users[0].(*tg.TLUser)
	if !ok {
		t.Fatalf("first user = %T", got.ViewerUsers[0].Users[0])
	}
	secondUser, ok := got.ViewerUsers[0].Users[1].(*tg.TLUser)
	if !ok {
		t.Fatalf("second user = %T", got.ViewerUsers[0].Users[1])
	}
	if firstUser.Id != 2 || secondUser.Id != 1 {
		t.Fatalf("user order = %#v", got.ViewerUsers[0].Users)
	}
	if len(got.MissingUserIds) != 1 || got.MissingUserIds[0] != 3 {
		t.Fatalf("missing = %v", got.MissingUserIds)
	}
	if len(got.Facts) != 2 || got.Facts[0].ToImmutableUser().User.ToUserData().Id != 1 || got.Facts[1].ToImmutableUser().User.ToUserData().Id != 2 {
		t.Fatalf("facts = %v", got.Facts)
	}
}

func testUserData(id int64, firstName string) tg.UserDataClazz {
	return tg.MakeTLUserData(&tg.TLUserData{Id: id, AccessHash: id * 10, FirstName: firstName, LastName: "User"})
}
