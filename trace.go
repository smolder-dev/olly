package olly

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

//nolint:ireturn
func getTracer(ctx context.Context) (trace.Tracer, bool) {
	recorder, ok := RecorderFromContext(ctx)
	if !ok {
		return nil, false
	}

	return recorder.tracer, true
}

// Nested creates a span and a context.Context containing the newly-created span.
//
// If the context.Context provided in `ctx` contains a Span then the newly-created
// Span will be a child of that span, otherwise it will be a root span. This behavior
// can be overridden by providing `WithNewRoot()` as a SpanOption, causing the
// newly-created Span to be a root span even if `ctx` contains a Span.
//
// When creating a Span it is recommended to provide all known span attributes using
// the `WithAttributes()` SpanOption as samplers will only have access to the
// attributes provided when a Span is created.
//
// Any Span that is created MUST also be ended. This is the responsibility of the user.
// Implementations of this API may leak memory or other resources if Spans are not ended.
//
//nolint:ireturn
func Nested(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if tracer, ok := getTracer(ctx); ok {
		//nolint:spancheck
		ctx, span := tracer.Start(
			ctx,
			name,
			opts...,
		)

		//nolint:spancheck
		return ctx, span
	} else {
		//nolint:exhaustruct
		return ctx, noop.Span{}
	}
}

// Nest wraps around Nested, calling the provided nestedFn with the context containing the newly-created span.
// It automatically ends the span after nestedFn returns.
func Nest(ctx context.Context, name string, nestedFn func(ctx context.Context), opts ...trace.SpanStartOption) {
	ctx, span := Nested(ctx, name, opts...)
	defer span.End()

	nestedFn(ctx)
}

// RecordError will record err as an exception span event for this span.
// If the Status of the Span should be set to Error, Fail or Failf should be used instead.
// If this span is not being recorded or err is nil then this method does nothing.
func RecordError(ctx context.Context, err error, opts ...trace.EventOption) {
	span := trace.SpanFromContext(ctx)
	span.RecordError(err, opts...)
}

// Fail records the error and sets the status of the Span to Error.
func Fail(ctx context.Context, err error, opts ...trace.EventOption) error {
	span := trace.SpanFromContext(ctx)
	span.RecordError(err, opts...)
	span.SetStatus(codes.Error, err.Error())

	return err
}

// Failf wraps around Fail, formatting the error message according to a format specifier and arguments.
func Failf(ctx context.Context, format string, args ...any) error {
	return Fail(ctx, fmt.Errorf(format, args...))
}
