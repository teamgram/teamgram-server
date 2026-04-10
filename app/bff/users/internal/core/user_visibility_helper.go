package core

import "github.com/teamgram/proto/mtproto"

func patchBotUsernameFromImmutable(dst *mtproto.User, src *mtproto.ImmutableUser) {
	if dst == nil || src == nil || !src.IsBot() {
		return
	}

	if len(src.Usernames()) > 1 {
		dst.Username = nil
		dst.Usernames = src.Usernames()
		return
	}

	if src.Username() != "" {
		dst.Username = mtproto.MakeFlagsString(src.Username())
		dst.Usernames = nil
	}
}
