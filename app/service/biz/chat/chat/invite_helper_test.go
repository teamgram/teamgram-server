package chat

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/env2"
	mtproto "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestGenChatInviteHashReturnsUsableHash(t *testing.T) {
	hash := GenChatInviteHash()

	if !IsChatInviteHash(hash) {
		t.Fatalf("GenChatInviteHash returned non-chat invite hash: %q", hash)
	}
	if got := GetChatTypeByInviteHash(hash); got != mtproto.PEER_CHAT {
		t.Fatalf("GetChatTypeByInviteHash(%q) = %d, want %d", hash, got, mtproto.PEER_CHAT)
	}

	link := BuildInviteLink(hash)
	if got := NormalizeInviteHash(link); got != hash {
		t.Fatalf("NormalizeInviteHash(BuildInviteLink(hash)) = %q, want %q", got, hash)
	}
}

func TestGenChannelInviteHashReturnsUsableHash(t *testing.T) {
	hash := GenChannelInviteHash()

	if !IsChannelInviteHash(hash) {
		t.Fatalf("GenChannelInviteHash returned non-channel invite hash: %q", hash)
	}
	if got := GetChatTypeByInviteHash(hash); got != mtproto.PEER_CHANNEL {
		t.Fatalf("GetChatTypeByInviteHash(%q) = %d, want %d", hash, got, mtproto.PEER_CHANNEL)
	}
}

func TestInviteHashTypeHelpers(t *testing.T) {
	tests := []struct {
		name        string
		hash        string
		wantType    int
		wantChat    bool
		wantChannel bool
	}{
		{"chat digit lower bound", "0bcDEF123456789abcde", mtproto.PEER_CHAT, true, false},
		{"chat digit upper bound", "4bcDEF123456789abcde", mtproto.PEER_CHAT, true, false},
		{"chat lowercase", "abcDEF123456789abcd0", mtproto.PEER_CHAT, true, false},
		{"channel digit lower bound", "5bcDEF123456789abcde", mtproto.PEER_CHANNEL, false, true},
		{"channel digit upper bound", "9bcDEF123456789abcde", mtproto.PEER_CHANNEL, false, true},
		{"channel uppercase", "AbcDEF123456789abcde", mtproto.PEER_CHANNEL, false, true},
		{"bad length", "abcDEF123", mtproto.PEER_UNKNOWN, false, false},
		{"bad first byte", "_bcDEF123456789abcde", mtproto.PEER_UNKNOWN, false, false},
		{"bad body byte", "0bcDEF123456789abcd_", mtproto.PEER_UNKNOWN, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetChatTypeByInviteHash(tt.hash); got != tt.wantType {
				t.Fatalf("GetChatTypeByInviteHash(%q) = %d, want %d", tt.hash, got, tt.wantType)
			}
			if got := IsChatInviteHash(tt.hash); got != tt.wantChat {
				t.Fatalf("IsChatInviteHash(%q) = %t, want %t", tt.hash, got, tt.wantChat)
			}
			if got := IsChannelInviteHash(tt.hash); got != tt.wantChannel {
				t.Fatalf("IsChannelInviteHash(%q) = %t, want %t", tt.hash, got, tt.wantChannel)
			}
		})
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
