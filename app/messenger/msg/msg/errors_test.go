package msg

import (
	"errors"
	"fmt"
	"testing"
)

func TestSemanticErrorsSupportErrorsIs(t *testing.T) {
	cases := []error{
		ErrRandomIdConflict,
		ErrSenderSyncFailed,
		ErrReceiverBackpressure,
		ErrMsgStorage,
	}
	for _, target := range cases {
		wrapped := fmt.Errorf("wrapped: %w", target)
		if !errors.Is(wrapped, target) {
			t.Fatalf("errors.Is did not match %v", target)
		}
	}
}
