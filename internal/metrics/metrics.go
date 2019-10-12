package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	AddEvent       = "add_event"
	NumberOfEvents = "number_of_events"
	PrintAll       = "print_all"
	RemoveEvent    = "remove_event"
	SendEvent      = "sent_event"
	UpdateEvent    = "udpate_event"
)

var (
	addEventSummary = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name:       AddEvent,
			Help:       "Add new event to the database",
			Objectives: map[float64]float64{0.5: 0.000000001, 0.9: 0.000000001, 0.99: 0.000000001},
		})
	removeEventSummary = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name:       RemoveEvent,
			Help:       "Remove event from the database",
			Objectives: map[float64]float64{0.5: 0.000000001, 0.9: 0.000000001, 0.99: 0.000000001},
		})
	updateEventSummary = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name:       UpdateEvent,
			Help:       "Update event in the database",
			Objectives: map[float64]float64{0.5: 0.000000001, 0.9: 0.000000001, 0.99: 0.000000001},
		})
	printAllSummary = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name:       PrintAll,
			Help:       "Print all events from database",
			Objectives: map[float64]float64{0.5: 0.000000001, 0.9: 0.000000001, 0.99: 0.000000001},
		})
	sentEventsCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: SendEvent,
			Help: "Sent events to users",
		})
	numberOfEvents = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: NumberOfEvents,
			Help: "Sent events to users",
		})
)

func Observe(name string, value float64) {
	switch name {
	case AddEvent:
		addEventSummary.Observe(value)
	case RemoveEvent:
		removeEventSummary.Observe(value)
	case UpdateEvent:
		updateEventSummary.Observe(value)
	case PrintAll:
		printAllSummary.Observe(value)
	case SendEvent:
		sentEventsCounter.Inc()
	case NumberOfEvents:
		numberOfEvents.Set(value)
	default:
	}
}

func ServeMetrics(address string) {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(address, nil)
}
