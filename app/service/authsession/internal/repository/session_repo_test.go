package repository

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
)

func TestSessionParamsDefaultToNull(t *testing.T) {
	row := authsFromClientSession(2002, authsession.MakeTLClientSession(&authsession.TLClientSession{
		AuthKeyId: 1001,
		Params:    "",
	}))
	if row.Params != "null" {
		t.Fatalf("Params = %q, want null", row.Params)
	}
}

func TestSetInitConnectionParamsDefaultToNull(t *testing.T) {
	row := authsFromInitConnection(2002, &authsession.TLAuthsessionSetInitConnection{
		AuthKeyId: 1001,
		Params:    "",
	})
	if row.Params != "null" {
		t.Fatalf("Params = %q, want null", row.Params)
	}
}

func TestSessionConversionDoesNotMutateInput(t *testing.T) {
	in := authsession.MakeTLClientSession(&authsession.TLClientSession{
		AuthKeyId: 1001,
	})
	row := authsFromClientSession(2002, in)
	if in.AuthKeyId != 1001 {
		t.Fatalf("input AuthKeyId mutated: got %d, want 1001", in.AuthKeyId)
	}
	if row.AuthKeyId != 2002 {
		t.Fatalf("row AuthKeyId = %d, want perm id 2002", row.AuthKeyId)
	}
}

func TestInitConnectionConversionDoesNotMutateInput(t *testing.T) {
	in := &authsession.TLAuthsessionSetInitConnection{AuthKeyId: 1001}
	row := authsFromInitConnection(2002, in)
	if in.AuthKeyId != 1001 {
		t.Fatalf("input AuthKeyId mutated: got %d, want 1001", in.AuthKeyId)
	}
	if row.AuthKeyId != 2002 {
		t.Fatalf("row AuthKeyId = %d, want perm id 2002", row.AuthKeyId)
	}
}
