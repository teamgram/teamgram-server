package xerr

import (
	"errors"
	"strings"
	"testing"
)

func TestNewDoesNotUseStringEqualityAsSentinelMatching(t *testing.T) {
	left := New("same text")
	right := New("same text")

	if errors.Is(left, right) {
		t.Fatalf("expected different New(\"same text\") values to not match via errors.Is")
	}
}

func TestWrapMsgPreservesCause(t *testing.T) {
	base := errors.New("base error")

	got := WrapMsg(base, "save user", "uid", 100)
	if got == nil {
		t.Fatal("expected wrapped error")
	}
	if !errors.Is(got, base) {
		t.Fatalf("expected wrapped error to preserve cause")
	}
	if !strings.Contains(got.Error(), "base error") {
		t.Fatalf("expected wrapped error text to contain base error, got %q", got.Error())
	}
	if !strings.Contains(got.Error(), "save user, uid=100") {
		t.Fatalf("expected wrapped error text to contain context, got %q", got.Error())
	}
}

func TestWrapDoesNotLeakStackFramesInErrorText(t *testing.T) {
	got := Wrap(errors.New("boom"))
	if got == nil {
		t.Fatal("expected wrapped error")
	}
	if strings.Contains(got.Error(), ".go:") {
		t.Fatalf("expected Error() to not include file paths, got %q", got.Error())
	}
	if strings.Contains(got.Error(), "->") {
		t.Fatalf("expected Error() to not include rendered stack frames, got %q", got.Error())
	}
}

func TestErrPanicMsgReturnsWrappedErrorWithoutCodePrefix(t *testing.T) {
	got := ErrPanicMsg("panic-value", 500, "panic error", 1)
	if got == nil {
		t.Fatal("expected panic error")
	}
	if strings.Contains(got.Error(), "500 ") {
		t.Fatalf("expected panic error text to not include transport code prefix, got %q", got.Error())
	}
	if !strings.Contains(got.Error(), "panic error") || !strings.Contains(got.Error(), "panic-value") {
		t.Fatalf("expected panic error text to include message and recovered value, got %q", got.Error())
	}
}
