package dialog

import (
	"errors"
	"testing"
)

func TestDomainErrorsAreStableSentinels(t *testing.T) {
	errs := []error{
		ErrDialogNotFound,
		ErrDialogStorage,
		ErrDialogInvalid,
	}

	for _, err := range errs {
		if err == nil {
			t.Fatalf("domain error is nil")
		}
	}
}

func TestWrapDialogStorage(t *testing.T) {
	cause := errors.New("mysql unavailable")
	err := WrapDialogStorage("list dialogs", cause)

	if err == nil {
		t.Fatal("WrapDialogStorage returned nil")
	}
	if !errors.Is(err, ErrDialogStorage) {
		t.Fatalf("errors.Is(%v, ErrDialogStorage) = false", err)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("errors.Is(%v, cause) = false", err)
	}
	if got := WrapDialogStorage("list dialogs", nil); got != nil {
		t.Fatalf("WrapDialogStorage nil input = %v, want nil", got)
	}
}
