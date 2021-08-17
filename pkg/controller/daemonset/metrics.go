package daemonset

import "github.com/prometheus/client_golang/prometheus"

var (
	dsCreateEvent = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "kruise_daemon_create_pod_events",
	}, []string{"status"})

	dsDeleteEvent = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "kruise_daemon_delete_pod_events",
	}, []string{"status"})

	dsUpdateEvent = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "kruise_daemon_update_pod_event_type",
	}, []string{"type"})
)

func init() {
	prometheus.MustRegister(
		dsCreateEvent,
		dsDeleteEvent,
		dsUpdateEvent,
	)
}
