package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
)

func main() {
	// Create an OTLP exporter and register it as a trace provider.
	ctx := context.Background()
	exporter, err := otlp.NewExporter(ctx, otlpgrpc.NewDriver(
		otlpgrpc.WithInsecure(),
		otlpgrpc.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")),
	))
	if err != nil {
		log.Fatal(err)
	}
	provider := trace.NewTracerProvider(
		trace.WithSpanProcessor(
			trace.NewBatchSpanProcessor(exporter),
		),
		trace.WithResource(
			semconv.ServiceNameKey.String(os.Getenv("OTEL_SERVICE_NAME")),
		),
	)
	otel.SetTracerProvider(provider)

	// Set up the web server and start listening for requests.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tracer := otel.Tracer("my-go-app")
		_, span := tracer.Start(ctx, "handleRequest")
		defer span.End()

		// Handle the request.
		fmt.Fprintf(w, "Hello, world!\n")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
