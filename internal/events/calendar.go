package events

import (
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"
)

// day is a calendar day
type Day struct {
	day   int
	month int
	year  int
}

// Calendar is a map for day in year to events in this day
type Calendar struct {
	sync.Mutex
	db map[Day][]Event
}

// NewCalendar in initialization of Calendat struct
func NewCalendar() Calendar {
	return Calendar{
		db: make(map[Day][]Event, 0)}
}

// Add is a function to add event to a calendar
func (c *Calendar) Add(e Event) error {
	c.Lock()
	defer c.Unlock()

	if plannedEvents, ok := c.db[e.day]; ok {
		if _, present := e.indexOfEvent(plannedEvents); present {
			return errors.New("event is already present")
		}
	}
	c.db[e.day] = append(c.db[e.day], e)
	return nil
}

func (c *Calendar) Print() string {
	c.Lock()
	defer c.Unlock()

	var builder strings.Builder
	for day, events := range c.db {
		builder.WriteString(strconv.Itoa(day.day) + "-" +
			strconv.Itoa(day.month) + "-" +
			strconv.Itoa(day.year) + "\n")
		for _, event := range events {
			builder.WriteString("\tfrom ")
			builder.WriteString(event.beginning.Format(time.UnixDate))
			builder.WriteString(" to ")
			builder.WriteString(event.end.Format(time.UnixDate))
			builder.WriteString(": " + event.name + "\n")
		}
	}
	return builder.String()
}

// Remove is a function to remove event from a calendar
func (c *Calendar) Remove(e Event) error {
	c.Lock()
	defer c.Unlock()

	if plannedEvents, ok := c.db[e.day]; ok {
		idx, present := e.indexOfEvent(plannedEvents)
		if !present {
			return errors.New("event is absent")
		}
		c.db[e.day] = append(c.db[e.day][:idx], c.db[e.day][idx+1:]...)
		return nil
	}
	return errors.New("no events on this date")
}

// Update is a function to update event in a calendar
func (c *Calendar) Update(e Event) error {
	c.Lock()
	defer c.Unlock()

	if plannedEvents, ok := c.db[e.day]; ok {
		idx, present := e.indexOfEvent(plannedEvents)
		if !present {
			return errors.New("event is absent")
		}
		c.db[e.day][idx] = e
	}
	return nil
}
