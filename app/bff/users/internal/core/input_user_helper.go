package core

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

func userIDFromInputUser(selfID int64, in tg.InputUserClazz) (int64, error) {
	switch user := in.(type) {
	case *tg.TLInputUserSelf:
		if selfID <= 0 {
			return 0, tg.ErrUserIdInvalid
		}
		return selfID, nil
	case *tg.TLInputUser:
		if user.UserId <= 0 {
			return 0, tg.ErrUserIdInvalid
		}
		return user.UserId, nil
	default:
		return 0, tg.ErrUserIdInvalid
	}
}

func requireSelfID(c *UsersCore) (int64, error) {
	if c == nil || c.MD == nil || c.MD.UserId <= 0 {
		return 0, tg.ErrUserIdInvalid
	}
	return c.MD.UserId, nil
}
