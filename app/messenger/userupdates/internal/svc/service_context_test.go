package svc

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
)

func TestServiceContextStartsNoWorkersWithoutKafkaConfig(t *testing.T) {
	c := config.Config{}
	ctx := NewServiceContext(c)
	if len(ctx.workers) != 0 {
		t.Fatalf("workers = %d, want 0", len(ctx.workers))
	}
}

func TestServiceContextDoesNotStartPushWorkerWhenDisabled(t *testing.T) {
	c := config.Config{
		PushOutboxWorker: config.PushOutboxWorkerConf{
			Enabled: false,
		},
	}
	ctx := NewServiceContext(c)
	for _, worker := range ctx.workers {
		if _, ok := worker.(*repository.PushOutboxWorker); ok {
			t.Fatal("push outbox worker should not be installed when disabled")
		}
	}
}

func TestServiceContextDoesNotStartAffectedOutboxWorkerWhenDisabled(t *testing.T) {
	c := config.Config{
		AffectedOutboxWorker: config.AffectedOutboxWorkerConf{
			Enabled: false,
		},
	}
	ctx := NewServiceContext(c)
	for _, worker := range ctx.workers {
		if _, ok := worker.(*repository.AffectedOutboxWorker); ok {
			t.Fatal("affected outbox worker should not be installed when disabled")
		}
	}
}

func TestServiceContextStartsAffectedOutboxWorkerWhenEnabled(t *testing.T) {
	c := config.Config{
		AffectedOutboxWorker: config.AffectedOutboxWorkerConf{
			Enabled:             true,
			PollIntervalMs:      200,
			BatchSize:           100,
			ProcessingTimeoutMs: 60000,
		},
	}
	ctx := NewServiceContext(c)
	found := false
	for _, worker := range ctx.workers {
		if _, ok := worker.(*repository.AffectedOutboxWorker); ok {
			found = true
		}
	}
	if !found {
		t.Fatal("affected outbox worker was not installed when enabled")
	}
}

type testCloser struct{ closed int }

func (c *testCloser) Close() error {
	c.closed++
	return nil
}

func TestServiceContextCloseClosesRegisteredClosers(t *testing.T) {
	ctx := &ServiceContext{}
	closer := &testCloser{}
	ctx.closers = append(ctx.closers, closer)
	if err := ctx.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if closer.closed != 1 {
		t.Fatalf("closed = %d, want 1", closer.closed)
	}
}

type waitableWorker struct {
	stopped bool
	waited  bool
}

func (w *waitableWorker) Run(context.Context) {}

func (w *waitableWorker) Stop() {
	w.stopped = true
}

func (w *waitableWorker) Wait() {
	if !w.stopped {
		panic("Wait called before Stop")
	}
	w.waited = true
}

func TestServiceContextCloseWaitsForWaitableWorkers(t *testing.T) {
	worker := &waitableWorker{}
	ctx := &ServiceContext{workers: []backgroundWorker{worker}}

	if err := ctx.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if !worker.stopped {
		t.Fatal("worker was not stopped")
	}
	if !worker.waited {
		t.Fatal("worker was not waited")
	}
}
