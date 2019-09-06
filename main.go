// The Calendar server
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/antik9/microservice-go/internal/backends/database"
	"github.com/antik9/microservice-go/internal/backends/memory"
	"github.com/antik9/microservice-go/internal/events"
	"github.com/antik9/microservice-go/internal/server"
	"github.com/antik9/microservice-go/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	mode := flag.String("mode", "server", "mode can be `server` or `client`")
	database := flag.String("db", "memory", "database can be `postgres` or `memory`")
	flag.Parse()

	if *mode == "server" {
		sock, err := net.Listen("tcp", "0.0.0.0:50051")
		if err != nil {
			log.Fatalf("failed to listen %v", err)
		}
		defer sock.Close()

		var calendar events.Calendar
		if *database == "memory" {
			calendar, err = memory.NewCalendar()
		} else {
			calendar, err = db.NewCalendar()
		}

		if err != nil {
			log.Fatalf("cannot instantiate calendar %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterEventServiceServer(grpcServer, &server.Server{Calendar: calendar})
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
