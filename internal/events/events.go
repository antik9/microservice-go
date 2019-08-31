// Package events provide API to add, remove and update events
// on particular calendar day
package events

import (
	"time"

	"github.com/antik9/microservice-go/pkg/pb"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Event struct {
	day       Day
	beginning time.Time
	end       time.Time
	name      string
	eventType int
}

// EventFromProto converts proto structure to Event struct
func EventFromProto(e *pb.Event) Event {
	t := time.Unix(e.Beginning.GetSeconds(), 0)
	d := Day{day: t.Day(), month: int(t.Month()), year: t.Year()}

	return Event{
		day:       d,
		beginning: time.Unix(e.Beginning.GetSeconds(), 0),
		end:       time.Unix(e.End.GetSeconds(), 0),
		name:      e.Name,
		eventType: int(e.EventType)}
}

// Proto method convert Event back to protobuf structure
func (e Event) Proto() *pb.Event {
	return &pb.Event{
		Beginning: &timestamp.Timestamp{Seconds: e.beginning.Unix(), Nanos: 0},
		End:       &timestamp.Timestamp{Seconds: e.end.Unix(), Nanos: 0},
		Name:      e.name,
		EventType: pb.Event_Type(e.eventType)}
}

// NewEvent creates Event instance
func NewEvent(beginning, end time.Time, name string, eventType int) Event {
	t := time.Unix(beginning.Unix(), 0)
	d := Day{day: t.Day(), month: int(t.Month()), year: t.Year()}

	return Event{
		day:       d,
		beginning: beginning,
		end:       end,
		name:      name,
		eventType: eventType}
}

func (e Event) indexOfEvent(es []Event) (int, bool) {
	for idx, event := range es {
		if e.name == event.name {
			return idx, true
		}
	}
	return -1, false
}
