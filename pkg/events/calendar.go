package events

import (
	"sync"

	"github.com/antik9/microservice-go/internal/pb"
)

// Day is a calendar day
type Day struct {
	day   int
	month int
	year  int
}

// Calendar is a map for day in year to events in this day
type Calendar struct {
	sync.Mutex
	db map[Day][]pb.Event
}

var (
	// MyCalendar is in-memory database of events
	MyCalendar = &Calendar{db: make(map[Day][]pb.Event, 0)}
)
