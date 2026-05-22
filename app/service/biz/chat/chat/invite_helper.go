package chat

import (
	"net/url"
	"strings"

	"github.com/teamgram/marmota/pkg/random2"
	"github.com/teamgram/teamgram-server/v2/pkg/env2"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func GenChatInviteHash() string {
	return randomInviteHashPrefix(isChatInviteHashPrefix) + random2.RandomAlphanumeric(19)
}

func GenChannelInviteHash() string {
	return randomInviteHashPrefix(isChannelInviteHashPrefix) + random2.RandomAlphanumeric(19)
}

func GetChatTypeByInviteHash(hash string) int {
	if !isInviteHash(hash) {
		return tg.PEER_UNKNOWN
	}

	switch {
	case isChatInviteHashPrefix(hash[0]):
		return tg.PEER_CHAT
	case isChannelInviteHashPrefix(hash[0]):
		return tg.PEER_CHANNEL
	default:
		return tg.PEER_UNKNOWN
	}
}

func IsChatInviteHash(hash string) bool {
	return isInviteHash(hash) && isChatInviteHashPrefix(hash[0])
}

func IsChannelInviteHash(hash string) bool {
	return isInviteHash(hash) && isChannelInviteHashPrefix(hash[0])
}

func randomInviteHashPrefix(match func(byte) bool) string {
	for {
		prefix := random2.RandomAlphanumeric(1)
		if len(prefix) == 1 && match(prefix[0]) {
			return prefix
		}
	}
}

func isInviteHash(hash string) bool {
	if len(hash) != 20 {
		return false
	}

	for i := 0; i < len(hash); i++ {
		if !isInviteHashChar(hash[i]) {
			return false
		}
	}
	return true
}

func isInviteHashChar(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isChatInviteHashPrefix(b byte) bool {
	return (b >= '0' && b <= '4') || (b >= 'a' && b <= 'z')
}

func isChannelInviteHashPrefix(b byte) bool {
	return (b >= '5' && b <= '9') || (b >= 'A' && b <= 'Z')
}

func NormalizeInviteHash(linkOrHash string) string {
	s := strings.TrimSpace(linkOrHash)
	if s == "" {
		return ""
	}

	if u, err := url.Parse(s); err == nil && u.Scheme != "" && u.Host != "" {
		s = u.EscapedPath()
		if unescaped, err := url.PathUnescape(s); err == nil {
			s = unescaped
		}
	}

	s = strings.TrimRight(s, "/")
	if idx := strings.LastIndex(s, "/"); idx >= 0 {
		s = s[idx+1:]
	}
	return strings.TrimPrefix(s, "+")
}

func BuildInviteLink(hash string) string {
	return "https://" + env2.TDotMe + "/+" + hash
}
