package kitex

import (
	"context"
	"testing"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func TestServerTracingMiddlewareAddsSpanToHandlerContext(t *testing.T) {
	restoreTracerProvider(t)

	var handlerTraceID string
	next := func(ctx context.Context, req, resp interface{}) error {
		spanCtx := oteltrace.SpanContextFromContext(ctx)
		if !spanCtx.IsValid() {
			t.Fatal("expected valid span context in server handler")
		}
		handlerTraceID = spanCtx.TraceID().String()
		return nil
	}

	err := serverTracingMiddleware(endpoint.Endpoint(next))(context.Background(), nil, nil)
	if err != nil {
		t.Fatalf("server tracing middleware returned error: %v", err)
	}
	if handlerTraceID == "" {
		t.Fatal("expected handler trace id to be recorded")
	}
}

func TestClientTracingMiddlewareInjectsTraceContextIntoMetainfo(t *testing.T) {
	restoreTracerProvider(t)

	var (
		traceparent string
		childTrace  string
	)
	parentCtx, parentSpan := otel.Tracer("test").Start(context.Background(), "parent")
	defer parentSpan.End()
	parentTrace := parentSpan.SpanContext().TraceID().String()

	next := func(ctx context.Context, req, resp interface{}) error {
		var ok bool
		traceparent, ok = metainfo.GetValue(ctx, traceparentKey)
		if !ok || traceparent == "" {
			t.Fatal("expected traceparent to be injected into metainfo")
		}
		childTrace = oteltrace.SpanContextFromContext(ctx).TraceID().String()
		return nil
	}

	err := clientTracingMiddleware(endpoint.Endpoint(next))(parentCtx, nil, nil)
	if err != nil {
		t.Fatalf("client tracing middleware returned error: %v", err)
	}
	if childTrace != parentTrace {
		t.Fatalf("expected child trace %s to match parent trace %s", childTrace, parentTrace)
	}
}

func TestTracingMiddlewaresPropagateTraceFromClientToServer(t *testing.T) {
	restoreTracerProvider(t)

	parentCtx, parentSpan := otel.Tracer("test").Start(context.Background(), "parent")
	defer parentSpan.End()
	parentSpanCtx := parentSpan.SpanContext()

	var outboundCtx context.Context
	clientNext := func(ctx context.Context, req, resp interface{}) error {
		outboundCtx = ctx
		return nil
	}
	if err := clientTracingMiddleware(endpoint.Endpoint(clientNext))(parentCtx, nil, nil); err != nil {
		t.Fatalf("client tracing middleware returned error: %v", err)
	}

	var serverSpanCtx oteltrace.SpanContext
	serverNext := func(ctx context.Context, req, resp interface{}) error {
		serverSpanCtx = oteltrace.SpanContextFromContext(ctx)
		return nil
	}
	if err := serverTracingMiddleware(endpoint.Endpoint(serverNext))(outboundCtx, nil, nil); err != nil {
		t.Fatalf("server tracing middleware returned error: %v", err)
	}

	if serverSpanCtx.TraceID() != parentSpanCtx.TraceID() {
		t.Fatalf("expected server trace %s to match parent trace %s", serverSpanCtx.TraceID(), parentSpanCtx.TraceID())
	}
	if serverSpanCtx.SpanID() == parentSpanCtx.SpanID() {
		t.Fatal("expected server span id to differ from parent span id")
	}
}

func restoreTracerProvider(t *testing.T) {
	t.Helper()

	orig := otel.GetTracerProvider()
	otel.SetTracerProvider(trace.NewTracerProvider(trace.WithSampler(trace.AlwaysSample())))
	t.Cleanup(func() {
		otel.SetTracerProvider(orig)
	})
}
