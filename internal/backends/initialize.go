package backend

import (
	"github.com/antik9/microservice-go/internal/backends/database"
	"github.com/antik9/microservice-go/internal/backends/memory"
	"github.com/antik9/microservice-go/internal/events"
)

func NewCalendar(dbType string) (events.Calendar, error) {
	if dbType == "postgres" {
		return db.NewCalendar()
	}
	return memory.NewCalendar()
}
