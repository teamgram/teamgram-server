package repository

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
)

func TestWrapStorageErrorReturnsStatusStorage(t *testing.T) {
	cause := errors.New("redis down")

	err := wrapStorageError("get user online sessions", cause)
	if !errors.Is(err, status.ErrStatusStorage) {
		t.Fatalf("wrapStorageError() error = %v, want ErrStatusStorage", err)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("wrapStorageError() error = %v, want original cause", err)
	}
}
