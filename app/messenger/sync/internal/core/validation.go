package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/repository"
	syncpb "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func validateUserID(method string, userID int64) error {
	if userID <= 0 {
		return fmt.Errorf("%w: %s invalid user_id %d", syncpb.ErrSyncInvalidArgument, method, userID)
	}
	return nil
}

func validatePositiveID(method, field string, value int64) error {
	if value <= 0 {
		return fmt.Errorf("%w: %s invalid %s %d", syncpb.ErrSyncInvalidArgument, method, field, value)
	}
	return nil
}

func validateNonZeroID(method, field string, value int64) error {
	if value == 0 {
		return fmt.Errorf("%w: %s invalid %s %d", syncpb.ErrSyncInvalidArgument, method, field, value)
	}
	return nil
}

func validateUpdates(method string, updates tg.UpdatesClazz) error {
	if updates == nil {
		return fmt.Errorf("%w: %s updates is nil", syncpb.ErrSyncInvalidArgument, method)
	}
	return nil
}

func validatePermKeyVector(method, field string, values []int64, max int) error {
	if len(values) > max {
		return fmt.Errorf("%w: %s %s length %d exceeds %d", syncpb.ErrSyncInvalidArgument, method, field, len(values), max)
	}
	for i, value := range values {
		if value == 0 {
			return fmt.Errorf("%w: %s %s[%d] invalid %d", syncpb.ErrSyncInvalidArgument, method, field, i, value)
		}
	}
	return nil
}

func validatePushUpdatesIfNot(method string, includes, excludes []int64) error {
	if err := validatePermKeyVector(method, "includes", includes, repository.MaxIncludePermKeys); err != nil {
		return err
	}
	return validatePermKeyVector(method, "excludes", excludes, repository.MaxExcludePermKeys)
}
