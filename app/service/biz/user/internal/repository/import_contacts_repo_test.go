package repository

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestInputContactPhonesSkipsNilEmptyAndDuplicates(t *testing.T) {
	phones := inputContactPhones([]tg.InputContactClazz{
		nil,
		tg.MakeTLInputPhoneContact(&tg.TLInputPhoneContact{Phone: ""}).ToInputContact(),
		tg.MakeTLInputPhoneContact(&tg.TLInputPhoneContact{Phone: "10001"}).ToInputContact(),
		tg.MakeTLInputPhoneContact(&tg.TLInputPhoneContact{Phone: "10001"}).ToInputContact(),
		tg.MakeTLInputPhoneContact(&tg.TLInputPhoneContact{Phone: "10002"}).ToInputContact(),
	})

	if len(phones) != 2 || phones[0] != "10001" || phones[1] != "10002" {
		t.Fatalf("unexpected phones: %#v", phones)
	}
}

func TestUserFromModelEncodesDecodableUsernamesVector(t *testing.T) {
	user := userFromModel(&model.Users{
		Id:         1001,
		AccessHash: 2002,
		FirstName:  "Test",
	}, true, true, true, nil)

	x := bin.NewEncoder()
	defer x.Release()
	if err := user.Encode(x, 223); err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	gotClazz, err := tg.DecodeUserClazz(bin.NewDecoder(x.Bytes()))
	if err != nil {
		t.Fatalf("DecodeUserClazz() error = %v", err)
	}
	got, ok := gotClazz.(*tg.TLUser)
	if !ok {
		t.Fatalf("DecodeUserClazz() = %T, want *tg.TLUser", gotClazz)
	}
	if got == nil || got.Usernames == nil || len(got.Usernames) != 0 {
		t.Fatalf("decoded usernames = %#v, want empty non-nil vector", got)
	}
}
