package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	ctx := context.Background()

	// Create the OTLP exporter
	exporter, err := otlptrace.New(
		ctx,
		otlptrace.WithEndpoint("localhost:55680"),
		otlptrace.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to create OTLP exporter: %v", err)
	}

	// Create the tracer provider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			otel.GetSemanticConventionAttributes(),
			// Add additional attributes as necessary
		)),
	)

	// Set the global tracer provider
	otel.SetTracerProvider(tp)

	// Create a tracer instance
	tracer := otel.Tracer("example")

	// Start a span
	ctx, span := tracer.Start(ctx, "example-operation")
	defer span.End()

	// Simulate some work
	time.Sleep(1 * time.Second)

	log.Println("Example completed successfully")
}
