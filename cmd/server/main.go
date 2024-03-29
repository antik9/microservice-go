// The Calendar server
package main

import (
	"log"
	"net"

	"github.com/antik9/microservice-go/internal/backends"
	"github.com/antik9/microservice-go/internal/config"
	"github.com/antik9/microservice-go/internal/metrics"
	"github.com/antik9/microservice-go/internal/server"
	"github.com/antik9/microservice-go/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	go metrics.ServeMetrics(config.Conf.Prometheus.Server)

	sock, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	defer sock.Close()

	calendar, err := backend.NewCalendar()
	if err != nil {
		log.Fatalf("cannot instantiate calendar %v", err)
	}
	go backend.ObserveEvents(calendar)

	grpcServer := grpc.NewServer()
	pb.RegisterEventServiceServer(grpcServer, &server.Server{Calendar: calendar})
	grpcServer.Serve(sock)
}
