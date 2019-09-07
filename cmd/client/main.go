// The Calendar client
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/antik9/microservice-go/internal/events"
	"github.com/antik9/microservice-go/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
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
