package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	corruptEntryTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "presence_corrupt_entry_total",
		Help: "Corrupt presence entries skipped by class.",
	}, []string{"error_class"})

	corruptEntryCleanupFailureTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "presence_corrupt_entry_cleanup_failure_total",
		Help: "Failed best-effort cleanup attempts for corrupt presence entries.",
	}, []string{"error_class"})

	permissionDeniedTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "presence_permission_denied_total",
		Help: "Rejected presence RPC calls by method and caller.",
	}, []string{"method", "caller"})
)

func CorruptEntry(errorClass string) {
	corruptEntryTotal.WithLabelValues(errorClass).Inc()
}

func CorruptEntryCleanupFailure(errorClass string) {
	corruptEntryCleanupFailureTotal.WithLabelValues(errorClass).Inc()
}

func PermissionDenied(method, caller string) {
	permissionDeniedTotal.WithLabelValues(method, caller).Inc()
}
