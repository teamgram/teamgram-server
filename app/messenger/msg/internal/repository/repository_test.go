package repository

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
)

func TestStorageErrorPreservesCause(t *testing.T) {
	cause := errors.New("mysql unavailable")

	err := storageError("query", cause)

	if !errors.Is(err, msg.ErrMsgStorage) {
		t.Fatalf("storageError() does not wrap ErrMsgStorage: %v", err)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("storageError() does not preserve cause: %v", err)
	}
}
