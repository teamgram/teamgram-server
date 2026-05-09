package core

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

const maxCaptionLen = 4096

func checkCaption(caption string) error {
	if utf16CodeUnitLen(caption) > maxCaptionLen {
		return tg.ErrMediaCaptionTooLong
	}
	return nil
}
