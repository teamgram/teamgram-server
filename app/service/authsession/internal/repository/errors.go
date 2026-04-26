package repository

import (
	"errors"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
)

func wrapStorage(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %w", authsession.ErrAuthSessionStorage, err)
}

func isNotFound(err error) bool {
	return errors.Is(err, model.ErrNotFound) || errors.Is(err, sqlx.ErrNotFound) || errors.Is(err, sqlc.ErrNotFound)
}
