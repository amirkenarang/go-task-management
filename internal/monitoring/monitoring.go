package monitoring

import (
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func Init() {
	prometheus.MustRegister(RequestCounter)
	prometheus.MustRegister(RequestDuration)
}

func MetricsHandler() http.Handler {
	return promhttp.Handler()
}

func ServeMetrics(c *fiber.Ctx) {
	handler := MetricsHandler()

	// Create a minimal http.Request
	req := &http.Request{
		Method: string(c.Context().Method()),
		URL:    &url.URL{Path: string(c.Context().Path())},
		Header: http.Header{},
	}

	writer := &fiberResponseWriter{ctx: c}

	// Serve the metrics
	handler.ServeHTTP(writer, req)
}

type fiberResponseWriter struct {
	ctx *fiber.Ctx
}

func (w *fiberResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w *fiberResponseWriter) Write(data []byte) (int, error) {
	return w.ctx.Context().Response.BodyWriter().Write(data)
}

func (w *fiberResponseWriter) WriteHeader(statusCode int) {
	w.ctx.Status(statusCode)
}
