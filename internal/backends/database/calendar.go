package db

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/antik9/microservice-go/internal/config"
	"github.com/antik9/microservice-go/internal/events"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DatabaseCalendar is a struct with connection to database
type DatabaseCalendar struct {
	conn *sqlx.DB
}

// NewCalendar in initialization of DatabaseCalendar struct
func NewCalendar() (*DatabaseCalendar, error) {
	connectionParams := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Conf.Database.Host,
		config.Conf.Database.Username,
		config.Conf.Database.Password,
		config.Conf.Database.Name)

	db, err := sqlx.Connect("postgres", connectionParams)
	if err != nil {
		log.Fatalf("cannot connect to database, %v", err)
	}
	return &DatabaseCalendar{conn: db}, nil
}

// Add is a function to add event to a calendar
func (c *DatabaseCalendar) Add(e events.Event) error {
	plannedEvents := getEventsAtDay(c.conn, e.Day)
	if _, present := e.IndexOfEvent(plannedEvents); present {
		return errors.New("event is already present")
	}
	c.conn.MustExec(
		`INSERT INTO events (day, beginning, endofevent, name, eventtype, is_sent)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		e.Day.Format(events.DayFormat), e.Beginning.Format(events.TimeFormat),
		e.End.Format(events.TimeFormat), e.Name, e.EventType, e.IsSent)
	return nil
}

func getEventsAtDay(db *sqlx.DB, d time.Time) []events.Event {
	plannedEvents := []events.Event{}
	err := db.Select(
		&plannedEvents, "SELECT * FROM events WHERE day = $1", d.Format(events.DayFormat))
	if err != nil {
		log.Fatalf("cannot execute query, %v", err)
	}
	return plannedEvents
}

func (c *DatabaseCalendar) GetImmediateEvents() []events.Event {
	plannedEvents := []events.Event{}
	err := c.conn.Select(
		&plannedEvents,
		"SELECT * FROM events WHERE beginning <= $1 AND is_sent IS FALSE",
		time.Now().Format(events.TimeFormat))
	if err != nil {
		log.Fatalf("cannot execute query, %v", err)
	}
	return plannedEvents
}

func (c *DatabaseCalendar) Print() string {
	var builder strings.Builder

	plannedEvents := []events.Event{}
	err := c.conn.Select(&plannedEvents, "SELECT * FROM events")

	if err != nil {
		log.Fatalf("cannot execute query, %v", err)
	}

	for _, event := range plannedEvents {
		builder.WriteString("\tfrom ")
		builder.WriteString(event.Beginning.Format(time.UnixDate))
		builder.WriteString(" to ")
		builder.WriteString(event.End.Format(time.UnixDate))
		builder.WriteString(": " + event.Name + "\n")
	}
	return builder.String()
}

// Remove is a function to remove event from a calendar
func (c *DatabaseCalendar) Remove(e events.Event) error {
	plannedEvents := getEventsAtDay(c.conn, e.Day)
	if _, present := e.IndexOfEvent(plannedEvents); !present {
		return errors.New("event is absent")
	}
	c.conn.MustExec(
		"DELETE FROM events WHERE name = $1 AND day = $2",
		e.Name, e.Day.Format(events.DayFormat))
	return nil
}

// Update is a function to update event in a calendar
func (c *DatabaseCalendar) Update(e events.Event) error {
	plannedEvents := getEventsAtDay(c.conn, e.Day)
	if _, present := e.IndexOfEvent(plannedEvents); !present {
		return errors.New("event is absent")
	}
	c.conn.MustExec(
		`UPDATE events SET beginning = $1, endofevent = $2, eventtype = $3, is_sent = $4
			WHERE name = $5 AND day = $6`,
		e.Beginning.Format(events.TimeFormat),
		e.End.Format(events.TimeFormat), e.EventType, e.IsSent,
		e.Name, e.Day.Format(events.DayFormat))
	return nil
}
