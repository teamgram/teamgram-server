package kitex

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	ztrace "github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

const traceparentKey = "traceparent"

func serverTracingMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		ctx, span := startServerSpan(ctx)
		defer span.End()

		ztrace.MessageReceived.Event(ctx, 1, req)
		err := next(ctx, req, resp)
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			span.SetAttributes(ztrace.StatusCodeAttr(gcodes.Unknown))
			return err
		}

		span.SetAttributes(ztrace.StatusCodeAttr(gcodes.OK))
		ztrace.MessageSent.Event(ctx, 1, resp)
		return nil
	}
}

func clientTracingMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		ctx, span := startClientSpan(ctx)
		defer span.End()

		ztrace.MessageSent.Event(ctx, 1, req)
		err := next(ctx, req, resp)
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			span.SetAttributes(ztrace.StatusCodeAttr(gcodes.Unknown))
			return err
		}

		span.SetAttributes(ztrace.StatusCodeAttr(gcodes.OK))
		ztrace.MessageReceived.Event(ctx, 1, resp)
		return nil
	}
}

func startServerSpan(ctx context.Context) (context.Context, oteltrace.Span) {
	md := metadataFromMetainfo(ctx)
	bags, spanCtx := ztrace.Extract(ctx, otel.GetTextMapPropagator(), &md)
	ctx = baggage.ContextWithBaggage(ctx, bags)

	name, attr := ztrace.SpanInfo(fullMethod(ctx), peerAddress(ctx, true))
	return otel.Tracer(ztrace.TraceName).Start(
		oteltrace.ContextWithRemoteSpanContext(ctx, spanCtx),
		name,
		oteltrace.WithSpanKind(oteltrace.SpanKindServer),
		oteltrace.WithAttributes(attr...),
	)
}

func startClientSpan(ctx context.Context) (context.Context, oteltrace.Span) {
	name, attr := ztrace.SpanInfo(fullMethod(ctx), peerAddress(ctx, false))
	ctx, span := ztrace.TracerFromContext(ctx).Start(
		ctx,
		name,
		oteltrace.WithSpanKind(oteltrace.SpanKindClient),
		oteltrace.WithAttributes(attr...),
	)

	md := metadata.MD{}
	ztrace.Inject(ctx, otel.GetTextMapPropagator(), &md)
	return metainfoFromMetadata(ctx, md), span
}

func metadataFromMetainfo(ctx context.Context) metadata.MD {
	md := metadata.MD{}
	for k, v := range metainfo.GetAllPersistentValues(ctx) {
		md.Set(k, v)
	}
	for k, v := range metainfo.GetAllValues(ctx) {
		md.Set(k, v)
	}
	return md
}

func metainfoFromMetadata(ctx context.Context, md metadata.MD) context.Context {
	for k, values := range md {
		if len(values) == 0 {
			continue
		}
		ctx = metainfo.WithValue(ctx, k, values[0])
	}
	return ctx
}

func fullMethod(ctx context.Context) string {
	ri := rpcinfo.GetRPCInfo(ctx)
	if ri == nil {
		return "/unknown/unknown"
	}

	serviceName := ""
	methodName := ""
	if inv := ri.Invocation(); inv != nil {
		serviceName = inv.ServiceName()
		methodName = inv.MethodName()
	}
	if serviceName == "" && ri.To() != nil {
		serviceName = ri.To().ServiceName()
	}
	if methodName == "" && ri.To() != nil {
		methodName = ri.To().Method()
	}
	if serviceName == "" {
		serviceName = "unknown"
	}
	if methodName == "" {
		methodName = "unknown"
	}

	return "/" + serviceName + "/" + methodName
}

func peerAddress(ctx context.Context, serverSide bool) string {
	ri := rpcinfo.GetRPCInfo(ctx)
	if ri == nil {
		return ""
	}

	var endpoint rpcinfo.EndpointInfo
	if serverSide {
		endpoint = ri.From()
	} else {
		endpoint = ri.To()
	}
	if endpoint == nil || endpoint.Address() == nil {
		return ""
	}
	return endpoint.Address().String()
}
