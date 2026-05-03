package core

import (
	"context"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestContactsGetBirthdays(t *testing.T) {
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		getBirthdays: func(_ context.Context, in *userpb.TLUserGetBirthdays) (*userpb.VectorContactBirthday, error) {
			if in.UserId != 1001 {
				t.Fatalf("birthdays user id = %d, want 1001", in.UserId)
			}
			return &userpb.VectorContactBirthday{Datas: []tg.ContactBirthdayClazz{
				tg.MakeTLContactBirthday(&tg.TLContactBirthday{ContactId: 2002}),
			}}, nil
		},
		getMutableUsersV2: func(_ context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
			if len(in.Id) != 1 || in.Id[0] != 2002 {
				t.Fatalf("mutable users request = %+v", in)
			}
			return tg.MakeTLMutableUsers(&tg.TLMutableUsers{Users: []tg.ImmutableUserClazz{
				immutableUserFixture(2002, "Ada", "", "ada"),
			}}).ToMutableUsers(), nil
		},
	}, nil, 1001)
	got, err := core.ContactsGetBirthdays(&tg.TLContactsGetBirthdays{})
	if err != nil {
		t.Fatalf("ContactsGetBirthdays returned error: %v", err)
	}
	if len(got.Contacts) != 1 || len(got.Users) != 1 {
		t.Fatalf("birthdays response = %+v", got)
	}
}

func TestUsersGetSavedMusicEmpty(t *testing.T) {
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		getSavedMusicIDList: func(_ context.Context, in *userpb.TLUserGetSavedMusicIdList) (*userpb.VectorLong, error) {
			if in.UserId != 1001 {
				t.Fatalf("saved music user id = %d, want 1001", in.UserId)
			}
			return &userpb.VectorLong{}, nil
		},
	}, &fakeMediaClient{}, 1001)
	got, err := core.UsersGetSavedMusic(&tg.TLUsersGetSavedMusic{Id: tg.MakeTLInputUserSelf(&tg.TLInputUserSelf{})})
	if err != nil {
		t.Fatalf("UsersGetSavedMusic returned error: %v", err)
	}
	saved, ok := got.Clazz.(*tg.TLUsersSavedMusic)
	if !ok || saved.Count != 0 || len(saved.Documents) != 0 {
		t.Fatalf("saved music response = %#v", got)
	}
}

func TestUsersGetSavedMusicByID(t *testing.T) {
	core := newUserChannelProfilesCoreForTest(nil, &fakeMediaClient{
		getDocumentList: func(_ context.Context, in *mediapb.TLMediaGetDocumentList) (*mediapb.VectorDocument, error) {
			if len(in.IdList) != 2 || in.IdList[0] != 10 || in.IdList[1] != 20 {
				t.Fatalf("document list request = %+v", in)
			}
			return &mediapb.VectorDocument{Datas: []tg.DocumentClazz{documentFixture(10), documentFixture(20)}}, nil
		},
	}, 1001)
	got, err := core.UsersGetSavedMusicByID(&tg.TLUsersGetSavedMusicByID{Documents: []tg.InputDocumentClazz{
		tg.MakeTLInputDocument(&tg.TLInputDocument{Id: 10}),
		tg.MakeTLInputDocument(&tg.TLInputDocument{Id: 20}),
	}})
	if err != nil {
		t.Fatalf("UsersGetSavedMusicByID returned error: %v", err)
	}
	saved, ok := got.Clazz.(*tg.TLUsersSavedMusic)
	if !ok || saved.Count != 2 || len(saved.Documents) != 2 {
		t.Fatalf("saved music by id response = %#v", got)
	}
}
