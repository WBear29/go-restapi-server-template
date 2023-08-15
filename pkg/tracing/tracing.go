package tracing

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// newExporter returns a jaeger exporter.
func newExporter(endpoint string) (trace.SpanExporter, error) {
	return jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)),
	)
}

// newResource returns a resource describing this application.
func newResource(name, version, env string) (*resource.Resource, error) {
	return resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(name),
			semconv.ServiceVersion(version),
			attribute.String("environment", env),
		),
	)
}

func NewTraceProvider(name, version, env, endpoint string) (*trace.TracerProvider, error) {
	exp, err := newExporter(endpoint)
	if err != nil {
		return nil, err
	}
	resource, err := newResource(name, version, env)
	if err != nil {
		return nil, err
	}
	provider := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource),
	)
	return provider, nil
}
