package backend

import (
	"github.com/antik9/microservice-go/internal/backends/database"
	"github.com/antik9/microservice-go/internal/backends/memory"
	"github.com/antik9/microservice-go/internal/config"
	"github.com/antik9/microservice-go/internal/events"
)

func NewCalendar() (events.Calendar, error) {
	if config.Conf.Database.Backend == "postgres" {
		return db.NewCalendar()
	}
	return memory.NewCalendar()
}
