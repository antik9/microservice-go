// Package events provide API to add, remove and update events
// on particular calendar day
package events

import (
	"errors"
	"time"

	"github.com/antik9/microservice-go/internal/pb"
)

func dayOfEvent(e pb.Event) Day {
	t := time.Unix(e.Beginning.GetSeconds(), 0)
	return Day{day: t.Day(), month: int(t.Month()), year: t.Year()}
}

func indexOfEvent(e pb.Event, es []pb.Event) (int, bool) {
	for idx, event := range es {
		if e.Name == event.Name {
			return idx, true
		}
	}
	return -1, false
}

// Add is a function to add event to a calendar
func Add(e pb.Event, c *Calendar) error {
	day := dayOfEvent(e)

	c.Lock()
	defer c.Unlock()

	if plannedEvents, ok := c.db[day]; ok {
		if _, present := indexOfEvent(e, plannedEvents); present {
			return errors.New("event is already present")
		}
	}
	c.db[day] = append(c.db[day], e)
	return nil
}

// Remove is a function to remove event from a calendar
func Remove(e pb.Event, c *Calendar) error {
	day := dayOfEvent(e)

	c.Lock()
	defer c.Unlock()

	if plannedEvents, ok := c.db[day]; ok {
		idx, present := indexOfEvent(e, plannedEvents)
		if !present {
			return errors.New("event is absent")
		}
		c.db[day] = append(c.db[day][:idx], c.db[day][idx+1:]...)
		return nil
	}
	return errors.New("no events on this date")
}

// Update is a function to update event in a calendar
func Update(e pb.Event, c *Calendar) error {
	day := dayOfEvent(e)

	c.Lock()
	defer c.Unlock()

	if plannedEvents, ok := c.db[day]; ok {
		idx, present := indexOfEvent(e, plannedEvents)
		if !present {
			return errors.New("event is absent")
		}
		c.db[day][idx] = e
		return nil
	}
	return errors.New("no events on this date")
}
