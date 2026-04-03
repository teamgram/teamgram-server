package dao

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

func TestWatchGatewayIgnoresSubscriberCreationFailure(t *testing.T) {
	orig := newSubscriber
	t.Cleanup(func() {
		newSubscriber = orig
	})

	newSubscriber = func(_ []string, _ string) (gatewaySubscriber, error) {
		return nil, errors.New("boom")
	}

	d := &Dao{
		eGateServers: map[string]*Gateway{
			"existing": {},
		},
	}

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("watchGateway should not panic when subscriber creation fails: %v", r)
		}
	}()

	d.watchGateway(kitex.RpcClientConf{})

	if len(d.eGateServers) != 1 {
		t.Fatalf("expected existing gateways to stay intact, got %#v", d.eGateServers)
	}
}
