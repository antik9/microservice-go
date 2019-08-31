// The Calendar server
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/antik9/microservice-go/internal/events"
	"github.com/antik9/microservice-go/pkg/pb"
	"google.golang.org/grpc"
)

// Server is implementation of EventService
type Server struct {
	calendar events.Calendar
}

// AddEvent adds event to inner calendar
func (s *Server) AddEvent(c context.Context, e *pb.Event) (*pb.Empty, error) {
	return &pb.Empty{}, s.calendar.Add(events.EventFromProto(e))
}

// RemoveEvent removes event to inner calendar
func (s *Server) RemoveEvent(c context.Context, e *pb.Event) (*pb.Empty, error) {
	return &pb.Empty{}, s.calendar.Remove(events.EventFromProto(e))
}

// UpdateEvent updates event to inner calendar
func (s *Server) UpdateEvent(c context.Context, e *pb.Event) (*pb.Empty, error) {
	return &pb.Empty{}, s.calendar.Update(events.EventFromProto(e))
}

// PrintAll prints days and theirs events from inner calendar
func (s *Server) PrintAll(c context.Context, e *pb.Empty) (*pb.Response, error) {
	return &pb.Response{Resp: s.calendar.Print()}, nil
}

func main() {
	mode := flag.String("mode", "server", "mode can be `server` or `client`")
	flag.Parse()

	if *mode == "server" {
		sock, err := net.Listen("tcp", "0.0.0.0:50051")
		if err != nil {
			log.Fatalf("failed to listen %v", err)
		}
		defer sock.Close()

		grpcServer := grpc.NewServer()
		pb.RegisterEventServiceServer(grpcServer, &Server{events.NewCalendar()})
		grpcServer.Serve(sock)
	} else {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("could not connect: %v", err)
		}
		defer conn.Close()

		var r *pb.Response
		c := pb.NewEventServiceClient(conn)
		e1 := events.NewEvent(time.Now(), time.Now(), "Some event", 1)
		e2 := events.NewEvent(time.Now(), time.Now(), "Some other event", 2)

		c.AddEvent(context.Background(), e1.Proto())
		r, _ = c.PrintAll(context.Background(), &pb.Empty{})
		fmt.Println(r.Resp)

		c.AddEvent(context.Background(), e2.Proto())
		r, _ = c.PrintAll(context.Background(), &pb.Empty{})
		fmt.Println(r.Resp)

		c.RemoveEvent(context.Background(), e1.Proto())
		r, _ = c.PrintAll(context.Background(), &pb.Empty{})
		fmt.Println(r.Resp)
	}
}
