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
	i.s.StartTime = time.Now()

	i.s.setupEnvironmentVars()

	log.Println("[Server] Setting interceptors")
	interceptorChain := grpcmiddleware.ChainUnaryServer(
		i.cb.circuitBreakerInterceptor,
		grpcrecovery.UnaryServerInterceptor(i.cb.errorHandler()),
	)

	log.Println("[Server] Starting gRPC", i.Name, "v"+i.Version, "("+i.s.Environment+") on", i.s.Address)
	i.s.Server = grpc.NewServer(
		grpc.UnaryInterceptor(interceptorChain),
		grpc.KeepaliveParams(kasp),
	)
	protobuf.RegisterTokenServer(i.s.Server, &service.Server{})
}

// setupEnvironmentVars - checks if environment variables exist, otherwise loads variables from .env
func (s *server) setupEnvironmentVars() {
	log.Println("[Server] Getting environment variables")
	hasEnvironmentVars := os.Getenv("HAS_ENV_VARS")
	if hasEnvironmentVars == "" {
		err := godotenv.Load()
		if err != nil {
			log.Panicln("Error loading .env file")
		}
	}
	i.Name = os.Getenv("APP_NAME")
	i.Version = os.Getenv("APP_VERSION")
	i.s.Environment = os.Getenv("ENV")
	i.s.Network = os.Getenv("NETWORK")
	i.s.Address = os.Getenv("ADDRESS") + ":" + os.Getenv("PORT")
}
