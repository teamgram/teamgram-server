package chat

import (
	"errors"
	"testing"
)

func TestDomainErrorsAreStableSentinels(t *testing.T) {
	errs := []error{
		ErrChatNotFound,
		ErrChatMigrated,
		ErrChatDeactivated,
		ErrChatAdminRequired,
		ErrChatTitleEmpty,
		ErrChatNotModified,
		ErrCreateChatFlood,
		ErrParticipantInvalid,
		ErrInputUserDeactivated,
		ErrUserAlreadyParticipant,
		ErrUserNotParticipant,
		ErrUsersTooFew,
		ErrUsersTooMuch,
		ErrInviteHashInvalid,
		ErrInviteHashExpired,
		ErrChatLinkExists,
		ErrChatStorage,
	}

	for _, err := range errs {
		if err == nil {
			t.Fatalf("domain error is nil")
		}
	}
}

func TestMessageActionErrorsAreStableSentinels(t *testing.T) {
	for _, err := range []error{
		ErrChatWriteForbidden,
		ErrMessageActionUnsupported,
	} {
		if err == nil {
			t.Fatalf("message action sentinel is nil")
		}
	}
}

func TestCreateChatFloodErrorCarriesWaitSecondsAndMatchesSentinel(t *testing.T) {
	err := NewCreateChatFloodError(37)

	if !errors.Is(err, ErrCreateChatFlood) {
		t.Fatalf("errors.Is(%v, ErrCreateChatFlood) = false", err)
	}

	var floodErr *CreateChatFloodError
	if !errors.As(err, &floodErr) {
		t.Fatalf("errors.As(%T) = false", floodErr)
	}
	if floodErr.WaitSeconds != 37 {
		t.Fatalf("WaitSeconds = %d, want 37", floodErr.WaitSeconds)
	}
}

func TestWrapChatStorage(t *testing.T) {
	cause := errors.New("disk unavailable")
	err := WrapChatStorage("load chat", cause)

	if err == nil {
		t.Fatal("WrapChatStorage returned nil")
	}
	if !errors.Is(err, ErrChatStorage) {
		t.Fatalf("errors.Is(%v, ErrChatStorage) = false", err)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("errors.Is(%v, cause) = false", err)
	}
	if got := WrapChatStorage("load chat", nil); got != nil {
		t.Fatalf("WrapChatStorage nil input = %v, want nil", got)
	}
}
