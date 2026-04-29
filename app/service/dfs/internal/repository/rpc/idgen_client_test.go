package rpc

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	idgenclient "github.com/teamgram/teamgram-server/v2/app/service/idgen/client"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeIDGenClient struct {
	id    int64
	calls int
	err   error
	idgenclient.IdgenClient
}

func (f *fakeIDGenClient) IdgenNextId(context.Context, *idgen.TLIdgenNextId) (*tg.Int64, error) {
	f.calls++
	if f.err != nil {
		return nil, f.err
	}
	return tg.MakeInt64(f.id), nil
}

func TestIDGeneratorMethodsCallNextID(t *testing.T) {
	fake := &fakeIDGenClient{id: 42}
	gen := NewIDGenClient(fake)

	for _, call := range []struct {
		name string
		fn   func(context.Context) (int64, error)
	}{
		{name: "photo", fn: gen.NextPhotoID},
		{name: "document", fn: gen.NextDocumentID},
		{name: "encrypted", fn: gen.NextEncryptedFileID},
	} {
		got, err := call.fn(context.Background())
		if err != nil {
			t.Fatalf("%s next id error = %v", call.name, err)
		}
		if got != 42 {
			t.Fatalf("%s next id = %d, want 42", call.name, got)
		}
	}
	if fake.calls != 3 {
		t.Fatalf("IdgenNextId calls = %d, want 3", fake.calls)
	}
}

func TestIDGeneratorWrapsDownstreamErrors(t *testing.T) {
	cause := errors.New("idgen down")
	gen := NewIDGenClient(&fakeIDGenClient{err: cause})
	_, err := gen.NextPhotoID(context.Background())
	if !errors.Is(err, dfs.ErrDfsDownstream) {
		t.Fatalf("error = %v, want ErrDfsDownstream", err)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("error = %v, want original cause", err)
	}
}
