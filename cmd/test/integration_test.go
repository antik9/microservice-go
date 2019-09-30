package main

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/antik9/microservice-go/internal/backends"
	"github.com/antik9/microservice-go/internal/events"
	"github.com/antik9/microservice-go/pkg/pb"
	"google.golang.org/grpc"
)

func TestIntegration(t *testing.T) {
	time.Sleep(5 * time.Second)

	calendar, err := backend.NewCalendar()
	if err != nil {
		t.Fatal(err)
	}

	calendar.Clear()

	conn, err := grpc.Dial("server:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewEventServiceClient(conn)

	r, _ := c.PrintAll(context.Background(), &pb.Empty{})

	if r.Resp != "" {
		t.Fatalf("expected: empty string; got: %s", r.Resp)
	}

	loc, _ := time.LoadLocation("Europe/Moscow")
	onTime := time.Now().In(loc).Add(time.Second * 5).Format(events.TimeFormat)
	s := onTime + " - " + onTime + " - Talk on the phone - REMINDER"
	event, _ := events.ParseFromInput(s)

	c.AddEvent(context.Background(), event.Proto())

	r, _ = c.PrintAll(context.Background(), &pb.Empty{})

	if !strings.Contains(r.Resp, "Talk on the phone") {
		t.Fatalf("expected: don't receive expected task")
	}

	evs := calendar.GetImmediateEvents()
	if len(evs) != 0 {
		t.Fatalf("got unexpected events before they should come")
	}

	time.Sleep(10 * time.Second)
	evs = calendar.GetAllEvents()
	if len(evs) != 1 {
		t.Fatalf("expected 1 event; got %d events", len(evs))
	}

	if !evs[0].IsSent {
		t.Fatalf("event unexpectedly is not sent")
	}
}
