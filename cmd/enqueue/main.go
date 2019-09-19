package main

import (
	"log"
	"time"

	"github.com/antik9/microservice-go/internal/backends"
	"github.com/antik9/microservice-go/internal/queue"
)

func main() {
	calendar, err := backend.NewCalendar()
	if err != nil {
		log.Fatalf("Cannot instantiate calendar %v", err)
	}

	client := queue.NewClient("publisher")
	defer client.Close()

	for {
		for _, event := range calendar.GetImmediateEvents() {
			client.SendMessage(event.Name)
			event.IsSent = true
			calendar.Update(event)
		}
		time.Sleep(time.Second * 1)
	}
}
