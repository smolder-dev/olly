package olly

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/contrib/exporters/autoexport"
	"go.opentelemetry.io/contrib/propagators/autoprop"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

type OTelProviders struct {
	TracerProvider *trace.TracerProvider
	MeterProvider  *metric.MeterProvider
	LoggerProvider *log.LoggerProvider
}

// OTelProviderConfig holds configuration options for the OpenTelemetry providers.
// An empty OTelProviderConfig will result in the default configuration.
type OTelProviderConfig struct {
	DisableDefaultTracerExporter bool
	TracerProviderOptions        []trace.TracerProviderOption
	DisableDefaultMeterExporter  bool
	MeterProviderOptions         []metric.Option
	DisableDefaultLoggerExporter bool
	LoggerProviderOptions        []log.LoggerProviderOption
	DisableGlobalProviders       bool
}

func newTracerProvider(
	ctx context.Context,
	config OTelProviderConfig,
) (*trace.TracerProvider, error) {
	if !config.DisableDefaultTracerExporter {
		traceExporter, err := autoexport.NewSpanExporter(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create OTLP trace exporter: %w", err)
		}

		config.TracerProviderOptions = append(config.TracerProviderOptions, trace.WithBatcher(traceExporter))
	}

	tracerProvider := trace.NewTracerProvider(config.TracerProviderOptions...)

	return tracerProvider, nil
}

func newMeterProvider(ctx context.Context, config OTelProviderConfig) (*metric.MeterProvider, error) {
	if config.DisableDefaultMeterExporter {
		metricExporter, err := autoexport.NewMetricReader(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create OTLP metric exporter: %w", err)
		}

		config.MeterProviderOptions = append(
			config.MeterProviderOptions,
			metric.WithReader(metricExporter),
		)
	}

	meterProvider := metric.NewMeterProvider(config.MeterProviderOptions...)

	return meterProvider, nil
}

func newLoggerProvider(ctx context.Context, config OTelProviderConfig) (*log.LoggerProvider, error) {
	logExporter, err := autoexport.NewLogExporter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP log exporter: %w", err)
	}

	if !config.DisableDefaultLoggerExporter {
		config.LoggerProviderOptions = append(
			config.LoggerProviderOptions,
			log.WithProcessor(log.NewBatchProcessor(logExporter)),
		)
	}

	loggerProvider := log.NewLoggerProvider(config.LoggerProviderOptions...)

	return loggerProvider, nil
}

// NewOTelProviders initializes and returns OpenTelemetry providers for tracing, metrics, and logging
// based on the environment variables and configuration options provided.
// The providers are returned in an OTelProviders struct, which includes a Shutdown method to clean up resources.
// The development parameter can be used to adjust the batch timeout for trace and log exporters
// making it shorter for development environments.
func NewOTelProviders(
	ctx context.Context,
	config OTelProviderConfig,
) (OTelProviders, error) {
	//nolint:exhaustruct
	providers := OTelProviders{}

	shutdown := false

	defer func() {
		if shutdown {
			//nolint:contextcheck // the context might be cancelled when shutdown is called
			// so we use a background context here.
			err := providers.Shutdown(context.Background())
			if err != nil {
				slog.ErrorContext(ctx, "failed to shutdown OpenTelemetry providers during initialization", "error", err)
			}
		}
	}()

	propagator := autoprop.NewTextMapPropagator()
	otel.SetTextMapPropagator(propagator)

	tracerProvider, err := newTracerProvider(ctx, config)
	if err != nil {
		shutdown = true

		return OTelProviders{}, fmt.Errorf("failed to create tracer provider: %w", err)
	}

	if !config.DisableGlobalProviders {
		otel.SetTracerProvider(tracerProvider)
	}

	providers.TracerProvider = tracerProvider

	// Set up meter provider.
	meterProvider, err := newMeterProvider(ctx, config)
	if err != nil {
		shutdown = true

		return OTelProviders{}, fmt.Errorf("failed to create meter provider: %w", err)
	}

	if !config.DisableGlobalProviders {
		otel.SetMeterProvider(meterProvider)
	}

	providers.MeterProvider = meterProvider

	// Set up logger provider.
	loggerProvider, err := newLoggerProvider(ctx, config)
	if err != nil {
		shutdown = true

		return OTelProviders{}, fmt.Errorf("failed to create logger provider: %w", err)
	}

	if !config.DisableGlobalProviders {
		global.SetLoggerProvider(loggerProvider)
	}

	providers.LoggerProvider = loggerProvider

	return providers, nil
}

func (p OTelProviders) Shutdown(ctx context.Context) error {
	var outerErr error

	if p.TracerProvider != nil {
		err := p.TracerProvider.Shutdown(ctx)
		if err != nil {
			outerErr = errors.Join(outerErr, fmt.Errorf("failed to shutdown tracer provider: %w", err))
		}
	}

	if p.MeterProvider != nil {
		err := p.MeterProvider.Shutdown(ctx)
		if err != nil {
			outerErr = errors.Join(outerErr, fmt.Errorf("failed to shutdown meter provider: %w", err))
		}
	}

	if p.LoggerProvider != nil {
		err := p.LoggerProvider.Shutdown(ctx)
		if err != nil {
			outerErr = errors.Join(outerErr, fmt.Errorf("failed to shutdown logger provider: %w", err))
		}
	}

	return outerErr
}
