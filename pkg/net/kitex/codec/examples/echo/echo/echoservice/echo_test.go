package echoservice

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/echo"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
)

func TestEchoArgsDecodePropagatesDecodeError(t *testing.T) {
	x := bin.NewEncoder()
	x.PutClazzID(echo.ClazzID_echo_echo)
	data := x.Clone()
	x.Release()

	var args EchoArgs
	err := args.Decode(bin.NewDecoder(data))
	if err == nil {
		t.Fatal("expected decode error for truncated request payload")
	}
	if args.Req != nil {
		t.Fatal("expected request to stay unset on decode failure")
	}
}
