package main

import (
	"github.com/ccesarfp/shrine/pkg/protobuf"
	"github.com/ccesarfp/shrine/pkg/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const address string = "localhost:8080"
const network string = "tcp"

func main() {
	start := time.Now()

	log.Println("Starting gRPC server on", address)

	log.Println("Starting listener")
	listener, err := net.Listen(network, address)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	protobuf.RegisterTokenServer(s, &service.Server{})

	t := time.Now()
	elapsed := t.Sub(start)
	log.Println("Application initialization took", elapsed)

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
