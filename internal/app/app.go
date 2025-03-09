package app

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sadhakbj/rag-with-go-ollama/internal/config"
	"github.com/sadhakbj/rag-with-go-ollama/internal/di"
	"github.com/sadhakbj/rag-with-go-ollama/internal/utils/logger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	Name      string
	Version   string
	Config    *config.Config
	Container *di.Container
	Logger    *slog.Logger
	Tracer    trace.Tracer
}

func NewApp(config *config.Config) *App {
	return &App{
		Name:    config.AppName,
		Version: config.AppVersion,
		Config:  config,
		Logger:  logger.InitializeLogger(config.AppName, false, slog.LevelInfo),
	}
}

func (a *App) Run() {
	a.Container = di.NewContainer(a.Config)
	a.Logger.Info("Starting application", "version", a.Version)

	// Set up OpenTelemetry and Jaeger
	tp, err := a.initTracer()
	if err != nil {
		a.Logger.Error("Failed to initialize tracer", "error", err)
		os.Exit(1)
	}
	defer func() { _ = tp.Shutdown(context.Background()) }()
	a.Tracer = otel.Tracer(a.Name)

	// Set up OpenTelemetry metrics and Prometheus exporter
	if err := a.initMetrics(); err != nil {
		a.Logger.Error("Failed to initialize metrics", "error", err)
		os.Exit(1)
	}

	// Instrument HTTP handlers
	http.Handle("/hello", otelhttp.NewHandler(http.HandlerFunc(a.handleHello), "handleHello"))
	http.Handle("/world", otelhttp.NewHandler(http.HandlerFunc(a.handleWorld), "handleWorld"))

	a.Logger.Info("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		a.Logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}

func (a *App) initTracer() (*sdktrace.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(a.Name),
			semconv.ServiceVersionKey.String(a.Version),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func (a *App) initMetrics() error {
	exporter, err := prometheus.New()
	if err != nil {
		return err
	}

	provider := metric.NewMeterProvider(
		metric.WithReader(exporter),
		metric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(a.Name),
			semconv.ServiceVersionKey.String(a.Version),
		)),
	)
	otel.SetMeterProvider(provider)

	http.Handle("/metrics", promhttp.Handler())
	return nil
}

func (a *App) handleHello(w http.ResponseWriter, r *http.Request) {
	_, span := a.Tracer.Start(r.Context(), "handleHello")
	defer span.End()

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func (a *App) handleWorld(w http.ResponseWriter, r *http.Request) {
	_, span := a.Tracer.Start(r.Context(), "handleWorld")
	defer span.End()

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "World, Hello!"}
	json.NewEncoder(w).Encode(response)
}
