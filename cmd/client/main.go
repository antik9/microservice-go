// The Calendar client
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/antik9/microservice-go/internal/events"
	"github.com/antik9/microservice-go/pkg/pb"
	"google.golang.org/grpc"
)

func readString(reader *bufio.Reader) string {
	msg, _ := reader.ReadString('\n')
	return strings.TrimRight(msg, "\n\t\r ")
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	var r *pb.Response
	var action string
	c := pb.NewEventServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(`Choose command:
	1 - print all events
	2 - add event
	3 - remove event`)
		action = readString(reader)

		switch action {
		case "1":
			r, _ = c.PrintAll(context.Background(), &pb.Empty{})
			fmt.Println(r.Resp)
		case "2":
			fmt.Println(`Print new task with such format:
2006/1/2 12:45:10 - 2006/1/2 13:45:10 - Walk the dog - REMINDER`)
			action = readString(reader)
			event, err := events.ParseFromInput(action)
			if err != nil {
				fmt.Println(err)
			} else {
				c.AddEvent(context.Background(), event.Proto())
			}
		case "3":
			fmt.Println(`Print task to remove with such format:
2006/1/2 12:45:10 - 2006/1/2 13:45:10 - Walk the dog - REMINDER`)
			action = readString(reader)
			event, err := events.ParseFromInput(action)
			if err != nil {
				fmt.Println(err)
			} else {
				c.RemoveEvent(context.Background(), event.Proto())
			}
		}
	}
}
