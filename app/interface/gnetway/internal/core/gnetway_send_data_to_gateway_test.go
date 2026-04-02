package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/gnetway"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestGnetwaySendDataToGatewayPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	okResult, err := c.GnetwaySendDataToGateway(&gnetway.TLGnetwaySendDataToGateway{
		AuthKeyId: 1,
		SessionId: 2,
		Payload:   []byte{1, 2, 3},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(okResult) {
		t.Fatalf("expected boolTrue placeholder, got %#v", okResult)
	}

	badResult, err := c.GnetwaySendDataToGateway(&gnetway.TLGnetwaySendDataToGateway{
		AuthKeyId: 1,
		SessionId: 2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if tg.FromBool(badResult) {
		t.Fatalf("expected boolFalse when payload is empty, got %#v", badResult)
	}
}
