package olly

import "context"

type otelContextKey struct{}

// WithOTel Registers an OTelRecorder in the context, which can be retrieved using RecorderFromContext.
// This is necessary to allow other functions in this package to operate.
func WithOTel(ctx context.Context, recorder OTelRecorder) context.Context {
	return context.WithValue(ctx, otelContextKey{}, recorder)
}

func RecorderFromContext(ctx context.Context) (OTelRecorder, bool) {
	tracer, ok := ctx.Value(otelContextKey{}).(OTelRecorder)

	return tracer, ok
}
