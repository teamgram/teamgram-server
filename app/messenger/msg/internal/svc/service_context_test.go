package svc

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
)

type closeablePublisher struct {
	closed int
}

func (p *closeablePublisher) Publish(context.Context, repository.ReceiverOperation) (repository.KafkaAck, error) {
	return repository.KafkaAck{}, nil
}

func (p *closeablePublisher) Close() error {
	p.closed++
	return nil
}

func TestServiceContextCloseClosesReceiverPublisher(t *testing.T) {
	publisher := &closeablePublisher{}
	ctx := &ServiceContext{ReceiverPublisher: publisher}
	if err := ctx.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if publisher.closed != 1 {
		t.Fatalf("closed = %d, want 1", publisher.closed)
	}
}
