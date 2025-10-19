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

	// GraphQLQueries counts GraphQL queries
	GraphQLQueries = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "graphql_queries_total",
			Help: "Total number of GraphQL queries",
		},
		[]string{"operation", "status"},
	)

	// GraphQLQueryDuration measures GraphQL query duration
	GraphQLQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "graphql_query_duration_seconds",
			Help:    "Duration of GraphQL queries in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation"},
	)

	// GRPCCallsTotal counts gRPC calls to backend services
	GRPCCallsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_calls_total",
			Help: "Total number of gRPC calls to backend services",
		},
		[]string{"service", "method", "status"},
	)

	// GRPCCallDuration measures gRPC call duration
	GRPCCallDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_call_duration_seconds",
			Help:    "Duration of gRPC calls to backend services",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"service", "method"},
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
		GraphQLQueries,
		GraphQLQueryDuration,
		GRPCCallsTotal,
		GRPCCallDuration,
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

// RecordGraphQLQuery records a GraphQL query metric
func RecordGraphQLQuery(operation, status string, duration time.Duration) {
	GraphQLQueries.WithLabelValues(operation, status).Inc()
	GraphQLQueryDuration.WithLabelValues(operation).Observe(duration.Seconds())
}

// RecordGRPCCall records a gRPC call metric
func RecordGRPCCall(service, method, status string, duration time.Duration) {
	GRPCCallsTotal.WithLabelValues(service, method, status).Inc()
	GRPCCallDuration.WithLabelValues(service, method).Observe(duration.Seconds())
}

// SetServiceHealth sets the service health status
func SetServiceHealth(healthy bool) {
	if healthy {
		ServiceHealth.Set(1)
	} else {
		ServiceHealth.Set(0)
	}
}
