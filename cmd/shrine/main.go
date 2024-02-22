package main

import (
	"github.com/ccesarfp/shrine/internal/protobuf"
	"github.com/ccesarfp/shrine/internal/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const address string = "0.0.0.0:3000"
const network string = "tcp"

var start = time.Now()

func main() {
	log.Println("Getting environment variables")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Starting gRPC shrine on", address)

	log.Println("Starting listener")
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	protobuf.RegisterTokenServer(s, &service.Server{})

	log.Println("Application initialization took", time.Since(start))

	log.Fatalf("Failed to serve: %v", s.Serve(listener))
}
