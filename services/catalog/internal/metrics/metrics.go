package metrics

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// RequestsTotal counts the total number of HTTP requests
	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	// RequestDuration measures the duration of HTTP requests
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// GRPCRequestsTotal counts the total number of gRPC requests
	GRPCRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",
			Help: "Total number of gRPC requests",
		},
		[]string{"method", "status"},
	)

	// GRPCRequestDuration measures the duration of gRPC requests
	GRPCRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_request_duration_seconds",
			Help:    "Duration of gRPC requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	// ProductsQueried counts product queries
	ProductsQueried = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "products_queried_total",
			Help: "Total number of product queries",
		},
		[]string{"operation"},
	)

	// CatalogUpdates counts catalog updates
	CatalogUpdates = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "catalog_updates_total",
			Help: "Total number of catalog updates",
		},
		[]string{"operation", "status"},
	)

	// ServiceHealth indicates if the service is healthy
	ServiceHealth = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "service_health",
			Help: "Health status of the service (1 = healthy, 0 = unhealthy)",
		},
	)
)

// Init initializes Prometheus metrics and starts the metrics server
func Init() {
	// Register metrics
	prometheus.MustRegister(
		RequestsTotal,
		RequestDuration,
		GRPCRequestsTotal,
		GRPCRequestDuration,
		ProductsQueried,
		CatalogUpdates,
		ServiceHealth,
	)

	// Set initial health status
	ServiceHealth.Set(1)

	// Setup metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Start metrics server in a goroutine
	go func() {
		log.Printf("🔍 Metrics server starting on :9091")
		if err := http.ListenAndServe(":9091", nil); err != nil {
			log.Printf("❌ Erro ao iniciar servidor de métricas: %v", err)
		}
	}()

	log.Printf("✅ Métricas Prometheus inicializadas")
}

// RecordHTTPRequest records an HTTP request metric
func RecordHTTPRequest(method, endpoint, status string, duration time.Duration) {
	RequestsTotal.WithLabelValues(method, endpoint, status).Inc()
	RequestDuration.WithLabelValues(method, endpoint).Observe(duration.Seconds())
}

// RecordGRPCRequest records a gRPC request metric
func RecordGRPCRequest(method, status string, duration time.Duration) {
	GRPCRequestsTotal.WithLabelValues(method, status).Inc()
	GRPCRequestDuration.WithLabelValues(method).Observe(duration.Seconds())
}

// RecordProductQuery records a product query metric
func RecordProductQuery(operation string) {
	ProductsQueried.WithLabelValues(operation).Inc()
}

// RecordCatalogUpdate records a catalog update metric
func RecordCatalogUpdate(operation, status string) {
	CatalogUpdates.WithLabelValues(operation, status).Inc()
}

// SetServiceHealth sets the service health status
func SetServiceHealth(healthy bool) {
	if healthy {
		ServiceHealth.Set(1)
	} else {
		ServiceHealth.Set(0)
	}
}
