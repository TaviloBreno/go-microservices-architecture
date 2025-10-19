package telemetry

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// InitTracer initializes OpenTelemetry tracer for the service
func InitTracer(serviceName string) func(context.Context) error {
	// Create Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint("http://jaeger:14268/api/traces"),
	))
	if err != nil {
		log.Fatalf("❌ Erro ao criar exporter Jaeger: %v", err)
	}

	// Create tracer provider
	tp := sdktrace.NewTracerProvider(
		// Always be sure to batch in production.
		sdktrace.WithBatcher(exp),
		// Record information about this application in a Resource.
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String("1.0.0"),
		)),
	)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	log.Printf("✅ Tracer inicializado para serviço: %s", serviceName)

	return tp.Shutdown
}

// GetTracer returns a tracer for the given name
func GetTracer(name string) sdktrace.Tracer {
	return otel.Tracer(name)
}
