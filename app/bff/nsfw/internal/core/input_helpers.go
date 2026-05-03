package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func requireSelfID(c *NsfwCore) (int64, error) {
	if c == nil || c.MD == nil || c.MD.UserId <= 0 {
		return 0, tg.ErrUserIdInvalid
	}
	return c.MD.UserId, nil
}

func requireUserClient(c *NsfwCore) error {
	if c == nil || c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.UserClient == nil {
		return fmt.Errorf("nsfw: user client is nil")
	}
	return nil
}
