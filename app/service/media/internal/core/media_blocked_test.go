package core

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestBlockedMediaRPCsReturnEnterpriseBlocked(t *testing.T) {
	c := &MediaCore{}

	tests := []struct {
		name string
		call func() error
	}{
		{name: "upload encrypted", call: func() error {
			_, err := c.MediaUploadEncryptedFile(&media.TLMediaUploadEncryptedFile{})
			return err
		}},
		{name: "get encrypted", call: func() error {
			_, err := c.MediaGetEncryptedFile(&media.TLMediaGetEncryptedFile{})
			return err
		}},
		{name: "wallpaper", call: func() error {
			_, err := c.MediaUploadWallPaperFile(&media.TLMediaUploadWallPaperFile{})
			return err
		}},
		{name: "theme", call: func() error {
			_, err := c.MediaUploadThemeFile(&media.TLMediaUploadThemeFile{})
			return err
		}},
		{name: "sticker", call: func() error {
			_, err := c.MediaUploadStickerFile(&media.TLMediaUploadStickerFile{})
			return err
		}},
		{name: "ringtone", call: func() error {
			_, err := c.MediaUploadRingtoneFile(&media.TLMediaUploadRingtoneFile{})
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.call(); !errors.Is(err, tg.ErrEnterpriseIsBlocked) {
				t.Fatalf("error = %v, want ErrEnterpriseIsBlocked", err)
			}
		})
	}
}
