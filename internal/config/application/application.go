package application

import (
	"github.com/ccesarfp/shrine/internal/protobuf"
	"github.com/ccesarfp/shrine/internal/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type Application struct {
	Name        string
	Version     string
	Environment string
	Address     string
	Network     string
	StartTime   time.Time
	Server      *grpc.Server
}

var (
	appOnce  sync.Once
	instance *Application
	kasp     = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	}
)

func New() *Application {
	appOnce.Do(func() {
		instance = &Application{}
	})
	return instance
}

// SetupServer - start application preparation
func (a *Application) SetupServer() {
	instance.StartTime = time.Now()

	instance.setupEnvironmentVars()

	log.Println("Starting gRPC", instance.Name, "v"+instance.Version, "("+instance.Version+") on", instance.Address)
	instance.Server = grpc.NewServer(grpc.KeepaliveParams(kasp))
	protobuf.RegisterTokenServer(instance.Server, &service.Server{})
}

// setupEnvironmentVars - checks if environment variables exist, otherwise loads variables from .env
func (a *Application) setupEnvironmentVars() {
	log.Println("Getting environment variables")
	hasEnvironmentVars := os.Getenv("HAS_ENV_VARS")
	if hasEnvironmentVars == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	instance.Name = os.Getenv("APP_NAME")
	instance.Version = os.Getenv("APP_VERSION")
	instance.Environment = os.Getenv("ENV")
	instance.Network = os.Getenv("NETWORK")
	instance.Address = os.Getenv("ADDRESS") + ":" + os.Getenv("PORT")
}

// Up - Start application
func (a *Application) Up() {
	log.Println("Starting listener")
	listener, err := net.Listen(instance.Network, instance.Address)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Application initialization took", time.Since(instance.StartTime))
	err = instance.Server.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}

// Down - shut down application
func (a *Application) Down() {
	log.Println("Shutting down application")
	instance.Server.GracefulStop()
}

// DownBrutally - forcefully shutdown application
func (a *Application) DownBrutally() {
	log.Println("Brutally shutting down application")
	instance.Server.Stop()
}
