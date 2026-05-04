package metadata

import (
	"context"
	"testing"
)

func TestRpcMetadataCarriesGatewayRoute(t *testing.T) {
	ctx := context.Background()
	want := &RpcMetadata{
		ServerId:          "gateway-dev-1",
		GatewayRpcAddr:    "127.0.0.1:20110",
		GatewayGeneration: "generation-1",
		AuthId:            1001,
		SessionId:         2002,
	}

	outgoing, err := RpcMetadataToOutgoing(ctx, want)
	if err != nil {
		t.Fatalf("RpcMetadataToOutgoing() error = %v", err)
	}
	got := RpcMetadataFromIncoming(outgoing)
	if got == nil {
		t.Fatal("RpcMetadataFromIncoming() returned nil")
	}
	if got.GatewayRpcAddr != want.GatewayRpcAddr || got.GatewayGeneration != want.GatewayGeneration {
		t.Fatalf("gateway route metadata = (%q, %q), want (%q, %q)",
			got.GatewayRpcAddr, got.GatewayGeneration, want.GatewayRpcAddr, want.GatewayGeneration)
	}
}
