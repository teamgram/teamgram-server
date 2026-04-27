package repository

import (
	"testing"

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
