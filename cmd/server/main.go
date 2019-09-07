// The Calendar server
package main

import (
	"flag"
	"log"
	"net"

	"github.com/antik9/microservice-go/internal/backends"
	"github.com/antik9/microservice-go/internal/server"
	"github.com/antik9/microservice-go/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	database := flag.String("db", "memory", "database can be `postgres` or `memory`")
	flag.Parse()

	sock, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	defer sock.Close()

	calendar, err := backend.NewCalendar(*database)
	if err != nil {
		log.Fatalf("cannot instantiate calendar %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEventServiceServer(grpcServer, &server.Server{Calendar: calendar})
	grpcServer.Serve(sock)
}
