package userupdates

import (
	"errors"
	"fmt"
	"testing"
)

func TestSemanticErrorsSupportErrorsIs(t *testing.T) {
	cases := []error{
		ErrNotOwner,
		ErrOwnerFenceFailed,
		ErrOperationPayloadConflict,
		ErrPtsContinuityViolation,
		ErrOperationTerminal,
		ErrUserupdatesStorage,
	}
	for _, target := range cases {
		wrapped := fmt.Errorf("wrapped: %w", target)
		if !errors.Is(wrapped, target) {
			t.Fatalf("errors.Is did not match %v", target)
		}
	}
}
