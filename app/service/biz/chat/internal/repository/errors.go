package repository

import (
	"errors"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
)

func isNotFound(err error) bool {
	return errors.Is(err, model.ErrNotFound)
}

func wrapStorage(op string, err error) error {
	return chatpb.WrapChatStorage(op, err)
}
