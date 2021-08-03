package processor

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metricsExecuteDuration = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: "storable_execute_duration_ms",
		Help: "Duration to execute the storable",
	}, []string{"storable"})

	metricsRollbackDuration = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: "storable_rollback_duration_ms",
		Help: "Duration to rollback the storable",
	}, []string{"storable"})

	metricsSaveDuration = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: "storable_save_duration_ms",
		Help: "Duration to save the storable to db",
	}, []string{"storable"})
)

func recordExecuteDuration(task string, start time.Time) {
	d := float64(time.Since(start) / time.Millisecond)
	metricsExecuteDuration.WithLabelValues(task).Observe(d)
}

func recordRollbackDuration(task string, start time.Time) {
	d := float64(time.Since(start) / time.Millisecond)
	metricsRollbackDuration.WithLabelValues(task).Observe(d)
}

func recordSaveDuration(task string, start time.Time) {
	d := float64(time.Since(start) / time.Millisecond)
	metricsSaveDuration.WithLabelValues(task).Observe(d)
}
