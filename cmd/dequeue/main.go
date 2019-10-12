package main

import (
	"fmt"

	"github.com/antik9/microservice-go/internal/config"
	"github.com/antik9/microservice-go/internal/metrics"
	"github.com/antik9/microservice-go/internal/queue"
)

func main() {
	go metrics.ServeMetrics(config.Conf.Prometheus.Dequeue)

	client := queue.NewClient("consumer")
	defer client.Close()

	for {
		fmt.Println(client.ReadMessage())
	}
}
