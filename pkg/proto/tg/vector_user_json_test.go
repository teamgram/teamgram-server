package tg

import (
	"encoding/json"
	"testing"
)

func TestVectorUserMarshalJSONDoesNotRecurse(t *testing.T) {
	got, err := json.Marshal(&VectorUser{
		Datas: []UserClazz{
			MakeTLUser(&TLUser{
				Id:        1001,
				Usernames: []UsernameClazz{},
			}),
		},
	})
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	if len(got) == 0 {
		t.Fatal("Marshal() returned empty JSON")
	}
}
