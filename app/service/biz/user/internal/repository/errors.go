package repository

import (
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
)

func isNotFound(err error) bool {
	var nf *model.NotFoundError
	return errors.As(err, &nf)
}

func wrapStorage(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", userpb.ErrUserStorage, op, err)
}
