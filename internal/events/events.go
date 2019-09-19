// Package events provide API to add, remove and update events
// on particular calendar day
package events

import (
	"errors"
	"strings"
	"time"

	"github.com/antik9/microservice-go/pkg/pb"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Event struct {
	Day       time.Time `db:"day"`
	Beginning time.Time `db:"beginning"`
	End       time.Time `db:"endofevent"`
	Name      string    `db:"name"`
	EventType int       `db:"eventtype"`
	IsSent    bool      `db:"is_sent"`
}

// EventFromProto converts proto structure to Event struct
func EventFromProto(e *pb.Event) Event {
	t := time.Unix(e.Beginning.GetSeconds(), 0)
	d, _ := time.Parse(DayFormat, t.Format(DayFormat))

	return Event{
		Day:       d,
		Beginning: time.Unix(e.Beginning.GetSeconds(), 0),
		End:       time.Unix(e.End.GetSeconds(), 0),
		Name:      e.Name,
		EventType: int(e.EventType)}
}

// ParseFromInput converts input string to Event or return an occuring error
func ParseFromInput(msg string) (*Event, error) {
	parts := strings.Split(msg, " - ")
	if len(parts) == 4 {
		location, _ := time.LoadLocation("Europe/Moscow")
		beggining, err1 := time.ParseInLocation(TimeFormat, parts[0], location)
		end, err2 := time.ParseInLocation(TimeFormat, parts[1], location)
		eventType := 2

		switch strings.Trim(parts[3], "\n\r\t ") {
		case "MEETING":
			eventType = 0
		case "REMINDER":
			eventType = 1
		default:
			eventType = 2
		}

		if err1 == nil && err2 == nil {
			return NewEvent(beggining.Local(), end.Local(), parts[2], eventType), nil
		}
	}
	return nil, errors.New("bad format")
}

// Proto method convert Event back to protobuf structure
func (e Event) Proto() *pb.Event {
	return &pb.Event{
		Beginning: &timestamp.Timestamp{Seconds: e.Beginning.Unix(), Nanos: 0},
		End:       &timestamp.Timestamp{Seconds: e.End.Unix(), Nanos: 0},
		Name:      e.Name,
		EventType: pb.Event_Type(e.EventType)}
}

// NewEvent creates Event instance
func NewEvent(beginning, end time.Time, name string, eventType int) *Event {
	t := time.Unix(beginning.Unix(), 0)
	d, _ := time.Parse(DayFormat, t.Format(DayFormat))

	return &Event{
		Day:       d,
		Beginning: beginning,
		End:       end,
		Name:      name,
		EventType: eventType}
}

func (e Event) IndexOfEvent(es []Event) (int, bool) {
	for idx, event := range es {
		if e.Name == event.Name {
			return idx, true
		}
	}
	return -1, false
}
