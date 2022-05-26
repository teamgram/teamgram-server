package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMedia(t *testing.T) {
	t.Run("#ObtainEncryptionKey", func(t *testing.T) {
		t.Run("Should get nil if encryptionKey is not set", func(t *testing.T) {
			mediaFile := Mediafile{}

			require.Nil(t, mediaFile.ObtainEncryptionKey())
		})

		t.Run("Should return file.keyinfo if it's set", func(t *testing.T) {
			mediaFile := Mediafile{encryptionKey: "file.keyinfo"}
			require.Equal(t, []string{"-hls_key_info_file", "file.keyinfo"}, mediaFile.ObtainEncryptionKey())
		})
	})
}
