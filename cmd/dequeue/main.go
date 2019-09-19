package main

import (
	"fmt"

	"github.com/antik9/microservice-go/internal/queue"
)

func main() {
	client := queue.NewClient("consumer")
	defer client.Close()

	for {
		fmt.Println(client.ReadMessage())
	}
}
