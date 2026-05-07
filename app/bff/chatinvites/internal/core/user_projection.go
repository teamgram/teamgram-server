package core

import (
	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *ChatInvitesCore) fetchUserClazzes(ids []int64, selfID int64) ([]tg.UserClazz, error) {
	ids = uniqueInt64s(ids)
	if len(ids) == 0 {
		return []tg.UserClazz{}, nil
	}

	users, err := userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, selfID, ids, userprojection.MissingStoredReference)
	if err != nil {
		c.Logger.Errorf("chatinvites.fetchUserClazzes - user.getUserProjectionBundle failed: ids: %v, self_id: %d, err: %v", ids, selfID, err)
		return nil, tg.ErrInternalServerError
	}
	return users, nil
}

func uniqueInt64s(ids []int64) []int64 {
	seen := make(map[int64]struct{}, len(ids))
	out := make([]int64, 0, len(ids))
	for _, id := range ids {
		if id == 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		out = append(out, id)
	}
	return out
}

func nonEmptyStringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
