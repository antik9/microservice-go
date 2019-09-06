package server

import (
	"context"

	"github.com/antik9/microservice-go/internal/events"
	"github.com/antik9/microservice-go/pkg/pb"
)

// Server is implementation of EventService
type Server struct {
	Calendar events.Calendar
}

// AddEvent adds event to inner Calendar
func (s *Server) AddEvent(c context.Context, e *pb.Event) (*pb.Empty, error) {
	return &pb.Empty{}, s.Calendar.Add(events.EventFromProto(e))
}

// RemoveEvent removes event to inner Calendar
func (s *Server) RemoveEvent(c context.Context, e *pb.Event) (*pb.Empty, error) {
	return &pb.Empty{}, s.Calendar.Remove(events.EventFromProto(e))
}

// UpdateEvent updates event to inner Calendar
func (s *Server) UpdateEvent(c context.Context, e *pb.Event) (*pb.Empty, error) {
	return &pb.Empty{}, s.Calendar.Update(events.EventFromProto(e))
}

// PrintAll prints days and theirs events from inner Calendar
func (s *Server) PrintAll(c context.Context, e *pb.Empty) (*pb.Response, error) {
	return &pb.Response{Resp: s.Calendar.Print()}, nil
}
