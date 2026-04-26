package repository

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
)

func TestSessionParamsDefaultToNull(t *testing.T) {
	row := authsFromClientSession(authsession.MakeTLClientSession(&authsession.TLClientSession{
		AuthKeyId: 1001,
		Params:    "",
	}))
	if row.Params != "null" {
		t.Fatalf("Params = %q, want null", row.Params)
	}
}

func TestSetInitConnectionParamsDefaultToNull(t *testing.T) {
	row := authsFromInitConnection(&authsession.TLAuthsessionSetInitConnection{
		AuthKeyId: 1001,
		Params:    "",
	})
	if row.Params != "null" {
		t.Fatalf("Params = %q, want null", row.Params)
	}
}
