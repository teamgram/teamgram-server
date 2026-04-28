package echo

import (
	"strings"
	"testing"
)

func TestTLEchoEchoStringIncludesPredicateName(t *testing.T) {
	got := (&TLEchoEcho{Message: "hi"}).String()
	if !strings.Contains(got, "\"@type\":\"echo_echo\"") {
		t.Fatalf("expected predicate type in string output, got %s", got)
	}
}
