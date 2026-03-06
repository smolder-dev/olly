package olly

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type OTelRecorder struct {
	name      string
	config    otelRecorderConfig
	providers OTelProviders
	tracer    trace.Tracer
	meter     metric.Meter
	logger    log.Logger
}

type otelRecorderConfig struct {
	InstrumentationVersion string
	SchemaURL              string
	Attrs                  attribute.Set
}

type otelRecorderOptions interface {
	apply(config otelRecorderConfig) otelRecorderConfig
}

type otelRecorderOptionInstrumentationVersion string

func (o otelRecorderOptionInstrumentationVersion) apply(config otelRecorderConfig) otelRecorderConfig {
	config.InstrumentationVersion = string(o)

	return config
}

func WithInstrumentationVersion(version string) otelRecorderOptionInstrumentationVersion {
	return otelRecorderOptionInstrumentationVersion(version)
}

var _ otelRecorderOptions = otelRecorderOptionInstrumentationVersion("")

type otelRecorderOptionSchemaURL string

func (o otelRecorderOptionSchemaURL) apply(config otelRecorderConfig) otelRecorderConfig {
	config.SchemaURL = string(o)

	return config
}

func WithInstrumentationSchemaURL(schemaURL string) otelRecorderOptionSchemaURL {
	return otelRecorderOptionSchemaURL(schemaURL)
}

var _ otelRecorderOptions = otelRecorderOptionSchemaURL("")

type otelRecorderOptionAttrs attribute.Set

func (o otelRecorderOptionAttrs) apply(config otelRecorderConfig) otelRecorderConfig {
	config.Attrs = attribute.Set(o)

	return config
}

func WithInstrumentationAttributeSet(attrs attribute.Set) otelRecorderOptionAttrs {
	return otelRecorderOptionAttrs(attrs)
}

func WithInstrumentationAttributes(attrs ...attribute.KeyValue) otelRecorderOptionAttrs {
	return WithInstrumentationAttributeSet(attribute.NewSet(attrs...))
}

//nolint:exhaustruct
var _ otelRecorderOptions = otelRecorderOptionAttrs{}

func NewOTelRecorder(
	ctx context.Context,
	providers OTelProviders,
	name string,
	options ...otelRecorderOptions,
) OTelRecorder {
	//nolint:exhaustruct
	config := otelRecorderConfig{}
	for _, option := range options {
		config = option.apply(config)
	}

	tracerOptions := []trace.TracerOption{
		trace.WithInstrumentationVersion(config.InstrumentationVersion),
		trace.WithSchemaURL(config.SchemaURL),
		trace.WithInstrumentationAttributeSet(config.Attrs),
	}

	tracer := providers.TracerProvider.Tracer(name, tracerOptions...)

	meterOptions := []metric.MeterOption{
		metric.WithInstrumentationVersion(config.InstrumentationVersion),
		metric.WithSchemaURL(config.SchemaURL),
		metric.WithInstrumentationAttributeSet(config.Attrs),
	}

	meter := providers.MeterProvider.Meter(name, meterOptions...)

	loggerOptions := []log.LoggerOption{
		log.WithInstrumentationVersion(config.InstrumentationVersion),
		log.WithSchemaURL(config.SchemaURL),
		log.WithInstrumentationAttributeSet(config.Attrs),
	}

	logger := providers.LoggerProvider.Logger(name, loggerOptions...)

	return OTelRecorder{
		name:      name,
		config:    config,
		providers: providers,
		tracer:    tracer,
		meter:     meter,
		logger:    logger,
	}
}

func (r OTelRecorder) NewLogger(handler slog.Handler, options ...otelslog.Option) *slog.Logger {
	loggerOptions := make([]otelslog.Option, 0, 5+len(options))

	loggerOptions = append(
		loggerOptions,
		otelslog.WithLoggerProvider(r.providers.LoggerProvider),
		otelslog.WithSchemaURL(r.config.SchemaURL),
		otelslog.WithVersion(r.config.InstrumentationVersion),
		otelslog.WithAttributes(r.config.Attrs.ToSlice()...),
		otelslog.WithSource(true),
	)
	loggerOptions = append(loggerOptions, options...)

	slogHandler := slog.NewMultiHandler(handler, otelslog.NewHandler(r.name, loggerOptions...))
	slogLogger := slog.New(slogHandler)

	return slogLogger
}
