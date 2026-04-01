package echo

import (
	"strings"
	"testing"
)

func TestTLEchoEchoStringIncludesPredicateName(t *testing.T) {
	got := (&TLEchoEcho{Message: "hi"}).String()
	if !strings.Contains(got, "\"_name\":\"echo_echo\"") {
		t.Fatalf("expected predicate name in string output, got %s", got)
	}
}
