package main

import (
	"github.com/ccesarfp/shrine/internal/protobuf"
	"github.com/ccesarfp/shrine/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const address string = "localhost:8080"
const network string = "tcp"

var start = time.Now()

func main() {
	log.Println("Starting gRPC shrine on", address)

	log.Println("Starting listener")
	listener, err := net.Listen(network, address)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	protobuf.RegisterTokenServer(s, &service.Server{})

	log.Println("Application initialization took", time.Since(start))

	log.Fatalf("Failed to serve: %v", s.Serve(listener))
}
