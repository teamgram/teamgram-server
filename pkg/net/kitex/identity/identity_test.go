package identity

import (
	"context"
	"testing"
)

func TestCallerServiceRoundTrip(t *testing.T) {
	ctx := WithCallerService(context.Background(), "gateway")
	got, ok := CallerService(ctx)
	if !ok {
		t.Fatal("CallerService() ok = false, want true")
	}
	if got != "gateway" {
		t.Fatalf("CallerService() = %q, want gateway", got)
	}
}

func TestCallerServiceEmptyIsAbsent(t *testing.T) {
	ctx := WithCallerService(context.Background(), "")
	_, ok := CallerService(ctx)
	if ok {
		t.Fatal("CallerService() ok = true, want false")
	}
}
