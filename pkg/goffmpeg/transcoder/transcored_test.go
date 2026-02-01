package transcoder

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/teamgram/teamgram-server/pkg/goffmpeg/models"
)

func TestTranscoder(t *testing.T) {
	t.Run("#SetWhiteListProtocols", func(t *testing.T) {
		t.Run("Should not set -protocol_whitelist option if it isn't present", func(t *testing.T) {
			ts := Transcoder{}

			ts.SetMediaFile(&models.Mediafile{})
			require.NotEqual(t, ts.GetCommand()[0:2], []string{"-protocol_whitelist", "file,http,https,tcp,tls"})
			require.NotContains(t, ts.GetCommand(), "protocol_whitelist")
		})

		t.Run("Should set -protocol_whitelist option if it's present", func(t *testing.T) {
			ts := Transcoder{}

			ts.SetMediaFile(&models.Mediafile{})
			ts.SetWhiteListProtocols([]string{"file", "http", "https", "tcp", "tls"})

			require.Equal(t, ts.GetCommand()[0:2], []string{"-protocol_whitelist", "file,http,https,tcp,tls"})
		})
	})
}
