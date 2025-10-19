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

	// NotificationsSent counts notifications sent
	NotificationsSent = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "notifications_sent_total",
			Help: "Total number of notifications sent",
		},
		[]string{"type", "status"},
	)

	// NotificationProcessingTime measures notification processing time
	NotificationProcessingTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "notification_processing_seconds",
			Help:    "Time taken to process notifications",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"type"},
	)

	// QueueLength tracks RabbitMQ queue length
	QueueLength = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rabbitmq_queue_length",
			Help: "Number of messages in RabbitMQ queue",
		},
		[]string{"queue"},
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
		NotificationsSent,
		NotificationProcessingTime,
		QueueLength,
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

// RecordNotificationSent records a notification sent metric
func RecordNotificationSent(notificationType, status string) {
	NotificationsSent.WithLabelValues(notificationType, status).Inc()
}

// RecordNotificationProcessingTime records notification processing time
func RecordNotificationProcessingTime(notificationType string, duration time.Duration) {
	NotificationProcessingTime.WithLabelValues(notificationType).Observe(duration.Seconds())
}

// UpdateQueueLength updates the queue length metric
func UpdateQueueLength(queue string, length float64) {
	QueueLength.WithLabelValues(queue).Set(length)
}

// SetServiceHealth sets the service health status
func SetServiceHealth(healthy bool) {
	if healthy {
		ServiceHealth.Set(1)
	} else {
		ServiceHealth.Set(0)
	}
}
