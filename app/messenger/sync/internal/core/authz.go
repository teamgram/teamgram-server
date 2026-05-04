package core

import (
	"fmt"

	syncpb "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/identity"
)

func (c *SyncCore) requireCaller(method string) error {
	if !c.svcCtx.Config.RequireCallerIdentity {
		return nil
	}
	caller, ok := identity.CallerService(c.ctx)
	if !ok {
		caller = ""
	}
	for _, allowed := range c.svcCtx.Config.AllowedCallers {
		if caller == allowed {
			return nil
		}
	}
	return fmt.Errorf("%w: %s caller %q", syncpb.ErrSyncPermissionDenied, method, caller)
}
