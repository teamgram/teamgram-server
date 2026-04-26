package repository

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
)

func wrapStorageError(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", status.ErrStatusStorage, op, err)
}
