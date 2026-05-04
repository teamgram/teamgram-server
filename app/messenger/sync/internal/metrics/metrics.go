package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	fanoutFailureTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "sync_fanout_failure_total",
		Help: "Realtime fanout failures by method, gateway, and class.",
	}, []string{"method", "gateway_id", "error_class"})
	rpcResultTargetMissingTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "sync_push_rpc_result_target_missing_total",
		Help: "pushRpcResult target session missing events.",
	}, []string{"gateway_id", "error_class"})
	addressRejectedTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "sync_gateway_address_rejected_total",
		Help: "Rejected direct gateway addresses.",
	}, []string{"error_class"})
)

func FanoutFailure(method, gatewayID, errorClass string) {
	fanoutFailureTotal.WithLabelValues(method, gatewayID, errorClass).Inc()
}

func RpcResultTargetMissing(gatewayID, errorClass string) {
	rpcResultTargetMissingTotal.WithLabelValues(gatewayID, errorClass).Inc()
}

func AddressRejected(errorClass string) {
	addressRejectedTotal.WithLabelValues(errorClass).Inc()
}
