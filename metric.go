package olly

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/metric"
)

var ErrNoRecorderInContext = fmt.Errorf("no recorder found in context")

//nolint:ireturn
func getMeter(ctx context.Context) (metric.Meter, error) {
	recorder, ok := RecorderFromContext(ctx)
	if !ok {
		return nil, ErrNoRecorderInContext
	}

	return recorder.meter, nil
}

// Int64Counter returns a new Int64Counter instrument identified by name
// and configured with options. The instrument is used to synchronously
// record increasing int64 measurements during a computational operation.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Int64Counter(
	ctx context.Context,
	name string,
	options ...metric.Int64CounterOption,
) (metric.Int64Counter, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	counter, err := meter.Int64Counter(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Int64Counter: %w", err)
	}

	return counter, nil
}

// Int64UpDownCounter returns a new Int64UpDownCounter instrument
// identified by name and configured with options. The instrument is used
// to synchronously record int64 measurements during a computational
// operation.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Int64UpDownCounter(
	ctx context.Context,
	name string,
	options ...metric.Int64UpDownCounterOption,
) (metric.Int64UpDownCounter, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	counter, err := meter.Int64UpDownCounter(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Int64UpDownCounter: %w", err)
	}

	return counter, nil
}

// Int64Histogram returns a new Int64Histogram instrument identified by
// name and configured with options. The instrument is used to
// synchronously record the distribution of int64 measurements during a
// computational operation.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Int64Histogram(
	ctx context.Context,
	name string,
	options ...metric.Int64HistogramOption,
) (metric.Int64Histogram, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	histogram, err := meter.Int64Histogram(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Int64Histogram: %w", err)
	}

	return histogram, nil
}

// Int64Gauge returns a new Int64Gauge instrument identified by name and
// configured with options. The instrument is used to synchronously record
// instantaneous int64 measurements during a computational operation.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Int64Gauge(
	ctx context.Context,
	name string,
	options ...metric.Int64GaugeOption,
) (metric.Int64Gauge, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	gauge, err := meter.Int64Gauge(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Int64Gauge: %w", err)
	}

	return gauge, nil
}

// Int64ObservableCounter returns a new Int64ObservableCounter identified
// by name and configured with options. The instrument is used to
// asynchronously record increasing int64 measurements once per a
// measurement collection cycle.
//
// Measurements for the returned instrument are made via a callback. Use
// the WithInt64Callback option to register the callback here, or use the
// RegisterCallback method of this Meter to register one later. See the
// Measurements section of the package documentation for more information.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Int64ObservableCounter(
	ctx context.Context,
	name string,
	options ...metric.Int64ObservableCounterOption,
) (metric.Int64ObservableCounter, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	counter, err := meter.Int64ObservableCounter(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Int64ObservableCounter: %w", err)
	}

	return counter, nil
}

// Int64ObservableUpDownCounter returns a new Int64ObservableUpDownCounter
// instrument identified by name and configured with options. The
// instrument is used to asynchronously record int64 measurements once per
// a measurement collection cycle.
//
// Measurements for the returned instrument are made via a callback. Use
// the WithInt64Callback option to register the callback here, or use the
// RegisterCallback method of this Meter to register one later. See the
// Measurements section of the package documentation for more information.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Int64ObservableUpDownCounter(
	ctx context.Context,
	name string,
	options ...metric.Int64ObservableUpDownCounterOption,
) (metric.Int64ObservableUpDownCounter, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	counter, err := meter.Int64ObservableUpDownCounter(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Int64ObservableUpDownCounter: %w", err)
	}

	return counter, nil
}

// Int64ObservableGauge returns a new Int64ObservableGauge instrument
// identified by name and configured with options. The instrument is used
// to asynchronously record instantaneous int64 measurements once per a
// measurement collection cycle.
//
// Measurements for the returned instrument are made via a callback. Use
// the WithInt64Callback option to register the callback here, or use the
// RegisterCallback method of this Meter to register one later. See the
// Measurements section of the package documentation for more information.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Int64ObservableGauge(
	ctx context.Context,
	name string,
	options ...metric.Int64ObservableGaugeOption,
) (metric.Int64ObservableGauge, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	gauge, err := meter.Int64ObservableGauge(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Int64ObservableGauge: %w", err)
	}

	return gauge, nil
}

// Float64Counter returns a new Float64Counter instrument identified by
// name and configured with options. The instrument is used to
// synchronously record increasing float64 measurements during a
// computational operation.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
//nolint:ireturn
func Float64Counter(
	ctx context.Context,
	name string,
	options ...metric.Float64CounterOption,
) (metric.Float64Counter, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	counter, err := meter.Float64Counter(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Float64Counter: %w", err)
	}

	return counter, nil
}

// Float64UpDownCounter returns a new Float64UpDownCounter instrument
// identified by name and configured with options. The instrument is used
// to synchronously record float64 measurements during a computational
// operation.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Float64UpDownCounter(
	ctx context.Context,
	name string,
	options ...metric.Float64UpDownCounterOption,
) (metric.Float64UpDownCounter, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	counter, err := meter.Float64UpDownCounter(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Float64UpDownCounter: %w", err)
	}

	return counter, nil
}

// Float64Histogram returns a new Float64Histogram instrument identified by
// name and configured with options. The instrument is used to
// synchronously record the distribution of float64 measurements during a
// computational operation.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Float64Histogram(
	ctx context.Context,
	name string,
	options ...metric.Float64HistogramOption,
) (metric.Float64Histogram, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	histogram, err := meter.Float64Histogram(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Float64Histogram: %w", err)
	}

	return histogram, nil
}

// Float64Gauge returns a new Float64Gauge instrument identified by name and
// configured with options. The instrument is used to synchronously record
// instantaneous float64 measurements during a computational operation.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Float64Gauge(
	ctx context.Context,
	name string,
	options ...metric.Float64GaugeOption,
) (metric.Float64Gauge, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	gauge, err := meter.Float64Gauge(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Float64Gauge: %w", err)
	}

	return gauge, nil
}

// Float64ObservableCounter returns a new Float64ObservableCounter
// instrument identified by name and configured with options. The
// instrument is used to asynchronously record increasing float64
// measurements once per a measurement collection cycle.
//
// Measurements for the returned instrument are made via a callback. Use
// the WithFloat64Callback option to register the callback here, or use the
// RegisterCallback method of this Meter to register one later. See the
// Measurements section of the package documentation for more information.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Float64ObservableCounter(
	ctx context.Context,
	name string,
	options ...metric.Float64ObservableCounterOption,
) (metric.Float64ObservableCounter, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	counter, err := meter.Float64ObservableCounter(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Float64ObservableCounter: %w", err)
	}

	return counter, nil
}

// Float64ObservableUpDownCounter returns a new
// Float64ObservableUpDownCounter instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// float64 measurements once per a measurement collection cycle.
//
// Measurements for the returned instrument are made via a callback. Use
// the WithFloat64Callback option to register the callback here, or use the
// RegisterCallback method of this Meter to register one later. See the
// Measurements section of the package documentation for more information.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Float64ObservableUpDownCounter(
	ctx context.Context,
	name string,
	options ...metric.Float64ObservableUpDownCounterOption,
) (metric.Float64ObservableUpDownCounter, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	counter, err := meter.Float64ObservableUpDownCounter(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Float64ObservableUpDownCounter: %w", err)
	}

	return counter, nil
}

// Float64ObservableGauge returns a new Float64ObservableGauge instrument
// identified by name and configured with options. The instrument is used
// to asynchronously record instantaneous float64 measurements once per a
// measurement collection cycle.
//
// Measurements for the returned instrument are made via a callback. Use
// the WithFloat64Callback option to register the callback here, or use the
// RegisterCallback method of this Meter to register one later. See the
// Measurements section of the package documentation for more information.
//
// The name needs to conform to the OpenTelemetry instrument name syntax.
// See the Instrument Name section of the package documentation for more
// information.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
//nolint:ireturn
func Float64ObservableGauge(
	ctx context.Context,
	name string,
	options ...metric.Float64ObservableGaugeOption,
) (metric.Float64ObservableGauge, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	gauge, err := meter.Float64ObservableGauge(name, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Float64ObservableGauge: %w", err)
	}

	return gauge, nil
}

// RegisterCallback registers f to be called during the collection of a
// measurement cycle.
//
// If Unregister of the returned Registration is called, f needs to be
// unregistered and not called during collection.
//
// The instruments f is registered with are the only instruments that f may
// observe values for.
//
// If no instruments are passed, f should not be registered nor called
// during collection.
//
// Implementations of this method need to be safe for a user to call
// concurrently.
//
// The function f needs to be concurrent safe.
//
//nolint:ireturn
func RegisterCallback(
	ctx context.Context,
	f metric.Callback,
	instruments ...metric.Observable,
) (metric.Registration, error) {
	meter, err := getMeter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get meter from context: %w", err)
	}

	registration, err := meter.RegisterCallback(f, instruments...)
	if err != nil {
		return nil, fmt.Errorf("failed to register callback: %w", err)
	}

	return registration, nil
}
