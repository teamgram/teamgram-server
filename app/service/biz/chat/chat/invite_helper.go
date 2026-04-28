package chat

import (
	"net/url"
	"strings"

	"github.com/teamgram/teamgram-server/v2/pkg/env2"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
)

func GenChatInviteHash() string {
	return crypto.RandomString(16)
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
