package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/presence/internal/metrics"
	presencepb "github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/identity"
)

func (c *PresenceCore) requireCaller(method string, allowed []string) error {
	_, err := c.authorizedCaller(method, allowed)
	return err
}

func (c *PresenceCore) authorizedCaller(method string, allowed []string) (string, error) {
	caller, ok := identity.CallerService(c.ctx)
	if !ok {
		caller = ""
	}
	if !c.svcCtx.Config.RequireCallerIdentity {
		if callerServiceAllowed(caller, allowed) {
			return caller, nil
		}
		return "local", nil
	}
	for _, service := range allowed {
		if caller == service {
			return caller, nil
		}
	}
	metrics.PermissionDenied(method, permissionDeniedCallerLabel(caller))
	return "", fmt.Errorf("%w: %s caller %q", presencepb.ErrPresencePermissionDenied, method, caller)
}

func callerServiceAllowed(caller string, allowed []string) bool {
	for _, service := range allowed {
		if caller == service {
			return true
		}
	}
	return false
}

func permissionDeniedCallerLabel(caller string) string {
	if caller == "" {
		return "missing"
	}
	return "unauthorized"
}

func (c *PresenceCore) requireQuota(method, caller string, limit int) error {
	if c.svcCtx.AllowPresenceCall(method, caller, limit) {
		return nil
	}
	return fmt.Errorf("%w: %s caller %q", presencepb.ErrPresenceQuotaExceeded, method, caller)
}

func allowedQueryCallers(syncCallers, adminCallers, debugCallers []string) []string {
	allowed := make([]string, 0, len(syncCallers)+len(adminCallers)+len(debugCallers))
	allowed = append(allowed, syncCallers...)
	allowed = append(allowed, adminCallers...)
	allowed = append(allowed, debugCallers...)
	return allowed
}

func allowedAdminDebugCallers(adminCallers, debugCallers []string) []string {
	allowed := make([]string, 0, len(adminCallers)+len(debugCallers))
	allowed = append(allowed, adminCallers...)
	allowed = append(allowed, debugCallers...)
	return allowed
}

func validateOnlineSession(method string, session *presencepb.OnlineSession) error {
	if session == nil {
		return fmt.Errorf("%w: %s session is nil", presencepb.ErrPresenceInvalidArgument, method)
	}
	if session.UserId <= 0 {
		return fmt.Errorf("%w: %s invalid user_id %d", presencepb.ErrPresenceInvalidArgument, method, session.UserId)
	}
	if session.PermAuthKeyId <= 0 {
		return fmt.Errorf("%w: %s invalid perm_auth_key_id %d", presencepb.ErrPresenceInvalidArgument, method, session.PermAuthKeyId)
	}
	if session.AuthKeyId <= 0 {
		return fmt.Errorf("%w: %s invalid auth_key_id %d", presencepb.ErrPresenceInvalidArgument, method, session.AuthKeyId)
	}
	if session.SessionId == 0 {
		return fmt.Errorf("%w: %s invalid session_id", presencepb.ErrPresenceInvalidArgument, method)
	}
	if session.GatewayId == "" {
		return fmt.Errorf("%w: %s gateway_id is empty", presencepb.ErrPresenceInvalidArgument, method)
	}
	if session.GatewayGeneration == "" {
		return fmt.Errorf("%w: %s gateway_generation is empty", presencepb.ErrPresenceInvalidArgument, method)
	}
	if session.GatewayRpcAddr == "" {
		return fmt.Errorf("%w: %s gateway_rpc_addr is empty", presencepb.ErrPresenceInvalidArgument, method)
	}
	return nil
}
