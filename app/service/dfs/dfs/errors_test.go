package dfs

import (
	"errors"
	"testing"
)

func TestDfsSemanticErrorsAreWrappable(t *testing.T) {
	cause := errors.New("kv down")
	err := WrapDfsStorage("upload part", cause)
	if !errors.Is(err, ErrDfsStorage) {
		t.Fatalf("WrapDfsStorage errors.Is ErrDfsStorage = false: %v", err)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("WrapDfsStorage lost cause: %v", err)
	}

	downstream := errors.New("idgen down")
	err = WrapDfsDownstream("next photo id", downstream)
	if !errors.Is(err, ErrDfsDownstream) {
		t.Fatalf("WrapDfsDownstream errors.Is ErrDfsDownstream = false: %v", err)
	}
	if !errors.Is(err, downstream) {
		t.Fatalf("WrapDfsDownstream lost cause: %v", err)
	}
}
