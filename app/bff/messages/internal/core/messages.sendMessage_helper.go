package core

import (
	"context"
	"errors"
	"strings"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface/ecode"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func utf16CodeUnitLen(s string) int {
	n := 0
	for _, r := range s {
		if r >= 0x10000 {
			n += 2
		} else {
			n++
		}
	}
	return n
}

const maxMessageLen = 4096

func checkMessage(message string) error {
	if utf16CodeUnitLen(strings.TrimSpace(message)) == 0 {
		return tg.ErrMessageEmpty
	}
	if utf16CodeUnitLen(message) > maxMessageLen {
		return tg.ErrMessageTooLong
	}
	return nil
}

func checkUnsupportedFields(in *tg.TLMessagesSendMessage) error {
	if in.NoWebpage {
		return tg.ErrInputRequestInvalid
	}
	if in.Silent {
		return tg.ErrInputRequestInvalid
	}
	if in.Background {
		return tg.ErrInputRequestInvalid
	}
	if in.Noforwards {
		return tg.ErrInputRequestInvalid
	}
	if len(in.Entities) > 0 {
		return tg.ErrInputRequestInvalid
	}
	if in.ReplyTo != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.ReplyMarkup != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.SendAs != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.QuickReplyShortcut != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.SuggestedPost != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.ScheduleDate != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.ScheduleRepeatPeriod != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.Effect != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.AllowPaidStars != nil {
		return tg.ErrInputRequestInvalid
	}
	return nil
}

func mapMsgSendError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, context.DeadlineExceeded), errors.Is(err, context.Canceled):
		return tg.ErrTimeout
	case errors.Is(err, msg.ErrRandomIdConflict):
		return tg.ErrRandomIdDuplicate
	case errors.Is(err, msg.ErrReceiverBackpressure),
		errors.Is(err, msg.ErrSenderSyncFailed),
		errors.Is(err, msg.ErrMsgStorage),
		errors.Is(err, msg.ErrSendStateConflict):
		return tg.ErrInternalServerError
	default:
		var codeErr ecode.CodeError
		if errors.As(err, &codeErr) {
			return err
		}
		return tg.ErrInternalServerError
	}
}
