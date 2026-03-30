package sess

import (
	"testing"

	"github.com/teamgram/proto/mtproto"
	rpcmetadata "github.com/teamgram/proto/mtproto/rpc/metadata"
)

func TestTakeoutRequiredError(t *testing.T) {
	err := takeoutRequiredError()
	if err.GetErrorMessage() != "TAKEOUT_REQUIRED" {
		t.Fatalf("unexpected error message: %s", err.GetErrorMessage())
	}
}

func TestTakeoutMetadataShape(t *testing.T) {
	guard := newTakeoutGuard()
	md, err := guard.Validate(1001)
	if err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if md.Id != 1001 {
		t.Fatalf("unexpected takeout metadata: %#v", md)
	}
}

func TestTakeoutGuardRejectsZeroId(t *testing.T) {
	guard := newTakeoutGuard()
	md, err := guard.Validate(0)
	if err == nil {
		t.Fatalf("Validate() error = nil, md = %#v", md)
	}
	if md != nil {
		t.Fatalf("Validate() md = %#v, want nil", md)
	}
}

func TestTakeoutMetadataStructShape(t *testing.T) {
	md := &rpcmetadata.Takeout{Id: 1001}
	if md.Id != 1001 {
		t.Fatalf("unexpected takeout metadata: %#v", md)
	}
}

func TestTakeoutGuardUnwrapsMessagesRangeIntoMetadata(t *testing.T) {
	guard := newTakeoutGuard()
	queryBytes := []byte{0x6b, 0x18, 0xf9, 0xc4}

	query, md, err := guard.ValidateWrappedQuery(1001, &mtproto.TLInvokeWithMessagesRange{
		Range: &mtproto.MessageRange{MinId: 5, MaxId: 10},
		Query: queryBytes,
	})
	if err != nil {
		t.Fatalf("ValidateWrappedQuery() error = %v", err)
	}
	if _, ok := query.(*mtproto.TLHelpGetConfig); !ok {
		t.Fatalf("ValidateWrappedQuery() query = %T", query)
	}
	if md == nil || md.Id != 1001 || md.Range == nil || md.Range.MinId != 5 || md.Range.MaxId != 10 {
		t.Fatalf("ValidateWrappedQuery() metadata = %#v", md)
	}
}

func TestNewTakeoutMetadataWithRangeShape(t *testing.T) {
	md := newTakeoutMetadata(0, &mtproto.MessageRange{MinId: 7, MaxId: 9})
	if md == nil || md.Id != 0 || md.Range == nil || md.Range.MinId != 7 || md.Range.MaxId != 9 {
		t.Fatalf("newTakeoutMetadata() = %#v", md)
	}
}
