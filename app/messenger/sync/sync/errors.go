package sync

import "errors"

var (
	ErrSyncInvalidArgument         = errors.New("sync: invalid argument")
	ErrSyncPermissionDenied        = errors.New("sync: permission denied")
	ErrSyncAddressRejected         = errors.New("sync: gateway address rejected")
	ErrSyncPresenceFailure         = errors.New("sync: presence failure")
	ErrSyncGatewayFailure          = errors.New("sync: gateway failure")
	ErrSyncTargetSessionMissing    = errors.New("sync: target session missing")
	ErrSyncRpcResultDeliveryFailed = errors.New("sync: rpc result delivery failed")
	ErrSyncMethodNotImplemented    = errors.New("sync: method not implemented")
)
