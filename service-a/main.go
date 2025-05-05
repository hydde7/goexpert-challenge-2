package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/hydde7/goexpert-challenge-2/service-a/cmd"
	"github.com/hydde7/goexpert-challenge-2/service-a/internal/cfg"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initTracer(ctx context.Context) func() {
	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(cfg.Otl.OTELPEndpoint),
	)

	if err != nil {
		log.Fatalf("failed to create exporter: %v", err)
	}
	res := sdkresource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(cfg.Otl.ServiceName),
	)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatalf("error shutting down tracer provider: %v", err)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "goexpert-challenge-2/service-a"
	app.Usage = "Challenge 2 for GoExpert"
	app.Flags = cfg.Flags
	app.Action = cli.ActionFunc(run)
	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Fatal("failed to run CLI app")
	}
	logrus.Info("Shutting down...")
}

func run(c *cli.Context) error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	level, err := logrus.ParseLevel(cfg.App.LogLevel)
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse log level")
	}
	logrus.SetLevel(level)
	logrus.Info("Starting application...")
	router := cmd.SetupRouter()
	logrus.Info("Router setup complete")

	shutdown := initTracer(context.Background())
	defer shutdown()

	err = router.Run(":8080")
	if err != nil {
		logrus.WithError(err).Fatal("failed to start server")
	}

	return err
}
