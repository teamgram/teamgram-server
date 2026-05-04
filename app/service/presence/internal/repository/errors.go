package repository

import (
	"fmt"

	presencepb "github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
)

func wrapStorageError(op string, err error) error {
	return fmt.Errorf("%w: %s: %w", presencepb.ErrPresenceStorage, op, err)
}
