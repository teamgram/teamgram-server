package chat

import (
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/env2"
)

func TestGenChatInviteHashReturnsUsableHash(t *testing.T) {
	hash := GenChatInviteHash()

	if hash == "" {
		t.Fatal("GenChatInviteHash returned empty string")
	}
	if strings.ContainsAny(hash, "/+ \t\r\n") {
		t.Fatalf("GenChatInviteHash returned non-hash characters: %q", hash)
	}
	if hash == GenChatInviteHash() {
		t.Fatalf("GenChatInviteHash returned duplicate hash %q", hash)
	}
}

func TestNormalizeInviteHash(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"raw hash", "abcDEF123", "abcDEF123"},
		{"plus hash", "+abcDEF123", "abcDEF123"},
		{"t.me plus link", "https://t.me/+abcDEF123", "abcDEF123"},
		{"telegram.me plus link", "https://telegram.me/+abcDEF123", "abcDEF123"},
		{"url final segment", "https://example.test/invite/abcDEF123", "abcDEF123"},
		{"url final plus segment with query", "https://example.test/invite/+abcDEF123?start=1", "abcDEF123"},
		{"trailing slash", "https://example.test/invite/+abcDEF123/", "abcDEF123"},
		{"spaces", "  +abcDEF123  ", "abcDEF123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeInviteHash(tt.in); got != tt.want {
				t.Fatalf("NormalizeInviteHash(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestBuildInviteLinkUsesTDotMe(t *testing.T) {
	old := env2.TDotMe
	t.Cleanup(func() { env2.TDotMe = old })
	env2.TDotMe = "telegram.example"

	if got, want := BuildInviteLink("abcDEF123"), "https://telegram.example/+abcDEF123"; got != want {
		t.Fatalf("BuildInviteLink = %q, want %q", got, want)
	}
}
