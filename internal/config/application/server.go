package application

import (
	"github.com/ccesarfp/shrine/internal/protobuf"
	"github.com/ccesarfp/shrine/internal/service"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

type server struct {
	Environment string
	Address     string
	Network     string
	StartTime   time.Time
	Server      *grpc.Server
}

func newServer() *server {
	return &server{}
}

// SetupServer - start application preparation
func (s *server) SetupServer() {
	i.S.StartTime = time.Now()

	i.S.setupEnvironmentVars()

	log.Println("Starting gRPC", i.Name, "v"+i.Version, "("+i.S.Environment+") on", i.S.Address)
	i.S.Server = grpc.NewServer(
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			grpcrecovery.UnaryServerInterceptor(i.cb.errorHandler()),
		)),
		grpc.KeepaliveParams(kasp),
	)
	protobuf.RegisterTokenServer(i.S.Server, &service.Server{})
}

// setupEnvironmentVars - checks if environment variables exist, otherwise loads variables from .env
func (s *server) setupEnvironmentVars() {
	log.Println("Getting environment variables")
	hasEnvironmentVars := os.Getenv("HAS_ENV_VARS")
	if hasEnvironmentVars == "" {
		err := godotenv.Load()
		if err != nil {
			log.Panicln("Error loading .env file")
		}
	}
	i.Name = os.Getenv("APP_NAME")
	i.Version = os.Getenv("APP_VERSION")
	i.S.Environment = os.Getenv("ENV")
	i.S.Network = os.Getenv("NETWORK")
	i.S.Address = os.Getenv("ADDRESS") + ":" + os.Getenv("PORT")
}
