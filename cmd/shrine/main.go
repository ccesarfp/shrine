package main

import (
	"github.com/ccesarfp/shrine/internal/config/application"
)

func main() {
	s := application.New()
	s.SetupServer()
	s.Up()
}
