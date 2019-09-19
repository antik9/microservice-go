package events

const (
	DayFormat  = "2006/1/2"
	TimeFormat = "2006/1/2 15:04:05"
)

// Calendar is an interface for keeping events
type Calendar interface {
	Add(e Event) error
	GetImmediateEvents() []Event
	Print() string
	Remove(e Event) error
	Update(e Event) error
}
