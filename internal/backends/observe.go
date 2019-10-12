package backend

import (
	"time"

	"github.com/antik9/microservice-go/internal/events"
	"github.com/antik9/microservice-go/internal/metrics"
)

func ObserveEvents(calendar events.Calendar) {
	for {
		metrics.Observe(metrics.NumberOfEvents, float64(calendar.Count()))
		time.Sleep(time.Minute)
	}
}
