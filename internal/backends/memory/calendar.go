package memory

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/antik9/microservice-go/internal/events"
)

// InMemoryCalendar is a map for day in year to events in this day
type InMemoryCalendar struct {
	sync.Mutex
	db map[time.Time][]events.Event
}

// NewCalendar in initialization of InMemoryCalendar struct
func NewCalendar() (*InMemoryCalendar, error) {
	return &InMemoryCalendar{
		db: make(map[time.Time][]events.Event, 0)}, nil
}

// Add is a function to add event to a calendar
func (c *InMemoryCalendar) Add(e events.Event) error {
	c.Lock()
	defer c.Unlock()

	if plannedEvents, ok := c.db[e.Day]; ok {
		if _, present := e.IndexOfEvent(plannedEvents); present {
			return errors.New("event is already present")
		}
	}
	c.db[e.Day] = append(c.db[e.Day], e)
	return nil
}

func (c *InMemoryCalendar) Clear() {
	c.db = map[time.Time][]events.Event{}
}

func (c *InMemoryCalendar) GetAllEvents() []events.Event {
	allEvents := []events.Event{}
	for _, eventsOnDay := range c.db {
		for _, event := range eventsOnDay {
			allEvents = append(allEvents, event)
		}
	}
	return allEvents
}

func (c *InMemoryCalendar) GetImmediateEvents() []events.Event {
	now := time.Now()
	d, _ := time.Parse(events.DayFormat, now.Format(events.DayFormat))
	plannedEvents := []events.Event{}
	for _, event := range c.db[d] {
		if event.Beginning.Unix() < now.Unix() && !event.IsSent {
			plannedEvents = append(plannedEvents, event)
		}
	}
	return plannedEvents
}

func (c *InMemoryCalendar) Print() string {
	c.Lock()
	defer c.Unlock()

	var builder strings.Builder
	for day, events := range c.db {
		builder.WriteString(day.Format("2006/01/02") + "\n")
		for _, event := range events {
			builder.WriteString("\tfrom ")
			builder.WriteString(event.Beginning.Format(time.UnixDate))
			builder.WriteString(" to ")
			builder.WriteString(event.End.Format(time.UnixDate))
			builder.WriteString(": " + event.Name + "\n")
		}
	}
	return builder.String()
}

// Remove is a function to remove event from a calendar
func (c *InMemoryCalendar) Remove(e events.Event) error {
	c.Lock()
	defer c.Unlock()

	if plannedEvents, ok := c.db[e.Day]; ok {
		idx, present := e.IndexOfEvent(plannedEvents)
		if !present {
			return errors.New("event is absent")
		}
		c.db[e.Day] = append(c.db[e.Day][:idx], c.db[e.Day][idx+1:]...)
		return nil
	}
	return errors.New("no events on this date")
}

// Update is a function to update event in a calendar
func (c *InMemoryCalendar) Update(e events.Event) error {
	c.Lock()
	defer c.Unlock()

	if plannedEvents, ok := c.db[e.Day]; ok {
		idx, present := e.IndexOfEvent(plannedEvents)
		if !present {
			return errors.New("event is absent")
		}
		c.db[e.Day][idx] = e
	}
	return nil
}
