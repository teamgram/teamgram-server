package iface_test

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestFindRPCContextTupleByClazzIDForGeneratedTGMethod(t *testing.T) {
	tuple := iface.FindRPCContextTupleByClazzID(tg.ClazzID_auth_sendCode)
	if tuple == nil {
		t.Fatal("FindRPCContextTupleByClazzID() = nil, want tuple")
	}
	if tuple.Method != "/tg.RPCAuthorization/auth.sendCode" {
		t.Fatalf("tuple.Method = %q, want %q", tuple.Method, "/tg.RPCAuthorization/auth.sendCode")
	}
	if got := tuple.ServiceName(); got != "RPCAuthorization" {
		t.Fatalf("tuple.ServiceName() = %q, want %q", got, "RPCAuthorization")
	}
}
